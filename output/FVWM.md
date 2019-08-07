**FVWM** (также **F(?)VWM**, значение буквы **F** утеряно) - второй в
истории оконный менеджер для иксов (первым был twm), до сих пор
актуальный и поддерживаемый (хоть и не так активно, но в
официальном списке рассылок проскакивают новые патчи),
насчитывающий свыше 250 000 строчек исходного кода.

# FVWM

Обладает большим набором функций, и, если так будет понятнее, FVWM --
это как конструктор или фреймворк, в котором уже реализованы
<u>все</u> необходимые возможности современных оконных менеджеров, а
вам лишь остается написать конфигурационный файл для своих задач. FVWM
позволяет создавать внешний вид (декорации) для каждого окна в
отдельности, или отключить декорации совсем, задавать поведение
каждого окна в отдельности и так далее... Помимо настраиваемых окон, из
коробки имеет 4 294 967 295 виртуальных рабочих столов, столь же
настраиваемое меню и возможность выполнять скрипты (на любом
языке)... И, к примеру, для настройки того же поведения окон, с
помощью своих скриптов можно реализовать тайлинг в FVWM. FVWM дает
вам полную свободу действий, и вы можете сделать свой оконный
менеджер, с нескучными декорациями и панелями, совсем не похожий
на другие. FVWM -- полноценный оконный менеджер, в котором реализованы
все существующие стандарты (EWMH, ICCCM-2 и GNOME Hints). С FVWM
возможен любой "высший пилотаж" в управлении окружением.

**Для кого предназначен FVWM?** Если вы не ленитесь узнать что-то новое,
готовы потратить время на чтение документации, чтобы однажды, но
навсегда настроить что-либо для себя, и чтобы затем это просто
стабильно работало -- FVWM для вас. Потратить день на изучение, чтобы
потом за 5 минут "долететь". Собственно, не ради ли этого вы
используете [Linux](Linux "wikilink")? Либо, даже если вам
просто нужен <u>полноценный</u> оконный менеджер, следующий
стандартам.

**Кому FVWM противопоказан?** Если вы любите нескучные 3D эффекты и
пошлый глянец (хотя и это настраиваемо, опять же), или
предпочитаете, чтобы за вас уже подумали другие, и
предоставили рабочее решение "под ключ", или вам жаль потратить
пару вечеров на чтение документации и настройку софта. Вам надоел
"линукс ради линукса", и надоело всё вечно настраивать и
допиливать -- FVWM не для вас.

## Установка и запуск

FVWM (пакет fvwm) доступен в официальных репозиториях множества
дистрибутивов, в некоторых может быть установлен уже даже из
коробки. Если нет, можно [скачать с официального
сайта](http://fvwm.org/download/) и, распаковав архив с
исходниками, собрать их самостоятельно.

    user@linux:~$ wget ftp://ftp.fvwm.org/pub/fvwm/version-2/fvwm-2.6.5.tar.bz2
    user@linux:~$ tar xf fvwm-2.6.5.tar.bz2
    user@linux:~$ cd fvwm-2.6.5
    user@linux:~$ ./configure && make
    user@linux:~$ su -c "make install"

После чего запустить еще одни иксы по соседству, но уже с FVWM.

    xinit fvwm -- :1

Для запуска FVWM как оконного менеджера по умолчанию вам уже необходимо
подправить ваш \~/.xinitrc или какой другой [Экранный
менеджер](http://ru.wikipedia.org/wiki/XDM), но в этом вы и сами
разберетесь, раз уж осилили взяться за изучение FVWM и собрали его...

Конфигурационный файл расположен в файле **\~/.fvwm/config**, но по
умолчанию сам он не создается, поэтому:

    user@linux:~$ mkdir .fvwm
    user@linux:~$ touch .fvwm/config

## Первое впечатление

Вы получаете чистый экран с малофункциональной менюшкой по клику в любом
месте на рабочем столе. С нуля, если вы не используете
[Debian](Debian "wikilink") и стандартный эмулятор терминала xterm -- вы
вообще вряд ли чего сможете запустить в fvwm. Чтобы такого не произошло,
запишем первую строку в конфигурационный файл, чтобы по нажатию
определенных клавиш запускался терминал, откуда затем мы уже
сможем запустить что угодно.

Добавим в \~/.fvwm/config:

    Key Return A CM Exec urxvt

В данном случае: (**Key**) по нажатию клавиши (**Return**) Enter (**A**)
в любом контексте, будь то рабочий стол или активное окно (**CM**) с
зажатыми клавишами Ctrl и Meta (Alt) следует запустить программу
(**urxvt**).

То есть, <s>пере</s>запускаем FVWM с одной такой строкой в конфиге, и
теперь по нажатию *Ctrl + Alt + Enter* мы запустим нужный нам
эмулятор терминала.

Глядя на декорации окон, кто-то пустит скупую слезу по былым временам и
предастся ностальгии, а кто-то в ужасе закроет FVWM и больше никогда
его не запустит. На время нашего первого знакомства с FVWM предлагаю
отключить декорации окон совсем, чтобы они вас не отвлекали.

Добавим в \~/.fvwm/config:

    Style * !Title, !Handles, !Borders

В данном случае: (**Style**) командой по управлению внешним видом и
поведением (**\***) всем окнам (**\!**) отключим (**Title**)
заголовок, (**Handles**) элементы управления по бокам окна
(**Borders**) и однопиксельную рамку.

Получилось минималистично, но хотя бы не так ужасно. И просто для того,
чтобы данным оконным менеджером можно было пользоваться, соберем
минимально необходимую, но удобную для того конфигурацию.

Добавим в \~/.fvwm/config:

    DesktopSize 3x1
    EdgeThickness 0

    OpaqueMoveSize 0
    MoveThreshold 0

    DestroyFunc FuncFvwmRaiseLowerX
    AddToFunc FuncFvwmRaiseLowerX
    + I Raise
    + M $0
    + D Lower

    Key Return A CM Exec urxvt
    Key Delete W CM Close
    Key Insert W CM RaiseLower
    Key BackSpace A CM Restart

    Key 1 A CM GotoPage 0 0
    Key 2 A CM GotoPage 1 0
    Key 3 A CM GotoPage 2 0

    Key x W CM Maximize

    Mouse 1 R A Menu MenuFvwmWindowOps
    Mouse 2 R A WindowList NoDeskSort
    Mouse 3 R A Menu MenuFvwmRoot

    Mouse 1 T A FuncFvwmRaiseLowerX "Move"
    Mouse 1 FS A FuncFvwmRaiseLowerX "Resize"
    Mouse 1 W M FuncFvwmRaiseLowerX "Move"
    Mouse 2 W M Lower
    Mouse 2 FST A Lower
    Mouse 3 WFS M FuncFvwmRaiseLowerX "Resize"

    Style * !Title, !Handles, !Borders
    Style URxvt PositionPlacement UnderMouse

Горячие клавиши:

  - *Ctrl + Alt + Enter* -- запустит терминал urxvt (измените название
    терминала, если необходимо);
  - *Ctrl + Alt + Delete* -- закроет активное окно;
  - *Ctrl + Alt + Insert* -- переключает фокус между окнами и выводит
    следующее окно на передний план;
  - *Ctrl + Alt + BackSpace* -- перезапустит FVWM (необходимо при
    редактировании конфигурационного файла);
  - *Ctrl + Alt + X* -- развернет окно на весь экран, при повторном
    нажатии восстановит его обратно;
  - *Ctrl + Alt + 1, 2, 3* -- переключает рабочие столы (1, 2, или 3
    соответственно);
  - *(Shift +) Alt + Tab* -- встроенные клавиши для переключения между
    окнами.

В случае, если вы решили не отключать декорации окон -- мышкой вы можете
управлять окнами, таская их за декорации, иначе, зажав клавишу Meta
(Alt), левой кнопкой мышки вы перемещаете окно, средней (колесиком, либо
двойным кликом левой) переключаете фокус между окнами, а правой
изменяете размер окна.

И таким образом, на первых порах, мы получили классический оконный
менеджер, который предоставляет весь необходимый минимум в
управлении окнами.

## Правила оформления кода

При оформлении конфигурационного файла вовсе не обязательно следовать
данным правилам, но так уже сложилось...

  - Названия функций следует начинать с **FuncFvwm**. Например,
    **FuncFvwmMaximize**, если вы желаете переопределить функцию
    Maximize, и затем поменять везде название данной функции на вызов
    вашей новой FuncFvwmMaximize. Исключение составляют стандартные
    встроенные функции -- как раз потому, что они встроенные, и чтобы
    ваши функции с ними не пересекались, добавляйте вышеуказанный
    префикс.
  - Названия меню следует начинать с **MenuFvwm**. **MenuFvwmRoot** --
    самое главное (корневое) меню, которое вызывается в FVWM по клику
    на рабочем столе, **MenuFvwmWindowOps** -- меню со списком действий
    над окнами, **StartMenu** используется в модуле FvwmTaskBar для
    вызова меню кнопки "Пуск". Так или иначе, названия меню не
    критичны, ведь вы сами определяете, какие и где меню вам
    вызывать.

# Особенности поведения

У каждого оконного менеджера имеются свои особенности поведения (как
задумали их авторы), и в большинстве они вшиты в исходный код. Меняя
оконные менеджеры, вам приходится привыкать к новому поведению окон и
прочим мелочам, которые могут казаться непривычными или даже
раздражать, но в FVWM все настраиваемо. FVWM не стал
исключением и по умолчанию имеет свой ряд особенностей.

## Инициализация

Имеется ряд встроенных функций для последующего автовыполнения команд,
таких, как **InitFunction**, **SessionInitFunction**,
**StartFunction**, **SessionStartFunction**, **RestartFunction**,
**SessionRestartFunction** и **ExitFunction**.

Однако двух функций с последующей проверкой условий вполне достаточно.

    DestroyFunc StartFunction
    AddToFunc StartFunction
    + I Test (Init) PipeRead 'test -d "/tmp/fvwm-$[UID]" || mkdir "/tmp/fvwm-$[UID]"; echo Nop'
    + I Test (Init) PipeRead 'xsetroot -solid "#400040"; echo Nop'

    DestroyFunc ExitFunction
    AddToFunc ExitFunction
    + I Test (Quit) PipeRead 'rm --force --recursive "/tmp/fvwm-$[UID]"; echo Nop'

Здесь командой **Test** мы проверяем, что сейчас конкретно происходит, и
при запуске, в зависимости от выполненного условия, создадим временную
директорию для FVWM (необязательно, просто она нам еще пригодится),
установим обои (фон) на рабочий стол. При закрытии FVWM удалим
временную директорию со всем содержимым.

## Alt + Tab, Alt + Shift + Tab

Как и во многих других оконных менеджерах, данные горячие клавиши
переключают окна, по умолчанию вызывая меню со списком окон. Не
всех может устраивать такое поведение, поэтому для "тихого"
переключения между окнами предлагаю переопределить их на свои
функции.

    DestroyFunc FuncFvwmNextWindow
    AddToFunc FuncFvwmNextWindow
    + I Next (CurrentPage, !Iconic, AcceptsFocus) FlipFocus NoWarp
    + I Current Raise

    DestroyFunc FuncFvwmPrevWindow
    AddToFunc FuncFvwmPrevWindow
    + I Prev (CurrentPage, !Iconic, AcceptsFocus) Focus NoWarp
    + I Current Raise

    Key Tab A M FuncFvwmNextWindow
    Key Tab A SM FuncFvwmPrevWindow

Таким образом, курсор мышки не будет скакать от окна к окну, а фокус
между окнами переключается без всяких меню.

## Автофокус окна

Представим, что вы используете все виртуальные рабочие столы в
количестве 4 294 967 295 штук, на каждом из которых крутятся
какие-то приложения, и вот одно из них "позвало вас к себе" -- когда,
например, вас хайлайтнули (обратились к вашему нику) в чате. Среди
такого зоопарка окон трудно найти нужное, а если это был просто
"бип" спикером? Сиди, гадай, что за окно может бибикать и по какой
причине.

В FVWM было решено мгновенно переключаться на такое окно, и это удобно,
но может раздражать, если вы в это время писали код, а какое-то
надоедливое окно вас отвлекает автофокусом на него.

Так вот, когда такое окно "звенит", вызывается встроенная функция
**UrgencyFunc**, и когда оно перестает "звенеть", вызывается
**UrgencyDoneFunc**. По умолчанию функции выглядят следующим образом:

    AddToFunc UrgencyFunc
    + I Iconify off
    + I FlipFocus NoWarp
    + I Raise

    AddToFunc UrgencyDoneFunc
    + I Nop

И понятное дело, вы можете их переопределить. К примеру, наделать
всплывающих уведомлений где-либо с наименованием окна, которое
вас "позвало"... Все же удобнее и проще будет просто их удалить, и
более "хайлайт в чатах" вас не будет отвлекать (в прямом смысле).

    DestroyFunc UrgencyFunc

    DestroyFunc UrgencyDoneFunc

Еще один момент, когда, помимо звенящих окон, стандарт EWMH требует
активации окна, и схож он с предыдущей функцией, но называется
EWMHActivateWindowFunc:

    AddToFunc EWMHActivateWindowFunc
    + I Iconify off
    + I Focus NoWarp
    + I Raise

Ее мы также удаляем, если вам совсем-совсем не хочется отвлекаться на
требующие к себе внимания окна.

    DestroyFunc EWMHActivateWindowFunc

## Меню

Есть два вида меню -- обычное (скажем, для запуска программ или
выполнения каких-то действий) по умолчанию называется
**MenuFvwmRoot**, и меню со списком всех окон, которое вызывается
функцией **WindowList**.

Если по заголовку меню кликнуть средней кнопкой мыши (колесиком), то
меню превратится в обычное окно **fvwm_menu**, и для него также
можно установить отдельные стили:

    Style fvwm_menu Title, Handles, Borders
    Style fvwm_menu UsePPosition

### WindowList

**WindowList** вызывается по клику мышкой на рабочем столе, либо
встроенными горячими клавишами *Alt + Tab* (чтобы переключить
следующее окно) и/или *Alt + Shift + Tab* (для переключения на
предыдущее окно).

Когда вы выбираете пункт в этом меню, вызывается функция
**WindowListFunc**, по умолчанию выполняющая следующее:

    AddToFunc WindowListFunc
    + I Iconify off
    + I FlipFocus NoWarp
    + I Raise

Можете переопределить ее по своему усмотрению.

## Иконки приложений

Обычно графические приложения уже имеют свои иконки, установленные в
/usr/share/icons/hicolor, но в большинстве случаев вам придется задавать
иконки приложениям вручную, если вы хотите использовать какие-то "темы"
иконок.

Для удобства задаем **ImagePath**, где FVWM будет искать картинки.

    ImagePath /usr/share/icons/Tango:+

И теперь к каждому приложению задаем иконки.

    Style * Icon "48x48/mimetypes/application-x-executable.png", MiniIcon "22x22/mimetypes/application-x-executable.png"
    Style URxvt Icon "48x48/apps/utilities-terminal.png", MiniIcon "22x22/apps/utilities-terminal.png"
    Style Firefox Icon "48x48/apps/internet-web-browser.png", MiniIcon "22x22/apps/internet-web-browser.png", IconOverride, EWMHMiniIconOverride
    Style MPlayer Icon "48x48/apps/mplayer.png", MiniIcon "22x22/apps/mplayer.png"
    Style Thunar Icon "48x48/apps/system-file-manager.png", MiniIcon "22x22/apps/system-file-manager.png"
    Style Leafpad Icon "48x48/apps/accessories-text-editor.png", MiniIcon "22x22/apps/accessories-text-editor.png", IconOverride, EWMHMiniIconOverride
    Style Deadbeef Icon "48x48/apps/audioplayer.png", MiniIcon "22x22/apps/audioplayer.png", IconOverride, EWMHMiniIconOverride

Обратите внимание, что некоторые графические приложения требуют
дополнительные параметры **IconOverride** и
**EWMHMiniIconOverride**, чтобы установить иконки.

## Свернутые окна

Сворачиваемые окна становятся иконками на рабочем столе, по клику на
которые разворачиваются обратно в окно. Если вы хотите отключить
такие иконки, то тогда установите следующий параметр:

    Style * NoIcon

При этом прежде у окон не должны быть заданы какие-либо иконки (Icon,
MiniIcon), и теперь свернутые окна не будут видны на рабочем столе, а
доступны лишь через какой-либо список окон (всевозможные TaskBar'ы,
либо меню WindowList).

# Советы и трюки

## Выполнить...

Если у вас установлен эмулятор терминала rxvt-unicode (urxvt), то мы
можем использовать его отдельно для запуска приложений.

    DestroyFunc FuncFvwmRun
    AddToFunc FuncFvwmRun
    <nowiki>+ I PipeRead 'test -a "/tmp/fvwm-$[UID]/!PS1" || echo "PS1=" > /tmp/fvwm-$[UID]/!PS1; urxvtc -geometry 50x1+30-30 \
                       -name fvwm_run --keysym.0xFF0D " &\\nexit\\n" -fn "9x15bold" +sb -e bash --init-file "/tmp/fvwm-$[UID]/!PS1"'</nowiki>

    Key r A CM FuncFvwmRun

    Style fvwm_run !Title, !Handles, Borders
    Style fvwm_run !Iconifiable, !Maximizable, FixedPosition, FixedSize
    Style fvwm_run GrabFocus

По нажатию *Ctrl + Alt + R* в левом нижнем углу у вас появится urxvt,
если что-нибудь ввести и нажать Enter -- оно запустится, а urxvt
закроется.

## Фокус окна в зависимости от положения курсора

**FocusFollowsMouse** -- фокус переключается на окно, над которым
расположен указатель мыши.

**MouseFocusClickRaises** -- Click на любой части окна поднимает его
наверх.

    Style * FocusFollowsMouse, MouseFocusClickRaises

Автоматическое всплытие окна через 1 секунду, при наведении указателя.

    Module FvwmAuto 1000 Raise Nop

## Alt-tab, unix way

Стандартное переключение окон через список -- это, конечно, удобно, но
порой хочется чего-то необычного и юниксового. Пример ниже работает
так: сначала показываем предыдущее окно, после -- переключаемся между
окнами по кругу. Звучит удобно, не так ли?

    InfoStoreAdd DIR Next

    DestroyFunc FuncSwitchWindow
    AddToFunc FuncSwitchWindow
    + I $[infostore.DIR] (CurrentDesk, !Iconic, !Sticky) SelectWindow
    + I Deschedule 134000
    + I PipeRead `[ "$[infostore.DIR]" == "Prev" ] && \
                  echo 'InfoStoreAdd NDIR Next' || \
                  echo 'InfoStoreAdd NDIR Prev'`
    + I Schedule 700 134000 InfoStoreAdd DIR $[infostore.NDIR]

    DestroyFunc SelectWindow
    AddToFunc   SelectWindow
        + I Iconify false
        + I Focus
        + I Raise
        + I WarpToWindow 50 50

    Key Tab A 4  FuncSwitchWindow

Обратите внимание на строчку "+ I WarpToWindow 50 50" -- она перемещает
курсор в центр окна. Если у вас фокус окна выделяется кликом по нему,
то строчку можно убрать.

Я не стал портить картину *Alt + Tab* и прописал сочетание *Win + Tab*
для переключения вышеописанным способом.

## Операции над окном (или окнами)

### Полноэкранный режим

Может быть удобным при развертывании окна на весь экран отключать его
декорации и фиксировать (ни свернуть, ни развернуть, ни поменять
размер или положение окна), и тем же методом восстанавливать окно в
обычное состояние.

    DestroyFunc FuncFvwmMaximize
    AddToFunc FuncFvwmMaximize
    + I ThisWindow (Maximized) DestroyWindowStyle
    + I TestRc (NoMatch) ThisWindow (!Maximized) WindowStyle !Title, !Handles, !Borders
    + I UpdateStyles
    + I ThisWindow (Maximized) Maximize
    + I TestRc (NoMatch) Maximize ewmhiwa 100 100
    + I TestRc (NoMatch) WindowStyle !Iconifiable, !Maximizable, FixedPosition, FixedSize
    + I TestRc (NoMatch) UpdateStyles

    Key x A CM FuncFvwmMaximize

Делаем это клавишами *Ctrl + Alt + X*.

### Свернуть/развернуть все окна

#### Свернуть/развернуть все окна

Можно свернуть все окна на рабочем столе, как это сделано в windows.

    DestroyFunc FuncSHDesktop
    AddToFunc FuncSHDesktop
        + I All (CurrentPage, Iconifiable, !Iconic, !WindowListSkip, !Shaded) Iconify on
        + I TestRc (Match) Break
        + I All (CurrentPage, Iconic, !WindowListSkip) Iconify off

    Key D A 4 FuncSHDesktop
    Mouse 1 R M FuncSHDesktop

Таким образом, всё, что у вас на текущем экране, будет свернуто в иконки
(сочетание *Win + D*), при повторном нажатии -- развернуто. А сочетание
*Alt + клик_левой_кнопкой* по рабочему столу выполняет ту же функцию.

#### Свернуть/развернуть все окна, но не разворачивать свернутые ранее

Бывает полезно свернуть то, что у вас открыто, но не разворачивать то,
что скрыли до этого. В трюке ниже мы будем использовать параметр
**State**.

    DestroyFunc ShowDesktop
    AddToFunc   ShowDesktop
        + I All (CurrentPage, Iconic, State 31) RestoreDesktop
        + I TestRc (Match) Break
        + I All (CurrentPage, !Iconic, !State 31) ThisWindow State 31 True
        + I All (CurrentPage, !Iconic) Iconify

    DestroyFunc RestoreDesktop
    AddToFunc   RestoreDesktop
        + I All (CurrentPage, Iconic, State 31) Iconify off
        + I All (CurrentPage, State 31) ThisWindow State 31 False

    Key H A 4 ShowDesktop

Таким образом, всё, что у вас на текущем экране, будет свернуто в иконки
(сочетание *Win + H*), при повторном нажатии -- развернуто, кроме тех
окон, что ранее были свернуты.

### AutoShade или свернуть в заголовок

Набор функций, реализующих autoshade, подсмотренный в kwin и стащенный у
Hoodoo. Окно сворачивается в заголовок. Как только курсор оказывается в
заголовке, окно разворачивается. Как только курсор покидает окно, оно
снова сворачивается. Используется команда **State**, позволяющая
задать для любого окна один или несколько статусов, которые
впоследствии можно проверить, и на основании проверки выполнить
ту или иную команду. Модуль **FvwmAuto** следит за курсором и
реагирует, когда тот входит в окно или покидает его. Всего
можно задать 32 статуса (0-31). Вообще этот модуль следовало бы
поместить в автостарт.

    Module FvwmAuto 0 -menter "Silent AutoUnshade" "Silent AutoShade"

    DestroyFunc ToggleAutoshaded
    AddToFunc ToggleAutoshaded
        + I State 0
        + I ThisWindow (!State 0) WindowShade false

    DestroyFunc AutoShade
    AddToFunc AutoShade
        + I ThisWindow (State 0) WindowShade True

    DestroyFunc AutoUnshade
    AddToFunc AutoUnshade
        + I ThisWindow (State 0) WindowShade False

    DestroyFunc FuncWindowShade
    AddToFunc FuncWindowShade
        + I ThisWindow WindowShade toggle
        + I ThisWindow (State 0) State 0
        + I ThisWindow WarpToWindow 50 50

    Mouse 3 T N ToggleAutoshaded

Функция в данном примере активируется по щелчку средней кнопкой мыши по
заголовку окна.

### Группировка окон

    DestroyFunc AddToGroup
    AddToFunc AddToGroup
        + I ThisWindow State $0

    DestroyFunc GroupSwitch
    AddToFunc GroupSwitch
        + I Next (State $0) SelectWindow

    DestroyFunc SelectWindow
    AddToFunc   SelectWindow
        + I Iconify false
        + I Focus
        + I Raise
        + I WarpToWindow 50 50

```
    Key 1 WTSF C4 AddToGroup 1
    Key 2 WTSF C4 AddToGroup 2
    Key 3 WTSF C4 AddToGroup 3
    Key 4 WTSF C4 AddToGroup 4
    Key 5 WTSF C4 AddToGroup 5
    Key 6 WTSF C4 AddToGroup 6
    Key 7 WTSF C4 AddToGroup 7
    Key 8 WTSF C4 AddToGroup 8
    Key 9 WTSF C4 AddToGroup 9
    Key 1 A 4 GroupSwitch 1
    Key 2 A 4 GroupSwitch 2
    Key 3 A 4 GroupSwitch 3
    Key 4 A 4 GroupSwitch 4
    Key 5 A 4 GroupSwitch 5
    Key 6 A 4 GroupSwitch 6
    Key 7 A 4 GroupSwitch 7
    Key 8 A 4 GroupSwitch 8
    Key 9 A 4 GroupSwitch 9
```

При нажатии *Ctrl + {цифра от 1 до 9}* -- окно добавляется в группу, при
нажатии *Win + {цифра от 1 до 9}* можно переключаться между окнами
данной группы. Повторное нажатие *Ctrl + {цифра от 1 до 9}*
исключает окно из группы.

### Окно поверх других окон

    DestroyFunc WinOnTop
    AddToFunc WinOnTop
        + I ThisWindow (Layer 6) Layer
        + I TestRc (NoMatch) Layer 0 6
        + I Raise

    Key T          WFST  4   WinOnTop

По сочетанию клавиш *Win + T* окно будет поднято над другими окнами, по
повторному нажатию -- окно примет прежние параметры.

### Прижатие окна к углам/сторонам экрана

    # Для прижатия окна к углам/сторонам страницы и перемещения в центр нажимаем win+
    # +---------+
    # ! q  w  e !
    # !         !
    # ! a  s  d !
    # !         !
    # ! z  x  c !
    # +---------+

    DestroyFunc NWWindow
    AddToFunc   NWWindow
        + I ThisWindow AnimatedMove 0 0
        + I WarpToWindow 50 50

    DestroyFunc NEWindow
    AddToFunc   NEWindow
        + I ThisWindow AnimatedMove -0 0
        + I WarpToWindow 50 50

    DestroyFunc SWWindow
    AddToFunc   SWWindow
        + I ThisWindow AnimatedMove 0 -0
        + I WarpToWindow 50 50

    DestroyFunc SEWindow
    AddToFunc   SEWindow
        + I ThisWindow AnimatedMove -0 -0
        + I WarpToWindow 50 50

    DestroyFunc CWindow
    AddToFunc   CWindow
        + I ThisWindow Piperead "echo AnimatedMove $(( $[vp.width]/2-$[w.width]/2 ))p $(( $[vp.height]/2-$[w.height]/2 ))p"
        + I WarpToWindow 50 50

    # Прижатие окна к середине верхней границы страницы
    DestroyFunc CNWindow
    AddToFunc   CNWindow
        + I ThisWindow Piperead "echo AnimatedMove $(( $[vp.width]/2-$[w.width]/2 ))p 0"
        + I WarpToWindow 50 50

    # Прижатие окна к середине нижней границы страницы
    DestroyFunc SCWindow
    AddToFunc   SCWindow
        + I ThisWindow Piperead "echo AnimatedMove $(( $[vp.width]/2-$[w.width]/2 ))p -0"
        + I WarpToWindow 50 50

    # Прижатие окна к середине левой границы страницы
    DestroyFunc CWWindow
    AddToFunc   CWWindow
        + I ThisWindow Piperead "echo AnimatedMove 0 $(( $[vp.height]/2-$[w.height]/2 ))p"
        + I WarpToWindow 50 50

    # Прижатие окна к середине правой границы страницы
    DestroyFunc CEWindow
    AddToFunc   CEWindow
        + I ThisWindow Piperead "echo AnimatedMove -0 $(( $[vp.height]/2-$[w.height]/2 ))p"
        + I WarpToWindow 50 50

        Key Q           WTSF    4      NWWindow
        Key W           WTSF    4      CNWindow
        Key E           WTSF    4      NEWindow
        Key A           WTSF    4      CWWindow
        Key S           WTSF    4      CWindow
        Key D           WTSF    4      CEWindow
        Key Z           WTSF    4      SWWindow
        Key C           WTSF    4      SEWindow
        Key X           WTSF    4      SCWindow

Вместо AnimatedMove вы можете использовать просто Move. В каждой функции
используется "+ I WarpToWindow 50 50", это перемещает указатель мышки в
центр окна. Если вам это не нужно -- можно смело убрать.

### Миниатюры окон

#### Миниатюры окон(простой вариант)

По умолчанию окна сворачиваются в небольшие иконки приложений на рабочем
столе. С помощью пакетов imagemagick и xorg-xwd возможно сделать так,
чтобы предварительно делался скриншот сворачиваемого окна, а затем
задавался как иконка.

    DestroyFunc FuncFvwmThumbnail
    AddToFunc FuncFvwmThumbnail
    + I Raise
    + I InfostoreAdd Icon-$[w.id] $[w.iconfile]
    + I PipeRead 'xwd -silent -id $[w.id] | convert -scale 128 -quality 0 xwd:- png:- | composite -gravity SouthEast $[w.iconfile] - /tmp/fvwm-$[UID]/Icon-$[w.id].png && echo WindowStyle Icon "/tmp/fvwm-$[UID]/Icon-$[w.id].png" || echo Nop'

    DestroyFunc FuncFvwmDeThumbnail
    AddToFunc FuncFvwmDeThumbnail
    + I PipeRead 'echo WindowStyle Icon \\$\\[Icon-$[w.id]\\]'
    + I InfostoreRemove Icon-$[w.id]
    + I PipeRead 'rm --force "/tmp/fvwm-$[UID]/Icon-$[w.id].png"'

    DestroyFunc FuncFvwmIconify
    AddToFunc FuncFvwmIconify
    + I ThisWindow (!Shaded, Iconifiable, !Iconic) FuncFvwmThumbnail
    + I ThisWindow (Iconic) FuncFvwmDeThumbnail
    + I Iconify
    + I All (CurrentPage, Iconic) PlaceAgain Icon

    Key z W CM FuncFvwmIconify

    Mouse 1 I M FuncFvwmRaiseLowerX "Move"
    Mouse 1 I A FuncFvwmIconify
    Mouse 2 I M Lower
    Mouse 3 I A Menu MenuFvwmWindowOps

    Style * IconFont "Shadow=0 0 BottomRight:StringEncoding=UTF-8:xft:Sans:Medium:Roman:size=8:minspace=False;fixed"
    Style * IconBox screen w 0 0 -0 -0, IconFill left bottom

По нажатию *Ctrl + Alt + Z* окно свернется в небольшую миниатюру в левом
нижнем углу, а по клику по ней развернется обратно.

#### Миниатюры окон(расширенный режим)

Мы делаем миниатюру окна и накладываем на неё иконку приложения (если
она прописана в стилях). Задействовал параметр **State 28**, чтобы не
давать слишком часто сворачивать окна, иначе иконка стиля замещается
миниатюрой окна (в таком случае помогает только DestroyWindowStyle).
Сделана соместимость с функцией **ShowDesktop** (которая использует
**State 31**).

    InfoStoreAdd fvwm_icon_size 160
    InfoStoreAdd fvwm_tmp "/tmp"

    DestroyFunc Thumbnail
    AddToFunc Thumbnail
    + I ThisWindow (State 28) Break
    + I State 28 True
    <nowiki>
        + I Raise
        + I InfostoreAdd Icon-$[w.id] $[w.iconfile]
        + I ThisWindow (!Shaded, Iconifiable, !Iconic) PipeRead "xwd -silent -id $[w.id] | \
                                                                  convert -scale $[infostore.fvwm_icon_size] -frame 1x1 \
        -mattecolor black -quality 0 xwd:- png:$[infostore.fvwm_tmp]/window-$[w.id].png && echo Nop"
        + I TestRc (Match) Test (f $[w.iconfile], f $[infostore.fvwm_tmp]/window-$[w.id].png) PipeRead \
            "composite -gravity SouthEast -geometry +4+6 \
        $[w.iconfile] $[infostore.fvwm_tmp]/window-$[w.id].png $[infostore.fvwm_tmp]/window-$[w.id].png || echo Nop"
        + I TestRc (Match) Test (f $[infostore.fvwm_tmp]/window-$[w.id].png) \
            Piperead "echo WindowStyle IconOverride, Icon $[infostore.fvwm_tmp]/window-$[w.id].png || echo Nop"
    </nowiki>
        + I Schedule 150 Iconify on
    + I Schedule 1500 State 28 False

    DestroyFunc DeThumbnail
    AddToFunc DeThumbnail
        + I Exec rm -f $[infostore.fvwm_tmp]/window-$[w.id].png
        + I PipeRead "echo Test \\(i \\$\\[infostore.Icon-$[w.id]\\]\\) WindowStyle Icon \\$\\[infostore.Icon-$[w.id]\\]"
        + I TestRc (NoMatch) WindowStyle NoIconOverride, Icon
        + I Schedule 100 InfostoreRemove Icon-$[w.id]
        + I ThisWindow (State 31) State 31 False

    # FvwmEvent
    DestroyModuleConfig FvwmEvent: *
    *FvwmEvent: deiconify DeThumbnail

    Module FvwmEvent

    Key z W CM Thumbnail

Сворачиваем окно так же, как и в предыдущем варианте, по *Ctrl + Alt +
z*.

Обратите внимание на модуль **FvwmEvent**, в данном случае все окна,
которые будут разворачиваться из иконки, почувствуют на себе эффект
**DeThumbnail**. :)

### Прозрачность окна

У вас должен быть запущен композитный менеджер, например xcompmgr.

    AddToFunc InitFunction
        + I Exec exec $(sleep 3 && xcompmgr -cC &> /dev/null)&

В примерах используется программа transset-df, которая, в отличие от
transset, считается более стабильной.

#### Устанавливаем прозрачность для текущего окна

А дальше установим сочетания для управления прозрачностью с помощью
transset-df:

    Key Prior WTSF 4 Exec exec transset-df --min 0.3 --inc 0.1 -i $[w.id] &> /dev/null
    Key Next  WTSF 4 Exec exec transset-df --min 0.3 --dec 0.1 -i $[w.id] &> /dev/null
    Key Home  WTSF 4 Exec exec transset-df --min 0.3 -i $[w.id] 1 &> /dev/null
    Key End   WTSF 4 Exec exec transset-df --min 0.3 -i $[w.id] 0 &> /dev/null

Минимальная прозрачность задана как 0.3.

Итак, сочетания клавиш:

  - *Win + PageUp* -- уменьшить прозрачность с шагом 0.1;
  - *Win + PageDown* -- увеличить прозрачность с шагом 0.1;
  - *Win + Home* -- сделать окно полностью непрозрачным;
  - *Win + End* -- сделать окно полностью прозрачным<sup>(с оговоркой,
    что полная прозрачность 0.3)</sup>.

#### Прозрачное окно при перемещении или изменении размера

Легко реализовать прозрачные окна, когда мы их перемещаем, либо изменяем
им размер. Никакой пользы данный эффект не несет, просто рюшечка. :)

    OpaqueMoveSize unlimited
    Style * ResizeOpaque

    DestroyFunc FuncFvwmRaiseLowerX
    AddToFunc FuncFvwmRaiseLowerX
    + I Raise
    + M PipeRead 'transset-df --min 0.3 -i $[w.id] 0 > /dev/null; echo Nop'
    + M $0
    + M PipeRead 'transset-df --min 0.3 -i $[w.id] 1 > /dev/null; echo Nop'
    + D Lower

    Mouse 1 T A FuncFvwmRaiseLowerX "Move"
    Mouse 1 FS A FuncFvwmRaiseLowerX "Resize"
    Mouse 1 W M FuncFvwmRaiseLowerX "Move"
    Mouse 2 W M Lower
    Mouse 2 FST A Lower
    Mouse 3 WFS M FuncFvwmRaiseLowerX "Resize"

Перво-наперво мы установили отображение содержимого окон при их
перемещении, а также при изменении размера (по умолчанию
рисовалась рамка в стиле FVWM). Затем к функции
FuncFvwmRaiseLowerX добавили две команды -- одна выполняется до
перемещения/ресайза, а вторая после. Теперь попробуйте
перемещать окна или изменять их размер -- в этот момент они
будут прозрачными\!

#### Прозрачность для окна по событию

К разделу рюшечек. Пусть главное окно имеет прозрачность 1, а все
второстепенные 0.92. Тогда:

    Module FvwmEvent

    DestroyModuleConfig FvwmEvent: *
    *FvwmEvent: add_window SetFocusTrans
    *FvwmEvent: leave_window SetUnfocusTrans
    *FvwmEvent: enter_window SetFocusTrans

    DestroyFunc SetFocusTrans
    AddToFunc SetFocusTrans
    + I Exec exec transset-df --inc .5 -x 1 -i $[w.id]

    DestroyFunc SetUnfocusTrans
    AddToFunc SetUnfocusTrans
    + I Exec exec transset-df --dec .5 -m .92 -i $[w.id]

Немного поясню по событиям:

  - add_window -- окно создано
  - leave_window -- c окна сняли фокус
  - enter_window -- окно получило фокус

## Управление плеером

Простой способ управления одним из плееров: MPD + mpc, cmus, xmms,
audacious и другими.

### MPD + mpc

MPD сервер должен быть запущен.

    Key XF86AudioNext A A Exec exec $(mpc next &> /dev/null)
    Key XF86AudioPrev A A Exec exec $(mpc prev &> /dev/null)
    Key XF86AudioPlay A A Exec exec $(mpc toggle &> /dev/null)
    Key XF86AudioStop A A Exec exec $(mpc stop &> /dev/null)

### cmus

    # запуск cmus или проигрывание/пауза
    Key XF86AudioPlay A N  Exec exec sh -c "if ( grep playing ~/np.txt >/dev/null ); then cmus-remote -u; else cmus-remote -p; fi"
    # следующий трек
    Key XF86AudioNext A N  Exec exec cmus-remote -n
    # предыдущий трек
    Key XF86AudioPrev A N  Exec exec cmus-remote -r
    # останов
    Key XF86AudioStop A N  Exec exec cmus-remote -s

### xmms

Управление xmms

    Key XF86AudioStop A N  Exec exec xmms -s
    Key XF86AudioPlay A N  Exec exec xmms -t
    Key XF86AudioNext A N  Exec exec xmms -f
    Key XF86AudioPrev A N  Exec exec xmms -r

### audacious

    Key XF86AudioPrev A N  Exec exec audacious --rew
    Key XF86AudioStop A N  Exec exec audacious --stop
    Key XF86AudioPlay A N  Exec exec audacious --play-pause
    Key XF86AudioNext A N  Exec exec audacious --fwd

### Quod Libet

    Key XF86AudioPrev A N  Exec exec quodlibet --previous
    Key XF86AudioStop A N  Exec exec quodlibet --stop
    Key XF86AudioPlay A N  Exec exec quodlibet --play-pause
    Key XF86AudioNext A N  Exec exec quodlibet --next

### banshee

    Key XF86AudioPrev A N  Exec exec banshee --previous
    Key XF86AudioStop A N  Exec exec banshee --stop
    Key XF86AudioPlay A N  Exec exec banshee --play-pause
    Key XF86AudioNext A N  Exec exec banshee --next

### Goggles Music Manager

    Key XF86AudioPrev A N  Exec exec gmm --previous
    Key XF86AudioStop A N  Exec exec gmm --stop
    Key XF86AudioPlay A N  Exec exec gmm --play-pause
    Key XF86AudioNext A N  Exec exec gmm --next

### Rhythmbox

#### Rhythmbox dbus

    Key XF86AudioNext A N  Exec exec dbus-send --type=method_call --dest=org.mpris.MediaPlayer2.rhythmbox /org/mpris/MediaPlayer2 org.mpris.MediaPlayer2.Player.Next
    Key XF86AudioPrev A N  Exec exec dbus-send --type=method_call --dest=org.mpris.MediaPlayer2.rhythmbox /org/mpris/MediaPlayer2 org.mpris.MediaPlayer2.Player.Previous
    Key XF86AudioPlay A N  Exec exec dbus-send --type=method_call --dest=org.mpris.MediaPlayer2.rhythmbox /org/mpris/MediaPlayer2 org.mpris.MediaPlayer2.Player.PlayPause
    Key XF86AudioStop A N  Exec exec dbus-send --type=method_call --dest=org.mpris.MediaPlayer2.rhythmbox /org/mpris/MediaPlayer2 org.mpris.MediaPlayer2.Player.Stop

#### rhythmbox-client

    Key XF86AudioNext A N  rhythmbox-client --next
    Key XF86AudioPrev A N  rhythmbox-client --previous
    Key XF86AudioPlay A N  rhythmbox-client --play-pause
    Key XF86AudioStop A N  rhythmbox-client --stop

## Управление громкостью (amixer)

Ниже пример с PCM:

    Key XF86AudioRaiseVolume  A N  Exec exec amixer sset PCM 2+
    Key XF86AudioLowerVolume  A N  Exec exec amixer sset PCM 2-
    Key XF86AudioMute         A N  Exec exec amixer sset PCM toggle

В варианте с PCM может отсутствовать возможность toggle. Ошибок в stdout
не выводится, но и не работает переключение mute/unmute.

Или в варианте с Master:

    Key XF86AudioRaiseVolume  A N  Exec exec amixer sset Master 2+
    Key XF86AudioLowerVolume  A N  Exec exec amixer sset Master 2-
    Key XF86AudioMute         A N  Exec exec amixer sset Master toggle

# Ссылки на сайты, связанные с fvwm

[FVWM(англ.)](http://fvwm.org) -- основной сайт.
[Форум FVWM(англ.)](http://www.fvwmforums.org) -- основной форум.
[zensites guide FVWM(англ.)](http://zensites.net/fvwm/guide/) -- гайд по
заточке системы, с некоторыми интересными функциями.
[CookBook(англ.)](http://www.xteddy.org/fvwm/fvwmcookbookfaq.html) --
рецепты. К сожалению, сайт больше не пополняется. Надеемся, что это
временно.
[box-look или примеры
настройки(англ.)](http://box-look.org/index.php?xcontentmode=7314)
-- давольно ценный сайт, можно посмотреть и скачать темы оформления
пользователей.
[советы по настройке
fvwm](http://www.altlinux.org/DotFiles/WindowManagers/FVWM) -- старые
советы, но ещё актуальные.
[Описание оконного менеджера
FVWM](http://linuxland.itam.nsc.ru/misc/fvwm/fvwm.html) -- описываются
модули, анатомия окон и прочее. Увы, всё старо, но не бесполезно.
[топик: Строим свой wm (с преферансом и
куртизанками)](http://forum.ubuntu.ru/index.php?topic=67218.0)
-- убунту-форум. Для многих из современности знакомство с fvwm началось
c просмотра этой темы.
[Комментарии этой статьи](http://www.linux.org.ru/wiki/en/Comments:FVWM)
содержат множество полезных идей и каких-либо решений по вопросам FVWM,
если вам интересно почитать неформальные обсуждения про FVWM.