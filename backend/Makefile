build:
	GOOS=linux go build -o app
	docker build -t michaelbi/say-grpc .
	rm -f app
	