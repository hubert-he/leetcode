package unclassified

func RemoveDuplicatesString(S string) string {
	return removeDuplicatesString(S)
}


func removeDuplicates(S string) string {
	var data []byte = []byte(S)
	for true {
		flag := true
		for i := 0; i < len(data);{
			if i > 0 && data[i] == data[i-1] {
				data = append(data[:i-1], data[i+1:]...)
				flag = false
			} else {
				i++
			}
		}
		if flag {
			break
		}
	}
	return string(data[:])
}

func removeDuplicatesString(S string) string { // 栈实现
	st := []byte{}
	for i := 0; i < len(S); i++ {
		if len(st) == 0 || st[len(st) - 1] != S[i]{
			st = append(st, S[i])
		}else {
			st = st[:len(st) - 1]
		}
	}
	return string(st)
}