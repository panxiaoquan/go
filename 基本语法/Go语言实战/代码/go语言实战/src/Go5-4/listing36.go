package main

import (
	"fmt"
)

type notifier interface {
	notify()
	//aaa()
}

type user1 struct {
	name string
	email string
}

/*如果使用指
针接收者来实现一个接口，那么只有指向那个类型的指针才能够实现对应的接口。如果使用值
接收者来实现一个接口，那么那个类型的值和指针都能够实现对应的接口。*/
func (u user1) notify()  { //用值接收者user实现接口的方法即为实现了接口notifier
	fmt.Printf("sending user email to %s<%s>",u.name,u.email)
}

func (u user) aaa()  {
	
}

func _()  {
	u := user{name: "Bill",email: "bill@email.com"}
	sendNotification1(u)

}

func sendNotification1(n notifier)  {
	n.notify()
}


