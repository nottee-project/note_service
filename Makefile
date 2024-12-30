args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`
env = .env
composefile = ./docker/docker-compose.yml

dc-up: dc-down
	docker compose -f $(composefile) --env-file $(env) up --build -d

dc-down:
	docker compose -f $(composefile) --env-file $(env) down

dc-config:
	docker compose -f $(composefile) --env-file $(env) config

dc-migrate:
	docker compose -f $(composefile) --env-file $(env) run --rm --entrypoint "/bin/sh" app -c "./migrations.sh $(call args)"

clean:
	rm -rf ./bin ./cover.out.tmp

.PHONY: lint lint-fast test test-cover dc-up dc-down dc-config dc-migrate clean
