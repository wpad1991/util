package slice

import (
	"math"
)

const epsilon = 1e-14

func Contains(slice interface{}, value interface{}) bool {

	switch slice.(type) {
	case nil:
		panic("SliceContains slice is nil")
	case bool:
	case int:
	case int64:
	case float32:
	case float64:
	case string:
	case []byte:
		conv := slice.([]byte)
		comp := value.(byte)
		for _, val := range conv {
			if val == comp {
				return true
			}
		}
	case []int:
		conv := slice.([]int)
		comp := value.(int)
		for _, val := range conv {
			if val == comp {
				return true
			}
		}
	case []int64:
		conv := slice.([]int64)
		comp := value.(int64)
		for _, val := range conv {
			if val == comp {
				return true
			}
		}
	case []float32:
		conv := slice.([]float64)
		comp := value.(float64)
		for _, val := range conv {
			if math.Abs(val-comp) <= epsilon {
				return true
			}
		}
	case []float64:
		conv := slice.([]float64)
		comp := value.(float64)
		for _, val := range conv {
			if math.Abs(val-comp) <= epsilon {
				return true
			}
		}
	case []string:
		conv := slice.([]string)
		comp := value.(string)
		for _, val := range conv {
			if val == comp {
				return true
			}
		}
	default:
	}

	return false
}
