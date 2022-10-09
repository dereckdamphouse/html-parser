package resp

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
)

const DefaultBody = "{}"

var (
	headers = map[string]string{
		"Content-Type":                "application/json",
		"Access-Control-Allow-Origin": allowOrigin(),
	}
	StatusOK = events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       DefaultBody,
		Headers:    headers,
	}
	StatusInternalServerError = events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       DefaultBody,
		Headers:    headers,
	}
	StatusBadRequest = events.APIGatewayProxyResponse{
		StatusCode: 400,
		Body:       DefaultBody,
		Headers:    headers,
	}
)

func allowOrigin() string {
	ao := os.Getenv("ALLOWORIGIN")
	if len(ao) < 2 {
		return ""
	}

	return ao[1 : len(ao)-1]
}
