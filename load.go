package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func LoadImages() map[string]*ebiten.Image {
	//assets location
	assetPath := "/assets"
	//dict for easy texture access
	Images := make(map[string]*ebiten.Image)
	if build {
		//get cwd of executable
		assetPath = GetCWDPath(assetPath) + "/assets"
		fmt.Println("asset path:" + assetPath)
	}
	if true {
		//loading images in the directory and saving the to images
		items, _ := os.ReadDir(assetPath)
		for _, item := range items {
			if !item.IsDir() {
				fmt.Println(assetPath + "/" + item.Name())
				appendImage, _, err := ebitenutil.NewImageFromFile(assetPath + "/" + item.Name())
				if err != nil {
					log.Fatal(err)
				}
				//like in PyKart, textures can be accessed with their file name
				fmt.Println("loaded image:" + item.Name())
				Images[item.Name()] = appendImage
			}
		}
	}
	return Images
}
func GetCWDPath(path string) string {
	//loading files in the directory of the executable
	filePath, _ := os.Executable()
	filePath = filepath.Dir(filePath)
	return filePath
}
