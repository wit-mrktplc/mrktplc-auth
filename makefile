deploy:
	docker build -t mrktplc-auth .
	docker run -p 8081:8081 -e AUTHORIZED_DOMAINS=wit.edu,alumni.wit.edu mrktplc-auth