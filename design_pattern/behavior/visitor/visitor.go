package visitor

import (
	"fmt"
)

type Action interface {
	GetManConclusion(p *Man)
	GetWomanConclusion(p *Woman)
}

type Person interface {
	Accept(visitor Action)
}

type Success struct {
	name string
}

func (s *Success) GetManConclusion(p *Man) {
	fmt.Printf("%s%s的时候，背后多半有一个伟大的女人。\n", p.name, s.name)
}

func (s *Success) GetWomanConclusion(p *Woman) {
	fmt.Printf("%s%s的时候，背后多半有一个不成功的男人。\n", p.name, s.name)
}

type Failing struct {
	name string
}

func (s *Failing) GetManConclusion(p *Man) {
	fmt.Printf("%s%s的时候，低头喝闷酒，谁也不用劝。\n", p.name, s.name)
}

func (s *Failing) GetWomanConclusion(p *Woman) {
	fmt.Printf("%s%s的时候，眼泪汪汪，谁也劝不了。\n", p.name, s.name)
}

type Man struct {
	name string
}

func (m *Man) Accept(visitor Action) {
	visitor.GetManConclusion(m)
}

type Woman struct {
	name string
}

func (w *Woman) Accept(visitor Action) {
	visitor.GetWomanConclusion(w)
}

type ObjectStructure struct {
	persons []Person
}

func (o *ObjectStructure) Attach(p Person) {
	o.persons = append(o.persons, p)
}

func (o *ObjectStructure) Display(visitor Action) {
	for _, p := range o.persons {
		p.Accept(visitor)
	}
}
