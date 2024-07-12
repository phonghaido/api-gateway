docker-build:
	docker build -t apigateway-nginx:latest .

deploy:
	kubectl apply -k ./k8s

all: docker-build deploy