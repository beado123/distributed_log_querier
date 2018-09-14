package main

import (
	"os"
	"fmt"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func deleteFile(f string) {
	if _, err := os.Stat("./" + f); err == nil {
		err := os.Remove("./" + f)
		check(err)
		fmt.Println("Deleted " + f)
	}
}

func main() {

	deleteFile("rareOne.log")
	deleteFile("rareSome.log")
	deleteFile("rareAll.log")
	deleteFile("somewhatFrequentOne.log")
	deleteFile("somewhatFrequentSome.log")
	deleteFile("somewhatFrequentAll.log")
	deleteFile("frequentOne.log")
	deleteFile("frequentSome.log")
	deleteFile("frequentAll.log")
}
