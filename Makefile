build:
		go build -o fuzzer *.go

run:
		go run *.go

clean:
		rm -rf fuzzer fuzzing-*

