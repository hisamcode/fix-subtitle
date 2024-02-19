package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func fixSub(path string) {
	ext := filepath.Ext(path)
	if ext == ".srt" {
		base := filepath.Base(path)
		dir := filepath.Dir(path)
		baseWithoutExtension := base[:len(base)-len(ext)]
		newName := baseWithoutExtension + "-fixed" + ext
		newPath := filepath.Join(dir, newName)
		fileNew, err := os.Create(newPath)
		if err != nil {
			log.Fatal(err)
		}
		defer fileNew.Close()

		openFile, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer openFile.Close()

		scanner := bufio.NewScanner(openFile)
		for scanner.Scan() {
			text := scanner.Text()
			regex := regexp.MustCompile(`[0-9]+:[0-9]+,[0-9]+`)
			regexResult := regex.FindAllString(text, -1)
			edited := ""
			if len(regexResult) > 1 {
				edited = "00:" + regexResult[0] + " --> 00:" + regexResult[1]
			} else {
				edited = text
			}
			edited += "\n"
			fileNew.WriteString(edited)
		}
		fmt.Println(newPath, "fixed")
	}
}

func listDir(rootPath string) error {
	paths, err := os.ReadDir(rootPath)
	if err != nil {
		return err
	}

	for _, path := range paths {

		fullPath := filepath.Join(rootPath, path.Name())
		if path.IsDir() {
			listDir(fullPath)
		} else {
			fixSub(fullPath)
		}

	}
	return nil
}

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	var input string
	fmt.Println("Confirm this dir? yes/no")
	_, err = fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}

	if input == "yes" {
		err = listDir(currentDir)
		if err != nil {
			fmt.Println(err)
		}
	}

}
