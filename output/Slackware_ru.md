<u>**Русификация Slackware c UTF-8**</u>

  - [Localization](http://docs.slackware.com/slackware:localization) на
    docs.slackware.com

Ниже описаны способы руссификации с учетом особенностей в выпуске
конкретного релиза

# Slackware-14

[Русификация
Slackware 14](http://www.slackware.ru/wiki/%D0%A0%D1%83%D1%81%D0%B8%D1%84%D0%B8%D0%BA%D0%B0%D1%86%D0%B8%D1%8F)
статья на slackware.ru/wiki

# Slackware-13

## Консоль

  - Раскладка по Ctrl+Shift

/etc/rc.d/rc.keymap

    #!/bin/sh
    # Load the keyboard map.  More maps are in /usr/share/kbd/keymaps.
    if [ -x /usr/bin/loadkeys ]; then
     /usr/bin/loadkeys /usr/share/kbd/keymaps/i386/qwerty/ruwin_ct_sh-UTF-8.map.gz
    fi

  - Шрифт, отображающий кириллицу

/etc/rc.d/rc.font

    #!/bin/sh
    setfont Cyr_a8x16.psfu.gz

  - Локаль

/etc/profile.d/lang.sh

    #!/bin/sh

    # en_US is the Slackware default locale:
    #export LANG=en_US

    # There is also support for UTF-8 locales, but be aware that
    # some programs are not yet able to handle UTF-8 and will fail to
    # run properly.  In those cases, you can set LANG=C before
    # starting them.  Still, I'd avoid UTF unless you actually need it.
    #export LANG=en_US.UTF-8
    export LANG=ru_RU.UTF-8

    # One side effect of the newer locales is that the sort order
    # is no longer according to ASCII values, so the sort order will
    # change in many places.  Since this isn't usually expected and
    # can break scripts, we'll stick with traditional ASCII sorting.
    # If you'd prefer the sort algorithm that goes with your $LANG
    # setting, comment this out.
    export LC_COLLATE=C

    # End of /etc/profile.d/lang.sh

Не забываем убедиться, что на выше приведенных файлах (rc.font,
rc.keymap, lang.sh) установлен атрибут выполнения.

Поставить же его можно следующей командой:

    chmod +x полный_путь_к_файлу

Или же, разом все файлы:

    chmod +x /etc/rc.d/rc.keymap /etc/rc.d/rc.font /etc/profile.d/lang.sh

  - Lilo

В /etc/lilo.conf нужно исправить строчку:

    append=" vt.default_utf8=0"

на:

    append=" vt.default_utf8=1"

и выполнить команду:

    lilo

## HAL , udev и X'ы

  - Раскладка через HAL (Slackware 13.0, 13.1)

<!-- end list -->

    cp /usr/share/hal/fdi/policy/10osvendor/10-keymap.fdi /etc/hal/fdi/policy/10-keymap.fdi

Правим строки файла /etc/hal/fdi/policy/10-keymap.fdi с input.xkb, а
именно options, layout, variant, задаем в них примерно следующее:

    <?xml version="1.0" encoding="ISO-8859-1"?>
    <deviceinfo version="0.2">
      <device>
        <match key="info.capabilities" contains="input.keymap">
          <append key="info.callouts.add" type="strlist">hal-setup-keymap</append>
        </match>

        <match key="info.capabilities" contains="input.keys">
          <merge key="input.xkb.rules"   type="string">base</merge>
          <merge key="input.xkb.model"   type="string">evdev</merge>
          <merge key="input.xkb.layout"  type="string">us,ru</merge>
          <merge key="input.xkb.variant" type="string">,winkeys</merge>
          <merge key="input.xkb.options" type="string">terminate:ctrl_alt_bksp,grp:ctrl_shift_toggle,grp_led:scroll</merge>
        </match>
      </device>
    </deviceinfo>

  - Раскладка через udev (Slackware 13.37 и Current)

в Slackware Current используется более новая версия xorg-server в
которой наконец то избавились от поддержки HAL, поэтому
переключение раскладок снова настраивается как и раньше,
только с небольшим отличием

Правим файл /etc/X11/xorg.conf.d/90-keyboard-layout.conf, (если этого
файла нет, то создаём) так, как нам нужно. Конечный результат должен
выглядеть примерно так:

    Section "InputClass"
        Identifier "keyboard-all"
        MatchIsKeyboard "on"
        Driver "evdev"
        Option "XkbLayout" "us,ru(winkeys)"
        Option "XkbOptions" "terminate:ctrl_alt_bksp,grp:alt_shift_toggle,grp_led:scroll"
    EndSection

Перезапускаем иксы, и переключаем раскладку по Alt+Shift, с подсветкой
лампочки "Scroll Lock" на клаве.

  - Типографская раскладка

Если вы часто-часто готовите тексты, в, упаси боже, OpenOffice и он
тупит с автозаменой — то вам пригодится «типографская раскладка»,
такая как [тут](http://sapegin.ru/typolayout)(раскладка Артёма
Сапегина) или вот
[здесь](http://ilyabirman.ru/typography-layout/)(раскладка Ильи
Бирмана), например.

Тогда файл /etc/X11/xorg.conf.d/90-keyboard-layout.conf должен выглядить
следующим образом:

    Section "InputClass"
        Identifier "keyboard-all"
        MatchIsKeyboard "on"
        Driver "evdev"
        Option "XkbLayout" "us+typo,ru(winkeys):2+typo"
        Option "XkbOptions" "lv3:ralt_switch,terminate:ctrl_alt_bksp,grp:alt_shift_toggle,grp_led:scroll"
    EndSection

Перезапускаем иксы, и радуемся, «типографские» символы набираются через
правый Alt, и правый Alt+Shift.

Раскладка немного отличается от двух приведенныйх выше, но
поэкспериметировать и выяснить что-как не составит труда в
том же OpenOffice’е. Да и во всяких жабберах можно выпендриться.

## Русские man-страницы

[1](http://www.slackware.ru/forum/viewtopic.php?f=8&t=234) Если кратко,
то в /usr/lib/man.conf заменить строку

    NROFF /usr/bin/nroff -Tlatin1 -mandoc

на

    NROFF iconv -f utf8 -t koi8r -c | /usr/bin/nroff -Tlatin1 -mandoc -c | iconv -f koi8r -t utf8 -c

# Slackware-12

## Начало

Сначала нужно убедиться что установлены следующие пакеты:

  - l/glibc-i18n (содержит locale)
  - x/x11-fonts-cyrillic (русские шрифты для старых тулкитов, не
    использующих freetype2, например, gtk1)
  - x/dejavu-fonts-ttf
  - x/liberation-fonts-ttf

Если используется kde:

  - kdei/kde-i18n-ru (русский перевод kde)
  - kdei/koffice-i18n-ru (русский перевод koffice)

Если их нет, то доустановить недостающее можно с дистрибутивного диска.
Дополнительно с него же можно поставить русский словарь для aspell, он
находится extra/aspell-word-lists/aspell-ru

## Поддержка UTF-8

Для поддержки UTF-8 понадобятся следующие вещи:

  - ru-utf keymap, можно взять с
    [2](http://mlclm.narod.ru/ru-utf.map.gz) и положить в
    /usr/share/kbd/keymaps/i386/qwerty

## Консоль

Русский шрифт в /etc/rc.d/rc.font

    #!/bin/sh
    #
    # This selects your default screen font from among the ones in
    # /usr/share/kbd/consolefonts.
    #
    unicode_start LatArCyrHeb-16
    for i in 1 2 3 4 5 6;do
            echo -ne "\033%G" >/dev/tty$i
    done

Примечание: если вам понравился шрифт, который стоял по-умолчанию,
вместо <b>LatArCyrHeb-16</b> пропишите <b>cyr-sun16</b>

Раскладка клавиатуры в /etc/rc.d/rc.keymap (переключение по правому Alt,
раскладка winkeys):

    #!/bin/sh
    # Load the keyboard map.  More maps are in /usr/share/kbd/keymaps.
    if [ -x /usr/bin/loadkeys ]; then
      /usr/bin/loadkeys /usr/share/kbd/keymaps/i386/qwerty/ru-utf.map.gz
    fi

Если раскладку ru-utf скачать не удалось, то можно попробовать
нижеприведенный вариант (DOSовская раскладка, переключение,
если верить автору, по CapsLock):

    #!/bin/sh
    # Load the keyboard map.  More maps are in /usr/share/kbd/keymaps.
     if [ -x /usr/bin/loadkeys ]; then
    #  /usr/bin/loadkeys /usr/share/kbd/keymaps/i386/qwerty/ru-utf.map.gz
    #ru4.map переключает раскладки по Caps Lock. можно использовать любой другой .map файл.
      /usr/bin/loadkeys ru4.map
      /usr/bin/dumpkeys -c koi8-r | loadkeys --unicode
     fi

Локаль в /etc/profile.d/lang.sh

    #!/bin/sh
    export LANG=ru_RU.UTF-8

и /etc/profile.d/lang.csh

    #!/bin/csh
    setenv LANG ru_RU.UTF-8

Не забудьте убедиться, что на файлы rc.font, rc.keymap, lang.sh,
lang.csh установлен атрибут выполнения. Поставить его можно командой
**chmod +x**

## Настройка X-ов

Сохранив старый конфиг X-ов cp /etc/X11/xorg.conf
/etc/X11/xorg.conf.old, можно создать новый c помощью xorgsetup. Для
стандартной русско-английской раскладки с переключением по
Ctrl+Shift надо выбрать раскладки us,none,ru,winkeys и Options
grp:ctrl_shift_toggle и grp_led:scroll

В -current вместо winkeys выбирается none, а в /etc/X11/xorg.conf надо
добавить следующий фрагмент:

    Section "ServerFlags"
       Option "AutoAddDevices" "False"
    EndSection

Это отключит автоопределение устройств ввода HALом, но зато настройки
клавиатуры будут читаться из конфига. Подробности в разделе
[X-сервер](X-сервер "wikilink").

Чтобы разрешить использовать 3D-ускорение всем туда же дописываются
строчки:

```

Section "DRI"
   Mode 0666
EndSection
```

## Установка шрифтов от Microsoft

В общем-то после появления в составе дистрибутива шрифтов Liberation
стал необязателен, но шрифты в нём всё равно пока лучше, и
называются привычными пользователю Windows названиями.

  - <http://downloads.sourceforge.net/dropline-gnome/cabextract-1.2-i686-3dl.tgz>
  - <http://downloads.sourceforge.net/dropline-gnome/webfonts-1.3-noarch-3dl.tgz>

Ставить нужно сначала cabextract, потом webfonts.

Другой, более простой по моему (JB) мнению способ это просто взять и
скопировать шрифты из Windows куда нибудь в
/usr/share/fonts/X11/winfonts. Плюс такого метода в том, что будут
доступны шрифты, не входящие в пакет webfonts, например Tahoma.

Если кого-то смущает слово microsoft, можно не смущаться, это шрифты
Agfa Monotype, просто одна малоизвестная софтверная контора купила на
них лицензию и выложила их в сети на наиболее либеральных из все�
купивших условиях.

## Настройка /etc/fstab для поддержки русских имен на компакт дисках и дисках с FAT/NTFS

Имена устройств подставьте свои. Посмотреть из можно в выводе команды
fdisk -l

CDROM/DVD:

```
 /dev/cdrom      /mnt/cdrom   iso9660   user,noauto,nosuid,noexec,nodev,ro,utf8   0   0
```

FAT32:

```
 /dev/sda1      /mnt/windows   vfat   showexec,noexec,nosuid,nodev,umask=000,utf8,codepage=866   0   0
```

NTFS:

    /dev/sda1      /mnt/windows   ntfs-3g   umask=000,locale=ru_RU.UTF-8   1   0

Для поддержки ntfs требуется пакет ntfs-3g (<b>в последних версия�
ntfs-3g опция locale не работает, ntfs-3g сейчас поддерживает только
UTF-8</b>)

Флешка:

```
 /dev/sdb1      /mnt/flash   vfat   showexec,noexec,nosuid,nodev,noauto,user,umask=000,utf8,codepage=866   0   0
```

Смысл каждой опции можно узнать из man mount

## Настройка TeX

По минимуму - запустить texconfig и выбрать PAPER-\>A4. Для красивы�
русских буковок в pdf поставить шрифты cm-super.

Для полноценной поддержки лучше поставить TeXlive

## Настройка man

Для корректного отображения man-страниц надо в /usr/lib/man.conf сделать
замену

    -NROFF /usr/bin/nroff -Tlatin1 -mandoc
    +NROFF /usr/bin/enconv -L ru -x KOI8-R | /usr/bin/nroff -mandoc -Tlatin1 -c | /usr/bin/enconv -L ru -x UTF8