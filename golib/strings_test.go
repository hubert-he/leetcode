package golib

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"unicode"
)

// Package strings implements simple functions to manipulate UTF-8 encoded strings.

/* 包级别处理函数
functions:
1. func Compare(a, b string) int  // a>b 返回1   a<b 返回-1   a==b 返回 0
5. func Count(s, substr string) int // Count counts the number of non-overlapping instances of substr in s.
6. func EqualFold(s, t string) bool // 在utf-8编码角度下，大小写不敏感的相等比较
22. func Repeat(s string, count int) string // Repeat returns a new string consisting of count copies of the string s.

 */
func TestFunctions(t *testing.T){
	// 字符串比较函数Compare
	t.Logf("ab vs bb = %d", strings.Compare("ab", "bb"))
	t.Logf("aa vs aa = %d", strings.Compare("aa", "aa"))
	t.Logf("bb vs ab = %d", strings.Compare("bb", "ab"))

	// Count 返回不重叠的子串个数, 如果子串是空串的话，返回总串个数+1
	t.Logf("Count(%s,%s)=%d", "cheese", "e", strings.Count("cheese", "e"))
	// 注意 是不重叠的计算， sts sts 之间有重叠，所以没有计算为2
	t.Logf("Count(%s,%s)=%d", "chststs", "sts", strings.Count("chststs", "sts"))
	// 空串 返回 待插入的位置 +1
	t.Logf("Count(%s,\"\")=%d", "chststs", strings.Count("chststs", ""))

	t.Logf("EqualFold(%s, %s)=%t", "go", "Go", strings.EqualFold("Go", "go"))
	t.Logf("EqualFold(%s, %s)=%t", "Goo", "Go", strings.EqualFold("Goo", "go"))

	// panic 如果count 是负数 或者 len(s) *count 整型溢出
	t.Logf("Repeat(\"na\", 2)=%s", strings.Repeat("na", 2))
}

/* 子串包含类
2. func Contains(s, substr string) bool  // Contains reports whether substr is within s.
3. func ContainsAny(s, chars string) bool // ContainsAny reports whether any Unicode code points in chars are within s
4. func ContainsRune(s string, r rune) bool // ContainsRune reports whether the Unicode code point r is within s
 */
func TestContains(t *testing.T){
	s := "seafood"
	//注意： 空串 返回 true
	t.Logf("Contains(%s, \"foo\") = %t", s, strings.Contains(s, "foo"))
	t.Logf("Contains(%s, \"bar\") = %t", s, strings.Contains(s, "bar"))
	t.Logf("Contains(%s, \"\") = %t", s, strings.Contains(s, ""))
	t.Logf("Contains(\"\", \"\") = %t", strings.Contains("", ""))

	t.Logf("ContainsAny(%s, %s)=%t", "team", "i", strings.ContainsAny("team", "i"))
	t.Logf("ContainsAny(%s, %s)=%t", "ure", "ui", strings.ContainsAny("ure", "ui"))
	t.Logf("ContainsAny(\"foo\", \"\")=%t", strings.ContainsAny("foo", ""))
	t.Logf("ContainsAny(\"\", \"\")=%t", strings.ContainsAny("", ""))

	t.Logf("ContainsRune(\"aardvark\", 97)=%t", strings.ContainsRune("aardvark", 97))
	t.Logf("ContainsRune(\"timeout\", 97)=%t", strings.ContainsRune("timeout", 97))
}

/* 前后缀系列
9. func HasPrefix(s, prefix string) bool 	// HasPrefix tests whether the string s begins with prefix.
10. func HasSuffix(s, suffix string) bool 	// HasSuffix tests whether the string s ends with suffix.
 */

func TestPreSuffix(t *testing.T){
	t.Logf("HasPrefix(\"Gopher\", \"Go\") = %t", strings.HasPrefix("Gopher", "Go"))
	t.Logf("HasPrefix(\"Gopher\", \"go\") = %t", strings.HasPrefix("Gopher", "go"))
	t.Logf("HasSuffix(\"Amigo\", \"go\") = %t", strings.HasSuffix("Amigo", "go"))
	t.Logf("HasSuffix(\"Amigo\", \"Go\") = %t", strings.HasSuffix("Amigo", "Go"))
	// 注意：空字符串 返回true 是任何字符串的 前后缀
	t.Logf("HasPrefix(\"Gopher\", \"\") = %t", strings.HasPrefix("Gopher", ""))
	t.Logf("HasSuffix(\"Amigo\", \"\") = %t", strings.HasSuffix("Amigo", ""))
}

/*
11. func Index(s, substr string) int // 返回首次出现的下标位置(从0计)，无则返回-1
12. func IndexAny(s, chars string) int
13. func IndexByte(s string, c byte) int
14. func IndexFunc(s string, f func(rune) bool) int
15. func IndexRune(s string, r rune) int
17. func LastIndex(s, substr string) int
18. func LastIndexAny(s, chars string) int
19. func LastIndexByte(s string, c byte) int
20. func LastIndexFunc(s string, f func(rune) bool) int
 */
func TestIndex(t *testing.T){
	t.Logf("Index(\"go gopher\", \"go\")=%d", strings.Index("go gopher", "go"))
	t.Logf("LastIndex(\"go gopher\", \"go\")=%d", strings.LastIndex("go gopher", "go"))
	t.Logf("Index(\"go gopher\", \"rodent\")=%d", strings.Index("go gopher", "rodent"))
	t.Logf("LastIndex(\"go gopher\", \"rodent\")=%d", strings.LastIndex("go gopher", "rodent"))

	/* IndexAny returns the index of the first instance of any Unicode code point from chars in s,
		or -1 if no Unicode code point from chars is present in s*/
	t.Logf("IndexAny(\"go gopher\", \"go\")=%d", strings.IndexAny("go gopher", "go"))
	t.Logf("LastIndexAny(\"go gopher\", \"go\")=%d", strings.LastIndexAny("go gopher", "go"))
	t.Logf("IndexAny(\"abc\", \"def\")=%d", strings.IndexAny("abc", "def"))

	/* IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s. */
	t.Logf("IndexByte(\"golang\", 'g')=%d", strings.IndexByte("golang", 'g'))
	t.Logf("LastIndexByte(\"golang\", 'g')=%d", strings.LastIndexByte("golang", 'g'))
	t.Logf("IndexByte(\"golang\", 'a')=%d", strings.IndexByte("golang", 'a'))
	t.Logf("IndexByte(\"golang\", 'x')=%d", strings.IndexByte("golang", 'x'))

	/* IndexFunc returns the index into s of the first Unicode code point satisfying f(c), or -1 if none do. */
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	t.Logf("IndexFunc(\"Hello, 世界\", f)=%d", strings.IndexFunc("Hello, 世界", f))
	t.Logf("LastIndexFunc(\"Hello, 世界\", f)=%d", strings.LastIndexFunc("Hello, 世界", f))
	t.Logf("IndexFunc(\"Hello, world\", f)=%d", strings.IndexFunc("Hello, world", f))
	t.Logf("LastIndexFunc(\"Hello, world\", unicode.IsNumber)=%d", strings.LastIndexFunc("Hello, world", unicode.IsNumber))

	/* IndexRune returns the index of the first instance of the Unicode code point r,
	   or -1 if rune is not present in s. If r is utf8.RuneError,
	   it returns the first instance of any invalid UTF-8 byte sequence.
	 */
	t.Logf("IndexRune(\"chicken\", 'k')=%d", strings.IndexRune("chicken", 'k'))
	t.Logf("IndexRune(\"chicken\", 'd')=%d", strings.IndexRune("chicken", 'd'))
}
/* 替换类
21. func Map(mapping func(rune) rune, s string) string
23. func Replace(s, old, new string, n int) string
24. func ReplaceAll(s, old, new string) string
 */
func TestReplace(t *testing.T){
	s := "'Twas brillig and the slithy gopher..."
	rot13 := func(r rune) rune{
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r - 'A' + 13) % 26
		case r >= 'a' && r <= 'z':
			return 'a' + (r - 'a' + 13) % 26
		}
		return r
	}
	t.Logf("Map(f, %s)=%s", s, strings.Map(rot13, s))

	/* Replace returns a copy of the string s with the first n non-overlapping instances of old replaced by new.
	** If old is empty, it matches at the beginning of the string and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune string
	** If n < 0, there is no limit on the number of replacements
	 */
	s = "oink oink oink"
	t.Logf("Replace(%s, \"k\", \"ky\", 2)=%s", s, strings.Replace(s, "k", "ky", 2))
	t.Logf("Replace(%s, \"\", \"mx\", 1)=%s", "t", strings.Replace("t", "", "mx", 1))
	t.Logf("Replace(%s, \"\", \"mx\", 2)=%s", "t", strings.Replace("t", "", "mx", 2))
	t.Logf("Replace(%s, \"\", \"mx\", 3)=%s", "t", strings.Replace("t", "", "mx", 3))
	t.Logf("Replace(%s, \"\", \"mx\", -1)=%s", "oink", strings.Replace("oink", "", "mx", -1))
	t.Logf("Replace(%s, \"\", \"mx\", 1)=%s", "oink", strings.Replace("oink", "", "mx", 1))
	t.Logf("Replace(%s, \"\", \"mx\", -1)=%s", s, strings.Replace(s, "", "mx", -1))
	t.Logf("Replace(%s, \"oink\", \"moo\", -1)=%s", s, strings.Replace(s, "oink", "moo", -1))
	t.Logf("ReplaceAll(%s, \"oink\", \"moo\")=%s", s, strings.ReplaceAll(s, "oink", "moo"))

}
/*
7. func Fields(s string) []string
8. func FieldsFunc(s string, f func(rune) bool) []string
16. func Join(elems []string, sep string) string
25. func Split(s, sep string) []string
26. func SplitAfter(s, sep string) []string
27. func SplitAfterN(s, sep string, n int) []string
28. func SplitN(s, sep string, n int) []string
 */
func TestSplit(t *testing.T){
/* Fields splits the string s around each instance of one or more consecutive(连续的) white space characters,
	as defined by unicode.IsSpace, returning a slice of substrings of s or an empty slice if s contains only white space.
*/
	t.Logf("Fields(\"  foo bar\tbaz   \")=%q", strings.Fields("  foo bar	baz   "))
/* FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c) and returns an array of slices of s.
   If all code points in s satisfy f(c) or the string is empty, an empty slice is returned
 */
	f := func(c rune)bool{  // f 参数为 rune 单字符
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	t.Logf("FieldsFunc(\"  foo1;bar2,baz3...\")=%q", strings.FieldsFunc("  foo1;bar2,baz3...", f))

	s := []string{"foo", "bar", "baz"}
	t.Logf("Join(s, \", \")=%s", strings.Join(s, ", "))

	/* If s does not contain sep and sep is not empty,  returns a slice of length 1 whose only element is s
	** If sep is empty, Split splits after each UTF-8 sequence
	** If both s and sep are empty, Split returns an empty slice
	** It is equivalent to SplitN with a count of -1
	 */
	ss := "a,b,c"
	t.Logf("Split(%s, \",\")=%q", ss, strings.Split(ss, ","))
	sep := "xx"
	t.Logf("Split(%s, %s)=%q", ss, sep, strings.Split(ss, sep))
	sep = ""
	t.Logf("Split(%s, %s)=%q", ss, sep, strings.Split(ss, sep))

	/* SplitAfter 与 Split 区别就是 sep SplitAfter 返回结果携带sep */
	ss = "a,b,c"
	sep = ","
	t.Logf("SplitAfter(%s, %s)=%q", ss, sep, strings.SplitAfter(ss, sep))

	/* n > 0: at most n substrings; the last substring will be the unsplit remainder.
	** n == 0: the result is nil (zero substrings)
	** n < 0: all substrings
	*/
	t.Logf("SplitAfterN(%s, %s, 2)=%q", ss, sep, strings.SplitAfterN(ss, sep, 2))
	t.Logf("SplitN(%s, %s, 2)=%q", ss, sep, strings.SplitN(ss, sep, 2))
}

/*
29. func Title(s string) string
30. func ToLower(s string) string
31. func ToLowerSpecial(c unicode.SpecialCase, s string) string
32. func ToTitle(s string) string
33. func ToTitleSpecial(c unicode.SpecialCase, s string) string
34. func ToUpper(s string) string
35. func ToUpperSpecial(c unicode.SpecialCase, s string) string
36. func ToValidUTF8(s, replacement string) string
 */
func TestToFunction(t *testing.T){
	s := "her royal highness"
	// 注： The rule Title uses for word boundaries does not handle Unicode punctuation properly.
	t.Logf("Title(%s) = %s", s, strings.Title(s)) // 首字母大写
	t.Logf("ToTitle(%s) = %s", s, strings.ToTitle(s)) // 全大写

	t.Logf("ToLower(%s) = %s", s, strings.ToLower(s))
	t.Logf("ToUpper(%s) = %s", s, strings.ToUpper(s))

	// 针对非ASCII unicode 字符的处理
	t.Logf("ToLowerSpecial(unicode.TurkishCase, \"Önnek İş\")=%s", strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))
	t.Logf("ToUpperSpecial(unicode.TurkishCase, \"Önnek İş\")=%s", strings.ToUpperSpecial(unicode.TurkishCase, "Önnek İş"))
	t.Logf("ToTitleSpecial(unicode.TurkishCase, \"Önnek İş\")=%s", strings.ToTitleSpecial(unicode.TurkishCase, "Önnek İş"))

	// ToValidUTF8 returns a copy of the string s with each run of invalid UTF-8 byte sequences replaced by the replacement string, which may be empty.
	t.Logf("ToValidUTF8(\"��\", \"\")=%s", strings.ToValidUTF8("hah��aa", ""))
}
/* Trim returns a slice of the string s with all leading and trailing Unicode code points contained in cutset removed.
37. func Trim(s, cutset string) string
38. func TrimFunc(s string, f func(rune) bool) string
39. func TrimLeft(s, cutset string) string
40. func TrimLeftFunc(s string, f func(rune) bool) string
41. func TrimPrefix(s, prefix string) string
 */
func TestTrim(t *testing.T){
	s := "¡¡¡Hello, Gophers!!!"
	tx := "¡!"
	t.Logf("Trim(\"%s\", \"%s\")=%q", s, tx, strings.Trim(s, tx))
	t.Logf("TrimLeft(\"%s\", \"%s\")=%q", s, tx, strings.TrimLeft(s, tx))
	t.Logf("TrimRight(\"%s\", \"%s\")=%q", s, tx, strings.TrimRight(s, tx))

	f := func(r rune)bool{
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}
	t.Logf("TrimFunc(\"%s\", \"%s\")=%q", s, tx, strings.TrimFunc(s, f))
	t.Logf("TrimLeftFunc(\"%s\", \"%s\")=%q", s, tx, strings.TrimLeftFunc(s, f))
	t.Logf("TrimRightFunc(\"%s\", \"%s\")=%q", s, tx, strings.TrimRightFunc(s, f))

	t.Logf("TrimPrefix(\"%s\", \"%s\")=%q", s, "¡¡¡Hello, ", strings.TrimPrefix(s, "¡¡¡Hello, "))
	t.Logf("TrimPrefix(\"%s\", \"%s\")=%q", s, "¡¡¡Howdy, ", strings.TrimPrefix(s, "¡¡¡Howdy, "))
	t.Logf("TrimSuffix(\"%s\", \"%s\")=%q", s, ", Gophers!!!", strings.TrimSuffix(s, ", Gophers!!!"))
	t.Logf("TrimSuffix(\"%s\", \"%s\")=%q", s, ", Marmots!!!", strings.TrimSuffix(s, ", Marmots!!!"))

	t.Logf("TrimSpace(\"\" \\t\\n Hello, Gophers \\n\\t\\r\\n\"\")=%s", strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))
}
/*
  A Builder is used to efficiently build a string using Write methods.
  It minimizes memory copying. The zero value is ready to use.
  Do not copy a non-zero Builder.
方法：
1. func (b *Builder) String() string  // String returns the accumulated string.
2. func (b *Builder) Cap() int // Cap returns the capacity of the builder's underlying byte slice.
							   //It is the total space allocated for the string being built and includes any bytes already written.
3. func (b *Builder) Grow(n int) // Grow grows b's capacity, if necessary, to guarantee space for another n bytes.
								 // After Grow(n), at least n bytes can be written to b without another allocation.
								 // If n is negative, Grow panics.
4. func (b *Builder) Len() int  // Len returns the number of accumulated bytes; b.Len() == len(b.String()).
5. func (b *Builder) Reset()	// Reset resets the Builder to be empty.
6. func (b *Builder) Write(p []byte) (int, error) // Write appends the contents of p to b's buffer. Write always returns len(p), nil.
7. func (b *Builder) WriteByte(c byte) error
8. func (b *Builder) WriteRune(r rune) (int, error)
9. func (b *Builder) WriteString(s string) (int, error)
 */
func TestStringBuilder(t *testing.T){
	var b strings.Builder
	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&b, "%d...", i)
	}
	b.WriteString("inition")
	t.Log(b.String())
	t.Log(b.Cap())
	t.Log(b.Len())
	b.Reset()
	t.Log(b.String())
	b.Write([]byte{'a', 'A'})
	t.Log(b.String())
}
/*
A Reader implements the io.Reader, io.ReaderAt, io.ByteReader, io.ByteScanner,
  io.RuneReader, io.RuneScanner, io.Seeker, and io.WriterTo interfaces by reading from a string.
 The zero value for Reader operates like a Reader of an empty string.
方法：
1. func NewReader(s string) *Reader // NewReader returns a new Reader reading from s. It is similar to bytes.NewBufferString but more efficient and read-only.
2. func (r *Reader) Len() int	// Len returns the number of bytes of the unread portion of the string.返回未读的字符串长度
3. func (r *Reader) Size() int64  //  返回字符串的长度
4. func (r *Reader) Read(b []byte) (n int, err error) // Read implements the io.Reader interface.
5. func (r *Reader) ReadAt(b []byte, off int64) (n int, err error) // ReadAt implements the io.ReaderAt interface.
6. func (r *Reader) ReadByte() (byte, error) // ReadByte implements the io.ByteReader interface. 从当前已读取位置继续读取一个字节
7. func (r *Reader) ReadRune() (ch rune, size int, err error) // ReadRune implements the io.RuneReader interface.
8. func (r *Reader) UnreadRune() error // UnreadRune implements the io.RuneScanner interface.
9. func (r *Reader) UnreadByte() error // UnreadByte implements the io.ByteScanner interface. 将当前已读取位置回退一位，当前位置的字节标记成未读取字节
10. func (r *Reader) Reset(s string) // Reset resets the Reader to be reading from s
11. func (r *Reader) Seek(offset int64, whence int) (int64, error) // Seek implements the io.Seeker interface.
12. func (r *Reader) WriteTo(w io.Writer) (n int64, err error) // WriteTo implements the io.WriterTo interface.
 */
func TestStringReader(t *testing.T){
	r := strings.NewReader("Hello, Gopher")
	t.Log(r.Len(), r.Size()) // 未读取， 未读长度等于字符串长度
	// 读取的 取决于  buff 大小
	buff := []byte{}
	readLen, err := r.Read(buff)
	t.Log(readLen, err)
	t.Log(r.Len(), r.Size())

	buff = make([]byte, 5)
	readLen, err = r.Read(buff)
	t.Log(readLen, err)
	t.Log(string(buff), r.Len(), r.Size())
	//ReadAt读取偏移off字节后的剩余信息到b中,
	bufAt := make([]byte, 2)
	r.ReadAt(bufAt, 2)
	t.Log(string(bufAt), r.Len()) // 与Read不同，ReadAt不会影响 Len() 即未读取长度

	c, err := r.ReadByte()
	t.Log(c, r.Len()) // 影响 Len()
	cx, size, err := r.ReadRune() // size 返回
	t.Log(cx, size, err, r.Len())
	r.UnreadByte()  // UnreadByte 与 Unreadrune 是分开的
	t.Log(r.Len())
	r.UnreadRune()
	t.Log(r.Len())
	r.UnreadByte()
	t.Log(r.Len())
	r.UnreadRune()
	t.Log(r.Len())
/*	前面有ReadAt方法可以将字符串偏移多少位读取剩下的字符串内容，但是该方法不会影响正在用Read方法读取的内容，
	如果相对Read方法读取的内容做偏移就可以使用seek方法，
	offset是偏移的位置，whence是偏移起始位置，支持三种位置（io.SeekStart起始位，io.SeekCurrent当前位，io.SeekEnd末位）
 	需要注意的是offset可以为负数，但是 偏移起始位 与offset相加得到的值不能小于0或者大于size()的长度
 */
	r.Seek(-2, io.SeekCurrent) // 从当前位置向前偏移2位, 影响到Len()
	t.Log(string(buff), r.Len())
	r.Read(buff)
	t.Log(string(buff), r.Len())
	r.Reset("hah")
	t.Log(r.Len())
}

/*
1. func NewReplacer(oldnew ...string) *Replacer  // NewReplacer returns a new Replacer from a list of old, new string pairs.
2. func (r *Replacer) Replace(s string) string // Replace returns a copy of s with all replacements performed.
3. func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error) // WriteString writes s to w with all replacements performed.
*/
func TestStringReplacer(t *testing.T){
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	t.Log(r.Replace("This is <b>HTML</b>!"))
	r.WriteString(os.Stdout, "This is <b>HTML</b>!\n")
}

