"Из коробки" пока, к сожалению, на этой клавиатуре движок "Zoom" не
работает. Однако, его можно настроить. Так как настройка делается
при помощи xdotool, можно генерировать zoom'ом любые клавиатурные или
"мышиные" последовательности. Итак, краткий алгоритм настройки:

1.  Создаем правило keymap (файл
    /lib/udev/keymaps/microsoft-ergonomic-keyboard) для назначения
    скан-кодов нераспознаваемым сигналам от zoom:

`   0xC022D 0xC1 # Zoom Up which we wish to be Scroll up`
`   0xC022E 0xC2 # Zoom Down which we wish to be Scroll down`

1.  Создаем файл /etc/udev/rules.d/ms_ergo.rules для автозапуска keymap
    с содержимым

`   SUBSYSTEM=="input", ATTRS{manufacturer}=="Microsoft", RUN+="keymap $name microsoft-ergonomic-keyboard"`

1.  Назначаем на новые скан-коды сигналы колеса мыши. Для этого создаем
    (если не было) файл \~/.xbindkeysrc и заносим в него

`   "xdotool click 4"   # Scroll Up`
`   c:201`
`   "xdotool click 5"   # Scroll Down`
`   c:202`

1.  В автозапуск своего DE/WM прописываем xbindkeys. Например, в IceWM
    для этого нужно добавить в файл \~/.icewm/startup строчку

`   xbindkeys &  # чтобы zoom на клавиатуре работал`

Естественно, чтобы все работало, нужно установить xdotool и xbindkeys.

После всех вышеперечисленных телодвижений "zoom" начинает работать как
прокрутка колесом мыши.

советы](Category:Полезные_советы)