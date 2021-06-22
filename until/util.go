package until

import (
	"fmt"
	"math"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	pid = os.Getpid()
)

// CalcBitValueByString 计算字符串位运算后的值(字符串以逗号分隔)
func CalcBitValueByString(s string) *int {
	val := 0
	if s == "" {
		return &val
	}

	for _, v := range strings.Split(s, ",") {
		val = val | 1<<S(v).DefaultUint(0)
	}

	return &val
}

// ConvStringToFloatInt 转换字符串为浮点数的整数（b为需要乘的基数）
func ConvStringToFloatInt(s string, b int) *int {
	var v int
	f, err := S(s).Float64()
	if err != nil {
		return &v
	}

	if b > 0 {
		v = int(f * float64(b))
	} else {
		v = int(f)
	}

	return &v
}

// FormatNumberString 格式化数字字符串
func FormatNumberString(s string, f int) string {
	if i := strings.Index(s, "."); i > -1 {
		if len(s[i:]) > f+1 {
			return s[:i+f+1]
		}
	}
	return s
}

// ContentDisposition implements a simple version of https://tools.ietf.org/html/rfc2183
// Use mime.ParseMediaType to parse Content-Disposition header.
func ContentDisposition(fileName, dispositionType string) (header string) {
	if dispositionType == "" {
		dispositionType = "attachment"
	}
	if fileName == "" {
		return dispositionType
	}

	header = fmt.Sprintf(`%s; filename="%s"`, dispositionType, url.QueryEscape(fileName))
	fallbackName := url.PathEscape(fileName)
	if len(fallbackName) != len(fileName) {
		header = fmt.Sprintf(`%s; filename*=UTF-8''%s`, header, fallbackName)
	}
	return
}

// FillZero 填充零
func FillZero(i int) string {
	if i < 10 {
		return fmt.Sprintf("0%d", i)
	}
	return fmt.Sprintf("%d", i)
}

// FracFloat 取小数
func FracFloat(f float64) float64 {
	if f == 0 {
		return 0
	}
	_, fr := math.Modf(f)
	return fr
}

// BoolToInt 布尔转int
func BoolToInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

// DecimalFloat64 保留两位小数
func DecimalFloat64(value float64, limit ...int) float64 {

	if len(limit) > 0 {
		switch limit[0] {
		case 3:
			value, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", value), 64)
		case 4:
			value, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", value), 64)
		default:
			value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
		}
	} else {
		value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	}

	return value
}

// StringToInt 字符串转换为Int
func StringToInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return int(i)
}

// Max 整型取大值
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func IsContainString(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
func IsContainInt(items []int, item int) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// Divide 除法(解决除零错误)
func Divide(dividend, divisor float64) float64 {
	if divisor == 0 {
		return 0.00
	}
	return DecimalFloat64(dividend / divisor)
}

func SortStringNumber(p *string) string {
	stringSlice := strings.Split(*p, ",")
	var intSlice []int
	for _, item := range stringSlice {
		num, err := strconv.Atoi(item)
		if err != nil {
			num = 0
		}
		intSlice = append(intSlice, num)
	}
	sort.Ints(intSlice)
	stringSlice = stringSlice[0:0]
	for _, item := range intSlice {
		str := strconv.Itoa(item)
		stringSlice = append(stringSlice, str)
	}
	*p = strings.Join(stringSlice, ",")
	return *p
}

func GenerateSN(size, num int) string {
	if num < 0 {
		num = 0
	}
	result := fmt.Sprint(num)
	for i := 0; len(result) < size && i < 100; i++ {
		result = fmt.Sprintf("0%s", result)
	}
	return result
}

// 获取 当前月的起始时间和下个月的起始时间
func GetCurrentMonthDateTime() (string, string) {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, 0)

	return firstOfMonth.String()[:19], lastOfMonth.String()[:19]
}

// Min 整型取小值
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// NewTraceID 创建追踪ID
func NewTraceID() string {
	return fmt.Sprintf("trace-id-%d-%s",
		pid,
		time.Now().Format("2006.01.02.15.04.05.999999"))
}

// ParseURL 修正URL 转义特殊字符
func ParseURL(url string) string {

	for i, v := range url {
		switch v {
		case '#':
			add := append([]byte{}, url[i+1:]...)
			url = url[:i] + `%23`
			url = url + string(add)
		default:
		}

	}
	return url
}
