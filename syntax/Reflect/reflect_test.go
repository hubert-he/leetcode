package Reflect

import (
	"bytes"
	"io"
	"testing"
)
type MyWriter struct {
	Name string
}
func (MyWriter) Write([]byte)(int, error){
	return 0,nil
}
/* go test -v 打印 t.Log
 */
func TestNil(t *testing.T){
	var w io.Writer // nil interface
	t.Log("nil")
	t.Log("==判断：", w == nil)
	t.Log("异常判断：", InterfaceIsNil1(w))
	t.Log("类型判断：", InterfaceIsNil2(w))

	var b *bytes.Buffer
	w = b
	t.Log("ptr type not nil; value nil")
	t.Log("==判断：", w == nil)
	t.Log("异常判断：", InterfaceIsNil1(w))
	t.Log("类型判断：", InterfaceIsNil2(w))

	var my_writer MyWriter
	w = my_writer
	t.Log("struct type not nil; value not nil")
	t.Log("==判断：", w == nil)
	t.Log("异常判断：", InterfaceIsNil1(w))
	t.Log("类型判断：", InterfaceIsNil2(w))
}
