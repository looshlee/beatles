// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeExportTasksInput struct {
	_ struct{} `type:"structure"`

	// The export task IDs.
	ExportTaskIds []string `locationName:"exportTaskId" locationNameList:"ExportTaskId" type:"list"`
}

// String returns the string representation
func (s DescribeExportTasksInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeExportTasksOutput struct {
	_ struct{} `type:"structure"`

	// Information about the export tasks.
	ExportTasks []ExportTask `locationName:"exportTaskSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeExportTasksOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeExportTasks = "DescribeExportTasks"

// DescribeExportTasksRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the specified export instance tasks or all your export instance
// tasks.
//
//    // Example sending a request using DescribeExportTasksRequest.
//    req := client.DescribeExportTasksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeExportTasks
func (c *Client) DescribeExportTasksRequest(input *DescribeExportTasksInput) DescribeExportTasksRequest {
	op := &aws.Operation{
		Name:       opDescribeExportTasks,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeExportTasksInput{}
	}

	req := c.newRequest(op, input, &DescribeExportTasksOutput{})
	return DescribeExportTasksRequest{Request: req, Input: input, Copy: c.DescribeExportTasksRequest}
}

// DescribeExportTasksRequest is the request type for the
// DescribeExportTasks API operation.
type DescribeExportTasksRequest struct {
	*aws.Request
	Input *DescribeExportTasksInput
	Copy  func(*DescribeExportTasksInput) DescribeExportTasksRequest
}

// Send marshals and sends the DescribeExportTasks API request.
func (r DescribeExportTasksRequest) Send(ctx context.Context) (*DescribeExportTasksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeExportTasksResponse{
		DescribeExportTasksOutput: r.Request.Data.(*DescribeExportTasksOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeExportTasksResponse is the response type for the
// DescribeExportTasks API operation.
type DescribeExportTasksResponse struct {
	*DescribeExportTasksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeExportTasks request.
func (r *DescribeExportTasksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
