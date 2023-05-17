package main

import "testing"

func TestMethod(t *testing.T) {
	b := Book{Name: "The Go Programming Language", Author: "Alan A. A. Donovan & Brian W. Kernighan"}
	want := "The Go Programming Language by Alan A. A. Donovan & Brian W. Kernighan"
	if b.String() != want {
		t.Errorf("Book.String() = %q, want %q", b.String(), want)
	}
}

func TestSetName(t *testing.T) {
	b := Book{Name: "The Go Programming Language", Author: "Alan A. A. Donovan & Brian W. Kernighan"}
	want := "The Go Programming Language (Edit) by Alan A. A. Donovan & Brian W. Kernighan"
	b.SetName("The Go Programming Language (Edit)")
	if b.String() != want {
		t.Errorf("Book.String() = %q, want %q", b.String(), want)
	}
}
