A very simple library for logging in golang
No ensurance for performance, it is just easy to use.

## Easy to go:

```go
slog.Error("error %s", "err")
slog.Error("error")
```

## Setup logging level
```go
slog.SetDefaultLevel(slog.TRACE)
slog.Error("error %s", "err")
slog.Error("error")
slog.Trace("error")
```

## Using a logging instance and log into a file

```go
log := slog.FileLogger(slog.TRACE, "out.log")
log.Error("error %s", "err")
log.Error("error")
```

