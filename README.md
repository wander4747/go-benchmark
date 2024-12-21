## How use
Build image
```sh 
 docker build --no-cache -t benchmark-tool .
```

Run comand:
```sh
docker run --rm benchmark-tool -url="https://test-api.k6.io/public/crocodiles/?format=api" -requests=100 -duration="1m" -method=GET
```

## Command line arguments
This document describes the command-line flags used in the benchmark tool and their corresponding functionality.

| Flag        | Type    | Default Value                                           | Description                                                                                       |
|-------------|---------|---------------------------------------------------------|---------------------------------------------------------------------------------------------------|
| `--url`     | `string`| `"https://test-api.k6.io/public/crocodiles/?format=api"`| The URL of the API to be used for the benchmark. This is the target endpoint for the GET or POST requests made during the benchmark test. |
| `--method`  | `string`| `"GET"`                                                 | The HTTP method used for the requests. This flag accepts either `GET` or `POST`. It defines the type of HTTP request (GET for fetching data or POST for sending data). |
| `--requests`| `int`   | `0`                                                     | Specifies the total number of requests to be made during the benchmark. If set to `0`, the tool will perform an infinite number of requests, running until the specified duration is reached. |
| `--duration`| `string`| `"30s"`                                                 | Defines the duration of the benchmark. This should be a string representing time, such as `30s` (seconds), `1m` (minute), or `2h` (hours). The benchmark will run for the given duration and stop afterward. |
| `--payload` | `string`| `{"key":"value"}`                   