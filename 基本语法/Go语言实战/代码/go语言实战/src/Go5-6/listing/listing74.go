package main

import (
	"Go5-6/entities"
	"fmt"
)

func main(){
	a := entities.Admin{
		Rights: 10,
	}                 //user未公开，Name，Email公开，Name,Email依然可以访问
	a.Name = "Bill" //user未公开，user内的变量公开，可以用这种方法赋值
	a.Email = "Bill@email.com"
	fmt.Printf("Admin:%v\n",a)
}
