# 1. Добавить новые операции
Добавить отдельные методы для работы с массивами (массивы хранить в другой мапе, не там же где базовые типы и предыдущих занятий). Далее идет текстовое описание каждой новой операции, как если бы работы с приложением происходила через интерактивную командную строку (например, как в оболочками SQL). Вам нужно написать только функции для выполнения этих операций, необходимую сигнатуру определите сами.

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
1
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

# 2. Добавить возможность сохранения и чтения с диска
Добавить возможность сохранения и чтения состояния нашей базы данных на диск/с диска. Сохранение на диск должно происходить перед выходом из main, чтение должно происходить перед началом работы, в начале main. Сами команды по работе с базой идут в блоке между загрузкой предыдущего состояния и сохранением текущего. Для сохранения и чтения с диска рекомендуется использовать json-encoded файл.

```
main(){
    // read from disk if exists

    // do operations

    // save to disk
}

```