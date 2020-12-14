package main
import (
	"fmt"
)

func test(arg string)error {
	fmt.Println(arg)
	return nil
}
func main() {
	f := test
	fmt.Printf("%T\n", f) //func(string) error
	f("f")

	f2 := func(arg string) {
		fmt.Println(arg)
	}
	f2("f2")

	//f3=0
	f3 := func(arg string) int{
		fmt.Println(arg)
		return 0
	}("f3")
	fmt.Println(f3)
}