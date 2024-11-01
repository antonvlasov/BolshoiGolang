# 1. Сохранение 5 последних версий стейта хранилища в Postgres

Необходимо добавить Postgres в ваш проект. 

Поднимите его в docker compose вместе с основым приложением. 

Строку подключения к базе нужно передавать через environment переменные.

В базе нужно хранить стейт в виде JSONB.

Описание таблицы:

```
CREATE TABLE IF NOT EXISTS core (
		version bigserial PRIMARY KEY,
		timestamp bigint NOT NULL,
		payload JSONB NOT NULL
	)
```


Поскольку нужно хранить последние 5 версий, разрешается удалять все версии ниже.
