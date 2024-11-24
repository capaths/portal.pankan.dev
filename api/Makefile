main_package_path = .
binary_name = main
jet_executable_win = C:/Users/pcamp/go/bin/jet.exe


.PHONY: test
test:
	go test -v -buildvcs ./...

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

.PHONY: generate_db
generate_db:
	${jet_executable_win} -dsn=postgres://postgres:postgres@localhost:5432/pankan_db?sslmode=disable -schema=portal -path=./.gen