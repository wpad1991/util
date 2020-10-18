package convert

import (
	"testing"
)

func TestUnitTest(t *testing.T) {

	test := make([][]interface{}, 1)

	test[0] = make([]interface{}, 16)

	test[0][0] = uint8(1)
	test[0][1] = uint16(2)
	test[0][2] = uint32(3)
	test[0][3] = uint64(3)
	test[0][4] = uint(3)
	test[0][5] = int8(3)
	test[0][6] = int16(3)
	test[0][7] = int32(3)
	test[0][8] = int64(3)
	test[0][9] = int(3)
	test[0][10] = float32(3)
	test[0][11] = float64(3)
	test[0][12] = true
	test[0][13] = byte(3)
	test[0][14] = rune(3)
	test[0][15] = string("3")

	slist := ConvertStringList(test)

	for _, row := range slist {
		for _, col := range row {
			println(col)
		}
	}

}
