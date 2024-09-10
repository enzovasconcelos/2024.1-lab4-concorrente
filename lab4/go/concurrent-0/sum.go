package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// read a file from a filepath and return a slice of bytes
func readFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v", filePath, err)
		return nil, err
	}
	return data, nil
}

// sum all bytes of a file
func sum(filePath string, chanel chan Tupla) {
	data, _ := readFile(filePath)
	//if err != nil {
	//}

	_sum := 0
	for _, b := range data {
		_sum += int(b)
	}
    tipo := Tupla { filePath, _sum }
    chanel <- tipo
}

type Tupla struct {
    path string
    sum int 
}

// print the totalSum for all files and the files with equal sum
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file1> <file2> ...")
		return
	}

	var totalSum int64
	sums := make(map[int][]string)
    chanel := make(chan Tupla, len(os.Args) - 2)
    for _, path := range os.Args[1:] {
		go sum(path, chanel)
	}

    for i := 0; i < len(os.Args[1:]); i++ {
        readValue := <-chanel
		totalSum += int64(readValue.sum)
		sums[readValue.sum] = append(sums[readValue.sum], readValue.path)
    }
    
	fmt.Println(totalSum)

	for sum, files := range sums {
		if len(files) > 1 {
			fmt.Printf("Sum %d: %v\n", sum, files)
		}
	}
}
