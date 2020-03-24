package util

import "strconv"

func FloatToString(number float64) string {
	return strconv.FormatFloat(number, 'f', 6, 64)
}
