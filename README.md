# Simple Go REST API

Deploy GO function to [Lambda](https://aws.amazon.com/lambda/) using [CDK](https://aws.amazon.com/cdk/).

## Installation

```bash
git clone https://github.com/danielcristho/cdk-simple-go-restapi.git
```

```bash
make install
```

Or

```bash
make setup
```

## Build

```bash
make build
```

After the deployment is successful, testing the API:

```bash
curl https://<api_gateway_url>/prod/hello-world
```

Response:

```json
{"Greeting":"Hello From API!"}
```

## CDK

The `cdk.json` file tells the CDK toolkit how to execute your app.

## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `cdk destroy`     destroy the stack
 * `go test`         run unit tests
