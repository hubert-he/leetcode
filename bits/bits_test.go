package bits

import "testing"

func Test_reverseBits(t *testing.T) {
	for caseID, testCase := range []struct{
		num uint32
		want uint32
	}{
		{0B00000010100101000001111010011100,0B00111001011110000010100101000000},
		{0B11111111111111111111111111111101, 0B10111111111111111111111111111111},
	}{
		result := reverseBits(testCase.num)
		if result != testCase.want{
			t.Errorf("case-%d result = %d, but want %d", caseID, result, testCase.want)
		}
	}
}

func TestToHex(t *testing.T) {
	for caseID, testCase := range []struct{
		num int
		want string
	}{
		{26, "1a"},
		{-1, "ffffffff"},
	}{
		result := ToHex(testCase.num)
		if result != testCase.want{
			t.Errorf("case-%d result=%s but want=%s", caseID, result, testCase.want)
			break
		}
	}
}
