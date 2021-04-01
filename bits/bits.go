package bits

// 190 Reverse Bits
const (
	m1 = 0x55555555
	m2 = 0x33333333
	m4 = 0x0F0F0F0F
	m8 = 0x00FF00FF
	m16 = 0x0000FFFF
)
// 分治，自底向上
func reverseBits(num uint32) uint32 {
	n := num
	n = n >> 1 & m1 | n & m1 << 1
	n = n >> 2 & m2 | n & m2 << 2
	n = n >> 4 & m4 | n & m4 << 4
	n = n >> 8 & m8 | n & m8 << 8
	n = n >> 16 & m16 | n & m16 << 16
	return n
}