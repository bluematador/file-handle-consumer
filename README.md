
BUILD
----------
docker run -v $(pwd):/app --workdir /app quay.io/deis/go-dev:latest go build -o consume-fh main.go


RUN
----------

Included binary is linux/amd64

sudo chown ${USER}:${USER} consume-fh
./consume-fh
