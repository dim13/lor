## Установка

<http://www.lissyara.su/articles/openbsd/syntony/installation_openbsd_4.2/>

С 2008 года инсталлятор практически не изменился, наверное он вообще
никогда не менялся :)

Настоятельно рекомендуется ознакомиться с
[afterboot(8)](http://www.openbsd.org/cgi-bin/man.cgi?query=afterboot) и
[FAQ](http://www.openbsd.org/faq/)\!

### Удаленная установка на сервер без KVM

В случае, если на удалённом сервере доступна только (linux) rescue
system, установка всё ещё возможна с использованием qemu и vnc (и
проброской 5900 порта на локальную машину).

    local$ slogin -L 5900:localhost:5900 root@remote.host
    remote# qemu-system-x86_64 -hda /dev/sda -cdrom install52.iso -vnc localhost:0

    local$ vncviewer localhost

Единственно, на что стоит обратить внимание: авторазбивка диска
[disklabel(8)](http://www.openbsd.org/cgi-bin/man.cgi?query=disklabel)
использует размер RAM для swap (256M + ramsize) и /var (4G + 2 \*
ramsize) (здесь подразумеваются диски рамером более 300G). При
использование qemu эти значения будут несоответствовать
актуальному размеру RAM и их надо подкорректировать вручную.

Так же при использование qemu-установки не всегда правильно определяется
количество процессоров. В таком случае следует добавить bsd.mp в выборе
пакетов установки и заменить SP kernel на MP kernel после установки:

    # cd /mnt && ln bsd bsd.sp && mv bsd.mp bsd

### Установка на [softraid(4)](http://www.openbsd.org/cgi-bin/man.cgi?query=softraid)

Начиная с 5.1 возможна загрузка напрямую с raid: [OpenBSD 5.1
installation on
softraid(4)](http://blog.cochard.me/2012/03/openbsd-51-installation-on-sofraid4.html),
включая установку на шифрованную партицию.

## Русификация

### Консоль

Если при установке выбрать раскладку ru, то инсталлятор сразу пропишет
ее в /etc/kdbtype, как и советуют почти во всех хауту. Правда у меня
([JB](http://www.linux.org.ru/whois.jsp?nick=JB)) из за этого были
проблемы с xdm и хотя я не уверен в их причине, тем не менее я
пошел другим путем - /etc/kbdtype удалил, а в /etc/wsconsctl.conf
добавил вот это:

    keyboard.encoding=ru # use different keyboard encoding

там же можно указать и клавишу для переключения:

    keyboard.map+="keycode 157 = Mode_Lock" # Right Ctrl

далее в /etc/rc.local (что бы консоль могла отображать русские буквы):

    wsfontload /usr/share/misc/pcvtfonts/koi8-r-8x16
    for cons in 1 2 3 4; do wsconscfg -dF $cons; wsconscfg -t 80x25bf $cons; done
    unset cons

и если в wsconsctl.conf не был указан keyboard.map, то сюда же надо
добавить **wsconsctl -w keyboard.map+="keycode 157=Mode_Lock"**. А
можно добавить и туда и туда

### X11

Иксы настраиваются как и везде добавлением в \~/.xsession:

    setxkbmap -layout us,ru -option grp:ctrl_shift_toggle -option grp_led:scroll

Еще надо добавить в .profile, .xinitrc или .xsession (в случае логина
через xdm) вот это:

    export LC_COLLATE=C
    export LC_CTYPE=ru_RU.KOI8-R
    export LC_MONETARY=C
    export LC_NUMERIC=C
    export LC_TIME=C

## Сеть

Каждый интерфейс настраивается в отдельном конфиге
[/etc/hostname.$interface](http://www.openbsd.org/cgi-bin/man.cgi?query=hostname.if).
Например:

    $ cat /etc/hostname.ath0
    inet 10.0.0.38 255.255.255.0 10.0.0.255

    $ cat /etc/hostname.athn0
    nwid SSID_точки_доступа
    wpakey пассфраза
    dhcp

Шлюз указывается в
[/etc/mygate](http://www.openbsd.org/cgi-bin/man.cgi?query=mygate) (в
случае с dhcp это не нужно):

    $ cat /etc/mygate
    10.0.0.1

Чтобы применить все изменения на лету без перезагрузки достаточно
запустить **sh /etc/netstart имя_интерфейса**

### IPv6 как клиент

Автоконфигурация:

    # cat <<EOF>> /etc/hostname.ath0
    rtsol
    EOF
    # ed /etc/sysctl.conf <<EOF
    /net.inet6.icmp6.rediraccept/s/^#//
    w
    EOF

Пример ручной настройки:

    # cat <<EOF>> /etc/hostname.re0
    inet6 alias 2001:db8:beef::1::1 64
    EOF
    # cat <<EOF>> /etc/mygate
    fe80::1%re0
    EOF

Проверка:

    $ ping6 ipv6.google.com

**Пометка** для тех, кто хостится у Hetzner: у них слегка странная
конфигурация — gateway находится в другой подсети (/59). Поэтому
prefixlen в hostname.if надо указывать не 64, а 59. (*quick and dirty
fix\!*) Это касается и других BSD и Linux. →
<http://marc.info/?t=130001685700003&r=1&w=2>

### Автоматическое переключение между LAN и WLAN

В случае, если LAN и WLAN находятся в одной подсети, что бывает часто,
можно организовать незаметное переключение туда и обратно используя
[trunk(4)](http://www.openbsd.org/cgi-bin/man.cgi?query=trunk)
интерфейс.

Например: проводной интерфейс re0 (без IP)

    $ cat /etc/hostname.re0
    up

беспроводной интерфейс ral0 (тоже без IP, только установки линка)

    $ cat /etc/hostname.ral0
    nwid myssid
    wpakey secret
    up

trunk0 (dhcp и по желанию IPv6)

    $ cat /etc/hostname.trunk0
    trunkproto failover
    trunkport re0
    trunkport ral0
    dhcp
    rtsol autoconfprivacy

Порядок имеет значение: re0 — это главный интерфейс и используется по
умолчанию, ral0 — это запасной интерфейс (wifi), на который будет
переключён трафик, если первый интерфейс потеряет линк (выдернули
кабель). Если опять воткнуть кабель, то трафик будет опять идти через
главный интерфейс. Переключение происходит незаметно. Что собственно
говоря и требовалось.

## Порты и пакеты

По умолчанию порты отсутствуют, поэтому скачиваем их вручную и
распаковываем:

    # ftp ftp://ftp.openbsd.org/pub/OpenBSD/`uname -r`/ports.tar.gz
    # cd /usr
    # tar zxf /root/ports.tar.gz

Установка и удаление аналогичны портам FreeBSD, единственное что стоит
запомнить это скрипт для проверки устаревших установленных портов:

    # cd /usr/ports/infrastructure/build/
    # ./out-of-date
    Collecting installed packages: ok
    Collecting port versions: ok
    Collecting port signatures: ok
    Outdated ports:

    net/wget # 1.12p1 -> 1.13.4

    # cd /usr/ports/net/wget
    # make update

Путь к репозиторию с пакетами устанавливается в
[/etc/pkg.conf](http://www.openbsd.org/cgi-bin/man.cgi?pkg.conf):

    # echo "installpath = ftp://ftp.openbsd.org/pub/OpenBSD/`uname -r`/packages/`uname -m`" > /etc/pkg.conf

он так же может быть альтернативно указан в переменной окружения
PKG_PATH:

    # export PKG_PATH=ftp://ftp.openbsd.org/pub/OpenBSD/`uname -r`/packages/`uname -m`/

Поиск пакетов (pkg.conf должен быть установлен):

    pkg_info -Q package_name

или

    cd /usr/ports && make search name=package_name

Пакеты ставятся так:

    # pkg_add -i package_name

Ключ **-i** - интерактивная установка, в случае неопределённости
(например у пакета несколько
[flavors](http://www.openbsd.org/cgi-bin/man.cgi?query=ports#FLAVORS))
будет запрошена помощь пользователя.

Для обновления есть ключ -u:

    # pkg_add -u package_name

Либо можно обновить все пакеты сразу:

    # pkg_add -iu

Проверить (не устанавливать) обновления:

    # pkg_add -sux

Эту команду можно добавить в /etc/daily.local что бы уведомления об
изменениях приходили с ежедневной почтой.

Переустановка всех пакетов (или для клонирования установленных пакетов
на другой хост):

    # pkg_info -m > packages.txt
    # pkg_delete -X
    # pkg_add -z -l packages.txt

Удалить все неиспользуемые зависимости (аналог *apt-get autoremove*):

    pkg_delete -a

Есть так же пакет с псевдографическим интерфейсом к пакетам:
[**pkg_mgr**](http://ports.su/sysutils/pkg_mgr) и два Web-интерфейса:
<http://www.openports.se/> и <http://ports.su/> для поиска и просмотра
пакетов ([подробнее о сайтах
пакетов](http://www.linux.org.ru/news/bsd/9132500)).

## Конфигурация

### rc.d

Начиная с 4.9 все сторонние сервисы находятся в
[/etc/rc.d](http://www.openbsd.org/cgi-bin/man.cgi?query=rc.d). Для их
запуска добавить в /etc/rc.conf.local параметр pkg_scripts:

    pkg_scripts="dbus_daemon"

### Редактирование /etc/rc.conf

Разработчики *настоятельно* рекомендуют избегать правки непосредственно
[/etc/rc.conf](http://www.openbsd.org/cgi-bin/man.cgi?query=rc.conf).
Все изменения вносятся в
[/etc/rc.conf.local](http://www.openbsd.org/cgi-bin/man.cgi?query=rc.conf.local),
переменные в котором имеют приоритет над переменными в
[/etc/rc.conf](http://www.openbsd.org/cgi-bin/man.cgi?query=rc.conf).

Это имеет смысл, поскольку /etc/rc.conf будет перезаписан при обновлении
системы. Аналогично, изменения
[/etc/rc](http://www.openbsd.org/cgi-bin/man.cgi?query=rc) следует
выносить в
[/etc/rc.local](http://www.openbsd.org/cgi-bin/man.cgi?query=rc.local).
Это же касается файлов
[/etc/daily](http://www.openbsd.org/cgi-bin/man.cgi?query=daily),
[/etc/weekly](http://www.openbsd.org/cgi-bin/man.cgi?query=weekly) и
[/etc/monthly](http://www.openbsd.org/cgi-bin/man.cgi?query=mounthly) —
имена файлов с локальными модификациями получают суффикс **.local**

## Пересборка ядра

Качаем тарбол с исходниками и разворачиваем куда надо:

    # ftp -C ftp://mirror.corbina.net/pub/OpenBSD/`uname -r`/sys.tar.gz
    # cd /usr/src
    # tar zxvf /root/sys.tar.gz

Конфиг лежит в /usr/src/sys/arch/\`uname -m\`/conf. Можно либо
отредактировать GENERIC (тут еще нужно обратить внимание на
то, что конфиг для многоядерных или многопроцессорных систем
GENERIC.MP состоит по сути из трех строчек - двух параметров отвечающих
за SMP и вставки содержимого GENERIC), либо скопировать его под другим
именем и править уже его:

    # cd /usr/src/sys/arch/`uname -m`/conf
    # cp GENERIC PUFFY
    # vi PUFFY
    # config PUFFY
    # cd ../compile/PUFFY
    # make clean && make depend && make -j3
    # make install
    # reboot

Для чего это нужно? Например чтобы [включить поддержку
bluetooth](http://wiki.openbsd.ru/Использование_Bluetooth.html)

## Обновление системы

[Описание процесса обновления системы до ветки
-stable](http://openbsd.ru/docs/steps/stable.html)

[Обновление до -current,
English](http://blather.michaelwlucas.com/archives/543)

Скрипт для обновления дерева исходного кода и портов:

    #!/bin/sh
    VERSION=$(uname -r | tr . _)
    MODULES='src ports xenocara'
    CVSROOT=anoncvs@anoncvs.openbsd.org:/cvs
    export ${CVSROOT}
    cd /usr
    cvs -d ${CVSROOT} up -C -rOPENBSD_${VERSION} -Pd ${MODULES}

Обновление конфигурационных файлов в /etc происходит интерактивно
семи-автоматически при помощи
[sysmerge(8)](http://www.openbsd.org/cgi-bin/man.cgi?query=sysmerge):

    sysmerge -s etcXX.tgz -x xetcXX.tgz

Обновление бинарных пакетов:

    # cat /etc/pkg.conf
    installpath = ftp://ftp.eu.openbsd.org/pub/OpenBSD/5.4/packages/amd64/
    # pkg_add -ui

## Звук

Все манипуляции с микшером делаются с помощью
[mixerctl(1)](http://www.openbsd.org/cgi-bin/man.cgi?query=mixerctl),
например изменение громкости:

    $ mixerctl outputs.master=130,130
    outputs.master: 126,126 -> 130,130

В портах так же имеется интерфейс на curses: **cmixer**.

На ноутбуках при подключении наушников звук может продолжать идти из
колонок. Лечится это так:

    $ mixerctl -a | grep output | grep slaves
    outputs.master.slaves=dac-0:1,line,hp
    $ mixerctl outputs.master.slaves=dac-0:1,line
    $ mixerctl outputs.master.mute=on

без второй команды не получится отдельно выключать колонки и наушники, а
третья собственно выключает вывод звука на колонки оставляя только
наушники. Насчет бескостыльности этого метода я не уверен

Все это можно сохранить в
[/etc/mixerctl.conf](http://www.openbsd.org/cgi-bin/man.cgi?query=mixerctl.conf):

    $ cat /etc/mixerctl.conf
    outputs.master=130,130
    outputs.master.slaves=dac-0:1,line
    outputs.master.mute=on

### Проигрывание музыки через сеть

Для этого следует на машине с колонками запустить
[sndiod(1)](http://www.openbsd.org/cgi-bin/man.cgi?query=sndiod) с
параметром **-L-**:

    # cat <<EOF>> /etc/rc.conf.local
    sndiod_flags="-L-"
    EOF
    # /etc/rc.d/sndiod restart

и выставить на машине, откуда будет идти звук, соответствующую
переменную окружения:

    $ export AUDIODEVICE=snd@hostname_or_IP/0

Работает со всеми программами.

## ACPI, энергосбережение

Для управления частотой процессора достаточно добавить запуск apmd в
/etc/rc.conf.local:

    cat <<EOF>> /etc/rc.conf.local
    apmd_flags="-A"
    EOF

**-A** - автоматическое управление частотой в зависимости от нагрузки,
**-C** - максимальное энергосбережение

В ждущий режим (suspend-to-ram) система отправляется командой **apm -z**
или просто **zzz**. Спящий режим (suspend-to-disk) поддерживается с
версии 5.2 командой **ZZZ**.

Для того, чтобы ноутбук автоматически засыпал при закрытии крышки,
следует включить в
[/etc/sysctl.conf](http://www.openbsd.org/cgi-bin/man.cgi?query=sysctl.conf)
*lidsuspend*:

    # ed /etc/sysctl.conf <<EOF
    /machdep.lidsuspend/s/^#//
    w
    EOF

## Автомонтирование

HAL в OpenBSD отсутствует и udisks тоже, поэтому для автомонтирования
хотя бы тех же флешек придется сооружать давно забытые костыли. В
первую очередь надо поставить hotplug-diskmount и разрешить запуск
hotplugd:

    # pkg_add hotplug-diskmount
    hotplug-diskmount-0.5p1: ok
    # echo "hotplugd_flags=" >> /etc/rc.conf.local
    # /etc/rc.d/hotplugd start
    hotplugd(ok)
    # /usr/local/libexec/hotplug-diskmount init

Потом создать
[/etc/hotplug/attach](http://www.openbsd.org/cgi-bin/man.cgi?query=hotplugd)
с таким содержимым:

    #!/bin/sh

    DEVCLASS=$1
    DEVNAME=$2

    case $DEVCLASS in
    2)
            /usr/local/libexec/hotplug-diskmount attach -u artem -m 700 "$DEVNAME"
            ;;
    esac

разумеется вместо **artem** подставить свое имя пользователя. И не
забыть сделать файл исполняемым - **chmod +x /etc/hotplug/attach**

Теперь можно вставлять флешку и проверять. По умолчанию все монтируется
в /vol/МЕТКА_ТОМА

## Разное

  - Чтобы научить ls выводить все в цвете нужен пакет fileutils и
    соответствующие алиас - **alias ls="gls --color"**. Готовый
    файл раскраски для gls можно взять
    [тут](http://openbsd.ru/files/etc/DIR_COLORS). Можно также
    воспользоваться пакетом colorls.
  - Умолчальный тип тепминала **vt220**, чтобы консоль стала цветной
    следует установить **export TERM=wsvt25**
  - Чтобы кнопка Delete в xterm работала именно как Del, а не Backspace
    в .Xdefaults понадобится такая строчка - **XTerm\*deleteIsDEL:
    false**
  - xsetbg почему-то не дружит с композитингом - при запуске xcompmgr
    теряется wallpaper. С feh же все нормально

<!-- end list -->

  - В наличии [имеется](http://www.linux.org.ru/news/bsd/9023141) сайт
    быстрого поиска по свежим исходникам kernel + non-gnu userland —
    [BXR.SU/OpenBSD/](http://BXR.SU/OpenBSD/)
  - Для удобства RTFM
    [существует](http://www.linux.org.ru/news/bsd/9165554)
    сайт коротких адресов системных руководств
    [„mdoc.su“](http://mdoc.su/)

<!-- end list -->

  - [Новости с пометкой „openbsd“ на
    ЛОРе](http://www.linux.org.ru/tag/openbsd?section=1)