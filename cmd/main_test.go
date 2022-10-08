package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/dereckdamphouse/html-parser/pkg/req"
	"github.com/dereckdamphouse/html-parser/pkg/res"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	tt := []struct {
		name     string
		d        *deps
		initDeps func(d *deps) error
		resBody  string
		resCode  int
		err      error
	}{
		{
			"handles initDeps error",
			&deps{},
			func(d *deps) error {
				return fmt.Errorf("some error")
			},
			res.DefaultBody,
			500,
			nil,
		},
		{
			"handles unmarshal error",
			&deps{
				unmarshal: func(reqBody string) (*req.Body, error) {
					return &req.Body{}, fmt.Errorf("some error")
				},
			},
			func(d *deps) error {
				return nil
			},
			res.DefaultBody,
			400,
			nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			initDeps = tc.initDeps
			res, err := tc.d.handler(
				context.TODO(),
				events.APIGatewayProxyRequest{},
			)
			assert.Equal(t, tc.resBody, res.Body)
			assert.Equal(t, tc.resCode, res.StatusCode)
			assert.Equal(t, tc.err, err)
		})
	}
}
