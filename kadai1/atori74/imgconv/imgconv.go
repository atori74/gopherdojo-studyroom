package imgconv

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

type ImageType int

const (
	JPG ImageType = iota
	PNG
	GIF
)

func GetImageType(t string) (ImageType, bool) {
	switch t {
	case "jpeg", "jpg":
		return JPG, true
	case "png":
		return PNG, true
	case "gif":
		return GIF, true
	default:
		return -1, false
	}
}

func (t ImageType) Ext() string {
	switch t {
	case JPG:
		return ".jpg"
	case PNG:
		return ".png"
	case GIF:
		return ".gif"
	default:
		panic("Invalid ImageType")
	}
}

func Convert(file *os.File, image image.Image, toType ImageType) error {
	switch toType {
	case JPG:
		if err := jpeg.Encode(file, image, nil); err != nil {
			return err
		}
		return nil
	case PNG:
		if err := png.Encode(file, image); err != nil {
			return err
		}
		return nil
	case GIF:
		if err := gif.Encode(file, image, nil); err != nil {
			return err
		}
		return nil
	default:
		return errors.New("invalid type")
	}
}
