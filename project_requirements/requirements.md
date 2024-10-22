В рамках первого проекта вам требуется сделать in-memory key-value базу данных с HTTP интерфейсом, возможностью персистентного хранения и автоматического удаления устаревших записей.

## Типы хранимых данных

Все ключи базы данных - строки. По одному ключу может быть только одно значение одного из типов.
Значения базы данных могут быть одним из следующих типов:

- скалярами
- словарь
- массивами

### Скаляр

Скаляр - единичное значения типа строка либо целое число. Способ внутреннего представления скаляров - на ваш выбор (можно хранить все как строки, как интерфейсы, и тп.). При чтении данных у пользователя должна быть возможность явно узнать тип скаляра (строка/число). Далее в описаниях операция числа без кавычек, а строки в кавычках. Подробнее возможные форматы передачи ответов и запросов описаны в разделе [интерфейс](#интерфейс)

#### Операции по работе со скалярами

##### GET key

Возвращает значение по ключу key. Если по ключу нет значения, возвращается nil. Если значение по ключу не типа скаляр, возвращается ошибка. Возвращаемые строки должны быть заключены в кавычки, возвращаемые числа должны быть без кавычек.
Пример:

```
SET name "Anton"
OK
SET number 42
OK
GET name
"Anton"
GET PUE
(nil)
GET number
42
```

##### SET key value [EX seconds]

Устанавливает значение по ключу key равным value. Если value указано в кавычках - это строка, иначе это число. Если переданное без кавычек значение не является числом, должна возвращаться ошибка. Если указан дополнительный параметр EX seconds, значение очищается через заданное количество секунд. Значение seconds 0 означает хранение без ограничений по времени.
Пример:

```
SET key1 "value1" EX 20
OK
SET key2 value1 EX 20
Requested value is not a number
```

### Словарь

Словарь - структура, хранящая в своих полях скаляры. Поле мапы задается строковым ключем. С помощью словаря можно по определенному ключу нашей базы данных положить не просто одно значение, а набор полей.

#### Операции по работе со словарями

##### HSET key field value [field value ...]

Устанавливает поле словаря field, являющимся значением по ключу key равным value. Возможно установить сразу несколько полей. Если значение указанного ключа (key) является другим типом, возвращается ошибка. Возвращает количество затронутых полей.
Пример:

```
SET key1 1
OK
HSET key1 hash1 "val1"
Requested field is of type scalar
HSET key1 hash1 "val1" hash2 "val2"
(integer) 2
```

##### HGET key field

Возвращает значение поля field словаря по ключу key. Если по ключу key находится другой тип, возвращается ошибка. Если значение поля field не задано или значение по ключу key не задано, возвращается nil.
Пример:

```
SET name Anton
OK
HSET key1 hash1 "val1" hash2 "val2"
(integer) 2

HGET key1 hash1
"val1"
HGET key1 hash2 
"val2"
HGET key1 hash3
(nil)
HGET key434 hash2
(nil)
HGET name hash
Requested field is of type scalar
```

### Массив

Массив позволяет по определенному ключу базы данных хранить упорядоченный массив скаляров.

#### Операции по работе с массивами

### LPUSH key element [element ...]

Вставляет элементы слева в список по ключу key. Если элементов несколько, они вставляются так, как будто для каждого из них по порядку была бы вызвана эта команда. Если значения по ключу не существовало, список создается.
Пример:

```
LPUSH list1 1 2 3
(integer) 3
LPOP list1 0 -1
1)3
2)2
3)1
```

### RPUSH key element [element ...]

Вставляет элементы справа в список по ключу key. Если элементов несколько, они вставляются так, как будто для каждого из них по порядку была бы вызвана эта команда. Если значения по ключу не существовало, список создается.
Пример:

```
RPUSH list1 1 2 3
(integer) 3
LPOP list1 0 -1
1)1
2)2
3)3
```

### RADDTOSET key element [element ...]

Вставляет справа в список по ключу key элементы, которых еще нет в списке. Если элементов несколько, они вставляются так, как будто для каждого из них по порядку была бы вызвана эта команда. Если значения по ключу не существовало, список создается.
Пример:

```
RPUSH list1 1 2 3
(integer) 3
RADDTOSET list1 3 5 8 4 8
LPOP list1 0 -1
1)1
2)2
3)3
4)5
5)8
6)4
```

### LPOP key [count]

Удаляет и возвращает элемент слева списка. Параметр count - количество удаляемых элементов, может быть либо единственным числом - тогда это количество элементов с края, либо двумя числами - тогда это индексы первого и последнего удаляемых элементов. Индексы могут быть отрицательными для доступа с конца списка. Если количество удаляемых элементов превышает количество элементов в списке, возвращается доступное количество.
Пример:

```
RPUSH list1 1 2 3 4 5 6 7 8 9 10
(integer) 10
LPOP list1 2
1)1
2)2

LPOP list1 2 -2
1)5
2)6
3)7
4)8
5)9

LPOP list1
3
```

### RPOP key [count]

Удаляет и возвращает элемент слева списка. Параметр count - количество удаляемых элементов, может быть либо единственным числом - тогда это количество элементов с края, либо двумя числами - тогда это индексы первого и последнего удаляемых элементов. Индексы могут быть отрицательными для доступа с конца списка. Если количество удаляемых элементо превышает количество элементов в списке, возвращается доступное количество.
Пример:

```
RPUSH list1 1 2 3 4 5 6 7 8 9 10
(integer) 10
RPOP list1 2
1)10
2)9

RPOP list1 2 -2
1)7
2)6
3)5
4)4
5)3

RPOP list1
8
```

### LSET key index element

Устанавливает значение элемента с индексом index списка по ключу key равным element. Если элемента с этим индексом не существует, возвращается ошибка.
Пример:

```
RPUSH list1 0 1 2 3 4 5 6 7 8 9
(integer) 10
LSET list1 3 30
OK
LGET list1 3
30
LSET list1 20 2
Index out of range
```

### LGET key index

Получает значение элемента с индексом index из списка по ключу key. Если элемента с этим индексом не существует, возвращается ошибка.
Пример:

```
RPUSH list1 0 1 2 3 4 5 6 7 8 9
(integer) 10
LSET list1 3 30
OK
LGET list1 3
30
LSET list1 20 2
Index out of range
```

## Дополнительные операции

### EXPIRE key seconds

Для любого ключа нашей базы данных можно явно задать время жизни в секундах. После истечения времени жизни ключа все операции по этому ключу должны работать так, как будто этого ключа нет в базе данных. Если по указанному ключу существует значение, возвращает 1, иначе 0. Должен быть предусмотрен механизм "garbage collection" - физически удалять автоматически из нашей базы данных сильно устаревшие ключи для освобождения места.
Пример:

```
SET key1 1   
OK
EXPIRE key3 30
(integer) 0
EXPIRE key1 20
(integer) 1
...after 20 seconds
GET key1
(nil)
```

### KEYS pattern

Возвращает список ключей, удовлетворяющих регулярному выражению.
Пример:

```
SET name Anton
OK
SET age 20
OK
KEYS .*?
1) "name"
2) "age"
```

## Сохранение данных

База данных должна переодически сохранять свое состояние на диск для восстановления после сбоев. Должно сохраняться состояние базы, устаревшее не более чем на одну минуту. Для сохранения состояния базы данных требуется использовать sqlite (пока не дошли до этого, сохранять в файл как JSON).
При запуске база данных должна проверять наличие состояния и восстанавливать данные из него (если файл состояния есть).

## Интерфейс.

База данных должна предоставлять HTTP интерфейс. Для получения данных использовать GET, для записи данных - POST. Каждая операция должна быть отдельным эндпоинтов HTTP. Схожие операции нужно группировать по сходим url путям (например, /array/set/l, /array/set/l, /array/get/l).
Если HTTP интерфейсы еще не изучили, достаточно иметь возможность задать последовательность операций как вызов последовательности функций в main()
Возможные варианты описания значений по ключам:

- Передавать все комманды через тело HTTP запроса как в примерах из этого документа в кодировке UTF-8 (стандартная go строка в байтах), различия между строками и числами определяются кавычками.
- Использовать структуру для описания значения, под строковое и числовое варианты будут 2 разных поля. Кодировка любая (JSON, msgp, ...)
- Использовать структуру для описания значения, любые значения кодировать как строки и добавить отдельное поле с явным указанием типа. Кодировка любая.
  Выбирайте то что считаете самым лучшим вариантом. Единственным объективным критерием тут может быть скорость чтения/записи значений.

## Docker-compose

Требуется поднимать наше приложение и его базу данных sqlite c помощью docker-compose.