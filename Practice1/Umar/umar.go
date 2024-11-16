package umar

import "fmt"

type Umar struct {
	name  string
	lname string
}

func (u *Umar) PrintName() {
	u.name = "umar"
	u.lname = "kasan"
	fmt.Println("My name is", u.name, u.lname)
}
