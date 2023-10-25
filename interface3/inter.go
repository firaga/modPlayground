package main

type T1 struct{}

func (T1) T1M1()   { println("T1's M1") }
func (*T1) PT1M2() { println("PT1's M2") }

type T2 struct{}

func (T2) T2M1()   { println("T2's M1") }
func (*T2) PT2M2() { println("PT2's M2") }

type T struct {
	T1
	*T2
}

func main() {
	t := T{
		T1: T1{},
		T2: &T2{},
	}
	pt := &t
	t.PT1M2()
	pt.PT1M2()
}
