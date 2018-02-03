package main

import (
	"flag"
	"os"
	"testing"
	"time"
)

type Ducker interface {
	Duck(int)
}

type Goose struct {
}

var sleep = 0

func TestMain(m *testing.M) {
	flag.IntVar(&sleep, "sleep", 0, "milliseconds to sleep (default 0)")
	os.Exit(m.Run())
}

func (g *Goose) Duck(i int) {
	if sleep > 0 {
		time.Sleep(time.Duration(i*sleep) * time.Millisecond)
	}
	return
}

func BenchmarkCallThroughPointerFunc(b *testing.B) {
	x := 1
	d := &Goose{}
	dF := d.Duck
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dF(x)
	}
}

func BenchmarkCallThroughInterfaceFunc(b *testing.B) {
	x := 1
	d := Ducker(&Goose{})
	dF := d.Duck
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dF(x)
	}
}

func BenchmarkCallThroughInterface(b *testing.B) {
	x := 1
	d := Ducker(&Goose{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Duck(x)
	}
}

func BenchmarkCallThroughPointer(b *testing.B) {
	d := &Goose{}
	b.ResetTimer()
	x := 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Duck(x)
	}
}

type EmbedDucker struct {
	Duck1 Ducker
	Duck2 *Goose
	f1    func(int)
}

func BenchmarkCallMemberInterface(b *testing.B) {
	x := 1
	e := EmbedDucker{Duck1: &Goose{}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.Duck1.Duck(x)
	}
}

func BenchmarkCallMemberPointer(b *testing.B) {
	x := 1
	e := EmbedDucker{Duck2: &Goose{}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.Duck2.Duck(x)
	}
}

func BenchmarkCallMemberFunc(b *testing.B) {
	x := 1
	e := EmbedDucker{f1: (&Goose{}).Duck}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.f1(x)
	}
}

func BenchmarkCallMemberFuncInterface(b *testing.B) {
	x := 1
	e := EmbedDucker{f1: Ducker(&Goose{}).Duck}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.f1(x)
	}
}
