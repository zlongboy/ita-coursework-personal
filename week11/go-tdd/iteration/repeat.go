package iteration

func main() {
	Repeat("help")
}

func Repeat(c string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += c
	}
	return repeated
}
