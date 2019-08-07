''' Оптимизация размеров gtk2 виджетов '''

<u>1. Отступы</u>

Проблема из-за чего в gtk2 все виджеты выглядят слишком большими, это
отступы от текста по горизонтали и вертикали на них. В gtkrc эти
отступы задаются параметрами xpadding и ypadding.

Для начала глянем что за размер иконок у нас выставлен в теме:

    gtk-icon-sizes = "panel-menu=24,24:panel=24,24:gtk-button=24,24:gtk-large-toolbar=24,24"

Это значит что на кнопках у нас будут рисоваться иконки 24x24 пиксела.
Меня устраивает следующая установка

    gtk-icon-sizes = "gtk-large-toolbar=16,16:gtk-small-toolbar=16,16:gtk-button=16,16"

Далее ищем дефолтовую секцию для настроек всех виджетов ибо в ней всегда
есть много интересного:

    class "GtkWidget"          style "default"

Окей, значит стиль называется default. Ищем его в теме:

    style "default"
    {
      xthickness                = 2
      ythickness                = 2
      fg[NORMAL]        = "#353535"
    }

Вот наши отступы. Во всех виджетах, отступы от текста будут 2 пикселя,
если для конкретного виджета не задан свой \[xy\]thickness. Можно
попробовать уменьшить отступ до 1 или до 0 пикселя, по вскусу. Но,
в некоторых виджетах этот отступ очень даже полезен. Поэтому лучше
редактировать отступы прямо в конкретном виджете.

Например, уменьшим размер виджета GtkButton (или просто "кнопка"). Ищем
описание нашего GtkButton виджета:

    class "GtkButton"                       style "button"

Стиль нашего виджета в теме завется button. Ищем его описание:

    style "button"          = "default"
    {
    }

Если вы нашли там \[xy\]padding, то смело меняйте их размеры на 0 или 1
пиксел, по вкусу. если их там нет, то просто добавьте 2 новые строчки:

    style "button"          = "default"
    {
      xthickness                = 0
      ythickness                = 0
    }

Сохраняем Ваши изменения и перезапускаем какую-нибудь gtk программу.
Размер кнопок должен быть значительно меньше.

<u>2. Лечим размер табов.</u>

Таким же образом можно поправить размер у табов. Виджет "табы"
называется GtkNotebook

Итак, рецепт. В default секцию добавляем:

    style "default"
    {
      GtkNotebook::tab-border = 0
      GtkNotebook::tab-hborder = 0
      GtkNotebook::tab-vborder = 0
      GtkNotebook::show-border = 0
      GtkNotebook::gtk-button-images = 0
      GtkNotebook::gtk-menu-images = 0
      GtkNotebook::arrow-spacing = 0
      GtkNotebook::tab-curvature = 0
      GtkNotebook::tab-overlap = 0
      GtkNotebook::focus-line-width = 0
    }

Сохраняем, перезапускаем gnome-terminal и видим отличные красивые табы.

<u>3. Полезные переменные</u>

Для уменьшения размеров виджетов в секцию default можно добавить
следующие переменные:

    gtk-button-images = 0

переменная отключает отображение иконок на кнопках

    gtk-menu-images = 0

переменная отключает отображение иконок в меню

<u>4. Убираем стрелочки на скроллбарах</u>

Никогда не понимал их пользы - прокрутка обычно ведётся либо колесом
мыши, либо перетаскиванием ползунка, либо с клавиатуры.

    style "default"
    {
      GtkScrollbar::has-backward-stepper=0
      GtkScrollbar::has-forward-stepper=0
    }

<u>5. Уменьшаем ширину скроллбаров</u>

Весьма полезно, если скроллбар служит только для индикации положения на
странице.

    style "default"
    {
        GtkRange::trough_border = 0
        GtkRange::slider_width = 10
    }

<u>6. Убираем границы вокруг меню, панелей инструментов и строки
состояния</u>

Если хочется избавиться от границ вокруг данных элементов, то добавьте
следующие настройки к стилю:

    style "default"
    {
        GtkStatusbar::shadow_type = GTK_SHADOW_NONE
        GtkToolbar::shadow-type = GTK_SHADOW_NONE
        GtkMenuBar::shadow-type = GTK_SHADOW_NONE
    }

Возможные значения параметра shadow-type:

    GTK_SHADOW_NONE
        No outline.

    GTK_SHADOW_IN
        The outline is bevelled inwards.

    GTK_SHADOW_OUT
        The outline is bevelled outwards like a button.

    GTK_SHADOW_ETCHED_IN
        The outline has a sunken 3d appearance.

    GTK_SHADOW_ETCHED_OUT
        The outline has a raised 3d appearance.

Описание синтаксиса gtkrc смотрите на
[1](http://library.gnome.org/devel/gtk/unstable/gtk-Resource-Files.html).
Небольшой урок по созданию тем можно посмотреть на
[2](http://live.gnome.org/GnomeArt/Tutorials/GtkThemes). Информация по
движкам GTK-тем представлена на
[3](http://live.gnome.org/GnomeArt/Tutorials/GtkEngines).

Предлагаю развивать дальше эту статью.

-----

\--[Alexander Butenko](User:mrdeath "wikilink") 09-Sep-2008 08:50 MSD

''' Настройка диалога выбора файлов '''

По-умолчанию, диалог выбора файлов показывает список последних открытых
документов. Чтобы изменить это поведение, необходимо в
$XDG_CONFIG_HOME/gtk-2.0/gtkfilechooser.ini поменять значение
параметра StartupMode с recent на cwd.

[Category:Программирование](Category:Программирование "wikilink")