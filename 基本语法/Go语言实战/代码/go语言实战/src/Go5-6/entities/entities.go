package entities

type  User struct {
	Name string
	email string  //email字段不能被包外访问,但是可以申明一个只含有Name字段的User对象
	Email string

}

type user struct{
	Name string
	Email string
}
type Admin struct {
	user
	Rights int
}
