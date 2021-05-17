package main

import "fmt"

type User struct{
	name string
	email string
}

func (u User) Notify(){
	fmt.Printf("sending user email to %s<%s>\n",u.name,u.email)
}

type Admin struct {
	User
	level string
}

func _main()  {
	ad := Admin{
		User:User{
			name: "john smith",
			email: "john smith@email.com",
		},
		level: "super",
	}
	ad.User.Notify()
	ad.Notify()
}


