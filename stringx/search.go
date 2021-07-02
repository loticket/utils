package stringx

import "sort"

//@action 二分查找算法
//@param  sortedList 整数数组 lookingFor 需要查找的整数
//@return  -1 没有查找到  其他的返回为 查找的key值
func binarySearch(sortedList []int, lookingFor int) int {
	var lo int = 0
	var hi int = len(sortedList) - 1
	sort.Ints(sortedList)
	for lo <= hi {
		var mid int = lo + (hi-lo)/2
		var midValue int = sortedList[mid]
		if midValue == lookingFor {
			return midValue
		} else if midValue > lookingFor {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}



//@action 排序查找算法
//@param  sint 整数数组 index 需要查找的整数
//@return false 没有查找到  true 查询到
func SearchInt(sint []int, index int) bool {
	sort.Ints(sint)
	pos := sort.Search(len(sint), func(i int) bool { return sint[i] >= index })
	if pos < len(sint) && sint[pos] == index {
		return true
	} else {
		return false
	}
}


//@action查询数组中是否包含某一个值 -- 字符串查找
//@param  strArr 字符串切片 index 需要查找的字符串
//@return false 没有查找到  true 查询到
func SearchString(strArr []string, index string) bool {
	sort.Strings(strArr)
	pos := sort.Search(len(strArr), func(i int) bool { return strArr[i] >= index })
	if pos < len(strArr) && strArr[pos] == index {
		return true
	} else {
		return false
	}
}


//@action 字符串切片去重
//@param  a 字符串切片
//@return  去重后的切片
func ArrayUnique(a []string) (ret []string) {
	sort.Strings(a)
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

//@action 比较两个切片是否相同
//@param  a 整形切片  b 整形切片
//@return  true 相同  false 不相同
func SliceEqualInt(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

//@action 比较两个切片是否相同
//@param  a interface切片  b interface切片
//@return  true 相同  false 不相同
func SliceEqualInterface(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}