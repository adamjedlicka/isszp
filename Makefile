default: clean build doc

build:
	mkdir -p ./dist
	go build ./main.go
	mv ./main ./dist/isszp
	cp -rf ./template ./static ./config ./dist

clean:
	rm -rf ./dist

doc:
	mkdir -p ./dist
	go build golang.org/x/tools/cmd/godoc
	mv ./godoc ./dist/godoc