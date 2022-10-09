# HTML Parser
This project is a tool for exposing HTML parsing via a REST API endpoint.

## Prerequisites

**CLI Tools**
- [Go](https://golang.org/doc/install) - v1.19.2
- [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html) - v2.7.6
- [AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html) - v1.51

## JavaScript Example

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
