package main

import "fmt"

type Car interface {
    Run()
}

type Driver interface {
    Drive(car Car)
}

type BMW struct {
}

func (b *BMW) Run() {
    fmt.Println("宝马车跑起来...")
}

type BenZ struct {
}

func (b *BenZ) Run() {
    fmt.Println("奔驰车跑起来...")
}

type Zhang3 struct {
}

func (z Zhang3) Drive(car Car) {
    fmt.Println("张三在开车...")
    car.Run()
}

type Li4 struct {
}

func (l *Li4) Drive(car Car) {
    fmt.Println("李四在开车...")
    car.Run()
}

func main() {
    var car Car
    car = &BMW{}

    var driver Driver
    driver = Zhang3{}

    driver.Drive(car)
}
