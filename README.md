# microservice-for-working-with-user-balance

## Техническое задание
Необходимо реализовать микросервис для работы с балансом пользователей (зачисление средств, списание средств, перевод средств от пользователя к пользователю, а также метод получения баланса пользователя). Сервис должен предоставлять HTTP API и принимать/отдавать запросы/ответы в формате JSON.

1. Метод начисления средств на баланс. Принимает id пользователя и сколько средств зачислить.
2. Метод списания средств с баланса. Принимает id пользователя и сколько средств списать.
3. Метод перевода средств от пользователя к пользователю. Принимает id пользователя с которого нужно списать средства, id пользователя которому должны зачислить средства, а также сумму.
4. Метод получения текущего баланса пользователя. Принимает id пользователя. Баланс всегда в рублях.

Детали по заданию:

1. Методы начисления и списания можно объединить в один, если это позволяет общая архитектура.
2. По умолчанию сервис не содержит в себе никаких данных о балансах (пустая табличка в БД). Данные о балансе появляются при первом зачислении денег.
3. Валидацию данных и обработку ошибок оставляем на усмотрение кандидата.
4. Список полей к методам не фиксированный. Перечислен лишь необходимый минимум. В рамках выполнения доп. заданий возможны дополнительные поля.
5. Механизм миграции не нужен. Достаточно предоставить конечный SQL файл с созданием всех необходимых таблиц в БД.
6. Баланс пользователя - очень важные данные в которых недопустимы ошибки (фактически мы работаем тут с реальными деньгами). Необходимо всегда держать баланс в актуальном состоянии и не допускать ситуаций когда баланс может уйти в минус.
8. Валюта баланса по умолчанию всегда рубли.

Дополнительные задания

Доп. задание 1:

Эффективные менеджеры захотели добавить в наши приложения товары и услуги в различных от рубля валютах. Необходима возможность вывода баланса пользователя в отличной от рубля валюте.

Задача: добавить к методу получения баланса доп. параметр. Пример: ?currency=USD. Если этот параметр присутствует, то мы должны конвертировать баланс пользователя с рубля на указанную валюту. Данные по текущему курсу валют можно взять отсюда https://exchangeratesapi.io/ или из любого другого открытого источника.

Примечание: напоминаем, что базовая валюта которая хранится на балансе у нас всегда рубль. В рамках этой задачи конвертация всегда происходит с базовой валюты.

Доп. задание 2:

Пользователи жалуются, что не понимают за что были списаны (или зачислены) средства.

Задача: необходимо предоставить метод получения списка транзакций с комментариями откуда и зачем были начислены/списаны средства с баланса. Необходимо предусмотреть пагинацию и сортировку по сумме и дате.

## Инструкция

Микросервис можно запустить в docker-compose командой: `docker-compose up -d --build`. Но перед этим в файле comfig.yaml нужно пеменять адрес хоста: с `host: localhost` на `host: host.docker.internal`. Файл main.go находиться в папке cmd/main/, скрипт создания таблиц в базе данных up.sql находится в папке postgres/, файл конфигурации config.yaml - в корневой папке, все поля конфигурации с коментариями. 
  

Примеры HTTP запросов:
1. Начисление средств на баланс:
- метод - POST;
- URL - `localhost:8080/api`;
- тело запроса - 
``{
    "id": "z",
    "amount": 7654
}``.

2. Метод списания средств с баланса: 
- метод - PUT;
- URL - `localhost:8080/api/write-off`;
- тело запроса - 
``{
    "id": "z",
    "amount": 1000
}``.

3. Метод перевода средств от пользователя к пользователю:
- метод - PUT;
- URL - `localhost:8080/api/transaction`;
- тело запроса - 
``{
   "sender": "z",
   "recipient": "x",
   "amount": 400
}``.

4. Метод получения текущего баланса пользователя:
- метод - GET;
- URL - `localhost:8080/api/z`;
- z - id пользователя

5. Метод получения списка операций:
- метод - GET
- URL - `localhost:8080/api/history/z`;
- z - id пользователя

6. Метод получения сортированного списка операций, сортировка зависит от последнего указанного параметра:
- метод - GET;
- URL - `localhost:8080/api/history/z/amount<`;
- z - id пользователя, amount< - способ сортировки, ниже приведены все возможные.

7. Метод конвертации баланса пользователя на выбранную валюту:
- метод - GET;
- URL - `localhost:8080/api/convert/z/eur`;
- z - id пользователя, eur - валюта конвертации.

Варинты сортировки: 
1. Параметр `date<` - по дате по возрастанию.
2. Параметр `date>` - по дате по убыванию.
3. Параметр `amount<` - по сумме операции по возрастанию.
4. Параметр `amount>` - по сумме операции по убыванию.
5. Параметр `balance<` - по сумме баланса по возрастанию.
6. Параметр `balance>` - по сумме баланса по убыванию.

## Документация Postman для запросов 
https://documenter.getpostman.com/view/21487929/2s83RzjasV#e578ea5f-1bb2-40dd-807c-913954c01b24