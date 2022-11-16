//Golang
// rw json file

package filefun

import (
	"encoding/json"
	"bufio"
	"fmt"
	"log"
	"os"
)

type Info struct {
	Name   string   `json:"name"`
	Mobile int      `json:"mobile"`
	Active bool     `json:"active"`
	Pets   []string `json:"pets"`
	//Address map[string]string `json:"address"`
	Address interface{} `json:"address"`
}

func StructuredWrite() {
	info := &Info{
		Name:   "David",
		Mobile: 12345,
		Active: true,
		Pets:   []string{"Dog", "Cat"},
		Address: map[string]string{"City": "Austin",
			"state": "Texas",
		},
	}


// indenting json during marshal, (1 space)
	jInfo, err := json.MarshalIndent(info, ""," ")
	if err != nil {
		log.Panic(err)
	}
	
	// either append/create file
	file, err := os.OpenFile("testing.json", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	bufferedWriter := bufio.NewWriter(file)

	bytesWritten, err := bufferedWriter.Write(jInfo)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Bytes written: %d\n", bytesWritten)

	// write to disk
	bufferedWriter.Flush()
	// reset
	bufferedWriter.Reset(bufferedWriter)

}

func StructuredRead() {

	content, err := os.ReadFile("test2.json")
	if err != nil {
		log.Panic(err)
	}

	var payload Info
	if err = json.Unmarshal(content, &payload); err != nil {
		log.Panic(err)
	}

	fmt.Println(payload)
}

func UnstructuredRead(filepath string) error {

	file, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	var payload map[string]interface{}
	err = json.Unmarshal(file, &payload)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(payload)

	return nil
}
