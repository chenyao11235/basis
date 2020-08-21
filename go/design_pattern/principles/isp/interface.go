package isp

/*接口隔离原则
接口的设计要尽可能的单一，不要啥乱七八糟的东西都赛道一个接口中
*/

//UserInfo 用户信息
type UserInfo struct {
}

//UserService 用户接口
type UserService interface {
	register() bool
	login() bool
	getUserInfoByID(ID int) UserInfo
	getUserInfoByPhone(phone string) UserInfo
}

// 此时需要实现一个删除用户的功能，但是这个删除功能不应该被普通用户调用，而只能被后台调用
// 这时就应该把这个接口隔离开来，放到另外一个接口中

//RestrictedUserService 受限用户接口
type RestrictedUserService interface {
	deleteByID(ID int) bool
	deleteByPhone(phone string) bool
}
