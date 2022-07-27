// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateLocalGatewayRouteTableVpcAssociationInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the local gateway route table.
	//
	// LocalGatewayRouteTableId is a required field
	LocalGatewayRouteTableId *string `type:"string" required:"true"`

	// The tags to assign to the local gateway route table VPC association.
	TagSpecifications []TagSpecification `locationName:"TagSpecification" locationNameList:"item" type:"list"`

	// The ID of the VPC.
	//
	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateLocalGatewayRouteTableVpcAssociationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLocalGatewayRouteTableVpcAssociationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateLocalGatewayRouteTableVpcAssociationInput"}

	if s.LocalGatewayRouteTableId == nil {
		invalidParams.Add(aws.NewErrParamRequired("LocalGatewayRouteTableId"))
	}

	if s.VpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateLocalGatewayRouteTableVpcAssociationOutput struct {
	_ struct{} `type:"structure"`

	// Information about the association.
	LocalGatewayRouteTableVpcAssociation *LocalGatewayRouteTableVpcAssociation `locationName:"localGatewayRouteTableVpcAssociation" type:"structure"`
}

// String returns the string representation
func (s CreateLocalGatewayRouteTableVpcAssociationOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateLocalGatewayRouteTableVpcAssociation = "CreateLocalGatewayRouteTableVpcAssociation"

// CreateLocalGatewayRouteTableVpcAssociationRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Associates the specified VPC with the specified local gateway route table.
//
//    // Example sending a request using CreateLocalGatewayRouteTableVpcAssociationRequest.
//    req := client.CreateLocalGatewayRouteTableVpcAssociationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateLocalGatewayRouteTableVpcAssociation
func (c *Client) CreateLocalGatewayRouteTableVpcAssociationRequest(input *CreateLocalGatewayRouteTableVpcAssociationInput) CreateLocalGatewayRouteTableVpcAssociationRequest {
	op := &aws.Operation{
		Name:       opCreateLocalGatewayRouteTableVpcAssociation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateLocalGatewayRouteTableVpcAssociationInput{}
	}

	req := c.newRequest(op, input, &CreateLocalGatewayRouteTableVpcAssociationOutput{})

	return CreateLocalGatewayRouteTableVpcAssociationRequest{Request: req, Input: input, Copy: c.CreateLocalGatewayRouteTableVpcAssociationRequest}
}

// CreateLocalGatewayRouteTableVpcAssociationRequest is the request type for the
// CreateLocalGatewayRouteTableVpcAssociation API operation.
type CreateLocalGatewayRouteTableVpcAssociationRequest struct {
	*aws.Request
	Input *CreateLocalGatewayRouteTableVpcAssociationInput
	Copy  func(*CreateLocalGatewayRouteTableVpcAssociationInput) CreateLocalGatewayRouteTableVpcAssociationRequest
}

// Send marshals and sends the CreateLocalGatewayRouteTableVpcAssociation API request.
func (r CreateLocalGatewayRouteTableVpcAssociationRequest) Send(ctx context.Context) (*CreateLocalGatewayRouteTableVpcAssociationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateLocalGatewayRouteTableVpcAssociationResponse{
		CreateLocalGatewayRouteTableVpcAssociationOutput: r.Request.Data.(*CreateLocalGatewayRouteTableVpcAssociationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateLocalGatewayRouteTableVpcAssociationResponse is the response type for the
// CreateLocalGatewayRouteTableVpcAssociation API operation.
type CreateLocalGatewayRouteTableVpcAssociationResponse struct {
	*CreateLocalGatewayRouteTableVpcAssociationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateLocalGatewayRouteTableVpcAssociation request.
func (r *CreateLocalGatewayRouteTableVpcAssociationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
