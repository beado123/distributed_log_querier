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
	deleteFile("machine01.i.log")
	deleteFile("machine02.i.log")
	deleteFile("machine03.i.log")
	deleteFile("machine04.i.log")
	deleteFile("machine05.i.log")
	deleteFile("machine06.i.log")
	deleteFile("machine07.i.log")
	deleteFile("machine08.i.log")
	deleteFile("machine09.i.log")
	deleteFile("machine10.i.log")
}
