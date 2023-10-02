# AIC-OCR-search

## Run 

Run on port `8080`:
```go
go run main.go
```

## API

### Health
`/health`

### Search:
`/search/{method}/{query}/{limit}`

#### Request:
- `method`: [`exact`, `fuzzy (not implemented)`, `advanced (not implemented)`]
- `query`: `unicode string`, the query to search for
- `limit`: return the top `limit` results

#### Response:
- `status`: `int`, HTTP code
- `message`: `string`, Error message if error
- `data`: 
    - `filename`: `string`, the filename of the video
    - `frameID (keyframe number)`: `string (xxxx)`, the frame number of the keyframe
    - `score`: `float`, the score of the keyframe

#### Example: 
Search for `Tổng` with `exact` method and return `2` results: `http://localhost:8080/search/exact/T%E1%BB%95ng/2`