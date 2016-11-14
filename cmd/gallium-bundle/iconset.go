package main

import (
	"errors"
	"os/exec"
)

// On macOS, app bundle icons are represents as a .icns file. These files
// are created using the "iconutil" tool. The input to this tool is a
// ".iconset" directory containing some combination of the files:
//    icon_16x16[@2x].png
//    icon_32x32[@2x].png
//    icon_128x128[@2x].png
//    icon_256x256[@2x].png
//    icon_512x512[@2x].png

func buildIconSet(dst, src string) error {
	cmd := exec.Command("iconutil", "-c", "icns", "-o", dst, src)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(output))
	}
	return nil
}
