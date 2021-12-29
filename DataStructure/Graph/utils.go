package Graph

import "math"

func max(nums ...int)int{
	m := nums[0]
	for _, u := range nums{
		if m < u{
			m = u
		}
	}
	return m
}

func min(nums ...int) int {
	m := nums[0]
	for _, c := range nums{
		if m > c{
			m = c
		}
	}
	return m
}
// 排除except 下标的 值
func min2(except int, nums ...int) int{
	m := math.MaxInt32
	for i := range nums{
		if m > nums[i] && i != except{
			m = nums[i]
		}
	}
	return m
}
type pair struct {d, x int}
type hp []pair
func(h hp)Len() int{return len(h)}
func(h hp)Less(i, j int)bool { return h[i].d < h[j].d}
func(h hp)Swap(i, j int) { h[i], h[j] = h[j], h[i]}
func(h *hp)Push(v interface{}){ *h = append(*h, v.(pair))}
func(h *hp)Pop() (v interface{}){ a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

