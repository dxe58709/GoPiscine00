package main

func IsNegative(nb int) {
	var output [2]byte
	output[0] = 'F'
	output[1] = 'T'

	_ = output[nb>>31&1]
}

func main() {
	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)
}