# Simple Go API

Deploy GO function to [Lambda](https://aws.amazon.com/lambda/) using [CDK](https://aws.amazon.com/cdk/).

## Build

```bash
git clone https://github.com/danielcristho/simple-go-api.git
```

```bash
go mod tidy
```

The `cdk.json` file tells the CDK toolkit how to execute your app.

## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `cdk destroy`     destroy the stack
 * `go test`         run unit tests
