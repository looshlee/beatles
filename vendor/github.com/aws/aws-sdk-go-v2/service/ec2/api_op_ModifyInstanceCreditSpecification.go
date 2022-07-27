// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyInstanceCreditSpecificationInput struct {
	_ struct{} `type:"structure"`

	// A unique, case-sensitive token that you provide to ensure idempotency of
	// your modification request. For more information, see Ensuring Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// Information about the credit option for CPU usage.
	//
	// InstanceCreditSpecifications is a required field
	InstanceCreditSpecifications []InstanceCreditSpecificationRequest `locationName:"InstanceCreditSpecification" locationNameList:"item" type:"list" required:"true"`
}

// String returns the string representation
func (s ModifyInstanceCreditSpecificationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyInstanceCreditSpecificationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyInstanceCreditSpecificationInput"}

	if s.InstanceCreditSpecifications == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceCreditSpecifications"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyInstanceCreditSpecificationOutput struct {
	_ struct{} `type:"structure"`

	// Information about the instances whose credit option for CPU usage was successfully
	// modified.
	SuccessfulInstanceCreditSpecifications []SuccessfulInstanceCreditSpecificationItem `locationName:"successfulInstanceCreditSpecificationSet" locationNameList:"item" type:"list"`

	// Information about the instances whose credit option for CPU usage was not
	// modified.
	UnsuccessfulInstanceCreditSpecifications []UnsuccessfulInstanceCreditSpecificationItem `locationName:"unsuccessfulInstanceCreditSpecificationSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s ModifyInstanceCreditSpecificationOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyInstanceCreditSpecification = "ModifyInstanceCreditSpecification"

// ModifyInstanceCreditSpecificationRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies the credit option for CPU usage on a running or stopped burstable
// performance instance. The credit options are standard and unlimited.
//
// For more information, see Burstable Performance Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using ModifyInstanceCreditSpecificationRequest.
//    req := client.ModifyInstanceCreditSpecificationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyInstanceCreditSpecification
func (c *Client) ModifyInstanceCreditSpecificationRequest(input *ModifyInstanceCreditSpecificationInput) ModifyInstanceCreditSpecificationRequest {
	op := &aws.Operation{
		Name:       opModifyInstanceCreditSpecification,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyInstanceCreditSpecificationInput{}
	}

	req := c.newRequest(op, input, &ModifyInstanceCreditSpecificationOutput{})

	return ModifyInstanceCreditSpecificationRequest{Request: req, Input: input, Copy: c.ModifyInstanceCreditSpecificationRequest}
}

// ModifyInstanceCreditSpecificationRequest is the request type for the
// ModifyInstanceCreditSpecification API operation.
type ModifyInstanceCreditSpecificationRequest struct {
	*aws.Request
	Input *ModifyInstanceCreditSpecificationInput
	Copy  func(*ModifyInstanceCreditSpecificationInput) ModifyInstanceCreditSpecificationRequest
}

// Send marshals and sends the ModifyInstanceCreditSpecification API request.
func (r ModifyInstanceCreditSpecificationRequest) Send(ctx context.Context) (*ModifyInstanceCreditSpecificationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyInstanceCreditSpecificationResponse{
		ModifyInstanceCreditSpecificationOutput: r.Request.Data.(*ModifyInstanceCreditSpecificationOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyInstanceCreditSpecificationResponse is the response type for the
// ModifyInstanceCreditSpecification API operation.
type ModifyInstanceCreditSpecificationResponse struct {
	*ModifyInstanceCreditSpecificationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyInstanceCreditSpecification request.
func (r *ModifyInstanceCreditSpecificationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
