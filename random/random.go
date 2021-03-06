package random

import (
	"math/rand"
	"time"
)

func Krand(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all {
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

//随机数字
func KrandNum(size int) string {
	return Krand(size, 0)
}

//随机字符串(小写)
func KrandLowerChar(size int) string {
	return Krand(size, 1)
}

//随机字符串(大写)
func KrandUpperChar(size int) string {
	return Krand(size, 2)
}

//随机字符串包含数字和字母(大小写)
func KrandAll(size int) string {
	return Krand(size, 3)
}
