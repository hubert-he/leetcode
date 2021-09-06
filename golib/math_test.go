package golib

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
	"text/tabwriter"
	"time"
)
/*
导入：import "math/rand"
描述：rand包实现了伪随机数生成器，随机数从Source生成。包水平的函数都使用的默认的公共Source，该Source会在程序每次运行时都产生确定的序列。
     如果需要每次运行产生不同的序列，应使用Seed函数进行初始化。
     Default Source可以安全的用于多go程并发， 但是 如果Source是通过NewSource生成的，则不能用在mutiple goroutine
注意：此包不是安全相关的，如果需要安全相关的，查阅 crypto/rand package
类型：
Source：Source代表一个生成均匀分布在范围[0, 1<<63)的int64值的（伪随机的）资源。
	type Source interface { // 接口
		Int63() int64
		Seed(seed int64)
	}
生成器： 使用给定的种子创建一个伪随机资源。
	func NewSource(seed int64) Source

Rand: Rand生成服从多种分布的随机数
	type Rand struct {  // 结构体
    	// 内含隐藏或非导出字段
	}
生成器：返回一个使用src生产的随机数来生成其他各种分布的随机数值的*Rand
	func New(src Source) *Rand
方法：
1.  func (r *Rand) Seed(seed int64) 			// 使用给定的seed来初始化生成器到一个确定的状态。
2.  func (r *Rand) Int() int					// 返回一个非负的伪随机int值
3.  func (r *Rand) Int31() int32				// 返回一个int32类型的非负的31位伪随机数。
4.  func (r *Rand) Int63() int64				// 返回一个int64类型的非负的63位伪随机数。
5.  func (r *Rand) Uint32() uint32				// 返回一个uint32类型的非负的32位伪随机数。
6.  func (r *Rand) Intn(n int) int				// 返回一个取值范围在[0,n)的伪随机int值，如果n<=0会panic。
7.  func (r *Rand) Int31n(n int32) int32
8.  func (r *Rand) Int63n(n int64) int64
9.  func (r *Rand) Float32() float32			// 返回一个取值范围在[0.0, 1.0)的伪随机float32值
10. func (r *Rand) Float64() float64			// 返回一个取值范围在[0.0, 1.0)的伪随机float64值。
11. func (r *Rand) NormFloat64() float64		// 返回一个服从标准正态分布（标准差=1，期望=0）、取值范围在[-math.MaxFloat64, +math.MaxFloat64]的float64值.
												// 如果要生成不同的正态分布值，调用者可用如下代码调整输出：sample = NormFloat64() * 标准差 + 期望
12. func (r *Rand) ExpFloat64() float64			// 返回一个服从标准指数分布（率参数=1，率参数是期望的倒数）、取值范围在(0, +math.MaxFloat64]的float64值
												// 如要生成不同的指数分布值，调用者可用如下代码调整输出：sample = ExpFloat64() / 率参数
13. func (r *Rand) Perm(n int) []int			// 返回一个有n个元素的，[0,n)范围内整数的伪随机排列的切片。
14. func (r *Rand) Shuffle(n int, swap func(i, j int))
15. func (r *Rand) Read([]byte) (int, error)
Zipf: Zipf生成服从齐普夫分布的随机数。
	type Zipf struct {
		// 内含隐藏或非导出字段
	}
生成器：func NewZipf(r *Rand, s float64, v float64, imax uint64) *Zipf
方法：
func (z *Zipf) Uint64() uint64

Package 级别函数
1. func Seed(seed int64)		// 使用给定的seed将默认资源初始化到一个确定的状态；如未调用Seed，默认资源的行为就好像调用了Seed(1)
2. func Int() int				// 返回一个非负的伪随机int值
3. func Int31() int32			// 返回一个int32类型的非负的31位伪随机数。
4. func Int63() int64			// 返回一个int64类型的非负的63位伪随机数。
5. func Uint32() uint32			// 返回一个uint32类型的非负的32位伪随机数。
   func Uint64() uint64
6. func Intn(n int) int			// 返回一个取值范围在[0,n)的伪随机int值，如果n<=0会panic。
7. func Int31n(n int32) int32
8. func Int63n(n int64) int64
9. func Float32() float32
10. func Float64() float64
11. func NormFloat64() float64
12. func ExpFloat64() float64
13. func Perm(n int) []int
14. func Read(p []byte) (n int, err error) 		// 随机填充：从默认Source生成len(p)个随机字节，然后写入p数组中，并发安全的
15. func Shuffle(n int, swap func(i, j int))	// 随机打乱， n为元素的长度
*/

func TestRand(t *testing.T){
	// Package 级别
	rand.Seed(time.Now().UnixNano()) // 如果使用固定数字生成的Seed，每次Run所产生的序列是相同的
	t.Logf("Int: %X Int31: %X Int63: %X Uint32: %X Uint64: %X\n",
		rand.Int(), rand.Int31(), rand.Int63(), rand.Uint32(), rand.Uint64())
	t.Logf("Intn: %d Int31n: %d Int63n: %d ", rand.Intn(100), rand.Int31n(100), rand.Int63n(100))
	t.Logf("Float32: %f Float64: %f", rand.Float32(), rand.Float64())
	t.Logf("NormFloat64: %f %f %f", rand.NormFloat64(), rand.NormFloat64(), rand.NormFloat64())
	t.Logf("ExpFloat64: %f %f %f", rand.ExpFloat64(), rand.ExpFloat64(), rand.ExpFloat64())
	t.Logf("Perm: %v", rand.Perm(13))

	words := strings.Fields("ink runs from the corners of my mouth")
	rand.Shuffle(len(words), func(i, j int){
		words[i], words[j] = words[j], words[i]
	})
	t.Logf("words shuffered: %v", words)

	bs := make([]byte, 10)
	rand.Read(bs)
	t.Logf("byte read: %v", bs)

	// Rand 级别
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // Create and seed the generator.
	// The tabwriter here helps us generate aligned output.
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	defer w.Flush()
	show := func(name string, v1, v2, v3 interface{}) {
		fmt.Fprintf(w, "%s\t%v\t%v\t%v\n", name, v1, v2, v3)
	}
	// Float32 and Float64 values are in [0, 1).
	show("Float32", r.Float32(), r.Float32(), r.Float32())
	show("Float64", r.Float64(), r.Float64(), r.Float64())
	// ExpFloat64 values have an average of 1 but decay exponentially.
	show("ExpFloat64", r.ExpFloat64(), r.ExpFloat64(), r.ExpFloat64())
	// NormFloat64 values have an average of 0 and a standard deviation of 1.
	show("NormFloat64", r.NormFloat64(), r.NormFloat64(), r.NormFloat64())
	// Int31, Int63, and Uint32 generate values of the given width.
	// The Int method (not shown) is like either Int31 or Int63
	// depending on the size of 'int'.
	show("Int31", r.Int31(), r.Int31(), r.Int31())
	show("Int63", r.Int63(), r.Int63(), r.Int63())
	show("Uint32", r.Uint32(), r.Uint32(), r.Uint32())
	// Intn, Int31n, and Int63n limit their output to be < n.
	// They do so more carefully than using r.Int()%n.
	show("Intn(10)", r.Intn(10), r.Intn(10), r.Intn(10))
	show("Int31n(10)", r.Int31n(10), r.Int31n(10), r.Int31n(10))
	show("Int63n(10)", r.Int63n(10), r.Int63n(10), r.Int63n(10))

	// Perm generates a random permutation of the numbers [0, n).
	show("Perm", r.Perm(5), r.Perm(5), r.Perm(5))
}

/*
%v 值的默认格式。
%+v   类似%v，但输出结构体时会添加字段名
%#v　 相应值的Go语法表示
%T 相应值的类型的Go语法表示
%% 百分号,字面上的%,非占位符含义

%b 二进制表示
%c 相应Unicode码点所表示的字符
%d 十进制表示
%o 八进制表示
%q 单引号围绕的字符字面值，由Go语法安全地转义
%x 十六进制表示，字母形式为小写 a-f
%X 十六进制表示，字母形式为大写 A-F
%U Unicode格式：123，等同于 "U+007B"

%b 无小数部分、二进制指数的科学计数法，如-123456p-78； 参见strconv.FormatFloat %e 科学计数法，如-1234.456e+78 %E 科学计数法，如-1234.456E+78 %f 有小数部分但无指数部分，如123.456 %F 等价于%f %g 根据实际情况采用%e或%f格式（以获得更简洁、准确的输出） %e 科学计数法，例如 -1234.456e+78
%E 科学计数法，例如 -1234.456E+78
%f 有小数点而无指数，例如 123.456
%g 根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出
%G 根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出

%t   true 或 false

%s 字符串或切片的无解译字节
%q 双引号围绕的字符串，由Go语法安全地转义
%x 十六进制，小写字母，每字节两个字符
%X 十六进制，大写字母，每字节两个字符

%p 十六进制表示，前缀 0x

+ 总打印数值的正负号；对于%q（%+q）保证只输出ASCII编码的字符。
- 左对齐
# 备用格式：为八进制添加前导 0（%#o），为十六进制添加前导 0x（%#x）或0X（%#X），为 %p（%#p）去掉前导 0x；对于 %q，若 strconv.CanBackquote 返回 true，就会打印原始（即反引号围绕的）字符串；如果是可打印字符，%U（%#U）会写出该字符的Unicode编码形式（如字符 x 会被打印成 U+0078 'x'）。
' '  （空格）为数值中省略的正负号留出空白（% d）；以十六进制（% x, % X）打印字符串或切片时，在字节之间用空格隔开
0 填充前导的0而非空格；对于数字，这会将填充移到正负号之后
*/