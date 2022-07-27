// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

type ModifySubnetAttributeInput struct {
	_ struct{} `type:"structure"`

	// Specify true to indicate that network interfaces created in the specified
	// subnet should be assigned an IPv6 address. This includes a network interface
	// that's created when launching an instance into the subnet (the instance therefore
	// receives an IPv6 address).
	//
	// If you enable the IPv6 addressing feature for your subnet, your network interface
	// or instance only receives an IPv6 address if it's created using version 2016-11-15
	// or later of the Amazon EC2 API.
	AssignIpv6AddressOnCreation *AttributeBooleanValue `type:"structure"`

	// Specify true to indicate that ENIs attached to instances created in the specified
	// subnet should be assigned a public IPv4 address.
	MapPublicIpOnLaunch *AttributeBooleanValue `type:"structure"`

	// The ID of the subnet.
	//
	// SubnetId is a required field
	SubnetId *string `locationName:"subnetId" type:"string" required:"true"`
}

// String returns the string representation
func (s ModifySubnetAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifySubnetAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifySubnetAttributeInput"}

	if s.SubnetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifySubnetAttributeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ModifySubnetAttributeOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifySubnetAttribute = "ModifySubnetAttribute"

// ModifySubnetAttributeRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies a subnet attribute. You can only modify one attribute at a time.
//
//    // Example sending a request using ModifySubnetAttributeRequest.
//    req := client.ModifySubnetAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifySubnetAttribute
func (c *Client) ModifySubnetAttributeRequest(input *ModifySubnetAttributeInput) ModifySubnetAttributeRequest {
	op := &aws.Operation{
		Name:       opModifySubnetAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifySubnetAttributeInput{}
	}

	req := c.newRequest(op, input, &ModifySubnetAttributeOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return ModifySubnetAttributeRequest{Request: req, Input: input, Copy: c.ModifySubnetAttributeRequest}
}

// ModifySubnetAttributeRequest is the request type for the
// ModifySubnetAttribute API operation.
type ModifySubnetAttributeRequest struct {
	*aws.Request
	Input *ModifySubnetAttributeInput
	Copy  func(*ModifySubnetAttributeInput) ModifySubnetAttributeRequest
}

// Send marshals and sends the ModifySubnetAttribute API request.
func (r ModifySubnetAttributeRequest) Send(ctx context.Context) (*ModifySubnetAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifySubnetAttributeResponse{
		ModifySubnetAttributeOutput: r.Request.Data.(*ModifySubnetAttributeOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifySubnetAttributeResponse is the response type for the
// ModifySubnetAttribute API operation.
type ModifySubnetAttributeResponse struct {
	*ModifySubnetAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifySubnetAttribute request.
func (r *ModifySubnetAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
