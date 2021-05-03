package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"gopherdojo-studyroom/kadai1/atori74/imgconv"
)

var flagDir = flag.String("dir", "", "translate images under this directory")
var flagFrom = flag.String("from", "jpg", "from format")
var flagTo = flag.String("to", "png", "to format")

func main() {
	flag.Parse()

	if *flagDir == "" {
		fmt.Println("dir flag cannot be empty")
		return
	}
	if info, err := os.Stat(*flagDir); err != nil {
		fmt.Println(err)
		return
	} else if !info.IsDir() {
		fmt.Println("invalid dir path")
		return
	}

	fromType, ok := imgconv.GetImageType(*flagFrom)
	if !ok {
		fmt.Println("Not supported file type")
		return
	}
	toType, ok := imgconv.GetImageType(*flagTo)
	if !ok {
		fmt.Println("Not supported file type")
		return
	}
	if fromType == toType {
		fmt.Println("from/to is same format")
		return
	}

	err := filepath.WalkDir(*flagDir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)
		t, ok := imgconv.GetImageType(strings.Replace(ext, ".", "", 1))
		if !ok {
			return nil
		}
		if t == fromType {
			fmt.Println(path)
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()
			image, _, err := image.Decode(f)
			if err != nil {
				return err
			}
			outpath := path[0:len(path)-len(filepath.Ext(path))] + toType.Ext()
			outfile, err := os.Create(outpath)
			if err != nil {
				return err
			}
			defer outfile.Close()
			if err := imgconv.Convert(outfile, image, toType); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}
