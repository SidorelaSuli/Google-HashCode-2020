// simple funtion how to read Input Data Sets in Golang

package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

// readLines function reads a whole file into memory and returns a slice of its lines.
func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    // check for errors when closing a file
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

// writeLines function writes the lines to the given file.
func writeLines(lines []string, path string) error {
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()

    w := bufio.NewWriter(file)
    for _, line := range lines {
        fmt.Fprintln(w, line)
    }
    return w.Flush()
}

// main funtion
func main() {
    lines, err := readLines("a_example.in")  // a_example.in is the Data Set file to be read
    if err != nil {
        log.Fatalf("readLines: %s", err)
    }
    for _, line := range lines {   // USE for line := range lines {  -------------  for printing indexes
        fmt.Println(line)
    }

    if err := writeLines(lines, "a_example_output.in"); err != nil {   // a_example_output.in is the file where read lines will be written/saved
        log.Fatalf("writeLines: %s", err)
	}
}