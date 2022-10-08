package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dereckdamphouse/html-parser/pkg/log"
	"github.com/dereckdamphouse/html-parser/pkg/psr"
	"github.com/dereckdamphouse/html-parser/pkg/req"
	"github.com/dereckdamphouse/html-parser/pkg/res"
	"go.uber.org/zap"
)

type deps struct {
	parse     func(data *req.Body) (map[string][]string, error)
	marshal   func(v any) ([]byte, error)
	unmarshal func(reqBody string) (*req.Body, error)
}

var initDeps = func(d *deps) error {
	if d.parse == nil {
		d.parse = psr.Parse
	}

	if d.marshal == nil {
		d.marshal = json.Marshal
	}

	if d.unmarshal == nil {
		d.unmarshal = req.Unmarshal
	}

	return nil
}

func (d *deps) handler(ctx context.Context, pxyReq events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	defer func() {
		_ = log.Instance.Sync()
	}()

	reqID := pxyReq.RequestContext.RequestID

	if err := initDeps(d); err != nil {
		log.Instance.Error("unable to init dependencies",
			zap.String("requestId", reqID),
			zap.Error(err))
		return res.StatusInternalServerError, nil
	}

	b, err := d.unmarshal(pxyReq.Body)
	if err != nil {
		log.Instance.Error("failed to unmarshal request body",
			zap.String("requestId", reqID),
			zap.Error(err))
		return res.StatusBadRequest, nil
	}

	data, err := d.parse(b)
	if err != nil {
		log.Instance.Error("failed to parse html",
			zap.String("requestId", reqID),
			zap.Error(err))
		return res.StatusInternalServerError, nil
	}

	jsonStr, err := d.marshal(data)
	if err != nil {
		log.Instance.Error("failed to marshal response body",
			zap.String("requestId", reqID),
			zap.Error(err))
		return res.StatusInternalServerError, nil
	}

	res.StatusOK.Body = string(jsonStr)

	return res.StatusOK, nil
}

func main() {
	d := &deps{}
	lambda.Start(d.handler)
}
