package http_utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"math/rand"
	"net/http"
	"os"
)

/*

func main() {
	// 图片URL
	imageURL := "https://example.com/path/to/your/image.png"

	// 获取图片
	resp, err := http.Get(imageURL)
	if err != nil {
		fmt.Println("Error fetching image:", err)
		return
	}
	defer resp.Body.Close()

	// 创建临时文件
	tmpFile, err := os.CreateTemp("", "downloaded_image_*.png")
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return
	}
	defer tmpFile.Close()

	// 将图片写入临时文件
	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		fmt.Println("Error writing image to file:", err)
		return
	}

	// 将临时文件指针重置到文件开头
	tmpFile.Seek(0, 0)

	fmt.Printf("Image saved as *os.File: %s\n", tmpFile.Name())
}
*/

func GetPngImageFileByUrl(oriPictureUrl string) (*os.File, error) {
	response, err := http.Get(oriPictureUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// 创建临时文件
	tmpFile, err := os.CreateTemp("", fmt.Sprintf("temp_img_%v.png", rand.Intn(100)))
	if err != nil {
		logrus.Errorf("[downloadImage] Error creating temp file:; %v\n", err)

		return nil, err
	}

	_, err = io.Copy(tmpFile, response.Body)
	if err != nil {
		logrus.Errorf("[downloadImage] Error writing image to file:; %v\n", err)
		return nil, err
	}
	// 将临时文件指针重置到文件开头
	tmpFile.Seek(0, 0)

	return tmpFile, nil
}
