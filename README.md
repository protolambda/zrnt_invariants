# ZRNT invariants

Invariants for ETH 2.0 fuzzing. Build on [ZRNT](https://github.com/protolambda/zrnt), the Go executable spec.


## Getting started

`go get` to install go modules. The repo should be outside of the go-path, to use the new GO-module support.

Build/run/test with `-tags preset_minimal` or `-tags preset_mainnet`


## Usage

`genesis.go`: to build a valid simulated genesis state with

`state.go`: state invariants, to be extended

`block.go`: pre-processing, to give a pure random block a chance to become part of the corpus of valid blocks 

`usage_test.go`: test run that applies a bunch of random blocks to the genesis state.
 Most (if not all) are invalid, even after pre-processing. Either improve pre-processing, or drop the idea,
 and mutate valid blocks (hard too, one block is nothing like the other, due to crypto (hashes!))

`util.go`, `util_test.go`: as a compromise, make roots in random blocks valid,
 but only some of the time, to cover the invalid case too. These contain utils + testing for that.

`genesis_test.go`: check if the simulated genesis state passes the invariants. 


