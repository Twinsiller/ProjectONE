# ProjectONE

--Удаление данных из таблицы и обнуление значение serial:
    TRUNCATE TABLE db_name RESTART IDENTITY;

--Библиотеки:
    _ "github.com/lib/pq"
    password "github.com/vzglad-smerti/password_hash"

--Docker:
    build: "docker build -t myapp:latest ."
    For launch: "docker run -e POSTGRES_PASSWORD=qwerty -e POSTGRES_DB=ProjectONE -p 8080:8080 -p 5432:5432 myapp:latest"

ProjectONE:
    для запуска:
        1. go mod init ProjectONE **Модули**
        2. go mod tidy **Обновление модулей**
        3. go get -u github.com/gin-gonic/gin **Сервер**
        4. go get -u github.com/lib/pq **БД PostgreSQL**
        5. go get -u github.com/vzglad-smerti/password_hash **Хэширование паролей**