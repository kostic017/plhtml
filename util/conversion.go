package util

import "strconv"

func FloatToString(number float64) string {
    return strconv.FormatFloat(number, 'f', 6, 64)
}

func StrToInt64(str string) (int64, error) {
    return strconv.ParseInt(str, 10, 64)
}