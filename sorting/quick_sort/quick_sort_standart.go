/*Условие задачи
Название: Сортировка массива по возрастанию методом быстрой сортировки (Quick Sort).
Входные данные:
В первой строке задаётся число N (1 ≤ N ≤ 10⁵) — количество элементов.
Во второй строке — массив из N целых чисел, разделённых пробелами.
Выходные данные:
Выведите массив, отсортированный по возрастанию.
Пример:text 
Вход: 5  
3 1 4 1 5
Выход:
1 1 3 4 5*/
package main
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
arr:=initialisation()
left:=0
right:=len(arr)-1
quickSort(arr,left,right)	
fmt.Println(arr)
}

/////////////////////////////////////////////////////// ФУНКЦИИ/////////////////////////////////////////////////////////////////////////
func quickSort(arr []int, left, right int) {
    if left >= right {
        return
    }
    p := partition(arr, left, right)
    quickSort(arr, left, p)     
    quickSort(arr, p+1, right)   
}

func partition(arr []int, left, right int) int {
    pivot := arr[mediana(arr, left, right)]
    i := left - 1
    j := right + 1
    
    for {
        for i++; arr[i] < pivot; i++ {}
        for j--; arr[j] > pivot; j-- {}
        if i >= j {
            return j 
        }
        arr[i], arr[j] = arr[j], arr[i]
    }
}
// функция вычисления медианы
func mediana(arr []int,left,right int) int{
	mid:=left+(right-left)/2
	a,b,c:=arr[left],arr[mid],arr[right]
if a>b{
b,a=a,b
}
if b>c{
b,c=c,b
}
if a>b{
b,a=a,b
}
if b == arr[left] {
		return  left}
if b == arr[mid]  {	
	return   mid}	
if b == arr[right]{	
	return right}
	return mid	
}

// ============================================
// ВСПОМОГАТЕЛЬНАЯ ФУНКЦИЯ
// ============================================
// parseLine преобразует строку с числами в слайс целых чисел.
// Возвращает nil для пустых строк (чтобы их можно было пропустить).
func parseLine(line string) []int {
	// Удаляем пробелы и символы перевода строки по краям.
	line = strings.TrimSpace(line)
	// Если строка пустая — возвращаем nil.
	if line == "" {
		return nil
	}
	// Разбиваем строку на отдельные подстроки по пробелам.
	field := strings.Fields(line)
	// Создаем слайс нужного размера.
	row := make([]int, len(field))
	// Преобразуем каждую подстроку в число.
	for j, val := range field {
		temp, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal("convert error")
		}
		row[j] = temp
	}
	return row
}

func initialisation() []int{
// ============================================
	// БЛОК 1: ОТКРЫТИЕ ФАЙЛА
	// ============================================
	index:=1
	if len(os.Args)>1{
		indx,err:=strconv.Atoi(os.Args[1])
		if err!=nil||indx<0{
			log.Fatal("wrong input")	
		}
		index=indx	
	}
    file, err := os.Open("../../input.txt")
	if err != nil {
		log.Fatal("didn't open file")
	}
	defer file.Close()

	// ============================================
	// БЛОК 2: ЧТЕНИЕ ДАННЫХ
	// ============================================
	// Используем буферизированный ввод для эффективной работы с файлом.
	// Размер буфера 4096 байт — оптимально для небольших файлов.
	readerFile := bufio.NewReaderSize(file, 4096)

	// Буферизированный вывод: накапливает данные и отправляет их одной операцией ввода-вывода.
	// Это ускоряет работу при большом количестве fmt.Println.
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// matrix — хранилище для всех строк файла.
	// Каждая строка — отдельный слайс целых чисел.
	var matrix [][]int

	// Читаем файл построчно до EOF (конца файла).
	for {
		line, err := readerFile.ReadString('\n')
		// Если ошибка не EOF — критическая, завершаем программу.
		if err != nil && err != io.EOF {
			log.Fatal("reading error")
		}
		// Парсим строку: удаляем пробелы, разбиваем на числа.
		row := parseLine(line)
		// Пропускаем пустые строки (parseLine возвращает nil).
		if row != nil {
			matrix = append(matrix, row)
		}
		// Если достигнут конец файла — выходим из цикла.
		if err == io.EOF {
			break
		}
	}

	// ============================================
	// БЛОК 3: ПОДГОТОВКА ДАННЫХ
	// ============================================
	// Первая строка файла — всегда N (количество элементов).
	n := matrix[0][0]

	// Берем нужный массив из файла.
	// Индекс можно менять в зависимости от того, какой массив нужно сортировать.
	// matrix[1] — первый массив, matrix[2] — второй, matrix[3] — третий и т.д.
	arr := matrix[index]

	// Защита от некорректных данных:
	// 1. Если в строке больше чисел, чем N — обрезаем до N.
	if len(arr) > n {
		arr = arr[:n]
	}
	// 2. Если в строке меньше чисел, чем N — программа не может работать.
	if len(arr) < n {
		log.Fatal("not enough elements in array")
	}
	return arr
}

