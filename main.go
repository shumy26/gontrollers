package main

func main() {
	ch := make(chan float64)
	go control(ch)
	plotting(ch)
}
