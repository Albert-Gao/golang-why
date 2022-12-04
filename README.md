# golang Why?

A simple algorithm:

a function that takes an array of int, mutate each item in the array by 1.

I have 2 versions, a concurrent version and a non-current version.

## Interesting result:

With a `time.Sleep()`, the concurrent version is way faster than the non-current version.

Without the `time.Sleep()`, the concurrent version is way slower than the non-current version.

What am I missing here?

## How to re-produce?

Try commenting out the `time.Sleep(ms(1))` in `shared.go` and executing the following script:

`go test -bench=. -benchtime=3s`

## Some other findings

Update the channel to a buffered version (`resultChannel := make(chan result, len(values))`) does not really help closing the gap, helps a lot in the no-sleep version by 40% faster, but still way slower than the normal version in the no-sleep case.
