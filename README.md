# GINting Boilerplate
[GIN](https://github.com/gin-gonic/gin) Boilerplate
untuk kamu yang dikejar waktu

## Run Command
```sh scripts/start-dev``` with [air hot reload](https://github.com/cosmtrek/air)
```go run ext/sample-grpc-server/main.go``` run sample grpc server

#### Development Work Flow with Dependecty Injection Google Wire
- ```go get github.com/google/wire/cmd/wire```
- ```$GOPATH/bin/wire pkg/wire/wire.go```

#### Enable Keycloak
- edit .env file and set:
- KEYCLOAK=1
- KEYCLOAK_CERTS={{keycloak_host}}/auth/realms/{{realms_name}}/protocol/openid-connect/certs
- KEYCLOAK_ISSUER={{keycloak_host}}/auth/realms/{{realms_name}}

### RUN Sample GRPC SERVER
- ```go run ext/sample-grpc-server/main.go```

## TODO
- sample GIN Routing Group :white_check_mark:
- sample ORM database with [GORM](https://gorm.io/) :white_check_mark:
- sample Dependecy Injection with [Google Wire](https://github.com/google/wire) :white_check_mark:
- sample [Opentelemetry](https://opentelemetry.io/) with exporter [Jaeger](https://www.jaegertracing.io/) :white_check_mark:
- sample [Keycloak](https://www.keycloak.org/) Middleware :white_check_mark:
- sample [GRPC](https://grpc.io/) Service Middleware :white_check_mark:
- sample logging with [Logrus](https://github.com/sirupsen/logrus) :white_check_mark:
- sample email notification
- sample Telegram Notification 
- sample redis cache
- sample google pub/subs
- sample google cloud tasks
- etc..
