package ft

func PrintRune(r rune) {
	b := []byte(string(r))
	syscall.Write(1, b)
}