package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"strings"
	"crypto/rand"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getIPAddrAndLogfile() (string, string){

	data, err := ioutil.ReadFile("ip_address")
	//info := string(data[:])
	arr := strings.Split(string(data[:]), " ") 
	if err != nil {
		panic(err)
	}
	if strings.HasSuffix(arr[1], "\n") {
		arr[1] = arr[1][:len(arr[1]) - 1]
	}
	//fmt.Println("ip address of current VM:" + arr[0])
	//fmt.Println(arr[1])
	return arr[0],arr[1]
}

//generates a random string of length 15
func randomStr() string{
	b := make([]byte, 15)
	_, err := rand.Read(b)
	check(err)
	return string(b[:15]) + "\n"
}

//query: apple
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

//query: apple
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

//query: apple
func rareAll(machineNum string, fp *os.File) {

	n, err := fp.Write([]byte("This is an apple.\nToday is a sunny day.\nI love pizza\n" + randomStr() + randomStr() + randomStr()))
	check(err)
	fmt.Printf("Wrote %d bytes\n", n)
	
}

//query: apple
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

//query: apple
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

//query: apple
func somewhatFrequentAll(machineNum string, fp *os.File) {

	n, err := fp.Write([]byte("This is an apple.\nToday is a sunny day.\nI love apple\napple is essential.\napple is not delicious.\n" + randomStr() + randomStr() + randomStr() + randomStr() ))
	check(err)
	fmt.Printf("Wrote %d bytes\n", n)
	
}

//query: apple
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

//query: apple
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

//query: apple
func frequentAll(machineNum string, fp *os.File) {

	n, err := fp.Write([]byte("apple is apple an apple.\napple is an apple apple.\nI love apple apple\napple is apple.\napple is not apple\n" + randomStr() + randomStr() ))
	check(err)
	fmt.Printf("Wrote %d bytes\n", n)
}
	

func main() {

	ip, _ := getIPAddrAndLogfile()
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
