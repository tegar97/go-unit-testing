package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	// BEFORE
	fmt.Println("Before Unit Test")

	m.Run()

	//after
	fmt.Println("After Unit Test")

}

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

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("World")
	//if result != "Hello World" {
	//	//error
	//	t.Error("Error : Result Must Be World")
	//}
	assert.Equal(t, "Hello F", result, "Result Must be 'Hello World'")
	fmt.Println("finish")

}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("OS TIDAK SUPPORT")
	}

	result := HelloWorld("World")
	assert.Equal(t, "Hello s", result, "Result Must be 'Hello World'")

}

func TestSubTest(t *testing.T) {
	t.Run("tegar", func(t *testing.T) {
		result := HelloWorld("tegar")
		require.Equal(t, "Hello Tegar", result, "Result must be 'Hello Tegar'")
	})
	t.Run("akmal", func(t *testing.T) {
		result := HelloWorld("akmal")
		require.Equal(t, "Hello Tegar", result, "Result must be 'Hello akmal'")
	})

}
