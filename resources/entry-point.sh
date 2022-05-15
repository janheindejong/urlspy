#!/bin/sh
set -e

if [ "$RUN_MIGRATION" = "1" ] 
then 
    echo "Running database migration..."
    alembic upgrade head 
fi

exec "$@"