### Установка через другой GNU/Linux (продвинутая)

#### Пролог

Если вам, по каким-то причинам, не угодила установка через официальный
установщик (например, экономия дисков и т.д.), то можно установить
любую редакцию Ubuntu через любой другой GNU/Linux с прямым доступом
к диску. Например, через LiveCD (Knoppix, даже сам Ubuntu и т.д.). Может
пригодиться, как пример, для установки Kubuntu 12.04 через Ubuntu 8.10
или наоборот.  
**Внимание\! Инструкция рассчитана на BIOS-машины с x86-архитектурами
(x86-32 (i\*86) и x86-64 (amd64, EMT64)).**

#### Начальная установка

  - Для начала, разметим файловую систему. Для этого подойдёт программа
    GParted или CFDisk. Для установки GParted в Debian-подобной системе
    (Knoppix, Ubuntu и прочие) введите в консоли (от root или sudo):

`apt install gparted`

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

  - Ну и сама установка (от root или sudo):

`/usr/sbin/debootstrap --include=aptitude,tasksel,dselect,nano,wget --arch $ARCH $DISTRO /media/hda0 `<http://ru.archive.ubuntu.com/ubuntu>

Тут, вместо «$ARCH» укажите архитектуру ОС (для x86 это i386, а для
x86-64 это amd64), а вместо «$DISTRO» кодовое имя нужного релиза Ubuntu
(например, trusty).

  - Чтобы система увидела устройства, введите в консоли:

`mount -o rbind /dev /media/hda0/dev`  
`mount -o rbind /sys /media/hda0/sys`

  - Теперь можно в новую систему и войти (от root):

`env LANG=C chroot /media/hda0 /bin/bash`

  - Создайте символические ссылки для CD-ROM (DVD-ROM):

`mkdir /media/cdrom0`  
`ln -s cdrom0 /media/cdrom`

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

  - Теперь можно и смонтировать:

`mount -t proc /proc /proc`  
`mount -a`

#### Настройка зеркала репозиториев и выбор ПО

  - Чтобы получить доступ к большему количеству программ из
    репозиториев, укажите в /etc/apt/sources.list (можно
    выбрать другое зеркало, например, archive.ubuntu.com/ubuntu или
    mirror.yandex.ru/ubuntu):

`deb `<http://ru.archive.ubuntu.com/ubuntu>` $DISTRO main universe restricted multiverse`  
`# deb-src `<http://ru.archive.ubuntu.com/ubuntu>` $DISTRO main universe`  
  
`deb `<http://security.ubuntu.com/ubuntu>` $DISTRO-security main universe restricted multiverse`  
`# deb-src `<http://security.ubuntu.com/ubuntu>` $DISTRO-security main universe`  
  
`deb `<http://ru.archive.ubuntu.com/ubuntu>` $DISTRO-updates main universe restricted multiverse`  
`# deb-src `<http://ru.archive.ubuntu.com/ubuntu>` $DISTRO-updates main universe`  
  
`deb `<http://ru.archive.ubuntu.com/ubuntu>` $DISTRO-proposed main universe restricted multiverse`  
`# deb-src `<http://ru.archive.ubuntu.com/ubuntu>` $DISTRO-proposed main universe`  
  
`deb `<http://ru.archive.ubuntu.com/ubuntu>` $DISTRO-backports main universe restricted multiverse`  
`# deb-src `<http://ru.archive.ubuntu.com/ubuntu>` $DISTRO-backports main universe`  
  
`deb `<http://archive.canonical.com/ubuntu>` $DISTRO partner`  
  
`deb `<http://extras.ubuntu.com/ubuntu>` $DISTRO main`  
`# deb-src `<http://extras.ubuntu.com/ubuntu>` $DISTRO main`

Где вместо $DISTRO кодовое имя дистрибутива.

  - На x86-64 (amd64) системах лучше предварительно включить поддержку
    установки i686-пакетов:

`dpkg --add-architecture i386`

  - После чего обновить список пакетов

`apt update`

  - Для использования команды add-apt-repository следует установить
    пакет software-properties-common (или
    python-software-properties в Ubuntu \< 12.10):

`aptitude install python-software-properties`

  - Если вы хотите участвовать в опросе популярности пакетов, то
    установите пакет popularity-contest:

`aptitude install popularity-contest`

  - Для выхода из опроса:

`dpkg-reconfigure popularity-contest`

  - Для установки некоторых стандартных пакетов можно воспользоваться
    утилитой tasksel:

`dselect update`  
`tasksel install standard`

  - Можно ещё запустить просто

`tasksel`

и установить что-нибудь из выпавшего списка.

Установка aptitude:

  - aptitude - мощная утилита для работы с пакетами. В случае, если вы
    не хотите её использовать, пропустите этот пункт, а в дальнейших
    указаниях заменяйте "aptitude" на "apt"

`apt install aptitude`

#### Установка языковых параметров

  - Для установки русской локали установите пакет language-pack-ru:

`aptitude install language-pack-ru`

  - Настройте шрифт в консоли:

`dpkg-reconfigure console-setup`

При настройке console-setup лучше выбрать как шрифт VGA.

  - А для применения:

`echo "LANG=\"ru_RU.UTF-8\"" > /etc/default/locale`

  - Для применения русской локали уже сейчас:

`export LANG=ru_RU.UTF-8`

#### Сетевые установки

  - Желательно выбрать имя своего ПК, ибо сейчас оно совпадает с именем
    ПК основной системы:

`echo '$PCName' > /etc/hostname`

Где вместо $PCName, имя ПК.

  - Также, создайте файл /etc/hosts с содержанием (nano /etc/hosts):

`127.0.0.1       localhost`  
`127.0.1.1       $PCName`  
  
`# The following lines are desirable for IPv6 capable hosts`  
`::1             localhost ip6-localhost ip6-loopback`  
`ff02::1         ip6-allnodes`  
`ff02::2         ip6-allrouters`

Где вместо $PCName имя ПК, а вместо localhost — имя домена.

  - Для настройки сети воспользуйтесь поисковиком, ибо здесь
    действительно много писать.

#### Настройка пользователей

  - Теперь пора создать себе пользователя:

`adduser $USER`

Где вместо $USER имя пользователя.

  - Также можно поставить пароль на root:

`passwd`

#### Установка ядра и загрузчика

  - Ну вот и предфинальное действие — установка ядра Linux. Для
    установки стандартного ядра достаточно:

`aptitude install linux-image-generic linux-headers-generic`

**Примечание:** Вместо GRUB2, который устанавливается как рекомендация
linux-image-generic, также можно установить Syslinux (пакет syslinux)
или LiLO (пакет lilo).

  - Для правильного функционирования системы могут понадобиться
    дополнительные прошивки:

`aptitude install linux-firmware-nonfree`

  - На этом всё. Поздравляю с установкой Ubuntu на Ваш компьютер.

### Установка «дополнительных» программ

#### Настройка sudo

Для совершения административных задач без переключения в root
используется sudo (SuperUserDO).

  - Далее, надо добавить нужного пользователя в список
    «администраторов». Для этого добавьте пользователя
    в группу «sudo»:

`usermod -a -G sudo $USER_NAME`

где $USER\_NAME — имя пользователя;

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

Для установки рабочей среды можно воспользоваться программой aptitude
(или apt) от root:

  - для KDE:

`aptitude install kubuntu-desktop`

  - для Unity в \>=11.04 или GNOME в =\<10.10:

`aptitude install ubuntu-desktop`

  - для GNOME Fallback между 11.10 и 13.04):

`aptitude install gnome-panel`

  - для GNOME Shell в \>=11.10):

`aptitude install gnome-shell gnome-themes-standard gnome-tweak-tool`

  - для Xfce:

`aptitude install xubuntu-desktop`

  - для LXDE:

`aptitude install lubuntu-desktop`

#### Настройка сети

Для настройки сетевого соединения существуют специализированные утилиты
как NetworkManager и Wicd.

  - Чтобы установить nm-applet (графический интерфейс к NetworkManager)
    (от root или sudo):

`aptitude install network-manager-gnome`

  - Чтобы установить Wicd-Gtk (графический интерфейс к Wicd) (от root
    или sudo):

`aptitude install wicd-gtk`

#### Курсор в терминале

Для появления курсора в виртуальном терминале установите пакет GPM (от
root или sudo):

`aptitude install gpm`

#### Отключение установки "рекомендаций" в aptitude

Если Вы желаете не устанавливать "рекомендации" пакетов по умолчанию с
aptitude, то введите следующую команду (от root):

`echo -e '`<APT::Install-Recommends>` "false";\naptitude::Recommends-Important "false";' | tee -a /etc/apt/apt.conf`

#### Подключение кэширования символьных таблиц X11

Для ускорения запуска программ в X-сервере можно включить кэширование
символьных таблиц, для этого введите следующую команду:

`mkdir /tmp ~/.compose-cache`

### Лицензия

Содержимое этой страницы предоставляется на условиях следующей лицензии:
[CC Attribution-Share
Alike 3.0](http://creativecommons.org/licenses/by-sa/3.0/).
