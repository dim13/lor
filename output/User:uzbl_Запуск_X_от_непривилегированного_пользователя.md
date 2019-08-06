Зачем? Меньше рутовых процессов — меньше гимора; неоднократно ловил фриз
иксов несколько лет назад.

Предложеный [здесь
метод](https://wiki.archlinux.org/index.php/Using_File_Capabilities_Instead_Of_Setuid#xorg-xserver)
даёт иксам излишние привилегии, на самом деле можно обойтись меньшим; на
примере закрытого драйвера nvidia и по-порядку:

    chmod -s    /usr/bin/Xorg
    setcap cap_sys_rawio=ep /usr/bin/Xorg

    chmod 666   /dev/input/event*
    chmod 660   /dev/tty{0,2}

    mknod -m660 /dev/nvidia0   c 195 0
    mknod -m660 /dev/nvidiactl c 195 255
    chgrp video /dev/nvidia*

юзер иксов должен быть в группах video и tty (либо, другая, обл.
доступом к /dev/tty\*), также, лучше написать правило udev, дабы
он создавал /dev/input\* принадлежащими группе xorg-a, например
/etc/udev/rules.d/111.rules:

    ACTION=="add", KERNEL=="event?", GROUP="video", MODE:="660"

Иксы запускаем так:

    Xorg :0 -logfile xlog_server -logverbose 999 ...другие опции...

перед этим необходимо удалить старые логи и файлы .X\* принадлежащие
руту из /tmp.

Проблем в работе не было обнаружено, 3d не проверялось.

Если не будет получаться — проще всего запускать X под strace и
смотреть, на чём он отваливается.