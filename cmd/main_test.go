package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/dereckdamphouse/html-parser/pkg/req"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	d := &deps{}
	d.init()
	assert.NotNil(t, d.marshal)
	assert.NotNil(t, d.parse)
	assert.NotNil(t, d.unmarshal)
}

func TestHandler(t *testing.T) {
	tt := []struct {
		name string
		d    *deps
		body string
		code int
		err  error
	}{
		{
			"handles unmarshal error",
			&deps{
				unmarshal: func(body string) (*req.Data, error) {
					return &req.Data{}, fmt.Errorf("some error")
				},
			},
			"{\"error\":\"failed to marshal request body ('html' or 'properties' field may be missing)\",\"found\":{}}",
			400,
			nil,
		},
		{
			"handles parsing error",
			&deps{
				unmarshal: func(body string) (*req.Data, error) {
					return &req.Data{}, nil
				},
				parse: func(data *req.Data) (map[string][]string, error) {
					return map[string][]string{}, fmt.Errorf("some error")
				},
			},
			"{\"error\":\"failed to parse html\",\"found\":{}}",
			400,
			nil,
		},
		{
			"handles marshal error",
			&deps{
				unmarshal: func(body string) (*req.Data, error) {
					return &req.Data{}, nil
				},
				parse: func(data *req.Data) (map[string][]string, error) {
					return map[string][]string{}, nil
				},
				marshal: func(found any) (string, error) {
					return "", fmt.Errorf("some error")
				},
			},
			"{\"error\":\"failed to marshal response body\",\"found\":{}}",
			500,
			nil,
		},
		{
			"handles successfull response",
			&deps{
				unmarshal: func(body string) (*req.Data, error) {
					return &req.Data{}, nil
				},
				parse: func(data *req.Data) (map[string][]string, error) {
					return map[string][]string{}, nil
				},
				marshal: func(found any) (string, error) {
					return "success!", nil
				},
			},
			"success!",
			200,
			nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res, err := tc.d.handler(
				context.TODO(),
				events.APIGatewayProxyRequest{},
			)
			assert.Equal(t, tc.body, res.Body)
			assert.Equal(t, tc.code, res.StatusCode)
			assert.Equal(t, tc.err, err)
		})
	}
}
