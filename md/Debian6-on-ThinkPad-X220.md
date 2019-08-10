## Введение

Небольшая инструкция, раскрывающая некоторые особенности по установке и
настройке `Debian6 (squeeze)` на ноутбук `Lenovo ThinkPad X220`.  
Будет кратко описан процесс установки (только моменты отличающиеся от
штатных), установки необходимых обновлений до версий
`backports/testing` (исключительно из-за неработоспособности некоторых
узлов ноутбука на стабильной ветке). Более подробно будет освещена
доработка функционирования горячих клавиш и сканера отпечатков
пальцев. Вариант изначальной установки `testing/sid` не
рассматривался в принципе из-за отсутствия Gnome2 (дискуссии
на тему альтернатив не уместны).

Немного о модели ноутбука (429058G, она же NYF58RT): 12.5" (1366x768)
IPS, Intel Core i7-2640M(2.8GHz), 4GB, 80GB SSD, 500GB (5400rpm), Intel
HD3000, WiFi, BT, TPM, FPR, WebCam, 6cell, Win 7 Pro 64

Хочу сразу акцент сделать на одном моменте: знаю, что это не "true \*nix
way", но работаю при установке/настройке преимущественно от root -
разумеется `sudo` также использую, но только на этапе
эксплуатации. Поэтому команды привожу без префикса `sudo`,
т.к. выполняются от root.

## Установка Debian

### Установка

Для установки был взят образ с официального сайта, актуальный на момент
установки
[debian-6.0.6-amd64-DVD-1.iso](http://cdimage.debian.org/debian-cd/6.0.6/amd64/iso-dvd/debian-6.0.6-amd64-DVD-1.iso),
залит с помощью dd на USB-flash носитель и запущен с него процесс
установки дистрибутива. Т.к. в данной модели ноутбука имеется и
SSD и SATA диск, то систему решил ставить на SSD, а рабочие разделы
разместить на жестком диске. Исходя из этого при установке вручную
разметил разделы:

  - <b>SSD:</b>
      - /boot (300Mb, ext4, noatime)
      - / (40Gb, ext4, noatime)
      - /home (\~40Gb, ext4, noatime)
  - <b>SATA:</b>
      - /var (20Gb, ext4, relatime)
      - Все рабочие и файловые разделы (ext4, relatime)

Swap раздел не создавал за ненадобностью (в ноут доустановил еще 4Gb
оперативной памяти - для работы вполне хватает).

Из предложенного списка метапакетов отметил: desktop, laptop, utils.

После установки Gnome запустился нормально, но иксы не "увидели"
разрешение экрана, также не определилась wifi карта, не работали
некоторые функциональные клавиши.

### Первичная настройка

Перво-наперво необходимо "починить" управление сетевыми интерфейсами,
если при установке сетевой интерфейс был сконфигурирован статически.
Если необходимо конфигурирование сети через network-manager, то
необходимо поправить конфигурацию интерфейсов в
`/etc/network/interfaces` - перевести сетевой интерфейс в режим авто:

    # This file describes the network interfaces available on your system
    # and how to activate them. For more information, see interfaces(5).
    
    # The loopback network interface
    auto lo
    iface lo inet loopback
    
    # The primary network interface
    auto eth0

Сразу же обратил внимание на наличие характерных периодических "кликов"
парковок головок жесткого диска. Проблема известная (извечная? на
предыдущем Lenovo была аналогичная ситуация), но решаемая просто.
Диагностируется также путём установки пакета `smartmontools` и
просмотром информации S.M.A.R.T по команде `smartctl -A
/dev/sd*` - а именно интересует параметр: `Load_Cycle_Count`, если его
значение от замера к замеру растёт (после "кликов" парковки), значит
оно и есть. Добавляем в файл `/etc/rc.local`, перед командой `exit 0`,
строку вида: `hdparm -S 244 -B 255 /dev/sda`. Также создаём скрипт для
обновления перезаписи параметров смарт при пробуждении ноутбука: файл
`/etc/pm/sleep.d/fix_sda_apm_param`

    #!/bin/sh
    
    case "${1}" in
        resume|thaw)
            hdparm -S 244 -B 255 /dev/sda
        ;;
    esac

Для оптимизации использования SSD (и продления его ресурса):

1.  Скорректировать fstab для задания оптимальных параметров
    монтирования разделов:
    
      - для /boot, / - discard,noatime,commit=60
      - для остальных - relatime,commit=60
      - временный раздел помещаем в память, изменяем на: tmpfs /tmp
        tmpfs defaults 0 0

2.  
    
    <div>
    
    Выставить в опциях загрузки ядра `elevator=noop`:

<!-- end list -->

  - Первый вариант - в конфигурационном файле `/boot/grub/grub.cfg`
    добавить к опциям загрузки ядра.
  - Второй вариант - в дефолтных настройках `/etc/default/grub` изменить
    параметр `GRUB_CMDLINE_LINUX_DEFAULT="quiet elevator=noop"` и после
    выполнить `update-grub`.

Так же можно и уменьшить время ожидания выбора ядра при старте (в
секундах): `GRUB_TIMEOUT=3`. Некоторые рекомендуют для ещё
большего уменьшения ожидания добавлять:

    GRUB_HIDDEN_TIMEOUT=0
    GRUB_HIDDEN_TIMEOUT_QUIET=true

Но это может быть чревато, т.к. при таймауте равном единице и двух этих
опциях у меня "проскакивало" без выбора даже с зажатым shift - в
результате пришлось править, загрузившись с LiveCD.

</div>

\# Добавить в конец `/etc/sysctl.conf` опции:

    vm.laptop_mode=5
    vm.dirty_writeback_centisecs=6000

Далее целесообразно перезагрузить ноутбук для применения всех внесённых
настроек.

### Обновление

\# Включить `contrib` и `non-free` репозитории squeeze. Для этого
добавить в файл `/etc/apt/sources.list` строку:

    deb http://ftp.ru.debian.org/debian/ squeeze main contrib non-free

1.  Обновить базу пакетов:
        # aptitude update
2.  Обновить систему - просто запустить графический менеджер обновлений.

\# Подключить `squeeze-backports`. Для этого добавить в файл
`/etc/apt/sources.list` строку:

    deb http://backports.debian.org/debian-backports squeeze-backports main contrib non-free

1.  Обновить базу пакетов:
    
        # aptitude update

2.  
    
    <div>
    
    Установить из backports:

\*

<div>

Свежее ядро (3.2.\*):

    # aptitude install -t squeeze-backports linux-image-3.2.0(посвежее).amd64

Если `aptitude` спросит "оставлять со старыми зависимостями"? Не
соглашаться, после чего спросит "обновить зависимости"?
Соглашатья - подтянет обновления для зависимых пакетов из
`backports`.

</div>

\* Пакеты для Wifi:

    # aptitude install -t squeeze-backports firmware-iwlwifi wireless-tools

\* Иксы:

    # aptitude install -t squeeze-backports xorg xserver-xorg xserver-xorg-core xserver-xorg-input-all xserver-xorg-video-all

Может что-то с собой подтянуть, что-то обновится, что-то установится -
соглашаться.

\* Поддержка opengl (2D/3D):

    # aptitude install -t squeeze-backports libgl1-mesa-dri libgl1-mesa-glx

</div>

1.  Загрузить ноутбук на новом ядре. Что по факту имеем в плюсе -
    разрешение экрана адекватное и wifi работает. В минусе - по
    прежнему не работают некоторые функциональные клавиши:
    <b>Fn+F6</b>, <b>Fn+F8</b>, <b>MicMute</b> (а также <b>VolumeUp</b>,
    <b>VolumeDown</b> - не включают звук при включённом Mute, хотя
    светодиод на кнопке Mute и гаснет).  
      
    Позже я обновил ядро, иксы, драйвера и некоторые утилиты до версий
    из репозиториев `wheezy` (`testing`). Не скажу, что было просто,
    т.к. часто пакеты тянут за собой очень много зависимостей и
    зачастую предлагается удалить Gnome2 и установить Gnome3,
    приходилось обновлять с помощью `apt-get` и подтягивая вручную
    нужные версии пакетов.

## Функциональные клавиши

Из-за особенностей конфигурации ядра Debian (конфигурационных опций - в
ядре не включена поддержка `CONFIG_ACPI_PROCFS` и
`CONFIG_ACPI_PROC_EVENT`, соответственно события `ibm/hotkey` посылаемые
на `/proc/acpi/event` начисто фильтруются). Таким образом, при
отсутствии реализации обработки клавиши в системе - свои
обработчики уже не навесить. В это упиралось большинство вопросов
в интернет на тему "как заставить работать функциональные клавиши". В
Ubuntu 12\*, Mint 13, Fedora 17 - эти клавиши работают. Различий в
исходниках модуля `thinkpad_acpi` почти нет (не принципиальные). В
исходниках acpi также никаких критических изменений на глаза не
попалось, так что причина может быть только в конфигурации ядра.
Пересобирать ядро для проверки я не стал, но работу `thinkpad_acpi`
проверил - исправно вызывает функцию генерации событий
`acpi_bus_generate_proc_event`, но та работает вхолостую, т.к.
`/proc/acpi/event` в системе отсутствует. Возможно решение двумя путями
- пересобрать ядро с нужными опциями и пропатчить `thinkpad_acpi` (я
выбрал второе - для включения светодиода его всёравно придётся
патчить, да и проще, чем пересобирать ядро). В принципе, стояла
задача заставить генерировать ACPI события нажатий клавиш
<b>Fn+F6</b>, <b>Fn+F8</b>, <b>MicMute</b>, и обрабатывать их в
юзерспейс, а также дополнительно обрабатывать нажатия
<b>VolumeUp</b>, <b>VolumeDown</b> (благо ACPI события на них
генерируются) - выключать <b>Mute</b>. Ну и заставить
работать светодиод на кнопке <b>MicMute</b>...

Уточню, что при наблюдении за событиями `ThinkPad Extra Buttons` с
помощью `evtest` выводились корректные события нажатий на все (\!)
функциональные клавиши даже на те, на которые `acpi_listen` молчал.

### Патч

1.  Скачать исходники ядра (находятся в
    `/usr/src/linux-source*.tar.bz2`).
2.  Скопировать из архива исходник `thinkpad_acpi.c` (в архиве
    `/linux-source-*/drivers/platform/x86/thinkpad_acpi.c`) в рабочий
    каталог где будем собирать.
3.  В этом каталоге создать файл `Makefile` (регистр\!) следующего
    содержания:

<!-- end list -->

    obj-m += thinkpad_acpi.o
    
    all:
            make -C /lib/modules/$(shell uname -r)/build M=$(PWD) modules
    
    clean:
            make -C /lib/modules/$(shell uname -r)/build M=$(PWD) clean

Те, кого напрягает отладочная информация, пристёгнутая к модулю - могут
добавить `stripe` для её удаления.

1.  Пропатчить `thinkpad_acpi.c` для корректной работы уведомлений о
    событиях:

<!-- end list -->

    static void hotkey_notify(struct ibm_struct *ibm, u32 event)
    {
    ...
            /* Legacy events */
            if (!ignore_acpi_ev &&
                (send_acpi_ev || hotkey_report_mode < 2)) {
    
                // sptim: Do nothing in Debian6 (without /proc/acpi/event)
                //acpi_bus_generate_proc_event(ibm->acpi->device, 
                //               event, hkey);
    
                // sptim: True generation (with bus id).
                acpi_bus_generate_netlink_event(
                        ibm->acpi->device->pnp.device_class,
                        ibm->acpi->device->pnp.bus_id,
                        event, hkey);
            }
    ...
    }

1.  Пропатчить `thinkpad_acpi.c` для работы индикатора на клавише
    <b>MicMute</b>. Этот патч и инструкцию нашёл
    [здесь](http://askubuntu.com/questions/125367/enabling-mic-mute-button-and-light-on-lenovo-thinkpads).

<!-- end list -->

    static const char * const tpacpi_led_names[TPACPI_LED_NUMLEDS] = {
        /* there's a limit of 19 chars + NULL before 2.6.26 */
        "tpacpi::power",
        "tpacpi:orange:batt",
        "tpacpi:green:batt",
        "tpacpi::dock_active",
        "tpacpi::bay_active",
        "tpacpi::dock_batt",
        "tpacpi::unknown_led",
        "tpacpi::standby",
        "tpacpi::dock_status1",
        "tpacpi::dock_status2",
        "tpacpi::unknown_led2",
        "tpacpi::unknown_led3",
        "tpacpi::thinkvantage",
        "tpacpi::unknown_led4", // sptim: add micmute led 
        "tpacpi::micmute",      // sptim: add micmute led  
    };
    #define TPACPI_SAFE_LEDS    0x5081U  // sptim: safe leds update
    
    ...
    
    static const struct tpacpi_quirk led_useful_qtable[] __initconst = {
    
        ...
    
        /* (1) - may have excess leds enabled on MSB */
    
        /* Defaults (order matters, keep last, don't reorder!) */
        { /* Lenovo */
          .vendor = PCI_VENDOR_ID_LENOVO,
          .bios = TPACPI_MATCH_ANY, .ec = TPACPI_MATCH_ANY,
          .quirks = 0x5fffU, // sptim: quirks update 
        },
    
        ...
    };

1.  Собрать модуль командой `make` (при этом возможно необходимо будет
    доустановить пакеты необходимых библиотек, не забываем, что должны
    быть установлены и части библиотек предназначенные для разработки -
    \*-dev пакеты\!).

\# На выходе получится файл `thinkpad_acpi.ko`, который необходимо
инсталлировать (в принципе, ничего такого смертельного не
случится, если инсталлировать сразу без предварительной
апробации):

\#\# Отключить модуль:

    # modprobe -r thinkpad_acpi

\#\# Сохранить оригинальный модуль (или переименовать
`/lib/modules/`<ядро>`/kernel/drivers/platform/x86/thinkpad_acpi.ko`
в `thinkpad_acpi.ko.orig`):

    # mv /lib/modules/<ядро>/kernel/drivers/platform/x86/thinkpad_acpi.ko \
         /lib/modules/<ядро>/kernel/drivers/platform/x86/thinkpad_acpi.ko.orig

\#\# Скопировать собранный модуль в
`/lib/modules/`<ядро>`/kernel/drivers/platform/x86`:

    # cp -f thinkpad_acpi.ko /lib/modules/<ядро>/kernel/drivers/platform/x86

1.  1.  Запустить модуль:
            # modprobe thinkpad_acpi

2.  Я для этих целей, когда экспериментировал, набросал шелл-скрипт
    запускаемый в рабочем каталоге (под root разумеется):

<!-- end list -->

    #!/bin/sh
    
    make clean
    make
    modprobe -r thinkpad_acpi
    cp -f thinkpad_acpi.ko /lib/modules/3.2.0-3-amd64/kernel/drivers/platform/x86
    modprobe thinkpad_acpi

\# Запустить `acpi_listen` (демон `acpid` должен быть установлен и
запущен) и понажимать функциональные клавиши. Вывод должен быть
такого вида:

    ibm/hotkey HKEY 00000080 00001006
    ibm/hotkey HKEY 00000080 00001008
    ibm/hotkey HKEY 00000080 0000101b

\# Проверить работу светодиода <b>MicMute</b> (должен появиться каталог
`/sys/devices/platform/thinkpad_acpi/leds/tpacpi::micmute`).  
Включение светодиода

    # echo 1 > /sys/devices/platform/thinkpad_acpi/leds/tpacpi::micmute/brightness

Выключение светодиода

    # echo 0 > /sys/devices/platform/thinkpad_acpi/leds/tpacpi::micmute/brightness

### Настройка

После успешного патча необходимо настроить скрипты для обработки
событий. <b>Fn+F8</b> (<b>Fn+F6</b> мне не нужен),
<b>MicMute</b>, <b>VolumeUp</b>, <b>VolumeDown</b>. Далее
подразумевается, что пакеты `acpi-support` и `xinput`
установлены (будут активно использоваться), а также пакет
`libnotify-bin` - для вывода уведомлений.

<div>

  - 
    
    <div>
    
    <b>`Fn+F8`</b> - Включение/отключение тачпада.

За основу был взят скрипт `asus-touchpad.sh` из пакета `acpi-support`.  
Файл `/etc/acpi/events/thinkpad-touchpad`:

    # This is called when the user presses Fn-F8.
    # This toggles the touchpad on and off.
    
    event=ibm/hotkey HKEY 00000080 00001008
    action=/etc/acpi/thinkpad-touchpad.sh

Файл `/etc/acpi/thinkpad-touchpad.sh`:

    #!/bin/sh
    
    set -e
    
    pff=/usr/share/acpi-support/power-funcs
    [ -f $pff ] || exit 0
    
    atp_error() {
            logger -t${0##*/} -perr -- $*
            exit 1
    }
    
    . $pff || atp_error "Sourcing $pff failed"
    
    [ -x /usr/bin/xinput ] || atp_error "Please install package xinput to enable toggling of touchpad devices."
    
    getXconsole
    
    ENABLEPROP="Synaptics Off"
    # Get the xinput device number and enabling property for the touchpad
    XINPUTNUM=$(xinput --list | sed -nr "s|.*[Tt]ouch[pP]ad.*id=([0-9]+).*|\1|p")
    [ "$XINPUTNUM" ] || atp_error "Invalid TouchPad id '$XINPUTNUM'"
    
    # Get the current state of the touchpad
    TPSTATUS=$(xinput --list-props $XINPUTNUM | awk "/$ENABLEPROP/ { print \$NF }")
    
    # Switch state
    case $TPSTATUS in
            0)
                    xinput --set-int-prop $XINPUTNUM "$ENABLEPROP" 8 1
                    ;;
            1)
                    xinput --set-int-prop $XINPUTNUM "$ENABLEPROP" 8 0
                    ;;
            *)
                    atp_error "Invalid TouchPad status '$TPSTATUS'"
                    ;;
    esac

upd: **ВАЖНО\!** Для тех, у кого скрипт не отрабатывает корректно, его
нужно немного изменить. Вместо вызова getXconsole - жёстко прописываем
дисплей и файл авторизации.

    ...
    #getXconsole
    export DISPLAY=:0
    export XAUTHORITY=/home/имя пользователя/.Xauthority
    ...

Лично у меня никогда не создавалось файла \~/.Xauthority (вместо этого
файл генерируется в /var/run/gdm3/... с рандомной частью в названии
каталога, так что необходимо создать следующий файл для создания
символьной ссылки при начале сессии (в случае использования среды
gnome2) или добавить в уже существующий строку с созданием ссылки перед
командой запуска среды. Файл: `~/.xsession`

    #!/usr/bin/env bash
    ln -f -s "$XAUTHORITY" ~/.Xauthority
    exec gnome-session

</div>

  - 
    
    <div>
    
    <b>`MicMute`</b> - Включение/отключение микрофона.

Файл `/etc/acpi/events/thinkpad-mutemic`:

    # This is called when the user presses MicMute.
    # This toggles the microphone mute on and off.
    
    event=ibm/hotkey HKEY 00000080 0000101b
    action=/etc/acpi/thinkpad-mutemic.sh

Файл `/etc/acpi/thinkpad-mutemic.sh`:

    #!/bin/sh
    
    set -e
    
    pff=/usr/share/acpi-support/power-funcs
    [ -f $pff ] || exit 0
    
    atp_error() {
            logger -t${0##*/} -perr -- $*
            exit 1
    }
    . $pff || atp_error "Sourcing $pff failed"
    getXconsole
    
    MICMUTE=/sys/devices/platform/thinkpad_acpi/leds/tpacpi::micmute/brightness
    CHANNEL="Capture"
    
    notify() {
            notify-send -t 1000 -i microphone-sensitivity-muted-symbolic "$1"
    }
    
    if amixer sget "$CHANNEL",0 | grep '\[on\]' ; then
        amixer sset "$CHANNEL",0 toggle
        echo 1 > $MICMUTE
        notify "Mic MUTE"
    else
        amixer sset "$CHANNEL",0 toggle
        echo 0 > $MICMUTE
        notify "Mic ON"
    fi

</div>

  - 
    
    <div>
    
    <b>`VolumeUp`</b> - Увеличение громкости звука.

Файл `/etc/acpi/events/thinkpad-volumeup`:

    # This is called when the user presses VolumeUp.
    # This toggles the mute off.
    
    event=button/volumeup VOLUP 00000080
    action=/etc/acpi/thinkpad-unmute.sh

Файл `/etc/acpi/thinkpad-unmute.sh`:

    #!/bin/sh
    
    [ -f /usr/share/acpi-support/key-constants ] || exit 0
    . /usr/share/acpi-support/key-constants
    
    if amixer sget Master,0 | grep '\[off\]' ; then
        acpi_fakekey $KEY_MUTE
    fi

</div>

  - 
    
    <div>
    
    <b>`VolumeDown`</b> - Уменьшение громкости звука.

Файл `/etc/acpi/events/thinkpad-volumedown`:

    # This is called when the user presses VolumeDown.
    # This toggles the mute off.
    
    event=button/volumedown VOLDN 00000080
    action=/etc/acpi/thinkpad-unmute.sh

</div>

</div>

## Сканер отпечатков пальцев

В ThinkPad X61s у меня стоял дистрибутив `Fedora 14.1 (RFRemix)`, там
всё работало из коробки в том числе и авторизация в `gdm` с помощью
встроенного сканера отпечатков пальцев. Т.е. в `Gnome` уже была
реализована GUI часть для работы с `FingerPrint` через расширение
темы gdm насколько я понял (или это есть в мейнстриме изначально?).

В `squeeze` и его `backports` необходимых библиотек не обнаружил, а из
`wheezy` многое (почти всё, что завязано на GUI) тянет за собой Gnome3.
Пришлось брать исходники пакетов, доустанавливать необходимые
dev-пакеты и средства разработки и собирать вручную.

### Сборка

Понадобятся следующие пакеты исходников:

  - [libfprint-0.4.0](http://people.freedesktop.org/~hadess/libfprint-0.4.0.tar.bz2)
    (freedesktop.org)
  - [fprintd-0.2.0](http://people.freedesktop.org/~hadess/fprintd-0.2.0.tar.bz2)
    (freedesktop.org)  
    Более свежие версии требуют более новых библиотек в том числе `GIO`
    которую мне не удалось подтянуть, последняя версия -
    [fprintd-0.4.1](http://people.freedesktop.org/~hadess/fprintd-0.4.1.tar.bz2).
  - [shared-mime-info-1.0](http://people.freedesktop.org/~hadess/shared-mime-info-1.0.tar.xz)
    (freedesktop.org)
  - [fprint\_demo](https://nodeload.github.com/dsd/fprint_demo/tarball/master)
    (тарбол с github.com/dsd/fprint\_demo)

Что необходимо установить со всеми подтягиваемыми зависимостями:

    # aptitude install gcc intltool make libxml2-dev libglib2.0-dev libusb-1.0-0-dev libnss3-dev libmagickcore-dev libdbus-1-dev libpolkit-gobject-1-dev libpam0g-dev libxv-dev

Собрать shared-mime-info:

    ./configure
    make
    make install

Собрать libfprint:

    ./configure --prefix=/usr --enable-exanples-build --enable-x11-examples-build 
    make
    make install

Собрать fprintd:

    ./configure —prefix=/usr --enable-pam
    make
    make install

Собрать fprint\_demo:

    ./autogen.sh
    ./configure —prefix=/usr
    make
    make install

### Настройка

1.  Внести в группу `plugdev` (после чего необходимо перелогиниться).

\# Создать новое правило для `udev`, файл
`/etc/udev/rules.d/91-upek_rule`:

    # Доступ к устройству для группы и создание для него симлинка.
    ATTRS{idVendor}=="147e", ATTRS{idProduct}=="2016", SYMLINK+="input/touchchip-%k", MODE="0664", GROUP="plugdev"
    # Включение энергосбережения.
    ATTRS{idVendor}=="147e", ATTRS{idProduct}=="2016", ATTR{power/control}=="*", ATTR{power/control}="auto"

\# Перезагрузить ноутбук или создать фейковое событие для `udev` чтобы
правило сработало:

    # udevadm trigger --attr-match=idVendor="147e" --attr-match=idProduct="2016"

### Проверка

Запустить от юзера `fprint_demo` (можно из консоли, можно через меню
гнома <i>Приложения-\>Стандартные-\>fprint project demo</i>). Должно
запускаться, находить устройство `UPEK Eikon 2`. Для пробы
зарегистрировать палец на первой вкладке и проверить на
второй. Если всё сделано правильно - должно работать корректно. Да,
устройство не поддерживает отдачу "снимков" (?) - поэтому "весёлых
картинок" в программе не видно.

Что не удалось достичь - чтобы при авторизации в gdm можно было зайти по
отпечатку пальца. Т.е. добавление `pam_fprintd.so` в
`/etc/pam.d/common-auth` никак не помогло. Также не помог
`gdm_plugin_fingerprint` от Fedora14, установка gdm вместо gdm3 тоже,
установка gdm от федоры вообще отказалась работать корректно (как
подозреваю из-за наслоения конфигов). Хотя конечно интерес был
скорее академический, т.к. авторизацию по отпечаткам не использую.

## Заключение

В принципе, результатом удовлетворён - всё, что необходимо - работает.

Хотя с сожалением понимаю, что рано или поздно (скорее второе - со
следующим апгрейдом ноутбука) придётся переходить на более свежий
дистрибутив и там уж выбирать "меньшее зло". Имхо это будет Xfce.
