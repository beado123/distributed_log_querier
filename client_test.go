package client

import "os"
import "testing"
import "os/exec"
import "fmt"
import "time"
import "bufio"

//This function checks the result of unit test
func TestDifference(t *testing.T){
	queries := [9]string{"apple", "apple", "apple", "apple", "apple", "apple", "apple", "apple", "apple"}
	filenames := [9]string{"rareOne.log","rareSome.log","rareAll.log","somewhatFrequentOne.log","somewhatFrequentSome.log","somewhatFrequentAll.log","frequentOne.log","frequentSome.log","frequentAll.log"}
	outputs := [9]int{1, 3, 10, 4, 12, 40, 5, 15, 50}
	numbers := [10]string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10"}
	
	for i := 0; i < len(queries); i++ {
		start := time.Now()
		exec.Command("./client", queries[i], filenames[i]).Run();
		ret := 0
		for j := 0; j < 10; j++ {
			filename := "machine" + numbers[j] + ".i.log"
			if _, err := os.Stat(filename); err == nil {	
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

//This function deletes a file
func deleteFile(f string) {
        if _, err := os.Stat("./" + f); err == nil {
                err := os.Remove("./" + f)
		_ = err
        }
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
