BinariesDirectory = ./bin
MainDirectory = ./cmd/main

linux-mkdir:
	mkdir -p "$(BinariesDirectory)"

windows-mkdir:
	if not exist "$(BinariesDirectory)" mkdir "$(BinariesDirectory)"

linux-build:
	make linux-mkdir
	make basic-build

basic-build:
	go build -o "$(BinariesDirectory)" "$(MainDirectory)/main.go"

build:
	make windows-mkdir
	make basic-build

linux-run:
	make linux-build
	make basic-build

run:
	make build
	cd "$(BinariesDirectory)" && main

run-os:
	go run cmd/osfilepathtesting/ostesting.go