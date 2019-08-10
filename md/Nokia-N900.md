## Developer Tools

Чтобы превратить N900 в настоящую линукс-машину, на которой можно
собирать софт, необходимо установить инструменты разработки. Для
этого подключите репозитории в **/etc/apt/sources.list**

    # TOOLS
    deb http://repository.maemo.org/ fremantle/tools free
    deb-src http://repository.maemo.org/ fremantle/tools free
    
    # SDK
    deb http://repository.maemo.org/ fremantle/sdk free
    deb-src http://repository.maemo.org/ fremantle/sdk free

Не забудьте выполнить **apt-get update**

Затем установите все необходимые пакеты

    apt-get install bzip2 cpio cpp dpkg-dev g++ g++-4.2 gcc \
    libc6-dev libstdc++6-4.2-dev libstdc++6-4.2-dbg patch \ 
    perl perl-modules autoconf automake1.9 libtool flex bison \ 
    gdb diffutils-gnu

В случае каких-то проблем выполните **apt-get clean** и повторите
установку, все должно пройти гладко.

## PC Suite Mode

### USB IP

Подключив Nokia N900 по USB дата-кабелю, вы можете поднять сетевой
интерфейс **usb0** для прямого доступа к телефону.

    # lsusb | grep N900
    Bus 003 Device 007: ID 0421:01c8 Nokia Mobile Phones N900 (PC-Suite Mode)

Для этого выполните команду на телефоне:

    ifconfig usb0 192.168.2.1 up

И выполните команду на компьютере:

    ifconfig usb0 192.168.2.2 up

А теперь можете "постучаться" на телефон:

    # ssh 192.168.2.1
    The authenticity of host '192.168.2.1 (192.168.2.1)' can't be established.
    RSA key fingerprint is af:db:d4:24:3b:03:4c:ea:3b:b1:59:cb:47:46:f0:d1.
    Are you sure you want to continue connecting (yes/no)? yes
    Warning: Permanently added '192.168.2.1' (RSA) to the list of known hosts.
    root@192.168.2.1's password: 
    
    
    BusyBox v1.10.2 (Debian 3:1.10.2.legal-1osso30+0m5) built-in shell (ash)
    Enter 'help' for a list of built-in commands.
    
    spoofing_n900:~# 

Алсо, при включенном по дата-кабелю компьютеры видят N900 как флешку -
вы можете этим воспользоваться, чтобы загружаться в LiveUSB образы.

## Tips & Tricks

Чтобы установить браузер в вертикальное положение (портретный режим),
запустите его и нажмите Ctrl + Shift + o
