// даны прямоугольники 9 шт 3х4. найти минимальную сторону квадрата, чтобы разместить их все
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

	line, _ := reader.ReadString('\n') // 1й эл-1я сторона, 2й элемент -2я сторона, 3й эл-количество дипломов
	q := strings.Fields(line)
	firstSize, _ := strconv.Atoi(q[0])
	secondSize, _ := strconv.Atoi(q[1])
	n, _ := strconv.Atoi(q[2])

	left := max(firstSize, secondSize)
	right := n * left
	target := n
	var res, answer, middle int

	for left <= right {
		middle = (right + left) / 2
		res = (middle / firstSize) * (middle / secondSize)
		if res >= target {

			answer = middle
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	writer.WriteString(strconv.Itoa(answer))
	writer.WriteByte('\n')
}
