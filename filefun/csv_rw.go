// Golang
// file fun
// r/w csv

package filefun

import (
	"encoding/csv"
	//	"fmt"
	"io"
	"log"
	"os"
)

func check(e error) {
	// currently looking more efficient way of handling errs
	if e != nil {
		log.Panic(e)
	}
}

// not returning errors right now
//

func CsvReadLine(f string) [][]string {

	file, err := os.Open(f)
	check(err)
	defer file.Close()

	reader := csv.NewReader(file)

	// change delim if necessary, default = ','
	//reader.Comma = ';'

	data := [][]string{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		check(err)

		//fmt.Println(record)
		data = append(data, record)
	}

	return data
}

func CsvReadAll(f string) [][]string {

	file, err := os.Open(f)
	check(err)
	defer file.Close()

	reader := csv.NewReader(file)

	// change delim if necessary, default = ','
	//reader.Comma = ';'

	data, err := reader.ReadAll()
	check(err)

	return data

}

func CsvWrite(records [][]string, writeFile string) {

	// create file
	//file, err := os.Create(writeFile)

	//create file if non-exsistent, append otherwise
	file, err := os.OpenFile(writeFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer file.Close()

	writer := csv.NewWriter(file)

	//change delim, default = ','
	//writer.Comma = ';'

	for _, record := range records {

		if err := writer.Write(record); err != nil {
			check(err)
		}
	}

	//write to disk, clear buffer
	writer.Flush()
	err = writer.Error()
	check(err)

}
