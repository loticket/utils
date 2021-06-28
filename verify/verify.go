package verify

import (
	"regexp"
	"strconv"
	"strings"
)

//判断时间格式是否正确 - xxxx-xx-xx 00:00:00
func VerifyTimeDate(content string) bool {
	reg := regexp.MustCompile(`^[1-9]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])\s+(20|21|22|23|[0-1]\d):[0-5]\d:[0-5]\d$`)
	return reg.Match([]byte(content))
}

//验证手机号码
func VerifyMobile(content string) bool {
	reg := regexp.MustCompile(`^[1][3-9][0-9]{9}$`)
	return reg.Match([]byte(content))
}

//验证邮箱
func VerifyEmail(content string) bool {
	reg := regexp.MustCompile(`^[^\@]+@.*\.[a-z]{2,6}$`)
	return reg.Match([]byte(content))
}

//验证价格
func VerifyPrice(content string) bool {
	reg := regexp.MustCompile(`(^[1-9]([0-9]+)?(\.[0-9]{1,2})?$)|(^(0){1}$)|(^[0-9]\.[0-9]([0-9])?$)`)
	return reg.Match([]byte(content))
}

//验证网络地址
func VerifyUrl(content string) bool {
	reg := regexp.MustCompile(`http(s)?:\/\/([\w-]+\.)+[\w-]+(\/[\w- .\/?%&=]*)?`)
	return reg.Match([]byte(content))
}

//验证银行卡
func VerifyBankCode(content string) bool {
	reg := regexp.MustCompile(`^\d{16,21}$`)
	return reg.Match([]byte(content))
}

//验证整数
func VerifyInteger(content string) bool {
	reg := regexp.MustCompile(`^-?\\d+$`)
	return reg.Match([]byte(content))
}

// 身份证号正确性检查
func VerifyIdcard(idCard string) bool {
	idCard = strings.ToUpper(idCard)

	reg := regexp.MustCompile(`^[0-9]{17}[0-9X]$`)
	if reg.MatchString(idCard) == false {
		return false
	}
	return checkIdCardCode(idCard)
}

// 身份证校验码的计算方法：
//  1、将身份证号码前面的17位数分别乘以不同的加权因子，见： weights
//  2、将这17位数字和加权因子相乘的结果相加，得到的结果再除以11，得到余数 m
//  3、余数m作为位置值，在校验码数组 codes 中找到对应的值，就是身份证号码的第18位校验码
func checkIdCardCode(id string) bool {
	var weights []int    = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var codes   []string = []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}

	var sum int = 0
	for i := 0; i < 17; i++ {
		n, _ := strconv.Atoi(string(id[i]))
		sum += n * weights[i]
	}

	m := sum % 11

	return codes[m] == id[17:]
}