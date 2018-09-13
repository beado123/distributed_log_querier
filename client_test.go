package client

import "os"
import "testing"
import "os/exec"
import "fmt"
import "strconv"
import "strings"

func TestDifference(t *testing.T){
	queries := [1]string{"5qd"}
	
	for i := 0; i < len(queries); i++ {
		exec.Command("./client", queries[i], "3000").Run();
		ret := 0
		for j := 1; j <= 10; j++ {
			filename := "machine" + strconv.Itoa(j)+ ".i.log"
			if _, err := os.Stat(filename); err == nil {	
				cmd := exec.Command("wc", "-l", filename)
				count, err := cmd.CombinedOutput()
				_ = err
				words := strings.Fields(string(count))
				c, err := strconv.Atoi(words[0])	
				ret += c
				_ = err
			}
		}
		fmt.Println(ret)
	}
}

