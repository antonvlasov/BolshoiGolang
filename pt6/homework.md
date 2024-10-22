# 1. Добавить TTL команде SET

Ограничения смотрите в секции Set requirements.md

Предлагаю посмотреть, как это реализовано в Redis: 

> How Redis expires keys.
>
> Redis keys are expired in two ways: a passive way and an active way.
> A key is passively expired when a client tries to access it and the key is timed out.
> However, this is not enough as there are expired keys that will never be accessed again.
> These keys should be expired anyway, so periodically, Redis tests a few keys at random amongst the set of keys with an expiration. 
> All the keys that are already expired are deleted from the keyspace.

Подробнее про expire [тут](https://redis.io/docs/latest/commands/expire).

Вам нужно сделать аналогично: протухание "пассивное" и "активное".

Пассивно удалять записи с установленным временем жизни (TTL/EX) при попытке получить их через SET.

Активно удалять записи фоновой горутиной, алгоритм выборки выбирайте сами, но постарайтесь нjе лочить мапу надолго (перебирать все значения).
