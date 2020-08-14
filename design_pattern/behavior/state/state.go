package state

import (
	"fmt"
)

/*状态模式
当一个对象的内在状态改变时允许改变其行为，这个对象看起来像是改变了其类

下面的例子是说：一个程序员在一天的不同时间写代码的状态是不一样的
*/

//Work 工作
type Work struct {
	currentState State // 当前的状态
	Finished     bool  // 是否完成
	Hour         int   // 时间，整点
}

//NewWork 新建
func NewWork() *Work {
	return &Work{
		currentState: new(ForenoonState),
	}
}

//SetState 设置状态
func (w *Work) SetState(s State) {
	w.currentState = s
}

//WriteProgram 写代码
func (w *Work) WriteProgram() {
	w.currentState.WriteProgram(w)
}

//State 状态
type State interface {
	WriteProgram(*Work)
}

//ForenoonState 上午
type ForenoonState struct {
}

//WriteProgram 写代码
func (s *ForenoonState) WriteProgram(w *Work) {
	if w.Hour < 12 {
		fmt.Printf("当前时间是%d, 精神百倍...\n", w.Hour)
	} else {
		w.SetState(new(AfternoonState))
	}
}

//AfternoonState 下午午
type AfternoonState struct {
}

//WriteProgram 写代码
func (s *AfternoonState) WriteProgram(w *Work) {
	if w.Hour < 18 {
		fmt.Printf("当前时间是%d, 有点累了...\n", w.Hour)
	} else {
		w.SetState(new(EveningState))
	}
}

//EveningState 下午午
type EveningState struct {
}

//WriteProgram 写代码
func (s *EveningState) WriteProgram(w *Work) {
	if w.Finished {
		w.SetState(new(RestState))
		w.WriteProgram()
	} else {
		if w.Hour < 18 {
			fmt.Printf("当前时间是%d, 疲惫不堪...\n", w.Hour)
		} else {
			w.SetState(new(SleepingState))
			w.WriteProgram()
		}
	}
}

//SleepingState 下午午
type SleepingState struct {
}

//WriteProgram 写代码
func (s *SleepingState) WriteProgram(w *Work) {
	fmt.Printf("当前时间是%d, 睡觉...\n", w.Hour)
}

//RestState 休息时间
type RestState struct {
}

//WriteProgram 写代码
func (s *RestState) WriteProgram(w *Work) {
	fmt.Printf("工作完成了，可以休息了...")
}
