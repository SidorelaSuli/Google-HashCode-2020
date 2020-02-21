package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// define stacktype
type stackType []int

// push to list
func (s *stackType) Push(v int) {
	*s = append(*s, v)
}

// pop from list
func (s *stackType) Pop() int {
	l := len(*s)
	ret := (*s)[l-1]
	*s = (*s)[:l-1]
	return ret
}

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

func saveAnswer(array []int, filename string) {
	// create file
	s := []string{filename, ".output"}
	saveFile := strings.Join(s, "")
	f, err := os.Create(saveFile)
	check(err)
	defer f.Close()

	t := strconv.Itoa(len(array)) + "\n"
	f.WriteString(t)
	for x := 0; x < len(array); x++ {
		v := strconv.Itoa(array[x]) + " "
		f.WriteString(v)
	}
	f.WriteString("\n")
	fmt.Println("File saved at", saveFile)
}

func sumUntilLimit(array []int, size int, target int) (stackType, int) {
	var sizes, bestSizes stackType
	var sum, temp, bestSum int

	sum = 0
	bestSum = 0
	for (len(sizes) > 0 && sizes[0] != 0) || len(sizes) == 0 {


		size = size - 1
		for index := size; index >= 0; index-- {
			temp = sum + array[index]
			if temp <= target {
				sizes.Push(index)
				sum = temp
				if sum == target {
					return sizes, sum
				}
			}
		}

		// save better solutions until now
		if sum > bestSum {
			bestSum = sum
			for _, element := range sizes {
				bestSizes.Push(element)
			}
		}

		if len(sizes) != 0 {
			last := sizes.Pop()
			sum = sum - array[last]
			size = last
		}
		
		if len(sizes) == 0 && size == 0 {
			break
		}
	}
	return bestSizes, bestSum
}

func main() {
	// log error when file is not passed
	if len(os.Args) < 2{
		fmt.Println("Specify the dataset file as argument")
		os.Exit(1)
	}

	// parse and calculate
	totalSlices, types, sizesPizza := parseDataset(os.Args[1])
	sizes, sum := sumUntilLimit(sizesPizza, types, totalSlices)

	// output general info
	fmt.Println("Max Slices:", totalSlices)
	fmt.Println("Sum:", sum)
	saveAnswer(sizes, os.Args[1])
}
