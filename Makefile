test:
	@echo 'Мы сделали Makefile'

up:
	sudo docker-compose up --build structs

stop:
	sudo docker-compose stop
