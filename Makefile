##
# go-nepali
#
# @file
# @version 0.1

test:
	go test -v ./...

coverage:
	go test -v ./... -cover

coverage-html:
	go test -v ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm -rf coverage.out

# end
