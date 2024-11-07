# Restful API with Go and Gin

REF:
<https://go.dev/doc/tutorial/web-service-gin>

MORE:

- [Effective GO][effective_go]
- [How to Write Go Code][go_code]
- [A Tour of Go][tour]
- [Gin Web Framework][gin_framework]
- [Gin Web Framework documentation][gin_framework_docs]
- [go-sqlite3 - sqlite3 driver for go using database/sql][go_sqlite3]

## Set up

```sh
go get .
```

## SQLITE

1. Create a postgres server
2. Connect to the postgres server
3. Run the following SQL to create the table:

    ```sql
    CREATE TABLE albums (
        id INTEGER PRIMARY KEY,
        title VARCHAR(255),
        artist VARCHAR(64),
        created DATE DEFAULT (CURRENT_DATE),
        price DECIMAL(10, 2)
    );
    ```

For steps 1, locally you can use Docker or Postgres app. For step 2 and 3, you can use the CLI, or alternatively, a Database GUI tool like [TablePlus][tableplus] for creating the table, and inserting and viewing data.

## Run application

```sh
export DATABASE_URL="postgres://postgres:postgres@localhost:5432/test"
```

```sh
go run .
```

```sh
curl http://localhost:8080/albums
```

```sh
curl http://localhost:8080/albums/2
```

```sh
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```

## Next

1. Deploy it

[tableplus]: https://tableplus.com/
[effective_go]: https://go.dev/doc/effective_go
[go_code]: https://go.dev/doc/code
[tour]: https://go.dev/tour/welcome/1
[gin_framework]: https://pkg.go.dev/github.com/gin-gonic/gin
[gin_framework_docs]: https://gin-gonic.com/docs/
[go_sqlite3]: https://github.com/mattn/go-sqlite3
