service ?= mars-rover`

build: 
	@go build -o "bin/${service}" "./cmd/${service}"

run: build
	@./bin/${service}

