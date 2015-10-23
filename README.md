# discovery-go

Programmatically find services.

This library currently is just an abstraction around reading environment variables used for dependent services.



## API

[Godoc Documentation](http://godoc.org/github.com/Clever/discovery-go)

### import

```go
import "gopkg.in/Clever/discovery-go.v1"
```

### Examples

```go
gearmanAdminURLString, err := discovery.URL("gearman-admin", "http")
if err != nil {
    log.Fatal("ERROR: " + err.Error())
}


redisHOST, err := discovery.Host("redis", "tcp")
if err != nil {
    logger.Fatal("ERROR: " + err.Error())
}
redisPORT, err := discovery.Port("redis", "tcp")
if err != nil {
    logger.Fatal("ERROR: " + err.Error())
}
```

### Environment Variables

This library looks up environment variables(eventually maybe not). For it to work, your environment variables need to adhere to the following convention:

SERVICE_\<SERVICE_NAME\>_\<PROTOCOL\>_\<PROTO|HOST|PORT\>

Here is an example using redis:
```bash
SERVICE_REDIS_TCP_PROTO = "tcp"
SERVICE_REDIS_TCP_HOST = "localhost"
SERVICE_REDIS_TCP_PORT = "6379"
```

