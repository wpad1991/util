package convert

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func ConvertStringList(pval [][]interface{}) [][]string {

	datalist := make([][]string, len(pval))

	if pval != nil {
		for i, row := range pval {

			strlist := make([]string, len(row))
			for j, col := range row {

				defer func() {
					if r := recover(); r != nil {
						panic("convertStringList unknown type value i : " + strconv.Itoa(i) + ", j : " + strconv.Itoa(j))
					}
				}()

				strlist[j] = ConvertInterfacetoString(col)
			}

			datalist[i] = strlist
		}
	}

	return datalist
}

func ConvertInterfacetoString(value interface{}) string {

	switch v := value.(type) {
	case nil:
		return ""
	case bool:
		if v {
			return "true"
		}
		return "false"
	case int:
		return strconv.Itoa(value.(int))
	case int8:
		return strconv.Itoa(int(value.(int8)))
	case int16:
		return strconv.Itoa(int(value.(int16)))
	case int32:
		return strconv.Itoa(int(value.(int32)))
	case int64:
		return strconv.FormatInt(value.(int64), 10)
	case uint:
		return strconv.Itoa(int(value.(uint)))
	case uint8:
		return strconv.Itoa(int(value.(uint8)))
	case uint16:
		return strconv.Itoa(int(value.(uint16)))
	case uint32:
		return strconv.FormatInt(int64(value.(uint32)), 10)
	case uint64:
		return strconv.FormatUint(value.(uint64), 10)
	case float32:
		return strconv.FormatFloat(float64(value.(float32)), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(value.(float64), 'f', -1, 64)
	case []byte:
		return string(value.([]byte))
	case string:
		return value.(string)
	case time.Time:
		return v.Format("2006-01-02 15:04:05.999")
	default:
		println(reflect.TypeOf(v))
		fmt.Printf("QWE? : %v", v)
		panic("convertStringList unknown type value")
	}
}
