package main

import "net"
import "fmt"
import "bufio"
import "os"
import "io"
import "encoding/json" 
import "io/ioutil"
//import "strconv"
import "sync"
import "time"


func main(){
	//get grep command and port number from command-line arguments
	if len(os.Args) < 3 {
		fmt.Println("Please type in grep command and port number!")
		return
	}
	grep_cmd := os.Args[1]
	port_num := "3000"
	file_name := os.Args[2]

	//parse json file get each server information
	jsonFile, err := os.Open("servers.json") 
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
	start := time.Now()
	for i := 0; i < len(server_info.Servsers); i++ {
		go func(Hostname string, Logfile string, grep_cmd string, port_num string, id string){
			//connect to server
			conn, err := net.Dial("tcp", Hostname + ":" + port_num)
			if err != nil {
				fmt.Println(err)
				wg.Done()
				return;
			}
			//send to socket
			name := "machine" + id
			fmt.Fprintf(conn, grep_cmd + " " + name + " " + file_name)
			//create log file
			f, err := os.Create(Logfile)
			if err != nil {
				fmt.Println(err)
				wg.Done()
                                return;
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
					wg.Done()
                                	return;
				}
				text := string(message[:n1])
				fmt.Println(text)
				n2, err := f.WriteString(text)
				if err != nil {
					fmt.Println(err)
					wg.Done()
                                	return;
				}
				_ = n1
				_ = n2
			}	
			wg.Done()	
		}(server_info.Servsers[i].Hostname, server_info.Servsers[i].Logfile, grep_cmd, port_num, server_info.Servsers[i].Id)		
	}
	wg.Wait()
	end := time.Now()	
	elipsed := end.Sub(start)
	//print total line number of each file
	ret := 0
	for i := 0; i < len(server_info.Servsers); i++ {
		lc, _ := lineCount(server_info.Servsers[i].Logfile)
		name := "machine" + server_info.Servsers[i].Id
		fmt.Println(name, lc)
                ret += lc
	}
	fmt.Println("total:", ret)
	fmt.Println("latency: ", elipsed)
}

type Servsers struct {
	Servsers []serverInfo `json:"server_list"`
}

type serverInfo struct {
	Id    string    `json:"id"`
	Hostname   string `json:"hostname"`
	Logfile   string `json:"logfile"`
}

//This function calculates line number of a file
func lineCount(filename string) (int, error) {
    count := 0
    f, err := os.Open(filename)
    if err != nil {
        return 0, err
    }
    defer f.Close()
    s := bufio.NewScanner(f)
    for s.Scan() {
        if len(s.Text()) > 0 {
                count++
        }
    }
    return count, s.Err()
}

