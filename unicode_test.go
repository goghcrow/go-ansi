package ansi

import (
	"fmt"
	"testing"
)

func TestTransform(t *testing.T) {
	fmt.Println(Green.Text(Transform("Struck-Double", "Hello, World")))

	fmt.Println(Purple.Text(Transform("Capitalized", "The Universal Answer is 42")))

	fmt.Println(Red.Text(Transform("SansSerif-Bold", "The Universal Answer is 42")))
}
