package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)
func LoadImages() []*ebiten.Image{
	assetPath := "./assets"
	Images := []*ebiten.Image{}
	if build {
        assetPath = GetCWDPath(assetPath) + "/assets"
		fmt.Println("asset path:"+assetPath)
    }
	if true{
		items, _ := os.ReadDir(assetPath)
		for _, item := range items {
			if !item.IsDir() {
                fmt.Println(assetPath + "/" + item.Name())
				appendImage,_,err:= ebitenutil.NewImageFromFile(assetPath + "/" + item.Name())
				if err!= nil {
                    log.Fatal(err)
                }
				Images = append(Images, appendImage)

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