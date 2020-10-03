package main

import (
	"GoStudyNote/sorter/algorithms/bubble"
	"GoStudyNote/sorter/algorithms/qsort"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

//Go语言标准库提供了用于快速解析命令行参数的flag包
var infile *string = flag.String("i", "unsorted.txt", "File contains values for sorting")
var outfile *string = flag.String("o", "sorted.txt", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")


//读取文件，设计输入文件的格式，输入文件是一个纯文本，每一行是一个需要被排序的数字
func readValues(infile string)(values []int, error error) {				//采用同名返回值

	file, error := os.Open(infile)										//注意此处的返回值
	if error != nil {
		fmt.Println("Failed to open the input file", infile)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)

	for true {
		line, isPrefix, err := br.ReadLine()							//line是字节数组

		if err != nil {
			if err != io.EOF {
				error = err
			}
			break
		}

		if isPrefix {													//判断该行数据是不是太长
			fmt.Println("A too long line, seems unexpected.")
			return
		}

		str := string(line)												//转换字符数组为字符串

		value, err := strconv.Atoi(str)
		if err != nil {
			error = err
			return
		}

		values = append(values, value)

	}

	return

}

//写文件,将排序后的结果写入文件
func writeValues(values []int, outfile string) error {

	file, err := os.Create(outfile)

	if err != nil {
		fmt.Println("Failed to create the output file", outfile)
		return err
	}

	defer file.Close()

	for _, value := range values{
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}

	return nil

}

func main() {

	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}

	values, err := readValues(*infile)
	if err == nil {
		t1 := time.Now()

		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubble.BubbleSort(values)
		default:
			fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsupported.")
		}

		t2 := time.Now()

		fmt.Println("The sorting process costs", t2.Sub(t1), "to complete.")

		writeValues(values, *outfile)

	}else{
		fmt.Println(err)
	}



}