package factory

/*工厂方法  就是工厂的工厂
在简单工厂中，使用if判断来返回不同的对象，严格来说是违背开闭原则的
为了让代码更加符合开闭原则，可以利用多态来实现
golang中没有继承，使用组合

工厂方法适用于 不同的struct的构造过程比较复杂，这些不同的struct不适合放在同一个工厂函数中，而需要各自有自己的工厂函数
*/

//IRuleConfigParserMethodFactory 工厂接口
type IRuleConfigParserMethodFactory interface {
	createParser() Parser
}

//JSONRuleConfigParserFactory yaml的工厂函数
type JSONRuleConfigParserFactory struct {
}

func (p *JSONRuleConfigParserFactory) createParser() Parser {
	return &JSONRuleConfigParser{}
}

//YamlRuleConfigParserFactory yaml的工厂函数
type YamlRuleConfigParserFactory struct {
}

func (p *YamlRuleConfigParserFactory) createParser() Parser {
	return &YamlRuleConfigParser{}
}

//XMLRuleConfigParserFactory yaml的工厂函数
type XMLRuleConfigParserFactory struct {
}

func (p *XMLRuleConfigParserFactory) createParser() Parser {
	return &XMLRuleConfigParser{}
}

//GetRuleConfigParserFactory 工厂的工厂 预先缓存
func GetRuleConfigParserFactory(configFormat string) IRuleConfigParserMethodFactory {
	cachedFactories := map[string]IRuleConfigParserMethodFactory{
		"json": new(JSONRuleConfigParserFactory),
		"xml":  new(XMLRuleConfigParserFactory),
		"yaml": new(YamlRuleConfigParserFactory),
	}
	return cachedFactories[configFormat]
}

//RuleConfigSource1 工厂方法的调用方法
type RuleConfigSource1 struct {
}

func (r *RuleConfigSource1) load() *ruleConfig {
	var parser Parser
	configFormat := r.getFileExtion("./config.json")
	factory := GetRuleConfigParserFactory(configFormat)
	parser = factory.createParser()
	ruleConfig := parser.parse()
	return ruleConfig
}

func (r *RuleConfigSource1) getFileExtion(filePath string) string {
	return "json"
}
