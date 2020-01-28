# go-wasabi
[![CircleCI](https://circleci.com/gh/inloco/go-wasabi/tree/master.svg?style=svg)](https://circleci.com/gh/inloco/go-wasabi/tree/master) [![GoDoc](https://godoc.org/github.com/inloco/go-wasabi?status.svg)](https://godoc.org/github.com/inloco/go-wasabi)

Intuit Wasabi API Client written in Go

### Instantiating a client

```go

client := NewHttpClient(
  "https://your-wasabi-host",
  "application_name",
  "username",
  "password",
)

```

### API doc

http://wasabi-for-apps.ubee.in/swagger/swaggerui/#/

### Generate an assignment

```go
assignment, err := client.GenerateAssignment(
  context.Background(),
  "ExperimentLabel",
  "userID",
)
```

Check for `assignment.Status`, and use the `assignment.Payload` on your application :)

### Creating an experiment

See the file `examples/experiment.go`
