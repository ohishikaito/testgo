attach:
	docker attach testgo_app_1

test:
	docker-compose exec app go test -v -cover ./...
