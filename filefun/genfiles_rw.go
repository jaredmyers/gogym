// Golang
// file fun
// r/w general files

/*
for reference
just some testing

bufferedWrite():
		using bufio.NewWriter
		with Write & WriteString
bufferedRead():
		using bufio.NewReader
		with ReadString('\n')
bufferedRead2():
		using bufio.NewScanner
		testing by lines & words
directWrite()
		using ioutil.WriteFile
		depreciated
directReadAll():
		using ioutil.ReadAll
		depreciated
directWrite2():
		using Write & WriteString
		direct on the file
cpFile():
		using io.Copy
directReadFull():
		using io.ReadFull
directReadAtLeast():
		using io.ReadAtLeast
readWithBuffer()
*/

package filefun

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func BufferedWrite(filepath string, b []byte) error {
	// ---------------------------
	// using NewWriter w/bufio
	// for writing general file
	// ---------------------------
	file, err := os.OpenFile(filepath, os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	bufferedWriter := bufio.NewWriter(file)

	bytesWritten, err := bufferedWriter.Write(b)
	check(err)
	fmt.Printf("Bytes written: %d\n", bytesWritten)

	bytesWritten, err = bufferedWriter.WriteString("Buffered String\n")
	check(err)
	fmt.Printf("Bytes written: %d\n", bytesWritten)

	unflushedBufferSize := bufferedWriter.Buffered()
	fmt.Printf("Bytes Buffered: %d\n", unflushedBufferSize)

	bytesAvailable := bufferedWriter.Available()
	fmt.Printf("Available Buffer: %d\n", bytesAvailable)

	// write to disk
	bufferedWriter.Flush()
	// reset
	bufferedWriter.Reset(bufferedWriter)

	// default 4mb (4096byte) buffer size
	bytesAvailable = bufferedWriter.Available()
	fmt.Printf("Available Buffer: %d\n", bytesAvailable)

	// changing buffer size
	// min size is the default of 4mb
	bufferedWriter = bufio.NewWriterSize(bufferedWriter, 8000)

	return nil
}

func BufferedRead(filepath string) ([]string, error) {
	// ---------------------------
	// using NewReader w/bufio
	// ---------------------------
	
	data := []string{}

	file, err := os.Open(filepath)
	if err != nil {
		return data, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	line, counter := "", 0

	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		counter++

		fmt.Printf("Line %v > Read %d chars %v\n>",
			counter, len(line), string(line))
			data = append(data, line)
	}

	if err != io.EOF {
		return data, err
	}
	return data, nil
}

func BufferedRead2(filepath string) ([]string, error) {

	// ---------------------------
	// using NewScanner w/bufio
	// default delim \n
	// can split words with 
	// scanner.Split(bufio.ScanWords)
	// ---------------------------
	file, err := os.Open(filepath)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	/*
	lineCount := 0

	// scanner default delimiter '\n'
	// if error, skip line
	for scanner.Scan() {
		lineCount++
		err := scanner.Err()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Line %v > Read %d Char > %v\n",
			lineCount, len(scanner.Text()), string(scanner.Text()))
	}
   */

	// by lines
	/*
		for scanner.Scan() {
				line := scanner.Text()
				fmt.Println(line)
		}
	*/

	// by words
	test := []string{}
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		test = append(test, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return test, err
	}

	return test, nil
}

func DirectWrite(filepath string, test []byte) error {

	// ------------
	// using WriteFile w/ioutil
	// apparently depreciated
	// ------------
	err := ioutil.WriteFile(filepath, test, 0644)
	if err != nil {
		return err
	}
	return nil

}

func DirectReadAll(filepath string) (string, error) {

	// ------------
	// using ReadAll w/ioutil
	// ioutil depreciated
	// ------------
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "",err
	}

	// outputs []uint8
	/*
	fmt.Printf("%T: %v",data,data)
	fmt.Println("data as string:")
	fmt.Printf("%s", data)
    */
	
	return string(data),nil

}

func directWrite2(filepath string) error {

	// ------------
	// using Write & WriteString on file
	// ------------
	f, err := os.Create("testing/test2.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	// ASCII 32 - space, 10 - newline

	d2 := []byte{115, 111, 109, 101, 32}

	n2, err := f.Write(d2)
	if err != nil {
		return err
	}
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	if err != nil {
		return err
	}
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	return nil

}

func cpFile(fromFile, toFile string) error {

	//---------
	// copy from one file to another
	// using io.Copy
	//---------

	originalFile, err := os.Open(fromFile)
	if err != nil {
		return err
	}
	defer originalFile.Close()

	newFile, err := os.Create(toFile)
	if err != nil {
		return err
	}
	defer newFile.Close()

	//copy bytes to dest from src
	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		return err
	}
	fmt.Printf("Copied %d bytes\n", bytesWritten)

	//flush mem to disk (not sure if necessary)
	err = newFile.Sync()
	if err != nil {
		return err
	}

	return nil
}


func directReadFull(filepath string) error {

	// -------------------
	// using ReadFull w/io
	// read exactly n bytes
	// -------------------

	file, err := os.Open(filepath)
	if err != nil {
		return err
	}

	// read just two bytes
	b := make([]byte, 2)
	data, err := io.ReadFull(file, b)
	check(err)
	if err != nil {
		return err
	}

	fmt.Println(data)

	fmt.Println("data as string:")
	fmt.Printf("%s", data)

	return nil
}

func directReadAtLeast(filepath string) error {

	// -------------------
	// using ReadAtLeast w/io
	// read at least n bytes, then fill buffer
	// -------------------
	file, err := os.Open(filepath)
	check(err)
	if err != nil {
		return err
	}

	byteSlice, minBytes := make([]byte, 32), 8
	// read at least 8 bytes from file, else error
	// if there are at least 8 bytes, then read as much as byte slice can hold (32)
	numBytesRead, err := io.ReadAtLeast(file, byteSlice, minBytes)
	check(err)
	if err != nil {
		return err
	}

	fmt.Printf("Num of bytes read: %d\n", numBytesRead)
	fmt.Printf("Data read: %s\n", byteSlice)

	return nil
}

func ReadDiyBuffer(filepath string) {

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	const maxSz = 10

	b := make([]byte, maxSz)

	for {
		readTotal, err := file.Read(b)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fmt.Println(string(b[:readTotal]))
	}
}
