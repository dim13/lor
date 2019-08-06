Если после очередного обновления сверхнестабильной ветки используемого
дистрибутива, установки вражеской операционной системы на соседний
раздел жёсткого диска или случайного выполнения команды head -c 446
/dev/zero \> /dev/sda Ваш GNU/Linux перестал загружаться, причём сбой
происходит до загрузки ядра (сразу после заставки BIOS POST),
загрузчик (в нашем случае GRUB) нужно переустановить.

## Вкратце

1.  Загрузиться с livecd
2.  Получить shell с правами суперпользователя

\#

    mount /dev/sdXY /mnt
    for f in proc sys dev; do mount -o bind /$f /mnt/$f; done
    chroot /mnt /usr/sbin/grub-install /dev/sdX
    for f in dev sys proc; do umount /mnt/$f; done
    umount /mnt
    reboot

## Как это сделать?

Поскольку загрузчик на жёстком диске уже не работает, нужно запустить
систему с другого носителя. Хорошим выбором является
[SystemRescueCD](http://sysresccd.org) или [Finnix](http://finnix.org),
но подойдёт любой livecd, на котором можно запустить shell,
подмонтировать корневой раздел и сделать в него chroot.

## Система на LiveCD запущена, что делать дальше?

Получаем shell (интерактивную командную строку).

В графическом окружении пользователя это делается запуском т.н.
"эмулятора терминала". Примеры: xterm, Konsole, GNOME Terminal,
LXTerm... Если система ориентирована на работу из консоли, то shell у
Вас уже есть. Убедитесь в том, что этот shell - не результат сбоя при
попытке запустить livecd (на экране не должно быть ошибок вида "gave up
waiting for root device").

Получаем права суперпользователя. Тут всё целиком и полностью зависит от
используемого livecd. Обычно помогают команды:

    sudo -i
    su

Также права root могут быть уже в наличии - это можно определить по
символу "\#" в конце приглашения командной строки:

    root@somelivecdname#
    #
    gentoo #

## Поиск корневой файловой системы

Для начала мы определим, где хранится корневая файловая система.

Если она хранится просто на разделе жёсткого диска (частый случай;
обычно так и есть, если Вы не знаете, что это ещё может быть), нам
помогут утилиты blkid и fdisk.

Типичный вывод blkid:

    # blkid
    /dev/sda5: UUID="988b9f74-60e0-43fc-a490-47ba1b81190d" TYPE="crypto_LUKS"
    /dev/sda6: UUID="6d085fca-0675-4fd8-bb14-59e1b06022f6" TYPE="crypto_LUKS"
    /dev/sda4: LABEL="root" UUID="6452ebcc-a034-40e6-81a4-148e2ef17a59" TYPE="ext3"
    /dev/mapper/sda5_crypt: UUID="8ddef706-e5af-48e1-b6be-4558d9472026" TYPE="swap"
    /dev/mapper/sda6_crypt: UUID="0ca95dfb-2c59-4d6c-83ae-4497bff0f78c" SEC_TYPE="ext2" TYPE="ext3"
    /dev/sda2: LABEL="USER" UUID="cd613ec2-ae98-4aa9-a15e-be0156a5a43b" TYPE="ext3" SEC_TYPE="ext2"
    /dev/sda1: LABEL="disk" UUID="9f26970e-7638-4925-9ba6-148eea4537e4" TYPE="ext3" SEC_TYPE="ext2"
    /dev/sdb1: LABEL="SAMSUNG" UUID="17F6-2E45" TYPE="vfat"
    /dev/sdc1: LABEL="WD Passport" UUID="1E29-4D6E" TYPE="vfat"
    /dev/sdb3: LABEL="STORAGE" UUID="52202708-6973-4101-8571-51c81acee409" TYPE="ext3"

Здесь видно кучу разделов на разных дисках и с разными файловыми
системами и метками. Здесь нужный раздел можно легко определить
по метке "root" на его файловой системе. В других случаях корневым
может быть единственный раздел с POSIX-совместимой файловой
системой (например, ext3, reiserfs, xfs...).

**Примечание:** blkid определяет ФС ext4 как ext3.

Или можно воспользоваться fdisk -l:

    # fdisk -l

    Disk /dev/sda: 250 GB, 250056737280 bytes
    255 heads, 63 sectors/track, 30401 cylinders
    Units = cylinders of 16065 * 512 = 8225280 bytes

       Device Boot      Start         End      Blocks   Id  System
    /dev/sda1   *           1       14302   114880783    7  HPFS/NTFS
    /dev/sda2           14303       26704    99611032    7  HPFS/NTFS
    /dev/sda4           26705       28082    11060752   83  Linux
    Warning: Partition 4 does not end on cylinder boundary.
    /dev/sda3           28083       30402    18627367    5  Extended
    Warning: Partition 3 does not end on cylinder boundary.
    /dev/sda5           28083       28331     1992060   82  Linux swap
    Warning: Partition 5 does not end on cylinder boundary.
    /dev/sda6           28332       30402    16627275   83  Linux
    Warning: Partition 6 does not end on cylinder boundary.

    Disk /dev/sdb: 250 GB, 250056737280 bytes
    255 heads, 63 sectors/track, 30401 cylinders
    Units = cylinders of 16065 * 512 = 8225280 bytes

       Device Boot      Start         End      Blocks   Id  System
    /dev/sdb1               1       12768   102558928    c  FAT32 LBA
    /dev/sdb3           12769       30401   141629040   83  Linux
    Disk /dev/sdc: 120 GB, 120031511040 bytes
    255 heads, 63 sectors/track, 14593 cylinders
    Units = cylinders of 16065 * 512 = 8225280 bytes

       Device Boot      Start         End      Blocks   Id  System
    /dev/sdc1               1       14593   117218241    c  FAT32 LBA

Здесь показаны размеры разделов, а сами разделы рассортированы по
жёстким дискам, к которым они принадлежат; также показаны размеры
разделов. Исходя из этих данных, мы тоже можем определить, какой из
разделов является корневым.

Наконец, мы могли просто запомнить имя корневого раздела (например,
/dev/sda4).

    # mount /dev/sda4 /mnt

Кроме того, root может быть зашифрованным, находиться на LVM или RAID.
Некоторые предназначенные для восстановления системы LiveCD
автоматически включают при запуске все эти вещи (тогда
возможный корневой раздел можно найти в выводе blkid), но ручная
работа также может потребоваться.

Если раздел зашифрован при помощи LUKS (Linux Unified Key Setup), в
выводе blkid будет присутствовать TYPE="crypto_LUKS". Такие диски
можно открыть следующей командой:

    # cryptsetup luksOpen /dev/sdXY root
    Enter passphrase: <здесь вводим пароль; пароль при вводе никак не показан>
    # mount /dev/mapper/root /mnt

LVM-тома подключаются следующим образом:

    # vgscan
      Reading all physical volumes.  This may take a while...
      Found volume group "<...>" using metadata type lvm2
    # vgchange -ay
      <...> logical volume(s) in volume group "<...>" now active
    # lvscan
      ACTIVE            '/dev/<...>/root' [<размер>] inherit
      ACTIVE            '/dev/<...>/someothername' [<размер>] inherit
    # mount /dev/<...>/root /mnt

Естественно, имя volume group и logical volume в Вашем случае может
оказаться другим.

\<FIXME: работа с soft-RAID/fake-RAID\>

Это всё? Нет.

Поскольку в современных дистрибутивах файлами в /dev управляет udev, на
выключенной системе директория /dev пуста. Соответственно, попытка
обратиться к /dev/sda (которого нет в /mnt/dev/) закончится
неудачно. Кроме того, в редких случаях система может потребовать
себе ещё и /proc и /sys. Это мы тоже учтём:

    # for f in dev proc sys; do mount -o bind /$f /mnt/$f; done

Более того, в некоторых установках /boot бывает вынесен на отдельный
раздел (часто в этом случае монтирование корня тоже затруднено
необходимостью выполнять cryptsetup, поднимать LVM или RAID).
Соответственно, /boot (на котором и хранится загрузчик) тоже нужно
монтировать:

    # chroot /mnt mount /boot

## Установка загрузчика

Пользователи Debian и производных дистрибутивов (Ubuntu, Mint) могут
просто воспользоваться одной командой и переходить к следующему
разделу:

    # chroot /mnt dpkg-reconfigure grub-pc

**Примечание:** на большинство вопросов (о параметрах ядра и т.п.) нужно
отвечать нажатием Enter; в вопросе о том, куда ставить загрузчик, нужно
ставить галочки клавишей ПРОБЕЛ, а подтверждать свой ответ - ENTER.

В общем же случае действия пользователя будут чуть более сложными:

    chroot /mnt
    grub-install /dev/sdX
    exit

...где sdX - целевой жёсткий диск. Обратите внимание: ставить загрузчик
нужно именно на жёсткий диск (в MBR), а не на его раздел. Второй
вариант иногда тоже нужен, но тогда читатель сам знает, что ему
делать.

Также может быть полезной выполненная в chroot'е (т.е. после chroot и
перед exit) команда:

`grub-mkconfig -o /boot/grub/grub.cfg`

...если используемый загрузчик - GRUB2, и список доступных для загрузки
пунктов меню загрузчика должен быть изменён.

## Завершение работы

Для начала размонтируем те файловые системы, которые были смонтированы
внутрь /mnt. Это может быть /boot, если мы его монтировали ранее:

    # chroot /mnt umount /boot

Также необходимо размонтировать /proc, /sys и /dev:

    for f in sys proc dev; do umount /mnt/$f; done

И, наконец, размонтируем саму корневую файловую систему:

    # umount /mnt

Если мы пользовались cryptsetup, шифрованный диск нужно отключить:

    # cryptsetup luksClose root

В случае LVM для завершения работы с ним тоже нужна отдельная команда:

    # vgchange -an
      0 logical volume(s) in volume group "<...>" now active

\<FIXME: завершение работы с RAID\>

Всё\! Осталось только перезагрузиться (например, командой reboot) и
надеяться, что загрузчик заработает.