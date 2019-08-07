Это один из первых дистрибутивов Linux. Прославился огромным количеством
пакетов, наибольшим списком поддерживаемых архитектур, повышенными
требованиями к лицензиям, качественным тестированием и фразой
"Debian выпускается, когда приходит время", что означает, что релиз-цикл
у Debian ОЧЕНЬ долгий.

[Официальный сайт](http://debian.org)

## Что ставить - stable/testing/unstable?

Новичкам лучше начинать со stable, стабильный релиз имеет проверенный
набор пакетов с небольшой вероятностью того, что придется встретиться
с какими-либо ошибками, регулярно и в кратчайшие сроки получает
обновления безопасности. С другой стороны в stable будет старая
версия ядра, которая, возможно, не сможет работать с современным
железом. Частично это решается использованием
[бекпортов](http://www.linux.org.ru/wiki/en/Debian#Откуда_брать_пакеты%2C_отсутствующие_в_дистрибутиве%3F).

После детального понимания дистрибутива десктопы можно обновить до
testing. Тестируемая ветка отличается более новыми версиями пакетов
программ, а поэтому есть вероятность того, что иногда встречаются
ошибки или что-то бывает сломано, также возможны конфликты при
обновлении пакетов из за меняющихся зависимостей и изменения
версий программ. В случае использования testing можно (на свой страх
и риск) смешивать пакеты из репозиториев testing, unstable или
experimental. Для ветки stable подобным заниматься не следует.

А unstable - для экстремалов или тех, кто знает, что делает.

## Squeeze, Wheezy, Jessie, Sid - что это такое?

Это кодовые имена релизов Debian, названия берутся из мультфильма "Toy
Story" компании Pixar. В каждый момент времени существует три
официальные версии Debian:

  - Stable - стабильная версия, в которой версии ПО не обновляют, а
    только исправляют найденные ошибки.
  - Testing - версия, в которой пакеты переносятся из unstable, после
    предварительной заморозки и тестирования.
  - Unstable - ветка, содержащая последние версии ПО и, соответственно,
    n-ое количество ошибок :).

Так же существует Oldstable (предыдущий Stable) и Experimental, который
не является полноценной веткой - в нем находятся пакеты, требующие
тщательного тестирования или которые повлекут серьезные изменения в
дистрибутиве (например новая версия gcc или Xorg). Через определенные
промежутки времени stable отправляется в утиль, а текущий testing
становится новым stable.

**Sid** - кодовое название unstable. Этот релиз никогда не будет
выпущен. (В "Toy Story" Sid Phillips - хулиганистый мальчишка,
ломающий игрушки. Достаточно символичное название для нестабильной
ветки)

**Jessie** - текущий testing. Спустя некоторое время станет новым
stable, Debian 8.0

**Wheezy** - Debian 7.0, текущий stable. вышел
[04.05.2013](https://www.linux.org.ru/news/debian/9129003)

**Squeeze** - Debian 6.0, предыдущий stable (или как его еще обычно
называют - oldstable)

Кодовые названия старых релизов: buzz: 1.1, rex: 1.2, bo: 1.3.x, hamm:
2.0, slink: 2.1, potato: 2.2, woody: 3.0, sarge: 3.1, etch: 4.0, lenny:
5.0

Само название **Debian** происходит от имени создателя Ian Murdock и его
жены Debra.\[1\]

## Вопросы по установке Debian (свободное ядро и несвободные драйвера)

В связи с [переносом несвободных драйверов устройств из
ядра](http://www.debian.org/News/2010/20101215) в пакет
**firmware-linux-nonfree** в разделе **non-free** у некоторых
пользователей могут возникнуть проблемы с оборудованием во
время установки Debian, например, из-за нехватки драйвера сетевой
карты (например от Broadcom) не будет связи с интернетом.

Перед установкой Debian прочтите [тут](http://wiki.debian.org/Firmware)
и [тут](http://www.debian.org/releases/stable/i386/ch06s04.html.en) для
инструкций. Кратко:

  - перед установкой проверьте ваше оборудование на наличие проблем с
    драйверами (по ссылке выше есть список устройств);
  - скопируйте на съёмный носитель, например на CD/DVD или USB, тарболл
    взятый
    [отсюда](http://cdimage.debian.org/cdimage/unofficial/non-free/firmware/)
    и подключите его перед установкой. Если во время установки интернет
    доступен, то будет предложено скачать недостающий firmware.

Также доступны неофициальные netinst-iso-образы со всеми драйверами [по
ссылке](http://cdimage.debian.org/cdimage/unofficial/non-free/cd-including-firmware/).

Почему так было сделано? Мифы и факты об этом [есть по
ссылке](http://blog.schmehl.info/Debian/myths-and-facts-about-firmware).

## Откуда брать пакеты, отсутствующие в дистрибутиве?

Всё зависит от того, что вы хотите поставить. Большая часть необходимых
пакетов для стабильного релиза обнаруживается на
[1](http://backports.debian.org/). Как следует из названия, это проект,
занимающийся пересборкой некоторых программ из testing для stable.
Остальные пакеты, как правило, можно найти, воспользовавшись
поисковиком [2](http://www.apt-get.org/).

В частности, заслуживает внимания репозиторий для мультимедийных
программ (skype, google earth, adobe reader, различные кодеки),
которые из-за проблем с патентами не включены в основной репозиторий -
[3](http://deb-multimedia.org/)

А вот [тут](http://linuxoid.in/Полезные_репозитории_для_Debian) собран
список популярных репозиториев Debian.

Все вышеупомянутые ресурсы, за исключением
[4](http://backports.debian.org/), неофициальны, разработчики Debian не
оказывают никакой поддержки этим пакетам.

Если вам необходимо достать старый пакет, доступен архив по ссылке -
[5](http://snapshot.debian.org/)

## Пример оформления sources.list

Пример для Debian Squeeze. В случае с Sid (unstable) репозиторий
security.debian.org нужно отключить, так как для unstable ветки не
выпускаются исправления ошибок в безопасности.

Пояснение:

  - **deb** - бинарные пакеты, **deb-src** - пакеты с исходниками (если
    нужно пересобрать пакет на свой вкус);
  - **линк** репозитория;
  - **squeeze** - релиз Дебиана, может быть как название релиза
    (squeeze, lenny, wheezy, sid, ...), или по-другому (oldstable,
    stable, testing, unstable, experimental). Текущий релиз - stable, в
    данный момент это Debian Squeeze;
  - **main** - основная секция репозитория; **contrib** - пакеты в этой
    части распространяются владельцем авторских прав на условиях
    свободной лицензии, но зависят от несвободного программного
    обеспечения; **non-free** - лицензии пакетов в этой части содержат
    условия, ограничивающие использование или распространение ПО.


    # Debian repository
    deb http://ftp.ru.debian.org/debian/ squeeze main non-free contrib
    deb-src http://ftp.ru.debian.org/debian/ squeeze main non-free contrib

    # Security fixes
    deb http://security.debian.org/ squeeze/updates main contrib non-free
    deb-src http://security.debian.org/ squeeze/updates main contrib non-free

    # Debian updates
    deb http://ftp.debian.org/debian squeeze-updates main contrib non-free
    deb-src http://ftp.debian.org/debian squeeze-updates main contrib non-free

    # Backports
    deb http://backports.debian.org/debian-backports squeeze-backports main non-free contrib
    deb-src http://backports.debian.org/debian-backports squeeze-backports main non-free contrib

**Примечание:** строки с deb-src необязательны и могут быть
закомментированы для экономии трафика.

## Как редактировать список автозагрузки сервисов?

Для поклонников основанных на RedHat и Fedora дистрибутивов есть
замечательная программа sysv-rc-conf. Установим ее:

    root@linux# aptitude install sysv-rc-conf

Однако можно сделать тоже самое и вручную. Для этого достаточно сделать
нужный файл исполняемым или наоборот, если нужно отключить автозагрузку
сервиса, снять бит исполняемости:

    root@linux# chmod +x /etc/init.d/foo
    root@linux# chmod -x /etc/init.d/bar

Остальные могут воспользоваться "изкоробочной" утилитой
[update-rc.d](http://manpages.ylsoftware.com/ru/update-rc.d.8.html)

Не забывайте, что по умолчанию Debian использует второй уровень
автозагрузки (man init, man inittab).

## Как правильно установить проприетарные драйвера nVidia/ATI?

### Настройка драйвера nVidia

Установить драйвера nVidia можно одним из следующих способов.

В репозиториях Debian есть уже готовые скомпилированные бинарные модули
для дистрибутивного ядра. Например, для версии x86_64 установка будет
выглядеть как:

    root@linux# aptitude install nvidia-kernel-amd64 nvidia-glx nvidia-settings nvidia-vdpau-driver

nvidia-xconfig не используем, ибо она пишет в конфиг много ненужного.
Вместо этого пропишем в xorg.conf следующее:

    Section "Device"
        Identifier "video"
        Driver     "nvidia"
    EndSection

В большинстве случаев, этого конфига хватает для нормальной работы
иксов.

Есть и альтернативный вариант: можно установить nvidia-glx и
nvidia-kernel-dkms, в этом случае будут скачаны исходники и собран
модуль для текущего ядра при помощи dkms. Так же этот способ
подходит в том случае, если используется самосборное ядро или
установлен testing или unstable дистрибутив

Еще есть старый способ сборки модуля ядра с помощью module-assistant:

    root@linux# aptitude install module-assistant
    root@linux# module-assistant prepare
    root@linux# module-assistant update
    root@linux# module-assistant auto-install nvidia-kernel-source
    root@linux# nvidia-xconfig

Для поддержки аппаратного ускорения в 32-битных программах нужно
установить пакет nvidia-glx-ia32.

Также в репозитории присутствуют исходники legacy драйверов для
поддержки старых моделей видеокарт:

  - nvidia-kernel-legacy-71xx-source - для карт Vanta/TNT/GeForce256 и
    [подобного
    антиквариата](http://www.nvidia.com/object/linux-display-ia32-71.86.14-driver.html)
    (не поддерживаются Xorg 1.9 и выше)
  - nvidia-kernel-legacy-96xx-source - для карт GeForce2/4 MX [и
    некоторых
    других](http://www.nvidia.com/object/linux-display-ia32-96.43.19-driver.html)
  - nvidia-kernel-legacy-173xx-source - от Quadro, GeForce 5 FX [и
    старше](http://www.nvidia.com/object/linux-display-ia32-173.14.28-driver.html)

помимо исходников есть скомпилированные версии и модули
[DKMS](http://ru.wikipedia.org/wiki/Dynamic_Kernel_Module_Support).

### Настройка драйвера ATI

[6](http://wiki.debian.org/ATIProprietary) - Установка драйвера из
репозитория

[7](http://my-note.ru/archives/136.html) - Установка более новой, чем в
репозитории, версии

[8](http://debian.oppserver.net/fglrxlegacy/) - Репозиторий с драйверами
для RadeonHD 2000, 3000 и 4000 серий

### Быстрый (хоть и идеологически неверный) вариант

Также для установки драйверов можно воспользоваться утилитой
[sgfxi](http://code.google.com/p/sgfxi/).

sgfxi это автоматический инсталлятор и конфигуратор видеодрайверов для
nvidia, ati/amd, а также свободных драйверов из состава Xorg для Debian
(с некоторых пор добавлена поддержка Ubuntu и Arch). При каждом запуске
скрипт автоматически обновляется, что позволяет с большой долей
вероятности получить работоспособные драйвера даже для самых
свежих ядер.

Скачиваем:

    user@linux$ wget http://sgfxi.googlecode.com/svn/trunk/sgfxi

Делаем исполняемым:

    user@linux$chmod +x sgfxi

Выключаем иксы. Внимательно изучив справку можно обойтись и без этого,
но в любом случае их придётся перезагружать:

    root@linux# init 3

Запускаем:

    root@linux# ./sgfxi

Скрипт проверит актуальность своей версии, обновится при необходимости,
выяснит модель видеокарты, версию ядра, установит необходимые пакеты,
применит патчи и соберет библиотеки и модули. Если все пройдет успешно
- сам предложит запустить иксы.

## Как добавить ключ, которым подписаны пакеты в репозитории?

Зачастую при попытке установить пакет из стороннего репозитория
выводится выводится сообщение об ошибке, например такое:

> W: GPG error: <http://download.virtualbox.org> lenny Release:
> Следующие подписи не могут быть проверены, так как недоступен
> открытый ключ: NO_PUBKEY DCF9F87B6DFBCBAE

Добавить недостающий ключ можно так:

    root@linux# gpg --keyserver wwwkeys.eu.pgp.net --recv-keys DCF9F87B6DFBCBAE
    root@linux# gpg --armor --export DCF9F87B6DFBCBAE | apt-key add -
    root@linux# apt-get update

Для удобства вы можете сделать скрипт, который будет добавлять в кеш
переданный ему ключ. Он поможет в будущем легко импортировать другие
ключи.

## Как быстро собрать deb пакет?

[Вот так](http://barrel-of-herring.blogspot.com/2010/08/deb.html)

## Как подкорректировать сглаживание шрифтов в системе?

[Настройка сглаживания шрифтов](http://wiki.debian.org/ru/PatchCairo)

## Где можно получить ответы на вопросы, которых здесь нет?

Ответы на остальные вопросы можно получить в [Debian
FAQ](http://www.debian.org/doc/FAQ/), [Debian
Forums](http://forums.debian.net/), [Debian Q/A](http://ask.debian.net/)
или в соответствующих рассылках (см. [9](http://lists.debian.org/)).

Существует так же Debian Wiki по адресу [10](http://wiki.debian.org/),
русскоязычный раздел [находится
здесь](http://wiki.debian.org/ru/DebianRussian).

## Статьи в этой wiki

  - [Установка Debian через другой GNU/Linux
    (debootstrap)](Установка_Debian_через_другой_GNU_Linux_\(debootstrap\))

## Ссылки

  - [Документация Debian](http://www.debian.org/doc/)
  - [Debian Wiki](http://wiki.debian.org)
  - [Поиск пакетов по официальному
    репозиторию](http://packages.debian.org/)
  - [Полезные репозитории для
    Debian](http://linuxoid.in/Полезные_репозитории_для_Debian)
  - [Debian
    Reference](http://www.debian.org/doc/manuals/debian-reference/)
  - [Debian GNU/Linux Installation
    Guide](http://www.debian.org/releases/stable/i386/index.html.ru)
  - [Debian Stable Release
    Notes](http://www.debian.org/releases/stable/i386/release-notes/index.ru.html)
  - [Debian FAQ](http://www.debian.org/doc/manuals/debian-faq/)
  - [Debian New Maintainers'
    Guide](http://www.debian.org/doc/manuals/maint-guide/)
  - [Debian Developer's
    Reference](http://www.debian.org/doc/manuals/developers-reference/)
  - [The 101 most important things when using Debian
    GNU/Linux](http://tangosoft.com/refcard/)
  - [Debian Administration](http://www.debian-administration.org/)
  - [Planet Debian](http://planet.debian.net/)
  - [Apt-get install debian-wizard: blog about
    debian](http://raphaelhertzog.com)

## Сноски


<references />



1.  <http://www.debian.org/doc/manuals/project-history/ch-intro.en.html#s1.2>