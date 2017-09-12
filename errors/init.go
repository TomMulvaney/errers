package errors

func init() {
	fromHTTP = make(map[int]int, len(toHTTP))

	for k, v := range toHTTP {
		fromHTTP[v] = k
	}
}
