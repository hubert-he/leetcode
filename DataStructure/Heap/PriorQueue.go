package Heap

import (
	"container/heap"
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
	"sort"
)

/* 253. Meeting Rooms II
** Given an array of meeting time intervals intervals where intervals[i] = [starti, endi],
** return the minimum number of conference rooms required.
 */
// 2022-04-14 刷出此题
func minMeetingRooms(intervals [][]int) int {
	const Start, End = 0, 1
	sort.Slice(intervals, func(i, j int)bool{
		return intervals[i][Start] < intervals[j][Start]
	})
	rooms := [][][]int{}
	isOverLap := func(timeLime [][]int, interval []int)bool{
		for i := range timeLime{
			if timeLime[i][Start] >= interval[End] || interval[Start] >= timeLime[i][End]{
				continue
			}
			return true
		}
		return false
	}
	for i := range intervals{
		j := 0
		for ; j < len(rooms); j++{
			if !isOverLap(rooms[j], intervals[i]){
				rooms[j] = append(rooms[j], intervals[i])
				break
			}
		}
		if j >= len(rooms){
			rooms = append(rooms, [][]int{intervals[i]})
		}
	}
	return len(rooms)
}

/* 使用heap 优化
** 上面的方法使用的是一个朴素的方法， 即每当有新会议时，就遍历所有房间，查看是否有空闲房间。
** 但是,通过使用优先队列（或最小堆）堆数据结构,我们可以做得更好。
** 我们可以将所有房间保存在最小堆中,堆中的键值是会议的结束时间，而不用手动迭代已分配的每个房间并检查房间是否可用。
** 这样，每当我们想要检查有没有 任何 房间是空的，只需要检查最小堆堆顶的元素，它是最先开完会腾出房间的。
** 如果堆顶的元素的房间并不空闲，那么其他所有房间都不空闲。这样，我们就可以直接开一个新房间。
 */
type rooms struct{
	r []int
	sort.Interface
}
func (this *rooms) Push(x interface{}){
	this.r = append(this.r, x.(int))
}
func (this *rooms) Pop()interface{}{
	ret := this.r[this.Len()-1]
	this.r = this.r[:this.Len()-1]
	return ret
}
func minMeetingRooms_heap(intervals [][]int) int {
	const Start, End = 0, 1
	n := len(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][Start] < intervals[j][Start]
	})
	h := &rooms{}
	heap.Init(h)
	// 初始化一个新的 最小堆，将第一个会议的结束时间加入到堆中。
	// 我们只需要记录会议的结束时间，告诉我们什么时候房间会空
	heap.Push(h, intervals[0][End])
	for i := 1; i < n; i++{
		//对每个会议，检查堆的最小元素（即堆顶部的房间）是否空闲
		if intervals[i][Start] >= h.r[0]{
			// 若房间空闲，则从堆顶拿出该元素，将其改为我们处理的会议的结束时间，加回到堆中
			h.r[0] = intervals[i][End]
			heap.Fix(h, 0)
		}else{//若房间不空闲。开新房间，并加入到堆中。
			heap.Push(h, intervals[i][End])
		}
	}
	return h.Len()
}

/* 1606. Find Servers That Handled Most Number of Requests
** You have k servers numbered from 0 to k-1 that are being used to handle multiple requests simultaneously.
** Each server has infinite computational capacity but cannot handle more than one request at a time.
** The requests are assigned to servers according to a specific algorithm:
	The ith (0-indexed) request arrives.
	If all servers are busy, the request is dropped (not handled at all).
	If the (i % k)th server is available, assign the request to that server.
	Otherwise, assign the request to the next available server (wrapping around the list of servers and starting from 0 if necessary).
	For example, if the ith server is busy, try to assign the request to the (i+1)th server, then the (i+2)th server, and so on.
** You are given a strictly increasing array arrival of positive integers,
** where arrival[i] represents the arrival time of the ith request, and another array load,
** where load[i] represents the load of the ith request (the time it takes to complete).
** Your goal is to find the busiest server(s). A server is considered busiest if it handled the most number of requests successfully among all the servers.
** Return a list containing the IDs (0-indexed) of the busiest server(s). You may return the IDs in any order.
 */
/* 2022-04-18 未刷出此题，主要问题是潜在的 有序集合，选机器要就有顺序
** 考虑这个case， 返回结果[0,1] 真正答案为 [1]
	3
	[1,2,3,4,8,9,10]
	[5,2,10,3,1,2,2]
 */
func busiestServers_Error(k int, arrival []int, load []int) []int {
	usage := make([][]int, k)
	h := TimeFreeHeap{}
	for i := 0; i < k; i++{
		h = append(h, free{0, i})
	}
	heap.Init(&h)
	for i := range arrival{
		if h[0].end > arrival[i]{
			continue // drop req
		}
		top := heap.Pop(&h).(free)
		usage[top.id] = append(usage[top.id], i)
		top.end = arrival[i] + load[i]
		heap.Push(&h, top)
	}
	fmt.Println(usage)
	ans := []int{}
	max := 0
	for i := range usage{
		t := len(usage[i])
		if max < t{
			ans = []int{i}
			max = t
		}else if max == t{
			ans = append(ans, i)
		}
	}
	return ans
}
type free struct{
	end     int
	id      int
}
type TimeFreeHeap []free
func(h TimeFreeHeap)Len()int{
	return len(h)
}
func(h TimeFreeHeap)Swap(i, j int){
	h[i], h[j] = h[j], h[i]
}
func(h TimeFreeHeap)Less(i, j int)bool{
	if h[i].end == h[j].end{
		return h[i].id < h[j].id
	}
	return h[i].end < h[j].end
}
func(h *TimeFreeHeap)Pop()interface{}{
	ret := (*h)[h.Len()-1]
	(*h) = (*h)[:h.Len()-1]
	return ret
}
func(h *TimeFreeHeap)Push(x interface{}){
	(*h) = append((*h), x.(free))
}
// 官方解答：借助red-black tree
func busiestServers(k int, arrival, load []int) (ans []int) {
	available := redblacktree.NewWithIntComparator() // 使用了红黑树
	for i := 0; i < k; i++{
		available.Put(i, nil)
	}
	busy := TimeFreeHeap{}
	requests := make([]int, k)
	maxRequest := 0
	for i, t := range arrival{
		// 先全部装载获得最新的available
		/* 假设当前到达的请求为第 i 个，如果 busy 不为空，
		** 那么我们判断 busy 的队首对应的服务器的结束时间是否小于等于当前请求的到达时间 arrival[i]，
		*** 如果是，那么我们将它从 busy 中移除，并且将该服务器的编号放入 available 中，然后再次进行判断。
		 */
		for busy.Len() > 0 && busy[0].end <= t{
			available.Put(busy[0].id, nil)
			heap.Pop(&busy)
		}
		if available.Size() == 0{
			continue // drop req
		}
		node, _ := available.Ceiling(i % k)
		if node == nil {
			node = available.Left()
		}
		id := node.Key.(int)
		requests[id]++
		if requests[id] > maxRequest{
			maxRequest = requests[id]
			ans = []int{id}
		}else if requests[id] == maxRequest{
			ans = append(ans, id)
		}
		heap.Push(&busy, free{t + load[i], id})
		available.Remove(id)
	}
	return
}
// 使用双优先队列，
// 设在第 i 个请求到来时，编号为 id 的服务器已经处理完请求，那么将 id 从 busy 中移除，并放入一个不小于 i 且同余于 id 的编号，
// 这样就能在保证 available 中，编号小于 i % k 的空闲服务器能排到编号不小于 i % k 的空闲服务器后面
//
type hi struct{ sort.IntSlice }
func(h *hi)Push(x interface{}){
	h.IntSlice = append(h.IntSlice, x.(int))
}
func(h *hi)Pop()interface{}{
	ret := h.IntSlice[h.Len()-1]
	h.IntSlice = h.IntSlice[:h.Len()-1]
	return ret
}
func busiestServers_2heap(k int, arrival, load []int) (ans []int) {
	available := hi{make([]int, k)} // 注意初始化
	// 首先 装载 available
	for i := 0; i < k; i++{
		available.IntSlice[i] = i
	}
	busy := TimeFreeHeap{}
	/*
	for i := range arrival{
		busy = append(busy, free{arrival[i] + load[i], i})
	}*/
	//heap.Init(&busy)
	heap.Init(&available)
	requests := make([]int, k)
	maxRequest := 0
	for i := range arrival{
		start, end := arrival[i], arrival[i] + load[i]
		for busy.Len() > 0 && busy[0].end <= start{
			// 保证得到的是一个不小于 i 的且与 id 同余的数
			heap.Push(&available, i + ((busy[0].id - i) % k + k) % k)// 这个是计算差值模值
			heap.Pop(&busy)
		}
		if available.Len() == 0 {
			continue
		}
		id := heap.Pop(&available).(int) % k
		requests[id]++
		if requests[id] > maxRequest{
			maxRequest = requests[id]
			ans = []int{id}
		}else if requests[id] == maxRequest{
			ans = append(ans, id)
		}
		heap.Push(&busy, free{end, id})
	}
	return
}