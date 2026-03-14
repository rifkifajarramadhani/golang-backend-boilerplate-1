export MYSQL_URL='mysql://root:rootpassword@tcp(db:3306)/forum_app'

migrate-create:
	docker compose exec web migrate create -ext sql -dir internal/db/migrations -seq $(name)

migrate:
	docker compose exec web migrate -database $(MYSQL_URL) -path internal/db/migrations $(args)