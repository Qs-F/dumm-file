package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("please set a number of files")
	} else if v, err := strconv.Atoi(os.Args[1]); err == nil && len(os.Args) == 2 {
		for i := 0; i < v; i++ {
			ioutil.WriteFile("dumm_file_"+strconv.Itoa(i), []byte{}, 0700)
		}
	}
}
