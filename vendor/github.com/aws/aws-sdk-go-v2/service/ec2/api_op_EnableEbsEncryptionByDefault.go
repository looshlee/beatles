// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type EnableEbsEncryptionByDefaultInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`
}

// String returns the string representation
func (s EnableEbsEncryptionByDefaultInput) String() string {
	return awsutil.Prettify(s)
}

type EnableEbsEncryptionByDefaultOutput struct {
	_ struct{} `type:"structure"`

	// The updated status of encryption by default.
	EbsEncryptionByDefault *bool `locationName:"ebsEncryptionByDefault" type:"boolean"`
}

// String returns the string representation
func (s EnableEbsEncryptionByDefaultOutput) String() string {
	return awsutil.Prettify(s)
}

const opEnableEbsEncryptionByDefault = "EnableEbsEncryptionByDefault"

// EnableEbsEncryptionByDefaultRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Enables EBS encryption by default for your account in the current Region.
//
// After you enable encryption by default, the EBS volumes that you create are
// are always encrypted, either using the default CMK or the CMK that you specified
// when you created each volume. For more information, see Amazon EBS Encryption
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
// You can specify the default CMK for encryption by default using ModifyEbsDefaultKmsKeyId
// or ResetEbsDefaultKmsKeyId.
//
// Enabling encryption by default has no effect on the encryption status of
// your existing volumes.
//
// After you enable encryption by default, you can no longer launch instances
// using instance types that do not support encryption. For more information,
// see Supported Instance Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#EBSEncryption_supported_instances).
//
//    // Example sending a request using EnableEbsEncryptionByDefaultRequest.
//    req := client.EnableEbsEncryptionByDefaultRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/EnableEbsEncryptionByDefault
func (c *Client) EnableEbsEncryptionByDefaultRequest(input *EnableEbsEncryptionByDefaultInput) EnableEbsEncryptionByDefaultRequest {
	op := &aws.Operation{
		Name:       opEnableEbsEncryptionByDefault,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableEbsEncryptionByDefaultInput{}
	}

	req := c.newRequest(op, input, &EnableEbsEncryptionByDefaultOutput{})

	return EnableEbsEncryptionByDefaultRequest{Request: req, Input: input, Copy: c.EnableEbsEncryptionByDefaultRequest}
}

// EnableEbsEncryptionByDefaultRequest is the request type for the
// EnableEbsEncryptionByDefault API operation.
type EnableEbsEncryptionByDefaultRequest struct {
	*aws.Request
	Input *EnableEbsEncryptionByDefaultInput
	Copy  func(*EnableEbsEncryptionByDefaultInput) EnableEbsEncryptionByDefaultRequest
}

// Send marshals and sends the EnableEbsEncryptionByDefault API request.
func (r EnableEbsEncryptionByDefaultRequest) Send(ctx context.Context) (*EnableEbsEncryptionByDefaultResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableEbsEncryptionByDefaultResponse{
		EnableEbsEncryptionByDefaultOutput: r.Request.Data.(*EnableEbsEncryptionByDefaultOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableEbsEncryptionByDefaultResponse is the response type for the
// EnableEbsEncryptionByDefault API operation.
type EnableEbsEncryptionByDefaultResponse struct {
	*EnableEbsEncryptionByDefaultOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableEbsEncryptionByDefault request.
func (r *EnableEbsEncryptionByDefaultResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
