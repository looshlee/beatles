// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyCapacityReservationInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Capacity Reservation.
	//
	// CapacityReservationId is a required field
	CapacityReservationId *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The date and time at which the Capacity Reservation expires. When a Capacity
	// Reservation expires, the reserved capacity is released and you can no longer
	// launch instances into it. The Capacity Reservation's state changes to expired
	// when it reaches its end date and time.
	//
	// The Capacity Reservation is cancelled within an hour from the specified time.
	// For example, if you specify 5/31/2019, 13:30:55, the Capacity Reservation
	// is guaranteed to end between 13:30:55 and 14:30:55 on 5/31/2019.
	//
	// You must provide an EndDate value if EndDateType is limited. Omit EndDate
	// if EndDateType is unlimited.
	EndDate *time.Time `type:"timestamp"`

	// Indicates the way in which the Capacity Reservation ends. A Capacity Reservation
	// can have one of the following end types:
	//
	//    * unlimited - The Capacity Reservation remains active until you explicitly
	//    cancel it. Do not provide an EndDate value if EndDateType is unlimited.
	//
	//    * limited - The Capacity Reservation expires automatically at a specified
	//    date and time. You must provide an EndDate value if EndDateType is limited.
	EndDateType EndDateType `type:"string" enum:"true"`

	// The number of instances for which to reserve capacity.
	InstanceCount *int64 `type:"integer"`
}

// String returns the string representation
func (s ModifyCapacityReservationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyCapacityReservationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyCapacityReservationInput"}

	if s.CapacityReservationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CapacityReservationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyCapacityReservationOutput struct {
	_ struct{} `type:"structure"`

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s ModifyCapacityReservationOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyCapacityReservation = "ModifyCapacityReservation"

// ModifyCapacityReservationRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies a Capacity Reservation's capacity and the conditions under which
// it is to be released. You cannot change a Capacity Reservation's instance
// type, EBS optimization, instance store settings, platform, Availability Zone,
// or instance eligibility. If you need to modify any of these attributes, we
// recommend that you cancel the Capacity Reservation, and then create a new
// one with the required attributes.
//
//    // Example sending a request using ModifyCapacityReservationRequest.
//    req := client.ModifyCapacityReservationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyCapacityReservation
func (c *Client) ModifyCapacityReservationRequest(input *ModifyCapacityReservationInput) ModifyCapacityReservationRequest {
	op := &aws.Operation{
		Name:       opModifyCapacityReservation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyCapacityReservationInput{}
	}

	req := c.newRequest(op, input, &ModifyCapacityReservationOutput{})
	return ModifyCapacityReservationRequest{Request: req, Input: input, Copy: c.ModifyCapacityReservationRequest}
}

// ModifyCapacityReservationRequest is the request type for the
// ModifyCapacityReservation API operation.
type ModifyCapacityReservationRequest struct {
	*aws.Request
	Input *ModifyCapacityReservationInput
	Copy  func(*ModifyCapacityReservationInput) ModifyCapacityReservationRequest
}

// Send marshals and sends the ModifyCapacityReservation API request.
func (r ModifyCapacityReservationRequest) Send(ctx context.Context) (*ModifyCapacityReservationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyCapacityReservationResponse{
		ModifyCapacityReservationOutput: r.Request.Data.(*ModifyCapacityReservationOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyCapacityReservationResponse is the response type for the
// ModifyCapacityReservation API operation.
type ModifyCapacityReservationResponse struct {
	*ModifyCapacityReservationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyCapacityReservation request.
func (r *ModifyCapacityReservationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
