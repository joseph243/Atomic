package main

type Element struct {
	AtomicNumber int
	Name         string
	FreezingTemp int
	BoilingTemp  int
}

func newElement(num int, name string, freeze int, boil int) *Element {
	e := Element{AtomicNumber: num, Name: name, FreezingTemp: freeze, BoilingTemp: boil}
	return &e
}
