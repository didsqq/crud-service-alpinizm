# crud-service-alpinizm
### Работа с ms sql в докере
docker exec -it --user root mssql_container bash - заходим под правами root

// установка утилиты sqlcmd в докере
apt update && apt install -y unixodbc-dev msodbcsql17 mssql-tools
echo 'export PATH="$PATH:/opt/mssql-tools/bin"' >> ~/.bashrc
source ~/.bashrc 

// можно установить sqlcmd в самой системе и подключаться к бд, не заходя в докер
// запуск sqlcmd 
sqlcmd -S localhost -U sa -P 'YourStrong!Passw0rd'
-S — сервер (например, localhost или 127.0.0.1).
-U — логин (по умолчанию sa).
-P — пароль.

// Выводит существующие бд
SELECT name FROM sys.databases;
GO

// Выводит таблицы в бд
SELECT name FROM sys.tables;
#: TODO: создание бд автоматически
