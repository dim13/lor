**LightScribe** - технология записи оптических носителей информации (CD,
DVD, ...), целью которой является получение какого-либо изображения на
лицевой стороне носителя. Разработана Hewlett-Packard. Програмное
обеспечение разрабатывается и поддерживается французской фирмой La
Cie.

Под понятием "печати" дисков внутри этой статьи будет подразумеваться
лазерная обработка лицевой стороны носителя по технологии
LightScribe.

## Немного о технологии

Для пользования всеми прелестями этой технологии вам необходим
записывающий привод с поддержкой LightScribe, а также особые
болванки, лицевая сторона которых покрыта специальным химическим
покрытием, восприимчивым к излучению 780нм-лазера. Наверно, уже
стало понятно, что лазером можно выжечь на поверхности диска четкую
картинку состоящую из отенков серого цвета. Вы выбираете только цвет
лицевого покрытия при покупке - иначе говоря, "задний фон".

Вы кладёте оптический носитель в привод "вверх ногами", с помощью
специальной программы отправляете картинку на привод и ждёте
некоторое время (от 15 до 21 минуты при раскраске всей площади
диска) и получаете носитель, который не придётся оформлять или
подписывать самостоятельно.

*Каждый носитель имеет постоянную точку отсчета*, поэтому одно и тоже
изображение может быть нанесено на диск несколько раз без сдвигов или
неточностей. С другой стороны, вы можете закрасить только интересующую
Вас область на поверхности диска, не боясь промахнуться и испортить
другие, ранее напечатанные области.

## Инструменты

Все программы, используемые для записи LightScribe дисков -
проприетарные, либо используют проприетарные компоненты.
Будьте внимательны, La Cie выпускает сборки библиотек, утилит и SDK
**только под x86-архитектуру** (т.е. 32-битные), поэтому владельцам
x86\_64/EM64T нужно иметь на вооружении рабочий 32-битный gtk.

  - Ядро системы LightScribe - библиотека **liblightscribe.so**. Её
    утилизируют все нижеследующие программы. Качать
    [тут](http://www.lightscribe.com/downloadSection/linux/index.aspx?id=814).
    Если Вы разработчик, вас может заинтересовать [SDK к этой
    библиотеке](http://www.lightscribe.com/downloadSection/linux/index.aspx?id=816).
    Как уже стало ясно, для пользования LightScribe не требуется
    собирать какие-либо ядерные модули или патчить ядро.
  - **4L** - простенькая программа для вывода на "печать" любых
    изображений (или их областей) в любом масштабе. Искать
    [тут](http://www.lacie.com/support/drivers/driver.htm?id=10094).
    В комплект входят:
      - **4L-cli** - консольная утилита для печати.
      - **4L-gui** - графическая, на gtk.
  - **SimpleLabeler** - как ясно из названия, это простая программа для
    простого выжигания этикеток. Используется для создания и печати
    подписей, узоров и других примитивных видов оформления
    носителей. Найти её можно
    [здесь](http://www.lightscribe.com/downloadSection/linux/index.aspx?id=815).
  - **K3b** - заслуженно известная программа для записи всех видов
    носителей в Linux. Фирма La Cie является спонсором разработки
    k3b, поэтому здесь есть поддержка и LightScribe. Однако, реализована
    она через проприетарную библиотеку, поэтому только x86-сборка
    наделена этой способностью. [Сайт
    программы](http://www.k3b.org).

Все лицензионные аспекты использования и распространения ПО указаны по
соответствующим закачке ссылкам.

## Печать дисков

Для успешной работы утилит 4L и SimpleLabeler требуются повышенные
привилегии и GTK(x86).

    sudo chmod +s /usr/4L/4L-gui /opt/lightscribeApplications/SimpleLabeler

С работой утилит разобраться совсем не трудно, однако следует помнить,
что есть два режима записи (контрастности):

  - **Normal** - получается светлая, бледная и нечеткая размазня. Полная
    печать диска занимает порядка *15 минут*.
  - **Best** - довольно четкая и детализованная картина получается
    спустя *21 минуту*. Картинка "впечатывается" более глубоко и
    поэтому имеет бОльшую устойчивость к старению, в отличие от
    Normal.

Режим контрастности вы можете выбрать непосредственно перед началом
печати. Режим по умолчанию задается этой консольной утилитой,
запущеной из под root:

    /usr/lib/lightscribe/elcu.sh

Производитель гарантирует, что изображение будет держаться на носителе
без выцветания в комнатных условиях минимум *два года*. Повторюсь,
печать диска не одноразовая в отличие от записи, и уже напечатанное
изображение в будущем можно будет допечатывать, уточнять или портить
самостоятельно неограниченное число раз.

## Цены

Приводы с поддержкой LightScribe не является редкостью в ноутбуках от HP
и не создают заметного увеличения цены. Сами диски, однако, стоят раза в
два дороже обычных DVD+R (сравнение в рамках одного производителя
дисков). Существуют разные цвета дисков, однако все они
перечислены в спецификациях HP.

## Ссылки

  - [Подробне о технологии в википедии и с иллюстрациями (на английском
    языке)](http://en.wikipedia.org/wiki/LightScribe)
  - [1](http://www.lightscribe.com/)
  - [LinuxDownload-секция на
    LightScribe.com](http://www.lightscribe.com/downloadSection/linux/index.aspx)
  - [Галерея на
    lightscribe.com](http://www.lightscribe.com/ideas/labelgallery.aspx?id=219)
  - [Галерея обложек.](http://www.lightscribecovers.com/) Для закачки
    файлов-обложек требуется бесплатная
    [регистрация](http://www.lightscribecovers.com/register.php).

[Category:Hardware](Category:Hardware "wikilink")