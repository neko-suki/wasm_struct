package main

import (
	"fmt"
	"syscall/js"
)

type Test struct {
	Num int
}

func (t *Test) Print(this js.Value, args []js.Value) interface{} {
	fmt.Println(t)
	return nil
}

func (t *Test) Twice(this js.Value, args []js.Value) interface{} {
	t.Num *= 2
	return nil
}

func (t *Test) Add(this js.Value, args []js.Value) interface{} {
	t.Num += args[0].Int()
	return nil
}

func (t *Test) GetNum(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(t.Num)
}

func registerCallbacks() {
	var test = &Test{
		Num: 1,
	}
	js.Global().Set("test", js.ValueOf(
		map[string]interface{}{
			"print":  js.FuncOf(test.Print),
			"twice":  js.FuncOf(test.Twice),
			"add":    js.FuncOf(test.Add),
			"getNum": js.FuncOf(test.GetNum),
		},
	))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
