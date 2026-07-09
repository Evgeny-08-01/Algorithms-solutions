// тернарный поиск числа по отсортированному по возрастанию массиву
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err) // если файла нет — программа упадет
	}
	defer file.Close()
	consoleReader := bufio.NewReader(os.Stdin)
	reader := bufio.NewReaderSize(file, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()
	targetIn, _ := consoleReader.ReadString('\n')
	target, _ := strconv.Atoi(strings.TrimSpace(targetIn))
	line, _ := reader.ReadString('\n')
	nStr := strings.Fields(line)
	length := len(nStr)
	n := make([]int, length, length)
	for i, j := range nStr {
		n[i], _ = strconv.Atoi(j)
	}
	left := 0
	right := length - 1
	targetIndex := -1
	for left <= right {
		m1 := left + (right-left)/3
		m2 := right - (right-left)/3
		if n[m1] == target {
			targetIndex = m1
			break
		}
		if n[m2] == target {
			targetIndex = m2
			break
		}

		if target < n[m1] {
			right = m1 - 1
		}
		if target > n[m2] {
			left = m2 + 1
		} else {
			left = m1 + 1
			right = m2 - 1
		}
	}

	writer.WriteString(strconv.Itoa(targetIndex))
	writer.WriteString("\n")
}
