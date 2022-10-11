package html

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/dereckdamphouse/html-parser/pkg/req"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tt := []struct {
		name       string
		htmlReader func(r io.Reader) (*goquery.Document, error)
		data       *req.Data
		parsed     Parsed
		err        error
	}{
		{
			"handles reader error",
			func(r io.Reader) (*goquery.Document, error) {
				return nil, fmt.Errorf("some error")
			},
			&req.Data{},
			Parsed{},
			fmt.Errorf("some error"),
		},
		{
			"handles missing property name or selector error",
			goquery.NewDocumentFromReader,
			&req.Data{
				Properties: []req.Property{
					{
						Name:     "",
						Selector: ".title",
					},
					{
						Name:     "title",
						Selector: "",
					},
				},
			},
			Parsed{},
			nil,
		},
		{
			"handles finding element text and attribute",
			goquery.NewDocumentFromReader,
			func() *req.Data {
				// https://movableink-inkredible-retail.herokuapp.com/product/2599191
				stub, err := os.ReadFile("html_stub.txt")
				if err != nil {
					log.Panic(err)
				}

				return &req.Data{
					HTML: string(stub),
					Properties: []req.Property{
						{
							Name:     "title",
							Selector: ".video-info h2",
						},
						{
							Name:      "image",
							Selector:  "img.ProductImage",
							Attribute: "src",
						},
						{
							Name:     "fabric",
							Selector: ".video-info > div *:nth-child(8) li",
						},
					},
				}
			}(),
			Parsed{
				"title":  {"Men's Columbia Flattop Ridge Fleece Jacket"},
				"image":  {"/images/clothing/2599191_ALT-1000.jpg"},
				"fabric": {"Polyester fleece", "Machine wash", "Imported"},
			},
			nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			htmlReader = tc.htmlReader
			p, err := Parse(tc.data)
			assert.Equal(t, tc.parsed, p)
			assert.Equal(t, tc.err, err)
		})
	}
}
