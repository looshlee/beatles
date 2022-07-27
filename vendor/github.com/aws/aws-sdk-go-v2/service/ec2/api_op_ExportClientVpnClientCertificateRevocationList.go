// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ExportClientVpnClientCertificateRevocationListInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Client VPN endpoint.
	//
	// ClientVpnEndpointId is a required field
	ClientVpnEndpointId *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`
}

// String returns the string representation
func (s ExportClientVpnClientCertificateRevocationListInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ExportClientVpnClientCertificateRevocationListInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ExportClientVpnClientCertificateRevocationListInput"}

	if s.ClientVpnEndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientVpnEndpointId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ExportClientVpnClientCertificateRevocationListOutput struct {
	_ struct{} `type:"structure"`

	// Information about the client certificate revocation list.
	CertificateRevocationList *string `locationName:"certificateRevocationList" type:"string"`

	// The current state of the client certificate revocation list.
	Status *CertificateRevocationListStatus `locationName:"status" type:"structure"`
}

// String returns the string representation
func (s ExportClientVpnClientCertificateRevocationListOutput) String() string {
	return awsutil.Prettify(s)
}

const opExportClientVpnClientCertificateRevocationList = "ExportClientVpnClientCertificateRevocationList"

// ExportClientVpnClientCertificateRevocationListRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Downloads the client certificate revocation list for the specified Client
// VPN endpoint.
//
//    // Example sending a request using ExportClientVpnClientCertificateRevocationListRequest.
//    req := client.ExportClientVpnClientCertificateRevocationListRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ExportClientVpnClientCertificateRevocationList
func (c *Client) ExportClientVpnClientCertificateRevocationListRequest(input *ExportClientVpnClientCertificateRevocationListInput) ExportClientVpnClientCertificateRevocationListRequest {
	op := &aws.Operation{
		Name:       opExportClientVpnClientCertificateRevocationList,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ExportClientVpnClientCertificateRevocationListInput{}
	}

	req := c.newRequest(op, input, &ExportClientVpnClientCertificateRevocationListOutput{})
	return ExportClientVpnClientCertificateRevocationListRequest{Request: req, Input: input, Copy: c.ExportClientVpnClientCertificateRevocationListRequest}
}

// ExportClientVpnClientCertificateRevocationListRequest is the request type for the
// ExportClientVpnClientCertificateRevocationList API operation.
type ExportClientVpnClientCertificateRevocationListRequest struct {
	*aws.Request
	Input *ExportClientVpnClientCertificateRevocationListInput
	Copy  func(*ExportClientVpnClientCertificateRevocationListInput) ExportClientVpnClientCertificateRevocationListRequest
}

// Send marshals and sends the ExportClientVpnClientCertificateRevocationList API request.
func (r ExportClientVpnClientCertificateRevocationListRequest) Send(ctx context.Context) (*ExportClientVpnClientCertificateRevocationListResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ExportClientVpnClientCertificateRevocationListResponse{
		ExportClientVpnClientCertificateRevocationListOutput: r.Request.Data.(*ExportClientVpnClientCertificateRevocationListOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ExportClientVpnClientCertificateRevocationListResponse is the response type for the
// ExportClientVpnClientCertificateRevocationList API operation.
type ExportClientVpnClientCertificateRevocationListResponse struct {
	*ExportClientVpnClientCertificateRevocationListOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ExportClientVpnClientCertificateRevocationList request.
func (r *ExportClientVpnClientCertificateRevocationListResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
