list:
	go run main.go list tender -c hu
get:
	go run main.go get tender 1234 -c hu
build:
	go build -o dist/tenders-cli
run:
	./dist/tenders-cli