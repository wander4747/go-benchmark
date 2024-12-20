# Wait. I'm still developing

## How use
Build image
```sh 
 docker build --no-cache -t benchmark-tool .
```

Run comand:
```sh
docker run --rm benchmark-tool -url="https://test-api.k6.io/public/crocodiles/?format=api" -requests=100 -duration="1m" -method=GET
```