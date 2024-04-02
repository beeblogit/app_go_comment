package handler

import (
	//"bytes"
	"context"
	//"encoding/base64"
	"encoding/json"
	//"errors"
	//"io"
	//"net/http"
	//"strings"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/beeblogit/app_go_interaction/internal/comment"
	"github.com/digitalhouse-tech/go-lib-kit/request"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/awslambda"
	"github.com/go-kit/log"
	"github.com/ncostamagna/go_http_utils/response"
	"gorm.io/gorm"
)

func NewLambdaCommentStore(endpoints comment.Endpoints) *awslambda.Handler {
	return awslambda.NewHandler(endpoint.Endpoint(endpoints.Create), decodeCommentStoreRequest, EncodeResponse,
		HandlerErrorEncoder(nil), awslambda.HandlerFinalizer(HandlerFinalizer(nil)))
}

func NewLambdaCommentGetAll(endpoints comment.Endpoints) *awslambda.Handler {
	return awslambda.NewHandler(endpoint.Endpoint(endpoints.GetAll), decodeCommentGetAllRequest, EncodeResponse,
		HandlerErrorEncoder(nil), awslambda.HandlerFinalizer(HandlerFinalizer(nil)))
}

func NewLambdaCommentUpdate(endpoints comment.Endpoints) *awslambda.Handler {
	return awslambda.NewHandler(endpoint.Endpoint(endpoints.Create), decodeCommentStoreRequest, EncodeResponse,
		HandlerErrorEncoder(nil), awslambda.HandlerFinalizer(HandlerFinalizer(nil)))
}

func NewLambdaCommentDelete(endpoints comment.Endpoints) *awslambda.Handler {
	return awslambda.NewHandler(endpoint.Endpoint(endpoints.Create), decodeCommentStoreRequest, EncodeResponse,
		HandlerErrorEncoder(nil), awslambda.HandlerFinalizer(HandlerFinalizer(nil)))
}

func decodeCommentStoreRequest(_ context.Context, payload []byte) (interface{}, error) {

	var gateway events.APIGatewayProxyRequest
	err := json.Unmarshal(payload, &gateway)
	if err != nil {
		return nil, response.BadRequest(err.Error())
	}
	var event events.SNSEvent
	err = json.Unmarshal(payload, &event)
	if err != nil {
		return nil, response.BadRequest(err.Error())
	}

	var body string
	switch {
	case gateway.Body != "":
		body = gateway.Body

	case len(event.Records) > 0 && event.Records[0].SNS.Message != "":
		body = event.Records[0].SNS.Message
	default:
		return nil, response.BadRequest("No body received")
	}

	var res comment.CreateReq
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return nil, response.BadRequest(err.Error())
	}
	return res, nil
}

func decodeCommentGetAllRequest(_ context.Context, payload []byte) (interface{}, error) {

	var event events.APIGatewayProxyRequest

	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, response.BadRequest(err.Error())
	}

	var req comment.GetAllReq
	err := request.DecodeMap(event.QueryStringParameters, &req)
	if err != nil {
		return nil, response.BadRequest(err.Error())
	}

	return req, nil
}

func EncodeResponse(_ context.Context, resp interface{}) ([]byte, error) {
	var res response.Response
	switch r := resp.(type) {
	case response.Response:
		res = r
	default:
		res = response.InternalServerError("unknown response type")
	}
	return APIGatewayProxyResponse(res)
}

// HandlerErrorEncoder
func HandlerErrorEncoder(log log.Logger) awslambda.HandlerOption {
	return awslambda.HandlerErrorEncoder(
		awslambda.ErrorEncoder(errorEncoder(log)),
	)
}

// HandlerFinalizer -
func HandlerFinalizer(log log.Logger) func(context.Context, []byte, error) {
	return func(ctx context.Context, resp []byte, err error) {

	}
}

func errorEncoder(log log.Logger) func(context.Context, error) ([]byte, error) {
	return func(_ context.Context, err error) ([]byte, error) {
		res := buildResponse(err, log)
		return APIGatewayProxyResponse(res)
	}
}

// buildResponse builds an error response from an error.
func buildResponse(err error, log log.Logger) response.Response {
	switch r := err.(type) {
	case response.Response:
		return r
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return response.NotFound("")
	}

	return response.InternalServerError("")
}

// APIGatewayProxyResponse
func APIGatewayProxyResponse(res response.Response) ([]byte, error) {
	bytes, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	awsResponse := events.APIGatewayProxyResponse{
		Body:       string(bytes),
		StatusCode: res.StatusCode(),
	}
	return json.Marshal(awsResponse)
}
