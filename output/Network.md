Вопросы, связанные с X-сервером и запуском программ по сети смотрите в
разделе [X-сервер](X-сервер)

Знаменитый iptables-tutorial. [Русский
перевод](http://gazette.linux.ru.net/rus/articles/index-iptables-tutorial.html).
[Оригинал на
английском](http://iptables-tutorial.frozentux.net/iptables-tutorial.html).
Перевод весьма хорош. Лучше всего скачать себе в виде архива, тогда вы
получите рабочие примеры скриптов.

Также весьма полезно почитать статьи в [разделе переводов и оригинальных
статей Linux Gazette](http://gazette.linux.ru.net/rus/index.html).

## Как поменять в KPPP тоновый набор на импульсный?

Либо перед номером телефона пишем букву 'P' (латинская\!), либо заходим
в свойства kppp, команды и пишем вместо ATDT команду ATDP.

## Что делать, если при запуске wvdial не от рута программа пишет device busy?

Добавьте пользователя в группу dialout

## А можно пример конфига для wvdial?

    /etc/wvdial.conf:

    [Dialer Defaults]
    Modem = /dev/ttyS0
    Baud = 112500
    Init = ATZ
    Init2 = AT S11=50
    Dial Command = ATDP
    Phone = 9955555
    Username = int1104221
    Password = LyQSptAy
    New PPPD = yes
    #Init3 = ATM0

    [Dialer phone2]
    Phone = 9955556

    [Dialer phone3]
    Phone = 1055555

    [Dialer phone4]
    Phone=9519024

    [Dialer phone5]
    Phone=9613332

    [Dialer norm]
    Init3 = ATM1

## Как включить сжатие ip-заголовков microsoft (MPPE/MPPC)?

В виндах этот алгоритм называется MPPC сжатием (компрессией), Обычно в
дистрибутивах linux поддержи по умолчанию нет,

Для того, чтобы ее включить, понадобится скачать патч для ядра (вроде бы
есть на [1](http://gaute.vetsj.com/?tag=mppc)).

Далее нужно пропатчить pppd (патч там же, где патч под ядро). А в
/etc/modprobe.conf или /etc/modules.conf нужно добавить пару строчек,
чтобы получилось следующее:

    alias ppp-compress-18   ppp_mppe_mppc
    alias char-major-108    ppp_generic
    alias ppp-compress-21   bsd_comp
    alias ppp-compress-24   ppp_deflate
    alias ppp-compress-26   ppp_deflate
    alias tty-ldisc-3       ppp_async
    alias tty-ldisc-14      ppp_synctty

## Как подключить GPRS? Как настроить соединение через bluetooth-адаптер?

На русском - [2](http://www.opennet.ru/base/net/gprs_linux.txt.html),
[3](http://dvtl.pisem.net/gprs_moto.html)

Если мобильный телефон определяется как AT-совместимый модем, идем на
[4](http://www.easyconnect.linuxuser.hu/) и качаем программу 'GPRS Easy
Connect', собираем и запускаем

В некоторых дистрибутивах rfcomm выдает ошибку "can't connect RFCOMM
socket: Connection refused". Это может быть связано с неправильной
передачей pin'а телефону. Проверьте чтобы **/usr/bin/bluepin**
давал ответ в виде "PIN:you_pin". В случае ответа "ERR" замените в
hcid.conf строчку "pin_helper /usr/bin/bluepin" на "pin_helper
/usr/bin/pin-helper"

    user@host:~$ cat /usr/bin/pin-helper
    #!/bin/sh
    echo -n "PIN:"
    cat /etc/bluetooth/pin

За совет благодарим Lumi. В новых версиях bluez может не сработать,
поэтому лучше предварительно спарить телефон и компьютер.

Хорошие статьи: [5](http://www.linuxrsp.ru/artic/Linux-Bluetooth.html)
[6](http://www.linuxrsp.ru/artic/Linux-OBEX-GPRS.html)
[7](http://www.linuxrsp.ru/artic/Siemens-S55.html)
[8](http://turtiainen.dna.fi/GPRS-HOWTO) - для лучшего понимания
принципов работы.

Еще о bluetooth:
[<http://www.linuxforum.ru/index.php?s=fbfbecaba95453655a567110503627a7&showtopic=3745&pid=62649&st=0>](http://www.smartweb.ru/news.php?act=c&page=27&cat=2&scat=6)

## Как настроить СТРИМ?

Как любое ADSL соединение:

  - Устанавливаем пакет rp-pppoe
  - Запускаем adsl-setup.


```
      USER NAME: имя для инета
      INTERFACE: свой интерфейс (для Zyxel OMNI ADSL USB EE это nas0)
      DNS: пропускаем (должно работать)
      PASSWORD: от соединения в инет.
```

Для некоего питерского провайдера, предоставляющего сходные услуги, есть
подробная статья.

О том, как настроить те или иные ADSL USB-модемы, читайте в разделе
[Hardware](Hardware).

## Подключаюсь к интернету по модему. Соединение проходит, но ни один браузер не может ничего открыть. Что делать?

Возможные проблемы:

  - Нет маршрута по умолчанию
  - Ядро не воспринимает динамические адреса
  - Не указаны адреса DNS-серверов провайдера
  - Проблема с файрволлом

Решить все, кроме последнего пункта, довольно просто:

  - Попробуйте выполнить в консоли


```
      root@linux~:# route add default ppp0
```

Если все заработало, то поставьте себе эту строчку в
**/etc/ppp/if-up.local** На всякий случай проверьте скрипт if-up - в
старых дистрибутивах иногда отсутствовал вызов if-up.local. Строчка
там должна быть примерно такой:

```

      [ -x /etc/ppp/ip-up.local ] && /etc/ppp/ip-up.local "$@"
```

  - Попробуйте выполнить в консоли


```
      echo "1" > /proc/sys/net/ipv4/ip_dynaddr
```

Разрешение на динамические адреса должно присутствовать в обязательном
порядке (если вы не подключаетесь к частной линии и адрес не задан у
вас статически). Поэтому поставьте себе эту строчку в
**/etc/rc.d/rc.local** (или аналогичный файл, стартующий при запуске
системы). Для дистрибутивов fedora core 2 или 3 вместо этого
добавьте в **/etc/sysctl.conf** строчку **net.ipv4.ip_dynaddr =
1**.

  - Для указания адресов DNS-серверов нужно зайти в файл
    **/etc/resolv.conf** и прописать их туда примерно в такой форме
    (пример для МТУ-Интел):


```
      nameserver      195.34.32.11
      nameserver      195.34.32.10
```

  - Проблемы с firewall'ом нужно решать в частном порядке, методом
    внимательного прочтения iptables-tutorial. Отмечу только, что
    обычно стандартные брандмауэры обычно позволяют выходить в
    интернет.

## После того как вхожу в инет машина начинает дико тормозить. Все открывается долго. Что делать?

Убрать в своей программе дозвона получение имя хоста от провайдера. В
том же kppp, например, есть галочка на эту тему. О параметрах pppd
читайте **man pppd**

## Какой выбрать сервер VPN?

Вот [здесь](http://www.opennet.ru/opennews/art.shtml?num=4499) лежат
ссылки на статью, сравнивающую функциональность различных
VPN-серверов, а также ведущая на небольшие статьи со ссылками на
полную документацию к большинству основных дистрибутивов. Не везде
документация свежая, но в сочетании с манами по ней можно легко
настроить нужное вам решение.

## Как настроить подключение к Microsoft VPN Server?

Прочитайте об этом -
[9](http://www.nixp.ru/cgi-bin/go.pl?q=articles;a=ms_vpn_client),
[10](http://fine.kalinovka.net/?q=node/4)

Как настроить pptpclient в Mandrake Linux смотрите здесь
[11](http://pptpclient.sourceforge.net/howto-mandrake-10.phtml).

## Как настроить шифрование в pptp, если сам он уже работает? (VPN)

Найти и скачать пакет kernel_ppp_mppe
[12](http://sourceforge.net/project/showfiles.php?group_id=44827&package_id=120221&release_id=244156)
или патчик в ядро [13](http://www.polbox.com/h/hs001/) (тогда его
придется пересобирать).

Узнать, поддерживает ли твой pppd mschap v2. Наверно нет. Наверно стоит
сходить на [mcmcc.bat.ru](http://mcmcc.bat.ru) - там они вроде
патченные. Или опять же скачать патчик.

Добавим в /etc/modules.conf строчку

    alias ppp-compress-18 ppp_mppe_mppc

А в /etc/ppp/options.pptpd запишем

    lock
    require-mppe
    require-mschap-v2

Остальное в меру своей испорченности - у меня (jackill), например,
прописан плагин к радиусу.

Далее заходим в наши винды, во вкладку соединения, там выставляем
"дополнительные" параметры вместо всяческих типовых и отмечаем
mschap v2. Галочку с mschap (который первой версии) убираем, иначе не
будет подключаться.

## Как разрешить клиентам из LAN соединяться по VPN через NAT?

Для поддержки passthrough VPN-соединений нужно установить patch-o-matic
расширение -
[14](http://netfilter.org/projects/patch-o-matic/index.html)

## Есть компьютер, подключенный к интернету и в нем есть сетевая карта. Есть вторая машина. Как бы на нее интернет подключить?

Я так понимаю, сделать это надо быстро, но неизвестно как и что.
Отлично. Нам нужно найти rc.firewall-strong. Проще всего искать в
[15](http://www.google.ru) по словосочетанию "FWVER=0.80" (эта версия
уже держится около полугода без особых модификаций).

Я его нашел, например, здесь:
[16](http://www.aerospacesoftware.com/firewall-howto.html).

Версия 0.75 содержится в Linux IP Masquerading How-To.

Качаем этот файл, называем его rc.firewall и кладем в **/etc/rc.d/**,

Далее пользуясь комментариями исправляем (интерфейсы - вдруг ppp0
потребуется вместо eth0, включаем/выключаем DHCP, смотрим надо ли
включать ip_dynaddr), запускаем.

Затем на второй машине прописываем машину с интернетом как gateway (см.
как это делается в вашем дистрибутиве), в **/etc/resolv.conf**
прописываете адреса DNS-серверов провайдера. Все.

Если машина под управлением MS Windows, выставляете все настройки в
"Настройке сети".

## Как мне защитить свою домашнюю машину под управлением ОС Linux от атак из интернета и из локальной сети?

Это очень обширная тема, в силу чего я не очень хотел бы помещать такие
вопросы тут, но вот кратенько об основном:

    #!/bin/bash

    for table in INPUT OUTPUT FORWARD ; do
            iptables -P $table DROP
    done

    iptables -A INPUT -p TCP -m state --state ESTABLISHED,RELATED -j ACCEPT
    iptables -A INPUT -p UDP -m state --state ESTABLISHED,RELATED -j ACCEPT
    iptables -A OUTPUT -p ALL -j ACCEPT

Этими командами мы разрешаем выход в мир с нашей машины и запрещаем все
остальное. Обычно, такой скрипт называется **/etc/rc.d/rc.firewall** и
запускается из стартовых скриптов (например, **/etc/rc.d/rc.local**).

Для более подробного ознакомления с возможностями iptables следует
почитать хотя-бы это -
[17](http://www.opennet.ru/docs/RUS/iptables/index.html).

После того, как вы прочитали хотя бы порядок прохождения пакетов в linux
при использовании iptables, можете попробовать графические интерфейсы
для настройки iptables.

И уж точно не повредит прочтение документов из раздела [Linux Security
FAQ](Linux_Security_FAQ).

Если хочется быстро настроить файрволл для себя, чтобы разрешались на
вход любые исходящие пакеты, можно найти уже готовые правила.

## В какой цепочке iptables считать трафик, в каком порядке проходятся цепочки, что будет раньше, NAT или шейпинг и т.д.?

После внимательного изучения Kernel Packet Traveling Diagram эти вопросы
должны отпасть - [18](http://www.docum.org/docum.org/kptd/)

Так же неплохо посмотреть iptables-tutorial (ссылка в начале этого
раздела).

## Как ограничить скорость прохождения трафика?

В общем и целом, это довольно сложная тема и если применять
соответствующие утилиты без относительного понимания, лучше
точно не будет.

Ограничивать скорость прохождения трафика можно и с помощью iptables
(при использовании [Patch-o-matic
расширений](http://www.netfilter.org/patch-o-matic/),
например).

Или с помощью шейперов.

Кроме того, в Squid имеется механизм ограничения скорости web-трафика.

Очень большие возможностыи дает использование утилит пакета iproute2 -
[19](http://developer.osdl.org/dev/iproute2/).

Документация по использованию iproute2:

  - Linux Advanced Routing & Traffic Control HOWTO -
    [20](http://lartc.org/howto/)
  - Docum.org - [21](http://www.docum.org/docum.org/)
  - Повесть о Linux и управлении трафиком -
    [22](http://gazette.linux.ru.net/rus/articles/taleLinuxTC.html)
  - OpenNet iproute2 info -
    [23](http://www.opennet.ru/keywords/iproute2.html)

Кроме того, для упрощения использования возможностей по управлению
трафиком, были написаны множество скриптов, самые популярные из
которых - CBQ.init ( [24](http://sourceforge.net/projects/cbqinit/))
и HTB.init ( [25](http://uf.kadan.cz/htb/)). Хотя они и обладают рядом
недостатков (например, нет возможности обслуживать динамические
интерфейсы - ppp+), все же они довольно удобны для решения
простых задач.

Кстати, не используйте CBQ.init - дисциплина обработки очереди HTB была
разработана на замену CBQ, является более функциональной и удобной в
использовании.

## Как соединить MS Windows (hyperterminal) и linux по телефонной линии?

Воспользоваться программой minicom.

Если под MS Windows мы делаем HyperTerminal-\>Call-\>телефон, то в
minicom просто набираем ATDP телефон (для импульсных линий) или ATDT
телефон (для тоновых).

Протокол подключения называется zmodem.

## Как соединить два компьютера через com-порты?

Выполняем на первой машине:

    pppd /dev/ttyS0 115200 local nocrtscts nocdtrcts noxonxoff 192.168.0.1:192.168.0.2 netmask 255.255.255.252 noauth

на второй машине:

    pppd /dev/ttyS0 115200 local nocrtscts nocdtrcts noxonxoff 192.168.0.2:192.168.0.1 netmask 255.255.255.252 noauth

Получаем локалку со всеми вытекающими.

## Как соединить два компьютера через lpt-порты?

Нужно использовать PLIP (в ядре должна быть поддержка).

Как именно - читайте PLIP-HOWTO.

## Как соединить два компьютера через USB-порты?

Купить PC-to-PC USB кабель. А вот какая у вас связь получится и что
дальше - не в курсе. (Кто может описать подробнее, отправьте,
пожалуйста, ответ).

## Как соединить два компьютера по Bluetooth?

Первый вариант:

На одном:

    pand -s -M

На втором:

    pand -c bt_адрес_первого

Второй вариант:

На одном:

    pand --listen --role NAP --persist
    sdptool add NAP
    hciconfig hci0 piscan

На втором:

    pand --search

## Как узнать, подключен ли сетевой кабель?

    user@linux:~$ /sbin/mii-tool
    eth0: negotiated 100baseTx-FD, link ok

где link ok - подключен.

Другой вариант:

    $ ifconfig
    eth0      Link encap:Ethernet  HWaddr 00:01:6c:de:6a:cb
              inet addr:192.168.xxx.yyy  Bcast:192.168.xxx.255  Mask:255.255.255.0
              UP BROADCAST MULTICAST RUNNING MTU:1500  Metric:1
              RX packets:0 errors:0 dropped:0 overruns:0 frame:0
              TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
              collisions:0 txqueuelen:1000
              RX bytes:0 (0.0 B)  TX bytes:0 (0.0 B)
              Interrupt:23

Где есть RUNNING --- там кабель подключен.

Ещё вариант (нужен):

    # ethtool eth0 | grep Link
    Link detected: yes

Ещё один:

    $ ifplugstatus eth0
    eth0: link beat detected

## Как поменять MAC-адрес сетевой карты?

С разными способами можно ознакомиться здесь:
<http://mydebianblog.blogspot.com/2007/02/blog-post_24.html> . Также
есть программа macchanger, которая умеет генерировать и заменять
адрес по заданным критериям.

## Как добавить несколько IP-адресов на интерфейс?

    root@localhost:~# ifconfig eth0:any_alias 1.2.3.4
    root@localhost:~# ip addr add dev eth0 1.2.3.4

## eth0 смотрит, к примеру, в сеть 10.0.0.0, eth1 смотрит в 192.168.0.0. Как создать маршрут, чтобы машины из сети 192\* видели сеть 10\*?

На маршрутизаторе

    root@linuxrouter:~# echo "1" > /proc/sys/net/ipv4/ip_forward

*(Чтобы последняя команда выполнялась на старте, в fedora core и debian
нужно выставить 1 на одноименном параметре в /etc/sysctl.conf, а в
redhat 7.3 - 9 в /etc/sysconfig/network)*

На хосте

    root@linuxhost:~# ip route add 10.0.0.0/8 via 192.168.0.1

где 192.168.0.1 - адрес интерфейса маршрутизатора в сети 192.168.0.0

## Как жестко привязать интерфейсы eth0 и eth1 к оборудованию

Если после каждой загрузки интерфейсы меняются местами, нужно добавить
правило udev (создать файл /etc/udev/rules.d/net.rules), которое еще
на этапе загрузки компьютера жестко определит имя интерфейса. Первый
вариант:

    # net rule for eth0 (00:23:54:7C:26:12)
    SUBSYSTEM=="net", ACTION=="add", ATTR{address}=="00:23:54:7C:26:12", KERNEL=="eth*", NAME="eth0"
    # net rule for eth1 (00:10:22:FD:C5:0C)
    SUBSYSTEM=="net", ACTION=="add", ATTR{address}=="00:10:22:FD:C5:0C", KERNEL=="eth*", NAME="eth1"

Здесь в поле ATTR{address} необходимо записать реальный mac-адрес
сетевой карты.

Если это правило не срабатывает, используем следующее правило:

    SUBSYSTEM=="net", ENV{ID_VENDOR_ID}=="0x10ec", NAME="eth0"
    SUBSYSTEM=="net", ENV{ID_VENDOR_ID}=="0x10b7", NAME="eth1"

Узнать, что писать в ENV{ID_VENDOR_ID}, можно из содержимого файла
`/sys/class/net/ethX/device/vendor`, где ethX - текущее имя данного
интерфейса.

## Как сделать прозрачный мост?

Воспользоваться статьёй Linux Bridge
[26](http://xgu.ru/wiki/Linux_Bridge)

Если после настройки моста компьютеры из объединённых сетей не видят
друг-друга (не пингуются), то возможно причина в настройках
iptables.

Пути решения:

  - отключить слежение iptables за мостом:


  -
    echo "0" \> /proc/sys/net/bridge/bridge-nf-call-iptables


  - разрешить FORWARD'инг трафика через мост следующим правилом:


  -
    iptables -I FORWARD -m physdev --physdev-is-bridged -j ACCEPT

## После прописывания правил iptables выдается ошибка, последнее сообщение которой Try \`iptables -h' or 'iptables --help' for more information. Что делать?

Убьем сразу кучу зайцев и рассмотрим несколько вопросов в одном.

Запомните: **Скорее всего у вас неправильно написана команда. Вы где-то
что-то забыли написать.**

Теперь успокойтесь, глубоко вздохните, выдохните и ищите ошибку. Самые
распространенные:

  - **multiple -d flags not allowed** Скорее всего вместо --dport
    написано -dport. Очень часто забывают второй "-" или пишут
    dports/sports.
  - При написании правил вида **iptables -t nat -A POSTROUTING -o $EXTIF
    -j SNAT --to-source $EXTIP** часто забывают прописать либо **-t
    nat**, либо **-A POSTROUTING**. Часто неверно пишут POSTROUTING.
  - При указании нескольких портов с помощью multiple часто забывают
    указать **-m state**
  - При использовании переменных в строчках, иногда путают их названия
    или забывают задать изначально
  - Нужный модуль не загружен (проблемы с ftp, irc, командой multiple,
    маркировкой пакетов идут отсюда).

## Как просматривать ресурсы сети Windows (сетевое окружение) из консоли? Как его смонтировать?

Для начала нужно посмотреть ресурсы при помощи команды **smbclient -L
winmachine**. На данный момент ресурсы с названием более 16 символов
поддерживаются некорректно, т.е. все символы в названии после 16-го
отсекаются.

Чтобы смонтировать нужный ресурс, дайте примерно такую команду:

    smbmount //winmachine/movie /home/jackill/mnt/movie -o  iocharset=koi8-r,rw,codepage=cp866,umask=0

Для ресурсов с паролем дайте команду:

    smbmount //winmachine/work /home/jackill/mnt/work -o  iocharset=koi8-r,rw,umask=0,codepage=cp866,username=_username_,password=_password_

Также существуют GUI-утилиты и специальный скрипт, искать ресурсы
которыми и монтировать значительно удобнее.

## Как настроить Samba? Как сделать, чтобы в ресурсах Samba были видны русские буквы?

Настройка Samba

Для простой одноранговой сети пример файла настроек Samba
(/etc/samba/smb.conf) можно скачать
[здесь](http://fine.kalinovka.net/articles/hobby/lorFAQ/files/smb.conf.tgz).

Из примера должно быть понятно, как сделать доступными свои ресурсы.

Русский язык

Чтобы русские названия файлов отображались правильно, ядро должна быть
собрано с определенными параметрами:

    File systems -> Network File Systems


    <M> SMB file system support (to mount Windows shares etc.)
    [*]   Use a default NLS
    Default Remote NLS Option: "cp866"

Иными словами:

    CONFIG_SMB_FS=m
    CONFIG_SMB_NLS_DEFAULT=y
    CONFIG_SMB_NLS_REMOTE="cp866"

(проверялось на ядрах серии 2.4.х и 2.6.х)

Далее собираем Samba с поддержкой locales. Для третьей Samba может
потребоваться установить libiconv. При этом вторая Samba соберет
себе файлы кодовых страниц в /$PREFIX/share/samba/codepages.

Затем в конфигурационном файле Samba (для RedHat-подобных систем это
/etc/samba/smb.conf) пишем

для Samba 2.2.x:

    character set = KOI8-R
    client code page = 866

для Samba 3.x.x:

    unix charset = KOI8-R
    display charset = KOI8-R
    dos charset = 866

Если локаль не koi8-r, а utf8, то в вышеприведенных примерах поменяйте
koi8-r на utf8.

Как просмотреть чужие ресурсы Теперь, чтобы получить доступ к ресурсам
сети MS Windows, ресурс можно "подмонтировать" (прямо как NFS, только
опций больше):

    root@linux# mount -t smbfs -o fmask=666,dmask=777,rw,iocharset=koi8-r,codepage=cp866 //winmachine/share /mnt/smb/share

или

    user@linux# smbmount //winmachine/share /mnt/smb/share -o iocharset=koi8-r,rw,codepage=cp866,username=your_name,password=your_password

Если ваша локаль не koi8-r, то поставьте в iocharset свою локаль.

**Примечание:** Настоятельно рекомендуется использовать не smbfs, а
cifs, дабы избежать проблем с правами и передачей файлов длиной более
2Gb. Читайте следующий вопрос.

## Не получается смонтировать ресурсы сервера MS Windows 2003 в домене. Что делать?

Сначала собираем в ядре поддержку Client VFS (CIFS):

    CONFIG_CIFS=m
    # CONFIG_CIFS_STATS is not set
    CONFIG_CIFS_XATTR=y
    CONFIG_CIFS_POSIX=y

Затем нужно обновить Samba до версии минимум 3.0.7.

Монтируются ресурсы командой **mount.cifs** (поставляется вместе с
Samba)

## Обновил Samba/дистрибутив и больше не могу получить доступ к ресурсам Windows, что делать? Как передать через Samba файл с размером больше 2Gb?

А это вам маячок, подсказывающий, что пора использовать cifs. Команда
для монтирования будет выглядеть примерно так:

    mount -t cifs //homeserver/temp /home/jackill/mnt -o  iocharset=koi8-r,rw,codepage=cp866,umask=0

Что писать в ядре написано здесь

## Как сделать в Samba каталог, чтобы в него могли все заходить без аутентификации?

Удивительно, но этот вопрос зачастил. В файле конфигурации Samba
(smb.conf) в разделе \[global\] пишем строчку:

    security = share

Пример описания папки:

    [media]
    comment = working
    path = /mnt/dos/media
    browseable = yes
    writable = yes
    create mode = 0644
    directory mode = 0777
    guest ok = yes
    public = yes

## Как управлять iptables с помощью Samba + ldap?

[27](http://www.invask.ru/linuxtools/ldap-iptables/) (битая)

## Хочу синхронизировать время с серверами времени в интернете. Как это сделать?

Довольно просто - берем пакет ntp и устанавливаем себе на машину. В
большинстве дистрибутивов (Fedora Core, Mandrake, openSUSE) он
поставляется сразу.

Список серверов можно взять здесь:

  - [28](http://ntp.isc.org/bin/view/Servers/StratumOneTimeServers)
  - [29](http://ntp.isc.org/bin/view/Servers/StratumTwoTimeServers)

## На 64-битной платформе время в 32-битном дистрибутиве идет слишком быстро. Что делать?

Скорее всего мать на чипсете ATI (RS480-M). Лечится добавлением
параметра к ядру clock=tsc (см. документацию к ядру -
Documents/kernel-parameters.txt).

## Как настроить бездисковую загрузку?

Тема довольно объемная. Чтобы не играть в глухой телефон, лучше
прочитайте вот это:

[30](http://ltsp.org/)

[31](http://thinstation.sourceforge.net/wiki/index.php/ThIndex)

[32](http://www.wtware.ru/)

## Как передать звук с одной машины на другую по сети, используя aRts?

Первый пользователь должен запустить artsd так:

    user1@localhost$ artsd -u -p xxxxx

Второй пользователь должен отключить свой уже запущенный atrsd (см.
настройки звука в Control Center) и дать такую команду:

    user2@localhost$ export ARTS_SERVER=hostname:xxxxx

xxxxx- порт arts сервера.

Теперь user2 сможет использовать сервер запущенный пользователем user1
(и не только он, но и любой, кто даст такую же команду).

Для уменьшения времени отклика сервера можно запустить artswrapper.

## Как передать трафик ssh/telnet/irc и т.п., если я сижу за HTTP-прокси (ssh over proxy)?

Если есть доступ к серверу, для передачи SSH можно настроить Apache -
[33](http://dag.wieers.com/howto/ssh-http-tunneling/).

Или воспользоваться утилитой Proxytunnel -
[34](http://proxytunnel.sourceforge.net/intro.php)

Существует также временно приостановленный проект туннеля des-proxy -
[35](http://desproxy.sourceforge.net/)

Также, существует простенькая программа connect-proxy. Настройка ssh c
использованием connect-proxy:

`~$ cat .ssh/config`
`Host *`
` ProxyCommand connect-proxy -H 192.168.0.1:8080 %h %p`

Теперь все соединения(Host \*) будут устанавливаться через прокси
192.168.0.1:8080.

## Как настроить WAP-сервер?

Настроить обычный apache, предварительно задавшись соответствующим
вопросом на любом поисковике.

Если коротко: нужно ему написать в конфиге

    DirectoryIndex что_угодно index.wml

И положить это самый wml.

## У меня несколько подключений к интернету. Как мне сделать чтобы они корректно работали одновременно?

[36](http://www.opennet.ru/docs/RUS/LARTC/x348.html) ([оригинальная
статья на английском
языке](http://lartc.org/lartc.html#LARTC.RPDB.MULTIPLE-LINKS))

## Что делать, если в программе отсутствует поддержка прокси?

Почти в каждом дистрибутиве есть консольная утилита proxychains, с
помощью которой программы могут работать через прокси серверы. К
ней есть графическая оболочка -
[37](http://sourceforge.net/projects/proxychainsgui/)

## 3g Мегафон/Билайн Коннект при помощи pppd

### Конфиги для Мегафон

    # cat /etc/ppp/peers/megafon-gprs
    lcp-echo-failure 0
    lcp-echo-interval 0

    connect /etc/ppp/peers/megafon-gprs-connect-chat
    # debug      # раскомментировать для отладки

    /dev/ttyUSB0 # Имя девайса, последняя цифра может быть 1, 2 или 3
    921600       # Скорость обмена с модемом. Еще варианты: 115200, 230400 и т.д.

    local
    noipdefault
    ipcp-accept-local
    defaultroute
    #replacedefaultroute
    #usepeerdns  # Использовать предлагаемые ДНС (отключено, использую свои)

    novj
    nobsdcomp
    novjccomp
    nopcomp
    noaccomp
    nodetach

    noauth

    # При плохом качестве канала полезно уменьшить размер пакетов, иногда и более жестоко
    #mtu 800
    #mru 800

    # cat /etc/ppp/peers/megafon-gprs-connect-chat
    exec chat -vS \
        '' \rAT \
        TIMEOUT 12 \
        OK ATH \
        OK ATE1 \
        OK 'AT+CGDCONT=1,"IP","internet"' \
        OK ATD*99***1# \
        TIMEOUT 22 \
        SAY "\nWaitinf for connect.....\n" \
        CONNECT "" \
        SAY "\nConnected!\n"

Этот файл нужно сделать исполняемым.

    # chmod +x /etc/ppp/peers/megafon-gprs-connect-chat

### Конфиги для Билайн

Почти то же самое (файлы можно скопировать и поправить). Изменена точка
доступа и добавлена аутентификация. Имена файлов пусть будут с
префиксом **beeline**

  - Точка доступа хранится в файле
    **/etc/ppp/peers/beeline-gprs-connect-chat**. У кого как, но с
    симкой, которая была куплена мною, **APN с сайта Билайна
    internet.beeline.ru не работает**. Рабочая строка:


```
     OK 'AT+CGDCONT=1,"IP","home.beeline.ru"' \
```

  - В файл **/etc/ppp/peers/beeline-gprs** добавлена строка:


    user beeline

  - В файл **/etc/ppp/pap-secrets** (на меге не нужна аутентификация)
    добавлена строка:


    beeline *       beeline

### Запуск

От рута:

`# pppd file /etc/ppp/peers/megafon-gprs`

В общем то, работает\! Однако, соединение иногда слетает. После добавил
строчки в **/etc/inittab** чтоб коннект поддерживался "на плаву"

``` bash
# GPRS-подключения. Нужное РАСКОММЕНТИРОВАТЬ
#mg:2345:respawn:/usr/sbin/pppd file /etc/ppp/peers/megafon-gprs >/var/log/megafon-gprs
#bg:2345:respawn:/usr/sbin/pppd file /etc/ppp/peers/beeline-gprs >/var/log/beeline-gprs
```

Ну и перечитать конфиг:

`# /sbin/init q`

При такой схеме все же приходится один-два раза в месяц вмешиваться
ручками - в том случае, когда требуется заново зарегистрироваться
на БС опсоса. Скорей всего, это проблемы местной базовой станции.

Для быстрого переподключения достаточно сделать:

`# killall pppd`

## Сетевые контейнеры (network namespace)

Network namespace это изолированный сетевой стек без запуска виртуальной
машины

[Описание
функционала](http://wiki.sirmax.noname.com.ua/index.php/NetworkNamespaces)
[Примеры
использования](http://net-labs.in/2014/04/06/%D0%BF%D1%80%D0%B8%D0%BC%D0%B5%D1%80%D1%8B-%D0%BF%D1%80%D0%B8%D0%BC%D0%B5%D0%BD%D0%B5%D0%BD%D0%B8%D1%8F-linux-network-namespaces-netns/)

## Маршрутизация multicast

Используются smcroute, pimd, mrouted, xorp, igmpproxy

[Маршрутизация multicast в
Linux](http://net-labs.in/2014/04/19/%D0%BC%D0%B0%D1%80%D1%88%D1%80%D1%83%D1%82%D0%B8%D0%B7%D0%B0%D1%86%D0%B8%D1%8F-multicast-%D0%B2-linux/)

