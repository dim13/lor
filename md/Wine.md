## Как мне настроить Wine, чтобы все работало?

Никак. **Всё** работать не будет. И об этом написано даже на сайте
проекта.

Если совсем ничего не помогает, обратитесь к [виртуальным машинам или
оболочкам для запуска операционных
систем](Поиск_ПО#Какие_существуют_виртуальные_машины_и_среды_для_запуска_других_операционных_систем_\(в_том_числе_MS_Windows%3F\) "wikilink").

## Какой использовать Wine?

Чтобы успешно запустить то или иное приложение, нужно знать, чем его
запускать.

Приложения MS-DOS (не путать с консольными приложениями под Windows) не
нужно запускать с помощью Wine или его аналогов, они работают крайне
плохо и неустойчиво. Для этого существуют [другие
программы](Поиск_ПО#Какие_существуют_«эмуляторы»_для_запуска_программ_MS-DOS%2C_MS_Windows_1-3.11%3F "wikilink").

### Wine

[Сайт проекта](http://winehq.com)

Wine это реализация Windows API, использующая графический сервер
XFree86/Xorg. Он позволяет запустить программу, написанную для Windows
на любых платформах x86 и x86\_64. Windows для работы не нужна (но ее
dll и ее реестр могут использоваться). Название является акронимом: WINE
— Wine Is Not Emulator.

Wine является основной веткой, на основе которой строятся проекты Cedega
и Crossover.

Эти проекты взаимосвязаны и изменения перетекают из проекта в проект.
Wine в общем случае способен запустить те приложения, которые давно
запускаются в Cedega и Crossover плюс те, которые последние проекты
потеряли возможность запускать в силу своей большей ориентированности на
определенные приложения.

### TransGaming Cedega

[Сайт проекта GameTree Linux](http://gametreelinux.com/), в рамках
которого продолжается разработка Cedega

[Неофициальная
wiki](http://cedegawiki.sweetleafstudios.com/wiki/Main_Page)

Cedega (бывший WineX) является коммерческим ответвлением Wine и
отличается от последнего ориентированностью на запуск игр.
Cedega поддерживает ряд защит (но не поддерживает KP2 или последние
версии Starforce). По сути Cedega — это Wine с костылями для
запуска популярных игр. Позднее, когда становится понятным что и
как работает, поддержка тех или иных механизмов переписывается и
попадает в Wine.

Изначально WineX/Cedega распространялась по платной подписке, однако с
некоторых пор распространяется бесплатно.

### CodeWeavers CrossOver

[Официальный сайт](http://www.codeweavers.com)

CrossOver используется для запуска небольшого числа офисных
Windows-приложений — Microsoft Office, Internet Explorer, Microsoft
Visio, Lotus Notes, Quicken, Adobe Photoshop и т.п. В целом довольно
качественный продукт, но платный.

Способен запускать и другие программы, но раз на раз не приходится.

Существует так же отдельный вариант CodeWeavers CrossOver Games,
ориентированный на запуск игр.

## У меня под Wine/Cedega в игре портится звук — шипит как-то. Что мне делать? Как включить поддержку ALSA?

Посмотреть, не запущен ли PulseAudio. Если да — выключить его и
попробовать еще раз или прочитать ответ ниже. Если это не
помогает — значит поддержка этой игры сделана недостаточно хорошо.
Как вариант, можно попробовать сменить частоту звука, если игра это
позволяет.

Чтобы выбрать, какой звуковой API будет использовать Wine запустите
утилиту winecfg, закладка Audio.

## Можно ли научить Wine выводить звук через PulseAudio?

Можно. Для этого надо выбрать в winecfg вывод звука через OSS, а
запускать wine через padsp так:

`user@linux$ padsp wine /path/to/game.exe `<argument>

Еще можно воспользоваться утилитой pasuspender. Используйте ее для того
что бы не надолго приостановить PulseAudio и запустить wine:

`user@linux$ pasuspender -- wine /path/to/game.exe `<argument>

\==Cedega вылетает с сообщением error=21, просто вылетает с ошибкой. Что
делать?==

Для начала придется разобраться в чем проблема, потому что это слишком
общая ошибка и вылетает она по каждому поводу.

  - У вас вылетает программа установки. Проблема: вы не имеете прав на
    запись на данный раздел или на нем нет места. Также может не быть
    места в разделе /tmp.
  - Игра находится на неисполняемом разделе. Если раздел, с которого вы
    запускаете игру смонтирован с опцией noexec или игра находится на
    FAT32/NTFS разделе, то Cedega вполне может отказаться запускать exe.
  - В вашем дистрибутиве используется защита. Например, патчи grsec,
    любой патч, делающий стек неисполняемым, exec shield, и т.п.
  - У вас установлено ядро 2.6.9. Проблема с VA layout.
  - Используется prelink.
  - Все вышеперечисленное сразу.

Проблему с защитой exec-shield и проблему с VA layout, которые
присутствуют в Fedora Core, можно решить следующим образом:

`echo 0 > /proc/sys/kernel/exec-shield`  
`echo 0 > /proc/sys/kernel/exec-shield-randomize`  
`echo 1 > /proc/sys/vm/legacy_va_layout`  
`/sbin/sysctl -p`

## Как запустить игру ... под Wine/Cedega?

Для начала стоит проверить [каталог программ](http://appdb.winehq.org/)
на наличие истории успеха и каких-либо особых инструкций. Также стоит
заглянуть в раздел [Games](Games "wikilink").

## Как запустить приложение в оконном режиме?

Запускаем winecfg, идем на закладку Graphics, ставим галочку напротив
"Emulate a virtual desktop" и указываем нужное разрешение для оконного
режиме. Игра будет думать что запущена в полноэкранном режиме.

## Как заставить Lingvo работать под Wine?

12ая версия работает под wine "из коробки" -
[1](http://appdb.winehq.org/objectManager.php?sClass=version&iId=11432).
Версия X3, она же 13ая, не работает никак.

## Как поставить Autocad под Wine?

/\* FIXME: Аналогично и с Autocad \*/

## Можно ли запустить 1С под Wine?

Вам [сюда](http://www.etersoft.ru/content/view/56/63/)

На L.O.R периодически всплывает информация об успешном запуске и работе.
Воспользуйтесь [поиском](http://www.linux.org.ru/search.jsp).

Доподлинно известно, что у 1C должен быть сразу отключен splashscreen,
иначе падает тут же.

## Как запускать консольные приложения? Как настроить в них шрифт?

Для этого нужно запускать не wine, а wineconsole с ключом --backend=user
для запуска в X Window или --backend=curses для работы в консоли.
Например, wineconsole --backend=user Far.

Для настройки шрифтов в реестре Wine нужно найти ключи
HKEY\_CURRENT\_USER\\Console\\FaceName и
HKEY\_CURRENT\_USER\\Console\\FontSize и поставить желаемый шрифт
(например, Courier New Bold) и размер (старшее слово - высота,
младшее - ширина).

## Как установить библиотеку DirectX, .Net Framework, Visual C++ Runtime, etc. ?

Для этого можно использовать скрипт
[winetricks](http://wiki.winehq.org/winetricks). Он существенно упрощает
установку библиотек Windows в Wine.

## Как использовать com-порт под Wine?

Для начала, необходимо добавить пользователя в группу, которая имеет
доступ к последовательному порту (dialout для Debian).

Далее, надо сделать символьную ссылку в директории \~/.wine/dosdevices/
на последовательный порт в системе. Например:

`user@linux$ ln -s /dev/ttyUSB0 ~/.wine/dosdevices/com1`

Тем программам, которые работают с com-портом напрямую (?), этого должно
быть достаточно. А для тех, которые пользуются помощью windows для
работы с последовательным портом, надо будет применить файл
[реестра](http://bugs.winehq.org/attachment.cgi?id=10210) со
[страницы](http://bugs.winehq.org/show_bug.cgi?id=9902)
соответствующего баг-репорта. Например, так:

`user@linux$ wine regedit file.reg`

## Как установить одновременно несколько версий Wine?

Да, некоторые программы лучше всего идут на особых версиях Wine. Для
удобного управления множеством версий Wine воспользуйтесь программой
Q4Wine [2](http://q4wine.brezblock.org.ua)

Какая версия нужна для вашей програмы, игры можно посмотреть на сайте
[3](http://appdb.winehq.org)

## Где взять старую версию Wine?

Если в [базе программ](http://appdb.winehq.org) рекомендуют именно
старую версию, то её можно скачать здесь
[4](http://wine.budgetdedicated.com/archive/index.html).

## У меня 64 бита, я не осилил multiarch, хочу OpenGL или нужна свежая версия Wine

Придется собирать то, что нужно из исходников. Оригинальная статья
находится здесь
[5](http://ubuntu-wine.ru/publ/sborka_32_bit_wine_na_64_bit_sisteme_ubuntu_12_04/1-1-0-13),
ниже замечания связанные с Debian/testing а также новыми версиями.

Меняется команда создания chroot:

` $ sudo debootstrap --variant=buildd --arch i386 testing ~/wine_chroot `<http://ftp.de.debian.org/debian>

Исходники лучше брать из Git:

` # git clone `<git://source.winehq.org/git/wine.git>

Чтобы ./configure собрал полноценную версию, установим некоторые
дополнительные пакеты (nvidia-opencl-dev если у нас nVidia):

` # apt-get install libosmesa6-dev nvidia-opencl-dev libgstreamer-plugins-base0.10-dev`

Теперь собираем пакет как в ссылке выше, устанавливаем в 64битную
систему и играемся.

[Category:Wine](Category:Wine "wikilink")
