package main

import "fmt"

type notifier interface {
	notify()
}
type user struct {
	name string
	email string
}

func (u user) notify()  {
	fmt.Printf("sending user email to %s<%s>\n",u.name,u.email)
}

type admin struct {
	user
	level string
}

func (a admin) notify()  {
	fmt.Printf("sending admin email to %s<%s>\n",a.name,a.email)
}

func main()  {
	ad := admin{
		user{
			"bill",
			"bill@email.com",
		},
			"super",
	}
	sendNotification(ad) //外部类型和内部类型都实现了某个方法，那么就不会调用内部类型的方法。
	ad.user.notify()
	ad.notify() //外部类型和内部类型都实现了某个方法，那么就不会调用内部类型的方法。

}

func sendNotification(n notifier)  {
	n.notify()
}

