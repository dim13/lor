### Установка через другой GNU/Linux (продвинутая)

#### Пролог

Если вам, по каким-то причинам, не угодила установка через официальный
установщик (например, экономия дисков и т.д.), то можно установить
Debian через любой другой GNU/Linux с прямым доступом к диску. Например,
через LiveCD (Knoppix, Ubuntu и т.д.).
**Внимание\! Инструкция рассчитана на BIOS-машины с x86-архитектурами
(x86-32 (i\*86) и x86-64 (amd64, EMT64)).**

#### Начальная установка

  - Для начала, разметим файловую систему. Для этого подойдёт программа
    GParted или CFDisk. Для установки GParted в Debian-подобной системе
    (Knoppix, Ubuntu и прочие) введите в консоли (от root или sudo):

`apt install gparted`

  - Для установки CFDisk в Debian-подобной системе (Knoppix, Ubuntu и
    прочие) введите в консоли (от root или sudo):

`apt install util-linux`

В дальнейшем, будет предполагаться, что основной диск размечен в ext4, и
создан swap.

  - Теперь надо примонтировать созданный диск куда-нибудь (от root или
    sudo):

`mkdir /media/hda0`
`mount -t ext4 /dev/sda6 /media/hda0`

Вместо «sda6» укажите выбранный диск (в GParted), а вместо «ext4»
укажите выбранную файловую систему.

  - Придётся поставить debootstrap:

1\. Для Debian-подобных (от root или sudo):

`apt install debootstrap`

2\. Для всех остальных (от root или sudo):

`cd /`
`wget -O debootstrap.deb `<http://http.debian.net/debian/pool/main/d/debootstrap/debootstrap_1.0.64_all.deb>
`ar -x debootstrap.deb`
`tar -xf data.tar.xz`
`rm -f debootstrap.deb control.tar.gz data.tar.xz`

  - Ну и сама установка:

`/usr/sbin/debootstrap --include=dselect --arch $ARCH $DISTRO /media/hda0 `<http://http.debian.net/debian>

Тут, вместо «$ARCH» укажите архитектуру ОС (для x86 это i386, а для
x86-64 это amd64), а вместо «$DISTRO» кодовое имя нужного релиза Debian
(например, jessie.

  - Чтобы система увидела устройства, введите в консоли:

`mount -o rbind /dev /media/hda0/dev`
`mount -o rbind /sys /media/hda0/sys`

  - Теперь можно в новую систему и войти (от root или sudo):

`env LANG=C chroot /media/hda0 /bin/bash`

  - Создайте символические ссылки для CD-ROM (DVD-ROM):

`mkdir /media/cdrom0`
`ln -s /media/cdrom0 /media/cdrom`

#### Настройки времени и дисков

  - Для вывода диалога настройки времени:

`dpkg-reconfigure tzdata`

  - Создайте файл /etc/fstab (nano /etc/fstab) с содержанием,
    образованным по данному шаблону:

`# /etc/fstab: static file system information.`
`#`
`# Use 'blkid' to print the universally unique identifier for a`
`# device; this may be used with UUID= as a more robust way to name devices`
`# that works even if disks are added and removed. See fstab(5).`
`#`
`# `<file system>` `<mount point>`   `<type>`  `<options>`       `<dump>`  `<pass>
`# /boot was on /dev/sda2 during installation `
`/dev/sda2       /boot           ext2    defaults        0       2`
`# / was on /dev/sda3 during installation `
`/dev/sda3       /               ext4    defaults        0       1`
`# /home was on /dev/sda4 during installation `
`/dev/sda4       /home           ext4    defaults        0       2`
`# swap was on /dev/sda5 during installation`
`/dev/sda5       none            swap    sw              0       0`
`/dev/sr0        /media/cdrom0   udf,iso9660 user,noauto 0       0`
`tmpfs           /tmp            tmpfs   defaults        0       0`
`/tmp            /var/tmp        none    bind            0       0`

Поправьте имена устройств, ненужные диски удалите, нужные добавьте.
**Примечание:** Вместо имён устройств можно (и даже лучше, ибо имена
могут измениться по разным причинам) использовать UUID или метку
устройства. Например, так:
"UUID=63a233ee-b9c2-4cef-a6ab-34f2f756cf1c" или "LABEL=Root" (без
кавычек).

  - Теперь смонтируем всё это дело:

`mount -t proc /proc /proc`
`mount -a`

#### Настройка зеркала репозиториев и выбор ПО

  - Чтобы получить доступ к большему количеству программ из
    репозиториев, укажите в /etc/apt/sources.list (можно
    выбрать другое зеркало, например, ftp.debian.org/debian или
    mirror.yandex.ru/debian):

`deb `<http://http.debian.net/debian>` $DISTRO main contrib non-free`
`# deb-src `<http://http.debian.net/debian>` $DISTRO main`

`deb `<http://security.debian.org>` $DISTRO/updates main contrib non-free`
`# deb-src `<http://security.debian.org>` $DISTRO/updates main`

`deb `<http://http.debian.net/debian>` $DISTRO-proposed-updates main contrib non-free`
`# deb-src `<http://http.debian.net/debian>` $DISTRO-proposed-updates main`

`# deb `<http://http.debian.net/debian>` $DISTRO-backports main contrib non-free`
`# deb-src `<http://http.debian.net/debian>` $DISTRO-backports main`

Где вместо $DISTRO кодовое имя дистрибутива.

Персонально для Debian Sid:

`deb `<http://http.debian.net/debian>` sid main contrib non-free`
`# deb-src `<http://http.debian.net/debian>` sid main contrib`

  - На x86-64 (amd64) системах лучше предварительно включить поддержку
    установки i686-пакетов:

`dpkg --add-architecture i386`

  - После чего обновить список пакетов

`apt update`

  - Если вы хотите участвовать в опросе популярности пакетов, то
    установите пакет popularity-contest:

`aptitude install popularity-contest`

Для выхода из опроса:

`dpkg-reconfigure popularity-contest`

  - Для установки некоторых стандартных пакетов можно воспользоваться
    утилитой tasksel:

`dselect update`
`tasksel`

Установка aptitude:

  - aptitude - мощная утилита для работы с пакетами. В случае, если вы
    не хотите её использовать, пропустите этот пункт, а в дальнейших
    указаниях заменяйте "aptitude" на "apt"

`apt install aptitude`

#### Установка языковых параметров

  - Скачайте, а потом настройте локали:

`aptitude install locales console-setup`
`dpkg-reconfigure locales`
`dpkg-reconfigure console-setup`

При настройке console-setup лучше выбрать как шрифт VGA.

#### Сетевые установки

  - Желательно выбрать имя своего ПК, ибо сейчас оно совпадает с именем
    ПК основной системы:

`echo '$PCName' > /etc/hostname`

Где вместо $PCName имя ПК.

  - Также, создайте файл /etc/hosts с содержанием (nano /etc/hosts):

`127.0.0.1       localhost`
`127.0.1.1       $PCName`

`# The following lines are desirable for IPv6 capable hosts`
`::1             localhost ip6-localhost ip6-loopback`
`ff02::1         ip6-allnodes`
`ff02::2         ip6-allrouters`

Где вместо $PCName имя ПК, а вместо localhost — имя домена.

#### Настройка пользователей

  - Теперь пора создать себе пользователя:

`adduser $USER`

Где вместо $USER имя пользователя.

  - Также рекомендую поставить пароль на root:

`passwd`

#### Установка ядра и загрузчика

  - Ну вот и предфинальное действие — установка ядра Linux. Для
    установки стандартного ядра достаточно:

`aptitude install linux-base linux-image-$ARCH linux-headers-$ARCH`

Где $ARCH может быть amd64, 686-pae или 486 (на x86-64 системах лучше
amd64, а на x86-32 — 686-pae).

  - Для правильного функционирования системы могут понадобиться
    различные прошивки. Большинство прошивок можно найти в
    метапакете firmware-linux:

`aptitude install firmware-linux`

Учтите, некоторые прошивки вынесены в отдельные пакеты. Например, для
работы сетевых карт на чипсете от Ra-Link (многие продукты D-Link)
нужен пакет firmware-ralink:

`aptitude install firmware-ralink`

Также продукты выпускались с чипсетами RealTek:

`aptitude install firmware-realtek`

  - Ну и на конец установка загрузчика:

`aptitude install grub-pc`

**Примечание:** Вместо GRUB2 также можно установить GRUB1 (пакет
grub-legacy), Syslinux (пакет syslinux) или LiLO (пакет lilo).

  - На этом всё. Поздравляю с установкой Debian на Ваш компьютер.

### Установка дополнительных программ

#### Установка sudo

Для совершения административных задач без переключения в root
используется sudo (SuperUserDO).

  - Установка (от root):

`aptitude install sudo`

  - Далее, надо добавить нужного пользователя в список
    «администраторов». Для этого добавьте пользователя
    в группу «sudo» от root:

`usermod -a -G sudo $USER_NAME`

где $USER_NAME — имя пользователя;

#### Установка X Window System

Пока основной графического стека GNU/Linux является проект X.org.
Для установки X.org вместе со всеми графическими драйверами,
использующими Mesa:

`aptitude install xorg`

#### Установка дисплейного менеджера

Для графического входа в графическую оболочку существуют дисплейные
менеджеры.
Чтобы установить универсальный дисплейный менеджер LightDM с
Gtk-фронтэндом:

`aptitude install lightdm lightdm-gtk-greeter`

Или с KDE-фронтэндом:

`aptitude install lightdm lightdm-kde-greeter`

Также существуют дисплейные менеджеры GDM (от среды GNOME), MDM (от
проекта Mint), KDM (от среды LDE) и пр.

#### Установка рабочей среды

Если вы не установили GNOME, то можно спокойно установить среду сейчас.
Для этого можно воспользоваться программой aptitude (или apt) от root:

  - для KDE:

`aptitude install kde-standard`

  - для GNOME:

`aptitude install gnome`

  - для Xfce:

`aptitude install xfce4`

  - для LXDE:

`aptitude install lxde`

**Примечание:** Для управления питанием нужно установить его поддержку
(от root):

`aptitude install acpid acpi-support`

#### Установка звукового сервера PulseAudio

Cуществует звуковой сервер PulseAudio, решающий (а по иному мнению —
создающий) проблемы со звуком, который можно поставить (совместно с
ALSA) так (от root или sudo):

`aptitude install pulseaudio libasound2-plugins`

#### Настройка сети

Для настройки сетевого соединения существуют специализированные утилиты
как NetworkManager и Wicd.

  - Чтобы установить nm-applet (графический интерфейс к NetworkManager)
    (от root или sudo):

`aptitude install network-manager-gnome`

  - Чтобы установить Wicd-Gtk (графический интерфейс к Wicd) (от root
    или sudo):

`aptitude install wicd-gtk`

#### Установка дополнительных кодеков

По разным причинам, в Debian отсутствуют некоторые кодеки. Чтобы их
установить:

  - Создайте /etc/apt/sources.list.d/deb-multimedia.list с содержимым:

`deb `<http://mirror.yandex.ru/debian-multimedia>` $DISTRO main non-free`
`# deb-src `<http://mirror.yandex.ru/debian-multimedia>` $DISTRO main`

Где $DISTRO — имя Вашего дистрибутива.

  - Выполните (от root или sudo):

`apt update`
`aptitude install deb-multimedia-keyring w32codecs libdvdcss2`
`apt update`

Для amd64:

`apt update`
`aptitude install deb-multimedia-keyring w64codecs libdvdcss2`
`apt update`

#### Установка другой системы инициализации

В Debian по умолчанию используется systemd, который некоторые считают
переусложнённым посланием из ада. Но в репозиториях также есть
системы инициализации OpenRC и SysVinit, которые посланием из ада
не считаются.

  - Чтобы установить OpenRC вместо systemd (от root или sudo):

`aptitude install openrc`

Настройки OpenRC хранятся в /etc/rc.conf.

  - Чтобы установить SysVinit вместо systemd (от root или sudo):

`aptitude install sysvinit-core`

Также для ускорения загрузки, если нет systemd, можно поставить пакет
readahead-fedora (от root или sudo):

`aptitude install readahead-fedora`

#### Курсор в терминале

Для появления курсора в виртуальном терминале установите пакет gpm (от
root или sudo):

`aptitude install gpm`

#### Установка видеодрайвера

Для установки видеодрайвера посетите:

  - Для видеокарт nVIDIA:
    [проприетарный](http://wiki.debian.org/NvidiaGraphicsDrivers#non-freedrivers)
    ([на русском](http://wiki.debian.org/ru/NvidiaProprietary)
    (возможно, устарело)),
    [свободный](http://wiki.debian.org/NvidiaGraphicsDrivers#freedrivers);
  - Для видеокарт AMD:
    [проприетарный](http://wiki.debian.org/ATIProprietary)
    ([на русском](http://wiki.debian.org/ru/ATIProprietary) (возможно,
    устарело)), [свободный](http://wiki.debian.org/AtiHowTo)([на
    русском](http://wiki.debian.org/ru/AtiHowTo) (возможно,
    устарело)).

#### Автомонтирование устройств

Если у Вас не монтируются устройства автоматически, то это можно
исправить установкой пакета pmount (от root):

`aptitude install pmount`

#### Отключение установки "рекомендаций" в aptitude

Если Вы желаете не устанавливать "рекомендации" пакетов по умолчанию с
aptitude, то введите следующую команду (от root):

`echo -e '`<APT::Install-Recommends>` "false";\naptitude::Recommends-Important "false";' | tee -a /etc/apt/apt.conf`

#### Подключение кэширования символьных таблиц X11

Для ускорения запуска программ в X-сервере можно включить кэширование
символьных таблиц, для этого введите следующую команду:

`mkdir ~/.compose-cache`

### Лицензия

Содержимое этой страницы предоставляется на условиях следующей лицензии:
[CC Attribution-Share
Alike 3.0](http://creativecommons.org/licenses/by-sa/3.0/).