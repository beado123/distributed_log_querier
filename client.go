package main

import "net"
import "fmt"
//import "bufio"
import "os"
import "io"
import "encoding/json" 
import "io/ioutil"
//import "strconv"
import "sync"


func main(){
	//get grep command and port number from command-line arguments
	if len(os.Args) < 4 {
		fmt.Println("Please type in grep command and port number!")
		return
	}
	grep_cmd := os.Args[1]
	port_num := os.Args[2]
	file_name := os.Args[3]

	//parse json file get each server information
	jsonFile, err := os.Open("test.json") 
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile) 
	var server_info Servsers
	json.Unmarshal(byteValue, &server_info)
	/*for i := 0; i < len(server_info.Servsers); i++ {
		fmt.Println("Server Id: " + server_info.Servsers[i].Id)
		fmt.Println("Server Hostname: " + server_info.Servsers[i].Hostname)
		fmt.Println("Server Logfile: " + server_info.Servsers[i].Logfile)
	}*/
	
	//build a connect with each server
	var wg sync.WaitGroup
	wg.Add(len(server_info.Servsers))
	for i := 0; i < len(server_info.Servsers); i++ {
		go func(Hostname string, Logfile string, grep_cmd string, port_num string, id string){
			//connect to server
			conn, _ := net.Dial("tcp", Hostname + ":" + port_num)
			//send to socket
			name := "machine" + id + ".i.log"
			fmt.Fprintf(conn, grep_cmd + " " + name + " " + file_name)
			//create log file
			f, err := os.Create(Logfile)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer f.Close()
	
			//read message from socket and write to log file
			for true {
				message := make([]byte, 5120)
				n1, err := conn.Read(message)
				if err != nil {
					if err == io.EOF {
						break	
					}
					fmt.Println(err)
					os.Exit(1)	
				}
				text := string(message[:n1])
				fmt.Println(text)
				n2, err := f.WriteString(text)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				_ = n1
				_ = n2
			}	
			wg.Done()	
		}(server_info.Servsers[i].Hostname, server_info.Servsers[i].Logfile, grep_cmd, port_num, server_info.Servsers[i].Id)		
	}
	wg.Wait()
	//fmt.Println("done")
}

type Servsers struct {
	Servsers []serverInfo `json:"server_list"`
}

type serverInfo struct {
	Id    string    `json:"id"`
	Hostname   string `json:"hostname"`
	Logfile   string `json:"logfile"`
}
