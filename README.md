# social-networks-parsing
Golang service that parse user's info from social networks while user is signing up


## Сервис регистрации пользователей
Написать REST сервис для регистрации пользователей и автоматического поиска пользователя в по различным социальным сетям. 

### Методы
* Создание пользователя. Валидация, запись с базу, отправка письма, поиск данных по соц.сетям в фоне, оповещение в телеграм о новом пользователе + ссылки на соц.сети
* Получения списка пользователей + фильтры
* Получение пользователя + вся найденная информация по соц.сетям

Вся информация должна храниться в БД
Поиск должен осуществляться в фоне
