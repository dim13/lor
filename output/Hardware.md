## Как использовать дополнительные клавиши и колесики? Как назначить любую клавишу на запуск чего-либо?

Большинство современных сред имеют инструменты для их использования.

В Gnome задействовать эти клавиши можно с помощью утилиты acme и
настроить с помощью acme-properties (Gnome Control
Center-\>Настройки Рабочего Стола-\>Мульмедийные Клавиши).

В KDE 3.1 и раньше в kmenuedit - на запуск приложений и настройка клавиш
в соответствующих программах.

В KDE 3.2 ещё и отдельный пункт в Центре управления KDE.

В WindowMaker это делается с помощью WPrefs

В других - или тоже используется утилита, или комбинации прописываются в
файлах конфигурации.

Кроме того, можно использовать lineakd -
[1](http://lineak.sourceforge.net), им же можно назначить клавиши на
запуск приложений.

Или Hotkeys ( [2](http://freshmeat.net/projects/hotkeys/)) - если
разобраться, то можно вообще любыми клавишами "рулить" :).

Или xev & xmodmap. О том, как узнать коды клавиш и их использовать -
[3](http://www.nixp.ru/cgi-bin/go.pl?q=articles;a=multikeyboard).

А как эффективно использовать сами клавиши, тут -
[4](http://www.linux.org.ru/jump-message.jsp?msgid=452963)

Если вы хотите, чтобы клавиши были доступны для всех пользователей,
внесите работающие изменения в файл /etc/X11/Xmodmap.

Чтобы начало работать колесо прокрутки на клавиатуре, необходимо указать
**atkbd.scroll=1** в параметрах ядра.

## Как в /dev именуются IDE-устройства (жесткие диски, накопители CD-ROM), SCSI-устройства, ?

  - hda - primary master
  - hdb - primary slave
  - hdc - secondary master
  - hdd - secondary slave

Цифры означают раздел - hda1 - первый раздел, hda2 - второй раздел и
т.п. Приводы IDE CDROM не имеют отдельного обозначения. Имя cdrom
является симлинком на действующее устройство (в общем случае hdX,
если не включена эмуляция SCSI)

Соответственно для SCSI, где нет master-slave, sdX, а цифры все так же
обозначают раздел, например, sda1 - первый раздел, sda2 - второй
раздел и т.п.

Приводы SCSI CDROM (или приводы IDE, работающие в режиме эмуляции SCSI)
имеют обозначение scdX, srX или sgX, где X - номер устройства.

## Как включить DMA (UDMA)?

ВНИМАНИЕ: Неверные настройки могут привести к порче жесткого диска и/или
пропаже данных\!\!\!

Обычно все просто. Существует специальная программа, поставляемая в
любом дистрибутиве, называется hdparm. В общем случае достаточно
сделать следующее: \`hdparm -c1 <устройство>\` - установки 32-битного
асинхронного режима - большинство дисков прекрасно работают с ним,
\`hdparm -d1 <устройство>\` - собственно, включаем сам режим DMA.

PIO режим и DMA режим можно изменять и одной командой. \`hdparm
-X<число>\`. Хочу обратить внимание, что если режим не
поддерживается, то команда не сработает.

В дистрибутивах от RedHat постоянные настройки hdparm хранятся в
/etc/sysconfig/harddisk.

Подробнее о параметрах hdparm можно прочитать на
[5](http://computerlib.narod.ru/html/hdparm.htm)

Вот пример использования hdparm:

    #!/bin/sh
    # hdparm script

    PATH=/usr/local/sbin:/usr/local/bin:/sbin:/bin:/usr/sbin:/usr/bin

    MAXTOR=/dev/hda
    WD=/dev/hdb
    DVD=/dev/hdc
    BURNER=/dev/hdd

    case "$1" in
      start)
            echo -n "Adjusting drive parameters using hdparm... "

            # d1 = dma on
            # c3 = 32bit io w/sync
            # m16 = read upto 16 sectors at a time
            # u1 = unmask other interrupts while processing disk interrupt

            if [ `hostname` == "marsala" ]
            then
                hdparm -d1 -c3 -m16 -u1 $MAXTOR
                hdparm -d1 -c3 -m16 -u1 $WD
                hdparm -d1 -u1 $DVD
                # hdparm -d1 -u1 $BURNER
            fi

            echo "done."
            ;;
      stop)
            ;;
      restart|force-reload)
            ;;
      *)
            ;;
    esac

    exit 0

ВНИМАНИЕ\! Если у вас возникли проблемы с жестким диском или стала
зависать система при загрузке, то уберите параметр -u1.

\==hdparm выдает ошибку HDIO_SET_DMA failed: Operation not permitted
using_dma = 0 (off), что делать?==

У вас не включена поддержка чипсета вашей материнской платы в разделе
поддержки винтов. Или же поддержка собрана модулем (для 2.6.х). В
первом случае нужно пересобрать ядро, включив поддержку, а во втором
случае - загружать при старте модуль, прописав его в modprobe.conf

Есть еще вариант для чипсетов с SATA, когда в биосе выставлен неверный
режим работы SATA-контроллера.

## Насколько хороша в linux поддержка SATA для ядер 2.4? Как поставить linux с ядром, не поддерживающим SATA?

Вот здесь море документации и список оборудования
[6](http://www.linuxmafia.com/faq/Hardware/sata.html) Стоит заглянуть и
сюда [7](http://linux.yyz.us/sata/faq-sata-raid.html)

## Какой покупать тюнер, чтобы нормально работал под Linux'ом?

Имеет смысл брать модели на чипе SAA7130/SAA7134. Например, Avermedia
305/307.

Если с деньгами не очень, то на чипе Bt878. Например, Avermedia 203.

## Как заставить работать тв-тюнер (на примере Avermedia tv studio 203)? Как подключить пульт (lirc)?

Для примера возьмем avermedia tvstudio 203.

Для начала сконфигурим ядро:

    make menuconfig:

    Раздел Character devices, I2C support

    I2C support
    I2C bit-banging interfaces
    < > ELV adapter
    < > Velleman K9000 adapter
    < > NatSemi SCx200 I2C using GPIO pins
    < > NatSemi SCx200 ACCESS.bus
    < > I2C PCF 8584 interfaces
    I2C device interface
    I2C /proc interface (required for hardware sensors)

    Раздел Multimedia devices:
    Video For Linux
    Video For Linux --->
    Radio Adapters --->

    Подраздел Video For Linux:

    [*] V4L information in proc filesystem
    --- Video Adapters
    BT848 Video For Linux
    < > Mediavision Pro Movie Studio Video For Linux
    < > CPiA Video For Linux
    <M> SAA5249 Teletext processor
    < > SAB3036 tuner
    < > Stradis 4:2:2 MPEG-2 video driver (EXPERIMENTAL)
    < > Zoran ZR36057/36060 Video For Linux
    < > Zoran ZR36120/36125 Video For Linux

    Раздел Sound:

    Sound card support

    BT878 audio dma

    TV card (bt848) mixer support

Конфиг для ядра 2.6 выглядит почти так же:

Пульт и тюнер на BT878:

    #
    # Linux InfraRed Controller
    #
    CONFIG_LIRC_SUPPORT=m
    CONFIG_LIRC_MAX_DEV=2
    CONFIG_LIRC_I2C=m
    CONFIG_LIRC_GPIO=m
    # CONFIG_LIRC_BT829 is not set
    # CONFIG_LIRC_IT87 is not set
    # CONFIG_LIRC_ATIUSB is not set
    # CONFIG_LIRC_PARALLEL is not set
    # CONFIG_LIRC_SERIAL is not set
    # CONFIG_LIRC_SIR is not set

    ......

    #
    # Multimedia devices
    #
    CONFIG_VIDEO_DEV=m

    #
    # Video For Linux
    #

    #
    # Video Adapters
    #
    CONFIG_VIDEO_BT848=m
    # CONFIG_VIDEO_PMS is not set
    CONFIG_VIDEO_BWQCAM=m
    CONFIG_VIDEO_CQCAM=m
    CONFIG_VIDEO_W9966=m
    # CONFIG_VIDEO_CPIA is not set
    CONFIG_VIDEO_SAA5246A=m
    CONFIG_VIDEO_SAA5249=m
    CONFIG_TUNER_3036=m
    # CONFIG_VIDEO_STRADIS is not set
    # CONFIG_VIDEO_ZORAN is not set
    CONFIG_VIDEO_SAA7134=m
    # CONFIG_VIDEO_MXB is not set
    CONFIG_VIDEO_DPC=m
    # CONFIG_VIDEO_HEXIUM_ORION is not set
    # CONFIG_VIDEO_HEXIUM_GEMINI is not set
    CONFIG_VIDEO_CX88=m
    CONFIG_VIDEO_OVCAMCHIP=m

Собираем ядро. Далее берем lirc (www.lirc.org) и собираем его. Lirc
представляет собой модули для ядра и полезные утилиты для
конфигурации и настройки пульта. В случае ядра 2.6 нужно брать
последний snapshot. Если вы пропатчили ядро, чтобы появился lirc, то из
пакета lirc драйверы собирать не стоит.

Теперь в /etc/modules.conf (для ядра 2.6 /etc/modprobe.conf) пропишем
наш тюнер и пульт:

    # i2c
    alias char-major-89 i2c-dev
    options i2c-core i2c_debug=1
    options i2c-algo-bit bit_test=1

    # lirc
    alias char-major-61 lirc_gpio

    # bttv
    alias char-major-81 videodev
    alias char-major-81-0 bttv
    options bttv card=41 tuner=5 radio=1 pll=1 automute=0
    options tuner debug=1

Как видно выше, сам тв-тюнер задается номером. В данном случае card=41.
Список карт и чипсетов есть в поставке bttv ( [8](http://bytesex.org)).
Если конкретно вашего тюнера в списке нет, попробуйте просто перебрать
номера карточек.

Наверно понятно, что после установки нового ядра придется перегрузиться
:)

Теперь займемся конфигурацией:

/etc/lircd.conf - это файл настройки параметров пульта. Для большей
части пультов такие файлы настройки уже есть. Они поставляются в
составе lirc. Вам лишь нужно найти свой пульт и переписать файл в
/etc.

/etc/.lircrc - файл настройки команд пульта. Поскольку толковых примеров
нет, а читать документацию все боятся, приведу свой (jackill) -
[9](http://fine.kalinovka.net/articles/hobby/lorFAQ/files/lircrc.tgz)

irexec и irxevent - утилиты из lirc. За более подробной информацией
обращайтесь к документации.

В $HOME/.xinitrc перед запуском своего оконного менеджера пропишите
irexec.

## Как настроить тюнер AverMedia Tv Studio 305/307?

Нашелся добрый человек, который описал процесс.

[10](http://oppserver.net/downloads/dvb/unix/faq/www.linuxtv.org/v4lwiki/index.php/AverTV_305/307_linux_user_guide.html)

## Не могу звук на TV-тюнер вывести. Что делать?

Это довольно просто. Подключаем колонки или наушники в сам тюнер.

1.  Работает? Отлично. Тогда запускаем, скажем, alsamixer и ползунок
    Line выкручиваем до упора, после чего вставляем наш кабель jack-jack
    между линейным входом в звуковую и выходом с тюнера обратно.
2.  Не работает? Плохо. Набираем **/sbin/lsmod**. Если видим btaudio, а
    тюнер у нас avermedia 305, то есть вариант, при котором его наличие
    так пагубно влияет на звук. Ну что тут делать? Заходим под рутом,
    удаляем в **/lib/modules/ядро/_дальше_найдете_ модуль**
    btaudio и выполняете команду depmod -a. На случай атомной войны
    перегружаемся и смотрим что вышло.

## Как настроить USB-устройство в Linux?

Прочитайте статью
[11](http://vikos.lrn.ru//kos.php?name=papers/usb/USB-Lin.html)

## Как отключить и подключить зависшее USB-устройство из консоли?

Для этого достаточно отключить на время питание устройства:

```
    echo suspend > /sys/bus/usb/devices/«номер устройства»/power/level
    sleep 10
    echo on > /sys/bus/usb/devices/«номер устройства»/power/level
```

## Как проверить жесткий диск на наличие бэд-блоков из линукса?

smartctl -l selftest /dev/hda

Посмотреть результат: smartctl -A /dev/hda

Так же можно воспользоваться программой badblocks (man badblocks)

## Что делать, если на жестком диске появились бэд-блоки?

Почитайте статью [Bad block HOWTO for
smartmontools](http://smartmontools.sourceforge.net/badblockhowto.html),
там подробно описаны следующие моменты:

  - проверку диска на бэд-блоки (smartctl)
  - ремонт файловых систем, пострадавших от появления бэд-блоков
    (ext2/3, ReiserFS
  - ремонт LVM
  - и прочее...

А вообще, если есть возможность, рекомендуется загрузить программу
[Victoria](http://www.mhdd.ru/download.shtml) (там есть iso-образ), и
работать уже ей.

## Как настроить и использовать USB-вебкамеру в Linux?

Общаться можно с помощью gnomemeeting (
[12](http://www.gnomemeeting.org/index.php?rub=2&pos=0))

О настройке написано здесь - [13](http://www.aboutdebian.com/webcam.htm)
. Для тех, у кого большие сложности с английским, быстро перескажу
(jackill):

Нужно собрать (если нет) следующие модули:

videodev (иными словами video4linux или 4vl); usbcore; input; usb-uhci
(если не получится установить этот модуль, попробуйте usb-ohcl или
uhcl); и модуль поддержки камеры ibmcam для камер ibm, ov511 (например,
для камер Creative WebCam III) или dc2xx (для камер Kodak).

Далее разрешаем всем обращаться к видео-устройству: chmod 666
/dev/video0

Ставим xawtv и перегружаемся.

После перезагрузки, если нет ошибок, запускаем xawtv. В TV-norm
указываем вид сигнала с вашей камеры (PAL/NTSC), на предложение
просканировать диапазон отвечаем нет (а в последних версиях xawtv эту
возможность вообще убрали).

По большой и светлой идее теперь мы должны получить изображение с камеры
в окне xawtv.

## Как включить поддержку MiniDV видеокамеры?

Поддержка включается так:

    IEEE 1394 (FireWire) support
     [M] IEEE 1394 (FireWire) support
     [M]   OHCI-1394 support
     [M]   OHCI-1394 Video support
     [M]   OHCI-DV I/O support
     [M]   Raw IEEE1394 I/O support

О том, как с ней работать, читайте соответствующий вопрос.

## Почему на новом ядре (начиная с 2.6.9) у меня медленно читаются данные с фотоаппарата/usb-flash карты?

Потому что у вас собрана поддержка Low Performance USB Block driver.

Отключите ее и пересоберите ядро:

Device Drivers ---\> Block devices ---\> Low Performance USB Block
driver

## Как заставить работать привод CD-RW/DVD-RW?

В современных дистрибутивах поддержка CD-RW/DVD-RW идет прямо из
коробки, как и программы для записи дисков. Тем не менее, если вы
сами решили настроить и собрать ядро, или что-то пошло не так, полезно
знать, что и где подкрутить, чтобы ваш привод мог писать диски.

1\. Для ядер серии 2.6.х эмуляция scsi для поддержи записи не требуется
и не работает. Поэтому ниже приводится сборка ядер серии 2.4.х:

    Раздел ATA/IDE/MFM/RLL support, IDE, ATA and ATAPI Block devices:

    <M>   Include IDE/ATAPI CDROM support

    <M>   SCSI emulation support

    Раздел SCSI support:

    <M> SCSI support
     --- SCSI support type (disk, tape, CD-ROM)
    <M>   SCSI disk support
    (40) Maximum number of SCSI disks that can be loaded as modules
    < >   SCSI tape support
    < >   SCSI OnStream SC-x0 tape support
    <M>   SCSI CD-ROM support
    [*]     Enable vendor-specific extensions (for SCSI CDROM)
    (8) Maximum number of CDROM devices that can be loaded as modules
    <M>   SCSI generic support
     --- Some SCSI devices (e.g. CD jukebox) support multiple LUNs
    [*]   Enable extra checks in new queueing code
    [*]   Probe all LUNs on each SCSI device
    [*]   Verbose SCSI error reporting (kernel size +=12K)
    [ ]   SCSI logging facility
    SCSI low-level drivers  --->

2\. Собираем и ставим в систему пакеты. Они входят в состав любого
дистрибутива, кроме узкоспециализированных на работу с сетью:

cdda2wav cdrdao cdrecord mkisofs

Это программы необходимы для создания образов дисков и записи.

3\. Для ядер серии 2.4.х нужно передать в загрузчик строку о том, что
наш привод работает как scsi:

а) если у нас стоит lilo, то в /etc/lilo.conf добавляем строчку

    /etc/lilo.conf:

    append="hdX=ide-scsi"

где X - буква вашего привода и выполняем lilo,

б) если у нас стоит grub, то в /etc/grub.conf добавляем напротив
названия ядра строчку hdX=ide-scsi, например, вот так:

    /etc/grub.conf:

    kernel /boot/vmlinuz-2.4.25 ro hdc=ide-scsi

Чем записывать, читайте в вопросе [Какие есть программы для записи
CD-R/CD-RW/DVD?](http://www.linux.org.ru/wiki/en/Поиск_ПО#%D0%9A%D0%B0%D0%BA%D0%B8%D0%B5_%D0%B5%D1%81%D1%82%D1%8C_%D0%BF%D1%80%D0%BE%D0%B3%D1%80%D0%B0%D0%BC%D0%BC%D1%8B_%D0%B4%D0%BB%D1%8F_%D0%B7%D0%B0%D0%BF%D0%B8%D1%81%D0%B8_CD-R/CD-RW/DVD%3F)

Как записывать, читайте в вопросе [Как записать компакт-диск
(CD-R/CD-RW/DVD-RW)](http://www.linux.org.ru/wiki/en/Диски%2C_приводы_CD/DVD-RW_и_файловые_системы#%D0%9A%D0%B0%D0%BA_%D0%B7%D0%B0%D0%BF%D0%B8%D1%81%D0%B0%D1%82%D1%8C_%D0%BA%D0%BE%D0%BC%D0%BF%D0%B0%D0%BA%D1%82-%D0%B4%D0%B8%D1%81%D0%BA_\(CD-R/CD-RW/DVD-RW\)%3F)

Дополнительные статьи на тему настройки и записи дисков:

[14](http://linuxshop.ru/linuxbegin/article307.html)

[15](http://www-106.ibm.com/developerworks/linux/library/l-cdburn.html?ca=dgr-lnxw16BurnCDs)

[16](http://linuxdoc.ru/HOWTO/html/CD-Writing-HOWTO.html)

## Можно ли под Linux'ом или Mac OS X перешить привод NEC?

Да. Есть утилита Binflash - [17](http://binflash.cdfreaks.com/)

## Где найти драйверы под модем (winmodem, HCF, HSF)? Какой стоит купить модем?

Если модем аппаратный т.е. не является win-модемом, то драйверы для него
не нужны - достаточно поддержки com-порта, собранной в ядре (по
умолчанию у всех есть). Это старые модемы, работающие через
шину ISA, некоторые новые и дорогие на pci (часто имеют в своем
составе абревиатуру HCF), внешние подключаемые через com-порт, и
некоторые подключаемые через USB (обычно имеют аналогичную модель,
подключаемую через компорт).

Если модем программный, смотрите на каком чипе он собран (можно найти в
документации или сразу посмотреть на маркировку микросхемы).

После чего отправляйтесь на сайт
[www.linmodems.org](http://www.linmodems.org) и посмотрите список
поддерживаемых модемов - есть ли среди них ваш модем, или чип как
у него.

**Лучше всего покупать внешние аппаратные модемы, подключаемые через
com-порт**

Ниже в этом faq даны инструкции по установке некоторых модемов.

## Как установить винмодем Modem: Intel Corp. 82801DB (ICH4) AC'97 на ядро 2.6.х?

Речь идет о модеме smartlink. Согласно lspci чип выглядит так:

    00:1f.6 Modem: Intel Corp. 82801DB (ICH4) AC'97 Modem Controller (rev 03)

Для установки драйвера под данный модем понадобится следующее ПО:
alsa-driver не ниже версии 1.0.3 ( [18](http://www.alsa-project.org))
slmodem версии 2.9.6 ( [19](http://www.smlink.com))

Часть 1 - ядерная.

В последних версиях драйверов alsa есть драйвер snd-intel8x0m (буква m
принципиальна), который нам и нужен. В alsa'е, включенной в ядра 2.6.х
(на момент 2.6.3) этого модуля нет, поэтому его нужно добавить туда
самому.

Итак, берем alsa-drivers, и slmodem-2.9.6.tar.gz и распаковываем в
подходящие для этого директории.

Далее накладываем на ядро патч, поставляемый в составе slmodem. Если
ядро из поставки RedHat, то придется подредактировать патч.
Исправленная версия есть здесь -
[20](http://fine.kalinovka.net/articles/hobby/lorFAQ/files/alsa-linux-2.6.0.patch.gz)

Если у нас ядро версии ниже 2.6.5, то заменяем файл intel8x0m.c на более
свежий из alsa-driver - иначе просто не соберется.

В конфиге ядра включаем соответствующий модуль
(CONFIG_SND_INTEL8X0M=m).

Собираем ядро.

Часть 2 - пользовательская.

После чего заменяем в modprobe.conf строчку

    alias sound-slot-0 snd-intel8x0

на

    alias sound-slot-0 snd-intel8x0m.

Перегружаемся.

Далее заходим в slmodem и делаем "make SUPPORT_ALSA=1". Получаем
приложение slmodemd.

Как только мы его запускаем, автоматически появляется устройство
/dev/ttySL0 :maj 136, min 1 (точнее, /dev/pts/1 т.к. /dev/ttySL0 -
только ссылка). Это и есть наш винмодем.

За инфо благодарим Sergey V. Udaltsov AKA svu и Sasha Khapyorsky из
smlink.com

## Как установить драйверы к модему на чипсете Lucent?

Для серии ядер 2.6:
[21](http://linmodems.technion.ac.il/packages/ltmodem/kernel-2.6/martian)

## Как установить USB-модем ZyXEL Omni 56K (ZyXEL Omni 56K PLUS и DUO, ZyXEL Omni 56K UNO)?

Прочитать об этом можно на страничке McMCC
[22](http://mcmcc.bat.ru/omniusb/index.html)

## Как установить Zyxel OMNI ADSL USB (СТРИМ)?

Качаем libatm-2.4.1( [23](http://linux-atm.sourceforge.net/)) и cxacru (
[24](http://accessrunner.sourceforge.net), собственно сам драйвер)

Далее нужно настроить поддержку atm-а в ядре как написано в how-to к
libatm ( [25](http://www.tldp.org/HOWTO/ATM-Linux-HOWTO/)) и собрать
ядро и libatm.

Перегружаемся с новым ядром.

Осталось собрать cxacru. Для начала надо его пропатчить таким патчем:

    diff -urN cxacru-orig/init/cxioctl.c cxacru/init/cxioctl.c
    --- cxacru-orig/init/cxioctl.c  2004-05-11 11:05:25.000000000 +0400
    +++ cxacru/init/cxioctl.c       2004-09-27 23:13:43.000000000 +0400
    @@ -295,6 +295,9 @@
       /* Vendor = Zoom, Product = 5510 */
       else if (vid == 0x1803 && pid == 0x5510)
         return 6;
    +  /* zyxel omni */
    +  else if (vid == 0x0586 && pid == 0x330a)
    +    return 5;

       return -1;
     }
    diff -urN cxacru-orig/init/cxload.c cxacru/init/cxload.c
    --- cxacru-orig/init/cxload.c   2004-05-11 11:05:25.000000000 +0400
    +++ cxacru/init/cxload.c        2004-09-27 23:16:59.000000000 +0400
    @@ -1136,6 +1136,10 @@
       /* Vendor = Zoom, Product = 5510 */
       else if (vid == 0x1803 && pid == 0x5510)
         return 6;
    +  /* zyxel omni */
    +  else if (vid == 0x0586 && pid == 0x330a)
    +    return 5;
    +

       return -1;
     }
    diff -urN cxacru-orig/module2/xdslusb.c cxacru/module2/xdslusb.c
    --- cxacru-orig/module2/xdslusb.c       2004-05-11 11:05:26.000000000 +0400
    +++ cxacru/module2/xdslusb.c    2004-09-27 23:18:54.000000000 +0400
    @@ -154,6 +154,9 @@
     #define CXACRU_PRODUCTID8               0x5510  /* Product = 5510 */
     #define CXACRU_VENDORID9                0x0675  /* Vendor = Draytek */
     #define CXACRU_PRODUCTID9               0x0200  /* Product = Vigor 318 */
    +
    +#define ZYXEL_VID 0x0586
    +#define ZYXEL_PID 0x330a

     /* 3Com reference design (Alcatel DSP) */
     #define CP4218_VENDORID                 0x0506  /* Vendor = 3Com */
    @@ -238,6 +241,7 @@

     #define hex2int(c) ( (c >= '0') && (c <= '9') ? (c - '0') : ((c & 0xf) + 9) )

    +
     static struct usb_device_id udsl_usb_ids [] = {
            { USB_DEVICE (SPEEDTOUCH_VENDORID, SPEEDTOUCH_PRODUCTID) },
            { USB_DEVICE (AME_VENDORID, AME_PRODUCTID) },
    @@ -251,6 +255,7 @@
            { USB_DEVICE (CXACRU_VENDORID8, CXACRU_PRODUCTID8) },
            { USB_DEVICE (CXACRU_VENDORID9, CXACRU_PRODUCTID9) },
            { USB_DEVICE (CP4218_VENDORID, CP4218_PRODUCTID) },
    +       { USB_DEVICE (ZYXEL_VID, ZYXEL_PID) },
            { }
     };

    @@ -1238,6 +1243,9 @@
            else if (vid == CP4218_VENDORID && pid == CP4218_PRODUCTID && cl == USB_CLASS_VENDOR_SPEC && ifn == 0)
                    return UDSL_MODEM_TYPE3;

    +       else if (vid == ZYXEL_VID && pid == ZYXEL_PID && cl == USB_CLASS_VENDOR_SPEC && ifn == 0)
    +               return UDSL_MODEM_TYPE2;
    +
            return -1;
     }

Дальше скомпилировать, он заодно и установится.

В /etc он засунет файл cxacru.conf

    #
    # Config file for Conexant AccessRunner
    #

    # Driver mode
    DRIVER_MODE=1  # 1 = normal, 2 = debug, 3 = normal+max speed (without ask adsl status), 4 = debug+max speed (without ask adsl status)

    # Protocol
    PROTOCOL_MODE=4  # 1 = RFC1483/2684 routed, 2 = PPP over ATM (pppoa), 3 = RFC1483/2684 bridged, 4 = PPP over Ethernet (pppoe)

    # Paths
    BINARY_PATH="/usr/sbin"
    ATM_PATH=""

    # ADSL
    #  if OPEN_MODE is blank then cxload uses default mode acoording VID & PID
    #  Values for OPEN_MODE are:
    #    0 = auto selection, G.Handshake
    #    1 = auto selection, T1.413
    #    2 = G.Handshake
    #    3 = ANSI T1.413
    #    4 = ITU-T G.992.1 (G.DMT)
    #    5 = ITU-T G.992.2 (G.LITE)
    OPEN_MODE=3

    # ATM
    VPI=1
    VCI=50

    # Specific for RFC1483/2684 routed/bridged
    #  if IP_ADDRESS is blank in bridged mode then it uses DHCP to get IP
    IP_ADDRESS=
    NETMASK=
    GATEWAY=

Затем нужно настроить pppoe на интерфейс nas0 и запустить cxstart.sh. У
меня (no1sm) slackware и pppoe настраивается командой adsl-setup из
пакета rp-pppoe.

Скорее всего придется еще что-нибудь подкрутить,но вобщем эта
конфигурация должна быть рабочей.

## Как установить модем AusLinx 2006? (СТРИМ)

Так же, как предыдущий, только нужно переписать firmware. Как это
проделать читайте здесь:

[26](http://www.evil-and.nm.ru/)

Для ядер 2.6.10 чуть измененная инструкция
[27](http://www.evil-and.nm.ru/instruction-2.6.10.html)

## Как настроить USB-мышь/скролл у USB-мышки?

У большинства людей устройство скорее всего будет называться
/dev/input/mice

Для подключения должно быть соответствующая нода. В большинстве
дистрибутивов она есть. Если нет, ее можно создать:

    mkdir /dev/input
    mknod /dev/input/mice c 13 63

Ссылка на хорошее английское руководство -
[28](http://www.linux-usb.org/USB-guide/x194.html)

Ядро должно быть собрано с опциями:

    /usr/src/linux/.config:

    CONFIG_INPUT                    #Обязательно.

    CONFIG_INPUT_MOUSEDEV   #Для использования USB-мыши.

    CONFIG_USB                              #Для использования USB-устройств вообще.

    CONFIG_USB_DEVICEFS             #После этого появится /proc/bus/usb/devices, где видно что подключено.

    CONFIG_USB_HID                  #Тоже нужно.

    CONFIG_USB_HIDINPUT             #Аналогично.

    CONFIG_USB_UHCI                 #Для компьютеров с материнками на базе чипсетов от Intel (intel 430TX, 440FX, 440LX, 440BX, i810, i820), VIA (VIA VP2, VP3, MVP3, Apollo Pro, Apollo Pro II or Apollo Pro 133).

    #
    # ИЛИ
    #

    CONFIG_USB_OHCI                 #Для SiS или ALi (ALi IV, ALi V, Aladdin Pro)

Если включена поддержка usbfs, ее нужно прописать в /etc/fstab:

    /etc/fstab:

    none  /proc/bus/usb   usbfs defaults  0   0

Далее - настройка программ, которым нужна мышь.

Для gpm:

    /etc/gpm.conf:

    device=/dev/input/mice
    responsiveness=
    repeat_type=
    type=autops2
    append=""
    sample_rate=

Для X-сервера (xorg, XFree86), в /etc/X11/xorg.conf (/etc/X11/XF86Config
или /etc/X11/XF86Config-4):

    XF86Config:

    Section "InputDevice"
            Identifier  "Configured Mouse"
            Driver      "mouse"
            Option      "CorePointer"
            Option      "Device" "/dev/input/mice"
            Option      "Protocol" "ImPS/2"
            Option      "ZAxisMapping" "4 5"
    EndSection

За ответы спасибо Zulu.

## У меня проблема с клавиатурой/мышью при использовании ядра 2.6.x. Как исправить?

[29](http://www.kerneltrap.org/node/2199)

И вопрос Как перейти на ядро 2.6? У меня не грузятся модули на ядре 2.6,
что делать? (QM_MODULES)

## Где можно найти материалы по подключению мобильных устройств к Linux?

[30](http://www.tuxmobile.org)

## Как подключить кардридер?

[При всех телодвижениях это не
сложно.](http://hot-orange.narod.ru/chtivo/card-reader.htm)

Но возможно Вы встретите некорректно работающий кардридер из-за ошибки в
железе или ядре.

Это можно определить по следующим признакам:

1.  Команда `dmesg` показывает что ядро нашло кардридер и слоты в нём.
2.  Команда `cat /proc/scsi/scsi` показывает что кардридер нормально
    сконфигурирован

<!-- end list -->

    Host: scsi9 Channel: 00 Id: 00 Lun: 00
      Vendor: Generic  Model: USB SD Reader    Rev: 1.00
      Type:   Direct-Access                    ANSI  SCSI revision: 00
    Host: scsi9 Channel: 00 Id: 00 Lun: 01
      Vendor: Generic  Model: USB CF Reader    Rev: 1.01
      Type:   Direct-Access                    ANSI  SCSI revision: 00
    Host: scsi9 Channel: 00 Id: 00 Lun: 02
      Vendor: Generic  Model: USB SM Reader    Rev: 1.02
      Type:   Direct-Access                    ANSI  SCSI revision: 00
    Host: scsi9 Channel: 00 Id: 00 Lun: 03
      Vendor: Generic  Model: USB MS Reader    Rev: 1.03
      Type:   Direct-Access                    ANSI  SCSI revision: 00

1.  Но втыкание карточки не вызывает ни строчки реакции со строны
    `dmesg`.

Выполните команду `modprobe -r ehci_hcd` от пользователя имеющего права
root.

Что бы не повторять все действия вручную, занесите модуль в чёрный
список `echo "blacklist ehci_hcd" >>
/etc/modprobe.d/blacklist.conf`.

## Кардридер читает только первый слот, что делать?

В вашем дистрибутиве чуть-чуть недокрутили настройки ядра. Дайте
команду:

    echo "scsi scsi-add-single-device <scsi instance> <scsi channel> <scsi id> <scsi lun>">/proc/scsi/scsi

Например, для шестипортового кардридера:

    root@localhost# echo "scsi add-single-device 1 0 0 5" > /proc/scsi/scsi
    root@localhost# echo "scsi add-single-device 1 0 0 4" > /proc/scsi/scsi
    root@localhost# echo "scsi add-single-device 1 0 0 3" > /proc/scsi/scsi
    root@localhost# echo "scsi add-single-device 1 0 0 2" > /proc/scsi/scsi
    root@localhost# echo "scsi add-single-device 1 0 0 1" > /proc/scsi/scsi
    root@localhost# echo "scsi add-single-device 1 0 0 0" > /proc/scsi/scsi

## Я не понимаю, как заставить работать свою звуковую карту

Ну тогда прочитайте
[mini-howto](http://www.linuxcenter.ru/lib/articles/system/alsa_minihowto.phtml)
на русском языке и вам все станет понятно.

## Что такое OSS? Чем OSS лучше/хуже ALSA? Как установить и настроить OSS?

Как-то [так](http://www.linux.org.ru/wiki/en/OSS)

## Как настроить alsa?

Заходим на [31](http://www.alsa-project.org/alsa-doc/)

Выбираем нужного производителя. В открывшемся окне щелкаем на название
модуля. Читаем что и куда прописывать.

Для большинства карт пойдет такой способ:

В файле **/etc/modprobe.conf** (для 2.6.х) или **/etc/modules.conf**(для
2.4.х):

    # ALSA native device support
    alias char-major-116 snd
    alias snd-card-0 _модуль_вашей_звуковой_
    alias snd-slot-0 snd-card-0
    options snd major=116 cards_limit=1

    # OSS/Free setup
    alias char-major-14 soundcore
    alias sound-service-0-0 snd-mixer-oss
    alias sound-service-0-1 snd-seq-oss
    alias sound-service-0-3 snd-pcm-oss
    alias sound-service-0-8 snd-seq-oss
    alias sound-service-0-12 snd-pcm-oss

    #Эти две строчки нужны только если у вас работает udev (fedora core 3, gentoo, debian, asp linux 10
    install snd-pcm modprobe -i snd-pcm ; modprobe snd-pcm-oss ; true
    install snd-seq modprobe -i snd-seq ; modprobe snd-seq-oss ; true

Модуль карты в случае sb life, например, это snd-emu10k1 (т.е. к
названию модуля прибавляется приставка snd-) Точные названия
модулей вы увидите в
**/lib/modules/версия_вашего_ядра/kernel/sound/pci**

Для проверки можете сделать **modprobe название_вашей_звуковой**.

Далее запускаете alsamixer и включаете звук (по умолчанию выключен).

Сохраняете настройки карты **/sbin/alsactl store**

Перегружаетесь и проверяете, получилось ли у вас - по команде
**/sbin/lspci** должен появиться список модулей, среди которых должна
быть ваша карта.

## Поставил альсу (alsa), а звука нет. Что делать?

По умолчанию после установки alsa (если ранее она не стояла), все каналы
находятся в заглушенном положении (mute). Поэтому запустите alsamixer и
снимите mute клавишей m с каналов Master и PCM.

Кнопки курсора вверх и вниз позволят установить требуемый уровень звука.

Сохраняются настройки командой alsactl store

Ещё возможен вариант, что звук уходит в HDMI, который ни к чему не
подключен. В системах с PulseAudio можно переключить устройство
вывода в настройках. В голой alsa (например, при использовании xfce
или лёгких оконных менеджеров) посмотреть список устройств можно
командами aplay -l и aplay -L, настроить в \~/.asoundrc для
пользователя и в /etc/asound.conf для всей системы следующим
образом:

    pcm.!default {
       type hw
       card 1
       device 0
    }
    ctl.!default {
       type hw
       card 1
       device 0
    }

номера нужных card и device можно найти в выводе предыдущих команд.

если нужен софтмикшер, то предыдущий пример надо исправить до
следующего:

    pcm.!default {
        type plug
        slave.pcm "dmixer"
    }

    pcm.dmixer  {
        type dmix
        ipc_key 1024
        slave {
            pcm "hw:1,0"
            period_time 0
            period_size 1024
            buffer_size 4096
            rate 44100
        }
        bindings {
            0 0
            1 1
        }
    }

    ctl.dmixer {
        type hw
        card 1
        device 0
    }

## Как сохранить настройки микшера при использовании ALSA? Как сделать, чтобы они восстанавливались?

Для сохранения настроек выполните команду

    user@linux# /usr/sbin/alsactl store

А чтобы они восстанавливались (и записывались) нужно, чтобы в ваших
инициализационных скриптах стартовал демон alsasound.
Устанавливается он вместе с alsa-driver.

В дистрибутивах Fedora Core 1 и 2 запись и восстановление настроек
прописано двумя строчками в /etc/modules.conf
(/etc/modprobe.conf), например:

    install snd-intel8x0 /sbin/modprobe --ignore-install snd-intel8x0 && /usr/sbin/alsactl restore >/dev/null 2>&1 ¦¦ &colon;
    remove snd-intel8x0 { /usr/sbin/alsactl store >/dev/null 2>&1 ¦¦ &colon; &colon; }&semi; /sbin/modprobe -r --ignore-remove snd-intel8x0

Удобнее убрать эти строчки и установить демон alsasound (его можно
взять, например, в пакете alsa-driver с
[www.alsaproject.org](http://www.alsaproject.org)). Это позволит
нормально выгружать драйверы в случае необходимости. Однако в
случае Fedora Core 3 этого делать не нужно - демон alsasound не
сработает.

## Где в Linux эквалайзер? Как мне поднять/прибрать высокие/низкие частоты для всех аудио-приложений?

Этот вопрос задается регулярно. Наиболее простой способ, если вы
пользуетесь ALSA, [описан
здесь](http://www.thedigitalmachine.net/alsaequal.html). В звуковых
картах Creative Labs (семейства Live\!, Audigy, X-Fi) встроен
двухполосный аппаратный эквалайзер, но пользоваться им не
рекомендуется из-за ощутимого ухудшения звучания при его
использовании.

## Когда у меня включен xmms/mplayer/другая программа, звук с kde/gnome или других программ не выводится. Приложение польностью занимает звуковую карту, что делать? Как заставить приложения в KDE воспроизводить звук одновременно? (устройство вывода занято)?

Начнем с карты. Раз такое происходит, значит ваша карточка не умеет
аппаратно смешивать звуковые потоки. Лучше бы от нее избавиться, а
если она встроенная, то купить внешнюю. Sb live 5.1 стоит копейки. Но
не суть.

В kde и gnome существуют так называемые звуковые серверы. Звуковые
серверы занимаются смешиванием звуковых потоков. Эти серверы в
случае выведения звука занимают звуковую карту и по большой и
светлой идее все программы должны работать через них.

Но вот незадача - если программа не умеет работать через звуковой
сервер, то она сама занимает карточку и после этого сервер уже не
может ею воспользоваться, равно как и любая другая программа.

Есть два решения этой проблемы - для xmms скачать (если не включен в
пакет) плагин xmms-arts или xmms-esd - первый для KDE, второй для
Gnome. В mplayer указать вывод звука -ao arts и так далее.

Можно сделать проще - отключить эти серверы и воспользоваться
программным микшером самой alsa (у кого OSS - это их
проблемы). Для того, чтобы проигрывать системные звуки KDE или
Gnome нужно указать проигрывать их внешней программой (например, play из
пакета sox).

Если вам необходим качественный звуковой сервер (вы занимаетесь
музыкой), то не повредит сочетание программного микшера и
звукового сервера [jack](http://jackit.sourceforge.net/),
который специально писался для работы со звуком.

## Как сделать программное микширование с ALSA на картах, не поддерживающих аппаратное?

Использовать dmix. Для alsa \> 1.0.9pre2 делать то, что написано ниже,
необязательно, т.к. alsa сама определяет, способна ли карта
микшировать аппаратно и если нет, включает dmix.

Кидаем в /etc такой конфиг (назвается asound.conf)

    /etc/asound.conf:

    pcm.!default {
        type plug
        slave.pcm "dmixer"
    }

    pcm.dsp0 {
        type plug
        slave.pcm "dmixer"
    }

    pcm.dmixer  {
        type dmix
        ipc_key 1024
        slave {
            pcm "hw:0,0"
            period_time 0
            period_size 1024
            buffer_size 8192
            #buffer_size 32768
            #periods 128
            rate 44100
        }
        bindings {
            0 0
            1 1
        }
    }

    ctl.mixer0 {
        type hw
        card 0
    }

За ответ спасибо McMcc. Готовый файл лежит на
[32](http://mcmcc.bat.ru/mypatches/asound.conf)

## Что за ошибка FATAL: Error running install command for sound_slot_0?

Есть три варианта - простой, сложный, очень сложный.

  - **Простой:** в /etc/modprobe.conf (или /etc/modules.conf для ветки
    ядер 2.4.х) не хватает упоминания snd-slot-0. Правильно написать
    так:

<!-- end list -->

```
      alias snd-slot-0 snd-card-0
```

  - **Сложный:** Невозможно загрузить модуль звуковой карты и как
    следствие, нельзя к нему обратиться. Может перепутано имя,
    может карта неплотно стоит в слоте (отвалилась, такое бывает).
  - **Очень сложный:** Вы работаете в дистрибутиве с поддержкой selinux
    (например, fedora core) и последний включен. Либо отключите его,
    либо настраивайте разрешения.

## Как автоматизировать работу MIDI-клавиатуры?

Во-первых, для этого нам необходимо будет заставить aconnect соединять
выходной MIDI-порт клавиатуры и входной MIDI-порт синтеза, для этого
создадим такое правило udev:

    # автозапуск aconnect'а при подключении клавы
    SUBSYSTEM!="sound", GOTO="my_usb_audio_end"
    ACTION!="add", GOTO="my_usb_audio_end"
    KERNEL=="midi[0-9]", DRIVERS=="usb", RUN+="/usr/bin/aconnect 20:0 17:0"
    LABEL="my_usb_audio_end"

Порты для aconnect, которые нужно записать в ключе RUN, можно посмотреть
при помощи команды `aconnect -io` (она отобразит все открытые порты
ввода/вывода MIDI).

Чтобы аппаратный синтез на карточках типа SoundBlaster работал, при
запуске компьютера необходимо добавить запуск jack (`jackd -d
alsa`), а также подгрузить "звуковой шрифт" командой `asfxload`.
Выполнение этих двух команд можно поместить в стартовый скрипт
`rc.local`, который в различных дистрибутивах может находится в разных
местах (например /etc/rc.local в Debian или /etc/rc.d/boot.local в
openSUSE)

## Как подключить wireless-карточку?

Желающим сделать это - прямая дорожка на Yandex.ru & Google.com со
словами "настройка wireless linux" на устах. Проверено -
информации вполне достаточно.

Для начала следует сходить сюда [33](http://acx100.sourceforge.net/)

Пример для Planet WL-8305. Эта карта собрана на базе чипа acx100.

Для начала смотрим что у нас есть:

    root@linux:~# lspci
    ...
    00:0a.0 Network controller: Texas Instruments ACX 100 22Mbps Wireless Interface
    ...

Да, устройство с чипом acx100 в системе имеется. Идем дальше.

Скачиваем модуль -
[34](http://www.houseofcraig.net/old_acx100_howto.php#downloads)

    root@fhome:~# cd /lib/modules/
    root@fhome:/lib/modules# mkdir acx100_fmwe
    root@fhome:/lib/modules# cd ./acx100_fmwe/
    root@fhome:/lib/modules# tar -xvzf ~/acx100/acx100_firmware.tar.gz
    root@fhome:/lib/modules/acx100_fmwe# cd ../2.4.22/kernel/drivers/net/
    root@fhome:/lib/modules/2.4.22/kernel/drivers/net# tar -xvzf ~/acx100/acx100_pci-v0.2.0pre7-2.4.22.tar.gz

Этими командами мы установили в систему модуль acx100_pci, собранный
под ядро 2.4.22 и firmware, необходиое для его работы.

Для того, чтоб модуль загружался во время старта правим
/etc/modules.conf.

    /etc/modules.conf:
    alias eth0 acx100_pci
    options acx100_cpi use_eth_name=1 debug=0x01 formware_dir=/lib/modules/acx100_fmwe
    post-install acx100_pci /etc/rc.d/rc.acx100_pci

Первая строка - объявляем псевдоним eth1 для беспроводного устройства (у
меня eth0 уже есть)

Вторая - устанавливаем параметры, без которых модуль не загрузится.

Третья - указываем, что после загрузки модуля нужно запустить скрипт
/etc/rc.d/rc.acx100_pci, в котором производится настройка беспроводного
соединения.

Далее:

    root@fhome:~# depmod -a
    root@fhome:~# grep acx100_pci /lib/modules/2.4.22/modules.pcimap
    acx100_pci           0x0000104c 0x00008400 0xffffffff 0xffffffff 0x00000000 0x00000000 0x00000000
    acx100_pci           0x0000104c 0x00008401 0xffffffff 0xffffffff 0x00000000 0x00000000 0x00000000
    acx100_pci           0x0000104c 0x00009066 0xffffffff 0xffffffff 0x00000000 0x00000000 0x00000000

Модуль нормально установился. Пробуем его загрузить:

    root@fhome:~# modprobe acx100_pci

Если получилось что-то такое:

    /lib/modules/2.4.22/kernel/drivers/net/acx100_pci.o: init_module: No such device/lib/modules/2.4.22/kernel/drivers/net/acx100_pci.o: Hint: insmod errors can be
    caused by incorrect module parameters, including invalid IO or IRQ parameters.
          You may find more information in syslog or the output from dmesg
    /lib/modules/2.4.22/kernel/drivers/net/acx100_pci.o: insmod /lib/modules/2.4.22/kernel/drivers/net/acx100_pci.o failed
    /lib/modules/2.4.22/kernel/drivers/net/acx100_pci.o: insmod acx100_pci failed

Это либо нету устройства в системе, либо в /etc/modules.conf ошибка и
"options acx100_pсi" это "options acx100_cpi" например, или... что-то
еще :)

Если все прошло нормально, проверяем, загрузился ли модуль:

    root@fhome:~# lsmod|grep acx100_pci
    acx100_pci            124512   0  (unused)

## Как узнать температуру процессора, чипсета? Как настроить lm_sensors?

Нужно установить и настроить lm sensors.

Пакет lm_sensors входит практически в любой дистрибутив. Документация к
нему идет подробная.

Чтобы установить [lm_sensors](http://www2.lm-sensors.nu/~lm78/) в общем
случае нужно:

  - Собрать ядро с поддержкой lm_sensors (пропатчить при необходимости
    - патчи есть на сайте), т.е. включить поддержку i2c, smbus и
    датчиков (если не знаете какой у вас датчик, собирайте все).
  - Собрать и установить сам пакет lm_sensors (или поставить готовый из
    своего репозитария).
  - Выполнить в командной строке sensors-detect и проделать все, что
    говорит эта программа.

Температуру в консоли можно посмотреть, набрав sensors.

В графическом режиме температуру можно снять с помощью \[gkrellm\].
Список программ и расширений DE/WM, работающих с lmsensors, можно
прочитать здесь.

## Могу ли я посмотреть температуру своего жесткого диска?

Конечно. Утилита называется
[hddtemp](http://www.guzu.net/linux/hddtemp.php). Прямо на сайте помимо
тарбола выложены пакеты для Gentoo, Debian, Alt Linux, Redhat/Fedora
Core, Slackware.

Там же можно взять плагин для gkrellm, показывающий температуру.

## Как и чем измерить температуру воздуха, например, в помещении?

Я так понимаю, что хочется свой датчик и програмку, которая бы снимала
температуру с него.

[35](http://www.kusto.com.ru/temperature/)

## Чем протестировать оперативную память?

Программой [memtest86](http://www.memtest86.com/). Обычно она входит в
состав большинства дистрибутивов, а также в rescue-часть.

Дополнительная ссылка на memtest - [36](http://www.memtest.org/)

## Система не видит всю память. У меня 1Gb, а видно только 883Mb. Что делать?

Стандартное решение простое - в конфиге ядра включить поддержку большого
объема памяти: processor type and features -\> high memory support

Или пропатчить специальным патчем с названием
**1g_lowmem1_i386.diff**. Взять его можно, например, отсюда
[37](http://ck.kolivas.org/patches/2.6/).

В первом случае будет некоторый оверхед. Во втором случае его не будет,
но некоторые программы, типа vmware прекратят работу.

## Система не видит всю память. У меня 4Gb, а видно только 3Gb. Что делать?

processor type and features -\> high memory support -\> 64G

Также иногда необходимо включить в BIOS компьютера опцию "Memory remap".
На современных материнских платах эта опция есть не во всех BIOS, но,
как правило, если этой опции нет, то это означает что нужное
поведение автоматически активируется когда система определяет
объем установленной оперативной памяти в 4 и более гигабайта.

## Как протестировать процессор под нагрузкой (например, если он разогнан)?

Есть программа под названием
[cpuburn](http://pages.sbcglobal.net/redelm/) и еще одна
[cpuburn-in](http://users.bigpond.net.au/cpuburn/).

Читаете ридми, собираете, запускаете. Ждете минут сорок.

Есть тест прошел, запускаете quake 3, который известен вылетами на
нестабильном железе в linux - минут сорока должно хватить. Обычно
вылетает на пятнадцатой-двадцатой минуте.

Далее ставите на сборку что-нибудь, что у вас раньше нормально
собиралось. Сначала ядро. Потом какой-нить kdelibs.

Если не вылетит, тогда можно на случай атомной войны попробовать пожать
что-нить огромное rar'ом.

Если все это работает, значит система стабильна.

## Какую взять видеокарточку под linux - ATI или Nvidia?

У nvidia есть прилично работающий закрытый драйвер, который более-менее
стабильно работает на большинстве стандартных ядер (всякие
нестандартные патчи, вроде поддержки realtime, могут
вогнать его в ступор), иногда показывает даже большую
производительность, чем в windows, поддерживает все функции
карточки. Открытый драйвер nv поддерживает только 2D, открытый драйвер
Nouveau позволит в дополнение к этому запустить некоторые 3D-программы и
порадоваться, что хоть что-то работает.

27.03.10 Развитие драйвера nv прекращено. Пользователям новых видеокарт
NVIDIA рекомендуется до установки проприетарного драйвера пользоваться
драйвером vesa. [38](http://www.linux.org.ru/news/hardware/4704381)

У ATI (теперь это подразделение AMD) закрытые драйвера. Существует
мнение, что рекордов скорости они не ставят, да и со стабильностью
у них нередко есть проблемы. После покупки ATI компанией AMD появилось
также мнение о том, что закрытый драйвер стал не хуже, чем аналогичный
у Nvidia. На деле проблемы бывают у обеих производителей в равной
степени, а рекомендация покупать видеокарты от Nvidia - не более,
чем традиция. Отчасти мнение о большей стабильности драйверов для Nvidia
основано на том, что тестирование проводят именно на этих видеокартах.
Есть стабильно, но пока ещё не очень быстро работающие с 3D открытые
драйвера, в настоящее время спецификации на карты открыты и активно
ведётся разработка драйверов для новых карт. Посмотреть текущую
ситуацию с поддержкой этих карточек открытыми драйверами можно
здесь: [39](http://www.x.org/wiki/RadeonFeature)

В случае с ноутбуками ситуация оказывается не в пользу Nvidia в любом
случае, так как официальных драйверов с поддержкой технологии Optimus
как не было, так и нет. Для этого существуют проекты bumblebee и primus,
которые имеют свои ограничения по поддерживаемым чипам и способу
рендеринга. С AMD же ситуация на ноутбуках значительно лучше:
есть поддержка переключаемой гибридной графики AMD+AMD и CrossFire,
хотя последнее на деле бесполезно, так как CrossFire включается только
для ограниченного списка избранных приложений.

В общем, если для вас открытость драйвера важнее высокой
производительности 3D-графики --- лучше взять не самую
последнюю карточку от ATI/AMD, если скорость в 3D приложениях вам
нужнее --- тогда имеет смысл купить видеокарту от nvidia. Если же
нужен ноутбук, то в сторону nvidia с технологией optimus смотреть
противопоказано.

## Как определить, какой чип стоит в видеокарте от ATI/AMD?

<http://radeon.ru/reference/cardtable/>

## Как обновить BIOS, имея только Linux и привод cd-rom?

Дано: линукс, сдром. Не дано: виндовс, флопповод.

Идем на [40](http://www.bootdisk.com/bootdisk.htm) и качаем образ
загрузочной дискетки. Я например, скачал этот
[41](http://1gighost.net/randy/622c.zip) распаковываем его и среди
прочих файлов находим файл с расширением .IMG

Монтируем этот файл:

    mount -o loop -t vfat 622C.IMG /mnt/floppy

Удаляем из образа все лишние файлы (например BASIC или драйверы мыши),
чтобы освободить место для файлов BIOS'а и утилиты прошивки:

    rm /mnt/floppy/qbasic.*

Затем копируем образ своего BIOS'а и утилиту прошивки и отмонтируем
образ дискеты:

    cp 18d110.BIN awdflash.exe /mnt/floppy
    umount /mnt/floppy

Теперь осталось создать образ компакт-диска и прожечь его. Образ создаем
примерно так:

    mkdir ~/bootcd # временная директория для загрузочного сд-диска
    cp 622С.IMG ~/bootcd/ # копируем измененный образ дискеты
    mkisofs -r -b 622C.IMG -c boot.cat -o ~/bootcd.iso ~/bootcd

В результате в домашней директории появится файл **bootcd.iso**,
прожигаем его и грузимся.

За ответ благодарим Karmadon.

## Как настроить USB CDMA модем?

CDMA2000 (EVDO) модемы, продаваемые SkyLink (Россия) и PeopleNet
(Украина) работают через модуль ядра usbserial. Всё что нужно
(если модем сам не нашелся) - указать модулю параметры (выдержка из
/etc/modprobe.conf):

    options  usbserial vendor=0x05c6 product=0x6539

Значения брать из вывода утилиты lsusb

hint: pcmcia модемы - тоже usb, на самом деле.

Само подключение ничем не отличается от телефонного (dialup) модема
подсоединёного по COM. В новых дистритубивах с network-manager и
вовсе достаточно выбрать своего оператора связи и ввести логин с
паролем, программа сделает всё остальное и подключит вас к
интернету через мобильник/модем.

Более подробная [инструкция (ссылка
вникуда)](http://muromec.org.ua/post/zte-my39-linux).

## Как принудительно указать для сетевой карты скорость?

Довольно полная таблица по параметрам карточек, которые можно прописать
в /etc/modules.conf (для 2.4.х) или в /etc/modprobe.conf (для 2.6.х).

[42](http://www.startcom.org/docs/en/Reference%20Guide%20StartCom%20Enterprise%20Linux%203.0.x/s1-modules-ethernet.html)

Для карт realtek 8139 можно указать параметр media=0x0030

К тому же есть инструменты ethertools и mii.

## Как сделать, чтобы старый компьютер видел новый винт большого размера?

Если BIOS от award, то скачать с <http://www.rom.by/> bios patcher,
получить образ старого BIOS, пропатчить и залить обратно.

## Как заставить работать wifi на ноутбуке dell 1521 (модуль Broadcom Corporation BCM94311MCG) в CentOS/RHEL?

1\. Нужно подключить репозиторий rpmforge:

    wget http://dag.wieers.com/rpm/packages/rpmforge-release/rpmforge-release-0.3.6-1.el5.rf.i386.rpm
    rpm -ihv rpmforge-release-0.3.6-1.el5.rf.i386.rpm

2\. Установить модуль ndiswrapper

    yum install dkms-ndiswrapper

3\. Установить firmware Вариант 1 (использовать firmware с openwrt.org и
утилиту b43-fwcutter)

    wget http://ubuntu.cafuego.net/pool/hardy-cafuego/broadcom/b43-fwcutter_008-1.tar.gz
    tar -xzvf ./b43-fwcutter_008-1.tar.gz
    cd b43-fwcutter-008/
    make

Если по ссылке выше не удаётся скачать b43-fwcutter, можно попробовать
поискать тут:
<http://developer.berlios.de/project/showfiles.php?group_id=4547> или
тут: <http://ubuntu.cafuego.net/pool/hardy-cafuego/broadcom/>

    wget http://downloads.openwrt.org/sources/wl_apsta-3.130.20.0.o
    mkdir -p /lib/firmware
    mkdir -p /lib/firmware/`uname -r`
    ./b43-fwcutter -w /lib/firmware ./wl_apsta-3.130.20.0.o
    ./b43-fwcutter -w /lib/firmware/`uname -r` ./wl_apsta-3.130.20.0.o

Вариант 2 (скачать уже готовое firmware)

    wget http://ubuntu.cafuego.net/pool/hardy-cafuego/broadcom/bcm43xx-firmware_1.4-0cafuego1.tar.gz
    tar -xzvf bcm43xx-firmware_1.4-0cafuego1.tar.gz
    mkdir -p /lib/firmware
    cp -p ./bcm43xx-firmware-1.4/firmware/*.fw /lib/firmware/
    mkdir -p /lib/firmware/`uname -r`
    cp -p ./bcm43xx-firmware-1.4/firmware/*.fw /lib/firmware/`uname -r`/

4\. Настроить wifi сеть с помощью утилиты

    system-config-network

PS: Копирую в две директории (/lib/firmware/ и /lib/firmware/\`uname
-r\`/) т.к. не знаю в какую из них правильно.

## Как заставить работать принтер Canon LBP\[XXX\] под Linux и Gentoo в частности?

Итак, что же нужно сделать, чтобы подружить ваш Canon LBP с Gentoo
GNU/Linux? Я (Demon37) объясню как это сделать на примере моего LBP3000.
Для сокращения руководства будем считать, что cups у Вас установлен (он
обычно "вытягивается" как зависимость, поэтому если у Вас настроенная
домашняя GNU/Linux-система, скорее всего cups уже установлен). Если
это не так, обратитесь к справочному руководству по установке и
настройке cups.

Для начала скачайте драйверы с официального
[сайта](http://software.canon-europe.com/software/0031118.asp?model=)
Canon (около 13 мб). Они поставляются в виде tar.gz архива. Распакуйте
скаченный архив:

    tar -xvzf CAPTDRV180.tar.gz

После этого должен появиться новый каталог CANON_UK (название может
быть другим), содержащий файлы документации, исходные тексты
драйвера, rpm- и deb-пакеты. Перейдите в подкаталог, содержайщий
rpm-пакеты и просмотрите список файлов в нём:

    cd CANON_UK/Driver/RPM
    ls

    cndrvcups-capt-1.80-1.i386.rpm  cndrvcups-common-1.80-1.i386.rpm

Затем установите программу rpm2targz и сконвертируйте с её помощью оба
rpm-пакета в tar.gz архивы:

    sudo emerge rpm2targz
    rpm2targz cndrvcups-capt-1.80-1.i386.rpm
    rpm2targz cndrvcups-common-1.80-1.i386.rpm

Распакуйте полученные архивы:

    tar -xvzf cndrvcups-capt-1.80-1.i386.tar.gz
    tar -xvzf cndrvcups-common-1.80-1.i386.tar.gz

После этого выполните следующую операцию (зачем -- разъяснено ниже):

    sudo cp -Rvf usr/lib/* /usr/libexec

Обычно файлы cups в системах GNU/Linux располагаются в /usr/lib/cups, но
в Gentoo они должны располагаться в /usr/libexec/cups. Поэтому мы и
выполнили копирование файлов из каталога lib туда. Теперь удалите
ставший ненужным каталог lib из CANON_UK/Driver/RPM/usr:

    sudo rm -Rvf lib

Далее переместитесь на уровень каталогом выше (в моём случае это
CANON_UK/Driver/RPM) и выполните копирование каталогов etc и usr в
корень системы (/):

    cd ..
    cp -Rvf etc usr /

**Акцентирую Ваше внимание: нужно именно скопировать etc и usr в /
поверх существующих файлов (всё будет в порядке\!). НЕ НУЖНО
предварительно удалять /etc и /usr\!**

Далее, подключаем наш принтер, включаем его и смотрим, видит ли его наша
система:

    sudo lsusb

В моём случае результатом команды будет вывод следующего текста:

    demon37@developer ~ $ sudo lsusb
    Bus 002 Device 001: ID 0000:0000
    Bus 001 Device 003: ID 04a9:266a Canon, Inc.
    Bus 001 Device 002: ID 22b8:4902 Motorola PCS E398 GSM Phone
    Bus 001 Device 001: ID 0000:0000

Если принтера в списке нет, то возможно Вам следует пересобрать ядро
Linux с поддержкой USB Printer (об этом в другой статье). Если же есть,
можно переходить к следующему шагу.

Теперь перезапустите CUPS:

    sudo /etc/init.d/cupsd restart

'' Примечание: в других системах GNU/Linux (не Gentoo), файл-скрипт
обычно носит имя, отличное от "cupsd". Поищите файл, начинающийся с
"cups" в каталоге /etc/init.d/ ''

Затем проверьте наличие каталога /var/ccpd и файла fifo0 в нем. Если
каталог отсутствует, создайте его, выполнив команду:

    sudo mkdir /var/ccpd

Если файл именованного канала fifo0 отсутствует, давайте создадим его:

    sudo mkfifo -m 777 /var/ccpd/fifo0

Если же файл на месте, то измените права доступа к нему таким образом,
чтобы любой пользовател имел возможность писать/читать в/из него. Для
этого выполните:

    sudo chmod 777 /var/ccpd/fifo0

Gentoo ищет файлы, описывающие характеристики принтера (\*.ppd) в
каталоге /usr/share/pppd. Поэтому нам требуется создать
символическую ссылку на нужный файл (нужный файл обычно
содержит в своём имени название модели Вашего принтера. В моём
случае это LBP3000):

    sudo ln -s /usr/share/cups/model/CNCUPSLBP3000CAPTK.ppd
    /usr/share/ppd/CNCUPSLBP3000CAPTK.ppd

Что ж, теперь следует добавить принтер в спул:

    sudo /usr/sbin/lpadmin -p LBP3000 -m CNCUPSLBP3000CAPTK.ppd -v
    ccp:/var/ccpd/fifo0 -E

'' Примечание: Повяление каких-либо ошибок на данном этапе обозначает
то, что Вы, скорее всего, неправильно скопировали файлы драйвера
(вернитесь к началу руководства) ''

После этого нужно зарегистрировать принтер в системе

    sudo /usr/sbin/ccpdadmin -p LBP3000 -o /dev/usb/lp0

Для того, чтобы принтер работал, необходимо создать стартовый скрипт.
Это можно сделать с помощью вашего любимого текстового редактора; к
примеру, **nano**.

    sudo nano /etc/init.d/ccpd

Введите или скопируйте следующий текст скрипта:

    #!/bin/sh
    DAEMON=/usr/sbin/ccpd
    LOCKFILE=/var/lock/subsys/ccpd
    PATH=/usr/local/sbin:/usr/local/bin:/sbin:/bin:/usr/sbin:/usr/bin
    NAME=ccpd
    DESC="Canon Printer Daemon for CUPS"
    test -f $DAEMON || exit 0
    case $1 in
     start)
      echo -n "Starting $DESC: $NAME"
      start-stop-daemon --start --quiet --exec $DAEMON
      echo "."
      ;;
     stop)
      echo -n "Stopping $DESC: $NAME"
      start-stop-daemon --stop --quiet --oknodo --exec $DAEMON
      echo "."
      ;;
    status)
     echo "$DESC: $NAME:" `pidof $NAME`
     ;;
    restart)
     echo -n "Restarting $DESC: $NAME"
     start-stop-daemon --stop --quiet --oknodo --exec $DAEMON
     sleep 1
     start-stop-daemon --start --quiet --exec $DAEMON
     echo "."
     ;;
    *)
     echo "Usage: ccpd {start|stop|status}"
     exit 1
    ;;
    esac
    exit 0

После этого нажмите комбинацию ctrl+o для сохранения и ctrl+x для выхода
из редактора nano.

Теперь выполните только что созданный скрипт

    sudo /etc/init.d/ccpd start

и добавьте его на уровень запуска по умолчанию (default run-level)
командой:

    sudo rc-update add ccpd default

Готово. Теперь, если Вы все сделали правильно, принтер можно
использовать по назначению :)

## Как ноутбук HP научить реагировать на закрытие крышки?

С некоторыми ноутбуками фирмы HP случается такая неприятность. Ноутбук
не реагирует на закрытие крышки, однако если программно отправить его
в hibernate, то после просыпания всё работает нормально.

Для решения проблемы, под пользователем root нужно выполнить команду
типа:

    echo 1 > /proc/acpi/video/*/DOS

Цифра после echo может отличаться в зависимости от модели. Для HP 6710b
это 1. А, для HP 530 цифра 0:

    echo 0 > /proc/acpi/video/C085/DOS

Перед уходом в спящий режим и другими особенными событиями acpi, вы
должны восстанавливать прежнее значение цифры, иначе компьютер
может повиснуть\! Пример как это сделать:
<http://ubuntuforums.org/showthread.php?t=587390&page=4>

Upd. В ядре 2.6.30 на HP530 работает без напильников.

## Как отключить системный динамик?

Добавить в файл /etc/modprobe.d/blacklist.conf строчку:

    blacklist pcspkr

## Как настроить Logitech G27 в wine?

<https://www.linux.org.ru/forum/games/8345541>

## Как работать с микроконтроллерами семейства STM32

ARM-микроконтроллеры STM32 полноценно программируются и прошиваются в
линукс.

Для работы вам понадобятся:

  - компилятор: arm-none-eabi или
    [gcc-arm-none-eabi](https://launchpad.net/gcc-arm-embedded) (должны
    быть в репозитории вашего дистрибутива);
  - JTAG-адаптер
    ([ST-LINK](http://www.st.com/internet/evalboard/product/251168.jsp)
    или китайские копеечные аналоги);
  - утилита для прошивки flash-памяти микроконтроллера:
    [stlink](https://github.com/texane/stlink).

Код писать можно в чем угодно, но при работе с большими проектами удобно
использовать IDE — например, Geany.

Для упрощения работы стоит на первых порах скачать с
[сайта-производителя](http://www.st.com) такие
библиотеки, как STDPeriphLib (работа с периферией), а также
библиотеки для работы с USB, ЖК-экранами, файловой системой VFAT и
т.п.

Компилировать можно вручную, а можно — при помощи make. Кстати,
существуют и решения на cmake для STM32. Простейший Makefile
можно скачать, например, из [моего
проекта](https://sourceforge.net/p/ircontroller/code/ci/3f1305c285c3d7a392364144a7428b3bdbe20e18/tree/Makefile).

При написании своего Makefile'а следует обратить внимание на следующее:

  - для удаления неиспользуемых функций указать сборщику
    `-Wl,--gc-sections -Wl,-s`;
  - после сборки бинарника следует обработать его в "понимаемый"
    микроконтроллером вид при помощи утилиты objcopy;
  - заливка полученного бинарника выполняется при помощи простой команды
    `st-flash write $(BIN) 0x08000000`, где `$(BIN)` — имя бинарника,
    полученного при помощи objcopy.

На первых порах советую приобрести для "опытов" макетку навроде
discovery. У китайцев эти макетки стоят от 10 до 100 долларов в
зависимости от модели. После отладки кода и разводки схемы в
kicad можно приступать к изготовлению своего изделия.

[Category:Hardware](Category:Hardware) [Category:L.O.R.
FAQ](Category:L.O.R._FAQ)