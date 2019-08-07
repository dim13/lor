По сути DE - набор из панели (или средства для её предоставления), WM и
стандартных программ.
Перед сборкой DE из разных кусков попробуйте существующие решения.

# DE

## Enlightenment

Сайт проекта: <http://www.enlightenment.org/>
Скриншот:
[1](http://upload.wikimedia.org/wikipedia/ru/thumb/b/ba/Eshot-1680x1050.png/800px-Eshot-1680x1050.png).
Библиотека: EFL (Enlightenment Foundation Libraries)
Менеджер входа: elsa
Менеджер окон: Enlightenment, \[возможен другой?\]
Менеджер сессий: \[???\]
Разработчики: [список
разработчиков](http://www.enlightenment.org/p.php?p=contact&l=en)
Позиционирует как оболочка для слабых компьютеров, мобильных телефонов,
и других систем с малой производительностью (известно, что на базе
библиотеки EFL компания Samsung сделала интерфейс для контрольной
панели серии холодильников).

## GNOME

Сайт проекта: <http://www.gnome.org/>
Библиотека: GTK+ 3
Менеджер входа: GNOME Display Manager, возможен другой
Менеджер сессий: GNOME 3 Session manager
Разработчики: [The GNOME Foundation](http://foundation.gnome.org/)

Ныне самое распространенное окружение рабочего стола для Линуксов,
включенная по-умолчанию во многие дистрибутивы.

В состав GNOME кроме оболочки входит также и набор всех необходимых
большинству людей приложений.

Стандартные средства настройки очень скудные (можно настроить только
обоину, заставку и ещё несколько мелочей). Настройка таких вещей,
как: тема GTK+, тема иконок, тема курсора, шрифты, поведение
менеджера окон и ещё несколько мелочей доступно через GNOME
Tweak Tool (пакет gnome-tweak-tool).

В состав GNOME входит две разных оболочки: GNOME Fallback и GNOME Shell.

### Gnome Fallback

Скриншот: [2](http://www.linux.org.ru/gallery/6690197.png).
Менеджер окон: по умолчанию Metacity, но могут использоваться и
другие.
Позиционирует как оболочка для слабых и среднестатистических
компьютеров, а также для тех, кому более приятен интерфейс
Gnome 2 нежели Gnome Shell.

### Gnome Shell

Скриншот:
[3](http://upload.wikimedia.org/wikipedia/commons/thumb/9/97/GNOME_Shell.png/800px-GNOME_Shell.png).
Менеджер окон: Mutter (по сути оболочка GNOME Shell это плагин к
Mutter).
Позиционирует как оболочка для среднестатистических компьютеров и
планшетов.
Возможна кастомизация оболочки через GNOME Tweak Tool посредством
расширений (так званные GNOME Shell Extionsions) и тем. Для
установки тем Шелла необходимо установить расширение User-Theme
Extension (пакет gnome-shell-extension-user-theme).

## KDE

### KDE 4

Сайт проекта: <http://kde.org/>
Скриншот:[4](http://upload.wikimedia.org/wikipedia/commons/5/54/KDE_4.png).
Библиотеки: Qt 4, KDELibs
Менеджер входа: KDE Display Manager, возможен другой
Менеджер окон: K Window manager, возможен другой
Разработчики: разрабатывается сообществом
Позиционирует как оболочка для среднестатистических компьютеров.

В состав окружения также, как и в GNOME входят все необходимые обычному
пользователю приложения.

#### Plasma Desktop

Plasma Desktop — это фреймворк для виджетов для рабочего стола,
написанный в рамках KDE 4.

### Trinity

Сайт проекта: <http://www.trinitydesktop.org/>
Скриншот:
[5](http://upload.wikimedia.org/wikipedia/commons/thumb/a/a0/Screenshot_of_Trinity_3.5.12.png/800px-Screenshot_of_Trinity_3.5.12.png).
Библиотеки: Qt 3, TDELibs
Менеджер входа: TDE Display Manager, возможен другой
Менеджер окон: T Window manager, возможен другой
Разработчики: разрабатывается сообществом
Позиционирует как оболочка для слабых и среднестатистических
компьютеров.
Форк KDE 3.

## LXDE

Сайт проекта: <http://lxde.org/>
Скриншот:
[6](http://upload.wikimedia.org/wikipedia/commons/4/4c/LXDE_desktop_full.png).
Библиотека: GTK+ 2
Менеджер входа: произвольный
Менеджер окон: по умолчанию Openbox, возможна установка других WM
Разработчики: Команда разработчиков LXDE
Позиционирует как оболочка для слабых компьютеров.

Название расшифровывается как **L**ightweight **X**11 **D**esktop
**E**nvironment. Примечательно то, что создатели не стремились тесно
интегрировать различные компоненты, наоборот, каждый из них может
использоваться отдельно с несколькими зависимостями.

## MATE

Сайт проекта: <https://github.com/Perberos/Mate-Desktop-Environment>
Скриншот: [7](http://www.linux.org.ru/gallery/6667506.jpg).
Библиотека: GTK+ 2 (планируют переходить на GTK+ 3)
Менеджер входа: произвольный.
Менеджер окон: Mate window manager, возможен другой
Менеджер сессий: Mate session manager
Разработчики: разрабатывается сообществом
Позиционирует как оболочка для среднестатистических компьютеров.
Форк 2-ой версии GNOME.

## Razor-Qt

Сайт проекта: <http://razor-qt.org/>
Скриншот: <http://razor-qt.org/screenshots/images/01.Desktop.png>.
Библиотека: Qt 4
Менеджер входа: произвольный
Менеджер сессий: Razot-Qt Session Maneger
Менеджер окон: произвольный (разработчики рекомендуют использовать
Openbox)
Разработчики: команда из нескольких разработчиков (в т.ч. и русские)
Позиционирует как оболочка для слабых компьютеров. В плане
производительности может сравниться с LXDE.

Форк ныне не развивающегося Antico. В состав оболочки входят только
панель, переключалка окон (по нажатию комбинации клавиш Альт+Таб),
запускалка приложений, рабочий стол, менеджер сессий и библиотека
обеспечивающая функционал всего этого. Зависимостей мало. KDELibs
не тянет.

Небольшой список рекомендованных легковесных приложений на Qt, которые
могут пригодиться, можно найти
[тут](https://github.com/Razor-qt/razor-qt/wiki/3rd-party-applications).

## Unity

Сайт проекта: <http://unity.ubuntu.com/>
Скриншот:
[8](http://upload.wikimedia.org/wikipedia/commons/4/4f/Ubuntu_11.04_ru.png).
Библиотека: GTK+ 3
Менеджер входа: LightDM, возможен другой
Менеджер окон: Compiz (Unity — это по сути плагин к Compiz)
Менеджер сессий: gnome-session
Разработчики: [Canonical](http://www.canonical.com/)
Позиционирует как оболочка для среднестатистических компьютеров и
планшетов. Планируется в далеком будущем перенос на мобильные
телефоны.

Оболочка рабочего стола по умолчанию в дистрибутиве Ubuntu. Основан на
GNOME.

## Xfce

Сайт проекта: <http://www.xfce.org/>
Скриншот:
[9](http://upload.wikimedia.org/wikipedia/commons/7/71/Xfce-4.4.png)
Библиотека: GTK+ 2
Менеджер входа: произвольный
Менеджер окон: Xfwm, возможен другой
Менеджер сессий: Xfce4 Session Manager
Разработчики: Проект Xfce
Позиционирует как оболочка для слабых и среднестатистических
компьютеров. Многие ошибочно полагают, что это оболочка
заточенная для работы на слабых машинах. Да, это правда, но по
потреблению ресурсов она находится между такими, как MATE и LXDE,
Razor-Qt.

# WM

Менеджеры окон делятся на стековые, композитные и тайловые, они же
фреймовые. Некоторые тайловые могут работать в режиме стековых
(awesome). Стековые управляют окнами в привычном нам режиме, композитные
ещё и придают им разных эффектов и декораций (например прозрачность и
куб рабочего стола), тайловые же разбивают рабочее пространство
экрана на взаимно не пересекающиеся прямоугольные области —
фреймы. Есть ещё динамические оконные менеджеры — они работают как
тайловые, только там размещение фреймов предварительно настраивается
(пример — Awesome). Также есть приложения, позволяющие добавить
стековым менеджерам некоторый функции характерные для композитных
(Xcompmgr).

## Композитные

### Compiz

Размер: \~25 кб. (много зависимостей)
Скриншот:
[10](http://upload.wikimedia.org/wikipedia/commons/thumb/c/cb/Fedora-Core-6-AIGLX.png/800px-Fedora-Core-6-AIGLX.png)
Настройка: графический конфигуратор.
Статус: актуальный.

### KWin

Размер: \~3 мб.
Скриншот:
[11](http://upload.wikimedia.org/wikipedia/commons/thumb/c/c6/Kwin-coverswitch4.1_beta_1.png/800px-Kwin-coverswitch4.1_beta_1.png)
Настройка: графические средства.
Тайловый: опционально.
Статус: актуальный.
Разрабатывается в рамках проекта KDE.

### Xfwm

Размер: \~9347 кб.
Скриншот:
[12](http://upload.wikimedia.org/wikipedia/commons/7/71/Xfce-4.4.png)
Настройка: графическая утилита.
Статус: актуальный. Разрабатывается в рамках проекта Xfce.

### Mutter

Размер: \~6.5 мб.
Скриншот:
[13](http://upload.wikimedia.org/wikipedia/commons/thumb/9/97/GNOME_Shell.png/800px-GNOME_Shell.png)
Настройка: графические средства (очень ограниченная настройка).
Статус: актуальный. Является продолжением Metacity. Название — это
сокращение от „Metacity Clutter“ (Clutter - это новая
библиотека-тулкит для отрисовки интерфейса приложений
средствами OpenGL). Служит основой для GNOME Shell (по сути GNOME
Shell это плагин для Mutter).
[Может](http://en.wikipedia.org/wiki/Mutter_\(window_manager\))
функционировать, как отдельный оконный менеджер отдельно без GNOME
Shell. Наличие GPU не обязательно.

## Стековые

### Openbox

Размер: \~300 кб.
Скриншот:
[14](http://upload.wikimedia.org/wikipedia/commons/thumb/b/b0/Openbox-elementary3.png/800px-Openbox-elementary3.png)
Настройка: правка текстовых конфигурационных файлов, возможно
использование графического средства настройки.
Статус: актуальный.

### Enlightenment

Размер: \~3 мб.
Скриншот:
[15](http://upload.wikimedia.org/wikipedia/ru/thumb/b/ba/Eshot-1680x1050.png/800px-Eshot-1680x1050.png)
Настройка: GUI
Тайловый: плагином.
Статус: актуальный

### Fluxbox

Размер: \~1.5 мб.
Скриншот:
[16](http://upload.wikimedia.org/wikipedia/commons/thumb/b/b1/Flux.jpg/800px-Flux.jpg)
Настройка: редактирования текстовых файлов конфигурации, или
[fluxconf](http://devaux.fabien.free.fr/flux/).
Статус: актуальный.

### FVWM

Размер: \~4 мб.
Скриншот:
[17](http://upload.wikimedia.org/wikipedia/commons/thumb/e/ed/Fvwm1-wikipedia-20050312.png/800px-Fvwm1-wikipedia-20050312.png)
Настройка: редактирования текстовых файлов конфигурации.
Статус: актуальный.
Примечание: прабатюшка большинства современных wm.

### Metacity

Размер: \~500 кб.
Скриншот:
[18](http://upload.wikimedia.org/wikipedia/commons/e/ec/Metacity-screenshot.png)
Настройка: графические средства (ограниченная настройка), gconf (там
настроек побольше).
Статус: устарел. Разрабатывался в рамках GNOME 2.

### Sawfish

Размер: ???
Скриншот:
[19](http://upload.wikimedia.org/wikipedia/commons/thumb/b/ba/Sawfish.png/750px-Sawfish.png)
Настройка: ???
Статус: актуальный.

### IceWM

Размер: \~750 кб.
Скриншот:
[20](http://upload.wikimedia.org/wikipedia/commons/thumb/9/97/IceWM.png/750px-IceWM.png)
Настройка: как прямое редактирование текстовых файлов так и графические
средства.
Статус: разработка прекращена (fixme, если что).

### AfterStep

Размер: \~400 кб.
Скриншот:
[21](http://upload.wikimedia.org/wikipedia/commons/thumb/6/68/AfterStep.png/750px-AfterStep.png)
Настройка: редактирования текстовых файлов конфигурации.
Статус: в разработке, но очень медленной. \[fixme\]

### JWM

Размер: \~100 кб.
Скриншот:
[22](http://upload.wikimedia.org/wikipedia/commons/thumb/4/48/Openbsd37withjwm.png/800px-Openbsd37withjwm.png)
Настройка: ???
Статус: актуальный.
Примечание: Используется на очень слабых системах и "легких"
дистрибутивах.

### Blackbox

Размер: \~250 кб.
Скриншот:
[23](http://upload.wikimedia.org/wikipedia/commons/thumb/3/30/Blackbox_on_Debian.png/750px-Blackbox_on_Debian.png)
Настройка: редактирования текстовых файлов конфигурации.
Статус: актуальный.

## Тайловые

### Awesome

Размер: \~850 кб.
Пакет: awesome
Скриншот:
[24](http://upload.wikimedia.org/wikipedia/commons/thumb/f/f0/Awesome_screenshot.png/750px-Awesome_screenshot.png)
Настройка: один конфигурационный файл (на Lua).
Тайловый: возможность переключения.
Статус: актуальный.

### DWM

Размер: \~40 кб
Пакет: dwm
Скриншот:
[25](http://upload.wikimedia.org/wikipedia/commons/thumb/b/b0/Dwm-shot.png/800px-Dwm-shot.png)
Настройка: правка исходных кодов и перекомпиляция.
Статус: актуальный.

### Ion

Размер: ???
Пакет: ???
Скриншот:
[26](http://ru.wikipedia.org/wiki/%D0%A4%D0%B0%D0%B9%D0%BB:Screenshot-1.png)
Настройка: ???
Статус: актуальный.

### Stumpwm

Размер: \~300 кб.
Скриншот:
[27](http://upload.wikimedia.org/wikipedia/commons/e/eb/Stumpwm_12-2-2006.png)
Настройка: ???
Статус: актуальный.

### Window Manager Improved 2

Размер: \~400 кб.
Скриншот:
[28](http://upload.wikimedia.org/wikipedia/commons/thumb/5/54/Wmii-20070228.png/750px-Wmii-20070228.png)
Настройка: консольная утилита
Статус: актуальный.

### Xmonad

Размер: \~600 кб.
Скриншот:
[29](http://upload.wikimedia.org/wikipedia/commons/thumb/6/64/Xmonad-tall-status-dons.png/800px-Xmonad-tall-status-dons.png)
Настройка: путем написания исходного кода на Haskell.
Статус: актуальный.

# Панель/Окружение

## Окружения

### Plasma

Настройка: Графические утилиты.
Тулкит: Qt (QML)
Часть проекта KDE4, которая предоставляет API для написания виджетов
(плазмоидов) рабочего стола, в том числе панель.
Скриншот:
[30](http://upload.wikimedia.org/wikipedia/commons/5/54/KDE_4.png)
\=== GNOME Shell ===

Cреда рабочего стола, созданная для GNOME3. Очень тесно интегрирована с
Mutter.
Настройка: Графические утилиты.
Скриншот:
[31](http://upload.wikimedia.org/wikipedia/commons/9/97/GNOME_Shell.png)

## Панели

### Xfce4 panel

Панель, написанная разработчиками xfce.
Настройка: Графические утилиты.
Тулкит: GTK+ 2.
Скриншот:
[32](http://upload.wikimedia.org/wikipedia/en/b/bc/XFCE-4.8-Desktop.png)

### tint2

Панель, созданная для openbox.
Настройка: Правка конфигурационного файла/GUI.
Скриншот:
[33](http://img252.imageshack.us/img252/1433/wallpaper2td.jpg)

### lxpanelx

Улучшенный форк lxpanel.
Настройка: Графические утилиты.
Тулкит: GTK+ 2.
Скриншот: [34](http://i015.radikal.ru/1201/52/bf358d3416ca.png)

### AWN

Док, наподобие Dock из Mac OS X.
Настройка: Графические утилиты.
Скриншот:
[35](http://upload.wikimedia.org/wikipedia/commons/thumb/b/bb/Awn_screenshot.png/250px-Awn_screenshot.png)

## WM, содержащие панель/окружение

  - BlackBox
  - WindowMaker
  - FluxBox
  - IceWM
  - JWM
  - Windowlab