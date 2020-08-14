package factory

type ruleConfig struct {
	configFormat string
}

//Parser 用于解析配置文件
type Parser interface {
	parse() *ruleConfig
	getConfigFormat() string
}

//JSONRuleConfigParser 解析json文件
type JSONRuleConfigParser struct {
}

func (p *JSONRuleConfigParser) parse() *ruleConfig {
	return &ruleConfig{
		configFormat: p.getConfigFormat(),
	}
}

func (p *JSONRuleConfigParser) getConfigFormat() string {
	return "json"
}

//XMLRuleConfigParser 解析xml文件
type XMLRuleConfigParser struct {
}

func (p *XMLRuleConfigParser) parse() *ruleConfig {
	return &ruleConfig{
		configFormat: p.getConfigFormat(),
	}
}

func (p *XMLRuleConfigParser) getConfigFormat() string {
	return "xml"
}

//YamlRuleConfigParser 解析yaml文件
type YamlRuleConfigParser struct {
}

func (p *YamlRuleConfigParser) parse() *ruleConfig {
	return &ruleConfig{
		configFormat: p.getConfigFormat(),
	}
}

func (p *YamlRuleConfigParser) getConfigFormat() string {
	return "yaml"
}
