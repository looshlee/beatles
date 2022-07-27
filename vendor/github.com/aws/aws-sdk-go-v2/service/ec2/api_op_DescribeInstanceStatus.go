// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeInstanceStatusInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The filters.
	//
	//    * availability-zone - The Availability Zone of the instance.
	//
	//    * event.code - The code for the scheduled event (instance-reboot | system-reboot
	//    | system-maintenance | instance-retirement | instance-stop).
	//
	//    * event.description - A description of the event.
	//
	//    * event.instance-event-id - The ID of the event whose date and time you
	//    are modifying.
	//
	//    * event.not-after - The latest end time for the scheduled event (for example,
	//    2014-09-15T17:15:20.000Z).
	//
	//    * event.not-before - The earliest start time for the scheduled event (for
	//    example, 2014-09-15T17:15:20.000Z).
	//
	//    * event.not-before-deadline - The deadline for starting the event (for
	//    example, 2014-09-15T17:15:20.000Z).
	//
	//    * instance-state-code - The code for the instance state, as a 16-bit unsigned
	//    integer. The high byte is used for internal purposes and should be ignored.
	//    The low byte is set based on the state represented. The valid values are
	//    0 (pending), 16 (running), 32 (shutting-down), 48 (terminated), 64 (stopping),
	//    and 80 (stopped).
	//
	//    * instance-state-name - The state of the instance (pending | running |
	//    shutting-down | terminated | stopping | stopped).
	//
	//    * instance-status.reachability - Filters on instance status where the
	//    name is reachability (passed | failed | initializing | insufficient-data).
	//
	//    * instance-status.status - The status of the instance (ok | impaired |
	//    initializing | insufficient-data | not-applicable).
	//
	//    * system-status.reachability - Filters on system status where the name
	//    is reachability (passed | failed | initializing | insufficient-data).
	//
	//    * system-status.status - The system status of the instance (ok | impaired
	//    | initializing | insufficient-data | not-applicable).
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// When true, includes the health status for all instances. When false, includes
	// the health status for running instances only.
	//
	// Default: false
	IncludeAllInstances *bool `locationName:"includeAllInstances" type:"boolean"`

	// The instance IDs.
	//
	// Default: Describes all your instances.
	//
	// Constraints: Maximum 100 explicitly specified instance IDs.
	InstanceIds []string `locationName:"InstanceId" locationNameList:"InstanceId" type:"list"`

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value. This
	// value can be between 5 and 1000. You cannot specify this parameter and the
	// instance IDs parameter in the same call.
	MaxResults *int64 `type:"integer"`

	// The token to retrieve the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeInstanceStatusInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeInstanceStatusOutput struct {
	_ struct{} `type:"structure"`

	// Information about the status of the instances.
	InstanceStatuses []InstanceStatus `locationName:"instanceStatusSet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeInstanceStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeInstanceStatus = "DescribeInstanceStatus"

// DescribeInstanceStatusRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the status of the specified instances or all of your instances.
// By default, only running instances are described, unless you specifically
// indicate to return the status of all instances.
//
// Instance status includes the following components:
//
//    * Status checks - Amazon EC2 performs status checks on running EC2 instances
//    to identify hardware and software issues. For more information, see Status
//    Checks for Your Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-system-instance-status-check.html)
//    and Troubleshooting Instances with Failed Status Checks (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/TroubleshootingInstances.html)
//    in the Amazon Elastic Compute Cloud User Guide.
//
//    * Scheduled events - Amazon EC2 can schedule events (such as reboot, stop,
//    or terminate) for your instances related to hardware issues, software
//    updates, or system maintenance. For more information, see Scheduled Events
//    for Your Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-instances-status-check_sched.html)
//    in the Amazon Elastic Compute Cloud User Guide.
//
//    * Instance state - You can manage your instances from the moment you launch
//    them through their termination. For more information, see Instance Lifecycle
//    (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-lifecycle.html)
//    in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DescribeInstanceStatusRequest.
//    req := client.DescribeInstanceStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeInstanceStatus
func (c *Client) DescribeInstanceStatusRequest(input *DescribeInstanceStatusInput) DescribeInstanceStatusRequest {
	op := &aws.Operation{
		Name:       opDescribeInstanceStatus,
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
		input = &DescribeInstanceStatusInput{}
	}

	req := c.newRequest(op, input, &DescribeInstanceStatusOutput{})
	return DescribeInstanceStatusRequest{Request: req, Input: input, Copy: c.DescribeInstanceStatusRequest}
}

// DescribeInstanceStatusRequest is the request type for the
// DescribeInstanceStatus API operation.
type DescribeInstanceStatusRequest struct {
	*aws.Request
	Input *DescribeInstanceStatusInput
	Copy  func(*DescribeInstanceStatusInput) DescribeInstanceStatusRequest
}

// Send marshals and sends the DescribeInstanceStatus API request.
func (r DescribeInstanceStatusRequest) Send(ctx context.Context) (*DescribeInstanceStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeInstanceStatusResponse{
		DescribeInstanceStatusOutput: r.Request.Data.(*DescribeInstanceStatusOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeInstanceStatusRequestPaginator returns a paginator for DescribeInstanceStatus.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeInstanceStatusRequest(input)
//   p := ec2.NewDescribeInstanceStatusRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeInstanceStatusPaginator(req DescribeInstanceStatusRequest) DescribeInstanceStatusPaginator {
	return DescribeInstanceStatusPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeInstanceStatusInput
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

// DescribeInstanceStatusPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeInstanceStatusPaginator struct {
	aws.Pager
}

func (p *DescribeInstanceStatusPaginator) CurrentPage() *DescribeInstanceStatusOutput {
	return p.Pager.CurrentPage().(*DescribeInstanceStatusOutput)
}

// DescribeInstanceStatusResponse is the response type for the
// DescribeInstanceStatus API operation.
type DescribeInstanceStatusResponse struct {
	*DescribeInstanceStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeInstanceStatus request.
func (r *DescribeInstanceStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
