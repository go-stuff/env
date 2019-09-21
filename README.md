# env

Load environment variables from a <Key, Value> pair file. The purpose of this is to load your environment for testing locally, sometimes you are stuck in a situation where that is all you have to test, you are off the network and running all your services and databases locally.

## Example <Key, Value> Pair File

```txt
VALUE_ONE = "One"
VALUE_TWO = "Two"
ENV_URI = "https://github.com/go-stuff/env"
```

## Example Usage

It is important to know that you should never commit your environment files into any git repository, trust me, it will cause you lots of security headaches.

If you do not already have a `.gitignore` file create one. Then make sure you have the environment file you are using for testing in `.gitignore`.

In my example the environment file is `.env` this is an example contents of `.gitignore`, maybe you have one file, or maybe you setup a flag so you can switch between files for different environments:

```git
.env
.env-dev
.env-qa
.env-prod
```

`env` uses the package `https://github.com/uber-go/zap` for logging purposes. You can pass it the path to your environment file and a `*zap.logger`.

This is how you would load the contents of the `.env` file into the environment:

```go
env.File(".env", logger)
```

Example output:

```bash
github.com/go-stuff/grpc$ go run main.go
{"level":"info","ts":1569075534.3478606,"caller":"grpc/main.go:106","msg":"environment file was loaded","path":".env"}
```

*Note: it the path does not exist it will only `warn` and just load environment variables, but if there is an error with the file or parsing the file it will call `fatal`*

```bash
{"level":"warn","ts":1569075616.2841346,"caller":"grpc/main.go:39","msg":"path does not exist, using environment variables","path":""}

{"level":"fatal","ts":1569075459.8256843,"caller":"grpc/main.go:47","msg":"path is a directory not a file","path":"./","stacktrace":"..."}
```