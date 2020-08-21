package main

import "fmt"

// 开闭原则
//  1. 对修改关闭： 保证最基础的接口功能不被影响
//  2. 对扩展开放： 扩展新的功能，无论新的功能修改成什么样，都不能影响原有的功能

type AbstractBanker interface {
    DoBusiness() // 核心的业务接口
}

type SaveBanker struct {
    AbstractBanker
}

func (sb *SaveBanker) DoBusiness() {
    fmt.Println("进行了存款...")
}

type TransBanker struct {
    AbstractBanker
}

func (tb *TransBanker) DoBusiness() {
    fmt.Println("进行了转账...")
}

func BankerBusiness(b AbstractBanker) {
    fmt.Println("进行了业务...")
    b.DoBusiness()
}

func main() {
    BankerBusiness(&SaveBanker{})
    BankerBusiness(&TransBanker{})
}
