# Usage:
# make					# compile all binaries
# make clean			# remove ALL binaries

run:
	go run main.go -count=1

.PHONY: dependencies build test cover clean run