package factory_method

import (
	"reflect"
	"testing"
)

func TestJsonRuleConfigParser_Parse(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := JsonRuleConfigParser{}
			j.Parse(tt.args.data)
		})
	}
}

func TestNewIRuleConfigParserFactory(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want IRuleConfigParserFactory
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIRuleConfigParserFactory(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParserFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYamlRuleConfigParser_Parse(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := YamlRuleConfigParser{}
			y.Parse(tt.args.data)
		})
	}
}

func Test_jsonRuleConfigParserFactory_CrateParser(t *testing.T) {
	tests := []struct {
		name string
		want IRuleConfigParser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := jsonRuleConfigParserFactory{}
			if got := j.CrateParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CrateParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yamlRuleConfigParserFactory_CrateParser(t *testing.T) {
	tests := []struct {
		name string
		want IRuleConfigParser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml := yamlRuleConfigParserFactory{}
			if got := yaml.CrateParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CrateParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonRuleConfigParser_Parse1(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := JsonRuleConfigParser{}
			j.Parse(tt.args.data)
		})
	}
}

func TestNewIRuleConfigParserFactory1(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want IRuleConfigParserFactory
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIRuleConfigParserFactory(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParserFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYamlRuleConfigParser_Parse1(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := YamlRuleConfigParser{}
			y.Parse(tt.args.data)
		})
	}
}

func Test_jsonRuleConfigParserFactory_CrateParser1(t *testing.T) {
	tests := []struct {
		name string
		want IRuleConfigParser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := jsonRuleConfigParserFactory{}
			if got := j.CrateParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CrateParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yamlRuleConfigParserFactory_CrateParser1(t *testing.T) {
	tests := []struct {
		name string
		want IRuleConfigParser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml := yamlRuleConfigParserFactory{}
			if got := yaml.CrateParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CrateParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
