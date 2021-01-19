package main

import (
	"fmt"
	"strconv"
)

const NumberFSMMaxState = 9
const NumberFSMMaxCondition = 6
const INVALID_STATE = -1
const (
	SPACE = iota
	SIGN
	DECNUM
	POINT
	EXPONENTIAL
	END
)
const (
	START_STATE = iota
	SIGN_STATE
	POINT_STATE
	NUMBER_STATE
	NUM_POINT_STATE
	E_STATE
	E_NUMBER_STATE
	E_SIGN_STATE
	END_STATE
)

func main() {
	//test_isNumber()
	test_isIP()
}

func test_isNumber() {
	fmt.Println(" 3. =>", isNumber("3.e2"))
	fmt.Println(" 0.1 =>", isNumber(" 0.1 "))
	fmt.Println("0=>", isNumber("2"))
	fmt.Println("abc=>", isNumber("abc"))
	fmt.Println("1 a=>", isNumber("1 a"))
	fmt.Println("2e10=>", isNumber("2e10"))
	fmt.Println(" -90e3   =>", isNumber(" -90e3   "))
	fmt.Println(" 1e=>", isNumber(" 1e"))
	fmt.Println("e3=>", isNumber("e3"))
	fmt.Println(" 6e-1=>", isNumber(" 6e-1"))
	fmt.Println(" 99e2.5 =>", isNumber(" 99e2.5 "))
	fmt.Println("53.5e93=>", isNumber("53.5e93"))
	fmt.Println(" --6 =>", isNumber(" --6 "))
	fmt.Println("+-3=>", isNumber("+-3"))
	fmt.Println("95a54e53=>", isNumber("95a54e53"))
}

func isNumber(s string) bool {
	FSMTable := [NumberFSMMaxState][NumberFSMMaxCondition]int{
		{START_STATE, SIGN_STATE, NUMBER_STATE, POINT_STATE, INVALID_STATE, INVALID_STATE},
		{INVALID_STATE, INVALID_STATE, NUMBER_STATE, POINT_STATE, INVALID_STATE, INVALID_STATE},
		{INVALID_STATE, INVALID_STATE, NUM_POINT_STATE, INVALID_STATE, INVALID_STATE, INVALID_STATE},
		{END_STATE, INVALID_STATE, NUMBER_STATE, NUM_POINT_STATE, E_STATE, END_STATE},
		{END_STATE, INVALID_STATE, NUM_POINT_STATE, INVALID_STATE, E_STATE, END_STATE},
		{INVALID_STATE, E_SIGN_STATE, E_NUMBER_STATE, INVALID_STATE, INVALID_STATE, INVALID_STATE},
		{END_STATE, INVALID_STATE, E_NUMBER_STATE, INVALID_STATE, INVALID_STATE, END_STATE},
		{INVALID_STATE, INVALID_STATE, E_NUMBER_STATE, INVALID_STATE, INVALID_STATE, INVALID_STATE},
		{END_STATE, INVALID_STATE, INVALID_STATE, INVALID_STATE, INVALID_STATE, END_STATE},
	}
	state := START_STATE
	for _, char := range s {
		switch {
		case char == ' ':
			state = FSMTable[state][SPACE]
			break
		case char == '+' || char == '-':
			state = FSMTable[state][SIGN]
			break
		case char >= '0' && char <= '9':
			state = FSMTable[state][DECNUM]
			break
		case char == '.':
			state = FSMTable[state][POINT]
			break
		case char == 'e':
			state = FSMTable[state][EXPONENTIAL]
			break
		default:
			state = INVALID_STATE
		}
		if state == INVALID_STATE {
			return false
		}
	}
	state = FSMTable[state][END]
	if state == INVALID_STATE {
		return false
	}
	return true
}

func test_isIP() {
	fmt.Println("0.0.0", isIPv4("0.0.0"))
	fmt.Println("0.0.0.256", isIPv4("0.0.0.256"))
	fmt.Println("01.01.01.01", isIPv4("01.01.01.01"))
	fmt.Println("1.1.1.01", isIPv4("1.1.1.01"))
	fmt.Println(" 192.168.1.1 =>", isIPv4(" 192.168.1.1 "))
	fmt.Println("256.3.4.1=>", isIPv4("256.3.4.1"))
	fmt.Println("1.3.4.1=>", isIPv4("1.3.4.1"))
	fmt.Println("1.3.4.1.=>", isIPv4("1.3.4.1."))
	fmt.Println("1.3.4.1.5=>", isIPv4("1.3.4.1.5"))
	fmt.Println("172.16.254.1 => ", isIPv4("172.16.254.1"))
	fmt.Println("0.0.0.0 => ", isIPv4("0.0.0.0"))
	fmt.Println("255.255.255.255 => ", isIPv4("255.255.255.255"))
	fmt.Println("3456.erdf.0.0 => ", isIPv4("3456.erdf.0.0"))

	fmt.Println("255.255.255 => ", isIPv4("255.255.255"))
	fmt.Println("2001:0410::FB00:1400:5000:45FF =>", isIPv6("2001:0410::FB00:1400:5000:45FF"))
	fmt.Println("::1 =>", isIPv6("::1"))
	fmt.Println(":::1 =>", isIPv6(":::1"))
	fmt.Println("::55ad:1 =>", isIPv6("::1"))
	fmt.Println("0:0:0:0:0:0:138.1.1.1 =>", isIPv6("0:0:0:0:0:0:138.1.1.1"))
	fmt.Println("2001:0db8:85a3:0:0:8A2E:0370:73341 => ", isIPv6("2001:0db8:85a3:0:0:8A2E:0370:73341"))
	fmt.Println("20EE:Fb8:85a3:0:0:8A2E:0370:7334:12 => ", isIPv6("20EE:Fb8:85a3:0:0:8A2E:0370:7334:12"))
	fmt.Println("2001:db8:85a3:0::8a2E:0370:7334 => ", isIPv6("2001:db8:85a3:0::8a2E:0370:7334"))
}
func isIPv4(s string) bool {
	const (
		SPACE  = iota
		DECNUM //1-9
		POINT
		ZERO //0
		END
	)
	const (
		START_STATE = iota
		NUMBER_STATE
		POINT_STATE
		ZERO_STATE
		END_STATE
	)
	FSMTable := [][]int{
		{START_STATE, NUMBER_STATE, INVALID_STATE, ZERO_STATE, INVALID_STATE},
		{END_STATE, NUMBER_STATE, POINT_STATE, NUMBER_STATE, END_STATE},
		{INVALID_STATE, NUMBER_STATE, INVALID_STATE, ZERO_STATE, INVALID_STATE},
		{END_STATE, INVALID_STATE, POINT_STATE, INVALID_STATE, END_STATE},
		{END_STATE, INVALID_STATE, INVALID_STATE, INVALID_STATE, END_STATE},
	}
	state := START_STATE
	pointNum := 0
	num := []byte{}
	for _, ch := range s {
		switch {
		case ch == ' ':
			state = FSMTable[state][SPACE]
			break
		case ch == '0':
			if pointNum >= 4 {
				fmt.Println(pointNum)
				state = INVALID_STATE
			} else {
				num = append(num, byte(ch))
				state = FSMTable[state][ZERO]
			}
		case ch >= '1' && ch <= '9':
			if pointNum >= 4 {
				fmt.Println(pointNum)
				state = INVALID_STATE
			} else {
				num = append(num, byte(ch))
				state = FSMTable[state][DECNUM]
			}
			break
		case ch == '.':
			// check num >=0 && num <=255
			pointNum++
			if pointNum >= 4 {
				//fmt.Println(pointNum)
				state = INVALID_STATE
			} else {
				value, err := strconv.Atoi(string(num[:]))
				if err != nil || value < 0 || value > 255 {
					fmt.Println(string(num[:]))
					state = INVALID_STATE
				} else {
					state = FSMTable[state][POINT]
				}
				num = nil
			}
			break
		default:
			state = INVALID_STATE
		}
		//fmt.Printf("%c, %d\n", ch, state)
		if state == INVALID_STATE {
			return false
		}
	}
	state = FSMTable[state][END]
	value, err := strconv.Atoi(string(num[:]))
	if err != nil || value < 0 || value > 255 {
		state = INVALID_STATE
	}
	if state == INVALID_STATE || pointNum != 3 {
		return false
	}
	return true
}
func isIPv6(s string) bool {
	const (
		SPACE = iota
		HEXNUM
		COLON
		END
	)
	const (
		START_STATE = iota
		NUMBER_STATE
		COLON_STATE
		DOUBLE_COLON_STATE
		HEAD_COLON_STATE
		AFTER_NUMBER_STATE
		AFTER_COLON_STATE
		END_STATE
	)
	FSMTable := [][]int{
		{START_STATE, NUMBER_STATE, HEAD_COLON_STATE, INVALID_STATE},
		{INVALID_STATE, NUMBER_STATE, COLON_STATE, END_STATE},
		{INVALID_STATE, NUMBER_STATE, DOUBLE_COLON_STATE, INVALID_STATE},
		{INVALID_STATE, AFTER_NUMBER_STATE, INVALID_STATE, INVALID_STATE},
		{INVALID_STATE, NUMBER_STATE, DOUBLE_COLON_STATE, INVALID_STATE},
//		{INVALID_STATE,	INVALID_STATE,INVALID_STATE,INVALID_STATE,}, // 如果不允许::1 前导缩写情况
		{INVALID_STATE, AFTER_NUMBER_STATE, AFTER_COLON_STATE, END_STATE},
		{INVALID_STATE, AFTER_NUMBER_STATE, INVALID_STATE, INVALID_STATE},
	}
	state := START_STATE
	doubleColon := false
	hexCnt := 0
	colonCnt := 0
	for _, ch := range s {
		switch {
		case ch == ' ':
			state = FSMTable[state][SPACE]
			break
		case ch == ':':
			colonCnt++
			if doubleColon || hexCnt > 4 || colonCnt > 7 {
				state = INVALID_STATE
			} else {
				state = FSMTable[state][COLON]
				hexCnt = 0
			}
			break
		case (ch >= '0' && ch <= '9') || (ch >= 'A' && ch <= 'F') || (ch >= 'a' && ch <= 'f'):
			hexCnt++
			state = FSMTable[state][HEXNUM]
		default:
			state = INVALID_STATE
		}
		//fmt.Printf("%c %v \n", ch, state)
		if state == INVALID_STATE {
			return false
		}
	} // 判断结束符号
	//fmt.Println(state, FSMTable[state][END], colonCnt)
	if state == AFTER_NUMBER_STATE && colonCnt > 6 {
		return false
	}
	if FSMTable[state][END] != END_STATE || hexCnt > 4 {
		return false
	}

	return true
}
