package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("World")
	if result != "Hello World" {
		//error
		t.Fail()
	}
}

func TestHelloWorldTegar(t *testing.T) {
	result := HelloWorld("Tegar")
	if result != "Hello Tegar" {
		//error
		t.Fail()
	}
}
