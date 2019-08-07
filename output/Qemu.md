## Как примонтировать образ диска в формате raw?

Самый простой способ - воспользоваться утилитой kpartx из набора
[multipath-tools](http://christophe.varoqui.free.fr/)

Пример использования(disk.img - образ диска в формате raw):

    losetup /dev/loop0 /path/to/disk.img
    kpartx -a /dev/loop0

В результате в /dev/mapper появятся файлы вида loop0p<b>x</b>, где
<b>x</b> - номер раздела в образе. К слову, /dev/loop0 в данном случае -
полноценное блочное устройство, которое можно разбивать, делать с него
dd и т.д.

## Как примонтировать образ диска в формате qcow2?

Для этого в вашем ядре должна быть включена поддержка nbd (Network Block
Devices).

Загружаем модуль nbd с параметром max_part=8 , чтобы создавались
блочные устройства для разделов.

    # modprobe nbd max_part=8

Потом запускаем qemu-nbd

    # qemu-nbd --connect=/dev/nbd0 imagename.qcow

и можно мониторовать разделы

    # mount /dev/nbd0p1 /mnt

