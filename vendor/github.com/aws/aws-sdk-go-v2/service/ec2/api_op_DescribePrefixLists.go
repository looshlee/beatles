// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes available AWS services in a prefix list format, which includes the
// prefix list name and prefix list ID of the service and the IP address range for
// the service. We recommend that you use DescribeManagedPrefixLists instead.
func (c *Client) DescribePrefixLists(ctx context.Context, params *DescribePrefixListsInput, optFns ...func(*Options)) (*DescribePrefixListsOutput, error) {
	if params == nil {
		params = &DescribePrefixListsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePrefixLists", params, optFns, addOperationDescribePrefixListsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePrefixListsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePrefixListsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// One or more filters.
	//
	// * prefix-list-id: The ID of a prefix list.
	//
	// *
	// prefix-list-name: The name of a prefix list.
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults int32

	// The token for the next page of results.
	NextToken *string

	// One or more prefix list IDs.
	PrefixListIds []string
}

type DescribePrefixListsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// All available prefix lists.
	PrefixLists []types.PrefixList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribePrefixListsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribePrefixLists{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribePrefixLists{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePrefixLists(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribePrefixListsAPIClient is a client that implements the DescribePrefixLists
// operation.
type DescribePrefixListsAPIClient interface {
	DescribePrefixLists(context.Context, *DescribePrefixListsInput, ...func(*Options)) (*DescribePrefixListsOutput, error)
}

var _ DescribePrefixListsAPIClient = (*Client)(nil)

// DescribePrefixListsPaginatorOptions is the paginator options for
// DescribePrefixLists
type DescribePrefixListsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribePrefixListsPaginator is a paginator for DescribePrefixLists
type DescribePrefixListsPaginator struct {
	options   DescribePrefixListsPaginatorOptions
	client    DescribePrefixListsAPIClient
	params    *DescribePrefixListsInput
	nextToken *string
	firstPage bool
}

// NewDescribePrefixListsPaginator returns a new DescribePrefixListsPaginator
func NewDescribePrefixListsPaginator(client DescribePrefixListsAPIClient, params *DescribePrefixListsInput, optFns ...func(*DescribePrefixListsPaginatorOptions)) *DescribePrefixListsPaginator {
	options := DescribePrefixListsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribePrefixListsInput{}
	}

	return &DescribePrefixListsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribePrefixListsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribePrefixLists page.
func (p *DescribePrefixListsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribePrefixListsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribePrefixLists(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribePrefixLists(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribePrefixLists",
	}
}
