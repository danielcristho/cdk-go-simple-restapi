package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkStackProps struct {
	awscdk.StackProps
}

type LambdaData struct {
	name        string
	httpMethods string
}

func NewCdkStack(scope constructs.Construct, id string, props *CdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	api := awsapigateway.NewRestApi(stack, jsii.String("fpl-intel-cdk"), &awsapigateway.RestApiProps{
		DeployOptions: &awsapigateway.StageOptions{
			StageName: jsii.String("dev"),
		},
	})

	helloWorldLambdaData := LambdaData{
		name:        "hello-world",
		httpMethods: "GET",
	}

	helloWorldLambda := CreateNewGoLambdaFunction(helloWorldLambdaData, stack)
	AddLambdaToApiGateway(helloWorldLambda, helloWorldLambdaData, api)

	return stack
}

func CreateNewGoLambdaFunction(lambdaConfig LambdaData, stack awscdk.Stack) awslambda.Function {
	lambdaPath := "./api"
	lambdaCode := awslambda.Code_FromAsset(&lambdaPath, nil)

	return awslambda.NewFunction(stack, jsii.String(lambdaConfig.name), &awslambda.FunctionProps{
		Runtime:    awslambda.Runtime_PROVIDED_AL2(),
		Code:       lambdaCode,
		Handler:    jsii.String("main"),
		MemorySize: jsii.Number(128),
		Timeout:    awscdk.Duration_Seconds(jsii.Number(28)),
	})
}

func AddLambdaToApiGateway(lambdaFunction awslambda.Function, lambdaConfig LambdaData, api awsapigateway.RestApi) {
	resource := api.Root().AddResource(jsii.String(lambdaConfig.name), nil)
	apiIntegration := awsapigateway.NewLambdaIntegration(lambdaFunction, &awsapigateway.LambdaIntegrationOptions{
		Proxy: jsii.Bool(true),
	})

	corsOptions := &awsapigateway.CorsOptions{
		AllowOrigins: awsapigateway.Cors_ALL_ORIGINS(),
		AllowMethods: awsapigateway.Cors_ALL_METHODS(),
	}
	resource.AddCorsPreflight(corsOptions)
	resource.AddMethod(jsii.String("GET"), apiIntegration, nil)
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewCdkStack(app, "CdkStack", &CdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	//return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
