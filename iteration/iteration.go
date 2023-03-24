package iteration

func Repeat(character string, number_iter int) string {
	var repeated string
	for i := 0; i < number_iter; i++ {
		repeated += character
	}
	return repeated
}