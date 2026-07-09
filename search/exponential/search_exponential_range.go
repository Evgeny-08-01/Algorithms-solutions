/*Дан отсортированный по возрастанию массив целых чисел и некоторое целое число. 
Ваша задача — найти минимальный диапазон индексов (подмассив), в котором может находиться заданное число. 
Используйте метод экспоненциального поиска.
Формат входных данных
В первой строке задаётся число N — количество элементов в массиве.
Во второй строке вводится массив из N целых чисел, разделённых пробелами.
В третьей строке задаётся целое число target, для которого нужно найти диапазон.
Формат выходных данных
Программа должна вывести два числа через пробел, представляющие диапазон индексов, 
где может находиться число target. Если число находится в массиве, это должен быть диапазон, содержащий индекс числа. 
Если число отсутствует, диапазон должен содержать индексы, где оно могло бы быть вставлено.*/
package main
import
(
"bufio"
"strconv"
"strings"
"os"
//"fmt"
)
func main(){
 reader:=bufio.NewReaderSize(os.Stdin,1<<20)
 writer:=bufio.NewWriterSize(os.Stdout,1<<20)
 defer writer.Flush()
lengthStr,_:=reader.ReadString('\n')
line,_:=reader.ReadString('\n')
targetStr,_:=reader.ReadString('\n')
nStr:=strings.Fields(line)
//переводим все в int
length, _ := strconv.Atoi(strings.TrimSpace(lengthStr))
target, _ := strconv.Atoi(strings.TrimSpace(targetStr))
var  leftIndex, rightIndex int
n:=make([]int,length)
for i,j:=range nStr{
	n[i],_=strconv.Atoi(j)
}
// сначала граничные условия-1. таргет вне диапазона 2. длина массива==0,1
if length==0{
 leftIndex=-1
rightIndex=-1
	goto outToConsol
}
if target<n[0]{
 leftIndex=-1
rightIndex=-1
	goto outToConsol
}
if target==n[0]{
 leftIndex=0
rightIndex=1
	goto outToConsol
}
if target>n[length-1]{
 leftIndex=-1
rightIndex=-1
	goto outToConsol
}
if target==n[length-1]{
 leftIndex=length-2
rightIndex=length-1
	goto outToConsol
}
//експотренциальный алгоритм:
leftIndex=0
rightIndex=length-1


for rightIndex-leftIndex>1{   
temp:= leftIndex         
border:=1
 for n[temp+border]<=target{

border=border*2	
if temp+border>=length-1{rightIndex=length-1;break}
	if  n[temp+border]==target{	
border=border/2	
break
	}
 }
leftIndex=temp+border/2
rightIndex=temp+border
if rightIndex>=length-1{rightIndex=length-1;break}
}

if n[leftIndex]==target{rightIndex=leftIndex+1;leftIndex=leftIndex-1;goto outToConsol}
if n[leftIndex+1]==target{goto outToConsol}
if n[rightIndex]==target{leftIndex=rightIndex-1;rightIndex=rightIndex+1;goto outToConsol}
if rightIndex-leftIndex==2{
	if n[leftIndex+1]>target{rightIndex=leftIndex+1
	                   }else{leftIndex=rightIndex-1}
	
}
outToConsol:
writer.WriteString(strconv.Itoa(leftIndex))
writer.WriteString(" ")
writer.WriteString(strconv.Itoa(rightIndex))
}