package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	dir := "/User/some" // 指定目录路径
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	counter := 7 // 开始的数字
	for _, file := range files {
		if !file.IsDir() {
			newName := fmt.Sprintf("download-%d%s", counter, filepath.Ext(file.Name()))
			oldPath := filepath.Join(dir, file.Name())
			newPath := filepath.Join(dir, newName)

			err := os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Println("Error renaming file:", err)
				return
			}
			fmt.Printf("Renamed %s to %s\n", oldPath, newPath)

			counter++
			if counter > 15 { // 当达到download15时停止
				break
			}
		}
	}
}
