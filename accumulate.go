package accumulate

func Accumulate(given []string, f func(string) string) []string {

	acc := make([]string, len(given))

	for i, v := range given {

		acc[i] = f(v)
	}

	return acc
}
