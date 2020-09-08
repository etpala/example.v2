package SubClass

import "fmt"

type Phone interface {
	Method0(inParm string) int
	method1(inParm int) int
}

type Nokia struct {
	Name string
	dept string
}

func SubMethod1() {
	fmt.Println("Call SubMethod1")
	subMethod2("SubMethod1")
	return
}

func subMethod2(from string) {
	fmt.Println("Call subMethod2", from)
	return
}

func Method0(inParm string) int {
	fmt.Println("call interface Methon0", inParm)
	return 0
}

func method1(inParm string) int {
	fmt.Println("call interface methon1", inParm)
	return 0
}
