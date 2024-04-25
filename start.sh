#! /bin/sh
# 使用的是 alpine 的 docker 镜像，所以不能使用 bash

set -e

# echo "run db migration"
# source  /app/app.env  # 加载环境变量后，将 DB_SOURCE 覆盖掉了，从原先的 postgre 变成 localhost，导致了错误
# /app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"
