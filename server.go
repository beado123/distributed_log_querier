package main

import (
	"net"
	"os"
	"fmt"
	"os/exec"
	"strings"
	"io/ioutil"
)

//This function helps printing out errors
func printErr(err error, s string) {
	if err != nil {
		fmt.Println("Error occurs on ", s , "\n" , err.Error())
		os.Exit(1)
	}
}

//This function helps printing out commads that are executing
func printCommand(cmd *exec.Cmd) {
    fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

//This function executes grep command
func executeGrep(query string, vm string) []byte{

	cmd := exec.Command("grep", "-nr", "--text", query, vm)
   	printCommand(cmd)
    output, err := cmd.CombinedOutput()

   	 if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Grep Error: %s\n", err.Error()))
	}
	return output
}

//This function parses request sent from client and sends the result back to client
//commadn format: query logfile_name
func parseRequest(conn net.Conn) {

	//create a buffer to hold transferred data and read incoming data into buffer
	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	printErr(err, "reading")

	//convert request command into array
	reqArr := strings.Split(string(buf[:reqLen]), " ")
		
	//execute grep
	output := executeGrep(reqArr[0], reqArr[2])
	
	//append vm name to each grep result
	out := ""
	arr := strings.Split(string(output), "\n")	
	for i := 0; i<len(arr)-1; i++ {
		if i == len(arr) - 2 {
			out = out + reqArr[1] + " " + "line " + arr[i]
		} else {
			out = out + reqArr[1] + " " + "line " + arr[i] + "\n"
		}
	}
	
	//send response
	conn.Write([]byte(out))
	//close connection
	conn.Close()
}

//This function extracts ip address of current VM from file "ip_address" in current directory
func getIPAddrAndLogfile() string{

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

//Main function that starts the server and listens for incoming connections
func main() {

	//get ip address from servers list	
	ip := getIPAddrAndLogfile()
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
	
