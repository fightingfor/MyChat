package util

import "strconv"

func Str2int64(str string) int64 {

	i, _ := strconv.ParseInt(str, 10, 64)

	return i
}

func Str2int(str string) int {

	i, _ := strconv.Atoi(str)
	return i
}

func Int2Str(i int64) string {

	formatInt := strconv.FormatInt(i, 10)
	return formatInt
}

func Int642Str(i int) string {
	itoa := strconv.Itoa(i)

	return itoa
}
