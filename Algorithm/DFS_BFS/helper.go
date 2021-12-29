package DFS_BFS
func min(nums ...int)int{
	m := nums[0]
	for _, c := range nums{
		if m > c {
			m = c
		}
	}
	return m
}

func max(nums ...int)int{
	m := nums[0]
	for _, c := range nums{
		if m < c {
			m = c
		}
	}
	return m
}
