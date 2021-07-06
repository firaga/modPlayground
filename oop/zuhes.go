package main

import "fmt"

type pp struct {
	logo string
}

func (p pp) name() {
	fmt.Println(p.logo)
}

type mm struct {
	logo string
}

func (m mm) name() {
	fmt.Println(m.logo)
}

type son struct {
	pp
	mm
	logo string
}

func (s son) name() {
	fmt.Println(s.logo)
}

func main() {
	s := son{
		pp:   pp{logo: "pp logo"},
		mm:   mm{logo: "mm logo"},
		logo: "son logo",
	}
	s.name()
	s.pp.name()
	s.mm.name()
}
