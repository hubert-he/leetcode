package DP

import (
	"testing"
	"../tree"
)

func TestCanPartition(t *testing.T) {

}

func TestCoinChange(t *testing.T){
	for caseId, testCase := range []struct{
		coins		[]int
		amount		int
		want		int
	}{
		{[]int{1}, 0, 0},
		{[]int{1}, 1, 1},
		{[]int{1, 2}, 3, 2},
		{[]int{1,2,5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{2}, 4, 2},
		{[]int{186,419,83,408}, 6249, 20},
		{[]int{2,5,10,1}, 27, 4},
		{[]int{224,2,217,189,79,343,101}, 2938, 11},
	}{
		result := CoinChange(testCase.coins, testCase.amount)
		if result != testCase.want{
			t.Errorf("case-%d result=%d want=%d", caseId, result, testCase.want)
			break
		}
	}
}

func TestFindMaxForm(t *testing.T){
	for caseId, testCase := range []struct{
		strs	[]string
		limit	[2]int
		want	int
	}{
		{[]string{"00", "000"}, [2]int{1,10}, 0},
		{[]string{"11", "111"}, [2]int{1,10}, 2},
		{[]string{"10", "0001", "111001", "1", "0"}, [2]int{5, 3}, 4},
		{[]string{"10", "0", "1"}, [2]int{1,1}, 2},
	}{
		result := FindMaxForm(testCase.strs, testCase.limit[0], testCase.limit[1])
		if result != testCase.want{
			t.Errorf("case-%d: result=%d want=%d", caseId, result, testCase.want)
			break
		}
	}
}

func TestFindTargetSumWays(t *testing.T){
	for caseId, testCase := range []struct{
		nums	[]int
		target	int
		want	int
	}{
		{[]int{1}, 1, 1},
		{[]int{0,0,0,0,0,0,0,0,1}, 1, 256},
		{[]int{1,1,1,1,1}, 3, 5},
		{[]int{2,107,109,113,127,131,137,3,2,3,5,7,11,13,17,19,23,29,47,53}, 1000, 0},
		{[]int{1000}, -1000, 1},
	}{
		result := FindTargetSumWays(testCase.nums, testCase.target)
		if result != testCase.want{
			t.Errorf("case-%d failed: result=%d want=%d", caseId, result, testCase.want)
			break
		}

	}
}

func TestProfitableSchemes(t *testing.T){
	for caseId, testCase := range []struct{
		n			int
		minProfit	int
		group		[]int
		profit		[]int
		want		int
	}{
		{1, 1, []int{1,1,1,1,2,2,1,2,1,1}, []int{0,1,0,0,1,1,1,0,2,2}, 4},
		{5, 3, []int{2,2}, []int{2,3}, 2},
		{10, 5, []int{2,3,5}, []int{6,7,8}, 7},
		{100, 100,
			[]int{2,5,36,2,5,5,14,1,12,1,14,15,1,1,27,13,6,59,6,1,7,1,2,7,6,1,6,1,3,1,2,11,3,39,21,20,1,27,26,22,11,17,3,2,4,5,6,18,4,14,1,1,1,3,12,9,7,3,16,5,1,19,4,8,6,3,2,7,3,5,12,6,15,2,11,12,12,21,5,1,13,2,29,38,10,17,1,14,1,62,7,1,14,6,4,16,6,4,32,48},
			[]int{21,4,9,12,5,8,8,5,14,18,43,24,3,0,20,9,0,24,4,0,0,7,3,13,6,5,19,6,3,14,9,5,5,6,4,7,20,2,13,0,1,19,4,0,11,9,6,15,15,7,1,25,17,4,4,3,43,46,82,15,12,4,1,8,24,3,15,3,6,3,0,8,10,8,10,1,21,13,10,28,11,27,17,1,13,10,11,4,36,26,4,2,2,2,10,0,11,5,22,6},
			692206787,
		},
	}{
		result := ProfitableSchemes(testCase.n, testCase.minProfit, testCase.group, testCase.profit)
		if result != testCase.want{
			t.Errorf("case-%d faild: result=%d want=%d", caseId, result, testCase.want)
			break
		}
	}
}

func TestChange(t *testing.T) {
	for caseId, testCase := range []struct{
		amount	int
		coins	[]int
		want	int
	}{
		{5, []int{1,2,5}, 4},
		{3, []int{2}, 0},
		{10, []int{10}, 1},
		{
			4000,
			[]int{200,217,234,251,268,285,302,319,336,353,370,387,404,421,438,455,472,489,506,523,540,557,574,591,608,625,642,659,676,693,710,727,744,761,778,795,812,829,846,863,880,897,914,931,948,965,982,999,1016,1033,1050,1067,1084,1101,1118,1135,1152,1169,1186,1203,1220,1237,1254,1271,1288,1305,1322,1339,1356,1373,1390,1407,1424,1441,1458,1475,1492,1509,1526,1543,1560,1577,1594,1611,1628,1645,1662,1679,1696,1713,1730,1747,1764,1781,1798,1815,1832,1849,1866,1883,1900,1917,1934,1951,1968,1985,2002,2019,2036,2053,2070,2087,2104,2121,2138,2155,2172,2189,2206,2223,2240,2257,2274,2291,2308,2325,2342,2359,2376,2393,2410,2427,2444,2461,2478,2495,2512,2529,2546,2563,2580,2597,2614,2631,2648,2665,2682,2699,2716,2733,2750,2767,2784,2801,2818,2835,2852,2869,2886,2903,2920,2937,2954,2971,2988,3005,3022,3039,3056,3073,3090,3107,3124,3141,3158,3175,3192,3209,3226,3243,3260,3277,3294,3311,3328,3345,3362,3379,3396,3413,3430,3447,3464,3481,3498,3515,3532,3549,3566,3583,3600,3617,3634,3651,3668,3685,3702,3719,3736,3753,3770,3787,3804,3821,3838,3855,3872,3889,3906,3923,3940,3957,3974,3991,4008,4025,4042,4059,4076,4093,4110,4127,4144,4161,4178,4195,4212,4229,4246,4263,4280,4297,4314,4331,4348,4365,4382,4399,4416,4433,4450,4467,4484,4501,4518,4535,4552,4569,4586,4603,4620,4637,4654,4671,4688,4705,4722,4739,4756,4773,4790,4807,4824,4841,4858,4875,4892,4909,4926,4943,4960,4977,4994},
			3435,
		},
		{500, []int{3,5,7,8,9,10,11}, 35502874},
	}{
		result := ChangeDP(testCase.amount, testCase.coins)
		if result != testCase.want{
			t.Errorf("case-%d failed: result=%d want=%d", caseId, result, testCase.want)
			break
		}
	}
}

func TestGroupingKnapsack(t *testing.T) {
	for caseId, testCase := range []struct{
		N		int
		C		int
		S		[]int
		v		[][]int
		w		[][]int
		want	int
	}{
		{2, 9, []int{2, 3}, [][]int{[]int{1, 2, -1}, []int{1, 2, 3}}, [][]int{[]int{2, 4, -1}, []int{1, 3, 6}}, 10},
	}{
		result := GroupingKnapsack(testCase.N, testCase.C, testCase.S, testCase.v, testCase.w)
		if result != testCase.want{
			t.Errorf("GroupingKnapsack-case-%d failed: result=%d want=%d", caseId, result, testCase.want)
			break
		}
		result = GroupingKnapsack1(testCase.N, testCase.C, testCase.S, testCase.v, testCase.w)
		if result != testCase.want{
			t.Errorf("GroupingKnapsack1-case-%d failed: result=%d want=%d", caseId, result, testCase.want)
			break
		}
	}
}

func TestNumRollsToTarget(t *testing.T) {
	for caseId, testCase := range []struct{
		d		int
		f		int
		target	int
		want	int
	}{
		{30, 30, 500, 222616187},
		{1, 6, 3, 1},
		{2, 6, 7, 6},
		{2, 5, 10, 1},
		{1, 2, 3, 0},
	}{
		result := NumRollsToTarget(testCase.d, testCase.f, testCase.target)
		if result != testCase.want{
			t.Errorf("NumRollsToTarget-case-%d failed: result=%d want=%d", caseId, result, testCase.want)
			break
		}
		result = NumRollsToTarget2(testCase.d, testCase.f, testCase.target)
		if result != testCase.want{
			t.Errorf("NumRollsToTarget2-case-%d failed: result=%d want=%d", caseId, result, testCase.want)
			break
		}
		result = NumRollsToTarget1(testCase.d, testCase.f, testCase.target)
		if result != testCase.want{
			t.Errorf("NumRollsToTarget1-case-%d failed: result=%d want=%d", caseId, result, testCase.want)
			break
		}
	}
}

func TestWordBreak(t *testing.T) {
	for caseId, testCase := range []struct{
		s			string
		wordDict	[]string
		want		bool
	}{
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
			[]string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"}, false},
		{"leetcode", []string{"leet","code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats","dog","sand","and","cat"}, false},
	}{
		result := WordBreakDFS(testCase.s, testCase.wordDict)
		if result != testCase.want{
			t.Errorf("TestWordBreakDFS-case-%d: result=%t but want=%t", caseId, result, testCase.want)
			break
		}
		result = WordBreakDP(testCase.s, testCase.wordDict)
		if result != testCase.want{
			t.Errorf("TestWordBreakDP-case-%d: result=%t but want=%t", caseId, result, testCase.want)
			break
		}
	}
}

func TestPaintBinaryTree(t *testing.T) {
	for caseId, testCase := range []struct{
		nums		[]interface{}
		k			int
		want		int
	}{
		{[]interface{}{5,2,3,4}, 2, 12},
		{[]interface{}{4,1,3,9,nil,nil,2}, 2, 16},
		{[]interface{}{8,1,3,9,9,9,nil,9,5,6,8}, 2, 52},
	}{
		tree := Tree.GenerateBiTree(testCase.nums)
		result := PaintBinaryTree(tree, testCase.k)
		if result != testCase.want{
			t.Errorf("PaintBinaryTree-case-%d: result=%d, but want=%d", caseId, result, testCase.want)
			break
		}
		result = PaintBinaryTree2(tree, testCase.k)
		if result != testCase.want{
			t.Errorf("PaintBinaryTree2-case-%d: result=%d, but want=%d", caseId, result, testCase.want)
			break
		}
	}
}
