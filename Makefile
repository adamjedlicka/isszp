default: clean build

build:
	mkdir ./dist
	go build ./main.go
	mv ./main ./dist/isszp
	cp -rf ./template ./static ./config ./dist

clean:
	rm -rf ./dist

doc:
	@echo Documentation running on: http://localhost:6060/pkg/gitlab.fit.cvut.cz/isszp/isszp/
	godoc -http localhost:6060
