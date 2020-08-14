package main

import "fmt"

type Phone interface {
    Call()
}

type NokiaPhone struct {
}

func (n *NokiaPhone) Call() {
    fmt.Println("我在给你打电话...from nokia")
}

type ApplePhone struct {
}

func (a *ApplePhone) Call() {
    fmt.Println("我在给你打电话...from apple")
}

func PhoneCall(phone Phone) {
    phone.Call()
}

func main() {
    PhoneCall(&NokiaPhone{})
    PhoneCall(&ApplePhone{})
}
