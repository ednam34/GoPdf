package backend

import "strconv"

func IntSliceToStr(tab []int) []string {
	strArray := make([]string, len(tab))

	for i, num := range tab {
		strArray[i] = strconv.Itoa(num)
	}

	return strArray
}
