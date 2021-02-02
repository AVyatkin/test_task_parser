# Тестовое задание
https://git.linxdatacenter.com/tt/test22/tree/master

Как развернуть решение:
1. В системе должены быть установлены утилиты: docker, docker-compose
2. Перейти в папку проекта
3. Указать порт для сервиса в файле .env - константа LISTEN_PORT
4. Выполнить запросы:
  - docker-compose build
  - docker-compose up -d
5. Если нужно приостановить сервис, то необходимо запустить команду
  - docker-compose down
