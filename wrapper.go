package main

import (
	"strconv"
	"syscall/js"
	"unsafe"
)

func (t *Test) PrintWrapper(this js.Value, args []js.Value) interface{} {
	t.Print()
	return nil
}

func (t *Test) TwiceWrapper(this js.Value, args []js.Value) interface{} {
	t.Twice()
	return nil
}

func (t *Test) AddWrapper(this js.Value, args []js.Value) interface{} {
	t.Add(args[0].Int())
	return nil
}

func (t *Test) GetNumWrapper(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(t.GetNum())
}

func (t *Test) getPtr(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(uintptr(unsafe.Pointer(t)))
}

func (t *Test) AddAnotherTestWrapper(this js.Value, args []js.Value) interface{} {
	ptrJSValue := args[0].Call("_ptr")
	uintVal, _ := strconv.ParseUint(ptrJSValue.String(), 10, 64)
	t2 := (*Test)(unsafe.Pointer(uintptr(uintVal)))
	t.AddAnotherTest(t2)
	return nil
}

func registerCallbacks() {
	var test = &Test{
		Num: 1,
	}
	js.Global().Set("test", js.ValueOf(
		map[string]interface{}{
			"_ptr":           js.FuncOf(test.getPtr),
			"print":          js.FuncOf(test.PrintWrapper),
			"twice":          js.FuncOf(test.TwiceWrapper),
			"add":            js.FuncOf(test.AddWrapper),
			"getNum":         js.FuncOf(test.GetNumWrapper),
			"addAnotherTest": js.FuncOf(test.AddAnotherTestWrapper),
		},
	))

	var test2 = &Test{
		Num: 2,
	}
	js.Global().Set("test2", js.ValueOf(
		map[string]interface{}{
			"_ptr":           js.FuncOf(test2.getPtr),
			"print":          js.FuncOf(test2.PrintWrapper),
			"twice":          js.FuncOf(test2.TwiceWrapper),
			"add":            js.FuncOf(test2.AddWrapper),
			"getNum":         js.FuncOf(test2.GetNumWrapper),
			"addAnotherTest": js.FuncOf(test2.AddAnotherTestWrapper),
		},
	),
	)
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
