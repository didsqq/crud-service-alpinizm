#!/bin/sh

echo "Жду, пока база и миграции завершатся..."
until pg_isready -h db -U $DB_USER; do
  sleep 2
done

# Проверим, что нужная таблица создана
until PGPASSWORD=$DB_PASSWORD psql -h db -U $DB_USER -d $DB_NAME -c "SELECT * FROM alpinists LIMIT 1;" > /dev/null 2>&1
do
  echo "Ожидание миграций..."
  sleep 2
done

echo "Миграции завершены. Запускаю приложение..."
exec ./main
