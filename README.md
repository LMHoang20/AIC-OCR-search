# AIC-OCR-search

## Go version
```
go version go1.21.1 linux/amd64
```

## Run 

### Run with Go:
```go
go run main.go
```

### Docker:
```bash
sudo docker build -t ocr-search .  
sudo docker run -dp 0.0.0.0:8080:8080 ocr-search
```

## API

### Health
`/health` (GET)

### Search:
`/search/{method}` (POST)

#### Request:

##### Params:
- `method`: [`exact`, `fuzzy`, `advanced (not implemented)`]

##### Body:
- `query_string`: `unicode string`, the query to search for
- `topk`: return the top `k` results

#### Response:
- `status`: `int`, HTTP code
- `message`: `string`, Error message if error
- `data`: 
    - `video`: `string`, the name of the video
    - `frame_name`: `string`, the frame image name
    - `score`: `int`, the score of the keyframe

#### Example: 
Search for `vietnam` with `fuzzy` method and return `4` results: 

Request: `http://localhost:8080/search/fuzzy`

Body: 
```json
{
    "query_text": "vietnam",
    "topk": 4
}
```
Response: 
```json
{
    "data":
    [
        {"video":"L23_V018","frame_name":"099.jpg","score":7},
        {"video":"L12_V014","frame_name":"0314.jpg","score":7},
        {"video":"L06_V022","frame_name":"0074.jpg","score":7},
        {"video":"L22_V027","frame_name":"049.jpg","score":7}
    ],
    "message":"OK",
    "status":200
}
```