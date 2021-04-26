package verify

import (
	"regexp"
)

//判断时间格式是否正确 - xxxx-xx-xx 00:00:00
func VerifyTimeDate(content string) bool {
   reg := regexp.MustCompile("^[1-9]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])\s+(20|21|22|23|[0-1]\d):[0-5]\d:[0-5]\d$")
   return reg.Match([]byte(content))
}

//验证手机号码
func VerifyMobile(content string) bool {
   reg := regexp.MustCompile("^[1][3-9][0-9]{9}$")
   return reg.Match([]byte(content))
}

//验证身份证号码
func VerifyIdcard(content string) bool {
   reg := regexp.MustCompile("^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$")
   return reg.Match([]byte(content))
}

//验证邮箱
func VerifyEmail(content string) bool {
   reg := regexp.MustCompile("^[^\@]+@.*\.[a-z]{2,6}$")
   return reg.Match([]byte(content))
}


//验证价格
func VerifyPrice(content string) bool {
   reg := regexp.MustCompile("(^[1-9]([0-9]+)?(\.[0-9]{1,2})?$)|(^(0){1}$)|(^[0-9]\.[0-9]([0-9])?$)")
   return reg.Match([]byte(content))
}

//验证网络地址
func VerifyUrl(content string) bool {
   reg := regexp.MustCompile("http(s)?:\/\/([\w-]+\.)+[\w-]+(\/[\w- .\/?%&=]*)?")
   return reg.Match([]byte(content))
}

//验证银行卡
func VerifyBankCode(content string) bool {
   reg := regexp.MustCompile("^\d{16,21}$")
   return reg.Match([]byte(content))
}

//验证整数
func VerifyInteger(content string) bool {
   reg := regexp.MustCompile("^-?\\d+$")
   return reg.Match([]byte(content))
}
