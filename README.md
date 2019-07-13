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

This is how you would load the contents of the `.env` file into the environment:

```go
err := env.File(".env")
if err != nil {
    log.Fatalf("failed to parse: %v", err)
}
```

Example output:

```bash
github.com/go-stuff/grpc$ go run main.go
2019/07/12 20:33:39 INFO > env.go > File(): .env loaded
```
