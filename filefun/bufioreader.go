// Golang
// bufio testing, for reference

package nonsense

import (
	//"fmt"
	"log"
	//sort "github.com/jaredmyers/gogym/sorting"
	//ds "github.com/jaredmyers/gogym/datastructures"
	//ff "github.com/jaredmyers/gogym/filefun"
	"bufio"
	"os"
//	"io"
//		"strings"
	"strconv"
)

func notmain() {
		
		filepath := "/home/dev/common/source_cbook/ch5/test.txt"
		file, err := os.Open(filepath)
		if err != nil {
				log.Fatal("OpeningFile:",err)
		}
		defer file.Close()


		// ====== Using Scanner =============================================
		
		scanner := bufio.NewScanner(file)

		lines := []string{}

		for scanner.Scan() {
				line := scanner.Text()
				lines = append(lines, line)
		}

		log.Println("Number of numbers in file:", len(lines))
		runningTotal := 0
		for _, num := range lines{

				i, _ := strconv.Atoi(num)
				runningTotal += i
		}

		avg := runningTotal / len(lines)
		log.Println("Running Total:", runningTotal)
		log.Println("Average:", avg)


		// ========= Using Reader w/ReadBytes ==============================

		reader := bufio.NewReader(file)
		lines := []string{}
		for {
				line, err := reader.ReadBytes('\n')

				if err != nil{
						break
				}
 
				// remove newline char
				line = line[:len(line)-1]
				lines = append(lines, string(line))
		}

		log.Println(lines)


		// ======== bufio NewReader w/ Readstring =========================

		reader := bufio.NewReader(file)
		lines := []string{}

		for {
				line, err := reader.ReadString('\n')
				if err != nil {
						break	
				}

				log.Print(len(line)," ", line)
				line = line[:len(line)-1]

				lines = append(lines, line)
		}

		log.Println(lines)


		// ======= bufio NewReader w/ReadLine ============================

		reader := bufio.NewReader(file)
		lines := []string{}

		for {
				line, prefix, err := reader.ReadLine()
				if err != nil || prefix != false {
						break
				}

				lines = append(lines, string(line))
		}

		log.Println(lines)


}
