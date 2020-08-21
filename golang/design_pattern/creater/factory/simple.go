package factory

/* 简单工厂
golang中没有构造函数一说，一般通过NewXxxx来初始化相关的类，如果Newxxx返回的是接口类型那么就是简单工厂模式
*/

//RuleConfigParserSimpleFactory 工厂
type RuleConfigParserSimpleFactory struct {
}

func (f *RuleConfigParserSimpleFactory) createParser(configFormat string) (parser Parser) {
	if configFormat == "json" {
		parser = new(JSONRuleConfigParser)
	} else if configFormat == "xml" {
		parser = new(XMLRuleConfigParser)
	} else if configFormat == "yaml" {
		parser = new(YamlRuleConfigParser)
	} else {
		parser = nil
	}
	return
}

//RuleConfigSource 工厂方法的调用方法
type RuleConfigSource struct {
}

func (r *RuleConfigSource) load() *ruleConfig {
	var parser Parser
	var factory = &RuleConfigParserSimpleFactory{}
	configFormat := r.getFileExtion("./config.json")
	parser = factory.createParser(configFormat)
	ruleConfig := parser.parse()
	return ruleConfig
}

func (r *RuleConfigSource) getFileExtion(filePath string) string {
	return "json"
}
