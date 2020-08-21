package stack

//Stack 栈的接口
type Stack interface {
	//Push 压栈
	Push(v interface{})
	//Pop 出栈
	Pop() interface{}
	//IsEmpty 判断是否为空
	IsEmpty() bool
	//Top 取栈顶
	Top() interface{}
	//Flush 清空
	Flush()
	//打印
	Print()
}
