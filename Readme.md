<h1>API для управления кошельками</h1>

<h2>Запуск</h2>
<p>docker-compose -f .docker/docker-compose.yml up --build</p>
<p>Необходимо освободить порты 8080 и 5432</p>

<h2>Выбор архитектуры</h2>
<p>Я выбрал Clean architecture, а не микросервисы потому что приложение слишком маленькое. Были использованы вспомогательные библиотеки 
для настройки роутинга и работы с http протоколом. Уровни доступа: </p>
<p>1. Delivery (внешний мир)</p>
<p>2. Use Cases (бизнес логика)</p>
<p>3. Repositories (работа с бд)</p>
<p>4. Domains (структура сущностей)</p>
<p>Также я вынес инициализация в bootstrap, чтобы функция main была локоничнее
</p>

<h2>Тесты</h2>
<p>1. Для просмотра сгенерированных кошельков использовать docker exec -it my_postgres bash -c "PGPASSWORD=password psql -U user -d transactions_db -c 'SELECT * FROM wallets;'"</p>
<p>2. Проверка команды send. Пример запроса curl --location 'http://localhost:8080/api/send' \
--header 'Content-Type: application/json' \
--data '{
  "from": "a3479dbbcc4deff7ccf5a6987a5a82cac479ae52f985e04f93f92f8f3a88fbd4",
  "to": "fe5de72ed40ca888f9736b6007d5aa91feb439ffbdf49c6d4d414ccfbd760d6d",
  "amount": 1
}
'</p>
<p>3. Проверка команды get-balance. Пример запроса curl --location 'http://localhost:8080/api/wallet/fe5de72ed40ca888f9736b6007d5aa91feb439ffbdf49c6d4d414ccfbd760d6d/balance'
</p>
<p>4. Проверка команды transactions. Пример запроса curl --location 'http://localhost:8080/api/transactions?count=2''
</p>

