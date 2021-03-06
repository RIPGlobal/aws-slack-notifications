module github.com/RIPGlobal/aws-slack-notifications/services/deployment-notifications

go 1.13

require (
	github.com/RIPGlobal/aws-slack-notifications v0.0.0-20191115011656-27fe591efa78
	github.com/aws/aws-lambda-go v1.13.3
	github.com/aws/aws-sdk-go v1.25.35
	github.com/nlopes/slack v0.6.0
	github.com/stretchr/testify v1.4.0
)

// Due to issues with `aws-lambda-go` use my own fork with the fixes
//  * CodePipelineEvent hardcoded to CodePipelineJob.
//    - https://github.com/aws/aws-lambda-go/issues/244
//  * CodePipelineEvent missing (Incorrectly referencing only jobs)
//    - https://github.com/aws/aws-lambda-go/issues/246
// TODO: Move back to official once PRs merged:
//  * Clean up CodePipeline Job Implementation
//    - https://github.com/aws/aws-lambda-go/pull/245
//  * Implement CodePipelineEvent
//    - https://github.com/aws/aws-lambda-go/pull/247
//
replace github.com/aws/aws-lambda-go => github.com/whithajess/aws-lambda-go v1.13.2-0.20191030023142-ba8d4131ff69

// replace also can be used to inform the go tooling of the relative or absolute on-disk location of modules in a multi-module project, such as:
// replace example.com/project/foo => ../foo
// we use this to include our internal packages outside the root of this service.
// see: https://docs.google.com/document/d/1e8kOo3r51b2BWtTs_1uADIA5djfXhPT36s6eHVRIvaU/edit
//
replace github.com/RIPGlobal/aws-slack-notifications => ../../
