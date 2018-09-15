package client

import "os"
import "testing"
import "os/exec"
import "fmt"
//import "strconv"
//import "strings"
import "time"
import "bufio"

func TestDifference(t *testing.T){
	queries := [9]string{"apple", "apple", "apple", "apple", "apple", "apple", "apple", "apple", "apple"}
	filenames := [9]string{"rareOne.log","rareSome.log","rareAll.log","somewhatFrequentOne.log","somewhatFrequentSome.log","somewhatFrequentAll.log","frequentOne.log","frequentSome.log","frequentAll.log"}
	outputs := [9]int{1, 3, 10, 4, 12, 40, 5, 15, 50}
	numbers := [10]string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10"}
	//queries := [1]string{"apple"}
	//filenames := [1]string{"rareOne.log"}
	//outputs := [1]int{1}
	
	for i := 0; i < len(queries); i++ {
		start := time.Now()
		exec.Command("./client", queries[i], "3000", filenames[i]).Run();
		ret := 0
		for j := 0; j < 10; j++ {
			filename := "machine" + numbers[j] + ".i.log"
			if _, err := os.Stat(filename); err == nil {	
				/*cmd := exec.Command("wc", "-l", filename)
				count, err := cmd.CombinedOutput()
				_ = err
				words := strings.Fields(string(count))
				c, err := strconv.Atoi(words[0])	
				ret += c
				_ = err*/
				lc, _ := lineCount(filename)
				ret += lc
				//fmt.Println(ret)
			}
		}
		end := time.Now()
		elipsed := end.Sub(start)
		fmt.Println(ret)
		fmt.Println(elipsed)
		if ret != outputs[i] {
			t.Errorf("Test failed", filenames[i])
		}
		for j := 0; j < 10; j++ {
                        filename := "machine" + numbers[j] + ".i.log"
			if _, err := os.Stat(filename); err == nil {
				deleteFile(filename)
			}
		}
	}
}

func deleteFile(f string) {
        if _, err := os.Stat("./" + f); err == nil {
                err := os.Remove("./" + f)
                //check(err)
                fmt.Println("Deleted " + f)
		_ = err
        }
}

func lineCount(filename string) (int, error) {
    lc := 0
    f, err := os.Open(filename)
    if err != nil {
        return 0, err
    }
    defer f.Close()
    s := bufio.NewScanner(f)
    for s.Scan() {
	if len(s.Text()) > 0 {
        	lc++
	}
    }
    return lc, s.Err()
}
