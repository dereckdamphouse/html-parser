package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dereckdamphouse/html-parser/pkg/html"
	"github.com/dereckdamphouse/html-parser/pkg/log"
	"github.com/dereckdamphouse/html-parser/pkg/req"
	"github.com/dereckdamphouse/html-parser/pkg/resp"
	"go.uber.org/zap"
)

type deps struct {
	parse     func(data *req.Data) (map[string][]string, error)
	marshal   func(v any) ([]byte, error)
	unmarshal func(body string) (*req.Data, error)
}

func (d *deps) init() {
	if d.parse == nil {
		d.parse = html.Parse
	}

	if d.marshal == nil {
		d.marshal = json.Marshal
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
		log.Instance.Error("failed to unmarshal request body",
			zap.String("requestId", reqID),
			zap.Error(err))
		return resp.StatusBadRequest, nil
	}

	res, err := d.parse(data)
	if err != nil {
		log.Instance.Error("failed to parse html",
			zap.String("requestId", reqID),
			zap.Error(err))
		return resp.StatusBadRequest, nil
	}

	jsonByte, err := d.marshal(res)
	if err != nil {
		log.Instance.Error("failed to marshal response body",
			zap.String("requestId", reqID),
			zap.Error(err))
		return resp.StatusInternalServerError, nil
	}

	resp.StatusOK.Body = string(jsonByte)

	return resp.StatusOK, nil
}

func main() {
	d := &deps{}

	d.init()

	lambda.Start(d.handler)
}
