package until

import (
	"regexp"
	"strconv"
	"time"
)

// S 字符串类型转换
type S string

func (s S) String() string {
	return string(s)
}

// Bytes 转换为[]byte
func (s S) Bytes() []byte {
	return []byte(s)
}

// Bool 转换为bool
func (s S) Bool() (bool, error) {
	b, err := strconv.ParseBool(s.String())
	if err != nil {
		return false, err
	}
	return b, nil
}

// DefaultBool 转换为bool，如果出现错误则使用默认值
func (s S) DefaultBool(defaultVal bool) bool {
	b, err := s.Bool()
	if err != nil {
		return defaultVal
	}
	return b
}

// Int64 转换为int64
func (s S) Int64() (int64, error) {
	i, err := strconv.ParseInt(s.String(), 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// DefaultInt64 转换为int64，如果出现错误则使用默认值
func (s S) DefaultInt64(defaultVal int64) int64 {
	i, err := s.Int64()
	if err != nil {
		return defaultVal
	}
	return i
}

// Int 转换为int
func (s S) Int() (int, error) {
	i, err := s.Int64()
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

// DefaultInt 转换为int，如果出现错误则使用默认值
func (s S) DefaultInt(defaultVal int) int {
	i, err := s.Int()
	if err != nil {
		return defaultVal
	}
	return i
}

// DefaultInt 转换为int，如果出现错误则使用默认值
func (s S) DefaultPInt(defaultVal *int) *int {
	i, err := s.Int()
	if err != nil {
		return defaultVal
	}
	return &i
}

// Uint64 转换为uint64
func (s S) Uint64() (uint64, error) {
	i, err := strconv.ParseUint(s.String(), 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// DefaultUint64 转换为uint64，如果出现错误则使用默认值
func (s S) DefaultUint64(defaultVal uint64) uint64 {
	i, err := s.Uint64()
	if err != nil {
		return defaultVal
	}
	return i
}

// Uint 转换为uint
func (s S) Uint() (uint, error) {
	i, err := s.Uint64()
	if err != nil {
		return 0, err
	}
	return uint(i), nil
}

// DefaultUint 转换为uint，如果出现错误则使用默认值
func (s S) DefaultUint(defaultVal uint) uint {
	i, err := s.Uint()
	if err != nil {
		return defaultVal
	}
	return uint(i)
}

// Float64 转换为float64
func (s S) Float64() (float64, error) {
	f, err := strconv.ParseFloat(s.String(), 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

// DefaultFloat64 转换为float64，如果出现错误则使用默认值
func (s S) DefaultFloat64(defaultVal float64) float64 {
	f, err := s.Float64()
	if err != nil {
		return defaultVal
	}
	return f
}

// Float32 转换为float32
func (s S) Float32() (float32, error) {
	f, err := s.Float64()
	if err != nil {
		return 0, err
	}
	return float32(f), nil
}

// DefaultFloat32 转换为float32，如果出现错误则使用默认值
func (s S) DefaultFloat32(defaultVal float32) float32 {
	f, err := s.Float32()
	if err != nil {
		return defaultVal
	}
	return f
}

// ToJSON 转换为JSON
func (s S) ToJSON(v interface{}) error {
	return json.Unmarshal(s.Bytes(), v)
}

func (s S) Time(forms ...string) *time.Time {
	var form string
	if len(forms) == 1 {
		form = forms[0]
	} else {
		form = time.RFC3339
		//form = `2006-01-02T15:04:05+08:00`
	}
	t, err := time.Parse(form, s.String())
	if err != nil || t.IsZero() {
		return nil
	}
	return &t
}

// 字符串转成时间对象
func (s S) TimeWithLoc(loc *time.Location, forms ...string) *time.Time {
	var form string
	if len(forms) == 1 {
		form = forms[0]
	} else {
		form = `2006-01-02T15:04:05.999Z`
	}

	if nil == loc {
		tempLoc, err := time.LoadLocation("Asia/Shanghai")
		if nil != err {
			return nil
		}
		loc = tempLoc
	}

	t, err := time.ParseInLocation(form, s.String(), loc)
	if err != nil || t.IsZero() {
		return nil
	}
	return &t
}

// 判断是不是电话号码（含手机或座机）
func (s S) MatchCellPhone() bool {
	matchPhone, err := regexp.MatchString("^(1[3-9][0-9])\\d{8}$|^\\d{8}$|^0\\d{2}-(\\d{7}|\\d{8})$|^0\\d{3}-(\\d{7}|\\d{8})$", s.String())
	if err != nil {
		return false
	}
	return matchPhone
}
