package string

import "strconv"

/* 415. Add Strings
Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.
You must solve the problem without using any built-in library for handling large integers (such as BigInteger).
You must also not convert the inputs to integers directly.
Example 1:
	Input: num1 = "11", num2 = "123"
	Output: "134"
Example 2:
	Input: num1 = "456", num2 = "77"
	Output: "533"
Example 3:
	Input: num1 = "0", num2 = "0"
	Output: "0"
 */
func addStrings(num1 string, num2 string) string {
	carry := 0
	ans := ""
	for i,j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0|| carry > 0; i,j = i-1,j-1{
		x,y := 0,0
		if i >= 0{
			x = int(num1[i] - '0')
		}
		if j >= 0{
			y = int(num2[j] - '0')
		}
		result := x + y + carry
		ans = strconv.Itoa(result%10) + ans
		carry = result / 10
	}
	return ans
}
func addStrings2(num1 string, num2 string) string {
	var carry byte
	ans := []byte{}
	for i,j :=len(num1)-1,len(num2)-1; i>=0 || j>=0 || carry>0; i,j=i-1,j-1{
		var x,y byte
		if i >= 0{
			x = num1[i] - '0'
		}
		if j >=0 {
			y = num2[j] - '0'
		}
		sum := x + y + carry
		ans = append([]byte{(sum%10)+'0'}, ans...)
		carry = sum / 10
	}
	return string(ans)
}

