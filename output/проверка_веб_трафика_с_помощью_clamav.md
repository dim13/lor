## Общее

Для онлайн проверки на вирусы веб трафика можно использовать
[HAVP](http://www.server-side.de) (HTTP AntiVirus proxy). Мы рассмотрим
использование его в паре со Squid:

> Пользователь-\>Squid-\>Havp-\>WWW

Здесь Squid является прозрачным кэширующим прокси с ведением статистики,
поэтому havp установлен после него. [Lissyara](http://lissyara.su)
приводит и другие аргументы в пользу такой схемы.

## Установка

Предположим, что squid уже установлен и настроен. Устанавливаем havp и
вносим изменения в конфиг havp (/etc/havp/havp.config). Закомментируем
строку <i>REMOVETHISLINE deleteme</i> и раскомментируем следующее:

    USER havp
    GROUP havp
    DAEMON true

Меняем значение переменной (в зависимости от нагрузки):

    SERVERNUMBER 20

Раскомментируем:

    PORT 8080
    BIND_ADDRESS 127.0.0.1

В */etc/havp/templates/en* поле *TEMPLATEPATH* меняем *en* на *ru* и
удостоверяемся, что незакомментирована строка *ENABLECLAMLIB true*

Стартуем havp - */etc/init.d/havp start*

Вносим изменения в настройки squid:

    cache_peer 127.0.0.1 parent 8080 0
    always_direct allow SSL_Ports

Рестартуем squid - */etc/init.d/squid restart*. (Если нет delay pools,
то можно reload).

## Проверка

Проверить, что havp слушает, можно командой *netstat -an |grep 8080*

Проверить связку можно по адресу
[](http://www.eicar.org/anti_virus_test_file.htm)

## Источники

  - [1](http://ylsoftware.com/?action=news&na=viewfull&news=390)


  - [2](http://www.lissyara.su/?id=1662)


  - [3](http://www.linux16.net/node/390)

