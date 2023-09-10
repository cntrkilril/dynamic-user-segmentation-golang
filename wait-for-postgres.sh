#!/bin/sh
# wait-for-postgres.sh

set -e

host="$1"
shift
cmd="$@"

until PGPASSWORD="pass" psql -h "$host" -U "root" -c '\q'; do
  echo >&2 "Postgres is unavailable - sleeping"
  sleep 1
done

echo >&2 "Postgres is up - executing command"
exec $cmd
