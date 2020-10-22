# batchutil

`batchutil` is a package for Go that provides functions to describe a concurrent batch processing easily.
If you just set configurations and a function, you can run a concurrent program.
It's very useful when you write a big batch processing like data migrations.

## Install

Install using `go get github.com/tkm-kj/batchutil`.

## Usage

### Define configuration

```go
cfg := batchutil.Config{
    ConcurrentLimit: 100,  // limit value of concurrency(It's assuming database max connection)
    StartNumber:     1,    // loop start number
    EndNumber:       1450, // loop end number
    BatchSize:       100,  // the size of batch(StartNumber is incremented by BatchSize until it exceeds EndNumber)
}
```

When you set `1` to `ConcurrentLimit` or nothing, batch processing runs serially.

### Run batch processing

```go
util := batchutil.Open(&cfg)
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
