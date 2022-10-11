package html

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dereckdamphouse/html-parser/pkg/req"
)

type Parsed map[string][]string

var htmlReader = goquery.NewDocumentFromReader

func Parse(d *req.Data) (Parsed, error) {
	parsed := make(Parsed)

	doc, err := htmlReader(strings.NewReader(d.HTML))
	if err != nil {
		return parsed, err
	}

	for _, prop := range d.Properties {
		name := prop.Name
		selector := prop.Selector
		attribute := prop.Attribute

		if name == "" || selector == "" {
			continue
		}

		doc.Find(selector).Each(func(_ int, s *goquery.Selection) {
			if attribute != "" {
				att, ok := s.Attr(attribute)
				if att != "" && ok {
					parsed[name] = append(parsed[name], att)
				}
			} else {
				text := s.Text()
				if text != "" {
					parsed[name] = append(parsed[name], text)
				}
			}
		})
	}

	return parsed, nil
}
