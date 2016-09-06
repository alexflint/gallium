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
		Template         string
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

	if args.Template == "" {
		fmt.Println("You must provide --template")
		os.Exit(1)
	}

	if args.BundleIdentifier == "" {
		args.BundleIdentifier = bundleName
	}

	// Copy the bundle
	err := copyTree(args.Output, args.Template)
	if err != nil {
		fmt.Println("error in copytree:", err)
		os.Exit(1)
	}

	// Overwrite the executable
	exeDst := filepath.Join(args.Output, "Contents", "MacOS", bundleName)
	fmt.Println("copy executable to", exeDst)

	err = copyFile(exeDst, args.Executable)
	if err != nil {
		fmt.Println("error in copyfile:", err)
		os.Exit(1)
	}

	// Overwrite the info.plist
	tplbuf := MustAsset("info.plist.tpl")
	tpl, err := template.New("info.plist.tpl").Parse(string(tplbuf))
	if err != nil {
		fmt.Println("error parsing info.plist.tpl:", err)
		os.Exit(1)
	}

	plistDst := filepath.Join(args.Output, "Contents", "Info.plist")
	fmt.Println("writing to", plistDst)
	w, err := os.Create(plistDst)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer w.Close()

	tpl.Execute(w, map[string]string{
		"BundleName":       bundleName,
		"BundleIdentifier": args.BundleIdentifier,
	})
}
