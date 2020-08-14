package iod

/*依赖反转原则之控制反转
 */

//UserServiceTest 用户服务测试
type UserServiceTest struct {
}

func (u *UserServiceTest) doTest() bool {
	return true
}

// 正常情况下，我们直接在main中new一个UserServiceTest再调用doTest就可以进行测试了

// func main() {
// ut := &UserServiceTest{}
// ut.doTest()
// }

// 下面实现一个测试框架, 通过测试框架来运行测试用例

//TestCase 测试接口
type TestCase interface {
	doTest() bool
}

//TestApplication 测试框架
type TestApplication struct {
	testCaseList []TestCase
}

// 注册一个测试用例到框架中
func (ta *TestApplication) register(t TestCase) {
	ta.testCaseList = append(ta.testCaseList, t)
}

// 运行所有的测试用例
func (ta *TestApplication) run() {
	for _, t := range ta.testCaseList {
		t.doTest()
	}
}

/*
这里的“控制”指的是对程序执行流程的控制，
而“反转”指的是在没有使用框架之前，程序员自己控制整个程序的执行。
在使用框架之后，整个程序的执行流程可以通过框架来控制。
流程的控制权从程序员“反转”到了框架。
*/

func main() {
	TA := &TestApplication{
		testCaseList: make([]TestCase, 0),
	}
	ut := &UserServiceTest{}

	TA.register(ut)
	TA.run()
}
