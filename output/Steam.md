**Steam** - система цифровой дистрибуции игр с элементами социальной
сети. С помощью приложения Steam можно легально приобретать игры и
немедленно загружать их из интернета. В Steam доступно несколько тысяч
игр. Steam стал популярным приложением среди геймеров. Были выпущены
версии для Windows, MAC OS X и Linux. Некоторые новые игры
невозможно легально приобрести нигде, кроме Steam.

Steam разрабатывается и поддерживается компанией **Valve**. Помимо
Steam, Valve создала много других программных продуктов, ставшие
популярными. Игровой движок Source, игры Counter Strike: Source,
Half Life 2, Portal и Team Fortress 2 созданы компанией Valve.

Во всех играх Valve, а также во многих играх от других разработчиков,
можно находить партнёров для сетевой игры в системе Steam. Также
можно играть с друзьями из "списка друзей" вашего аккаунта. Кроме
поиска партнёров для сетевой игры Steam также обеспечивает VoIP. Кроме
того, у многих игр, достуапных в Steam, доступна система достижений.

Ниже перечислены проблемы при запуске Steam для Linux и некоторых игр,
доступных в нём. При написании этого текста использовались материалы с
Gentoo и Arch Wiki's, а также личный опыт авторов этой статьи. Это не
все доступные в Steam игры - ознакомиться с полным списком можно [на
этой странице](http://steamdb.info/linux/).

# Клиент Steam для Linux

На момент написания этих слов (конец апреля 2013 года) единственный
поддерживаемый дистрибутив Linux дистрибутивом Steam для Linux -
Ubuntu Linux. Однако ничто не мешает запустить бинарные файлы Steam в
любом другом дистрибутиве, который соответствует минимальным
системным требованиям:

  - Соответствие текущему стандарту
    [LSB-Desktop 4.1](http://refspecs.linuxfoundation.org/LSB_4.1.0/LSB-Desktop-IA32/LSB-Desktop-IA32/requirements.html#RLIBRARIES)
  - GLIBC 2.15 и старше
  - libstdc++46 (из GCC 4.6) и старше
  - Процессорная архитектура x86 (бинарные файлы Steam могут быть
    запущены в системе для архитектуры x86_64)

С сайта разработчика необходимо загрузить файл установки (пакет DEB либо
архив tar.gz), после запуска которого будет загружено 150 Мб данных,
необходимых для Steam.

**Steam Runtime** - набор системных библиотек, гарантирующий запуск игр
Valve в Linux независимо от используемого дистрибутива.

## Установка в дистрибутивах Linux, отличных от Ubuntu

### Arch Linux

Steam находится в официальном репозитории.

` # pacman -S steam`

Дополнительную информацию по решению проблем можно найти в
[ArchWiki/Steam](https://wiki.archlinux.org/index.php/Steam)

### Debian

#### Если используется открытый драйвер видеокарты

Если вы используете открытый драйвер видеокарты, то для нормальной
работы Steam в Debian Wheezy рекомендуется Mesa версии 9.0 и выше.
В момент написания этих слов (середина мая 2013 года) в репозитории
Debian доступна только восьмая версия. Установить Mesa 9 можно из
исходного кода. Для этого дополните список репозиториев в файле
/etc/apt/sources.list строками DEB-SRC и выполните apt-get update.
Выполните команду apt-get build-dep libdrm mesa. Загрузите [с этой
страницы](http://packages.ubuntu.com/raring/libdrm2) пакеты DEB-SRC с
libdrm, [а с этой](http://packages.ubuntu.com/raring/libgl1-mesa-glx) -
пакеты с Mesa. Ниже показана установка libdrm из консоли, установка
Mesa осуществляется таким же способом.

` $ su`
` Пароль:`
` # apt-get build-dep libdrm mesa`
` # exit`
` $ mkdir build && cd build`
` $ wget `<http://archive.ubuntu.com/ubuntu/pool/main/libd/libdrm/libdrm_2.4.43-0ubuntu1.dsc>
` $ wget `<http://archive.ubuntu.com/ubuntu/pool/main/libd/libdrm/libdrm_2.4.43.orig.tar.gz>
` $ wget `<http://archive.ubuntu.com/ubuntu/pool/main/libd/libdrm/libdrm_2.4.43-0ubuntu1.diff.gz>
` $ wget `<http://archive.ubuntu.com/ubuntu/pool/main/m/mesa/mesa_9.1.1-0ubuntu3.dsc>
` $ wget `<http://archive.ubuntu.com/ubuntu/pool/main/m/mesa/mesa_9.1.1.orig.tar.gz>
` $ wget `<http://archive.ubuntu.com/ubuntu/pool/main/m/mesa/mesa_9.1.1-0ubuntu3.diff.gz>
` $ dpkg-source -x libdrm_2.4.43-0ubuntu1.dsc`
` $ cd libdrm_2.4.43-0ubuntu1`
` $ fakeroot ./debian/rules binary`
` $ cd ..`
` $ rm *dbg*.deb`
` $ su`
` Пароль:`
` # dpkg -i *.deb`
` # exit`
` $ rm -f *.deb`

Знак "$" означает что команда выполняется от имени пользователя, а знак
"\#" - от имени администратора. Также желательно установить
[libtxc_dxtn](http://packages.debian.org/wheezy/libtxc-dxtn-s2tc0), а
если у вас видеокарта Intel, то ещё и [последнюю версию драйвера
Intel](https://01.org/linuxgraphics/downloads).

### Slackware

  - [Valve’s Steam
    client](http://alien.slackbook.org/blog/valves-steam-client-for-linux/)
    - заметка в блоге AlienBOB'а
  - [steamclient](http://www.slackware.com/~alien/slackbuilds/steamclient/)
    - Пакеты и SlackBuild'ы

### Fedora

Steam находится в [RPMFusion](http://rpmfusion.org/). После
[подключения](http://rpmfusion.org/Configuration) достаточно
установить пакет steam, и он установит себя вместе со всеми
зависимостями

` # yum install steam`

### Gentoo

Проще всего клиент steam поставить через оверлей "steam" или "gamerlay".
О том, как в gentoo устанавливаются оверлеи можно почитать в
gentoo-wiki: <http://wiki.gentoo.org/wiki/Layman>

Добавляем оверлей steam

`layman -a steam`

или gamerlay

`layman -a gamerlay`

и обновляем добавленные в систему оверлеи:

`layman -S`

Чтобы установить клиент steam достаточно выполнить одну команду:

`emerge -v steam-meta`

Пакет steam-meta устанавливает все пакеты необходимые для запуска steam:

  - `steam-client-meta`: мета-пакет, проверяющий наличие всех
    зависимостей для steam-клиента
  - `steam-games-meta`: мета-пакет, проверяющий наличие зависимостей для
    конкретных игр
  - `steam-launcher`: непосредственно лаунчер steam-клиента

USE-флаг `steamruntime` определяет будут ли использоваться поставляемые
в комплекте библиотеки для запуска клиента и игр. Но раз уж вы выбрали
себе генту, то видимо хотите использовать свои собственные, любовно
собранные системные библиотеки. Steam использует sdl-2, которая
отсутствует в официальном дереве, ее придется установить отдельно,
из того же самого оверлея steam или gamerlay:

`ABI_X86="32 64" emerge -v libsdl:2`

После этого установка и запуск steam без рантайма должны пройти без
проблем.

<b>Использование 32-bit chroot в 64-битной системе</b>

Если в 64-битной системе не хочется связываться с библиотеками
"emul-linux-x86-\*" или "multilib", то можно использовать 32-битное
окружение chroot. Подробнее о его создании для запуска steam можно
прочитать gentoo-wiki:
<http://wiki.gentoo.org/wiki/Steam#32-bit_chroot_on_amd64>

При создании 32-битного окружения для нормальной работы opengl с
открытыми драйверами скорее всего понадобится собрать mesa c
USE-флагом `xa` (помимо флагов для используемой в системе видеокарты).

### Mageia Linux

` # urpmi steam`

## Игры работают, запускаются, но текстур нет или все черные

Убедитесь, что стоит libtxc_dxtn и её 32-битная версия

## Video Driver Problem: Your system is running older proprietary nVidia video drivers. Steam requires driver version 304.22 or higher.

Valve установила минимально необходимые версии драйверов видеокарт
NVIDIA, это 304.22.

## Я купил диск с игрой, которая требует привязки к Steam, и я знаю что она есть для Linux. Как её установить?

[Решение](https://support.steampowered.com/kb_article.php?ref=5357-FSQM-0382)

` steam -install "/media/Left 4 Dead"`

Где "/media/Left 4 Dead" - каталог, в который примонтирован диск.
Кавычки нужны только если в пути каталога есть пробелы.

## Не работают видеоролики во встроенном браузере

Скорее всего у вас 64-битная система, так как в 32-битной Steam сам
находит библиотеку Adobe Flash Player. Загрузите с сайта
<http://adobe.com/> tar.gz-файл с 32-битным Adobe Flash Player.
Распакуйте файл libflashplayer.so. Положите его в один из двух
каталогов: /usr/lib/browser-plugins/ или
\~/.local/share/Steam/ubuntu12_32/plugins/

Под "\~/" подразумевается домашний каталог пользователя, а чтобы увидеть
скрытый каталог ".local" необходимо поставить галочку "Показывать
скрытые файлы и каталоги" в третьем пункте меню файлового
менеджера. Периодически обновляйте этот файл, так как в Adobe
Flash Player иногда находят сетевые уязвимости.

## Не запускается игра. Ошибок нет. В консоли ошибка "No such file or directory"

Вы используете профиль Steam из Windows или Wine? Посмотрите на ошибку
подробнее. Steam пытается найти каталог "SteamApps", а у вас
"steamapps". Дело в том что для файловой системы Linux большая и
маленькая буква - разные, а для файловой системы Windows -
одинаковая. Для Linux эти два каталога разные, а для Windows -
один и тот же. Переименуйте. Затем переименуйте "team fortress 2" в
"Team Fortress 2". И так пока не будет идентично пути каталога в
сообщении об ошибке.

# Игры на движке Source

# Общие проблемы

## Required OpenGL extension "GL_EXT_texture_sRGB_decode" is not supported. Please update your OpenGL driver.

[Скриншот ошибки](http://img824.imageshack.us/img824/764/v5in.png). Если
при запуске игр на движке Source возникает такая ошибка, необходимо
обновить драйвер видеокарты. NVIDIA: 304.xx, ATi/AMD: 12.11. Если
ваша видеокарта уже не поддерживается новым драйвером, воспользуйтесь
[этим способом обойти
проблему](http://steamcommunity.com/app/221410/discussions/0/846938351012409765/).
Кстати, AMD выпустила обновление драйвера для видеокарт Radeon HD
2xxx-4xxx версии
[13.1-legacy](http://support.amd.com/us/gpudownload/linux/legacy/Pages/legacy-radeon_linux.aspx).

## Ещё одно исправление для игр на движке Source

Если вы используете открытый драйвер видеокарты, установите библиотеку
[libtxc_dxtn](http://pkgs.org/search/?keyword=libtxc_dxtn). Тогда
появится расширение OpenGL GL_S3_S3TC, которое используют игры
на движке Source.

## Team Fortress 2 останавливает загрузку на вступительной заставке, жёсткий диск простаивает

Просто нажмите Enter. Причина в невидимой ошибке, закрытой полноэкранной
игрой. [Вот её
скриншот](http://img834.imageshack.us/img834/6592/fnq.png).

*SetLocale('en_US.UTF-8') failed. Using 'C'. You may have limited glyph
support. Please install 'en_US.UTF-8' locale.*

Решается запуском Steam не командой:

` steam`

А командой:

` LC_ALL=en_US.UTF-8 steam`

Можно отредактировать ярлык запуска программы. Если у вас Gentoo,
убедитесь что эта локаль в системе существует. Если у вас нет
звука, то это не полная команда запуска Steam. Речь об этом пойдёт
ниже.

## Нет звука

Возможно, у вас в системе нет звукового сервера PulseAudio, или он
выключен в настройках системы? Steam пока не научился определять
звуковую систему ALSA сам. Запускайте его командой:

` SDL_AUDIODRIVER=alsa steam`

И тогда звук появится. А лучше совместить два совета вместе:

` SDL_AUDIODRIVER=alsa LC_ALL=en_US.UTF-8 steam`

Так же звук может отсутствовать, если не установлен OpenAL, требуемый
для некоторых игр (Half-Life 1, например). Решается ручной установкой
соответствующего пакета.

## Нет звука (если не помог первый вариант)

Добавляем эту строку в \~/.bashrc:

` export SDL_AUDIODRIVER=pulseaudio`

Добавляем эту строку в /etc/openal/alsoft.conf:

` drivers = pulse,alsa,core,oss,solaris,sndio,mmdevapi,dsound,winmm,port,opensl,null,wave`

## Сворачивание в трей

Добавить в .bashrc

` export STEAM_FRAME_FORCE_CLOSE=1`

Чтобы восстановить окно программы, необходимо выбрать один из пунктов в
меню, вызываемом кликом правой кнопкой мыши на иконке в трее.

## Запуск игр с использованием видеокарты Nvidia Optimus

Установите пакет `bumblebee` для вашего дистрибутива в соответствии с
инструкцией на официальном сайте:
<http://bumblebee-project.org/install.html>

Рекомендуется установить пакет `primusrun` (пакет и инструкции по
установке - <https://github.com/amonakov/primus>), как более
новый и производительный backend к bumblebee. После его установки
достаточно будет в свойствах выбранной игры в Steam для "SET LAUNCH
OPTIONS" прописать команду

`primusrun %command%`

# Team Fortress 2

Начиная с 2011 года эта игра распространяется бесплатно. Добавьте её в
список игр и нажмите "Установить". Если свободного места на диске
достаточно (установленная игра занимает 12 Гб), начнётся загрузка
игры. Рассмотрим проблемы работы установленной игры.

## Квадратики вместо букв в некоторых местах игры

[Скриншот](http://ic.pics.livejournal.com/leperflesh/16536437/188897/188897_original.jpg).
Если у вас Gentoo, выполните эту команду:

` eselect fontconfig enable 70-no-bitmaps.conf`

[Скриншот, где всё выглядит
правильно](http://ic.pics.livejournal.com/leperflesh/16536437/189411/189411_original.jpg).
[Совет взят
отсюда](http://steamcommunity.com/app/221410/discussions/2/846939615090806472/).
Также я [встречал
совет](http://www.linux.org.ru/news/games/8848991?cid=8849504)
удалить скрытый каталог ".fonts" в /home/user, и
[совет](http://www.linux.org.ru/news/games/8848991?cid=8849793)
включить локаль en_US.UTF-8.

## Игра не запускается, если каталог SteamApps - симлинк

[Описание
проблемы](http://steamcommunity.com/app/221410/discussions/1/882966056549274409/?l=russian).
Сделайте каталог обычным, и всё заработает.

Или просто переместите всю директорию со steam на другой раздел и
запустите хотя бы раз скрипт steam.sh из нового
месторасположения.

А еще можно просто примонтировать каталог при помощи mount -o bind.

# Описание проблем с играми

# Left 4 Dead 2

## При запуске появляется ошибка Failed to create OpenGL device\! (fglrx-legacy)

Решение описано
[тут](http://steamcommunity.com/app/550/discussions/0/864971765979513826#c864972620729210207)
(при условии, что драйвера установлены нормально, TF2 и "Left 4 Dead 2
Beta", без скобок запускаются нормально):

Обновлённый файл конфигурации в официальном выпуске Left 4 Dead 2 в
параметре anti-alias (setting.mat_antialias) установлен в "2", что
в некоторых случаях может привести к возникновению этой ошибки. После
установки значения в "1" или "0", L4D должен запуститься.

При стандартной установке Steam, путь к файлу будет
\~/.local/share/Steam/SteamApps/common/Left\\ 4\\ Dead\\
2/left4dead2/cfg/video.txt

Или

  - Запустить Steam, перейти в "Library";
  - ПКМ на "Left 4 Dead 2 (Beta)" в списке библиотеки;
  - выбрать "Properties", потом вкладку "Local files";
  - нажать "browse local files...";
  - в файловом менеджере перейти в директорию "left4dead2", затем в
    "cfg";
  - открыть файл "video.txt";


  - найти параметр "setting.mat_antialias" и заменить значение "2" на
    "1";
  - сохранить;
  - перезапустить Steam (на всякий случай);
  - попытаться запустить "Left 4 Dead 2 (Beta)".

# Portal

## Portal: Prelude

Описание взято
[отсюда](http://steamcommunity.com/app/221410/discussions/2/828938532549623359/?l=russian#c828938532626681170).
Описания по необходимым изменениям для других модов игр на движке
SourceMods (для HL2 и Portal) собраны
[здесь](http://steamcommunity.com/app/221410/discussions/0/828938532638150996/)

При запуске мода Portal: Prelude происходит попытка вызова hl.exe вместо
hl.sh Перейдите в директорию с игрой Portal и создайте соответствующий
symlink, например:

`cd ~/.local/share/Steam/SteamApps/common/Portal`

`ln -s hl2.sh hl2.exe`

Так как ресурсы игры были перемещены из файлов .gcf в .vpk, необходимо
также внести изменения в файл `gameinfo.txt` (далее приводятся
изменения внесённые автором упомянутой выше заметки):

` "GameInfo"`
`     {`
`           game                                            "Portal: Prelude"`
`           title                                           "Portal: Prelude"`
`           type                                            singleplayer_only`
`          `
`           nodifficulty                                    1`
`           hasportals                                      1`
`           nocrosshair                                     1`
`           nomodels                                        1`
`          `
`           developer                                       "Nicolas 'NykO18' Grevet"`
`           developer_url                                   "`<http://www.portalprelude.com>`"`
`           manual                                          "`<http://www.portalprelude.com/forum>`"`
`           icon                                            "resource/prelude"`
`          `
`           FileSystem`
`           {`
`                   SteamAppId                              400`
`                  `
`                   SearchPaths`
`                   {`
`                           game+mod                        |gameinfo_path|.`
`                           platform                        |gameinfo_path|.`
`    `
`                           // We search VPK files before ordinary folders, because most files will be found in`
`                           // VPK and we can avoid making thousands of file system calls to attempt to open files`
`                           // in folders where they don't exist.  (Searching a VPK is much faster than making an operating`
`                           // system call.)`
`                           game_lv                         portal/portal_lv.vpk`
`                           game+mod                        portal/portal_english.vpk`
`                           game+mod                        portal/portal_pak.vpk`
`                           game                            |all_source_engine_paths|hl2/hl2_textures.vpk`
`                           game                            |all_source_engine_paths|hl2/hl2_sound_vo_english.vpk`
`                           game                            |all_source_engine_paths|hl2/hl2_sound_misc.vpk`
`                           game                            |all_source_engine_paths|hl2/hl2_misc.vpk`
`                           platform                        |all_source_engine_paths|platform/platform_misc.vpk`
`    `
`                           // Now search loose files.  We'll set the directory containing the gameinfo.txt file`
`                           // as the first "mod" search path (after any user customizations).  This is also the one`
`                           // that's used when writing to the "mod" path.`
`                           mod+mod_write+default_write_path        |gameinfo_path|.`
`    `
`                           // Add the mod directory as a game search path.  This is also where where writes`
`                           // to the "game" path go.`
`                           game+game_write                 |gameinfo_path|.`
`    `
`                           // Where the game's binaries are`
`                           gamebin                         portal/bin`
`    `
`                           // Last, mount in shared Portal loose files`
`                           game                            |all_source_engine_paths|portal`
`    `
`                           // Last, mount in shared HL2 loose files`
`                           game                            |all_source_engine_paths|hl2`
`                           platform                        |all_source_engine_paths|platform`
`    `
`                   }`
`           }`
`   }`

# Serious Sam 3

# Counter Strike

# Counter Strike: Source

# Half Life

## У меня драйвер Catalyst 13.4 и игра не запускается, не выдавая ошибку

Откатитесь до драйвера версии 13.1. С 13.4 вообще полно ошибок, например
Oil Rush выдаёт гораздо меньший FPS, чем с 13.1. Если уже есть драйвер
новее, тогда попробуйте его.

## Русификация

Автор этих слов не сравнивал русификаторы и не знает какой брать лучше.
Я пробовал только один, [доступный по этой
ссылке](http://forum.csmania.ru/viewtopic.php?f=8&t=279).
протестирован на Linux.

## Курсор мыши отображается поверх игры, угол обзора не полный

Причина этой проблемы в какой-то ошибке при использовании Raw Input -
новой технологии, появившейся в X Input 2. Чтобы отключить Raw Input,
откройте в вашем любимом текстовом редакторе файл

` ~/.local/share/Steam/SteamApps/common/Half-Life/valve/config.cfg`

"\~/" означает "домашний каталог пользователя". В этом файле найдите
строчку

` m_rawinput "1"`

и измените 0 на 1.

# Killing Floor

1.  Русифицируем чат в Killing Floor
2.  Заменяем \~ на русский язык.
3.  Ищем папку `System`:
    `<`*`PATH_TO_STEAM`*`>/steamapps/common/killingfloor/System`
    где `<`*`PATH_TO_STEAM`*`>` директория, куда установлен Steam. По
    дефолту в `./local/share/Steam/`.
4.  Открываем в папке `System` следующие файлы (например через gedit,
    kate):
      - `Engine.int`
      - `GUI2K4.int`
      - `KFMod.int`
      - `ROEngine.int`
      - `UnrealGame.int`
      - `XInterface.int`
5.  Делаем замену в следующем порядке:

| До                     | После                  |
| ---------------------- | ---------------------- |
| ROFonts.ROBtsrmVr      | ROFonts_RUS.ROArial   |
| ROFontsTwo.ROArial24DS | ROFonts_Rus.ROArial24 |
| ROFontsTwo.ROArial22DS | ROFonts_Rus.ROArial22 |
| ROFontsTwo.ROArial18DS | ROFonts_Rus.ROArial18 |
| ROFontsTwo.ROArial14DS | ROFonts_Rus.ROArial14 |
| ROFontsTwo.ROArial12DS | ROFonts_Rus.ROArial12 |
| ROFontsTwo.ROArial9DS  | ROFonts_Rus.ROArial9  |
| ROFontsTwo.ROArial7DS  | ROFonts_Rus.ROArial7  |

Если Вы всё сделали правильно, то теперь у вас в чате будет отображаться
русский текст.

Но, если Вы не хотите заморачиваться этими делами, то можете скачать их
отсюда: <http://rusfolder.com/36354275> и скопировать в папку `System`
с заменой.

# SpaceChem

## Не запускается

Ошибка в консоли такая:

` Setting breakpad minidump AppID = 92800`
` Steam_SetMinidumpSteamID:  Caching Steam ID:  76561197994161996 [API loaded no]`
` ERROR: ld.so: object 'gameoverlayrenderer.so' from LD_PRELOAD cannot be preloaded: ignored.`
` ERROR: ld.so: object 'gameoverlayrenderer.so' from LD_PRELOAD cannot be preloaded: ignored.`
` Stacktrace:`
` `
`   at (wrapper managed-to-native) Tao.Sdl.Sdl.__SDL_Init (int) <0xffffffff>`
`   at Tao.Sdl.Sdl.SDL_Init (int) <0x000a3>`
`   at #=qQdKRcUfE46$efAqMsrzhkw==.#=qgLEANBhVfIuIevx4ZYI9dQ== () <0x0005b>`
`   at #=qQdKRcUfE46$efAqMsrzhkw==.#=qtAieMJa3O4$ICce59T5vcg== () <0x000f7>`
`   at #=qT5m9zMMOGgeXuvB5RBp1CjKxG2TMp9VAFSsa0vglCLY=.#=qXYEX0UCMEzgZT5Za8zdXCw== (#=qQdKRcUfE46$efAqMsrzhkw==) <0x00093>`
`   at #=q_kNnXHjBp6OH1879peO3_DruPfChZ2vzI_L2drgdLTw=.#=qC1N1QIcG3BaOsMiqVVYv5w== (string[]) <0x001d3>`
`   at (wrapper runtime-invoke) `<Module>`.runtime_invoke_int_object (object,intptr,intptr,intptr) <0xffffffff>`
`   at (wrapper managed-to-native) System.AppDomain.ExecuteAssembly (System.AppDomain,System.Reflection.Assembly,string[]) <0xffffffff>`
`   at System.AppDomain.ExecuteAssemblyInternal (System.Reflection.Assembly,string[]) <0x0002b>`
`   at System.AppDomain.ExecuteAssembly (string,System.Security.Policy.Evidence,string[]) <0x00027>`
`   at (wrapper remoting-invoke-with-check) System.AppDomain.ExecuteAssembly (string,System.Security.Policy.Evidence,string[]) <0xffffffff>`
`   at (wrapper xdomain-dispatch) System.AppDomain.ExecuteAssembly (object,byte[]&,byte[]&,string,string[]) <0xffffffff>`
`   at (wrapper xdomain-invoke) System.AppDomain.ExecuteAssembly (string,System.Security.Policy.Evidence,string[]) <0xffffffff>`
`   at (wrapper remoting-invoke-with-check) System.AppDomain.ExecuteAssembly (string,System.Security.Policy.Evidence,string[]) <0xffffffff>`
`   at #=q_kNnXHjBp6OH1879peO3_DruPfChZ2vzI_L2drgdLTw=.#=qC1N1QIcG3BaOsMiqVVYv5w== (string[]) <0x000cb>`
`   at (wrapper runtime-invoke) `<Module>`.runtime_invoke_int_object (object,intptr,intptr,intptr) <0xffffffff>`
` `
` Native stacktrace:`
` `
` /home/username/.local/share/steam/SteamApps/common/SpaceChem/spacechem-launcher.sh: line 2: 15906 Ошибка сегментирования                   MONO_CFG_DIR=etc MONO_PATH=monolib ./mono SpaceChem.exe`

Чтобы решить проблему, откройте в любимом файловом менеджере каталог с
игрой, создайте новый каталог и перенесите в него файлы
libSDL-1.2.so.0, libSDL_image-1.2.so.0, libSDL_mixer-1.2.so.0. После
этого игра должна заработать. Теперь игра будет использовать версии
этих библиотек, установленные в дистрибутив Linux.

# Anna Extended edition

Чтобы игра нормально заработала, и Вы смогли совершать какие-либо
действия, нужно зайти в свойства игры. Установить параметры
запуска:

` LC_ALL=C %command%`

Нажать ОК. Удачной игры

# The Book of Unwritten Tales (+ The Book of Unwritten Tales: The Critter Chronicles)

## Игра вылетает при попытке зайти в меню Settings

Скачайте пакет libogre-1.7.4 для архитектуры i386
(http://packages.ubuntu.com/precise/libogre-1.7.4 - с другими версиями
может не сработать). Найдите в архиве файл `RenderSystem_GL.so` и
поместите его в директорию игры `./lib/32/`, заменив файл, который
там находится.

# DOTA 2

## Ключи запуска

Ключи запуска вводятся тут --\>

`клиент steam / библиотека / свойства DOTA 2 / общие / установить параметры запуска`

Пример:

`-language russian -novid -nojoy -console +mat_autoload_glshaders 0 +net_graph 1`

Со знаком "-" ключи вводятся только в свойствах игры. Со знаком "+"
ключи также можно ввести и внутри игры в консоли.

  - `-language russian` --- включить внутри игры русский язык. Текст
    будет по-русски, звуки останутся по-английски.
  - `-novid` --- отключить показ видео-заставки при старте игры.
  - `-nojoy` --- отключить поиск джойстика при старте игры.
  - `-console` --- включить внитриигровую консоль для ввода команд.
  - `+mat_autoload_glshaders 0` --- отключить предварительную компиляцию
    и кэширование шэйдеров при старте игры. Этот параметр очень сильно
    влияет на скорость загрузки клиента игры (startup time). При
    `+mat_autoload_glshaders 1` --\> медленная загрузка `~80 сек`. При
    `+mat_autoload_glshaders 0` --\> быстрая загрузка `~28 сек`.
    Обсуждения по следующим ссылкам на github разработчика:
    <https://github.com/ValveSoftware/Dota-2/issues/661>
    <https://github.com/ValveSoftware/Dota-2/issues/774>
  - `+net_graph 1` --- включить показ внитриигровой статистики
    (техинформация: FPS и т.д.).

Ещё команды:

  - `-override_vpk` --- игра будет искать ресурсы в
    \\steamapps\\common\\dota 2\\dota, заменять ими стандартные и
    загружать их в игру. Используется с модами.
  - `-autoconfig` --- восстанавливает настройки графики по умолчанию.
    Игнорирует любые установленные конфиги до удаления данного
    параметра.
  - `-dev` --- включает режим разработчика. Также отключает
    автоматическую загрузку фоновой картинки меню и
    перестает запрашивать подтверждение о выходе.
  - `-32bit` --- запускает движок в 32-разрядном режиме. Полезно только
    для 64-х разрядных ОС.
  - `-safe` --- запускает игру в окне без рамки с разрешением 640 х 480.
  - `-full или -fullscreen` --- запускает игру в полноэкранном режиме.
  - `-windowed или -sw` --- запускает игру в экранном режиме.
  - ` -w  `<width>`  -h  `<height> --- запускает игру с указанным
    разрешением экрана (ширина и высота). К примеру: -w 1280 -h
    1024. Используется для нестандартных расширений экранов.
  - `-nosound` --- полностью отключает звук в игре. Полезно для обхода
    проблем со звуком.
  - `-nosync` --- отключает вертикальную синхронизацию.
  - `-low` --- запускает игру с пониженным приоритетом.
  - `-high` --- запускает игру с повышенным приоритетом.
  - `+exec "имя cfg файла" (без кавычек)` --- автоматически подгружает
    указанный конфигурационный файл при запуске. Используется для
    загрузки специальных конфигов с настройками.
  - `-noaafonts` --- отключить TrueType шрифты.
  - `-noborder` --- не отображать рамку и заголовок окна при запуске
    игры в оконном режиме.
  - `-freq "ЧИСЛО" или -refresh "ЧИСЛО"` --- устанавливает частоту
    обновления экрана.

## Проблемы

1.  Внутри игры часто разрывается связь с серверами и на экране
    появляется надпись "auto-disconnect".
      - Частичное решение: внутри клиента DOTA 2 зайти в "`Настройки /
        Игра / Сеть / Качество соединения`" и выставить значение
        "низкое".

# The Witcher 2

## Проблемы

1.  При нажатии на Запуск игры ничего не происходит или появляется
    сообщение об ошибке.
      - Решение для Fedora: открываем `/etc/sysconfig/prelink` и
        заменяем `PRELINKING=yes` на `PRELINKING=no` и выполняем
        под рутом `prelink -ua`

# Garry's Mod

## Проблемы

1.  В дистрибутивах amd64 в консоли пишет ошибку **Create Stream Failed
    error 44**
      - Решение:


  - Установить libmpg123-0:i386

Для Ubuntu это выглядит так (файл libmpg123-0:i386 скачать в домашнюю
папку):

  - sudo dpkg --add-architecture i386
  - sudo apt-get update
  - sudo dpkg -i libmpg123-0_1.18.0-1ubuntu1_i386.deb

# Limbo

## Проблемы

1.  Игра не запускается. При запуске из консоли выдаёт подобную ошибку:

``Error in `./limbo': free(): corrupted unsorted chunk``

\#\* Если используете KDE, отключите "Desktop effects" в настройках:
Settings -\> System settings -\> Desktop effects -\> снять галочку
"Enable desktop effects at startup", примените изменения, нажав кнопку
\[Apply\]. Чтобы сработало наверняка, можно разлогиниться и залогиниться
снова.

\#\* Игру можно попробовать запустить в оконном режиме: для этого
откройте файл с настройками игры
"./SteamApps/common/Limbo/settings.txt"; найти и раскомментировать
параметр "windowedmode"; присвоить ему значение "true" и созранить
файл. После этого попытаться снова запустить игру.