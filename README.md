# HTML Parser

[![Go Report Card](https://goreportcard.com/badge/github.com/dereckdamphouse/html-parser)](https://goreportcard.com/report/github.com/dereckdamphouse/html-parser)

This project is a tool for providing HTML parsing via a RESTful API endpoint.

## Prerequisites
**CLI Tools**
- [Go](https://golang.org/doc/install) - v1.19.2
- [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html) - v2.7.6
- [AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html) - v1.51

## Example
**API Request**
```javascript
const url = 'https://.../v1/html-parser';

const body = {
  html: '<html>...</html>',
  properties: [
    {
      name: 'title',
      selector: '.product .title',
    },
    {
      name: 'details',
      selector: '.product .details ul li',
    },
    {
      name: 'image',
      selector: '.product img',
      attribute: 'src',
    },
  ],
};

const options = {
  method: 'POST',
  body: JSON.stringify(body),
};

return fetch(url, options)
  .then((response) => response.json())
  .catch((e) => console.log(e));
```
**Successful Response**
```yaml
{
  "title": [
    "Men's Columbia Flattop Ridge Fleece Jacket"
  ],
  "image": [
    "https://....com/images/clothing/2599191_ALT-1000.jpg"
  ],
  "details": [
    "Polyester fleece",
    "Machine wash",
    "Imported",
  ],
}
```
**Failure Response**
```yaml
{
  "error": "no properties found"
}
```
