initdb :
	docker run -d -p 5432:5432 -e POSTGRES_PASSWORD='root' --rm  --name=postgres postgres

postgres:
	docker exec -it postgres psql -U postgres

.PHONY: initdb postgres