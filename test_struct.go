package main

import (
	"fmt"
)

type Test struct {
	Num int
}

func (t *Test) Print() {
	fmt.Println(t)
}

func (t *Test) Twice() {
	t.Num *= 2
}

func (t *Test) Add(val int) {
	t.Num += val
}

func (t *Test) GetNum() int {
	return t.Num
}

func (t *Test) AddAnotherTest(t2 *Test) {
	t.Num += t2.Num
}
