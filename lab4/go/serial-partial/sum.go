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
func sum(filePath string, chuncks_g map[string][]string) (int, error) {
	data, err := readFile(filePath)
	if err != nil {
		return 0, err
	}

	_sum := 0
    var counter int = 0 
    var chunck string = ""
    var chuncks []string
	for _, b := range data {
        counter += 1
        if counter == 101 {
            counter = 0
            chuncks = append(chuncks, chunck)
            chunck = ""
        }
        chunck += string(b)
		_sum += int(b)
	}

    chuncks_g[filePath] = chuncks
	return _sum, nil
}

// print the totalSum for all files and the files with equal sum
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file1> <file2> ...")
		return
	}

	var totalSum int64
	sums := make(map[int][]string)
    chuncks_g := make(map[string][]string)
	for _, path := range os.Args[1:] {
		_sum, err := sum(path, chuncks_g)

		if err != nil {
			continue
		}

		totalSum += int64(_sum)

		sums[_sum] = append(sums[_sum], path)
	}

	fmt.Println(totalSum)

	for sum, files := range sums {
		if len(files) > 1 {
			fmt.Printf("Sum %d: %v\n", sum, files)
		}
	}

    fmt.Println(chuncks_g);
    //for i := 0; i < len(os.Args[1:); i++ {
            
    //}
}
