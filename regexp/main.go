package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	src := "Hello World!"
	match, _ := regexp.MatchString("H(.*)d!", src) // 优先匹配更多
	fmt.Println(match)

	match, _ = regexp.Match("H(.*)d!", []byte(src))
	fmt.Println(match)

	// 通过Compile来使用优化过得正则对象
	r, _ := regexp.Compile("H(.*)d!")
	fmt.Println(r.MatchString(src))

	// 返回匹配的子串
	fmt.Println(r.FindString(src))
	fmt.Println(string(r.Find([]byte(src))))

	// 这个方法查找首次匹配的索引，即起始索引和结束索引(前闭后开集合)
	src2 := "Hello World! world"
	index := r.FindStringIndex(src2)
	fmt.Println(index)
	fmt.Println(src2[index[0]:index[1]])

	// 返回全局匹配的字符串和局部匹配的字符串，匹配最大的字符串一次。
	// 与 r.FindAllString(src2, 1) 等价，返回匹配(.*)的子串
	sub := r.FindStringSubmatch(src2)
	fmt.Println(sub)
	fmt.Println(r.FindStringSubmatchIndex(src2))

	// 返回所有正则匹配的字符，不仅仅是第一个
	fmt.Println(r.FindAllString("Hello World! Held! world", -1))
	fmt.Println(r.FindAllString("Hello World! Held! world", 2))

	src3 := "Hello World! Held! Helloworld! world"
	r2, _ := regexp.Compile("H([a-z]+)d!")
	fmt.Println(r2.FindAllStringSubmatchIndex(src3, -1))
	fmt.Println(r2.FindAllStringSubmatch(src3, -1))
	// 指定正整数来限制匹配数量
	fmt.Println(r2.FindAllStringSubmatchIndex(src3, 1))

	re := regexp.MustCompile(`a.`)
	// n 表示查找前n个匹配项，若 n < 0 表示可查找所有匹配项
	fmt.Println(re.FindAllString("paranormal", -1))
	fmt.Println(re.FindAllIndex([]byte("paranormal"), -1))
	fmt.Println(re.FindAllString("paranormal", 2))
	fmt.Println(re.FindAllString("graal", -1))
	fmt.Println(re.FindAllString("none", -1))

	fmt.Println(r.ReplaceAllString(src2, "html"))
	in := []byte(src2)
	// Func变量可以将所有匹配的字符串都经过该函数处理转变为所需要的值
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

	b := bytes.NewReader([]byte("Hello World!"))
	reg := regexp.MustCompile(`\w+`)
	fmt.Println(reg.FindReaderIndex(b))

	// 将 template 的内容经过处理后，追加到 dst 的尾部。
	// template 中要有 $1、$2、${name1}、${name2} 这样的“分组引用符”
	// match 是由 FindSubmatchIndex 方法返回的结果，里面存放了各个分组的位置信息
	// 如果 template 中有“分组引用符”，则以 match 为标准，
	// 在 src 中取出相应的子串，替换掉 template 中的 $1、$2 等引用符号。
	reg = regexp.MustCompile(`(\w+),(\w+)`)
	bsrc := []byte("Golang,World!")
	dst := []byte("Say: ")
	template := []byte("Hello $1, Hello $2") // 模板
	m := reg.FindSubmatchIndex(bsrc) // 解析源文本
	fmt.Printf("%q", reg.Expand(dst, template, bsrc, m)) // 填写模板，并将模板追加到目标文本中

	reg = regexp.MustCompile(`Hello[\w\s]+`)
	fmt.Println(reg.LiteralPrefix())

	reg = regexp.MustCompile(`Hello`)
	fmt.Println(reg.LiteralPrefix())

	text := `Hello World! hello world`
	// 正则表达式非贪婪模式(?U)
	reg = regexp.MustCompile(`(?U)H[\w\s]+o`)
	fmt.Printf("%q\n", reg.FindString(text))
	reg.Longest() // 切换到贪婪模式
	fmt.Printf("%q\n", reg.FindString(text))

	//统计regexp中的分组个数 不包含非捕获的分组
	fmt.Println(r.NumSubexp())
	//返回r中的regexp字符串
	fmt.Printf("regexp=> %s\n", r.String())

	// 在字符串中搜索匹配项，并以匹配项为分隔符，将字符串分割成多个子串
	// 最多分割出n个子串,第n个子串不再分割, 若n<0 分割所有子串
	// 返回分割后的子串slice
	fmt.Printf("%q\n", r.Split(src3,-1))

	// 搜索匹配项，并替换为repl指定的内容
	// 如果rep中有分组引用符($1 $name)，则将分组引用符当普通字符处理
	// 返回全部替换后的结果字符串
	s := "Hello World, hello!"
	reg = regexp.MustCompile(`(Hell|h)o`)
	rep := "${1}"
	fmt.Printf("%q\n", reg.ReplaceAllLiteralString(s, rep))

	ss := []byte("Hello World!")
	// MustConplie 与 Complile 区别是  没有一个返回值err
	reg = regexp.MustCompile("(H)ello")
	repb := []byte("$0$1") // 没有Literal的，分组引用符号会起作用
	fmt.Printf("%s\n", reg.ReplaceAll(ss, repb))

	fmt.Printf("%s\n", reg.ReplaceAllFunc(ss,
		func(b []byte) []byte {
			rst := []byte{}
			rst = append(rst, b...)
			rst = append(rst, "$1"...)
			return rst
		}))


}
