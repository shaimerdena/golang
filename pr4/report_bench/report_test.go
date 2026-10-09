package main

import "testing"

var rep = Report{Title: "Large Report"}
var pt = Point{X: 1.0, Y: 2.0}

// Переменная Sink предотвращает удаление вызова функции оптимизатором компилятора
var Sink float64

func BenchmarkReportValue(b *testing.B) {
	var v float64
	for i := 0; i < b.N; i++ {
		v = rep.FirstValue()
	}
	Sink = v
}

func BenchmarkReportPointer(b *testing.B) {
	var v float64
	for i := 0; i < b.N; i++ {
		v = rep.FirstPointer()
	}
	Sink = v
}

func BenchmarkPointValue(b *testing.B) {
	var v float64
	for i := 0; i < b.N; i++ {
		v = pt.PointValue()
	}
	Sink = v
}

func BenchmarkPointPointer(b *testing.B) {
	var v float64
	for i := 0; i < b.N; i++ {
		v = pt.PointPointer()
	}
	Sink = v
}