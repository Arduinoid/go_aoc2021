package main

import (
	"fmt"
	"io/ioutil"
)

const fileNameFmt = `./data/d%d.txt`

func readInput(day int) []byte {
	b, err := ioutil.ReadFile(fmt.Sprintf(fileNameFmt, day))
	if err != nil {
		panic(err)
	}
	return b
}
