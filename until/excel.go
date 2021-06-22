package until

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

type ExcelPosition struct {
	X int //列数
	Y int //行数
}

type ExcelTemp struct {
	End  string
	Data map[string]ExcelData
	List map[string]ExcelListTemp
}

//Excel坐标
type ExcelListTemp struct {
	Start string //开始位置
	End   string //结束位置（方形）
	Data  map[string]ExcelData
}

// ExcelData 数单元包括
type ExcelData struct {
	Position string
	DataFunc func(data interface{}) interface{}
}

func (a *ExcelData) GetData(d interface{}) interface{} {
	if a.DataFunc == nil {
		return d
	}
	return a.DataFunc(d)
}

func ExcelWriteSheet(sheet *xlsx.Sheet, start string, temp ExcelTemp, data interface{}) string {
	sp := NewPosition(start) //起始坐标
	ep := NewPosition(start) //结束坐标
	xAdd, yAdd := 0, 0       //坐标补充值

	dataM := analysisData(data, "json")
	//添加定点数据（不影响补充坐标）
	for k, v := range temp.Data {
		dp := sp.AddString(v.Position)
		writeInPosition(sheet, dp, v.GetData(dataM[k]))
		ep.SetEnd(dp)
	}
	//添加列表数据
	for k, v := range temp.List {
		ya := NewPosition(v.End).Y - NewPosition(v.Start).Y + 1 //单条数据行增量
		if ya < 1 {
			ya = 1
		}
		d := dataM[k]
		if d != nil {
			dd, ok := d.([]map[string]interface{})
			if ok {
				if len(dd) > 1 && v.Start != "" { //插入新行
					copyData(sheet, v.Start, v.End, len(dd)-1)
				}
				for _, w := range dd {
					for j, z := range v.Data {
						dp := sp.AddPostion(xAdd, yAdd).AddString(z.Position)
						writeInPosition(sheet, dp, z.GetData(w[j]))
						ep.SetEnd(dp)
						yAdd += ya
					}

				}
			}
		}
	}
	return ep.ToString()
}

//向坐标内填充数据
func writeInPosition(sheet *xlsx.Sheet, position *ExcelPosition, data interface{}) {
	//确认行数
	row := sheet.Row(position.Y)
	for i := position.X - len(row.Cells) + 1; i > 0; i-- {
		row.AddCell()
	}
	row.Cells[position.X].SetValue(data)
}

//向坐标内填充数据
func WriteInPosition(sheet *xlsx.Sheet, ps string, data interface{}) {
	p := NewPosition(ps)
	writeInPosition(sheet, p, data)
}

//复制数据
func copyData(sheet *xlsx.Sheet, start, end string, n int) {
	if start == "" { //起始位置为空则不进行复制
		return
	}
	rs := 1 //单次行数
	sp := NewPosition(start)
	ep := NewPosition(end)
	if end == "" || ep.Y < sp.Y { //仅加行
		for i := n; i > 0; i-- {
			sheet.AddRowAtIndex(sp.Y + 1)
		}
		return
	}
	rs = ep.Y - sp.Y + 1

	//补全模板漏洞
	for j := 1; j <= rs; j++ {
		sheet.Row(ep.Y + j)
		for k := ep.X - len(sheet.Row(sp.Y+j-1).Cells) + 1; k > 0; k-- {
			sheet.Row(sp.Y + j - 1).AddCell()
		}
	}

	for i := n; i > 0; i-- {
		//加行
		for j := 1; j <= rs; j++ {

			row, _ := sheet.AddRowAtIndex(ep.Y + j)

			for i := ep.X - len(row.Cells) + 1; i > 0; i-- {
				row.AddCell()
			}
			for k := sp.X; k <= ep.X; k++ {
				row.Cells[k].SetString(sheet.Row(sp.Y + j - 1).Cells[k].Value)
			}
		}
	}
}

//新建坐标
func NewPosition(s string) *ExcelPosition {
	x, y := getPosition(s)
	return &ExcelPosition{
		X: x,
		Y: y,
	}
}

//移动坐标
func (a *ExcelPosition) AddString(s string) *ExcelPosition {
	if s == "" {
		return a
	}
	x, y := getPosition(s)
	return &ExcelPosition{
		X: a.X + x,
		Y: a.Y + y,
	}
}
func (a *ExcelPosition) AddPostion(x, y int) *ExcelPosition {
	if x == 0 && y == 0 {
		return a
	}
	return &ExcelPosition{
		X: a.X + x,
		Y: a.Y + y,
	}
}

func (a *ExcelPosition) SetEnd(e *ExcelPosition) {
	if e.X >= a.X {
		a.X = e.X + 1
	}
	if e.Y >= a.Y {
		a.Y = e.Y + 1
	}
}

//坐标转文字
func (a *ExcelPosition) ToString() string {
	return getPositionString(a.X, a.Y)
}

//获取Excel位置INDEX
func getPosition(s string) (int, int) {
	if s == "" {
		return 0, 0
	}
	var xs, ys string
	var x, y int

	for i, v := range []byte(s) {
		if v < 58 {
			xs = s[:i]
			ys = s[i:]
			break
		}
	}
	for i, v := range []byte(xs) {
		j := int(v) - 64
		for z := len(xs) - i - 1; z > 0; z-- {
			j = j * 26
		}
		x += j
	}
	if x <= 0 {
		x = 1
	}
	y, _ = strconv.Atoi(ys)
	if y <= 0 {
		y = 1
	}
	return x - 1, y - 1
}

func getPositionString(x, y int) string {
	x++
	y++
	var xs, ys string
	for i := 0; x > 0 && i < 10; i++ {
		a := x / 26
		b := x % 26
		if b == 0 {
			a--
			b = 26
		}
		if a == 0 {
			xs += string([]byte{byte(b + 64)})
			break
		} else {
			x = b
			xs += string([]byte{byte(a + 64)})
		}
	}

	ys = strconv.Itoa(y)
	return xs + ys
}

// analysisData分析数据
func analysisData(data interface{}, tag string) map[string]interface{} {

	var (
		ty     reflect.Type
		va     reflect.Value
		result = map[string]interface{}{}
	)

	switch reflect.TypeOf(data).Kind() {
	case reflect.Ptr:
		ty = reflect.TypeOf(data).Elem()
		va = reflect.ValueOf(data).Elem()
	case reflect.Struct:
		ty = reflect.TypeOf(data)
		va = reflect.ValueOf(data)
	default:
		return result
	}

	for i := 0; i < ty.NumField(); i++ {
		var key string
		if tag != "" {
			key = ty.Field(i).Tag.Get(tag)
			key = strings.Split(key, ",")[0]
		} else {
			key = ty.Field(i).Name
		}

		switch va.Field(i).Kind() {
		case reflect.Slice:
			d := va.Field(i)
			data := []map[string]interface{}{}
			for j := 0; j < d.Len(); j++ {
				data = append(data, analysisData(d.Index(j).Interface(), tag))
			}
			result[key] = data
		default:
			result[key] = va.Field(i).Interface()
		}

	}
	return result
}
