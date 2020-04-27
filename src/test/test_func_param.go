package testing

import (
	"fmt"
	"strconv"
)

func TestingParam() {
	person := make(map[string]int)
	person["allen"] = 123
	mp := &person

	fmt.Println(person)

	fmt.Printf("map address: %p\n", mp)

	modify(person)

	fmt.Println(person)

}

// can change mmp!
/*
// makemap implements a Go map creation make(map[k]v, hint)
// If the compiler has determined that the map or the first bucket
// can be created on the stack, h and/or bucket may be non-nil.
// If h != nil, the map can be created directly in h.
// If bucket != nil, bucket can be used as the first bucket.
func makemap(t *maptype, hint int64, h *hmap, bucket unsafe.Pointer) *hmap {
//省略无关代码
}
通过查看src/runtime/hashmap.go源代码发现，make函数返回的是一个hmap类型的指针*hmap。
也就是说map==*hmap。
现在看func modify(p map)这样的函数，其实就等于func modify(p *hmap)，
和我们前面第一节什么是值传递里举的func modify(ip *int)的例子一样，可以参考分析。
所以在这里，Go语言通过make函数，字面量的包装，为我们省去了指针的操作，让我们可以更容易的使用map。
这里的map可以理解为引用类型，但是记住引用类型不是传引用。

chan类型本质上和map类型是一样的
func makechan(t *chantype, size int64) *hchan {
//省略无关代码
}
chan也是一个引用类型，和map相差无几，make返回的是一个*hchan。

*/
func modify(mmp map[string]int) {
	fmt.Printf("map1 address: %p\n", &mmp)
	mmp["allen"] = 456
}

//--------------------------------------------------------------------------
type Person struct {
	name string
}

func TestingParam1() {
	P := Person{"hehe"}
	fmt.Println(P)
	fmt.Printf("address: %p\n", &P)
	modifypp(&P)
	fmt.Println(P)
}

// can not change p, pass by value
func modifyp(p Person) {
	fmt.Printf("address1: %p\n", &p)
	p.name = "fuck"
}

// can change p, pass by pointer
func modifypp(p *Person) {
	fmt.Printf("address1: %p\n", p)
	p.name = "fuck"
}

//--------------------------------------------------------------------------
/*
slice和map、chan都不太一样的，一样的是，它也是引用类型，它也可以在函数中修改对应的内容。
运行打印结果，发现的确是被修改了，而且我们这里打印slice的内存地址是可以直接通过%p打印的,
不用使用&取地址符转换。
这就可以证明make的slice也是一个指针了吗？不一定，也可能fmt.Printf把slice特殊处理了。

通过源代码发现，对于chan、map、slice等被当成指针处理，通过value.Pointer()获取对应的值的指针。
// If v's Kind is Slice, the returned pointer is to the first
// element of the slice. If the slice is nil the returned value
// is 0. If the slice is empty but non-nil the return value is non-zero.
func (v Value) Pointer() uintptr {
	// TODO: deprecate
	k := v.kind()
	switch k {
	//省略无关代码
	case Slice:
		return (*SliceHeader)(v.ptr).Data
	}
}
很明显了，当是slice类型的时候，返回是slice这个结构体里，字段Data第一个元素的地址。
type SliceHeader struct {
	Data uintptr
	Len int
	Cap int
}
type slice struct {
	array unsafe.Pointer
	len int
	cap int
}
所以我们通过%p打印的slice变量的地址其实就是内部存储数组元素的地址，slice是一种结构体+元素指针的混合类型，
通过元素array(Data)的指针，可以达到修改slice里存储元素的目的。
所以修改类型的内容的办法有很多种，类型本身作为指针可以，类型里有指针类型的字段也可以。
单纯的从slice这个结构体看，我们可以通过modify修改存储元素的内容，但是永远修改不了len和cap，
因为他们只是一个拷贝，如果要修改，那就要传递*slice作为参数才可以。

*/

func TestingSlice() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("address: %p\n", s)
	modifys(s)
	fmt.Println(s)
}

func modifys(ss []int) {
	fmt.Printf("address1: %p\n", ss)
	ss[1] = 100
}

//--------------------------------------------------------------------------

type PP struct {
	name string
	age  *int
}

func (p PP) String() string {
	return "name:" + p.name + " age:" + strconv.Itoa(*p.age)
}

func modifyy(p PP) {
	p.name = "heihei"
	*p.age = 100
}

func TestingPP() {
	i := 10
	pp := PP{name: "gg", age: &i}
	fmt.Println(pp)
	modifyy(pp)
	fmt.Println(pp)
}

/*
通过这个Person和slice对比，就更好理解了，Person的name字段就类似于slice的len和cap字段，
age字段类似于array字段。在传参为非指针类型的情况下，只能修改age字段，name字段无法修改。
要修改name字段，就要把传参改为指针。
所以slice类型也是引用类型。

Go语言中所有的传参都是值传递（传值），都是一个副本，一个拷贝。
因为拷贝的内容有时候是非引用类型（int、string、struct等这些），这样就在函数中就无法修改原内容数据；
有的是引用类型（指针、map、slice、chan等这些），这样就可以修改原内容数据。
是否可以修改原内容数据，和传值、传引用没有必然的关系。在C++中，传引用肯定是可以修改原内容数据的，
在Go语言里，虽然只有传值，但是我们也可以修改原内容数据，因为参数是引用类型。
这里也要记住，引用类型和传引用是两个概念。
再记住，Go里只有传值（值传递）。
*/
//--------------------------------------------------------------------------

func TestingSlice1() {
	s := make([]rune, 20)
	fmt.Println(s)
	modifyrs(s)
	fmt.Println(s)
}

func modifyrs(ss []rune) {
	i := 0
	for _, r := range "hello world" {
		// ss = append(ss, r)
		ss[i] = r
		i += 1
	}
}
