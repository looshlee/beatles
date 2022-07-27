// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateSnapshotsInput struct {
	_ struct{} `type:"structure"`

	// Copies the tags from the specified volume to corresponding snapshot.
	CopyTagsFromSource CopyTagsFromSource `type:"string" enum:"true"`

	// A description propagated to every snapshot specified by the instance.
	Description *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The instance to specify which volumes should be included in the snapshots.
	//
	// InstanceSpecification is a required field
	InstanceSpecification *InstanceSpecification `type:"structure" required:"true"`

	// Tags to apply to every snapshot specified by the instance.
	TagSpecifications []TagSpecification `locationName:"TagSpecification" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s CreateSnapshotsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSnapshotsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSnapshotsInput"}

	if s.InstanceSpecification == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceSpecification"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateSnapshotsOutput struct {
	_ struct{} `type:"structure"`

	// List of snapshots.
	Snapshots []SnapshotInfo `locationName:"snapshotSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s CreateSnapshotsOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateSnapshots = "CreateSnapshots"

// CreateSnapshotsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates crash-consistent snapshots of multiple EBS volumes and stores the
// data in S3. Volumes are chosen by specifying an instance. Any attached volumes
// will produce one snapshot each that is crash-consistent across the instance.
// Boot volumes can be excluded by changing the parameters.
//
//    // Example sending a request using CreateSnapshotsRequest.
//    req := client.CreateSnapshotsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateSnapshots
func (c *Client) CreateSnapshotsRequest(input *CreateSnapshotsInput) CreateSnapshotsRequest {
	op := &aws.Operation{
		Name:       opCreateSnapshots,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateSnapshotsInput{}
	}

	req := c.newRequest(op, input, &CreateSnapshotsOutput{})
	return CreateSnapshotsRequest{Request: req, Input: input, Copy: c.CreateSnapshotsRequest}
}

// CreateSnapshotsRequest is the request type for the
// CreateSnapshots API operation.
type CreateSnapshotsRequest struct {
	*aws.Request
	Input *CreateSnapshotsInput
	Copy  func(*CreateSnapshotsInput) CreateSnapshotsRequest
}

// Send marshals and sends the CreateSnapshots API request.
func (r CreateSnapshotsRequest) Send(ctx context.Context) (*CreateSnapshotsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSnapshotsResponse{
		CreateSnapshotsOutput: r.Request.Data.(*CreateSnapshotsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSnapshotsResponse is the response type for the
// CreateSnapshots API operation.
type CreateSnapshotsResponse struct {
	*CreateSnapshotsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSnapshots request.
func (r *CreateSnapshotsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
