package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	var delimiter = flag.String("d", ";", "delimiter")
	fmt.Println(*delimiter)
	path := os.Args[1]
	flag.Parse()
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	colMaxLen := make([]int, len(strings.Split(lines[0], *delimiter)))

	var table [][]string
	for _, line := range lines {
		cols := strings.Split(line, *delimiter)
		for i, col := range cols {
			if colMaxLen[i] < utf8.RuneCountInString(col) {
				colMaxLen[i] = utf8.RuneCountInString(col)
			}
		}
		table = append(table, cols)
	}

	printCSV(path, table, colMaxLen)

}

func printCSV(filename string, table [][]string, colMaxLen []int) {
	widthNums := 1
	widthCols := sum(colMaxLen)
	colCount := len(colMaxLen)
	widthTotal := (widthNums + 2) + widthCols + (colCount * 2)
	fmt.Printf("┏%s┓\n", strings.Repeat("━", widthTotal+colCount))
	fmt.Printf("┃ %-*s ┃\n", widthTotal+3, filename)
	for i, row := range table {
		fmt.Printf("┣━%s━", strings.Repeat("━", widthNums))
		if i == 0 {
			for j := range row {
				fmt.Printf("┳%s", strings.Repeat("━", colMaxLen[j]+2))
			}
		} else {
			for j := range row {
				fmt.Printf("╋%s", strings.Repeat("━", colMaxLen[j]+2))
			}
		}
		fmt.Printf("┫\n")
		fmt.Printf("┃ %v ", i+1)
		for j, col := range row {
			fmt.Printf("┃ %-*s ", colMaxLen[j], col)
		}
		fmt.Printf("┃\n")
	}
	fmt.Printf("┗━%s━", strings.Repeat("━", widthNums))
	for j := range table[0] {
		fmt.Printf("┻%s", strings.Repeat("━", colMaxLen[j]+2))
	}
	fmt.Printf("┛\n")
}

func sum(nums []int) int {
	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}
