package main

type Report struct {
	Title  string
	Values [1024]float64
}

type Point struct {
	X, Y float64
}

//go:noinline
func (r Report) FirstValue() float64 { return r.Values[0] }

//go:noinline
func (r *Report) FirstPointer() float64 { return r.Values[0] }

//go:noinline
func (p Point) PointValue() float64 { return p.X }

//go:noinline
func (p *Point) PointPointer() float64 { return p.X }