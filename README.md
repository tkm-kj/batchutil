# batchutil

![Go](https://github.com/tkm-kj/batchutil/workflows/Go/badge.svg)

`batchutil` is a package for Go that provides functions to describe a concurrent batch processing easily.
If you just set configurations and a function, you can run a concurrent program.
It's very useful when you write a big batch processing like data migrations.

## Install

Install using `go get github.com/tkm-kj/batchutil`.

## Usage

### Define configuration

```go
cfg := batchutil.NewConfig(
    3,  // [concurrentLimit] limit value of concurrency(It's assuming database max connection)
    1,    // [startNumber] loop start number
    1450, // [endNumber] loop end number
    100,  // [batchSize] the size of batch(StartNumber is incremented by BatchSize until it exceeds EndNumber)
)
```

When you set `1` to `ConcurrentLimit`, batch processing runs serially.

### Run batch processing

```go
util := batchutil.NewUtil(cfg)
f := func(min, max int64) error {
    log.Printf("min: %d, max: %d", min, max)
    // I assume that you'll find list by `min` and `max` and do something
    return nil
}
err := util.Run(f)
```

The result is below.

```
2020/10/22 18:33:36 min: 1401, max: 1501
2020/10/22 18:33:36 min: 301, max: 401
2020/10/22 18:33:36 min: 1, max: 101
2020/10/22 18:33:36 min: 101, max: 201
2020/10/22 18:33:36 min: 501, max: 601
2020/10/22 18:33:36 min: 1301, max: 1401
2020/10/22 18:33:36 min: 201, max: 301
2020/10/22 18:33:36 min: 901, max: 1001
2020/10/22 18:33:36 min: 701, max: 801
2020/10/22 18:33:36 min: 801, max: 901
2020/10/22 18:33:36 min: 601, max: 701
2020/10/22 18:33:36 min: 1101, max: 1201
2020/10/22 18:33:36 min: 1001, max: 1101
2020/10/22 18:33:36 min: 1201, max: 1301
2020/10/22 18:33:36 min: 401, max: 501
```

Batch processing runs concurrently!

### Run batch processing with context

Also, you can use batchutil with `context.Context` too.
In other words, you can just skip batch process when an error occurs.

```go
util := batchutil.NewUtil(cfg)
f := func(ctx context.Context, min, max int64) error {
    select {
    case <-ctx.Done():
        return ctx.Err()
    default:
    }

    if min == 101 {
        return errors.New("error")
    }
    log.Printf("min: %d, max: %d", min, max)
    return nil
}
err := util.RunWithContext(context.Background(), f)
```

The result is below.

```
2020/12/29 00:07:33 min: 201, max: 301
2020/12/29 00:07:33 min: 1, max: 101
```
