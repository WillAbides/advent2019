.PHONY gobuildcache:

bin/intcomputer: gobuildcache
	go build -o bin/intcomputer ./cmd/intcomputer
