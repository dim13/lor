**D-Bus** — это система межпроцессного взаимодействия, которая
предоставляет приложениям несколько шин для передачи
сообщений. Она обеспечивает беспроблемную связь десктопных
приложений между собой и связь между десктопными приложениями и
системными сервисами. Поддерживается не только широковещательная
рассылка сообщений (сигналов), но и удалённый вызов методов.

Некоторые системные сервисы, например как [HAL](HAL "wikilink") или WPA
Supplicant, предоставляют соответствующие D-Bus сервисы на системной
шине сообщений, таким образом мы можем с ними общаться с помощью
любого D-Bus клиента.

## Основные понятия

  - **Системная шина (System bus)** — данная шина создается при старте
    сервера D-Bus. К данной шине подключены системные сервисы (HAL,
    NetworkManager, bluez, WPA Supplicant...). На этой шине работают
    сервисы, которые нельзя отнести к какой-то определённой
    пользовательской сессии, и которые относятся к системе в
    целом.
  - **Сессионная шина (Session bus)** — данная шина создается на каждый
    вход пользователя в систему. К данной шине подключаются приложения
    залогинившегося пользователя, через нее проходит общение программ
    запущенных в данной рабочей сессии пользователя\[1\].
  - **Имя на шине (Bus name), Имя сервиса (Service name)** — каждая
    программа, подключенная к шине получает свое уникальное имя,
    которое начинается с двоеточия (":"). Если использовать аналогию
    с сетью, то это - IP адрес. Однако связь через IP адрес не очень
    удобна, поэтому используются доменные имена. Также и в D-Bus —
    приложение может взять несколько дополнительных имен, чтобы
    другие программы смогли связываться. Имя сервиса представляет
    собой набор из символов, разделенных точками ("."). Пример: сервис
    имеет реальное имя ":1.5", и символьное представление -
    "org.freedesktop.NetworkManager". Мы можем подсоединится и к ":1.5"
    и к "org.freedesktop.NetworkManager" - это один и тот же сервис.
    Создавать или нет символьное представление имени каждый клиент
    решает самостоятельно.
  - **Путь к объекту (Object path)** — представляет собой путь к некому
    объекту внутри адресуемого сервиса. Если продолжать аналогию с
    сетью, то путь к объекту — это "/wiki" в адресе
    <http://linux.org.ru/wiki>. Пример:
    "/org/freedesktop/NetworkManager/Device/eth0". Сервис может
    обслуживать несколько объектов.
  - **Интерфейс (Interface)** — каждый объект предоставляет доступ к
    некоторому набору методов и сигналов, который называется
    интерфейсом. При этом один объект может предоставлять
    несколько интерфейсов, которые не накладываются друг на
    друга. Пример:
    "org.freedesktop.NetworkManager.Device.Wireless".
  - **Метод (Method)** — некоторое действие, которое может производить
    объект в данной программе на данном интерфейсе. Аналогично
    функциям с языке Си, метод может вернуть набор некоторых
    данных, код ошибки, либо может вообще не возвращать данных
    (void).
  - **Сигнал (Signal)** — некоторое сообщение, которое распространяется
    среди всех программ, подписанных на этот сигнал на этом интерфейсе
    данного объекта данной программы. Сигналы могут содержать набор
    данных.
  - **Сообщение (Message)** — каждая передача данных на шине
    представляется в виде сообщений. Они могут быть 4х
    типов: вызовы методов, сигналы, результаты методов, ошибки.
  - **Прокси-объект (Proxy object)** — объект языка программирования
    (C++, Python, Java), вызовы методов которого проецируются на вызовы
    методов D-Bus. \[2\]

Отметим также, что архитектура сервис-объект-интерфейс свойственна
многим компонентым системам типа COM, XPCOM, CORBA и т.д.

## Принципы работы

### Соединение с шиной

После соединения с шиной приложение, которое создаёт на шине новый
сервис, получает уникальный идентификатор для соединения в рамках
шины. Также любой D-Bus сервис может присвоить себе дополнительное,
более понятное имя (например, ru.org.linux.trollhouse). Этот
механизм схож с созданием символьных ссылок на файловой системе
- у вас может быть один реальный файл и несколько ссылок на него с
разными именами. Имя сервиса, автоматически создаваемое для него
системой D-Bus, имеет вид ":X.Y", где X и Y - компоненты, определяющие
уникальность данного сервиса, обычно числа, например ":1.5".

Следующим шагом при инициализации приложения является установка
фильтров, указывающих, какие сообщения приложением принимаются.
Фильтровать можно по типу сообщения (signal, method_call), отправителю
(sender), интерфейсу, пути к объекту и имени метода.

Затем задаются обработчики сигналов и методов (будет рассмотрено ниже).

Пример на Tcl:

    :
    package require dbus 0.7

    proc processMethodCall {info thing} {
        puts "Method call: $info"
        puts "Argument: $thing"

        return "$thing ne nuzhno!"
    }

    set conn [ dbus::connect session ]
    puts "Connection ID: $conn"

    dbus name session "ru.org.linux.trollhouse"

    dbus filter session add -interface ru.org.linux.Troll -member SaySomething -path /troll
    dbus register session /troll processMethodCall

    vwait forever

Тестирование:

    dbus-send --session --dest=ru.org.linux.trollhouse --type=method_call --print-reply /troll ru.org.linux.Troll.SaySomething string:KDE

### Рассылка сигналов

Сигнал с точки зрения шины представляет собой сообщение с
destination=NULL, т.е. которое надо доставить всем (с учётом фильтров,
разумеется).

Пример:

    package require dbus 0.7

    dbus connect session
    dbus name session ru.org.linux.trollhouse
    dbus signal session \
        -signature si \
        /troll ru.org.linux.troll.Eat NeedSugar \
        "I want more sugar!" 2

Получение сигналов от HAL при появлении/извлечении устройства:

    package require dbus 0.7

    proc processDeviceAddRemove {info udi} {
        if {[ dict get $info member ] == "DeviceAdded"} {
            puts "Device added: $udi"
        } else {
            puts "Device removed: $udi"
        }
    }

    dbus connect system
    foreach member {DeviceAdded DeviceRemoved} {
        dbus filter system add \
            -sender org.freedesktop.Hal \
            -interface org.freedesktop.Hal.Manager \
            -member $member \
            -type signal \
            -path /org/freedesktop/Hal/Manager
    }

    dbus register system \
        /org/freedesktop/Hal/Manager processDeviceAddRemove

    vwait forever

### Удалённый вызов методов

Вызов метода состоит из отправки двух сообщений:

  - Сообщение с типом method_call и аргументами функции.
  - Сообщение с типом method_result или error, содержащее результат
    исполнения метода. В случае, если у исходного сообщения
    выставлен флаг noreply, ответ не посылается. При этом у
    ответного сообщения поле replyserial устанавливается в значение
    поля serial исходного сообщения, чтобы вызывающий смог определить,
    какой ответ к какому вызову метода относится.

Исходя из вышеописанного, существует два способа удалённого вызова
методов через D-Bus: синхроннный (вызывающий процесс блокируется
до прихода ответа) и асинхронный (ответ обрабатывается по мере
получения).

Пример синхронного вызова:

    package require dbus 0.7

    dbus connect session
    dbus call session \
        -dest ru.org.linux.trollhouse \
        -signature s \
        /troll ru.org.linux.Troll SaySomething \
        "Gnome"

Пример асинхронного вызова:

    package require dbus 0.7

    proc resultHandler {info result} {
        if {[ dict get $info messagetype ] == "error"} {
            puts "Unable to speak to troll: $result"
        } else {
            puts "Troll says: $result"
        }
    }

    dbus connect session
    dbus call session \
        -dest ru.org.linux.trollhouse \
        -handler resultHandler \
        -signature s \
        /troll ru.org.linux.Troll SaySomething \
        "Gnome"

    vwait forever

Пример асинхронной обработки вызова метода:

    package require dbus 0.7

    proc processMethodCall {info thing} {
        puts "Method call: $info"
        puts "Argument: $thing"

        set bus session
        set dest [ dict get $info sender ]
        set serial [ dict get $info serial ]
        if [ string match -nocase $thing unix ] {
            # возврат ошибки
            set action [ list dbus error $bus $dest $serial "UNIX rulez!" ]
        } else {
            # возврат результата
            set action [ list dbus return $bus $dest $serial "$thing is not needed!" ]
        }

        # специально вставим задержку, чтобы показать неблокирующий режим работы
        after 10000 $action
    }

    set conn [ dbus::connect session ]
    puts "Connection ID: $conn"

    dbus name session "ru.org.linux.trollhouse"

    dbus filter session add -interface ru.org.linux.Troll -member SaySomething -path /troll
    dbus register session -async \
        /troll processMethodCall

    vwait forever

### Стандартные интерфейсы

Существует несколько предопределенных интерфейсов, использование которых
может быть полезно многим программам\[3\]\[4\].

#### Интроспекция

Почти все библиотеки для работы с dbus позволяют автоматически
генерировать методы для интроспекции (кроме низкоуровневых
вроде libdbus или dbus-tcl).

Для того, чтобы получить информацию о методах, сигналах и свойствах,
надо вызвать метод Introspect из интерфейса
org.freedesktop.DBus.Introspectable. Результат отдаётся в виде строки,
которая представляет собой XML документ. Этот документ содержит
корневой элемент типа node. Под ним располагаются элементы типов
interface (содержащие описание реализуемых интерфейсов; имя интерфейса
хранится в атрибуте name) и элементы типа node, указывающие, что в
дереве объектов есть объекты ниже данного (атрибут name показывает,
что надо дописать к пути, чтобы туда достучаться).

Описание интерфейса содержит элементы типа method (имя задаётся в
атрибуте name), signal и property. Внутри каждого из этих
элементов находится произвольное количество элементов типа arg,
обозначающих аргументы функции или сигнала.

Атрибуты: name --- имя параметра (может отсутствовать), direction ---
направление (in --- аргументы метода, out --- возвращаемое значение;
только для методов), type --- тип элемента (например: i, s, as...).

Пример:

    $ qdbus --system org.freedesktop.Hal / org.freedesktop.DBus.Introspectable.Introspect
    <!DOCTYPE node PUBLIC "-//freedesktop//DTD D-BUS Object Introspection 1.0//EN"
    "http://www.freedesktop.org/standards/dbus/1.0/introspect.dtd">
    <node>
      <interface name="org.freedesktop.DBus.Introspectable">
        <method name="Introspect">
          <arg name="data" direction="out" type="s"/>
        </method>
      </interface>
      <node name="org"/>
    </node>

<table style="vertical-align:middle;padding:8px;margin:10px 50px 10px 50px;border:1px solid yellow;">

<tr>

<td>

<b>TODO</b>

</td>

<td>

пример с интроспекцией

</td>

</tr>

</table>

#### Свойства

Для работы со свойствами существует интерфейс
org.freedesktop.DBus.Properties, у которого определены 3 метода: Get(in
STRING interface_name, in STRING property_name, out VARIANT value),
Set(in STRING interface_name, in STRING property_name, in VARIANT
value) и GetAll(in STRING interface_name, out DICT\<STRING,VARIANT\>
props);

<table style="vertical-align:middle;padding:8px;margin:10px 50px 10px 50px;border:1px solid yellow;">

<tr>

<td>

<b>TODO</b>

</td>

<td>

пример

</td>

</tr>

</table>

#### org.freedesktop.DBus.Ping

<table style="vertical-align:middle;padding:8px;margin:10px 50px 10px 50px;border:1px solid yellow;">

<tr>

<td>

<b>TODO</b>

</td>

<td>

описать

</td>

</tr>

</table>

## Ссылки

  - [D-Bus tutorial](http://dbus.freedesktop.org/doc/dbus-tutorial.html)
  - [D-Bus python
    tutorial](http://dbus.freedesktop.org/doc/dbus-python/doc/tutorial.html)
  - [D-Bus bindings for Tcl](http://dbus-tcl.sourceforge.net/) содержит
    также высокоуровневый биндинг dbus-intf.
  - [Высокоуровневая привязка к объектной системе Snit для
    Tcl](http://code.google.com/p/dbus-snit)
  - [Привязки к Qt4](http://doc.trolltech.com/4.4/intro-to-dbus.html)
  - [D-Bus
    Specification](http://dbus.freedesktop.org/doc/dbus-specification.html)

## Отладочные утилиты

  - dbus-send\[5\] и dbus-listen из комплекта dbus
  - qdbus из пакета libqt-dbus. В отличие от dbus-send использует
    интроспекцию, поэтому не требует указания типов аргументов и
    позволяет просматривать объекты на шине.
  - qdbusviewer из пакета qt4-dev-tools. Отображает иерархию объектов на
    шине графически, показывает методы, сигналы и свойства (properties).
    Позволяет с ними работать.
  - d-feet из одноимённого пакета. Позиционируется как d-bus отладчик.

## Сноски

По мере исправления сноски необходимо удалять.

<references />

[Category:Системное ПО](Category:Системное_ПО "wikilink")
[Category:Программирование](Category:Программирование "wikilink")

1.  Это надо полностью переписать.
2.  Добавить свойства
3.  ЩИТО?
4.  [Второе предложение в первом
    абзаце](http://dbus.freedesktop.org/doc/dbus-specification.html#standard-interfaces)
5.  Поддерживает не все типы данных. По этому с её помощью нельзя к
    примеру открыть всплывающую подсказку
    [Подробности→](http://ubuntuforums.org/showthread.php?t=959664)