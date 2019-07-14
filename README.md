# ParkingLot

Help you manage your pariking a lot easier, the program is written by Golang.

## Requirements

- Install Golang and make sure you have GOPATH variable in PATH environment.
- Build file will be created under `$GOPATH/bin`, so make sure you have this bin folder in your PATH environment too.

## Install

- Place the solution in parking_lot folder under `$GOPATH/src`.
- At parking_lot folder, run `bin/setup` to install dependencies.

## Run

- At parking_lot folder, run `bin/parking_lot` to enter to interactive mode.
- At parking_lot folder, run `bin/parking_lot file_input.txt` to run with input file.
Please make sure to specify correct path of the input file, you can either use absolute
path or relative path to folder you currently stand on.
For example: you are standing at `parking_lot` and `file_input.txt` is under `parking_lot` too.
Then you can run `bin/parking_lot file_input.txt`.

## Test

Go offers excellent built-in `testing` package, to run the tests.
- `cd test`
- Run `go test ./...` to run all tests.
- To run specific test, just go into a specific folder (`test/parking` or `test/utils`) then run `go test`.
