package resp

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

var headers = map[string]string{
	"Content-Type":                "application/json",
	"Access-Control-Allow-Origin": allowOrigin(),
}

func Error(code int, msg string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Body:       fmt.Sprintf("{\"error\":\"%s\",\"found\":{}}", msg),
		Headers:    headers,
	}
}

func Success(body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       body,
		Headers:    headers,
	}
}

func allowOrigin() string {
	ao := os.Getenv("ALLOWORIGIN")
	if len(ao) < 2 {
		return ""
	}

	return ao[1 : len(ao)-1]
}
