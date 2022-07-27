// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetTransitGatewayAttachmentPropagationsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters. The possible values are:
	//
	//    * transit-gateway-route-table-id - The ID of the transit gateway route
	//    table.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token for the next page of results.
	NextToken *string `type:"string"`

	// The ID of the attachment.
	//
	// TransitGatewayAttachmentId is a required field
	TransitGatewayAttachmentId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetTransitGatewayAttachmentPropagationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTransitGatewayAttachmentPropagationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTransitGatewayAttachmentPropagationsInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if s.TransitGatewayAttachmentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayAttachmentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetTransitGatewayAttachmentPropagationsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the propagation route tables.
	TransitGatewayAttachmentPropagations []TransitGatewayAttachmentPropagation `locationName:"transitGatewayAttachmentPropagations" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s GetTransitGatewayAttachmentPropagationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetTransitGatewayAttachmentPropagations = "GetTransitGatewayAttachmentPropagations"

// GetTransitGatewayAttachmentPropagationsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Lists the route tables to which the specified resource attachment propagates
// routes.
//
//    // Example sending a request using GetTransitGatewayAttachmentPropagationsRequest.
//    req := client.GetTransitGatewayAttachmentPropagationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/GetTransitGatewayAttachmentPropagations
func (c *Client) GetTransitGatewayAttachmentPropagationsRequest(input *GetTransitGatewayAttachmentPropagationsInput) GetTransitGatewayAttachmentPropagationsRequest {
	op := &aws.Operation{
		Name:       opGetTransitGatewayAttachmentPropagations,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetTransitGatewayAttachmentPropagationsInput{}
	}

	req := c.newRequest(op, input, &GetTransitGatewayAttachmentPropagationsOutput{})
	return GetTransitGatewayAttachmentPropagationsRequest{Request: req, Input: input, Copy: c.GetTransitGatewayAttachmentPropagationsRequest}
}

// GetTransitGatewayAttachmentPropagationsRequest is the request type for the
// GetTransitGatewayAttachmentPropagations API operation.
type GetTransitGatewayAttachmentPropagationsRequest struct {
	*aws.Request
	Input *GetTransitGatewayAttachmentPropagationsInput
	Copy  func(*GetTransitGatewayAttachmentPropagationsInput) GetTransitGatewayAttachmentPropagationsRequest
}

// Send marshals and sends the GetTransitGatewayAttachmentPropagations API request.
func (r GetTransitGatewayAttachmentPropagationsRequest) Send(ctx context.Context) (*GetTransitGatewayAttachmentPropagationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTransitGatewayAttachmentPropagationsResponse{
		GetTransitGatewayAttachmentPropagationsOutput: r.Request.Data.(*GetTransitGatewayAttachmentPropagationsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetTransitGatewayAttachmentPropagationsRequestPaginator returns a paginator for GetTransitGatewayAttachmentPropagations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetTransitGatewayAttachmentPropagationsRequest(input)
//   p := ec2.NewGetTransitGatewayAttachmentPropagationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetTransitGatewayAttachmentPropagationsPaginator(req GetTransitGatewayAttachmentPropagationsRequest) GetTransitGatewayAttachmentPropagationsPaginator {
	return GetTransitGatewayAttachmentPropagationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetTransitGatewayAttachmentPropagationsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// GetTransitGatewayAttachmentPropagationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetTransitGatewayAttachmentPropagationsPaginator struct {
	aws.Pager
}

func (p *GetTransitGatewayAttachmentPropagationsPaginator) CurrentPage() *GetTransitGatewayAttachmentPropagationsOutput {
	return p.Pager.CurrentPage().(*GetTransitGatewayAttachmentPropagationsOutput)
}

// GetTransitGatewayAttachmentPropagationsResponse is the response type for the
// GetTransitGatewayAttachmentPropagations API operation.
type GetTransitGatewayAttachmentPropagationsResponse struct {
	*GetTransitGatewayAttachmentPropagationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTransitGatewayAttachmentPropagations request.
func (r *GetTransitGatewayAttachmentPropagationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
