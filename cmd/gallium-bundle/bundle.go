//go:generate go-bindata -o bindata.go info.plist.tpl

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	arg "github.com/alexflint/go-arg"
)

func must(err error, info ...interface{}) {
	if err != nil {
		fmt.Println(append(info, err.Error())...)
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
		Executable string `arg:"positional,required"`
		Output     string `arg:"-o"`
		Identifier string `arg:"help:The bundle identifier (CFBundleIdentifier)"`
		Name       string `arg:"help:The bundle name (CFBundleName)"`
		Icon       string `arg:"help:Path to a .icns file or a .iconset dir"`
	}
	arg.MustParse(&args)

	// If output is empty then use the app name if there is one, or the executable otherwise
	if args.Output == "" {
		if args.Name == "" {
			args.Output = filepath.Base(args.Executable) + ".app"
		} else {
			args.Output = args.Name + ".app"
		}
	}

	if !strings.HasSuffix(args.Output, ".app") {
		fmt.Println("output must end with .app")
		os.Exit(1)
	}

	// If the bundle name is empty then use the app name
	if args.Name == "" {
		args.Name = strings.TrimSuffix(filepath.Base(args.Output), ".app")
	}

	// If the bundle identifier is empty then use the bundle name
	if args.Identifier == "" {
		args.Identifier = args.Name
	}

	// extras for the Info.plist
	extraProps := make(map[string]string)

	// get the path to the gallium package
	golistCmd := exec.Command("go", "list", "-f", "{{.Dir}}", "github.com/alexflint/gallium")
	golistOut, err := golistCmd.CombinedOutput()
	if err != nil {
		fmt.Printf("go list github.com/alexflint/gallium failed:\n%s\n", string(golistOut))
		os.Exit(1)
	}

	// Find Gallium.framework
	galliumDir := strings.TrimSpace(string(golistOut))
	fwSrc := filepath.Join(galliumDir, "dist", "Gallium.framework")
	st, err := os.Stat(fwSrc)
	if err != nil {
		fmt.Printf("framework not found at %s: %v\n", fwSrc, err)
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
	exeDst := filepath.Join(tmpBundle, "Contents", "MacOS", args.Name)
	must(os.MkdirAll(filepath.Dir(exeDst), 0777))
	must(copyFile(exeDst, args.Executable))

	// Copy the icon in
	if args.Icon != "" {
		st, err := os.Stat(args.Icon)
		must(err)

		iconExt := filepath.Ext(args.Icon)
		iconName := strings.TrimSuffix(filepath.Base(args.Icon), iconExt) + ".icns"
		iconDst := filepath.Join(tmpBundle, "Contents", "Resources", iconName)
		must(os.MkdirAll(filepath.Dir(iconDst), 0777))
		extraProps["CFBundleIconFile"] = iconName

		// There are three kinds of source icons
		switch {
		case iconExt == ".icns":
			if !st.Mode().IsRegular() {
				fmt.Println("Icon had extension .icns but was not a regular file")
				os.Exit(1)
			}
			must(copyFile(iconDst, args.Icon))
		case iconExt == ".iconset":
			if !st.IsDir() {
				fmt.Println("Icon had extension .icns but was not a directory")
				os.Exit(1)
			}
			must(buildIconSet(iconDst, args.Icon), "error building iconset:")
		case iconExt == ".png":
			fmt.Println("Building icons from raw images not implemented yet")
			os.Exit(1)
		default:
			fmt.Println("Unrecognized icon extension:", iconExt)
			os.Exit(1)
		}
	}

	// Write Info.plist
	tpl, err := template.New("info.plist.tpl").Parse(string(MustAsset("info.plist.tpl")))
	must(err)

	plistDst := filepath.Join(tmpBundle, "Contents", "Info.plist")
	w, err := os.Create(plistDst)
	must(err)

	tpl.Execute(w, map[string]interface{}{
		"BundleName":       args.Name,
		"BundleIdentifier": args.Identifier,
		"Extras":           extraProps,
	})
	must(w.Close())

	// Write PkgInfo. I copied this verbatim from another bundle.
	pkginfo := []byte{0x3f, 0x3f, 0x3f, 0x3f, 0x3f, 0x3f, 0x3f, 0x3f}
	pkginfoDst := filepath.Join(tmpBundle, "Contents", "PkgInfo")
	must(ioutil.WriteFile(pkginfoDst, pkginfo, 0777))

	// Delete the bundle.app dir if it already exists
	must(os.RemoveAll(args.Output))

	// Move the temporary dir to the bundle.app location
	must(os.Rename(tmpBundle, args.Output))
}
