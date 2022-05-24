package main

import (
	"flag"
	"fmt"

	"github.com/harunalfat/cf-template-to-cli/lib"
)

var templateFilePath string
var stackName string
var region string

func main() {
	flag.StringVar(&templateFilePath, "p", "cf.yml", "Path of AWS Cloudformation template file")
	flag.StringVar(&stackName, "s", "default-stack-name", "Name of the stack")
	flag.StringVar(&region, "r", "eu-west-1", "Name of the region")
	flag.Parse()

	template := lib.ParseYaml(templateFilePath)
	text := lib.ToCliCommand(template, stackName, region, templateFilePath)
	fmt.Println(text)
}
