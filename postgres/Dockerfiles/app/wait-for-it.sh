#! /bin/sh

set -ex


sleep 5

echo "--start app--"

exec "$@"

#host="$1"
#shift
#cmd="$@"
#
#until psql -h "$host" -U "postgres" -c '\l'; do
#	  >&2 echo "Postgres is unavailable - sleeping"
#	    sleep 1
#done
#
#>&2 echo "Postgres is up - executing command"
#exec $cmd
