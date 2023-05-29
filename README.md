# simple-atm-controller
Implement a simple ATM controller
- No hardware control is involved
- Doesn't work with a real bank
- We're thinking about adding hardware control in the future
- We are considering integrating with a real bank in the future


## Test

```shell
$ make test
```
- Run all Go tests under the project.

## Build

```shell
$ make build
```

- First, run the tests and build the executable.
- If no changes have been made since the last build, we don't rebuild.
- If you want to build without testing, use the command below.

```shell
$ make build-force
```

## Run

```shell
$ make run
```

- Test and build, then run it.
- If you want to build and run without testing, use the command below.

```shell
$ make run-force
```

- If you want to run it without rebuilding, run the command below.

```shell
$ make run-now
```

-  If you don't have an executable file, it won't run properly. Use the commands above.

## Caution!

The test code fails because the current implementation of the test code is incomplete. Please run it with the force method.
