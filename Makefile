default: clean build

build:
	mkdir ./dist
	go build ./main.go
	mv ./main ./dist/isszp
	cp -rf ./template ./static ./config ./dist

clean:
	rm -rf ./dist
