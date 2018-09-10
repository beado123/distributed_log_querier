package main

import (
	"net"
	"os"
	"fmt"
	"os/exec"
	"strings"
	"io/ioutil"
)

func printErr(err error, s string) {
	if err != nil {
		fmt.Println("Error occurs on ", s , "\n" , err.Error())
		os.Exit(1)
	}
}

func printCommand(cmd *exec.Cmd) {
  fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printOutput(outs []byte) {
  if len(outs) > 0 {
    fmt.Printf("==> Output:\n%s\n", string(outs))
  }
}

func executeGrep(query string) []byte{

	cmd := exec.Command("grep", "-nr", query, "machine.i.log")
    printCommand(cmd)
    output, err := cmd.CombinedOutput()
	//print error
    if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
	return output
}

func parseRequest(conn net.Conn) {

	//create a buffer to hold transferred data
	buf := make([]byte, 1024)
	//read incoming data into buffer
	reqLen, err := conn.Read(buf)
	printErr(err, "reading")
	fmt.Println("reqLen:",reqLen)

	//put request command into array
	reqArr := strings.Split(string(buf[:reqLen]), " ")
	
	fmt.Println("received query:", reqArr[0])
		
	//execute grep
	output := executeGrep(reqArr[0])

	//append vm name to each grep result
	arr := strings.Split(string(output), "\n")
	out := ""
	for i := 0; i<len(arr)-1; i++ {
		if i == len(arr) - 2 {
			out = out + reqArr[1] + " " + "line " + arr[i]
		} else {
			out = out + reqArr[1] + " " + "line " + arr[i] + "\n"
		}
	}
	fmt.Println(out)
	
	//send response
	conn.Write([]byte(out))
	//close connection
	conn.Close()
}

func getIPAddr() string{

	data, err := ioutil.ReadFile("ip_address")
	ip := string(data[:])
	if err != nil {
		panic(err)
	}
	if strings.HasSuffix(ip, "\n") {
		ip = ip[:len(ip) - 1]
	}
	fmt.Println("ip address of current VM:" + ip)
	return ip
}

func main() {

	//get ip address from servers list	
	ip := getIPAddr()
	//listen for incoming connections
	l, err := net.Listen("tcp", ip + ":3000")
	printErr(err, "listening")
	
	//close the listener when app closes
	defer l.Close()
	fmt.Println("Listening on port 3000")

	//Listen for incoming connections
	for {
		conn, err := l.Accept()
		fmt.Println("Accept:", conn.RemoteAddr().String())
		printErr(err, "accepting")

		go parseRequest(conn)
	}
}
	
