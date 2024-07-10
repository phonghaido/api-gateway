FROM nginx:alpine

COPY . /etc/nginx
COPY ./ssl/apigateway.crt /etc/ssl/certs/apigateway.crt
COPY ./ssl/apigateway.key /etc/ssl/private/apigateway.key

RUN chmod 600 /etc/ssl/private/apigateway.key && \
	chmod 644 /etc/ssl/certs/apigateway.crt

EXPOSE 443

CMD ["nginx", "-g", "daemon off;"]