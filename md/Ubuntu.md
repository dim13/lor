Ubuntu — достаточно молодой дистрибутив, основанный на Debian, но
несмотря на это он сумел завоевать большую популярность и встать
на один уровень с Fedora, openSUSE и Mandriva.

[Официальный сайт](http://ubuntu.com/)

[Русское сообщество пользователей Ubuntu](http://ubuntu.ru)

## Что означает «Ubuntu»?

“Ubuntu” is an ancient African word, meaning “humanity to others”

«Ubuntu» — древнее африканское слово, означающее «гуманность к другим»

## В чем разница между Ubuntu, Kubuntu, Xubuntu и другими?

В Ubuntu по умолчанию в качестве рабочей среды используются GNOME (с
версии 11.04 — Unity с возможностью выбрать классический режим GNOME
2), в Kubuntu — KDE, а в Xubuntu — Xfce.
[Edubuntu](http://www.edubuntu.org/) ориентирован на использования в
образовательных учреждениях и содержит большое количество программ
образовательного направления. [UbuntuStudio](http://ubuntustudio.org/)
предназначен для творческих людей, занимающихся работой с графикой,
аудио и видео. Также содержит множество мультимедийных утилит, в
качестве DE — GNOME. Все эти дистрибутивы основаны на общем
репозитории и с легкостью превращаются друг в друга, для этого
существуют мета-пакеты ubuntu-desktop, kubuntu-desktop, xubuntu-desktop
и т.д.

Помимо официальных существует множество вариантов дистрибутивов,
разрабатываемых сообществом. Например
[Mythbuntu](http://www.mythbuntu.org/), [Ubuntu
ME](http://www.ubuntume.com/) и т.д.

## Почему я не могу зайти в систему под пользователем root?

По умолчанию возможность входа под root-ом отключена, вместо этого для
получения прав суперпользователя предлагается использовать sudo
(например *sudo ifconfig eth0 up* или *sudoedit
/etc/apt/sources.list*). Если вы привыкли выполнять административные
задачи под рутом, используйте *sudo -i*.

Подсказка: здесь sudo спрашивает пароль <u>Вашего</u> пользователя.

## Как «включить» аккаунт root?

В процессе установки Ubuntu пользователя, в отличие от других
дистрибутивов, не просят установить пароль root’а. Именно по
этой причине аккаунт недоступен. Впрочем особой нужды в возможности
залогиниться root'ом нет — все необходимые действия по настройке
системы можно выполнять через sudo. Но если уж так очень хочется,
то:

    sudo passwd root

В результате этой команды система предложит ввести новый пароль для
пользователя root. После этого данный аккаунт можно будет
полноценно использовать для входа в систему.

Чтобы «отключить» аккаунт root запустите следующую команду:

    sudo passwd -l root

Не забывайте, что
[работать](http://www.linux.org.ru/wiki/en/General#.D0.9F.D0.BE.D1.87.D0.B5.D0.BC.D1.83_.D0.B2.D1.81.D0.B5_.D0.B3.D0.BE.D0.B2.D0.BE.D1.80.D1.8F.D1.82.2C_.D1.87.D1.82.D0.BE_.D0.BD.D0.B5.D0.BB.D1.8C.D0.B7.D1.8F_.D1.81.D0.B8.D0.B4.D0.B5.D1.82.D1.8C_.D0.BF.D0.BE.D0.B4_root%27.D0.BE.D0.BC.3F)
под root’ом не рекомендуется.

## Интерфейс не полностью на русском\!

Если в инсталляторе выбрать русский язык, то при наличии Интернета во
время установки будут скачаны и установлены все необходимые пакеты с
русской локализацией. В ином случае придется вручную установить
мета-пакет language-support-ru либо поставить соответствующие
галочки в System -\> Administration -\> Language Support.

## Где взять недостающие аудио и видео кодеки, flash, java и т.д.?

Надо установить мета-пакет ubuntu-restricted-extras
(kubuntu-restricted-extras для KUbuntu, xubuntu-restricted-extras для
XUbuntu). Дополнительно об этом почитать можно
[тут](https://help.ubuntu.com/community/RestrictedFormats).

Начиная с версии 9.04 появилась возможность автоматической доустановки
кодеков — при необходимости система автоматически скачает и установит
необходимые для воспроизведения пакеты.

Если же на компьютере с Ubuntu нет доступа в интернет, то спасением
будет вот
[этот](http://hacktolive.org/wiki/Ubuntu-restricted-extras_offline_installer)
скрипт.

## Как вернуть кнопки сворачивания и закрытия обратно в правую часть окна?

Начиная с Ubuntu 10.04 кнопки сворачивания и закрытия перемещены в левую
часть окна, что часто неудобно для бывалых пользователей Linux.

Для придания кнопкам привычного вида, достаточно запустить gconf-editor
и в иерархии /apps/metacity/general установить переменной
«button\_layout» значение "menu:minimize,maximize,close".

Или из консоли:

    user@linux$ gconftool-2 --type string --set /apps/metacity/general/button_layout "menu:minimize,maximize,close"

## Что делать если после установки проприетарных драйверов «ломается» сплешскрин при загрузке системы?

В первую очередь ставим недостающие пакеты — <i>sudo apt-get install
v86d hwinfo</i>. Далее смотрим какие разрешение и с каким количеством
цветов поддерживается видеокарта — <i>sudo hwinfo --framebuffer</i>.
Потом редактируем несколько конфигов.

/etc/default/grub:

    GRUB_CMDLINE_LINUX_DEFAULT="quiet splash video=uvesafb:mode_option=1280x1024-16"
    ...
    GRUB_GFXMODE=1280х1024

В /etc/initramfs-tools/modules добавляем:

    uvesafb mode_option=1280x1024-16 

В /etc/initramfs-tools/conf.d/splash:

    FRAMEBUFFER=y 

После чего запускаем <i>sudo update-grub2</i> и <i>sudo update-initramfs
-u</i>

При наличии нескольких мониторов после перезагрузки возможно
переопределение приоритетов видеовыходов, по крайней мере в
Kubuntu 13.04. Решение — в настройках используемой DE самостоятельно
указать порядок расположения мониторов.

## Как добавить ключ для репозитория PPA?

Если система ругается, что не найден ключ репозитория:

> W: Ошибка: <http://ppa.launchpad.net> jaunty Release: Следующие
> подписи не могут быть проверены, так как недоступен открытый
> ключ: NO\_PUBKEY 5A9BF3BA4E5E17B5

Это легко исправить одной командой:

    user@linux$ sudo apt-key adv --recv-keys --keyserver keyserver.ubuntu.com 5A9BF3BA4E5E17B5

Начиная с Ubuntu 9.10 можно автоматизировать процесс добавления нужных
репозиториев и ключей PPA. Вот небольшой пример добавления репозитория
с браузером Chromium:

    user@linux$ sudo add-apt-repository ppa:chromium-daily

## Как полностью удалить GNOME/KDE/Xfce?

  - <http://www.psychocats.net/ubuntu/puregnome>
  - <http://www.psychocats.net/ubuntu/purekde>
  - <http://www.psychocats.net/ubuntu/purexfce>

## Статьи в этой wiki

  - [Установка Ubuntu через другой GNU/Linux
    (debootstrap)](Установка_Ubuntu_через_другой_GNU/Linux_\(debootstrap\) "wikilink")

## Ссылки

\[[http://ubuntuforums.org\] ](http://ubuntuforums.org%5D%C2%A0);—
официальные форумы.

\[[http://forum.ubuntu.ru\] ](http://forum.ubuntu.ru%5D%C2%A0);—
русскоязычный форум.

\[[http://ubuntuguide.org\] ](http://ubuntuguide.org%5D%C2%A0);—
подробный FAQ, в том числе и на русском языке

[Category:Дистрибутивы](Category:Дистрибутивы "wikilink")
