## CF Template to CLI

This is a simple lib that converts an existing [Cloudformation](https://aws.amazon.com/cloudformation/) template as a `aws cloudformation deploy ...` command with parameter overrides provided.

### How to Install

Run command below
```
go get -u github.com/harunalfat/cf-template-to-cli/cmd/cfToCli
```

### How to Use
Run command below
```
cfToCli -p path-to-cf-template.yml -s my-stack-name -r eu-west-1
```
