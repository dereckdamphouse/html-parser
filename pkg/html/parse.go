package html

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dereckdamphouse/html-parser/pkg/req"
)

var htmlReader = goquery.NewDocumentFromReader

func Parse(d *req.Data) (map[string][]string, error) {
	res := make(map[string][]string)

	doc, err := htmlReader(strings.NewReader(d.HTML))
	if err != nil {
		return res, err
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
					res[name] = append(res[name], att)
				}
			} else {
				text := s.Text()
				if text != "" {
					res[name] = append(res[name], text)
				}
			}
		})
	}

	return res, nil
}
