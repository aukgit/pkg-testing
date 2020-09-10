BuildPath = "./bin"

build:
	echo "building code"
	mkdir -p $(BuildPath)
	go build -o $(BuildPath) main.go

run:
	echo "running code"
	make build
	cd $(BuildPath) && clear && ./main

clean:
	rm -rf $(BuildPath)

labels:
	echo "build"
	echo "run"
	echo "clean"
