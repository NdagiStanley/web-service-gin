# Restful API with Go and Gin

REF:
<https://go.dev/doc/tutorial/web-service-gin>

MORE:
- <https://go.dev/doc/effective_go>
- <https://go.dev/doc/code>
- <https://go.dev/tour/welcome/1>
- <https://pkg.go.dev/github.com/gin-gonic/gin>
- <https://gin-gonic.com/docs/>

```sh
go get .
```

```sh
go run .
```

```sh
curl http://localhost:8080/albums
```

```sh
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```

```sh
curl http://localhost:8080/albums/2
```
