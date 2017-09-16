package errors

func init() {
	FromHTTPMap = make(map[int]int, len(ToHTTPMap))

	for k, v := range ToHTTPMap {
		FromHTTPMap[v] = k
	}
}
