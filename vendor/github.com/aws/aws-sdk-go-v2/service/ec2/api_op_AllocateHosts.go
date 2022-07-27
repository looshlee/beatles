// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AllocateHostsInput struct {
	_ struct{} `type:"structure"`

	// Indicates whether the host accepts any untargeted instance launches that
	// match its instance type configuration, or if it only accepts Host tenancy
	// instance launches that specify its unique host ID. For more information,
	// see Understanding Instance Placement and Host Affinity (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/how-dedicated-hosts-work.html#dedicated-hosts-understanding)
	// in the Amazon EC2 User Guide for Linux Instances.
	//
	// Default: on
	AutoPlacement AutoPlacement `locationName:"autoPlacement" type:"string" enum:"true"`

	// The Availability Zone in which to allocate the Dedicated Host.
	//
	// AvailabilityZone is a required field
	AvailabilityZone *string `locationName:"availabilityZone" type:"string" required:"true"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. For more information, see How to Ensure Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string `locationName:"clientToken" type:"string"`

	// Indicates whether to enable or disable host recovery for the Dedicated Host.
	// Host recovery is disabled by default. For more information, see Host Recovery
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-recovery.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	// Default: off
	HostRecovery HostRecovery `type:"string" enum:"true"`

	// Specifies the instance family to be supported by the Dedicated Hosts. If
	// you specify an instance family, the Dedicated Hosts support multiple instance
	// types within that instance family.
	//
	// If you want the Dedicated Hosts to support a specific instance type only,
	// omit this parameter and specify InstanceType instead. You cannot specify
	// InstanceFamily and InstanceType in the same request.
	InstanceFamily *string `type:"string"`

	// Specifies the instance type to be supported by the Dedicated Hosts. If you
	// specify an instance type, the Dedicated Hosts support instances of the specified
	// instance type only.
	//
	// If you want the Dedicated Hosts to support multiple instance types in a specific
	// instance family, omit this parameter and specify InstanceFamily instead.
	// You cannot specify InstanceType and InstanceFamily in the same request.
	InstanceType *string `locationName:"instanceType" type:"string"`

	// The number of Dedicated Hosts to allocate to your account with these parameters.
	//
	// Quantity is a required field
	Quantity *int64 `locationName:"quantity" type:"integer" required:"true"`

	// The tags to apply to the Dedicated Host during creation.
	TagSpecifications []TagSpecification `locationName:"TagSpecification" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s AllocateHostsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AllocateHostsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AllocateHostsInput"}

	if s.AvailabilityZone == nil {
		invalidParams.Add(aws.NewErrParamRequired("AvailabilityZone"))
	}

	if s.Quantity == nil {
		invalidParams.Add(aws.NewErrParamRequired("Quantity"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of AllocateHosts.
type AllocateHostsOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the allocated Dedicated Host. This is used to launch an instance
	// onto a specific host.
	HostIds []string `locationName:"hostIdSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s AllocateHostsOutput) String() string {
	return awsutil.Prettify(s)
}

const opAllocateHosts = "AllocateHosts"

// AllocateHostsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Allocates a Dedicated Host to your account. At a minimum, specify the supported
// instance type or instance family, the Availability Zone in which to allocate
// the host, and the number of hosts to allocate.
//
//    // Example sending a request using AllocateHostsRequest.
//    req := client.AllocateHostsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AllocateHosts
func (c *Client) AllocateHostsRequest(input *AllocateHostsInput) AllocateHostsRequest {
	op := &aws.Operation{
		Name:       opAllocateHosts,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AllocateHostsInput{}
	}

	req := c.newRequest(op, input, &AllocateHostsOutput{})
	return AllocateHostsRequest{Request: req, Input: input, Copy: c.AllocateHostsRequest}
}

// AllocateHostsRequest is the request type for the
// AllocateHosts API operation.
type AllocateHostsRequest struct {
	*aws.Request
	Input *AllocateHostsInput
	Copy  func(*AllocateHostsInput) AllocateHostsRequest
}

// Send marshals and sends the AllocateHosts API request.
func (r AllocateHostsRequest) Send(ctx context.Context) (*AllocateHostsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AllocateHostsResponse{
		AllocateHostsOutput: r.Request.Data.(*AllocateHostsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AllocateHostsResponse is the response type for the
// AllocateHosts API operation.
type AllocateHostsResponse struct {
	*AllocateHostsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AllocateHosts request.
func (r *AllocateHostsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
