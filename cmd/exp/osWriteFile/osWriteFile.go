package osWriteFile

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// i want to know what happen when copying text from another file to a created file
func Run() {
	// current path when running it, run on root path
	curpath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(curpath)
	paths1, _ := os.ReadDir(curpath)
	paths2, _ := os.ReadDir(paths1[0].Name())
	paths3 := filepath.Join(curpath, paths1[0].Name(), paths2[0].Name())
	fmt.Println(paths3)
	b, err := os.ReadFile(paths3)
	if err != nil {
		fmt.Println(err)
	}

	curpath = filepath.Join(curpath, "cmd", "exp", "osWriteFile", "tes.txt")
	err = os.WriteFile(curpath, b, 0666)
	if err != nil {
		fmt.Println(err)
	}

	f, err := os.OpenFile(curpath, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		fmt.Println(err)
	}

	_, err = f.WriteString("coba coba coba testing")
	if err != nil {
		fmt.Println(err)
	}

	_, _ = f.WriteAt([]byte("test1"), 0)
	_, _ = f.WriteAt([]byte("test2"), 1)
	_, _ = f.WriteAt([]byte("test3"), 30)

	f2, _ := os.Open(curpath)
	defer f2.Close()
	f3, _ := os.Create(curpath + "2.txt")
	defer f3.Close()
	scanner := bufio.NewScanner(f2)
	for scanner.Scan() {
		text := scanner.Text()
		// fmt.Println(text)
		// regex1 := regexp.MustCompile(`[0-9]+:[0-9]+,[0-9]+/g`)
		regex := regexp.MustCompile(`[0-9]+:[0-9]+,[0-9]+`)
		s1 := regex.FindAllString(text, -1)
		// fmt.Println(s1)
		// fmt.Println(len(s1))
		s2 := ""
		if len(s1) > 1 {
			s2 = "00:" + s1[0] + " --> 00:" + s1[1]
		} else {
			s2 = text
		}
		s2 += "\n"
		f3.WriteString(s2)

	}
}
