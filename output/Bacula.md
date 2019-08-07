Система резервного копирования "bacula"

«It comes in the night and sucks the essence from your computers». Kern
Sibbald

# Обзор

## Отзывы

О системе можно встретить много положительных отзывов:

  -
    масштабируема,
    автоматизирована,
    возможно восстановление с нуля (bare metal, не восстанавливаются
    атрибуты файловых систем и программный RAID)
    и многое другое.

Из отрицательных отзывов:

  -
    плохо переносит резкие смены конфигурации,
    плохая обработка ошибок при работе с диском (заполненный диск, глюки
    с созданием нового тома при наличии файла),
    много ручной работы (там-же: после настройки работает в полностью
    автоматическом режиме).

источник: [Bog BOS](http://bog.pp.ru/work/bacula.html)

## Компоненты

[Изображение](http://www.bacula.org/5.0.x-manuals/en/main/main/img7.png).

Здесь:

  -
    Bacula Director
    (<span style="background-color: rgb(222, 203, 182);">bacula-dir</span>):
    управляющий демон
    Bacula Console
    (<span style="background-color: rgb(221, 185, 192);">bconsole</span>):
    программа управления (на машине администратора)
    Bacula File
    (<span style="background-color: rgb(201, 210, 223);">bacula-fd</span>):
    собственно демон бекапа (на машине с которой выполняется резервное
    копирование)
    Bacula Storage
    (<span style="background-color: rgb(221, 222, 182);">bacula-sd</span>):
    хранилище
    Catalog: база данных
    (<span style="background-color: rgb(207, 233, 224);">MySQL, PgSQL,
    SQlite</span>)
    Bacula Monitor: монитор
    (<span style="background-color: rgb(221, 185, 192);">значок в
    трее</span>)

## Документация

  - Официальная документация на
    [www.bacula.org](http://www.bacula.org/en/?page=documentation).
  - Частичный перевод руководства есть на
    [www.opennet.ru](http://www.opennet.ru/opennews/art.shtml?num=7385).
  - Статья на
    [ibm.com](http://www.ibm.com/developerworks/ru/library/l-Backup_4/).
  - Материалы на [Bog BOS](http://bog.pp.ru/work/bacula.html).

# Установка

Основная статья: [Bacula/Быстрый старт](Bacula_Быстрый_старт "wikilink")

Установка по умолчанию, только до состояния "работает".

Все настройки по умолчанию.

# Pools, Volumes and Labels

Основная статья: [Bacula/Pools, Volumes and
Labels](Bacula_Pools,_Volumes_and_Labels "wikilink")

Понимание использования пулов, томов и меток.

# Планирование

## Использование накопителей

Теперь, на имеющемся материале, имеет смысл определиться со структурой
резервного копирования:

  - распределить какие бекапы на какие накопители писать (в случае
    файлов: винчестеры и разделы)
  - какие пулы при этом использовать
  - и описать это, распределив по соответственным конфигурационным
    файлам.

[Bacula/Использование
накопителей](Bacula_Использование_накопителей "wikilink")

## Вынос части конфигурации в отдельные файлы

Bacula
[FAQ](http://wiki.bacula.org/doku.php?id=faq#can_bacula-dirconf_include_others_files):

Yes, the Director configuration doesn't have to be in just one file. You
can do this:

    @/path/to/file1
    @/path/to/file2

In fact, the @filename can appear anywhere within the conf file where a
token would be read, and the contents of the named file will be
logically inserted in the place of the @filename. What must be in the
file depends on the location the @filename is specified in the conf
file.

    #
    # Include subfiles associated with configuration of clients.
    # They define the bulk of the Clients, Jobs, and FileSets.
    # Remember to "reload" the Director after adding a client file.
    #
    @|"sh -c 'for f in /etc/bacula/clientdefs/*.conf ; do echo @${f} ; done'"

<http://wiki.bacula.org/doku.php?id=sample_configs>

<http://www.bacula.org/5.0.x-manuals/en/main/main/Customizing_Configuration_F.html#SECTION001723000000000000000>

# Подключение fd с другого компьютера

## Безопасность

Установка на компьютере fd означает что к нему можно обратиться и
получить с него любые файлы. По мере необходимости, помимо
достаточно больших паролей следует озаботиться и другими
аспектами безопасности:

  -
    Ограничить адреса с которых можно обратиться.
    Закрыть брандмауером.
    Организовать шифрованный канал:
    [пример](http://www.bog.pp.ru/work/bacula.html#TLSexample).

## Настройка

Для подключения удалённого fd:

1.  В конфиге fd на компьютере, с которого нужно делать резервную копию:
      - Указываем директора которому можно подключаться к этому клиенту.
2.  В конфиге директора, который будет управлять этим fd (это можно
    вынести в отдельный файл, что бы не раздувать общий конфиг
    директора):
      - Описываем клиента (fd), к которому будем подключаться.
      - Описываем пул (необязательно).
      - Перечисляем что копировать (fileset) (необязательно).
      - Описываем расписание (schedule) (необязательно).
      - Описываем задание (job), включая в него секции:
          - client
          - pool
          - fileset
          - schedule
          - *часть описания можно вынести в секцию JobDefs*

## Восстановление

При восстановлении по умолчанию файлы будут восстановлены на том же
клиентском компьютере, в каталог указанный в конфигурации директора
в задании (job) для восстановления.

# бекап MySQL базы

В настройках по умолчанию есть один пример копирования базы данных
(собственной базы бакулы), но скорее всего чаще будут встречаться
случаи резервного копирования с удалённых машин. При этом можно как
выполнять mysqldump локально, обращаясь к удалённой базе, так и
выполнять непосредственно на машине с СУБД, забирая затем файл
оттуда. Для последнего случая понадобятся директивы:

    ClientRunBeforeJob = /path/to/script/create.dump
    ClientRunAfterJob  = /path/to/script/remove.dump

См. также:

  - <http://sozinov.blogspot.com/2008/05/bacula-backup-mysql.html>
  - Так-же, в
    /usr/share/doc/bacula-common-\[version\]/examples/database/ для
    этого есть примеры.

[Category:Bacula](Category:Bacula "wikilink")