package psr

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dereckdamphouse/html-parser/pkg/req"
)

func Parse(b *req.Body) (map[string][]string, error) {
	res := make(map[string][]string)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(b.HTML))
	if err != nil {
		return res, err
	}

	for _, sel := range b.Selectors {
		key := sel.Key
		selector := sel.Selector
		attribute := sel.Attribute

		if key == "" || selector == "" {
			continue
		}

		doc.Find(selector).Each(func(_ int, s *goquery.Selection) {
			if attribute != "" {
				att, ok := s.Attr(attribute)
				if att != "" && ok {
					res[key] = append(res[key], att)
				}
			} else {
				res[key] = append(res[key], s.Text())
			}
		})
	}

	return res, nil
}
