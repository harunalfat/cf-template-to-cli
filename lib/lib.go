package lib

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type CFParameter struct {
	Type        string `yaml:"Type"`
	Description string `yaml:"Description"`
	Default     string `yaml:"Default"`
}

type CFTemplate struct {
	Parameters map[string]CFParameter `yaml:"Parameters"`
}

func ParseYaml(cfTemplatePath string) CFTemplate {
	data, err := ioutil.ReadFile(cfTemplatePath)
	if err != nil {
		panic(err)
	}

	var template CFTemplate
	err = yaml.Unmarshal(data, &template)
	if err != nil {
		panic(err)
	}

	return template
}

func toParameterOverrides(parameters map[string]CFParameter) string {
	var text string
	idx := 0
	for name, val := range parameters {
		idx++
		text += fmt.Sprintf("%s=%s", name, val.Default)
		if idx != len(parameters) {
			text += " \\\n\t\t"
		}
	}

	return text
}

func ToCliCommand(template CFTemplate, stackName, region, pathToFile string) string {
	parameterOverrides := toParameterOverrides(template.Parameters)
	text := fmt.Sprintf(`
aws cloudformation deploy \
	--region %s \
	--template-file %s \
	--stack-name %s \
	--capabilities CAPABILITY_NAMED_IAM \
	--parameter-overrides \
		%s
	`, region, pathToFile, stackName, parameterOverrides)

	return text
}
