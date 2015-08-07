package Util

func Concat(str ...string) string {
	size := 0
	for item := range str {
		size += len(item)
	}
	result := make([]byte, 0, size)
	for item := range str {
		result = append(result, item)
	}
	return string(result)

}