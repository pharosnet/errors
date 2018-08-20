# errors
Stack errors for golang

Package errors provides simple error handling primitives.

`go get github.com/pharosnet/errors`

The traditional error handling idiom in Go is roughly akin to
```go
if err != nil {
        return err
}
```
which applied recursively up the call stack results in error reports without context or debugging information. The errors package allows programmers to add context to the failure path in their code in a way that does not destroy the original value of the error.

## Adding context to an error

The errors.Wrap function returns a new error that adds context to the original error. For example
```go
_, err := ioutil.ReadAll(r)
if err != nil {
        return errors.Wrap(err)
}
```
## Retrieving the cause of an error

Using `errors.Wrap` constructs a stack of errors, adding context to the preceding error. Depending on the nature of the error it may be necessary to reverse the operation of errors.Wrap to retrieve the original error for inspection. Any error value which implements this interface can be inspected by `errors.Cause`.
```go
type Errors interface {
	Error() string
	Cause() error
	OccurTime() time.Time
	PCS() []uintptr
	Contains(error) bool
	Format(fmt.State, rune)
}
```
`errors.Contains` will search target in error which implements `Errors`, For example:
```go
    e1 := io.EOF
    e2 := errors.With(e1, "error2")
    e3 := errors.WithF(e2, "%s", "error3")
     
    if errors.Contains(e3, e2) {
            // TODO .. 
    }
        
    if errors.Contains(e3, e1) {
        // TODO ...
    }
```

[Read the examples for more usages.](https://github.com/pharosnet/errors/example) 

[Read the package documentation for more information](https://godoc.org/github.com/pharosnet/errors).

## Contributing

We welcome pull requests, bug fixes and issue reports. With that said, the bar for adding new symbols to this package is intentionally set high.

Before proposing a change, please discuss your change by raising an issue.

## License

GNU GENERAL PUBLIC LICENSE(v3)