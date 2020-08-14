package main

import "fmt"

// 依赖倒转设计原则
// 如何正确的解决不同对象之间复杂的以来关系

/*
   模拟组装两台电脑
    ----抽象层----
    有显卡Card  方法display
    有内存 Memory 方法storage
    有处理器CPU    方法calculate
    ----实现层----
    有intel公司  产品有(显卡，cpu，内存)
    有kingston公司  产品有(内存)
    有NVIDIA公司 产品有(显卡)
    -----逻辑层----
    1. 组装一台intel系列电脑
    2. 组装一台intel的cpu，kingston的内存，NVIDIA显卡的电脑
*/

// -------------------抽象层-------------------

// 显卡
type Card interface {
    Display()
}

// 内存
type Memory interface {
    Storage()
}

// 处理器
type CPU interface {
    Calculate()
}

type Computer struct {
    card Card
    mem  Memory
    cpu  CPU
}

func (c *Computer) Run() {
    c.card.Display()
    c.mem.Storage()
    c.cpu.Calculate()
}

func NewComputer(card Card, mem Memory, cpu CPU) *Computer {
    return &Computer{
        card: card,
        mem:  mem,
        cpu:  cpu,
    }
}

// -------------------实现层-----------------
type IntelCard struct {
    Card
}

func (this *IntelCard) Display() {
    fmt.Println("Intel Card 开始显示了...")
}

type IntelMem struct {
    Memory
}

func (this *IntelMem) Storage() {
    fmt.Println("Intel Mem 开始存储了...")
}

type IntelCPU struct {
    CPU
}

func (this *IntelCPU) Calculate() {
    fmt.Println("Intel CPU 开始计算了...")
}

type KingstonMem struct {
    Memory
}

func (this *KingstonMem) Storage() {
    fmt.Println("Kingston mem 开始存储了...")
}

type NVIDIACard struct {
    Card
}

func (this *NVIDIACard) Display() {
    fmt.Println("NVIDIA Card 开始显示了...")
}

// ------------逻辑层--------
func main() {
    com1 := NewComputer(&IntelCard{}, &IntelMem{}, &IntelCPU{})
    com1.Run()

    com2 := NewComputer(&NVIDIACard{}, &KingstonMem{}, &IntelCPU{})
    com2.Run()
}
