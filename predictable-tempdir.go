package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	tempDirName := func(pattern string) string {
		logsDir, err := ioutil.TempDir(os.TempDir(), pattern)
		if err != nil {
			panic(err)
		}
		return logsDir
	}

	fmt.Println(tempDirName("*-my-app"))
	fmt.Println(tempDirName("my-app-*"))
}
