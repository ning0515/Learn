package iteration

func Repeat(a string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += a
	}
	return repeated
}
