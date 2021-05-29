package utils

import "strconv"

func EchoCount(listSize int, item int) string {
	return "[" + strconv.Itoa(item+1) + "\\" + strconv.Itoa(listSize) + "] "
}
