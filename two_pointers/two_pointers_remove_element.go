/*ТЗ  Ваша задача — удалить все вхождения числа element из массива, сохранив порядок остальных чисел.
Формат входных данных
В первой строке задаётся число  N  — размер массива (количество элементов).
Во второй строке вводится массив из  N  чисел, разделённых пробелами.
В третьей строке задаётся число element, которое нужно удалить из массива.
Формат выходных данных
На выходе программа должна вывести строку из элементов массива, из которого удалены все вхождения числа element. Оставшиеся числа должны быть записаны через пробел.
*/
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
	q := strings.Fields(line)
	line, _ = reader.ReadString('\n')
	element, _ := strconv.Atoi(strings.TrimSpace(line))
	qArr := make([]int, n)
	for i, v := range q {
		qArr[i], _ = strconv.Atoi(v)
	}
	var slow int
	for i, v := range qArr {
		if v != element {
			qArr[slow] = qArr[i]
			slow++
		}
	}
	for i := 0; i < slow; i++ {
		if i > 0 {
			writer.WriteByte(' ')
		}
		writer.WriteString(strconv.Itoa(qArr[i]))
	}
	writer.WriteByte('\n')
}
