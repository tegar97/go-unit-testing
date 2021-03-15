package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("World")
	if result != "Hello World" {
		//error
		t.Error("Error : Result Must Be World")
	}
	fmt.Println("finish")

}

func TestHelloWorldTegar(t *testing.T) {
	result := HelloWorld("sa")
	if result != "Hello Tegar" {
		//error
		t.Fatal("Error : Result Must Be Tegar")

	}
	fmt.Println("finish")
}
