docker-build:
	docker build -t nginx:latest ./nginx
	docker build -t common-api:latest ./app

deploy:
	kubectl apply -k ./k8s/_base
	kubectl rollout restart -n artifactory deployment/apigateway

all: docker-build deploy