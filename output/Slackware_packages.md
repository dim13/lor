## А есть ли в Slackware менеджер пакетов?

Да. Не слишком навороченный, но со своими функциями вполне справляется.
Основным пакетным менеджером Slackware является pkgtools, на
носителе/репозитории поставляется ap/slackpkg.

Установить новые пакеты, удалить, просмотреть список установленных, а
также провести базовую настройку системы (мышь, сеть, временная зона
и т.д.) можно с помощью утилиты pkgtool. Не хуже с этими процедурами
справляются и отдельные утилиты - installpkg, removepkg или
upgradepkg.

-----

Так, например, для того, чтобы обновить glibc можно сделать
(предварительно прочитав UPGRADE.TXT):

`root@linux# upgradepkg /root/slackware*/a/glibc-solibs-*.t?z`

  -
    А для обновления всей системы с установкой новых пакетов:

`root@linux# upgradepkg --install-new /root/slackware*/*/*.t?z`

  -

      -
        правда при этом поставятся все пакеты нового дистрибутива.

-----

Далее, как уже упоминалось, в каталоге /ap присутствует фронтенд для
pkgtools - slackpkg. К [slackpkg](http://www.slackpkg.org) есть плагин
поддержки нескольких репозиториев
[slackpkg+](http://slakfinder.org/slackpkg+.html)

Существуют и другие менеджеры пакетов для Slackware:

[spkg](http://spkg.megous.com/) - "The Unofficial Slackware Linux
Package Manager", тоже довольно неплохой менеджер пакетов и который
может быть необходим, если вы предпочтете gslapt/slapt-get от
Salix, там именно через него выполняются некоторые операции, если не во
всех версиях этого дистрибутива, то по крайне мере в 14.1 точно.

Также весь функционал slackpkg предоставляет фронтенд к целому ряду
популярных пакетных менеджеров под названием
[smartpm](http://labix.org/smart).

Дополнительно можно установить фронтенды с поддержкой **зависимостей**
пакетов [swaret](http://swaret.sourceforge.net/) и
[slapt-get](http://software.jaos.org). У slapt-get есть возможность
указания приоритета для каждого репозитория: "default \< official
\< preferred \< custom"
[1](http://forum.salixos.org/viewtopic.php?f=23&t=15#p159) .

Для pkgtools также существовал графический интерфейс на Qt -
[SlackIns](http://slackins.sourceforge.net), а для slapt-get -
[gslapt](http://software.jaos.org/).

Для **сборки** пакетов из исходников есть [sbopkg](http://sbopkg.org/) и
[slapt-src](http://software.jaos.org/). К slapt-src есть программа с
графическим интерфейсом -
[Sourcery](http://www.salixos.org/wiki/index.php/Sourcery).

## Как собрать пакет?

Написать SlackBuild, скрипт для сборки пакета - [SlackBuild
scripts](http://docs.slackware.com/slackware:slackbuild_scripts). Таким
способом собираются пакеты в основной системе, можно залезть в /source
и посмотреть примеры, н-р:
[bash](http://mirror.yandex.ru/slackware/slackware-current/source/a/bash/).

С помощью утилиты **makepkg**, где-то так:

    user@linux:~$ cd /usr/src/program-name
    user@linux:/usr/src/program_name$ ./configure
    user@linux:/usr/src/program_name$ make
    user@linux:/usr/src/program_name$ su
    root@linux:/usr/src/program_name# make install DESTDIR=/tmp/program-name
    root@linux:/usr/src/program_name# cd /tmp/program-name
    root@linux:/tmp/program_name# makepkg ../program-name-ver-arch-build.txz
    root@linux:/tmp/program_name# cd ..
    root@linux:/tmp/program_name# installpkg program-name-ver-arch-build.txz
    root@linux:/tmp/program_name# cd && rm -R /tmp/program-name

Если ничего из вышеперечисленного не нравится, использовать
[src2pkg](http://src2pkg.net/).

## Где найти готовые пакеты?

В первую очередь это /packages на
[slackware.com](http://www.slackware.com/packages/) . Здесь в принципе
все пакеты, которые вошли в релиз или входят в current.

Репозитории [AlienBOB](http://www.slackware.com/~alien/) :

  - [Alien's Slackware
    packages](http://taper.alienbase.nl/mirrors/people/alien/sbrepos/) -
    репозиторий в который время от времени Alien собирает и закидывает
    пакеты;
  - [Alien's 'ktown'
    repository](http://taper.alienbase.nl/mirrors/alien-kde/) - его
    репозиторий для KDE. Через этот репозиторий Вы сможете обновить
    Ваше KDE. Как обновлять в нем же для вашей ветки дистрибутива
    читайте README там все пошагово описано, н-р, последняя
    сборка для current:
    [README](http://taper.alienbase.nl/mirrors/alien-kde/current/latest/README)
    (и даже рассказано, если захотите пересобрать KDE сами на своей
    машине). Репозиторий применим для релиза (т.е. н-р: 14.1) и
    -current.

Salix:

  - <http://download.salixos.org/> или одно из их зеркал
    [Repository_mirrors](http://docs.salixos.org/wiki/Repository_mirrors)
    . Т.к. дистрибутив берет основную пакетную базу у slackware, далее
    они уже сами собирают отдельные пакеты и закидывают в свой реп, то
    почему бы не воспользоваться. Только тут стоит понимать различия,
    вот, н-р, срез репозитория от slackware:
    [download.salixos.org/x86_64/slackware-14.1](http://download.salixos.org/x86_64/slackware-14.1/)
    , конечно возможно что-то там пересобрали они, но в основном
    пакетная база слаки, а вот собранное ими
    [download.salixos.org/x86_64/14.1](http://download.salixos.org/x86_64/14.1/).

Slackel:

  - [Slackel Repository](http://www.slackel.gr/repo/)- репозиторий
    дистрибутива который базируется на -current и salix, но
    встречаются пакеты, которых нет в этих 2 репозиториях, н-р:
    [2](http://zenway.ru/page/slackware-torbrowser)

MSB:

  - [MATE SlackBuilds](http://mateslackbuilds.github.io/) (MSB) - там же
    адрес на репозиторий пакетов MATE (форка GNOME2). Как устанавливать
    читаем [MSB_SLACKPKG+](http://slackware.org.uk/msb/MSB_SLACKPKG+),
    в принципе установка ничем не отличается от той как если бы Вы
    ставили, обычные пакеты или к примеру пакеты из KTown, но если
    затруднения делайте, как они советуют.

SlackE17:

  - [SlackE17 - e17 packages for
    Slackware](http://slacke17.sourceforge.net/) - пакеты для установки
    e17 в Вашей Slackware.

SlaxXBMC:

  - [SlaxXBMC](http://slaxbmc.blogspot.ru/) - проект хоть и
    предоставляет свою версию дистрибутива, но в тоже время
    и собранные пакеты xbmc доступны и совместимы со slackware

WINE:

  - [Wine Is Not an Emulator (WINE) -
    пакеты](http://sourceforge.net/projects/wine/files/Slackware%20Packages/)
    для Slackware

Далее, на <http://linuxpackages.net> и <http://slacky.eu>. Для поиска
пакетов можно воспользоваться сайтом
[slakfinder.org](http://slakfinder.org/). Оба репозитория можно
подключить к slapt-get и swaret, для этого читайте faq на
соответствующих сайтах

Еще есть <http://darkstar.ist.utl.pt/slackware/addon>, на котором
собрано довольно много из выше перечисленных бинарных пакетов и
скриптов SlackBuild (имеются также альтернативные сборки Gnome).

Кроме того можно брать пакеты из Zenwalk и, вероятно, Vector Linux
(новые пакеты сжимаются lzma и имеют расширение .tlz. Установка
.tlz пакетов требует наличие Slackware версии старше 12.2. Slackware
13.0 использует второе поколение lzma - xz , пакеты для 13-й версии
имеют расширение .txz)

## Где найти готовые SlackBuild'ы для сборки пакетов?

  - Множество готовых SlackBuild'ов лежат на
    [SBo](http://www.slackbuilds.org) (slackbuilds.org), с их помощью
    можно легко и просто собрать программу с нужными опциями,
    достаточно лишь положить в ту же директорию архив с
    исходниками и запустить скрипт.


  - В большинстве случаев в самих собранных пакетах сборщики кладут
    SlackBuild'ы в пакет, открыв его в своем любимом архиваторе, найти
    можно в usr/doc/$PRGNAM-$VERSION. Те, кто не боятся проблем, могут
    использовать Slackware
    [Current](http://www.slackware.com/changelog/current.php?cpu=i386) -
    текущее состояние разработки дистрибутива.


  - Так же, в практически во всех репозиториях (которые были приведены
    выше) можно найти эти билды. Лежать они могут в том месте, где и
    сам пакет или в каталоге с названием src или source. Вот, как это
    реализовано у самого дистрибутива slackware в качестве примера
    [MPlayer](http://mirrors.slackware.com/slackware/slackware-current/source/xap/MPlayer/)


  -
    Стоит отметить, если Вы будете искать SlackBuild'ы у Salix , то
    скорей всего Вы их не найдете, т.к. там используется
    [SLKBUILD](http://docs.salixos.org/wiki/Building_packages_with_slkbuild)
    , ничего страшного нет пакеты все равно совместимы, просто
    скрипт/системы сборки у них такой/ва и это все равно не
    должно Вам мешать подсмотреть ключи, операции над исходниками во
    время сборки, если есть необходимость при написании SlackBuild'а.

Плюс еще несколько дополнительных репозиториев:

  - [Alien's SlackBuilds](http://www.slackware.com/~alien/slackbuilds/)
    содержит пакеты и билды, собранные при помощи них
  - [PhantomX](https://github.com/PhantomX/slackbuilds) , a collection
    of personal SlackBuilds
  - [willysr](https://github.com/willysr/SlackHacks/tree/master/SlackBuilds)
    repositories SlackBuilds

## Где ещё могут быть найдены уже собранные пакеты и написанные SlackBuild'ы?

Пользоваться готовыми пакетами от сторонних источников крайне не
рекомендуется, т.е. это может противоречить в первую очередь
безопасности и, вообще-то, стабильности уже собранного пакета, но
никто не запрещает воспользоваться готовым слакбилдом, просмотреть
его (и файлы к нему: описание, патчи, скрипты какие-нибудь) и на
основе его собрать пакет, так что, как говорится: "доверяй, но
проверяй"(c).

  - <http://slakfinder.org/slackpkg+/src/repositories.txt> - Список
    поддерживаемых репозиториев для slackpkg+, в которых имеются
    уже собранные пакеты и которые также можно установить при помощи
    slackpkg+ (что это такое - читать выше)
  - <http://slackware.org.uk/> - Зеркало, на котором хранятся
    отзеркаленные репозитории различных
    slack-дистров-проектов, в том числе есть и пакеты, и
    слакбилды
  - <http://www.microlinux.fr/slackware/> - Extra Software for Slackware
    Linux, пакеты и slackbuild'ы проекта
    [MLED](http://zenway.ru/page/meld), ссылка на
    [git](https://github.com/kikinovak/slackware)
  - <http://rlworkman.net/pkgs/> - Robby Workman's Slackware Packages
  - <http://schoepfer.info/slackware.xhtml> - походу, реп. пользователя,
    совместим со slapt-get,slackpkg
  - <http://www.slackers.it/packages/> - пакеты и slackbuild'ы ,
    [git](https://github.com/conraid/SlackBuilds) написанных
    SlackBuild'ов из этого ресурса
  - <http://ponce.cc/slackware/> - пакеты и slackbuild'ы, присутствуют
    lxde и razor-qt, есть поддержка rsync, информация от автора на
    [LQ](http://www.linuxquestions.org/questions/slackware-14/sbo-git-slackbuilds-org-on-git-with-patches-for-current-794521/)
  - <http://www.dawoodfall.net/> - David Woodfall's SlackBuilds and
    Scripts
  - <https://slackonly.com/> - [Slackonly: packages built from
    SBo](https://slackalaxy.wordpress.com/2014/04/12/slackonly-packages-built-from-sbo/)
    - [SBo в бинарном
    варианте](http://slackware.su/forum/index.php?topic=1705.0)


  - <http://slakfinder.org/>
  - <http://pkgs.org/>
  - <http://people.salixos.org/>


  - <http://slack.isper.sk/pub/>
  - <http://www.droplinegnome.org/#download>
  - <http://sourceforge.net/projects/portable/files/>
  - <http://www.nielshorn.net/slackware/slack_pkg.php>


  - <http://www.andreasliebe.de/pkgs/>
  - <http://slackware.sukkology.net/>

## Где взять GNOME?

Несмотря на то, что Патрик начиная с версии 10.2 удалил Gnome из
дистрибутива, на данный момент существуют сторонние сборки гнома
для Slackware:

  - [Dropline](http://www.droplinegnome.org) - Следует помнить что
    Dropline дополнительно устанавливает в систему PAM и заменяет
    некоторые системные пакеты из категорий l/, x/ и xap/.

Экстремалы могут воспользоваться системой сборки
[Garnome](http://www.gnome.org/projects/garnome/)

### MATE

  - [MSB](https://www.linux.org.ru/wiki/en/Slackware#a_.D0.93.D0.B4.D0.B5_.D0.BD.D0.B0.D0.B9.D1.82.D0.B8_.D0.B3.D0.BE.D1.82.D0.BE.D0.B2.D1.8B.D0.B5_.D0.BF.D0.B0.D0.BA.D0.B5.D1.82.D1.8B.3F)
    (MATE SlackBuilds) - пакеты и SlackBuild'ы
  - [Salix Mate](http://www.salixos.org/screenshots.html) - версия Salix
    с MATE, для установки на систему без данного DE можно
    воспользоваться их нем же репозиторием, установив
    пакеты из salix-репозитория
  - [MATE installation](http://wiki.mate-desktop.org/download#slackware)
    - документация и информация по репозиториям и установке MATE в
    Slackware и Salix на Mate Desktop Environment wiki

### Cinnamon

  - [Cinnamon SlackBuilds](https://github.com/willysr/csb) от willysr