
BUILD
----------

```
docker run -v $(pwd):/app --workdir /app quay.io/deis/go-dev:latest go build -o consume-fh main.go
```

RUN
----------

Included binary is linux/amd64

```
curl -LO https://github.com/bluematador/file-handle-consumer/raw/master/consume-fh
sudo chmod +x consume-fh
sudo chown ${USER}:${USER} consume-fh
./consume-fh
```