package main

import "fmt"

// 通知类行为的接口
type notifier interface {
	notify()
	notify2()
}

type user struct {
	name string
	email string
}

func (u *user) notify(){
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func (u user) notify2(){
	fmt.Printf("Sending user email to %s<%s> in 2nd way\n", u.name, u.email)
}

type admin struct {
	user // 嵌入类型
	level string
}

func (a *admin) notify2() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

func main(){
	u := user{"Bill", "bill@ex.com"}
	// sendNotification(u) 注意notify接收者是指针类型，如果参数为类型值类型的话，是无法找到interface对应方法实现的
	sendNotification(&u)
	//sendNotification(u)
	// 在直接使用类型实例调用类型的方法时，无论值类型变量还是指针类型变量都可以调用所有方法，原因是编译器辅助完成了自动转换；而在接口值中，无法自动转换，原因如下
	// 接口是在动态运行时解析的，产生指针 就会产生引用
	// 在接口中，因为不是总能获取一个值的地址，因此接口值的方法集只包括使用值接收者实现的方法！
	//var n notifier = u // cannot use u (type user) as type notifier in assignment:user does not implement notifier (notify method has pointer receiver)
	//n.notify2()
	u.notify()
	u.notify2()

	(&u).notify2()
	(&u).notify()
	//sendNotification(u)

	su := admin{user: user{name: "john snow", email:"snow@xx.com"}, level:"super"}
	su.user.notify() // 可以直接访问内部类型的方法
	su.notify()      // 内部类型的方法被提升到外部类型
	sendNotification(&su)
}

func sendNotification(n notifier){
	n.notify()
	n.notify2()
}
