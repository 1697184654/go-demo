package main

func A1() {
	println("A1")
}

func A2() {
	p := recover()
	println("A2")
	println(p)
}

func main() {
	defer A1()
	defer A2()
	panic(struct {
		Msg string
	}{Msg: "panic A"})
}
