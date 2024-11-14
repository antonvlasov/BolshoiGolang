# Занятие 10

## Архитектурные подходы в Go

Все основные подходы базируются на принципе Dependency Injection из [SOLID](https://en.wikipedia.org/wiki/SOLID).

Рассмотрим два подхода, которые говорят про одно и то же, но немного в разных терминах.

### Go clean template

[https://github.com/evrone/go-clean-template](https://github.com/evrone/go-clean-template)

### Go hexagonal architecture

[https://github.com/LordMoMA/Hexagonal-Architecture](https://github.com/LordMoMA/Hexagonal-Architecture)

## JWT

Основным ресурсом для знакомства с JSON Web Token является [сайт](https://jwt.io) и [RFC7519](https://datatracker.ietf.org/doc/html/rfc7519).

Основным достоинством JWT токенов является их подписание выдающей стороной. Это позволяет проверять информацию в таком токене с помощью публичного ключа сервиса выпустившего такой токен. Из этого напрямую вытекает применение - после выпуска JWT токен можно не хранить в базе, а проверить его содержимое с помощью ключа.

## gRPC

https://grpc.io/docs/what-is-grpc/introduction/
