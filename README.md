# golang Why?

A simple algorithm:

a function that takes an array of int, mutate each item in the array by 1.

I have 2 versions, a concurrent version and a non-current version.

## Interesting reesult:

With a `time.Sleep()`, the concurrent version is way faster than the non-current version.

Without the `time.Sleep()`, the concurrent version is way slower than the non-current version.

What am I missing here?
