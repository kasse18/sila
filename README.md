https://github.com/kasse18/silabot - тг бот

Deploy:

docker build -t sila-bot . 

docker run sila-bot

Использование: Если человек пишет в боту вопрос, то бот присылает ему ответ, как и в веб-версии, если человек задает некорректный вопрос, то бот перенаправляет ID чата и вопрос в беседу с сотрудниками тех-поддержки.

https://github.com/kasse18/sila - бекенд для админ панели

GET /container/get_all_containers - получить все контейнеры

POST /container/create_container - Content-type application\json, body: {"name": "name", "link_small": "small iframe link", "link_big": "big iframe link"} - создать контейнер

POST /container/login - Content-type application\json, body: {"login":"login", "password": "password"} - авторизоваться

POST /container/upload/{containerID} - Content-type multipart/form-data, body: {"file": your_file} - загрузить новый документ для контейнера

Deploy:

docker-compose up -d —build

Если контейнер sila-app падает сразу после билда, то повторить билд(зависит от постгри) - ошибка: 
panic: dial tcp 172.19.0.2:5432: connect: connection refused

goroutine 1 [running]:
sila-app/pkg/db/postgres.New({0xa841a8, 0xc0001ba510}, {0xc000028059?, 0xa7c5d0?})
        /app/pkg/db/postgres/postgres.go:16 +0xf4
main.main()
        /app/cmd/app/main.go:17 +0x8f
