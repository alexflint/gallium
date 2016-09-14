//go:generate go-bindata -o bindata.go info.plist.tpl

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	arg "github.com/alexflint/go-arg"
)

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func copyFile(dst, src string) error {
	st, err := os.Stat(src)
	if err != nil {
		return err
	}
	buf, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dst, buf, st.Mode())
}

func copyTree(dst, src string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		// re-stat the path so that we can tell whether it is a symlink
		info, err = os.Lstat(path)
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		targ := filepath.Join(dst, rel)

		switch {
		case info.IsDir():
			return os.Mkdir(targ, 0777)
		case info.Mode()&os.ModeSymlink != 0:
			referent, err := os.Readlink(path)
			if err != nil {
				return err
			}
			return os.Symlink(referent, targ)
		default:
			return copyFile(targ, path)
		}
	})
}

func main() {
	var args struct {
		Executable       string `arg:"positional,required"`
		Output           string `arg:"-o"`
		BundleIdentifier string
	}
	arg.MustParse(&args)

	var bundleName string
	if args.Output == "" {
		bundleName = filepath.Base(args.Executable)
		args.Output = bundleName + ".app"
	} else if !strings.HasSuffix(args.Output, ".app") {
		fmt.Println("output must end with .app")
		os.Exit(1)
	} else {
		bundleName = strings.TrimSuffix(filepath.Base(args.Output), ".app")
	}

	if args.BundleIdentifier == "" {
		args.BundleIdentifier = bundleName
	}

	// Find gallium.framework
	fwSrc := os.ExpandEnv("$GOPATH/src/github.com/alexflint/gallium/lib/build/Debug/Gallium.framework")
	st, err := os.Stat(fwSrc)
	if err != nil {
		fmt.Printf("framework not found at %s\n", fwSrc)
		os.Exit(1)
	}
	if !st.IsDir() {
		fmt.Printf("%s is not a directory\n", fwSrc)
		os.Exit(1)
	}

	// Create the bundle in a temporary dir
	tmpBundle, err := ioutil.TempDir("", "")
	must(err)

	// Create the bundle.app dir
	must(os.MkdirAll(tmpBundle, 0777))

	// Copy the framework in
	fwDst := filepath.Join(tmpBundle, "Contents", "Frameworks", "Gallium.framework")
	must(os.MkdirAll(filepath.Dir(fwDst), 0777))
	must(copyTree(fwDst, fwSrc))

	// Copy the executable in
	exeDst := filepath.Join(tmpBundle, "Contents", "MacOS", bundleName)
	must(os.MkdirAll(filepath.Dir(exeDst), 0777))
	must(copyFile(exeDst, args.Executable))

	// Overwrite the info.plist
	tpl, err := template.New("info.plist.tpl").Parse(string(MustAsset("info.plist.tpl")))
	must(err)

	plistDst := filepath.Join(tmpBundle, "Contents", "Info.plist")
	w, err := os.Create(plistDst)
	must(err)

	tpl.Execute(w, map[string]string{
		"BundleName":       bundleName,
		"BundleIdentifier": args.BundleIdentifier,
	})
	must(w.Close())

	// Delete the bundle.app dir if it already exists
	must(os.RemoveAll(args.Output))

	// Move the temporary dir to the bundle.app location
	must(os.Rename(tmpBundle, args.Output))
}
