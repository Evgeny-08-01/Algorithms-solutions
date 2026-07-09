//ТЗ  пересортировать оценки учеников. Все оценки, равные 0, должны быть перемещены в конец списка, при этом порядок остальных оценок должен остаться неизменным.

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
	qArr := make([]int, n)
	for i, v := range q {
		qArr[i], _ = strconv.Atoi(v)
	}
	var slow, fast int
	for i, v := range qArr {
		if v != 0 {
			fast = i
			qArr[slow], qArr[fast] = qArr[fast], qArr[slow]
			slow++
		}
	}
	for i, val := range qArr {
		if i > 0 {
			writer.WriteByte(' ')
		}
		writer.WriteString(strconv.Itoa(val))
	}
	writer.WriteByte('\n')
}
