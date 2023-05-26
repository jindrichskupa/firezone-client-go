# docker run --rm firezone/firezone bin/gen-env > .env
docker-compose down && rm -rf firezone caddy
docker-compose run --rm firezone bin/migrate
docker-compose run --rm firezone bin/create-or-reset-admin
docker-compose up
