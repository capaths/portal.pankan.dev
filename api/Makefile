main_package_path = .
binary_name = main


.PHONY: tidy
tidy:
	go mod tidy -v
	go fmt ./...

.PHONY: build
build:
	go build -o=./out/${binary_name} ${main_package_path}

.PHONY: run
run: build
	./out/${binary_name}
