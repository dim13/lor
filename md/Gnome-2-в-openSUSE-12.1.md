Официально openSUSE 12.1 не предоставляет пакеты со вторым гномом для
данного релиза, но установить его все же можно и это не потребует
много труда или утомительных "плясок с бубном" =\]  
Это справедливо так же и для openSUSE 12.2 или openSUSE 12.3 (актуальный
релиз).  

## Предупреждение

**Важно\!**

  - Все ниже перечисленные операции проводятся на Ваш страх и риск,
    помните это.
  - Крайне не рекомендуется ставить второй гном в систему, где
    установлен третий гном. Установка Gnome 2 успешно
    проверена из KDE4, LXDE и Xfce.

## Установка

Запускаем эмулятор терминала и вводим следующие команды:

    su -

    zypper ar http://download.opensuse.org/distribution/11.4/repo/oss/suse/ g2 

(добавляем репозиторий openSUSE 11.4, где есть нужный нам Gnome 2)

    zypper mr -p 80 g2 

(задаем приоритет для нашего репозитория выше, чем у всех остальных)

    zypper ref

(обновляем информацию о пакетах)

    zypper in metacity gnome-settings-daemon gnome-control-center \
    nautilus gnome-desktop gnome-menus gnome-menus-lang gnome-session \
    gnome-panel gnome-icon-theme gnome-themes gnome-terminal gdm gedit 

(ставим основные пакеты Gnome 2. **Важно\!** Если zypper спросит о
замене пакетов из третьего гнома версиями из второго –
соглашайтесь)

    yast2 sysconfig set DISPLAYMANAGER=gdm 

(делаем GDM менеджером входа по-умолчанию)

``` 
yast2 sysconfig set DEFAULT_WM=gnome  
```

(делаем Metacity менеджером окон по-умолчанию)

## Что делать дальше?

Русифицируем гном

    su -
    zypper install --force bundle-lang-gnome-ru-11.4-5.13.1.noarch 

Ставим дополнительные компоненты для второго гнома (например, totem и
file-roller)

    zypper install totem file-roller

Как только мы всё поставим, репозиторий от openSUSE 11.4 нам уже не
нужен и мы смело можем его удалить из списка используемых:

    zypper rr g2

(но если он вдруг понадобиться, Вы уже знаете, как его добавить)  
И, конечно же, **не забудьте** перезагрузить систему, когда все операции
будут проведены =\]

## Как теперь обновляться?

Да, если сейчас Вы попробуете обновиться, zypper непременно затрет
пакеты от второго гнома пакетами от третьего. Поэтому нам
необходимо заблокировать изменения этих пакетов. Итак,
добавляем блокировки:

    su -
    zypper addlock *gnome* *gconf* *nautilus* *metacity* gedit* gdm

Плюс ко всему, поставьте блокировки и для мелочи, которую ещё ставили от
второго гнома (totem, evince, file-roller и т.д.). После этого можно
обновляться, но при первом обновлении все же будьте внимательны и
проверяйте список того, что zypper хочет обновить.

    zypper ref && zypper up

## Решение проблемы со звуком

Я установил GNOME2 по этому руководству и получил проблему со звуком. В
оригинальном Opensuse 11.4 то же самое. Несколько игр в Wine работало
без звука, выдавая в консоль то что OpenAL не может соединиться с
PulseAudio. ICQ for Linux выводил свои уведомления, захватив устройство
звука монопольно. VLC Media Player перед воспроизведением выдавал
"невозможно использовать устройство default" и [такой лог в
консоль](http://paste.org.ru/?eqzsfw). При этом в MATE всё
работает.

В Wine PulseAudio отключаем в YAST одной галочкой, и для KDE этого
достаточно. В GNOME оказалось что нет. Нужно выполнить "gnomesu
gedit" и отредактировать /etc/asound-pulse.conf. Закомментировать знаком
\# все строчки. После этого всё заработает. Осталось только пересобрать
пакет с gnome-sound без PulseAudio, чтобы вернуть на место регулятор
громкости.

Не забудьте вернуть как было, если включите PulseAudio. А вообще это не
решение проблемы, а "костыль". Решение проблемы - перенастройка GNOME2
так, чтобы он не брал настройки из того файла.

[Category:openSUSE](Category:openSUSE "wikilink")
