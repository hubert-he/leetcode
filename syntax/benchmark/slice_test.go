package benchmark
import (
	"crypto/rand"
	"testing"
)
// dst作为全局变量是防止编译器优化for-loop
var (
	src = make([]byte, 512)
	dst = make([]byte, 512)
)

func genSource() {
	rand.Read(src)
}
/* 测试 copy 与 append
命令：uptime;go version;go test -bench=. ./
环境：go version go1.15.2 darwin/amd64
结果：
go version go1.15.2 darwin/amd64
goos: darwin
goarch: amd64
BenchmarkCopy-8          7401756               145 ns/op
BenchmarkAppend-8         781444              2410 ns/op
 */
func BenchmarkCopy(b *testing.B){
	for n := 0; n < b.N; n++{
		b.StopTimer()
		genSource()
		b.StartTimer()
		copy(dst, src)
	}
}

func BenchmarkAppend(b *testing.B){
	for n := 0; n < b.N; n++{
		b.StopTimer()
		genSource()
		b.StartTimer()
		dst = append(dst, src...)
	}
}

/* for range benchmark
for _, x := range X  产生x 的拷贝
结果：
goos: darwin
goarch: amd64
BenchmarkCopy-8                  8235602               148 ns/op
BenchmarkAppend-8                 712618              1728 ns/op
BenchmarkRangeIndex-8               8097            145191 ns/op
BenchmarkRangeValue-8                 43          31341447 ns/op
BenchmarkFor-8                      7645            138039 ns/op
 */
var X [1<<15]struct{
	val		int
	_		[4096]byte
}
var Result int
func BenchmarkRangeIndex(b *testing.B){
	r := 0
	for n := 0; n < b.N; n++{
		for i := range X{
			x := &X[i]
			r += x.val
		}
	}
	Result = r  // 防止编译器优化
}
func BenchmarkRangeValue(b *testing.B){
	r := 0
	for n := 0; n < b.N; n++{
		for _, x := range X{
			r += x.val
		}
	}
	Result = r
}
func BenchmarkFor(b *testing.B){
	r := 0
	for n := 0; n < b.N; n++{
		for i := 0; i < len(X); i++{
			x := &X[i]
			r += x.val
		}
	}
	Result = r
}