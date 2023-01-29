package factory_method

type JsonRuleConfigParser struct{}

func (j JsonRuleConfigParser) Parse(data []byte) {

}

type YamlRuleConfigParser struct{}

func (y YamlRuleConfigParser) Parse(data []byte) {

}

type IRuleConfigParser interface {
	Parse(data []byte)
}

// IRuleConfigParserFactory 工厂方法接口
type IRuleConfigParserFactory interface {
	CrateParser() IRuleConfigParser
}

type yamlRuleConfigParserFactory struct {
}

func (yaml yamlRuleConfigParserFactory) CrateParser() IRuleConfigParser {
	return YamlRuleConfigParser{}
}

// jsonRuleConfigParserFactory jsonRuleConfigParser 的工厂类
type jsonRuleConfigParserFactory struct {
}

// CreateParser CreateParser
func (j jsonRuleConfigParserFactory) CrateParser() IRuleConfigParser {
	return JsonRuleConfigParser{}
}

// NewIRuleConfigParserFactory 用一个简单工厂封装工厂方法
func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		return jsonRuleConfigParserFactory{}
	case "yaml":
		return yamlRuleConfigParserFactory{}
	}
	return nil
}
