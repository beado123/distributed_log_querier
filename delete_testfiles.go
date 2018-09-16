package main

import (
	"os"
	"fmt"
)

//This function checks for errors and prints them
func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

//This function first checks whether file exists and if yes, then delete it
func deleteFile(f string) {
	if _, err := os.Stat("./" + f); err == nil {
		err := os.Remove("./" + f)
		check(err)
		fmt.Println("Deleted " + f)
	}
}

//Main function that deletes a set of log files
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
