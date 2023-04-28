package changetool

import "strconv"

// InterfaceToString 接口转字符串
func InterfaceToString(i interface{}) string {
	if s, ok := i.(string); ok {
		return s
	}
	return ""
}

// IntToString int 转字符串
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// StringToInt 字符串转int
func StringToInt(str string) (int, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// StringToInt64 字符串转int64
func StringToInt64(str string) (int64, error) {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, err
}

// Int64ToString int64转字符串
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
