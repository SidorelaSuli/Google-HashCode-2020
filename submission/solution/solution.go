package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		fmt.Println("Error with file",e)
		os.Exit(1)
	}
}

func parseDataset(filePath string) (int, int, []int) {

	// open fd
	fileDescriptor, err := os.Open(filePath)
	check(err)
	defer fileDescriptor.Close()

	// create buffer for read words
	scanner := bufio.NewScanner(fileDescriptor)
	scanner.Split(bufio.ScanWords)

	// read file
	var numbers []int
	for scanner.Scan() {
		item, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, item)
	}
	return numbers[0], numbers[1], numbers[2:]
}


func main() {
	// log error when file is not passed
	if len(os.Args) < 2{
		fmt.Println("Specify the dataset file as argument")
		os.Exit(1)
	}

	// parse
	totalSlices := parseDataset(os.Args[1])

	// output general info
	fmt.Println("Max Slices:", totalSlices)
}