composefile = ./docker/docker-compose.yml
project = task-service
env = .env

dc-up:
	docker compose -p $(project) -f $(composefile) --env-file $(env) up --build -d

dc-down:
	docker compose -p $(project) -f $(composefile) --env-file $(env) down

dc-config:
	docker compose -p $(project) -f $(composefile) --env-file $(env) config

dc-migrate:
	docker compose -p $(project) -f $(composefile) --env-file $(env) run --rm --entrypoint "/bin/sh" app -c "./migrations.sh $(call args)"

clean:
	rm -rf ./bin ./cover.out.tmp

.PHONY: lint lint-fast test test-cover dc-up dc-down dc-config dc-migrate clean
