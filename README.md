# Restaurants
1. Описание: Приложение для бронирования слоликов в ресторанах
2. Процесс установки на Ubuntu
3. Установить Golang
4. В консоли выполняем  эти команды
    wget https://go.dev/dl/go1.18.4.linux-amd64.tar.gz; 
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.4.linux-amd64.tar.gz;
    export PATH=$PATH:/usr/local/go/bin;
5. Проверяем версию командой go version
6. устанавливаем докер: sudo apt-get install docker
7. Скачиваем образ постргеса: docker pull postgres
8. Создадим образ: docker run --name=restaurants-db -e POSTGRES_PASSWORD="zxcvbn" -p 5437:5432
9. Запускаем контейнер с БД: docker start  restaurants-db
10. Установим migrate sudo apt-get install migrate
11. Для создания структуры таблиц внутри контейнера: migrate ./schema -database 'postgres://postgres:zxcvbn@localhost:5437/postgres?sslmode-disable' up
11. Скачиваем репозиторий: git clone git@github.com:goodGopher/Restaurants.git
12. Запускаем программу go run ./Restaurants/cmd/main.go

PS: приложение не закончено. Основная проблема - главная функция бронирования не работает. Также не реализована обработка некоторых запросов. 
Реализован функционал добавления ресторанов и столиков, а также некоторые функции манипуляции с ними.
