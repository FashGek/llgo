package main

var hi = 0xFF00
var lo = 0xFF00

func main() {
	println(hi & 0x1000)
	println(hi & 0x0100)
	println(hi & 0x0010)
	println(hi & 0x0001)
	println(lo & 0x1000)
	println(lo & 0x0100)
	println(lo & 0x0010)
	println(lo & 0x0001)
	println(hi | lo)
	println(hi ^ hi)
	println(hi ^ lo)
	println(lo ^ lo)
	println(hi ^ 0x1000)
	println(hi ^ 0x0100)
	println(hi ^ 0x0010)
	println(hi ^ 0x0001)
	println(lo ^ 0x1000)
	println(lo ^ 0x0100)
	println(lo ^ 0x0010)
	println(lo ^ 0x0001)
	println(-123 >> 1)
	println(-123 << 1)
}

