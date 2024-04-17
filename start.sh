#! /bin/sh
# 使用的是 alpine 的 docker 镜像，所以不能使用 bash

set -e

echo "run db migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"
