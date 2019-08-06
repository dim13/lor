# Спецификации

На сайте производителя:
[1](http://www.dell.com/ru/business/p/vostro-3360/pd)

Вывод lspci:

    00:00.0 Host bridge: Intel Corporation 3rd Gen Core processor DRAM Controller (rev 09)
    00:02.0 VGA compatible controller: Intel Corporation 3rd Gen Core processor Graphics Controller (rev 09)
    00:14.0 USB controller: Intel Corporation 7 Series/C210 Series Chipset Family USB xHCI Host Controller (rev 04)
    00:16.0 Communication controller: Intel Corporation 7 Series/C210 Series Chipset Family MEI Controller #1 (rev 04)
    00:1a.0 USB controller: Intel Corporation 7 Series/C210 Series Chipset Family USB Enhanced Host Controller #2 (rev 04)
    00:1b.0 Audio device: Intel Corporation 7 Series/C210 Series Chipset Family High Definition Audio Controller (rev 04)
    00:1c.0 PCI bridge: Intel Corporation 7 Series/C210 Series Chipset Family PCI Express Root Port 1 (rev c4)
    00:1c.4 PCI bridge: Intel Corporation 7 Series/C210 Series Chipset Family PCI Express Root Port 5 (rev c4)
    00:1d.0 USB controller: Intel Corporation 7 Series/C210 Series Chipset Family USB Enhanced Host Controller #1 (rev 04)
    00:1f.0 ISA bridge: Intel Corporation HM77 Express Chipset LPC Controller (rev 04)
    00:1f.2 SATA controller: Intel Corporation 7 Series Chipset Family 6-port SATA Controller [AHCI mode] (rev 04)
    00:1f.3 SMBus: Intel Corporation 7 Series/C210 Series Chipset Family SMBus Controller (rev 04)
    01:00.0 Network controller: Atheros Communications Inc. AR9485 Wireless Network Adapter (rev 01)
    02:00.0 Ethernet controller: Atheros Communications Inc. AR8161 Gigabit Ethernet (rev 10)

Вывод lsusb:

    Bus 001 Device 002: ID 8087:0024 Intel Corp. Integrated Rate Matching Hub
    Bus 002 Device 002: ID 8087:0024 Intel Corp. Integrated Rate Matching Hub
    Bus 001 Device 001: ID 1d6b:0002 Linux Foundation 2.0 root hub
    Bus 002 Device 001: ID 1d6b:0002 Linux Foundation 2.0 root hub
    Bus 003 Device 001: ID 1d6b:0002 Linux Foundation 2.0 root hub
    Bus 004 Device 001: ID 1d6b:0003 Linux Foundation 3.0 root hub
    Bus 001 Device 003: ID 138a:0011 Validity Sensors, Inc. VFS5011 Fingerprint Reader
    Bus 001 Device 004: ID 0bda:0129 Realtek Semiconductor Corp.
    Bus 001 Device 005: ID 0c45:648b Microdia
    Bus 002 Device 004: ID 0cf3:e004 Atheros Communications, Inc.

# Устройства

## Wi-Fi

Работает из коробки.

## Ethernet

На текущий момент (ядро 3.8.1) поддержки карты

    02:00.0 Ethernet controller: Atheros Communications Inc. AR8161 Gigabit Ethernet (rev 10)

в ядре нет. Есть сторонний драйвер, который двигается в ядро, но,
поскольку нам нужно использовать его сейчас и немедленно, то
используем AUR:

    yaourt dkms-alx

Чтобы загрузить драйвер, введите команду **sudo modprobe alx**. В выводе
dmesg появится что-то типа такого:

    [ 9897.021035] Compat-drivers backport release: compat-drivers-v3.8-rc5-1-u
    [ 9897.021044] Backport based on linux.git v3.8-rc5
    [ 9897.021048] compat.git: linux.git
    [ 9897.021873] Qualcomm Atheros(R) AR816x/AR817x PCI-E Ethernet Network Driver
    [ 9897.022317] alx 0000:02:00.0: PHY has been configured.
    [ 9897.029040] alx 0000:02:00.0: alx(84:8f:69:d3:a6:09): Qualcomm Atheros Ethernet Network Connection

Дальше конфигурим сеть как обычно.

## Тачпад

В оригинальной Убунте от производителя тачпад работает через сторонний
неизвестный демон, а утилита для управления им написана на Java. Что
это за конь — раскопать не удалось, тем более найти пакеты не под
Убунту. Печальней всего оказалось, что на текущий момент (ядро
3.7.7) родной поддержки этого тачпада нет. Он определяется как обычная
мышка PS/2, но с недавнего времени это поправимо сторонними модулями:

    yaourt psmouse-alps-driver

После компиляции и установки драйвера следует перегрузить модуль
**psmouse**:

    sudo rmmod psmouse
    sudo modprobe psmouse

Вуаля — тачпад работает полноценно: есть прокрутка двумя пальцами в
любом направлении, управляемость через gsynaptics и тому подобные
прелести. Кстати, светодиод над тачпадом тоже можно включать и
выключать, эта функциональность в ядре есть «из коробки».

## Веб-камера

Работает из коробки.

## Кардридер

Работает из коробки. Одно «но»: его драйвер rts5139 пока живёт в
staging, это нужно учитывать, если хочется пересобрать ядро вручную, а
также если вылезут какие-то глюки.

## Видеокарта

Intel HD4000 начала нормально работать в ядре 3.4, более ранние версии
ядер используйте на свой страх и риск.

## Звуковая карта

Вот так работает:

    CONFIG_SND_HDA_INTEL=m
    CONFIG_SND_HDA_PREALLOC_SIZE=4096
    CONFIG_SND_HDA_HWDEP=y
    CONFIG_SND_HDA_RECONFIG=y
    CONFIG_SND_HDA_INPUT_BEEP=y
    CONFIG_SND_HDA_INPUT_BEEP_MODE=1
    CONFIG_SND_HDA_INPUT_JACK=y
    CONFIG_SND_HDA_PATCH_LOADER=y
    CONFIG_SND_HDA_CODEC_HDMI=y
    CONFIG_SND_HDA_CODEC_CIRRUS=y
    CONFIG_SND_HDA_GENERIC=y
    CONFIG_SND_HDA_POWER_SAVE=y
    CONFIG_SND_HDA_POWER_SAVE_DEFAULT=0

Остальные кодеки HDA можно выключить.

Завести встроенный микрофон пока не удалось.

## Мультимедиа-клавиши

Хардварные в полном объёме пока завести не удалось. Управление яркостью,
громкостью и медиаплеером работает нормально. Для переключения тачпада
можно использовать мои desktop-scripts:
[2](https://github.com/pfactum/desktop-scripts)

## Сканер отпечатков пальцев

Драйверов пока нет :(.

# Действия

## Регулирование яркости

Ядру при загрузке желательно передать

    acpi_osi=Linux acpi_backlight=vendor

чтобы использовался родной модуль для регулировки яркости. Иначе
сохранение её уровня не будет работать. Однако если используется
KDE, то опцию

    acpi_backlight=vendor

лучше убрать, так как иначе KDE не сможет программно регулировать
яркость.

## Гибернация

Гибернация работает из коробки, правда, с некоторыми странностями по
части USB. Поэтому рекомендуется поместить файлик tricks.sh с таким
вот содержимым в каталог /usr/lib/systemd/system-sleep/:

    #!/bin/sh

    case $1/$2 in
    pre/hibernate)
        echo 0 >/sys/power/tuxonice/replace_swsusp
        echo 0 >/sys/power/tuxonice/no_multithreaded_io
        echo 1 >/sys/power/tuxonice/no_flusher_thread
        echo 1 >/sys/power/tuxonice/compression/enabled
        echo 1 >/sys/power/tuxonice/user_interface/enable_escape
        echo 1 >/sys/power/tuxonice/user_interface/default_console_level
        echo lzo >/sys/power/tuxonice/compression/algorithm
        echo shutdown >/sys/power/disk
            ifconfig wlan0 down
            rmmod ehci_hcd
            rmmod ehci_pci
        sync
        echo 3 >/proc/sys/vm/drop_caches
        sync
        uksmctl -d
        sync
        ;;
    post/hibernate)
            modprobe ehci_pci
            modprobe ehci_hcd
        hdparm -B 254 /dev/sda
        uksmctl -a
        ;;
    pre/suspend)
        uksmctl -d
        sync
        ;;
    post/suspend)
        hdparm -B 254 /dev/sda
        uksmctl -a
        ;;
    esac

и сделать его исполняемым. Если не используете pf-kernel, все строчки,
относящиеся к TuxOnIce и UKSM, можно удалить.

Также ядру желательно передавать при загрузке параметр

    i8042.nopnp

иначе при гибернации будут сыпаться странные ошибки. Впрочем, без этого
параметра тоже всё работает, ошибки ни на что не влияют.

## Засыпание в память

Работает из коробки.