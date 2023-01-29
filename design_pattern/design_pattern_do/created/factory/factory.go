package factory

type IRuleConfigParser interface {
	Parse(data []byte)
}

type JsonRuleConfigParser struct{}

func (jsonParser JsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

type YamlRuleConfigParser struct {
}

func (yamlParser YamlRuleConfigParser) Parse(data []byte) {
	panic("implement me2")
}

//
func NewIRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "json":
		return JsonRuleConfigParser{}
	case "yaml":
		return YamlRuleConfigParser{}
	}
	return nil
}
