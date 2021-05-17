package counters

type  alertCounter int //当一个标识符的名字以小写字母开头时，这个标识符就是未公开的，即只能在包里使用，不能被包外代码所见；

type AlterCounter int//当一个标识符的名字以大写字母开头时，这个标识符就是公开的，即可为包外代码所见

func New(value int) alertCounter {
	return alertCounter(value)
}

