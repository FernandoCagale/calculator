# calculator

[![Build Status](https://travis-ci.org/FernandoCagale/calculator.svg?branch=master)](https://travis-ci.org/FernandoCagale/calculator)

### Docker

`running docker multi-stage builds and publish calculator to gRPC and HTTP`

```sh
$   ./scripts/publish-grpc.sh
```

```sh
$   ./scripts/publish-http.sh
```

### Running docker-compose

```sh
$   docker-compose up -d
```

*   `gRPC` testing

```sh
$   grpcurl -plaintext -d '{"user_id":"001", "product_id":"001"}' localhost:5000 grpc.Calculator.Calculator
```

*   `HTTP` testing

```sh
$   curl --location --request GET 'localhost:8080/calculator/user/001/product/003'
```

### Running local

```sh
$   docker-compose up mongo
```

#### Standard Go Project [Layout](https://github.com/golang-standards/project-layout)

```sh
$   go mod download
```

```sh
$   go mod vendor
```

`download "dependency injection"` [wire](https://github.com/google/wire)

```sh
$   go get -u github.com/google/wire/cmd/wire
```

```sh
$   ./scripts/start-grpc.sh
```

```sh
$   ./scripts/start-http.sh
```
*   `gRPC` testing

```sh
$   grpcurl -plaintext -d '{"user_id":"001", "product_id":"001"}' localhost:5000 grpc.Calculator.Calculator
```

*   `HTTP` testing

```sh
$   curl --location --request GET 'localhost:8080/calculator/user/001/product/003'
```