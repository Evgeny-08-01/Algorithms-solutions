/*Дан отсортированный по возрастанию массив целых чисел и заданное число. Если заданное число уже находится в массиве, верните его индекс.
Если числа в массиве нет, верните индекс, где оно должно находиться, чтобы сохранить порядок сортировки.
Формат входных данных
В первой строке задаётся число N — количество элементов в массиве.
Во второй строке вводится массив из N целых чисел, разделённых пробелами.
В третьей строке задаётся число target, для которого нужно найти индекс.*/

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
	line, _ = reader.ReadString('\n')
	target, _ := strconv.Atoi(strings.TrimSpace(line))

	var qMiddle int
	left := 0
	right := n - 1
	middle := (left + right) / 2
	min, _ := strconv.Atoi(q[0])
	max, _ := strconv.Atoi(q[n-1])
	if target < min || target > max {
		writer.WriteString("out of range")
		writer.WriteByte('\n')
		return
	}

	for left <= right {
		qMiddle, _ = strconv.Atoi(q[middle])
		if qMiddle == target {
			break
		}
		if qMiddle > target {
			right = middle - 1
			middle = (right + left) / 2
		} else {
			left = middle + 1
			middle = (right + left) / 2
		}
	}
	foundTarget, _ := strconv.Atoi(q[middle])
	if foundTarget < target {
		middle++
	}

	writer.WriteString(strconv.Itoa(middle))
	writer.WriteByte('\n')
}
