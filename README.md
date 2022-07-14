# Restaurants
1. Описание проекта
2. Процесс установки
    (пока что очень обще как это делал я, потом надо описать подробнее/удобнее)
    2.1 Установить Golang
        2.1.1 Заходим на страницу https://go.dev/doc/install
        2.1.2 Копируем ссылку на скачивание установщика https://go.dev/dl/go1.18.4.linux-amd64.tar.gz
        2.1.3 В консоли выполняем  эти команды
        wget https://go.dev/dl/go1.18.4.linux-amd64.tar.gz; 
        rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.4.linux-amd64.tar.gz;
        export PATH=$PATH:/usr/local/go/bin;
        2.1.4 Проверяем версию командой
        go version
        
