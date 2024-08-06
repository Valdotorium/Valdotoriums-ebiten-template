package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)
func fetchGameImagePaths() []string {
	imagePaths := []string{
		"assets/stone.png",
		"assets/packet.png",
	    "assets/grass.png"}
	return imagePaths
}

func LoadImages(imagePaths []string) map[string]*ebiten.Image {
	//assets location
	assetPath := "assets"
	//dict for easy texture access
	Images := make(map[string]*ebiten.Image)
	if true {
		//get cwd of executable
		assetPath = GetCWDPath() + "assets"
		fmt.Println("asset path:" + assetPath)
	}
	fmt.Println("asset path:" + assetPath)
	if true {
		//loading images in the list and saving the to images

		for image := range imagePaths {
			var path string = imagePaths[image]
			if true{
				path = JoinCWDPath(imagePaths[image])
			}
			fmt.Println(path)
			img, _, err := ebitenutil.NewImageFromFile(path)
            if err!= nil {
                panic(err)
            }
            Images[imagePaths[image]] = img
		}

	}
	return Images
}
func GetCWDPath() string {
	//loading files in the directory of the executable
	filePath, _ := os.Executable()
	filePath = filepath.Dir(filePath)
	return filePath
}
func JoinCWDPath(path string) string {
	filePath, _ := os.Executable()
	filePath = filepath.Dir(filePath)
	filePath = filepath.Join(filePath, path)
	return filePath
}
