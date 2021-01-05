# echo-service
## REST API Contract

echo-service provides the following endpoints

`/health`
- HTTP Method: GET
- Description: Returns a JSON object comprised of a timestamp of when the request was processed, the health status, and the http status returned by the endpoint

Response body schema:
```
{
    "TimeStamp" : string,
    "HealthStatus" : string,
    "HttpStatus": string 
}
```

`/echo`
- HTTP Method: POST
- Description: Returns a JSON object comprised of a timestamp of when the request was processed, and the string POSTed to the endpoint
  
Request body schema:
```
{
    "EchoStr" : string
}
```
 
Response body schema:
```
{
    "TimeStamp" : string,
    "EchoStr" : string
}
```



## Config

echo-service fetches it's config through the environment. The following environment variables are expected (or supported):
- SERVICE_PORT
  - optional: false
  - description: The port the REST API runs on.
- SSL_CERT_PATH
  - optional: true
  - description: The full path of the SSL cert, must be set alongside SSL_KEY_PATH to enable HTTPS.
- SSL_KEY_PATH
  - optional: true
  - description: The full path of the SSL key, must be set alongside SSL_CERT_PATH to enable HTTPS.

## Makefile Usage

`make clean`

Deletes ./bin folder and associated artifacts.

`make test`

Runs `go test` on all packages in echo-service.

`make run`

Runs the echo-service locally. Ensure SERVICE_PORT is exported into your local environment.

`make build`

Compiles the echo-service binary, storing it in ./bin/echo-service.

`make docker-build`

Builds the docker container for echo-service, tags with the semver found in the VERSION file, and enables HTTPS. If compiling on linux, will reuse any previously recompiled binaries.

`make docker-run`

Runs the docker container built in the above command, and maps the service to port 8080 on the host by default. The port on the host can be overriden by setting the `local-port ` variable when calling the make target.



## Local Testing

Test it with curl!
```
$> make docker-build docker-run
$> curl -k https://localhost:8080/health # use -k because we've issued the SSL cert ourselves!
$> curl -d '{"EchoStr":"hello world"}' -H 'Content-Type: application/json' -k  https://localhost:8080/echo
```


## TODO

- Refactoring: Why do we have three seperate loggers? Can also provide Errf, Infof, Fatalf, etc. functions for ease of use
- Fowarding from HTTP to HTTPS should be implemented
- Basic auth