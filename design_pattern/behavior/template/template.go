package template

import "fmt"

/*  模板模式一般是需要通过继承的的方式实现，在父类中定义抽象方法，可复用方法，来搭建某个功能的骨架
然后不同的子类去实现父类中抽象方法，或者复用父类中的通用方法

golang因为没有继承这一特性，所以通过组合来实现，不同的组合对象之间互相调用，来实现模板模式
这种不同对象通过组合互相调用也叫做回调
*/

/*  以下代码的说明：
对外的接口是Handler(处理)， handler需要调用接口Operator
Template是模版类，这个类需要实现Operator接口和Handler接口，但是它只实现了Operator的save方法作为默认方法
并没有实现download方法，这个方法留给它的“子类”去实现

HTTPhandler， FtpHandler是子类，它们通过组合Template模板类来实现Handler接口
FtpHandler并没有实现save方法，所以它会调用Template的save方法（默认方法）
*/

//Handler 接口
type Handler interface {
	Handle(url string)
}

//Operater 接口
type Operater interface {
	download(url string)
	save()
}

//TemplateHandler 模板实现类，功能骨架
type TemplateHandler struct {
	Operater
	URL string
}

//Handle 处理方法
func (t *TemplateHandler) Handle(url string) {
	t.URL = url
	t.Operater.download(url)
	t.Operater.save()
}

//Save 保存
func (t *TemplateHandler) save() {
	fmt.Println("default save...")
}

//NewTemplate 新建
func NewTemplate(operator Operater) *TemplateHandler {
	return &TemplateHandler{
		Operater: operator,
	}
}

//HTTPHandler 针对http的操作实现类
type HTTPHandler struct {
	*TemplateHandler
}

//NewHTTPHandler 新建
func NewHTTPHandler() Handler {
	handler := &HTTPHandler{}
	template := NewTemplate(handler)
	handler.TemplateHandler = template
	return handler
}

//Handle 处理
func (d *HTTPHandler) download(url string) {
	fmt.Println("download " + url)
}

//Save 保存
func (d *HTTPHandler) save() {
	fmt.Println("http save...")
}

//Ftphandler 针对ftp的操作实现类
type Ftphandler struct {
	*TemplateHandler
}

//NewFtpHandler 新建
func NewFtpHandler() Handler {
	handler := &Ftphandler{}
	template := NewTemplate(handler)
	handler.TemplateHandler = template
	return handler
}

//Download 下载
func (o *Ftphandler) download(url string) {
	fmt.Println("download " + url)
}
