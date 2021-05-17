package main

import (
	"Go5-6/entities"
	"fmt"
)

func main2(){
	u := entities.User{
		Name:"Bill",
		//email:"bill@email.com",
		Email: "Bill@email.com",
	}
	fmt.Printf("User:%v\n",u)
}
