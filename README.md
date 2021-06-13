
![Go](https://github.com/ayushjha-fm/tmai.server.mock/workflows/Go/badge.svg?branch=master)
![Publish](https://github.com/ayushjha-fm/tmai.server.mock/workflows/Publish%20Docker%20image/badge.svg)

## Running the mock server

```bash
docker run -p 3000:3000 ayushjhafm/tmai.server.mock
```

### How does it work

The server serves JSON files in the `tmai` directory, as per the rules
described in `tmai/api.json`. This server is slightly more fancy than a simple JSON server.
The features for this mock server are described below.

#### Features
* Recieve and transmit the "query" from `request->response` for all the `message` type responses.
* Manage state of suggestions based on availability of `CONVERSATION_TOKEN`
* Read and send JSON responses by reading from JSON files with same name as the endpoint path.
  For example, the `/text` endpoint will read from `tmai/text.json` file.

### Structure of api.json

`api.json`

```javascript
{
    "endpoints": [{
            "path": "/text",
            "status": 200,
            "type": "message",
            "methods": [
                "GET",
                "POST"
            ]
    }],
    "port": 3000
}
```

#### Structure of Endpoint

| Field Name | Definition                                               |
|------------|----------------------------------------------------------|
| path       | The path for which this endpoint will serve a file       |
| status     | The status returned in response                          |
| type       | There are two types of responses, "message" and "trivia" |
| methods    | Which HTTP Methods should the endpoints support          |

## License

MIT ✨
