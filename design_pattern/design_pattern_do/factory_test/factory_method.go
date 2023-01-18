package factory_test

type IRuleConfigParser interface {
	Parse(data []byte)
}

// IRuleConfigParserFactory 工厂方法接口
type IRuleConfigParserFactory interface {
	CrateParser() IRuleConfigParser
}

type yamlRuleConfigParserFactory struct {
}

func (yaml yamlRuleConfigParserFactory) createParser() IRuleConfigParserFactory {
	return nil
}
