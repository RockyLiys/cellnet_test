package main
import (
	"fmt"
)

type f1 interface {
	fun1()
}
type node1 struct {
	
}
func (n *node1)fun1() {
	fmt.Println("fun1")
}

func (n *node1)Anonymous(arg string) int{
	fmt.Printf("%s\n", arg)
	return 0
}

//注意，返回的是f1接口类型
func NewNode() f1 {
	return &node1{}
}

func main1() {
	f := NewNode()
	f.fun1()

	/*
	这是一个匿名接口,因为结构体node1实现了Anonymous方法，所以f1接口可以转换为这个匿名接口
	interface {
		Anonymous(arg string) int
	}
	*/
	//1
	f.(interface {
		Anonymous(arg string) int
	}).Anonymous("qwe")

	//2
	fx := f.(interface {
		Anonymous(arg string) int
	})
	fmt.Println(fx.Anonymous("qwe"))
}

func main() {

	f := struct {
		id int
		name string
	}{
		id : 1,
		name : "qwe",
	}
	fmt.Printf("%+v\n", f)

}