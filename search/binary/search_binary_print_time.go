// легкая задача. 2 ксерокса со скоростями x,y. можно использовать оба, можно ксерить копии. найти минимальное время для создания n копий.
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

	line, _ := reader.ReadString('\n') // 1й элемент-n копий, 2й элемент скорость X копий в минуту, 3й элемент-скорость Y копий в минуту
	q := strings.Fields(line)
	n, _ := strconv.Atoi(q[0])
	x, _ := strconv.Atoi(q[1])
	y, _ := strconv.Atoi(q[2])

	maX := min(x, y) // максимальная скорость из двух ксероксов
	miN := max(x, y) // минимальная скорость из двух ксероксов

	left := 0              // минимальное время за которое будут распечатаны n  копий и равно, соответственно 0, при n=1-ничего печатать не надо-она уже есть
	right := miN * (n - 1) // максимальное время за которое будут распечатаны n  копий и равно, соответственно 0, при n=1-ничего печатать не надо-она уже есть
	var middle int         // middle-среднее время за которое распечатается n-1 копий
	for left+1 < right {
		middle = (left + right) / 2
		if (middle/x + middle/y) < n-1 {
			left = middle
		} else {
			right = middle
		}
	}
	t := left + 1*maX
	if n == 1 {
		t = 0
	}
	writer.WriteString(strconv.Itoa(t))
	writer.WriteByte('\n')
}
