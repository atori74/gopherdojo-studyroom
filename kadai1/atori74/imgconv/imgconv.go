// imgconv package is for converting image to different type.
package imgconv

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

// ImageType is the type of image supported to convert.
type ImageType int

const (
	JPG ImageType = iota
	PNG
	GIF
)

// GetImageType receives file type of images and returns ImageType.
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

// Ext returns extension to output as a file.
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

// Convert encodes image as given ImageType and write that to given file.
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
