#!/bin/bash
# wait-for-mssql.sh

set -e

host="$1"
shift
cmd="$@"

until /opt/mssql-tools/bin/sqlcmd -S mssql -U sa -P 'YourStrong!Passw0rd' -Q "SELECT 1;" > /dev/null 2>&1; do
  >&2 echo "SQL Server is unavailable - sleeping"
  sleep 2
done

>&2 echo "SQL Server is up - executing command"
exec $cmd
