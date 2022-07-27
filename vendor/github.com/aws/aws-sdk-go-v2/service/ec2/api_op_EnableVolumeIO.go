// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

type EnableVolumeIOInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the volume.
	//
	// VolumeId is a required field
	VolumeId *string `locationName:"volumeId" type:"string" required:"true"`
}

// String returns the string representation
func (s EnableVolumeIOInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableVolumeIOInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EnableVolumeIOInput"}

	if s.VolumeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VolumeId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type EnableVolumeIOOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s EnableVolumeIOOutput) String() string {
	return awsutil.Prettify(s)
}

const opEnableVolumeIO = "EnableVolumeIO"

// EnableVolumeIORequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Enables I/O operations for a volume that had I/O operations disabled because
// the data on the volume was potentially inconsistent.
//
//    // Example sending a request using EnableVolumeIORequest.
//    req := client.EnableVolumeIORequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/EnableVolumeIO
func (c *Client) EnableVolumeIORequest(input *EnableVolumeIOInput) EnableVolumeIORequest {
	op := &aws.Operation{
		Name:       opEnableVolumeIO,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableVolumeIOInput{}
	}

	req := c.newRequest(op, input, &EnableVolumeIOOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return EnableVolumeIORequest{Request: req, Input: input, Copy: c.EnableVolumeIORequest}
}

// EnableVolumeIORequest is the request type for the
// EnableVolumeIO API operation.
type EnableVolumeIORequest struct {
	*aws.Request
	Input *EnableVolumeIOInput
	Copy  func(*EnableVolumeIOInput) EnableVolumeIORequest
}

// Send marshals and sends the EnableVolumeIO API request.
func (r EnableVolumeIORequest) Send(ctx context.Context) (*EnableVolumeIOResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableVolumeIOResponse{
		EnableVolumeIOOutput: r.Request.Data.(*EnableVolumeIOOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableVolumeIOResponse is the response type for the
// EnableVolumeIO API operation.
type EnableVolumeIOResponse struct {
	*EnableVolumeIOOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableVolumeIO request.
func (r *EnableVolumeIOResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
