APP_BIN = app/build/app

lint:
	golangci-lint run

build: clean $(APP_BIN)

$(APP_BIN):
	go build -o $(APP_BIN) ./app/cmd/main.go

clean:
	rm -rf ./app/build || true

swagger:
	swag init -g ./app/cmd/main.go -o ./app/docs

migrate:
	$(APP_BIN) migrate -version $(version)

migrate.down:
	$(APP_BIN) migrate -seq down

migrate.up:
	$(APP_BIN) migrate -seq up

git:
	git add .
	git commit -a -m "$m"
	git push -u origin master

get_reg:
	curl -X POST localhost:10000/rgfdart -H "Content-Type: application/json" -d '{"regulation_id": 1}' -o regulation.dart

get_links:
	curl -X POST localhost:10000/lfdart -H "Content-Type: application/json" -d '{"regulation_id": 1}' -o alllinks.dart