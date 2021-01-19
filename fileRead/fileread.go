package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// 方法1: 使用io/ioutil库的ReadFile()， 一次性读取
func ReadFile1(filePath string) []byte {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("Read Err", err)
	}
	return content
}

// 方法2：使用io/ioutil库的ReadAll()
func ReadAllTest(){
	//ioutil.ReadAll()
}
// bufio 包实现了带缓存的 I/O 操作, 封装一个 io.Reader 或 io.Writer 对象
/*
type Reader struct {
    		buf          []byte
    		rd           io.Reader // reader provided by the client
    		r, w         int       // buf read and write positions
    		err          error
    		lastByte     int
    		lastRuneSize int
}
NewReaderSize 将 io.Reader rd 封装成一个拥有 size 大小缓存的 bufio.Reader 对象
 如果 rd 的基类型就是 bufio.Reader 类型，而且拥有足够的缓存
 则直接将 rd 转换为基类型并返回
func NewReaderSize(rd io.Reader, size int) *Reader

// NewReader 相当于 NewReaderSize(rd, 4096)
func NewReader(rd io.Reader) *Reader
*/
// 方法3： 使用bufio库的ReadAll()

// 方法4： 使用bufio库的ReadLine()
func ReadFile4(filepath string, handle func([]byte))  {
	f, err := os.Open(filepath)
	defer f.Close()
	if err != nil {
		log.Println("Read Err", err)
	}
	buf := bufio.NewReader(f)
	for {
		line, isPrefix, err := buf.ReadLine()
		if err != nil {
			if err != io.EOF{
				log.Println("ReadLine Err", err)
			}  else {
				return
			}
		}
		if isPrefix == false {
			handle(line)
		}
	}
}

func print(str []byte) {
	fmt.Printf("==>%s\n", str)
}
/*
	当每次读取块的大小小于4KB，建议使用bufio.NewReader(f),
    大于5KB用 bufio.NewReaderSize(f, 缓存大小)
*/
func test_bufio(){
	/*
	  Peek返回缓存的一个切片，该切片引用缓存中前n字节数据
	  该操作不会将数据读出，只是引用
	  引用的数据在下一次读取操作之前是有效的，若引用的数据长度小于n，则返回一个错误信息
	  如果n大于缓存的总大小，则返回ErrBufferFull
	  通过Peek的返回值，可以修改缓存中的数据，但是不能修改底层io.Reader中的数据
	 */
	s := strings.NewReader("Abcdkaf")
	br := bufio.NewReader(s)
	b, _ := br.Peek(5)
	fmt.Printf("peek: %s\n", b)
	b[0] = '0'
	b, _ = br.Peek(5)
	fmt.Printf("%s\n",b)

	// func (b *Reader) Read(p []byte) (n int, err error)
	/*
	Read从b中读取数据到p，返回读出的字节数
	如果缓存不为空，则只能读出缓存中的数据，不会从底层io.Reader中提取数据
	如果缓存为空，则
		1. len(p) >= 缓存大小，则跳过缓存，直接从底层io.Reader中读入到p
		2. len(p) < 缓存大小，则先将数据从底层io.Reader中读取到缓存中，再从缓存读取到p
	 */
}
func main() {
	fname := "/Users/hezh/workspace/awesomeProject/fileRead/fileread.go"
	//c := ReadFile1(fname)
	//fmt.Printf("%s \n%d\n", c, len(c))
	ReadFile4(fname, print)
	test_bufio()
}
