default: clean build doc

build:
	mkdir -p ./dist
	GOOS=linux go build -o ./dist/isszp_linux ./main.go
	GOOS=windows go build -o ./dist/isszp_windows.exe ./main.go
	GOOS=darwin go build -o ./dist/isszp_macos ./main.go
	cp -rf ./template ./static ./config ./dist

clean:
	rm -rf ./dist

doc:
	mkdir -p ./dist
	GOOS=linux go build -o ./dist/godoc_darwin golang.org/x/tools/cmd/godoc
	GOOS=windows go build -o ./dist/godoc_windows golang.org/x/tools/cmd/godoc
	GOOS=darwin go build -o ./dist/godoc_macos golang.org/x/tools/cmd/godoc
