package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dereckdamphouse/html-parser/pkg/html"
	"github.com/dereckdamphouse/html-parser/pkg/log"
	"github.com/dereckdamphouse/html-parser/pkg/req"
	"github.com/dereckdamphouse/html-parser/pkg/resp"
	"go.uber.org/zap"
)

type deps struct {
	parse     func(data *req.Data) (html.Parsed, error)
	marshal   func(parsed html.Parsed) (string, error)
	unmarshal func(body string) (*req.Data, error)
}

func (d *deps) init() {
	if d.parse == nil {
		d.parse = html.Parse
	}

	if d.marshal == nil {
		d.marshal = resp.Marshal
	}

	if d.unmarshal == nil {
		d.unmarshal = req.Unmarshal
	}
}

func (d *deps) handler(ctx context.Context, pxyReq events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	defer func() {
		_ = log.Instance.Sync()
	}()

	reqID := pxyReq.RequestContext.RequestID

	data, err := d.unmarshal(pxyReq.Body)
	if err != nil {
		errMsg := "failed to marshal request body ('html' or 'properties' field may be missing)"
		log.Instance.Error(errMsg, zap.String("requestId", reqID), zap.Error(err))
		return resp.Error(http.StatusBadRequest, errMsg), nil
	}

	parsed, err := d.parse(data)
	if err != nil {
		errMsg := "failed to parse html"
		log.Instance.Error(errMsg, zap.String("requestId", reqID), zap.Error(err))
		return resp.Error(http.StatusBadRequest, errMsg), nil
	}

	if len(parsed) == 0 {
		errMsg := "no properties found"
		log.Instance.Error(errMsg, zap.String("requestId", reqID), zap.Error(err))
		return resp.Error(http.StatusBadRequest, errMsg), nil
	}

	body, err := d.marshal(parsed)
	if err != nil {
		errMsg := "failed to marshal response body"
		log.Instance.Error(errMsg, zap.String("requestId", reqID), zap.Error(err))
		return resp.Error(http.StatusInternalServerError, errMsg), nil
	}

	return resp.Success(body), nil
}

func main() {
	d := &deps{}

	d.init()

	lambda.Start(d.handler)
}
