package http_utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"image"
	"image/png"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
)

func GetImageFileByUrl(oriPictureUrl string) (*os.File, error) {
	img, err := downloadImage(oriPictureUrl)
	if err != nil {
		logrus.Errorf("[VariablePictures] downloadPict err; %v\n", err)
		return nil, err
	}
	tempFile, err := ioutil.TempFile("", fmt.Sprintf("img_%v.png", rand.Intn(100)))
	if err != nil {
		logrus.Errorf("[VariablePictures] TempFile err; %v\n", err)
		return nil, err
	}
	defer tempFile.Close()
	err = png.Encode(tempFile, img)
	if err != nil {
		logrus.Errorf("[VariablePictures] encode png err; %v\n", err)
		return nil, err
	}
	return tempFile, nil
}

func downloadImage(url string) (image.Image, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	img, _, err := image.Decode(response.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}
