package bitcal

import (
	"strconv"
	"strings"
)

func ParseByBitOrAfter(id int) string {
	var bitArr []string = ParseByBitOrAfterArr(id)
	if len(bitArr) == 0 {
		return ""
	}

	return strings.Join(bitArr, ",")
}

func ParseByBitOrAfterArr(id int) []string {
	var arr []string = make([]string, 0)
	for i := 0; i < id; i++ {
		tmp := pow(2, i)
		if id&tmp == tmp {
			arr = append(arr, strconv.Itoa(tmp))
		}
	}

	return arr
}

func pow(x float64, n int) int {
	if x == 0 {
		return 0
	}
	result := calPow(x, n)
	if n < 0 {
		result = 1 / result
	}

	return int(result)
}

func calPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	//向右移动一位
	result := calPow(x, n>>1)
	result *= result

	//如果n是奇数
	if n&1 == 1 {
		result *= x
	}

	return result
}
