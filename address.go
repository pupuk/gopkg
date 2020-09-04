package address

import (
	"fmt"
	"regexp"
	"strings"
)

// Decompose a input string to map[idn, mobile, postcode, name, addr]
func Decompose(str string) map[string]string {
	m := make(map[string]string)

	//1. 过滤掉收货地址中的常用说明字符，排除干扰词
	str = strings.Replace(str, "身份证号", " ", -1)
	str = strings.Replace(str, "地址", " ", -1)
	str = strings.Replace(str, "收货人", " ", -1)
	str = strings.Replace(str, "收件人", " ", -1)
	str = strings.Replace(str, "收货", " ", -1)
	str = strings.Replace(str, "邮编", " ", -1)
	str = strings.Replace(str, "电话", " ", -1)
	str = strings.Replace(str, "身份证号码", " ", -1)
	str = strings.Replace(str, "身份证号", " ", -1)
	str = strings.Replace(str, "身份证", " ", -1)
	str = strings.Replace(str, "：", " ", -1)
	str = strings.Replace(str, ":", " ", -1)
	str = strings.Replace(str, "；", " ", -1)
	str = strings.Replace(str, ";", " ", -1)
	str = strings.Replace(str, "，", " ", -1)
	str = strings.Replace(str, ",", " ", -1)
	str = strings.Replace(str, "。", " ", -1)
	str = strings.Replace(str, ".", " ", -1)

	//2. 多个空白字符(包括空格\r\n\t)换成一个空格
	reg := regexp.MustCompile(`\s{1,}`)
	str = strings.TrimSpace(reg.ReplaceAllString(str, " "))

	//3. 去除手机号码中的短横线 如0136-3333-6666 主要针对苹果手机
	reg = regexp.MustCompile(`0-|0?(\d{3})-(\d{4})-(\d{4})`)
	str = reg.ReplaceAllString(str, "$1$2$3")

	//4. 提取中国境内身份证号码
	reg = regexp.MustCompile(`(?i)\d{18}|\d{17}X`)
	idnMatch := reg.FindString(str)
	str = strings.Replace(str, idnMatch, "", -1)
	m["idn"] = strings.ToUpper(idnMatch)

	//5. 提取11位手机号码或者7位以上座机号
	reg = regexp.MustCompile(`\d{7,11}|\d{3,4}-\d{6,8}`)
	mobile := reg.FindString(str)
	str = strings.Replace(str, mobile, "", -1)
	m["mobile"] = mobile

	//6. 提取6位邮编 邮编也可用后面解析出的省市区地址从数据库匹配出
	reg = regexp.MustCompile(`\d{6}`)
	postcode := reg.FindString(str)
	str = strings.Replace(str, postcode, "", -1)
	m["postcode"] = postcode

	//再次把2个及其以上的空格合并成一个，并首位TRIM
	reg = regexp.MustCompile(` {2,}`)
	str = strings.TrimSpace(reg.ReplaceAllString(str, " "))

	//7. 按照空格切分 长度长的为地址 短的为姓名 因为不是基于自然语言分析，所以采取统计学上高概率的方案
	r := strings.Split(str, " ")

	name := r[0]
	if len(r) > 1 {
		for i := 1; i < len(r); i++ {
			if len(r[i]) < len(name) {
				name = r[i]
			}
		}
	}
	m["name"] = name

	addr := strings.TrimSpace(strings.Replace(str, name, "", -1))
	m["addr"] = addr

	return m
}

// Parse a tring to a map
func Parse(address string) string {
	return "Parse"
}

// Smart fucniton, include decompose ,then Parse
func Smart(address string) string {
	return "Smart"
}

// Test then echo a string
func Test(str string) string {
	fmt.Println("This is test string")
	return "call success"
}
