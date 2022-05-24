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

### Example

Existing Cloudformation *my-template.yml* with template as:
```
AWSTemplateFormatVersion: "2010-09-09"
Description:
  Stack for my service
Parameters:
  RedisEndpoint:
    Type: String
    Description: Endpoint for Redis
    Default: my-redis.something.clustercfg.euw1.cache.amazonaws.com:6379
  LogFile:
    Type: String
    Description: Path for log file
    Default: /var/log/app.log
```

Run command as
```
cfToCli -p my-template.yml -s my-service -r eu-west-1
```

Command Result:
```
aws cloudformation deploy \
	--region eu-west-1 \
	--template-file my-template.yml \
	--stack-name my-service \
	--capabilities CAPABILITY_NAMED_IAM \
	--parameter-overrides \
	  RedisEndpoint=my-redis.something.clustercfg.euw1.cache.amazonaws.com:6379 \
    LogFile=/var/log/app.log
```
