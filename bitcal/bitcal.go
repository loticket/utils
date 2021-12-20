package bitcal

import (
	"strconv"
	"strings"
)

//字符串数组求和
func AssignByBitOrAfter(ids string) int64 {
	var id int64 = 0
	idArr := strings.Split(strings.Trim(ids, ","), ",")
	for _, val := range idArr {
		tmp, _ := strconv.Atoi(val)
		id |= int64(tmp)
	}

	return id
}

//解析一个整数得到一个字符串，每个用,隔开
//该数组中全为2的倍数
func ParseByBitOrAfter(id int64) string {
	var bitArr []string = ParseByBitOrAfterArr(id)
	if len(bitArr) == 0 {
		return ""
	}

	return strings.Join(bitArr, ",")
}



func ParseByBitToString(id int64, maps map[string]string) string {
	var arr []string = make([]string, 0)
	var i int64 = 0
	for i = 0; i < id; i++ {
		var tmp int64 = pow(2, i)
		if id&tmp == tmp {
			var idStr string = strconv.Itoa(int(tmp))
			if v, ok := maps[idStr]; ok {
				arr = append(arr, v)
			}

		}
	}

	return strings.Join(arr, ",")
}

//解析一个整数得到一个字符串数组
//该数组中全为2的次方数 1，2，4，8，16，32，64，128，256
func ParseByBitOrAfterArrInt(id int64) []int64 {
	var arr []int64 = make([]int64, 0)
	var i int64 = 0
	for i = 0; i < id; i++ {
		var tmp int64 = pow(2, i)
		if id&tmp == tmp {
			arr = append(arr, tmp)
		}
	}

	return arr
}

//解析一个整数得到一个字符串数组
//该数组中全为2的次方数 1，2，4，8，16，32，64，128，256
func ParseByBitOrAfterArr(id int64) []string {
	var arr []string = make([]string, 0)
	var i int64 = 0
	for i = 0; i < id; i++ {
		var tmp int64 = pow(2, i)
		if id&tmp == tmp {
			arr = append(arr, strconv.Itoa(int(tmp)))
		}
	}

	return arr
}

func pow(x float64, n int64) int64 {
	if x == 0 {
		return 0
	}
	result := calPow(x, n)
	if n < 0 {
		result = 1 / result
	}

	return int64(result)
}

func calPow(x float64, n int64) float64 {
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
