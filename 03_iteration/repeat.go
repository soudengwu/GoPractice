package iteration

// Repeat a string character a RepeatCount Number of times
func Repeat(char string, repeatCount int) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += char
	}
	return repeated
}
