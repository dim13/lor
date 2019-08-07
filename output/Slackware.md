Slackware - один из старейших дистрибутивов GNU/Linux. Он отличается от
других популярных дистрибутивов (таких как Debian, Ubuntu, Mandriva,
RedHat/Fedora и OpenSUSE) тем, что Slackware был и остаётся "самым
UNIX'овым" - его чертами являются стабильность и безопасность, а также
некоторая аскетичность.

# Почему именно Slackware, чем она так выделяется на фоне других диструбутивов?

На этот вопрос Вам ответят ряд статей по следующим ссылкам:
[1](http://www.slackware.ru/wiki/%D0%9F%D0%BE%D1%87%D0%B5%D0%BC%D1%83_%D0%B8%D0%BC%D0%B5%D0%BD%D0%BD%D0%BE_Slackware),
[2](http://linuxforum.ru/viewtopic.php?id=5950)

# История

Будучи студентом в Миннесотском Государственном Университете города
Мурхед, Патрик Фолькердинг получил от своего преподавателя по ИИ
задание установить самый популярный в то время дистрибутив линукс —
Softlanding Linux System на компьютеры в их лаборатории. Всё началось с
записи методов исправления ошибок после установки дистрибутива и правки
конфигов. Позже Патрик начал вносить изменения прямо в установочные
диски с SLS, чтобы сразу получать нужные изменения. Он частично
поменял оригинальные сценарии для установки дистрибутива и добавил
механизм для автоматической установки таких важных компонентов,как
ядро и библиотеки.

Патрик не планировал показывать свою модифицированную версию SLS широкой
публике, рассчитывая, что все его изменения скоро появится в SLS.
Однако, увидев, что этого не происходит, и люди в интернете ждут
нового релиза, он написал пост с заголовком «Want an SLS-like .99pl11A
system?» («Нужна SLS-подобная .99pl11A система?»).

Патрик получил множество ответов на это сообщение, и 17 июля 1993 года
он выложил версию 1.00 своего дистрибутива, получившего название
Slackware, на университетский ftp-сервер.

# Релиз

Последняя версия дистрибутива доступна для загрузки через различные
зеркала , в том числе и через BitTorrent.
[Скачать](http://www.slackware.com/getslack/)\*

*\* Примечание. В свое время когда-то Патрик просил не грузить основной
сервер и загружать из зеркал, так что лучше загружаем через какое-либо
зеркало, н-р:
[yandex](http://mirror.yandex.ru/slackware/slackware-iso/), или тянем
через bittorrent.*

# Дистрибутивы, основанные на Slackware

Есть и такие. Вот небольшой список наиболее популярных форков:

  - [Salix OS](http://www.salixos.org/wiki/index.php/Home/ru/) -
    дистрибутив являющийся полностью обратно совместимым с
    Slackware. Наличествует как 32-х битная так и 64-х битная версия
    дистрибутива, в том числе сборки LiveCD. Основная графическая
    среда — Xfce. Также есть редакции дистрибутива с графическими
    средами KDE, Fluxbox, MATE, Ratpoison и LXDE.

<!-- end list -->

  - [Slax](http://slax.org) или его форк [Porteus](http://porteus.org/)
    - Один из лучших LiveCD/USB и имеет возможность установки на жесткий
    диск. Знаменит с легкостью внесения изменений (на одном лишь
    оффсайте находятся сотня дополнительных пакетов с модулями
    ядра и дополнительным ПО) и скромным размером (помещается на
    mini-cd). Существует в нескольких редакциях - Standart, KillBill
    (wine, dosbox, qemu), Server и PopCorn (помещается на 128
    мегабайтовую флешку).

<!-- end list -->

  - [Slackel](http://www.slackel.gr/) - Греческий дистрибутив,
    базирующийся на current-ветке и репозиториях salix.
    Варианты: KDE и Openbox. Доступен так же в LiveCD/DVD.
    Присутствует выбор русского языка.

<!-- end list -->

  - [DeepStyle](http://deepstyle.org.ua) - Украинский форк Slackware,
    отличается поддержкой русского и украинского языка "из коробки".
    Так же содержит множество пакетов, отсутствующих в оригинальной
    дистрибутиве.

<!-- end list -->

  - [MOPSLinux](http://www.mopslinux.org/) - Российский дистрибутив. В
    плюсы проекта можно записать полностью русифицированный интерфейс
    (включая консоль). Больше [не
    развивается](https://www.linux.org.ru/news/linux-general/4763032),
    последняя версия 7.0 и выходила она в 2010 году.

<!-- end list -->

  - [SlaXBMC](http://slaxbmc.blogspot.ru/) - Дистрибутив базирующийся на
    slackware и включающий в себя медиа-станцию xbmc, так же в
    дистрибутив в качестве WM включен FluxBox. На сайте
    имеются ссылки для загрузки образов, исходников и самих
    пакетов для установки xbmc на slackware.

<!-- end list -->

  - [Vector Linux](http://vectorlinux.com) - Нацелен прежде всего на
    офисные рабочие станции.

<!-- end list -->

  - [Zenwalk](http://zenwalk.org) - Бывший MiniSlack, десктопный
    дистрибутив, ориентированный на маломощные ПК. В качестве
    основной графической оболочки - XFCE.

# Существует ли версия для x86_64 и других платформ?

Официально Slackware поддерживает x86, x86_64 (начиная с релиза 13.0),
S/390 и ARM. Существуют также неофициальные версии дистрибутива для
x86_64 и других платформ:

  - x86_64: [Bluewhite64 Linux](http://www.bluewhite64.com/) ,
    [Slamd64](http://slamd64.com/) ;
  - Macintosh/PowerPC: [Slackintosh](http://workaround.ch/) ;
  - SPARC: [Splack](http://www.splack.org/) ;
  - ARM: [ARMedslack](http://www.armedslack.org) .

# SlackBook

Книга с таким названием поможет Вам познакомиться с основами Slackware
Linux, в том числе поможет в установке системы и даже сборке первого
пакета.

1.  [slackbook.org](http://slackbook.org/) - англ.версия книги,
    оригинал, книга доступна в различных форматах: HTML, Post
    script, PDF

<!-- end list -->

1.  [jack.kiev.ua](http://jack.kiev.ua/docs/slackbook/index.html) -
    русский перевод SlackBook версии 2

<!-- end list -->

1.  [OpenNET](http://www.opennet.ru/docs/RUS/slackware) - та же версия,
    материал опубликованный на OpenNET

# Помогите установить Slackware, пожалуйста\!

1.  Читайте SlackBook\!

<!-- end list -->

1.  [Установка
    Slackware](http://www.slackware.ru/wiki/%D0%A3%D1%81%D1%82%D0%B0%D0%BD%D0%BE%D0%B2%D0%BA%D0%B0_Slackware)
    статья на slackware.ru/wiki

<!-- end list -->

1.  Скорей всего, следующая статья: [Установка
    Slackware](http://www.slackguide.com/content/view/20/11/) на ресурсе
    slackguide.com еще должна решить вашу проблему

## Запуск установщика из образа, лежащего на жестком диске

  - [Booting the Installation Environment from
    HDD](http://docs.slackware.com/howtos:slackware_admin:booting_install_from_hdd)
    (LILO, GRUB2, GRUB, GRUB-legacy)

# Как русифицировать?

Этот вопрос часто задают потому, что это один из немногих популярных
дистрибутивов, в котором он еще не решен "из коробки". Но есть
статья, которая поможет Вам с русификацией Slackware:

  - [Русификация
    Slackware](https://www.linux.org.ru/wiki/en/Slackware/ru) с UTF-8.

Также можно использовать один из русскоязычных клонов:

  - MOPSLinux, увы больше [не
    развивается](https://www.linux.org.ru/news/linux-general/4763032),
    последняя версия 7.0 и выходила она в 2010 году ;
  - [DeepStyle](http://deepstyle.org.ua/)

или клон Slackware, поддерживающий локализацию на этапе установки
(например, [Salix
OS](http://www.salixos.org/wiki/index.php/Home)).

Можно воспользоваться многоязычным установщиком:
[SLINT](http://slint.fr/ru/) - проект интернационализации Slackware.

Информация по локализации на английском языке: [Localization: Adapt
Slackware to your own
Language](http://docs.slackware.com/slackware:localization)

## firefox

1.  [Инструкция по русификации
    Firefox](https://mozilla-russia.org/products/firefox/l10n.html)

<!-- end list -->

1.  Slackware SlackBuild script:
    [1](http://slackware.su/forum/index.php/topic,106.0.html),
    [2](http://taper.alienbase.nl/mirrors/people/alien/slackbuilds/mozilla-firefox-l10n/)

<!-- end list -->

1.  И последний пункт, воспользоваться версией от дистрибутива в котором
    предусмотрена локализация

<!-- end list -->

  -
    \* Salix, только учтите, они собирают esr-версию, которую
    понадобится установить и потом на нее прилагающийся пакет
    вида: mozilla-firefox-l10n-LOCALE.
    (Информация для пользователей slapt-get: [Get localized Firefox and
    Open
    Office](http://docs.salixos.org/wiki/Get_localized_Firefox_and_Open_Office)).

<!-- end list -->

  -
    \* В репозитории
    [DeepStyle](ftp://download.deepstyle.org.ua/pub/slackware/slackboost-14.1/slackboost/sbxap/)
    так же могут быть найдены пакеты для локализации продуктов mozilla:
    firefox-l10n и thunderbird-l10n.

## mp3-теги

  - [taglib-rusxmms-1.9.1 для
    slackware-14.1_current](http://www.linuxhub.ru/viewtopic.php?p=5858#p5858)
    - патчим taglib для русских тегов mp3 отличных от utf-8 в Amarok,
    Clementine \&etc.

## архивы: ru_Кодировка

Некоторые архиваторы при обращении к zip, работают с архивами не через
infozip, а через p7zip, н-р: архивы можно просматривать через mc, он
работает через infozip, аналогично и xarchiver, а вот engrampa
(mate) по-умолчанию через p7zip.

  - infozip , способы:

<!-- end list -->

1.  [пересборка infozip](http://slackware.su/forum/index.php?topic=33.0)
    с патчами (libnatspec) для zip-архивов;

<!-- end list -->

1.  подключаем репозиторий
    [deepstyle](http://slackboost.blogspot.ru/2014/10/slackware.html) ,
    как это выглядит для slackpkg+:

<!-- end list -->

  -
    MIRRORPLUS\['deepstyle'\]=<ftp://download.deepstyle.org.ua/pub/slackware/slackboost-14.1/slackboost/>

<!-- end list -->

  -
    устанавливаем приоритет: PKGS_PRIORITY=( deepstyle:infozip
    deepstyle:libnatspec )

<!-- end list -->

  -
    и устанавливаем: infozip и libnatspec оттуда

<!-- end list -->

  - [p7zip](https://www.linux.org.ru/forum/desktop/11489814?lastmod=1428716743137#comment-11493559)
    - И в случае необходимости здесь бы я Вам посоветовал бы собрать
    p7zip с libnatspec, который кстати даст вам возможность работы с
    7zip-архивами.

## man и русский язык

  - [Патчим
    man](https://www.linux.org.ru/forum/desktop/11489814?lastmod=1428770216417#comment-11500943)
    и устанавливаем
    [man-pages-ru](http://slackbuilds.org/result/?search=man-pages-ru&sv=).
    Для последнего не забудьте прописать в
    /usr/lib${LIBDIRSUFFIX}/man.conf для 14.1 и /etc/man.conf -
    14.1_current

`#NROFF     /usr/bin/nroff -mandoc`
`NROFF      /usr/bin/groff -Dutf8 -Tutf8 -mandoc`

  - Собранный пропатченный man от
    [DeepStyle](ftp://download.deepstyle.org.ua/pub/slackware/slackboost-14.1/slackboost/ap/)
    - уже готовый пропатченый пакет man с поддержкой utf8 может быть
    найден у них в репозитории

<!-- end list -->

  - Еще шпаргалка, если вас угораздило использовать man-pages-ru на
    koi8r от ALT или те которые собраны были в пакет во времена
    MOPSLinuх, то в man.conf необходимо прописать следующее:

`NROFF      /usr/bin/nroff -Tlatin1 -mandoc -c | iconv -f koi8r`

  -
    (порядок и действия для сборки man-pages-ru-alt подсмотреть можно в
    [Arch/AUR](https://aur.archlinux.org/packages/man-pages-ru-alt/))

<!-- end list -->

  - И если все так на каких-то руководствах (man xfce4-terminal)
    встречается непростая кодировка и не хочется заморачиваться
    с унификацией хранения, есть смысл глянуть в сторону сторонней enca
    (есть на SBo), которая умеет угадывать кодировку, через нее гнать
    заведомо в utf8, а дальше форматировать через groff -Dutf8 -Tutf8
    -mandoc , в man.conf пропишите:

`NROFF /usr/bin/enconv -L ru -x utf8 | /usr/bin/groff -Dutf8 -Tutf8 -mandoc`

  -
    У DeepStyle, аналогичное реализовано по подобной схеме
    [slackboost/sba/set_locale](ftp://download.deepstyle.org.ua/pub/slackware/slackboost-current/slackboost/sba/set_locale)

# Как создать пользователя?

Если возникают затруднения воспользуйтесь adduser , введя в терминале
данную команду. Вам будет задан ряд вопросов и будет создан
пользователь, добавленный в соответствующие группы.

## Хочу другое приглашение для пользователя, а не что-то типа bash-4.3$

  - [bash\#customization](http://docs.slackware.com/slackbook:bash#customization)

## Ставил KDE, а графического входа в систему не вижу

Что бы он был необходимо назначить runlevel 4 в /etc/inittab

    # These are the default runlevels in Slackware:
    #   0 = halt
    #   1 = single user mode
    #   2 = unused (but configured the same as runlevel 3)
    #   3 = multiuser mode (default Slackware runlevel)
    #   4 = X11 with KDM/GDM/XDM (session managers)
    #   5 = unused (but configured the same as runlevel 3)
    #   6 = reboot

    # Default runlevel. (Do not set to 0 or 6)
    id:4:initdefault:

## Не все команды выполняются от пользователя, пишет: команда не найдена

Скорей всего вы здесь имеете в виду:

    $ lspci
    bash: lspci: команда не найдена

в то время как от root все выполняется. А не выполняется оно по одно
простой причине, т.к. исполняемый файл находится в /sbin/lspci

    # which lspci
    /sbin/lspci

в принципе, это правильно, но если Вам все таки хочется что бы было
автозаполнение и какой-то вывод, то переменная
[PATH](http://docs.slackware.com/howtos:cli_manual:shells#path_variable),
в помощь.

Если же вкратце, то можно поступить так, создать файл .bashrc в домашнем
каталоге пользователь (\~/.bashrc) с содержимым:

    . /etc/profile
    PATH=$PATH:/sbin/

пропишите строчки выше и перезапустить терминал (перезайдите)

# Пакеты

  - [Slackware/packages](https://www.linux.org.ru/wiki/en/Slackware/packages)
    - эта страница ответит на ряд вопросов относительно пакетов, в том
    числе: А есть ли в Slackware менеджер пакетов?

# Как можно обновить Slackware, не переустанавливая ее?

Правильный способ обновления до текущей стабильной версии можно прочесть
в UPGRADE.txt. Так, н-р, для версии 14.1 было
[UPGRADE.TXT](http://ftp.gwdg.de/pub/linux/slackware/slackware-14.1/UPGRADE.TXT),
отдельные рекомендации также написаны в CHANGES_AND_HINTS.TXT, н-р:
[CHANGES_AND_HINTS.TXT](http://slackware.at/data/slackware64-14.1/CHANGES_AND_HINTS.TXT).

-----

Перед обновлением до Current надо **обязательно** прочитать
[Changelog](http://slackware.com/changelog/current.php?cpu=i386), обычно
Патрик указывает возможные проблемы при обновлении и пути их решения. И
еще глянуть желательно бы
[UPGRADE.TXT](http://mirror.yandex.ru/slackware/slackware-current/UPGRADE.TXT)
.

  - [Обновление Slackware до
    current](http://www.slackware.ru/wiki/%D0%9E%D0%B1%D0%BD%D0%BE%D0%B2%D0%BB%D0%B5%D0%BD%D0%B8%D0%B5_Slackware_%D0%B4%D0%BE_current)
    статья на slackware.ru/wiki

Также можно обновлять дистрибутив средствами slackpkg (update,
install-new, upgrade-all, clean-system), но этот метод менее надёжен.

  - [Обновление до current с использование
    slackpkg](http://slackalaxy.wordpress.com/2012/08/03/upgrading-to-slackware-current/)
    - Внимание\! 2012/08/03 - возможно что-то изменилось.

# А имеется ли образ slackware-current.iso и где его можно скачать?

Патрик не собирает в образ current, но никто не запрещает Вам самим
собрать образ на основе текущего среза репозитория или скачать уже
готовые собранные образы (с установщиком) для установки
slackware-current:

  - [сurrent-ISO 32](http://taper.alienbase.nl/mirrors/slackware/slackware-current-iso/)
    и
    [64](http://taper.alienbase.nl/mirrors/slackware/slackware64-current-iso/)
    от AlienBOB
  - [current-ISO 32
    и 64](http://ftp.slackware.no/slackware/slackware-iso/slackware-current-iso/)
    от slackware.no (сайт поддерживает так же ftp и rsync)

Скрипты для сборки ISO, которые так же будут полезны при необходимости в
написании своего скрипта:

  - [dvd.mkisofs](http://taper.alienbase.nl/mirrors/slackware/slackware64-current-iso/readme_dvd.mkisofs)
    и
    [mirror-slackware-current.sh](http://slackware.com/~alien/tools/mirror-slackware-current.sh)
    от AlienBOB
  - [makeSlackISOs.sh](http://slackware.no/makeSlackISOs.sh) от
    slackware.no

# Срез репозитория

Если есть необходимость локально у себя держать реп, то для получения
его у себя на машине можно воспользоваться rsync. О том как зеркалить
написано здесь [Mirroring
Guidelines](http://mirrors.slackware.com/guidelines/) .

Да, и если только осваиваетесь и у Вас затруднения какие ключи для рсинк
прописывать - не ленитесь, смотрите на самом сайте
дистрибутива/проекта/зеркала, чаще всего про это
рассказано. Вот, н-р, у Salix: [How to create a public Salix
mirror](http://docs.salixos.org/wiki/How_to_create_a_public_Salix_mirror)

Зеркала: [slackware.org.uk](http://slackware.org.uk/) ,
[mirror.yandex.ru](http://mirror.yandex.ru/slackware/) и т.д.

# Как добавить поддержку multilib в slackware64?

  - [Добавление мультибиблиотечности в Slackware
    x86_64](http://docs.slackware.com/ru:slackware:multilib)

# Железо и оборудование

## Звук: ALSA

По-умолчанию в Slackware для ввода/вывода звука используется ALSA. Далее
про основные команды для настройки:

  - **alsamixer** - настройка уровня громкости и прочее

<!-- end list -->

  - **alsactl init** - автоопрделение карточки

<!-- end list -->

  - **alsactl store** - cохранение настроек, что бы они после
    перезагрузки системы были сохранены

### Skype

В последних версиях Skype работает через pulse (pulseaudio).

По поводу pulse в skype AlienBOB отписался в своем блоге: [Skype drops
support for
ALSA](http://alien.slackbook.org/blog/skype-drops-support-for-alsa/) ,
[перевод](http://slackware.su/forum/index.php?topic=1471.0).

Найти же pulse, как и skype можно в готовом виде в различных
репозиториях или собрать через SBo. Но есть и другое
решение, если неохота возиться с установкой и настройкой этого
сервиса: поставить apulse (из репозиториев или через sbopkg) и
запустить skype:

`apulse skype`

## Видео

### NVidia-Optimus

  - [SlackBuilds (Slackware build
    scripts)](https://github.com/WhiteWolf1776/Bumblebee-SlackBuilds)
    for Bumblebee and related dependencies
  - <http://docs.slackware.com/howtos:hardware:nvidia_optimus>

## Принтеры

Настройка и определение принтеров происходит через CUPS. Запустим
сервис:

`chmod +x /etc/rc.d/rc.cups`
`/etc/rc.d/rc.cups restart`

и в браузере перейдите к настройкам CUPS: <http://localhost:631/>

Если не нравится способ через веб-браузер, то есть утилитла для
подключения и настройки принтеров: **system-config-printer**

### Samsung

Если есть проблемы с подключение принтеров от этой фирмы не торопитесь
обновлять самостоятельно CUPS, попробуйте воспользоваться следующим
решением:

  - [SpliX](http://splix.sourceforge.net) (CUPS printer drivers for SPL)

Данный пакет найти можно в репозиториях salix (в 14.1 есть точно) или
самостоятельно собрать из
[SBo](http://slackbuilds.org/result/?search=splix&sv=)

# VirtualBox

Если необходимо, чтобы программа была в системе, то самый правильный
способ - это собрать пакет с VBox и установить его. Еще есть готовый
вариант .run-файл от Sun/Orcale, но обо всем этом ниже.

## SBo/SlackBuilds Repository

Чтобы у вас в системе был Vbox, собираем вот эти пакеты: virtualbox и
virtualbox-kernel, отсюда
<http://slackbuilds.org/result/?search=virtualbox&sv=14.1> перед сборкой
может потребоваться acpica или его сборка с установкой, если нет его в
системе.

`sbopkg -i "acpica virtualbox-kernel virtualbox virtualbox-extension-pack"`

Создаем группу и добавляем в него пользователя:

`groupadd -g 215 vboxusers`
`usermod -a -G vboxusers ИмяПользователя`

(Пере)запустим сервис:

`/etc/rc.d/rc.vboxdrv restart`

С запуском сервиса при старте системы возможны проблемы: [ссылка на
форум
slackware.su](http://slackware.su/forum/index.php/topic,146.msg703.html?PHPSESSID=mbclcp436horg7f8uo7u12f936#msg703).
Более подробнее:
[README.SLACKWARE](http://slackbuilds.org/slackbuilds/14.1/system/virtualbox/README.SLACKWARE)
для virtualbox на SBo.

Так что, чтобы происходил автозапуск, если не запускается, прописываем в
rc.local и rc.local_shutdown - /etc/rc.d/rc.vboxdrv start и
/etc/rc.d/rc.vboxdrv stop , соответственно. Перезапускаемся.

## VirtualBox binaries

Другой вариант установки, в принципе все прекрасно ставится и работает ,
но возможны проблемы с совместимостью qt-текущей версией в Вашей
системе: [Qt themes are not
applied](https://www.virtualbox.org/ticket/6830), да и в обще как-то
неправильно запускать что-то вида "съешь-выпей-меня.run"

1.  Скачиваем: [VirtualBox for Linux
    hosts](https://www.virtualbox.org/wiki/Downloads) - All
    distributions i386 | AMD64

<!-- end list -->

1.  Устанавливаем:

`# sh VirtualBox-ВЕРСИЯ.run`

1.  Для удаления:

`# sh VirtualBox-ВЕРСИЯ.run --noexec --keep`
`# cd install`
`# sh uninstall.sh`

Здесь же, в этом вариант добавление в группу так же необходимо и для
обнаружения usb-устройств установка через VirtualBox - Extension
Pack на странице загрузки файл для загрузки доступен.

## Другое

Если в виртуалке у вас Slackware и необходимы различные дополнения для,
н-р, изменения экрана виртуалки и монтирования устройств: [Virtualbox
need
kernel-sources](http://www.slackel.gr/slackelmulti/xoops20171/htdocs/modules/newbb/viewtopic.php?post_id=1011&topic_id=304&forum=4&lang=english)

# LILO

## Спящий режим (Гибернация) и LILO

Для спящего режима, как и в других дистрибутивах, нам необходим
swap-раздел (больший или равный нашей ОЗУ).

Чтобы вернуть систему, у которой загрузчик LILO, из спящего режима,
необходимо прописать в файле /etc/lilo.conf в строчке append=
что-то вроде:

    append=" resume=/dev/sdaX"

где sdaX - раздел подкачки, X - номер его в таблице разделов. Не
забываем после всех действий дать команду lilo.

Ниже ссылки на различные источники, в которых представлены способы
заставить работать спящий режим (скажем для тех случаев, если у
Вас установлено не DE):

  - <http://docs.slackware.com/howtos:slackware_admin:hibernation>
  - <http://www.slackwiki.com/Hibernate>

## GRUB в MBR, а LILO был установлен на раздел со Slackware

Предположим система была установлена на раздел sda3 и туда же установили
загрузчик, но в MBR у Вас установлен GRUB от другой системы, которую вы
не удалили по различным причинам. Подмонтируете раздел с этой системой
или просто зайдите в нее и в /boot/grub/grub.cfg пропишите рядом, ниже
со строчкой загрузки ее:

    menuentry 'Slackware' --class slackware --class gnu-linux --class gnu --class os {
            set root='(hd0,msdos3)'
        chainloader +1
    }

Опять таки, это одно из решение, еще есть os-prober и т.д.

## Изменяем картинку в загрузчике

  - [Replace Lilo Splash with Mayan
    Slackware 14.0](http://slackblogs.blogspot.ru/2012/10/replace-lilo-splash-with-mayan.html)

# Шрифты

Ниже приведены ссылки на темы с различных источников, в которых вы
найдете рабочие версии билдов и пакеты которые можно установить.

  - [How to Optimize Fonts in
    Slackware](http://www.linuxquestions.org/questions/slackware-14/how-to-optimize-fonts-in-slackware-640468/)
    - тема на LQ (linuxquestions.org)

Infinality:

  - Infinality в
    [Slackware-14.1](http://linuxforum.ru/viewtopic.php?pid=410345)
  - The packages and SlackBuilds are for
    [Slackware 13.37/14.0](https://someslack.wordpress.com/infinality-on-slackware/)

# Игры

  - SBo: проходим сюда <http://slackbuilds.org/repository/>, выбираем
    версию системы и далее репозиторий (SlackBuild'ов) - Games или же
    вверху в поиске вбиваем необходимое нам название искомой игры,
    библиотеки, эмулятора и т.д.

<!-- end list -->

  - [Gens/GS](http://segaretro.org/Gens/GS), эмулятор SEGA/Genesis,
    доступен собранный пакет Gens/GS Release 7 (Slackware 13.0,
    i486), в пакете прилагается SlackBuild.

## Steam

  - [Valve’s Steam
    client](http://alien.slackbook.org/blog/valves-steam-client-for-linux/)
    - Заметка в блоге AlienBOB'а.
  - [steamclient](http://www.slackware.com/~alien/slackbuilds/steamclient/)
    - Пакеты и SlackBuild'ы.
  - [steamcommunity - Сообщество
    Steam](http://steamcommunity.com/groups/slackware) - Группа
    Slackware в Steam.

# Хочу быть в курсе последних новостей, что и где читать, смотреть?

В первую очередь это старая, добрая рассылка: [Slackware Linux Project
Mailing Lists](http://www.slackware.com/lists/) .

Также можно ещё время от времени поглядывать историю внесенных изменений
в дерево дистрибутива: [Slackware
ChangeLogs](http://www.slackware.com/changelog/) или через файл
ChangeLog.txt в репозитории пакетов. В официальном репозитории нет
ChangeLog.rss, но есть зеркала которые делают свой rss, н-р: [зеркало
slackware на SBo](http://slackbuilds.org/mirror/rss/) , это если вам
нужна подписка через RSS.

Временами и AlienBOB информируем о чем-то новом в своем блоге и ещё по
своим решениям в Slackware: [Alien Pastures, blog
articles](http://alien.slackbook.org/blog/) .

Ну и ещё есть множество личных блогов пользователей, вот, например, один
из них: [Unofficial Slackware Blog](http://slackblogs.blogspot.ru/),
который точно на данный момент (2014 г.) работает, обновляется.

# Различные советы

## Настройка часов BIOS

  - [How To Sync Your System Time to Hardware Clock Consistently How To
    Sync Your System Time to Hardware Clock
    Consistently](http://docs.slackware.com/howtos:hardware:syncing_hardware_clock_and_system_local_time)

## Что делать, если в -current не запускаются X с сегфолтом?

FIXME: временный workaround пока не починят более глобально

Включить KMS, для этого добавить "radeon.modeset=1" к строчке с
параметрами ядра append="vt.default_utf8=1" (чтобы стало
append=" vt.default_utf8=1 radeon.modeset=1"). Здесь radeon - название
видеодрайвера, вместо него может быть i915 (для intel) или nouveau (для
nvidia). С проприетарными драйверами этой проблемы нет.

При этом отвалится framebuffer и консоль будет радовать vga-разрешением.
Если этот вариант принципиально не нравится, что неудивительно, то надо
убрать или закомментировать из /etc/lilo.conf все строчки про
vga=циферки и дописать в append="vt.default_utf8=1
radeon.modeset=1 ещё до кучи video=1024x768@85" где 1024x768 ---
разрешение, а 85 --- частота обновления в Hz. После этого нужно,
чтобы драйвер radeon, i915 или nouveau грузился раньше, чем будет
настраиваться третий или четвёртый ранлевел. Для этого его нужно
добавить или в тот файл, на который ссылается /etc/rc.d/rc.modules
(rc.modules-версия ядра, в конец дописывается строчка /sbin/modprobe
radeon или нужный драйвер), или включить его в initial ramdisk.

## У меня <что-то> работает от рута, а от пользователя нет

Нужно добавить пользователя в соответствующую группу, например, audio
для работы со звуковыми устройствами, scanner для работы со сканером
и т.п. Список всех групп можно получить с помощью команды

    cat /etc/group

добавить пользователя командой

    useradd -aG список_групп_через_запятую имя_пользователя

## В свежеустановленной Slackware непривычно работает (g)vim. Что делать?

Скопировать файл (g)vimrc_example.vim из /usr/share/vim/vim${VERSION} в
\~/.(g)vimrc, где ${VERSION\] - версия vim.

А также проверить, на что указывает символическая ссылка /usr/bin/vi. В
установке по умолчанию она ссылается на /usr/bin/elvis.

## Как настроить воспроизведение MIDI файлов?

Заставить MIDI работать в Slackware не так уж и просто, как может
показаться. Предлагается простое решение этой проблемы.

Испытывалось на полностью установленном дистрибутиве Slackware 12,
дополнительно ни одного постороннего пакета не устанавливалось,
кроме пакетов поставляемых с дистрибутивом.

**Загрузка SlackBuild'а:**

Скачайте архив со
[SlackBuild](http://www.vcn.bc.ca/%7Edugan/timidity/TiMidity++_SlackBuild.tar.bz2)'ом
и разархивируйте его в рабочий каталог, где будет создаваться пакет. В
результате получите каталог “TiMidity++_SlackBuild”, с необходимыми
для создания пакета файлами.

**Загрузка исходников:**

Скачайте в каталог “TiMidity++_SlackBuild” два пакета:

  - [TiMidity++-2.13.2.tar.bz2](http://easynews.dl.sourceforge.net/sourceforge/timidity/TiMidity++-2.13.2.tar.bz2)
  - [eawpats12_full.tar.gz](http://gentoo.mirrors.pair.com/distfiles/eawpats12_full.tar.gz)

**Запуск SlackBuild'а:**

Под Root'ом перейдите в каталог со SlackBuild'ом и сделайте его
запускаемым:

    chmod +x TiMidity++.SlackBuild

и запустите его:

    ./TiMidity++.SlackBuild

**Установка пакета:**

После окончания работы SlackBuild'а перейдите в каталог /tmp и
установите созданный пакет с помощью команды installpkg.

**Добавление запуска TiMidity++ в скрипты инициализации:**

В каталоге /etc/rc.d создайте файл rc.local, если его там нет, и
сделайте запускаемым. Затем добавьте в него строку запуска
миди-сервера:

    /etc/rc.d/rc.timidity start

**Наслаждайтесь своим MIDI**

Теперь у вас есть рабчий порт MIDI - port 128:0. Можете проигрывать MIDI
файлы непосредственно. С помощью утилиты
[mozplugger](http://mozplugger.mozdev.org/) можно обеспечить потоковое
воспроизведение MIDI файлов в SeaMonkey и Firefox.

**"Внедрение" TiMidity++ в Firefox:**

После установки утилиты [mozplugger](http://mozplugger.mozdev.org/), Вы
легко можете добавить TiMidity++ в SeaMonkey и Firefox. После
добавления, проверьте файл /etc/mozpluggerrc на предмет
наличия фрагмента текста :

    audio/mid: midi,mid: MIDI audio file
    audio/x-mid: midi,mid: MIDI audio file
    audio/midi: midi,mid: MIDI audio file
    audio/x-midi: midi,mid: MIDI audio file
        embed controls: timidity -EFreverb=0 -EFchorus=0 "$file"
        noembed swallow(timidity) hidden: timidity -EFreverb=0 -EFchorus=0 -ig "$file"

Запустить GUI проигрывателя timidity можно из командной строки эмулятора
терминала:

    timidity -iA -B2,8 -Os -EFreverb=0 -ig

[см. рисунок](http://jpe.ru/1/max/141209/0ouokv8cea.jpg)

## Настройка службы akonadi

В Slackware 13 под KDE 4 появился пакет akonadi. Для того чтобы
правильно запускалась служба akonadi, нужно запустить сервер
MySQL. Для этого нужно сделать исполняемым файл /etc/rc.d/rc.mysqld и
запустить службу командой:

    /etc/rc.d/rc.mysqld start

Если MySQL запротестовал, значит у него нет разрешения на использование
своих же каталогов и файлов. Исправить эту оплошность можно следующей
командой:

    chown -R mysql:mysql /var/lib/mysql

Затем под рутом, изменить файл /usr/bin/startkde, в котором нужно
добавить путь /usr/share/akonadi в строку, которая определяет
пути XDG_DATA_DIRS.

    # Make sure that D-Bus is running
    if test -z "$XDG_DATA_DIRS"; then
      XDG_DATA_DIRS="`kde4-config --prefix`/share:/usr/share:/usr/local/share:/usr/share/akonadi"
      export XDG_DATA_DIRS
    fi

После этого в файл /home/юзер/.local/share/akonadi/mysql.conf нужно
добавить в строку "user=юзер" под заголовком "\[mysqld\]".

Для рута нужно добавить в файл /root/.local/share/akonadi/mysql.conf
строку "user=root" под заголовком "\[mysqld\]".

    [mysqld]
    skip_grant_tables
    skip_networking
    user=root

После двух перезагрузок служба "самонастроится" и сервер akonadi
перестанет при пуске KDE выбрасывать сообщения об ошибках.
Способ настройки конечно грубый, но простой. О правильной и более
тонкой настройке нужно заглянуть в /usr/doc/mysql, там можно найти
подробную информацию как настроить MySQL сервер. Но это для тех,
кто хорошо знает предмет и владеет английским на достаточном уровне.

## Совет по настройке загрузки на машине с 8Мб ОЗУ

<http://www.flaterco.com/kb/8MiB.html>

В принципе всё просто: пересобираем ядро под самый минимум (зависит от
реально установленного на машине железа), и обходим проблему с udev,
который пытается разместить /dev в tmpfs и не может этого сделать, по
причине малой памяти. Решение: /dev создаётся на диске, как в старые
времена. Грузится минуты 3, но потом работает.

# Ссылки

1.  [slackware.com](http://slackware.com) - основной ресурс

<!-- end list -->

1.  [docs.slackware.com](http://docs.slackware.com) - мультиязычный
    проект документации Slackware

<!-- end list -->

1.  [alien.slackbook.org](http://alien.slackbook.org/dokuwiki/doku.php)
    - Alien BOB's Wiki

<!-- end list -->

1.  [slackwiki.org](http://slackwiki.org) - SlackWiki

<!-- end list -->

1.  [собрание различных ссылок](http://www.slackwiki.com/Links) от
    SlackWiki

<!-- end list -->

1.  [linuxquestions.org/questions/slackware-14](http://www.linuxquestions.org/questions/slackware-14/)
    - LQ-Slackware - официальный форум (en) на Linux Questions

<!-- end list -->

1.  [slackware.ru](http://slackware.ru) - информационный ресурс,
    посвящённый дистрибутиву Slackware Linux и информационным
    технологиям в целом.

<!-- end list -->

1.  [slackware.su](http://slackware.su) - статьи и форум, посвящённые
    Slackware Linux.

<!-- end list -->

1.  [slackware.pp.ru](http://slackware.pp.ru) - xороший сайт с подробным
    описанием установки, русификации и настройки. Может быть временно в
    дауне.

<!-- end list -->

1.  [slackworld.berlios.de/links](http://slackworld.berlios.de/links.html)
    - ссылки на дополнительные материалы по Slackware.

<!-- end list -->

1.  [slackblogs.blogspot.com](http://slackblogs.blogspot.com/) -
    комментарии про последние обновления в -current и
    дополнительных репозиториях.

<!-- end list -->

1.  [MLED](http://www.microlinux.fr/mled.php) - Microlinux Enterprise
    Desktop и описание на русском про проект MLED на
    [zenway.ru](http://zenway.ru/page/meld).

<!-- end list -->

1.  [переводы README](http://wiki.slackware.su/wiki:od) различных файлов
    из дерева дстрибутива на wiki.slackware.su

<!-- end list -->

1.  [alv.me](http://alv.me/?p=5911) или [IBM
    dW](http://www.ibm.com/search/csass/search/?q=%D0%9F%D0%BE%D0%B3%D1%80%D1%83%D0%B6%D0%B5%D0%BD%D0%B8%D0%B5+%D0%B2+Salix&dws=rudw&ibm-search.x=0&ibm-search.y=0&sn=dw&lang=ru&cc=RU&ddr=&en=utf&lo=ru&hpp=20)
    - Погружение в Salix. Данный цикл посвящён дистрибутиву Salix – его
    установке, настройке и дистрибутив-специфическим особенностям
    (управление пакетами, утилиты конфигурирования), а также его
    назначению и месту в общей картине мира Linux.

<!-- end list -->

1.  [koptev.ru](http://koptev.ru/docs/oracle10) - установка Oracle на
    Slackware

<!-- end list -->

1.  [Slint](http://slint.fr/ru/slint.html) - Проект интернационализации
    Slackware

<!-- end list -->

1.  [Recommended UID/GIDs](http://slackbuilds.org/uid_gid.txt) for use
    with SlackBuilds.org scripts

<!-- end list -->

1.  Slackware [Banners, Logos,
    Propaganda](http://www.slackware.com/~msimons/slackware/grfx/)

<!-- end list -->

1.  [Slackware Linux
    Graphics](http://slackware-linux-graphics.blogspot.ru/)

<!-- end list -->

1.  [Создание минимального CD
    Slackware](http://unixforum.org/index.php?s=c6fb2114f3593e7d88dc23ebe3c38730&showtopic=135998&view=findpost&p=1252280)
    и сами образы на основе релизов от AlienBOB:
    [slackware-mini-install.iso](http://www.slackware.com/~alien/slackboot/mini/)

[Category:Дистрибутивы](Category:Дистрибутивы "wikilink")