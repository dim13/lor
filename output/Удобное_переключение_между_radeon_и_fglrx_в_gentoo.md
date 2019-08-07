# Введение

Для видеокарт AMD существует два драйвера: открытый
[radeon](http://en.gentoo-wiki.com/wiki/Radeon) и проприетарный
[fglrx](http://en.gentoo-wiki.com/wiki/Fglrx). У каждого из них есть
свои плюсы и минусы. Задача использовать эти драйверы параллельно,
так, чтобы один не мешал другому, при максимально удобном переключении
между ними.

# Входные данные

Данную задачу будем решать в gentoo linux. Будем считать, что драйвер
[radeon](http://en.gentoo-wiki.com/wiki/Radeon) уже установлен.

# Реализация

## Добавим драйвер [fglrx](http://en.gentoo-wiki.com/wiki/Fglrx) в */etc/make.conf*

    # grep VIDEO /etc/make.conf
    VIDEO_CARDS="radeon fglrx"

## Установим необходимые компоненты

    root@linux# emerge -uDNv world

## Настроим загрузчик

Нужно, чтобы для каждого ядра генерировалось два пункта загрузчика:

  - Для открытого драйвера (с использованием
    [KMS](https://wiki.archlinux.org/index.php/Kernel_Mode_Setting)) в
    cmdline необходимо добавить


    radeon.modeset=1 modprobe.blacklist=fglrx

  - Для закрытого драйвера нужно добавить


    modprobe.blacklist=drm,drm_kms_helper,radeon,uvcvideo,ttm,videodev,videobuf2_core,videobuf2_memops,videobuf2_vmalloc

Сгенерируйте файл настроек загрузчика после внесения изменений.-

## Настроим X11

    root@linux# /opt/bin/aticonfig --initial --input=/etc/X11/xorg.conf

Скопируем полученный файл в *xorg.conf.fglrx-0*

    root@linux# cp /etc/X11/xorg.conf /etc/X11/xorg.conf.fglrx-0

## Обеспечим автоматическое переключение между драйверами, в зависимости от загруженных модулей ядра

Создадим файл */etc/local.d/radeon.start*

    root@linux# cat /etc/local.d/radeon.start
    #!/bin/sh
    if ((`lsmod | grep fglrx | wc -l` > 0))
            then if ((`eselect opengl list | grep ati | grep \* | wc -l` == 0))
                    then eselect opengl set ati
                    cp -f /etc/X11/xorg.conf.fglrx-0 /etc/X11/xorg.conf
            fi
    else
            if ((`eselect opengl list | grep xorg-x11 | grep \* | wc -l` == 0))
                    then eselect opengl set xorg-x11
            fi
            rm /etc/X11/xorg.conf
            #следующая строка необязательна - она меняет профиль на low
            echo low > /sys/class/drm/card0/device/power_profile
    fi

и сделаем его исполняемым

    chmod +x /etc/local.d/radeon.start

Разумеется, сервис local должен быть добавлен в уровень исполнения
default

    root@linux# rc-update add local default

Теперь после перезагрузки вы можете выбрать в меню вашего любимого
загрузчика, какой драйвер запускать. Для удобства также можно
установить из оверлея stuff (спасибо megabaks'у) dkms-gentoo.