package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"strings"
	"crypto/rand"
)

//This function helps check for errors
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//This function extracts ip address of current VM from file "ip_address" in current directory
func getIPAddr() string{

	data, err := ioutil.ReadFile("ip_address")
	if err != nil {
		panic(err)
	}

	ip := string(data[:len(data)])
	
	//remove \n from end of line
	if strings.HasSuffix(ip, "\n") {
		ip = ip[:(len(ip) - 1)]
	}
	fmt.Println("ip address of current VM:\n", ip)
	return ip
}

//This function generates a random string of length 15
func randomStr() string{
	b := make([]byte, 15)
	_, err := rand.Read(b)
	check(err)
	return string(b[:15]) + "\n"
}

//This function writes lines to file for test case rare(frequency) and one(occurs in one log file)
func rareOne(machineNum string, fp *os.File) {
	if machineNum == "01" {
		n, err := fp.Write( []byte("This is an apple.\nToday is a sunny day.\nI love pizza.\n" + randomStr()+ randomStr() + randomStr()))
		check(err)
		fmt.Printf("Wrote %d bytes\n", n)
	} else {
		n, err := fp.Write([]byte("This is an banana.\nToday is a sunny day.\nI love pizza.\n" + randomStr() + randomStr() + randomStr()))
		check(err)
		fmt.Printf("Wrote %d bytes\n", n)
	}
}

//This function writes lines to file for test case rare(frequency) and Some(occurs in some log files)
func rareSome(machineNum string, fp *os.File) {
	if (machineNum == "01" || machineNum == "02" || machineNum == "06") {
		n, err := fp.Write([]byte("This is an apple.\nToday is a sunny day.\nI love pizza.\n" + randomStr() + randomStr() + randomStr()))
		check(err)
		fmt.Printf("Wrote %d bytes\n", n)
	} else {
		n, err := fp.Write([]byte("This is an banana.\nToday is a sunny day.\nI love pizza\n" + randomStr() + randomStr() + randomStr()))
		check(err)
		fmt.Printf("Wrote %d bytes\n", n)
	}
}

//This function writes lines to file for test case rare(frequency) and All(occurs in all log files)
func rareAll(machineNum string, fp *os.File) {

	n, err := fp.Write([]byte("This is an apple.\nToday is a sunny day.\nI love pizza\n" + randomStr() + randomStr() + randomStr()))
	check(err)
	fmt.Printf("Wrote %d bytes\n", n)
	
}

//This function writes lines to file for test case somewhatFrequent(frequency) and one(occurs in one log file)
func somewhatFrequentOne(machineNum string, fp *os.File) {
	if machineNum == "01" {
		n, err := fp.Write([]byte("This is an apple.\nToday is a sunny day.\nI love apple\napple is essential.\napple is not delicious.\n" + randomStr() + randomStr() + randomStr() + randomStr() ))
		check(err)
		fmt.Printf("Wrote %d bytes\n", n)
	} else {
		n, err := fp.Write([]byte("This is an banana.\nToday is a sunny day.\nI love banana\nbanana is essential.\nbanana is not delicious.\n" + randomStr() + randomStr() + randomStr() + randomStr() ))
		check(err)
		fmt.Printf("Wrote %d bytes\n", n)
	}
}

//This function writes lines to file for test case somewhatFrequent(frequency) and some(occurs in some log files)
func somewhatFrequentSome(machineNum string, fp *os.File) {
	if (machineNum == "01" || machineNum == "02" || machineNum == "06") {
		n, err := fp.Write([]byte("This is an apple.\nToday is a sunny day.\nI love apple\napple is essential.\napple is not delicious.\n" + randomStr() + randomStr() + randomStr() + randomStr() ))
		check(err)
		fmt.Printf("Wrote %d bytes\n", n)
	} else {
		n, err := fp.Write([]byte("This is an banana.\nToday is a sunny day.\nI love banana\nbanana is essential.\nbanana is not delicious.\n" + randomStr() + randomStr() + randomStr() + randomStr() ))
		check(err)
		fmt.Printf("Wrote %d bytes\n", n)
	}
}

//This function writes lines to file for test case somewhatFrequent(frequency) and all(occurs in all log files)
func somewhatFrequentAll(machineNum string, fp *os.File) {

	n, err := fp.Write([]byte("This is an apple.\nToday is a sunny day.\nI love apple\napple is essential.\napple is not delicious.\n" + randomStr() + randomStr() + randomStr() + randomStr() ))
	check(err)
	fmt.Printf("Wrote %d bytes\n", n)
	
}

//This function writes lines to file for test case frequent(frequency) and one(occurs in one log file)
func frequentOne(machineNum string, fp *os.File) {
	if machineNum == "01" {
		n, err := fp.Write([]byte("apple is apple an apple.\napple is an apple apple.\nI love apple apple\napple is apple.\napple is not apple\n" + randomStr() + randomStr() ))
		check(err)
		fmt.Printf("Wrote %d bytes\n", n)
	} else {
		n, err := fp.Write([]byte("This is an banana.\nToday is a sunny day.\nI love banana\nbanana is essential.\nbanana is not delicious.\n" + randomStr() + randomStr() + randomStr() + randomStr() ))
		check(err)
		fmt.Printf("Wrote %d bytes\n", n)
	}
}

//This function writes lines to file for test case frequent(frequency) and some(occurs in some log files)
func frequentSome(machineNum string, fp *os.File) {
	if (machineNum == "01" || machineNum == "02" || machineNum == "06") {
		n, err := fp.Write([]byte("apple is apple an apple.\napple is an apple apple.\nI love apple apple\napple is apple.\napple is not apple\n" + randomStr() + randomStr() ))
		check(err)
		fmt.Printf("Wrote %d bytes\n", n)
	} else {
		n, err := fp.Write([]byte("This is an banana.\nToday is a sunny day.\nI love banana\nbanana is essential.\nbanana is not delicious.\n" + randomStr() + randomStr() + randomStr() + randomStr() ))
		check(err)
		fmt.Printf("Wrote %d bytes\n", n)
	}
}

//This function writes lines to file for test case frequent(frequency) and all(occurs in all log files)
func frequentAll(machineNum string, fp *os.File) {

	n, err := fp.Write([]byte("apple is apple an apple.\napple is an apple apple.\nI love apple apple\napple is apple.\napple is not apple\n" + randomStr() + randomStr() ))
	check(err)
	fmt.Printf("Wrote %d bytes\n", n)
}
	
//Main function that creates 9 testing log files
func main() {

	ip := getIPAddr()
	fmt.Println("current ip address:", ip)
	machineNum := ip[15:17]
	fmt.Println("machine number:", machineNum)

	//rare pattern, exists in one log file
	f, err := os.Create("./rareOne.log")
	check(err)
	defer f.Close()
	rareOne(machineNum, f)

	//rare pattern, exists in some log files(01,02,06)
	f, err = os.Create("./rareSome.log")
	check(err)
	defer f.Close()
	rareSome(machineNum, f)

	//rare pattern, exists in all log files
	f, err = os.Create("./rareAll.log")
	check(err)
	defer f.Close()
	rareAll(machineNum, f)

	//somewhat frequent pattern(appears 4 times), exists in one log file
	f, err = os.Create("./somewhatFrequentOne.log")
	check(err)
	defer f.Close()
	somewhatFrequentOne(machineNum, f)

	f.Sync()

	//somewhat frequent pattern(appears 4 times), exists in some log file(01,02,06)
	f, err = os.Create("./somewhatFrequentSome.log")
	check(err)
	defer f.Close()
	somewhatFrequentSome(machineNum, f)

	//somewhat frequent pattern(appears 4 times), exists in all log files
	f, err = os.Create("./somewhatFrequentAll.log")
	check(err)
	defer f.Close()
	somewhatFrequentAll(machineNum, f)

	//frequent pattern(appears 12 times), exists in one log file
	f, err = os.Create("./frequentOne.log")
	check(err)
	defer f.Close()
	frequentOne(machineNum, f)

	//frequent pattern(appears 12 times), exists in some log files(01,02,06)
	f, err = os.Create("./frequentSome.log")
	check(err)
	defer f.Close()
	frequentSome(machineNum, f)

	//frequent pattern(appears 12 times), exists in all log files
	f, err = os.Create("./frequentAll.log")
	check(err)
	defer f.Close()
	frequentAll(machineNum, f)
	
}
