# Book Service
#### _test project for Fidibo_


## Service

- Simple http server for search books by title
- Fake data generator.(see docker-compose.yaml)

You can find project description in this [link](https://github.com/alichz2001/fidibo/blob/master/task.pdf).


## Project requirements
- [x] Book service:
  - [x] Search books by its title.
  - [x] Cache search results.
  - [x] Retrieve Search from cache if exists.
  - [x] Mysql DB.
  - [x] Redis.

## Used technologies
- **Fiber**: used as webserver and gateway. [link](https://docs.gofiber.io/)
- **Mysql**: used as primary database.
- **Redis**: used as cache server.
- **Docker & Docker-compose**: for containerization and dev and test mode run.
- **Makefile**: for manage commands.


## Notes
- In this project, I've employed a layered architecture with three logical layers: Model, Service, and Controller. The simplicity of the current task doesn't necessitate additional abstractions or complex packaging. However, for future development, some refactoring may be beneficial. This structure is well-suited for testing purposes.
- Why 'Fiber'? So while both Gin and Echo are excellent frameworks with their respective advantages and drawbacks, my preference for a web server is Fiber. see this [link](https://medium.com/deno-the-complete-reference/go-gin-vs-fiber-hello-world-performance-6863e597b654)


## RUN

```shell
make run
```
-----------------
## API Documentation

**Note:** You can find the full API Postman collection in this [link](https://github.com/alichz2001/fidibo/blob/master/fidibo.postman_collection.json).

#### 1\. Search Book

* POST http://[server]/v1/books/search?q=[search]
  > List all matched books.


* Sample logs of webserver:
  ```shell
    {"time":"2024-01-08T11:59:52.670071132Z","level":"INFO","msg":"try fetch books from cache, key: 'tit'"}
    {"time":"2024-01-08T11:59:52.670198612Z","level":"INFO","msg":"try fetch books from db, key: 'tit'"}
    {"time":"2024-01-08T11:59:52.670936363Z","level":"INFO","msg":"db hit, key: 'tit'"}
    {"time":"2024-01-08T11:59:52.67094211Z","level":"INFO","msg":"try put books to cache, key: 'tit'"}
    11:59:52 | 200 |    1.341938ms | 172.24.0.1 | GET | /v1/books/search | -
    {"time":"2024-01-08T11:59:53.536264743Z","level":"INFO","msg":"try fetch books from cache, key: 'tit'"}
    {"time":"2024-01-08T11:59:53.536603262Z","level":"INFO","msg":"cache hit, key: 'tit'"}
    11:59:53 | 200 |     386.519µs | 172.24.0.1 | GET | /v1/books/search | -
    {"time":"2024-01-08T12:00:07.197378678Z","level":"INFO","msg":"try fetch books from cache, key: 'tit'"}
    {"time":"2024-01-08T12:00:07.19749869Z","level":"INFO","msg":"try fetch books from db, key: 'tit'"}
    {"time":"2024-01-08T12:00:07.198028935Z","level":"INFO","msg":"db hit, key: 'tit'"}
    {"time":"2024-01-08T12:00:07.198036304Z","level":"INFO","msg":"try put books to cache, key: 'tit'"}
    12:00:07 | 200 |     839.388µs | 172.24.0.1 | GET | /v1/books/search | -
    ```


## Personal TODO
- [ ] add validations.
- [ ] add pagination! now all matched data will return in result.
- [ ] make better response messages for web client.
- [ ] rate-limiting and heath-checking.
- [ ] make config more useful. Add reloading config options by signal, file monitor or even REST API.
- [ ] make better error handling. create const errors and return error messages in fiber ErrorHandler.
- [ ] add interfaces for more usable and maintainable codes.
- [ ] add graceful shutdown scenario.
- [ ] add other query params.
