/*
1.结构体嵌套组合共同实现一个接口，子类结构体指针可以赋值给接口类型吗？<可以>
2.结构体内嵌一个匿名接口，另一个结构体指针赋值给这个匿名接口，
那本结构体如何调用这个匿名接口的方法？<通过已有的接口变量转换成这个匿名接口>
*/
/*
匿名结构体与匿名接口

*/

package main
import (
	"fmt"
)

type f1 interface {
	func1()
	func2()
}
type f2 interface {
	myfun()
}

//基类
type base struct {
	b_id int
	f f1 //接口变量
}
//派生类
type son struct {
	base
	f2 //匿名接口, 将被初始化为 &other{}
	
	s_id int
}
//第三方结构体
type other struct {
	
}
//实现f2接口的方法
func (b *other)myfun() {
	fmt.Println("myfun")
}
//base、son分别实现f1接口的方法
func (b *base)func1() {
	fmt.Println("func1")
}
func (s *son)func2() {
	fmt.Println("func2")
}


func main() {
	//o := &other{}

	//实例化派生类对象
	s := &son{
		f2 : &other{},
	}
	s.f = s //s的f接口变量赋值为自身s
	s.func1()
	s.func2()

	//测试结构体指针赋值给接口
	var f f1
	f = s 
	f.func1()
	f.func2()

	//测试结构体内部接口变量的转换，调用匿名接口的方法
	//o.myfun()
	s.f.(f2).myfun()
}