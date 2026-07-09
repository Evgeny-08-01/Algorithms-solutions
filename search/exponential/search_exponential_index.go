// експотенциальный поиск числа по отсортированному по возрастанию массиву
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
	// сначала в мэйн бинарный поиск и вызов функции для определения отрезка под бинарный поиск
	middle, left, right := 0, 0, 0
	index := -1
	if n[0] > target || length == 0 {
		index = -1
		goto label
	}

	left, right = expSearch(n, length, target)
	if left == -1 && right == -1 {
		goto label
	}
	// бинарный поиск:

	for left <= right {
		middle = (left + right) / 2
		if n[middle] == target {
			index = middle
			break
		}
		if n[right] == target {
			index = right
			break
		}
		if n[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
label:
	writer.WriteString(strconv.Itoa(index))
	writer.WriteString("\n")
}

func expSearch(n []int, length, target int) (int, int) {
	if length <= 1 {
		return 0, 0
	}
	border := 1
	right := length - 1
	if target > n[right] {
		return -1, -1
	}
	for n[border] <= target {
		if n[border] == target {
			return border, border
		}
		border = border * 2
		if border >= right {
			border = right
			break
		}
	}
	borderRight := border
	if right < border {
		borderRight = right
	}
	return border / 2, borderRight
}
