build:
	docker build -t mrktplc-auth .
deploy:
	docker run -d -p 8081:8081 -e AUTHORIZED_DOMAINS=wit.edu --name auth-service mrktplc-auth
