package hello

import (
	"fmt"
)

type user struct {
	name string
	gender string
	age int
}

type animal struct {
	weight float32
	height float32
}

type admin struct {
	person user
	level string
}
var pxq user
var tiger animal


//lisa1 := {"lisa","male",25}



func main()  {
	lisa := user{
		name:"lisa",
		gender:"female",
		age:15,
	}
	fred := admin{
		person: user{
			name:"lisa",
			age:12,
			gender:"female",
		},
		level:"high",
	}
	hilton := admin{
		person: lisa,
		level: "medium",
	}
	fmt.Println(fred.person.age)
	fmt.Println(hilton.person.name,hilton.level)

	arr := [4]int{1,2,3,4}
	arr1 := [...]int{1,2,4,5,7,8}
	//arr2 := [6]int{1,2,3,5,67,8}
	fmt.Print(arr[0])
	fmt.Println(arr1[0])
	var _ [5]string
	ar1 := [5]string{"red","black","white","green","blue"}
	_ = ar1;
	//map1 := make(map[string]int)
	//slice := make([]int,3,5)
	slice1 := []int{1,2,3,4}
	//slice2 := []int{88:10}
	//var slice3 []int        //nil切片
	//slice4 := make([]int,0) //空切片
	//slice5 := []int{}       //空切片
	slice := slice1[1:3]
	slice[1] = 10
	fmt.Println(slice1[2])
	slice = append(slice,30)
	fmt.Println(slice[2])
	fmt.Println(slice1[3])
	fmt.Print("hello")
	slice3 := []int{1,2}
	slice4 := []int{5,6}
	fmt.Println(append(slice3,slice4...))
	for index,value := range slice1{
		fmt.Printf("index:%d value:%d\n",index,value)
		fmt.Printf("value:%d fakeaddr:%X realaddr:%X\n",value,&value,&slice1[index])
	}
	for index := 2;index<len(slice1);index++{
		fmt.Printf("value is %d\n",slice1[index])
	}
	fmt.Println(len(slice1),cap(slice1))

	fmt.Println("==========================")
	slice6 := [][]int{{1},{10,20}}
	for index,_ := range slice6{
		for index1,value := range slice6[index]{
			fmt.Printf("切片%d的第%d个值是%d\n",index, index1,value)
		}
	}
	fmt.Println("========================")
	slice5 := []int{4,6,89,0}
	slice7 := foo(slice5)
	fmt.Println(slice7)
	dicex()

	type Duration int64
	lisa.test()
	fmt.Printf("%s%d岁", lisa.name,lisa.age)

}
func foo(slice []int) []int {
	return append(slice, 11)
}

func dicex()  {
	//dict := make(map[int]string)
	dict1 := map[string]string{"red":"红色","blue":"蓝色","green":"绿色"}
	dict2 := map[string][]int{}
	slice := []int{1,2,3,5}
	dict2["red"] = slice
	//var colors map[string]string

	for key,value := range dict1{
		fmt.Printf("key:%s, value:%s\n", key,value)
	}

	delete(dict1,"red")
	for key,value := range dict1{
		fmt.Printf("key:%s, value:%s\n", key,value)
	}
}
func (u *user) test(){   //如果传u的副本，函数不改变u的值，则使用值传递；如果传u本身，函数改变u的值，则使用指针。

	fmt.Printf("%s%d岁", u.name,u.age)
	u.age = 10
}
