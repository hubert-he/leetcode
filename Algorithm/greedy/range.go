package greedy

import "sort"

/* 452. Minimum Number of Arrows to Burst Balloons
** There are some spherical balloons taped onto a flat wall that represents the XY-plane.
** The balloons are represented as a 2D integer array points
	where points[i] = [xstart, xend] denotes a balloon whose horizontal diameter stretches between xstart and xend.
** You do not know the exact y-coordinates of the balloons.
** Arrows can be shot up directly vertically (in the positive y-direction) from different points along the x-axis.
** A balloon with xstart and xend is burst by an arrow shot at x if xstart <= x <= xend.
** There is no limit to the number of arrows that can be shot.
** A shot arrow keeps traveling up infinitely, bursting any balloons in its path.
** Given the array points, return the minimum number of arrows that must be shot to burst all balloons.
 */
/* 思路从箭的角度思考
** 首先随机地射出一支箭，再看一看是否能够调整这支箭地射出位置，使得我们可以引爆更多数目的气球。向右移动（为何向右，因为会先把气球排序好）
** 查看是否能在保证原本引爆气球的基础上，可以进一步引爆更多气球。
** 那么我们最远可以将这支箭往右移动多远呢？我们唯一的要求就是：原本引爆的气球只要仍然被引爆就行了
** 这样一来，我们找出原本引爆的气球中「右边界位置最靠左的那一个」，将这支箭的射出位置移动到这个右边界位置，这也是最远可以往右移动到的位置。
** 因为只要我们再往右移动一点点，这个气球就无法被引爆了。
** 因此，我们可以断定：
	「一定存在一种最优（射出的箭数最小）的方法，使得每一支箭的射出位置都恰好对应着某一个气球的右边界。」
** 有了这样一个有用的断定，我们就可以快速得到一种最优的方法了。
** 考虑所有气球中右边界位置最靠左的那一个，那么一定有一支箭的射出位置就是它的右边界（否则就没有箭可以将其引爆了）。
** 当我们确定了一支箭之后，我们就可以将这支箭引爆的所有气球移除，并从剩下未被引爆的气球中，再选择右边界位置最靠左的那一个，确定下一支箭，
** 直到所有的气球都被引爆。
** 复杂度分析：
   最坏情况， 当 n 个气球对应的区间互不重叠， 两个for 是重叠进行的，其中内部的for循环纯属多余计算
   因此造成复杂度为 O(n^2)
 */
func findMinArrowShots(points [][]int) int {
	const Start, End = 0, 1
	n := len(points)
	burst, cnt := make([]bool, n), n
	sort.Slice(points, func(i, j int)bool{ return points[i][End] < points[j][End]})
	ans := 0
	i := 0 // 最小的满足burst[i] = false 的索引 i
	for cnt > 0{ // 还有未爆的气球
		for i < n && burst[i] { // 选取 i 为 右边界位置最靠左的那一个
			i++
		}
		for j := i; j < n; j++{ // 凡是左边界的 气球 都在 i 右边界之内的 都是会引爆的
			if !burst[j] && points[j][Start] <= points[i][End]{
				burst[j] = true
				cnt--
			}
		}
		ans++
	}
	return ans
}
func findMinArrowShots_improve(points [][]int) int {
	const Start, End = 0, 1
	n := len(points)
	burst, cnt := make([]bool, n), n
	sort.Slice(points, func(i, j int)bool{ return points[i][End] < points[j][End]})
	ans := 0
	i := 0 // 最小的满足burst[i] = false 的索引 i
	for cnt > 0{ // 还有未爆的气球
		for i < n && burst[i] { // 选取 i 为 右边界位置最靠左的那一个
			i++
		}
		for j := i; j < n; j++{ // 凡是左边界的 气球 都在 i 右边界之内的 都是会引爆的
			if !burst[j] && points[j][Start] <= points[i][End]{
				burst[j] = true
				cnt--
			}else if !burst[j]{
				break
			}
		}
		ans++
	}
	return ans
}
/* 针对内层循环的优化，当我们遇到第一不满足 points[j][Start] <= points[i][End] 的时候，就完全可以跳出循环，
** 并且这个 points[j][End] 就是下一支箭的射出位置。
** 也就是说：当前这支箭在索引 j(第一个无法引爆的气球）之后所有可以引爆的气球，下一支箭也都可以引爆
*/
func findMinArrowShots2(points [][]int) int {
	const Start, End = 0, 1
	n := len(points)
	if n <= 0{ return 0 }
	sort.Slice(points, func(i, j int)bool{ return points[i][End] < points[j][End]})
	maxRight := points[0][End]
	ans := 1
	for i := 1; i < n; i++{
		p := points[i]
		if p[Start] > maxRight{
			maxRight = p[End]
			ans++
		}
	}
	return ans
}