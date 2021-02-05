package utils

import (
	"bufio"
	"bytes"
	"github.com/pkg/errors"
	"image"
	"image/png"
	"os"
)

func decode(filename string) (image.Image, string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()
	return image.Decode(bufio.NewReader(f))
}

func getImageBytes(filename string) ([]byte, error) {
	img, _, err := decode(filename)
	if err != nil {
		return []byte{}, errors.Wrap(err, "error decoding image")
	}

	buf := new(bytes.Buffer)
	err = png.Encode(buf, img)
	if err != nil {
		return []byte{}, errors.Wrap(err, "error encoding image")
	}

	return buf.Bytes(), nil
}
