// Бинарный поиск целочисленного квадратного корня
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	target := n
	left := 0
	right := n
	middle := (right + left) / 2
	for left <= right {
		if middle*middle == target {
			break
		}

		if middle*middle > target {
			right = middle - 1
			middle = (right + left) / 2
		} else {
			left = middle + 1
			middle = (right + left) / 2
		}
	}
	writer.WriteString(strconv.Itoa(middle))
	writer.WriteByte('\n')
}
