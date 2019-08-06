# Немного о технологии

Технология nVidia Optimus, которая используется в ноутбуках, стала
популярной в последнее время. Главная проблема - то, что поддержки
для Linux со стороны nVidia нет и не будет.

## Как работает

В ноутбуке с Optimus сразу 2 видеокарты: одна слабая и встроенная в CPU
от intel и одна мощная от nVidia. Идея состоит в том, что при слабы�
нагрузках на видеокарту используется только intel, а при сложны�
задачах включается nVidia. Польза в том, что время работы ноутбука
без подзарядки заметно увеличивается. Есть 2 типа реализации: в одной,
вы можете менять карточки в BIOS, а в другой nVidia не имеет выхода к
экрану, а значит вам обычные драйвера от nVidia не подойдут. Для того,
чтобы заставить это всё работать, был создан проект под названием
Bumblebee.

# Bumblebee

Это демон bumblebeed,клиент к нему, модуль ядра bbswitch и драйвера,
реализующий технологию Optimus для Linux. Правда пока не полностью.
Сейчас есть поддержка запуска приложения с карточки nVidia
самостоятельно (естественно все остальное работает на
intel) и автоматическое включение/выключение карточки nvidia, при
запуске приложений. Перед установкой следует настроить карточку
от intel.

## Где достать?

### Форки и адаптации

Если вы используете один из дистрибутивов ниже, то для вас процесс
установки другой.

#### CentOS

[Bumblebee](https://build.opensuse.org/package/show?project=home%3ABumblebee-Project%3ABumblebee&package=bumblebee),
[Bumblebee-develop](https://build.opensuse.org/package/show?project=home%3ABumblebee-Project%3ABumblebee-develop&package=bumblebee),
[Bumblebee-unstable](https://build.opensuse.org/package/show?project=home%3ABumblebee-Project%3ABumblebee-unstable&package=bumblebee).

#### Red Hat Enterprise Linux

[Bumblebee](https://build.opensuse.org/package/show?project=home%3ABumblebee-Project%3ABumblebee&package=bumblebee),
[Bumblebee-develop](https://build.opensuse.org/package/show?project=home%3ABumblebee-Project%3ABumblebee-develop&package=bumblebee),
[Bumblebee-unstable](https://build.opensuse.org/package/show?project=home%3ABumblebee-Project%3ABumblebee-unstable&package=bumblebee).

#### Fedora

[Bumblebee](https://build.opensuse.org/package/show?project=home%3ABumblebee-Project%3ABumblebee&package=bumblebee),
[Bumblebee-develop](https://build.opensuse.org/package/show?project=home%3ABumblebee-Project%3ABumblebee-develop&package=bumblebee),
[Bumblebee-unstable](https://build.opensuse.org/package/show?project=home%3ABumblebee-Project%3ABumblebee-unstable&package=bumblebee).

#### Mandriva

[Bumblebee](https://build.opensuse.org/package/show?project=home%3ABumblebee-Project%3ABumblebee&package=bumblebee),
[Bumblebee-develop](https://build.opensuse.org/package/show?project=home%3ABumblebee-Project%3ABumblebee-develop&package=bumblebee),
[Bumblebee-unstable](https://build.opensuse.org/package/show?project=home%3ABumblebee-Project%3ABumblebee-unstable&package=bumblebee).

#### openSUSE, SLE

[Bumblebee](https://build.opensuse.org/package/show?project=home%3ABumblebee-Project%3ABumblebee&package=bumblebee),
[Bumblebee-develop](https://build.opensuse.org/package/show?project=home%3ABumblebee-Project%3ABumblebee-develop&package=bumblebee),
[Bumblebee-unstable](https://build.opensuse.org/package/show?project=home%3ABumblebee-Project%3ABumblebee-unstable&package=bumblebee).

#### Debian и Linux Mint

Тоже имеет свой форк.
[debumblebee](https://github.com/z0rc/debumblebee). Тут вам понадобится
acpi_call (он туда не входит), я расскажу об этом позже.(Устарело.
Рекомендую использовать оригинальный bumblebeee.)

#### Ubuntu

Сразу стоит заметить, что если вы пользователь Ubuntu, то Bumblebee не
для вас. Был создан [форк под названием
ironhide](https://launchpad.net/~mj-casalogic/+archive/ironhide/),
который делает Optimus очень приятным и всё настроится за вас.
(Устарело. Рекомендую использовать оригинальный bumblebeee.)

#### Gentoo

Ebuild x11-misc/bumblebee доступен в официальном Portage, а также в
оверлея�
[bumblebee](https://github.com/Bumblebee-Project/bumblebee-gentoo) и
sabayon.

#### Arch

Имеет PKGBUILD для
[bbswitch](https://aur.archlinux.org/packages.php?ID=55799) и
[bumblebee](http://aur.archlinux.org/packages.php?ID=49469) в AUR. Все
остальное вытягивается зависимостями.

### Сам проект и как пользоваться

[Официальный Bumblebee](https://github.com/Bumblebee-Project/Bumblebee).
[VirtualGL](http://www.virtualgl.org/Downloads/VirtualGL), nVidia
driver, и модуль ядра
[bbswitch](https://github.com/Bumblebee-Project/bbswitch).

#### Установка

1.  Добавьте модуль ядра bbswitch.
2.  Соберите/Установите bumblebee.

## Конфигурация

Для Ubuntu и (Debian и Linux Mint возможно) он отличен.

1.  Добавьте пользователя в группу bumblebee :
        sudo usermod -a -G bumblebee USERNAME
2.  Отредактируйте bumblebee.conf под свои вкусы, выбрав нужный вам
    драйвер.

## Использование

1.  Запустите демон bumblebee.

Запускать программы с помощью nVidia командой:

```
 optirun program_name
```

И наслаждаться optimus на linux. :)

[Category: Hardware](Category:_Hardware "wikilink") [Category:
nVidia](Category:_nVidia "wikilink") [Category:
Bumblebee](Category:_Bumblebee "wikilink")