package main

import (
	"fmt"
)

type notifier1 interface {
	notify()
}

type user struct {
	name string
	email string
}

type admin struct {
	name string
	email string
}
func (u user) notify() {
	fmt.Printf("sending user email to %s<%s>\n",u.name,u.email)
}

func (a admin) notify()  {
	fmt.Printf("Sending admin email to %s<%s>",a.name,a.email)
}

func _() {
	bill := user{"Bill","bill@email.com"}
	sendNotification(bill)
	lisa := admin{"lisa", "lisa@email.com"}
	sendNotification(lisa)
}
func sendNotification(n notifier)  {
	n.notify()
}


