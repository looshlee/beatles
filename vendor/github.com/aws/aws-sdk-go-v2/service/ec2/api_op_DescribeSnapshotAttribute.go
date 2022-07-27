// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeSnapshotAttributeInput struct {
	_ struct{} `type:"structure"`

	// The snapshot attribute you would like to view.
	//
	// Attribute is a required field
	Attribute SnapshotAttributeName `type:"string" required:"true" enum:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the EBS snapshot.
	//
	// SnapshotId is a required field
	SnapshotId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeSnapshotAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSnapshotAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSnapshotAttributeInput"}
	if len(s.Attribute) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Attribute"))
	}

	if s.SnapshotId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnapshotId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeSnapshotAttributeOutput struct {
	_ struct{} `type:"structure"`

	// The users and groups that have the permissions for creating volumes from
	// the snapshot.
	CreateVolumePermissions []CreateVolumePermission `locationName:"createVolumePermission" locationNameList:"item" type:"list"`

	// The product codes.
	ProductCodes []ProductCode `locationName:"productCodes" locationNameList:"item" type:"list"`

	// The ID of the EBS snapshot.
	SnapshotId *string `locationName:"snapshotId" type:"string"`
}

// String returns the string representation
func (s DescribeSnapshotAttributeOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeSnapshotAttribute = "DescribeSnapshotAttribute"

// DescribeSnapshotAttributeRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the specified attribute of the specified snapshot. You can specify
// only one attribute at a time.
//
// For more information about EBS snapshots, see Amazon EBS Snapshots (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSSnapshots.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DescribeSnapshotAttributeRequest.
//    req := client.DescribeSnapshotAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeSnapshotAttribute
func (c *Client) DescribeSnapshotAttributeRequest(input *DescribeSnapshotAttributeInput) DescribeSnapshotAttributeRequest {
	op := &aws.Operation{
		Name:       opDescribeSnapshotAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSnapshotAttributeInput{}
	}

	req := c.newRequest(op, input, &DescribeSnapshotAttributeOutput{})
	return DescribeSnapshotAttributeRequest{Request: req, Input: input, Copy: c.DescribeSnapshotAttributeRequest}
}

// DescribeSnapshotAttributeRequest is the request type for the
// DescribeSnapshotAttribute API operation.
type DescribeSnapshotAttributeRequest struct {
	*aws.Request
	Input *DescribeSnapshotAttributeInput
	Copy  func(*DescribeSnapshotAttributeInput) DescribeSnapshotAttributeRequest
}

// Send marshals and sends the DescribeSnapshotAttribute API request.
func (r DescribeSnapshotAttributeRequest) Send(ctx context.Context) (*DescribeSnapshotAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSnapshotAttributeResponse{
		DescribeSnapshotAttributeOutput: r.Request.Data.(*DescribeSnapshotAttributeOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSnapshotAttributeResponse is the response type for the
// DescribeSnapshotAttribute API operation.
type DescribeSnapshotAttributeResponse struct {
	*DescribeSnapshotAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSnapshotAttribute request.
func (r *DescribeSnapshotAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
