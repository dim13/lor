Unreal — серия компьютерных игр в жанре «шутер от первого лица»,
действие которой происходит в одноимённой вселенной. Следует
различать игры, предназначенную для одиночного прохождения (Unreal,
Unreal II) и созданные специально для сетевой игры (Unreal Tournament,
Unreal Tournament 2003/2004, Unreal Tournament 3). Разработчик всех игр
серии и одноимённого игрового движка — компания Epic MegaGames (позже
переименованная в Epic Games).

## Unreal

Unreal — первая игра серии, вышла в 1998 году. Долгое время поиграть в
неё можно было только [поверх](http://icculus.org/~ravage/unreal/)
Unreal Tournament (который имел тот же игровой движок и существовал в
версии для Linux) Ещё есть
[скрипт](http://www.margrave.myzen.co.uk/downloads/installationutilities/unreal-gold-anthology-linux-installer_v.0.2.7z)
для установки Unreal Gold с диска Unreal Anthology. Однако, с недавних
пор появилась возможность играть под Linux в нативный Unreal — на
сайте Old Unreal был выложен неофициальный
[патч](http://www.oldunreal.com/oldunrealpatches.html), содержащий,
в том числе, и бинарные файлы для Linux. К сожалению, патч
устанавливается только на windows-версию Unreal при помощи
wine, но работает достаточно стабильно.

[Linux
FAQ](http://www.oldunreal.com/wiki/index.php?title=Oldunreal_227_Linux_FAQ)

Существую неофициальные инсталляторы. Подробности на [форумах
OldUnreal](http://www.oldunreal.com/cgi-bin/yabb2/YaBB.pl)

## Unreal Tournament (UT99)

Unreal Tournament — игра-продолжение серии Unreal, специально созданная
для сетевых баталий. Оригинальная двухдисковая версия Unreal Tournament
содержит на первом диске установщик для Linux. А вот издание Unreal
Anthology такого установщика не имеет (кстати, в это издание еще и
«забыли» положить набор S3TC текстур высокого разрешения).
Установщик можно взять здесь -
[1](http://www.liflg.org/?catid=6&gameid=51) , существует неофициальный
[патч](http://www.utpg.org/patches/UTPGPatch451.tar.bz2) до версии 451.
Проблема состоит в том, что родной установщик от loki требует
оригинальный диск, но к счастью, добрые люди уже позаботились
о нас и написали готовый установочный
[скрипт](http://www.margrave.myzen.co.uk/downloads/installationutilities/ut-anthology-linux-installer_v.0.1.7z)
для Unreal Anthology, который работает в том числе и с антологией от
«Нового диска». Для работы скрипта понадобится установленный в
системе unshield. Текстуры высокого разрешения можно найти
[тут](http://www.unrealtexture.com/UT/UT.htm). Так же можно
установить UT используя установленную версию для Windows с
диска Unreal Anthology при помощи [обновлённого
установщика](http://necrovision.ru/mfiles/unreal/ut-install-436-anthology.run),
для этого надо указать откуда ставить UT: <code>

` $ export SETUP_CDROM=/путь/к/UnrealTournament`
` $ sh ut-install-436-anthology.run`

</code>

### В Unreal Tournament (UT99) все носятся как сумасшедшие, невозможно играть

Это происходит из-за того, что игра слишком быстро работает на
современных процессорах. Мне помогло включение в настройках
OpenGL вертикальной синхронизации (опция Sync to VBlank во вкладке
OpenGL Settings панели управления nvidia-setting), правда fps стал
постоянно равен 60, есть еще стартовый скрипт, который, по-идее,
загружает процессор и снижает его производительность:
[2](http://icculus.org/lgfaq/files/ut).

### В Unreal Tournament (UT99) нет звука

Существует несколько причин. Установите
[патч 451](http://www.utpg.org/patches/UTPGPatch451.tar.bz2). В
конфигурационном файле \~/.loki/ut/System/UnrealTournament.ini
попробуйте заменить строчку `AudioDevice=ALAudio.ALAudioSubsystem`
на `AudioDevice=Audio.GenericAudioSubsystem` и наоборот. Убедитесь, что
у вас в системе установлен пакет alsa-oss (запуск: `aoss ./ut-bin`).
Если в системе стоит PulseAudio, то можно OSS отправлять сразу в
него: `padsp ./ut-bin`

## Unreal II: The Awakening

Unreal II: The Awakening — прямое продолжение первой части. Никогда не
портировалась под Linux. Под wine работает с переменным успехом.

## Unreal Tournament 2004 (UT2k4)

Unreal Tournament 2004 — улучшенный UT2003, очень удачное продолжение
серии. Аналогично UT99, оригинальная версия UT2k4 содержала
установщик для Linux, и, точно также, в Unreal Antology его не
включили. Чтобы поиграть, вам потребуется последний патч, например,
[отсюда](http://download.beyondunreal.com/fileworks.php/official/ut2004/ut2004-lnxpatch3369-2.tar.bz2).
Скопируйте рабочий UT2k4 с windows-раздела (если есть) в каталог
\~/UT2k4 или установите с его помощью wine/cedega. Затем распакуйте патч
3369-2 в этот же каталог с заменой файлов. Создайте ссылки на библиотеки
libopenal.so и libSDL-1.2.so.0 из /usr/lib в UT2004/System с названием
openal.so и libSDL-1.2.so.0 соответственно. Создайте
\~/.ut2004/System/cdkey и впишите в него свой ключ для игры. Для
удобства можете создать в /usr/bin следующий скрипт для запуска
игры:

```

#!/bin/sh
cd "/home/<user>/UT2004/System"
exec ./ut2004-bin "$@"
```

Разумеется, для ut2004-bin должны быть выставлены права на исполнение.

### Как запустить под Linux пользовательские модификации для Unreal Tournament 2004 (например Tactical Ops)?

Если пользовательская модификация (мод) распространяется в виде \*.zip,
или, ещё лучше, \*.tar.bz2 файла, то вам повезло. От вас требуется лишь
распаковать содержимое архива в директории \~/ut2k4/<название мода> и
создать в каталоге \~/ut2k4/System скрипт вида

```

#!/bin/sh
cd "/home/<user>/UT2004/System"
 ./ut2004-bin -MOD=TOCrossfire -log=TOCrossfire.log
```

для запуска модификации и создать на него ссылку в /usr/bin . Еще лучше,
если архив с модом содержит готовый установочный скрипт. Вам останется
только запустить его.

## Unreal Tournament 3 (UT3)

Самая новая и технологичная игра в серии на сегодняшний день. К
сожалению, портирование под Linux прекращено.

## Ссылки

  - [Planet Unreal](http://planetunreal.gamespy.com)
  - [BeyondUnreal](http://www.beyondunreal.com/) — сайт сообщества с
    множеством дочерних сайтов.
  - [UnrealAdmin.org](http://www.unrealadmin.org/) — сайт и форум для
    админов UT-серверов
  - [Unreal.ru](http://www.unreal.ru/) — Центральный сайт UT-сообщества
    России.
  - [OldUnreal.com](http://www.oldunreal.com/) — Сайт OldUnreal,
    сообщайте о багах в версии 227 сюда.
  - [Hyper.nl Unreal Services](http://hypernl.thenerdnetwork.net) — Сайт
    Hyper.nl, множество файлов.
  - [Русский Planet Unreal](http://www.planetunreal.ru/) — Русское
    сообщество анрилеров.

[category:Игры](category:Игры)