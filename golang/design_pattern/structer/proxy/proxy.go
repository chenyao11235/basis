package proxy

import (
	"fmt"
	"time"
)

/*  需求：
    需要统计登陆操作要耗费多长时间
*/

//IUser 统一接口
type IUser interface {
	login()
}

//UserController 被代理类
type UserController struct {
}

func (u *UserController) login() {
}

//UserControllerProxy 代理类  通过组合的方式实现代理模式
type UserControllerProxy struct {
	userController UserController
}

func (up *UserControllerProxy) login() {
	startTime := time.Now()

	up.userController.login()

	costTime := time.Since(startTime)
	fmt.Println(costTime)
}
