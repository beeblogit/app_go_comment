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

	"github.com/aws/aws-lambda-go/events"
	"github.com/beeblogit/app_go_interaction/internal/comment"
	//"github.com/digitalhouse-tech/go-lib-kit/request"
	"github.com/digitalhouse-tech/go-lib-kit/response"
	"github.com/digitalhouse-tech/go-lib-util/lambda"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/awslambda"
)

func NewLambdaCommentStore(endpoints comment.Endpoints) *awslambda.Handler {
	return awslambda.NewHandler(endpoint.Endpoint(endpoints.Create), decodeCommentStoreRequest, lambda.EncodeResponse,
		lambda.HandlerErrorEncoder(nil), awslambda.HandlerFinalizer(lambda.HandlerFinalizer(nil)))
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
