# Usage:
# make					# compile all binaries
# make clean			# remove ALL binaries

run:
	go run main.go

.PHONY: dependencies build test cover clean run