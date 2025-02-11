package lambda

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	aws_config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/common-fate/granted-approvals/accesshandler/pkg/config"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/common-fate/apikit/logger"
	"github.com/common-fate/granted-approvals/accesshandler/pkg/providers"
	"github.com/common-fate/granted-approvals/accesshandler/pkg/types"
)

// calls out to the provider to revoke access to the grant and disables execution to the granter state function
func (r *Runtime) RevokeGrant(ctx context.Context, grantID string, revoker string) (*types.Grant, error) {

	logger.Get(ctx).Infow("revoking grant", "grant", grantID)

	//using the grantID we need to work out all of the grant data info from the previous invocation when the grant was created

	//we can grab all this from the execution input for the step function we will use this as the source of truth
	c, err := aws_config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}
	sfnClient := sfn.NewFromConfig(c)

	//build the execution ARN
	exeARN := BuildExecutionARN(r.GranterStateMachineARN, grantID)

	out, err := sfnClient.DescribeExecution(ctx, &sfn.DescribeExecutionInput{ExecutionArn: aws.String(exeARN)})
	if err != nil {
		return nil, err
	}

	//build the previous grant from the execution input
	var grantInput WorkflowInput

	err = json.Unmarshal([]byte(*out.Input), &grantInput)
	if err != nil {
		return nil, err
	}
	grant := grantInput.Grant

	prov, ok := config.Providers[grant.Provider]
	if !ok {
		return nil, &providers.ProviderNotFoundError{Provider: grant.Provider}
	}

	args, err := json.Marshal(grant.With)
	if err != nil {
		return nil, err
	}

	//if the state function is in the active state then we will stop the execution
	statefn, err := sfnClient.GetExecutionHistory(ctx, &sfn.GetExecutionHistoryInput{ExecutionArn: &exeARN})
	if err != nil {
		return nil, err
	}
	lastState := statefn.Events[len(statefn.Events)-1]
	//if the state of the grant is in the active state
	if lastState.Type == "WaitStateEntered" && *lastState.StateEnteredEventDetails.Name == "Wait for Window End" {
		err = prov.Provider.Revoke(ctx, string(grant.Subject), args, grant.ID)
		if err != nil {
			return nil, err
		}
	}

	_, err = sfnClient.StopExecution(ctx, &sfn.StopExecutionInput{ExecutionArn: &exeARN})
	//if stopping the execution failed we want return with an error and not continue with the flow
	if err != nil {
		return nil, err
	}

	//update the grant status
	grant.Status = types.GrantStatusREVOKED

	return &grant, nil
}

//takes a grant id and state machine arn and returns an execution arn for a given state machine
//eg state machine arn: "arn:aws:states:us-east-1:123456789012:stateMachine:StatemachineNameASfgeg-dfvrfgafg"
//eg execution arn: 	"arn:aws:states:us-east-1:123456789012:execution:StatemachineNameASfgeg-dfvrfgafg:{grantID}"

func BuildExecutionARN(stateMachineARN string, grantID string) string {

	splitARN := strings.Split(stateMachineARN, ":")

	//position 5 is the location of the arn type
	splitARN[5] = "execution"
	splitARN = append(splitARN, grantID)

	return strings.Join(splitARN, ":")

}
