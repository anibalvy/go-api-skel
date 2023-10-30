[![Create and publish Docker Image](https://github.com/anibalvy/go-api-skel/actions/workflows/build_image.yaml/badge.svg)](https://github.com/anibalvy/go-api-skel/actions/workflows/build_image.yaml)
# GO API

skeleton used as a starting template for apis, a small subset of what I used on production.

## running

Use [Air](https://github.com/cosmtrek/air)
or
```sh
go run main.go
```

## Invokations

Get token:
```sh
curl --location --request POST http://127.0.0.1:3000/get-token --header 'Content-Type: application/json'  --data-raw '{"username": "user", "password": "pass1234"}'
```

Call authenticated user api:
```sh
curl --location --request GET http://127.0.0.1:3000/v1/users --header 'Content-Type: application/json' --header 'Authorization: bearer eyJhbGciOaioommsdk234nR5cCI6IkpXVCJ9.eyJleHdasfjkadsjxxc3MDAsIm5hbWUiOiJrYW5pYmFsdiadsfajkxI6MX0.wI1_LQ3Rafsdasdfasdfjjjl9B9JtiBXxjHsH2SNJ3s'
```

