package main
//
//import "fmt"
//
//type notifier interface {
//	notify()
//}
//
//func (u User) notify()  {
//	fmt.Printf("sending user email to %s<%s>",u.name,u.email)
//}
//
//func _()  {
//	ad := Admin{
//		User{
//			"john smith",
//			"john smith@email.com",
//		},
//		"super",
//	}
//	sendNotification(ad)   //内部类型出现的地方都可以用外部类型替代，go可以从外部类型中自行挑选出内部类型部分。
//}
//
//func sendNotification(n notifier)  {
//	n.notify()
//}
