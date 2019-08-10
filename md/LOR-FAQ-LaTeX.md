LaTeX - распространённая система вёрстки в основном предназначенная для
создания научной литературы с большим количеством формул, также латех
зачастую используется линуксоидами при написании разнообразных
дипломов. Произносится - *лате**х***

## Что можно почитать новичку о latex на русском?

Есть замечательная книга - С.М.Львовский. Набор и верстка в системе
LaTeX. Последнее на данный момент, 3-е издание включая исходные тексты
книги доступно для скачивания
[здесь](http://www.mccme.ru/free-books/llang/newllang.pdf).

Другая не менее замечательная книга - [Не очень краткое введение в
LaTeX2ε](http://zelmanov.ptep-online.com/ctan/lshort_russian.pdf).

И третья - [Компьютерная типография
LaTeX](http://ctan.org/tex-archive/info/russian/Computer_Typesetting_Using_LaTeX)

## Как создавать русскоязычные документы в latex?

Всё давно уже работает "из коробки". Ставим texlive и используем
следующий шаблон:

    \documentclass[a4paper,12pt]{article}
    \usepackage{mathtext}
    \usepackage[utf8x]{inputenc}
    \usepackage[russian]{babel}
    \usepackage[T2A]{fontenc}
    \usepackage{amsmath,amssymb,amsthm,amscd,amsfonts}
    
    \begin{document}
    
    
    \end{document}

Предполагается что файл пишется в юникоде.

В Debian, Ubuntu и их клонах стоит помнить, что поддержка кириллицы в
TeXLive идёт отдельным пакетом texlive-lang-cyrillic, который по
умолчанию не ставится.

## Какие стили использовать для курсовой/диплома/реферата в latex?

Основной стиль [dissert](http://ppg.ice.ru/files/59553/dissert.tgz), для
библиографии
[gost780u.bst](http://www.inp.nsk.su/~baldin/Cyrillic-HOWTO-russian/ch08s02.html)
(начиная с TeXLive 2012 gost2003 или gost2008).

Также можно использовать стиль eskdx, который входит в стандартную
поставку TeX Live.

Шаблон титульной страницы:
[1](http://www.linux.org.ru/jump-message.jsp?msgid=1305696&cid=1305748)

## Почему русские буквы в pdf-файлах, созданных в LaTeX, выглядят так страшно на экране

Шрифты стоят такие. Для создания красиво выглядящих документов лучше
поставить cm-super. Также есть другой вариант шрифтов, pscyr, но их
лицензионность весьма сомнительна.

## Русский текст из pdf копируется крякозябрами, как исправить?

В этом случае может помочь использование пакета cmap до вызова пакетов
fontenc и inputenc, и компиляция документа pdflatex. Пример:

    \documentclass[a4paper,titlepage,10pt]{article}
    \usepackage{cmap}
    \usepackage[T2A]{fontenc}
    \usepackage[utf8x]{inputenc}
    \usepackage[russian]{babel}

## Как сделать ссылки на литературу вида \[1,3-5\] вместо \[1,3,4,5\]?

    \usepackage{cite}

## А есть специальные редакторы для LaTeX?

  - [Emacs](http://www.gnu.org/software/emacs/) +
    [AUCTeX](http://www.gnu.org/software/auctex/)
  - [Kile](http://kile.sourceforge.net/), в KDE
  - [Texmaker](http://www.xm1math.net/texmaker/)

кроме того, есть WYSIWYG-редакторы, использующие TeX в качестве движка:

  - [LyX](http://www.lyx.org/)
  - [TeXmacs](http://www.texmacs.org/)

## Чем в latex нарисовать схему базы данных?

С этой задачей вполне справится пакет tikz, или metauml. Примеры
применения tikz, вместе с исходными текстами примеров можно
найти [здесь](http://www.texample.net/tikz/examples/).

## Можно ли вставить графику pdf как отдельную страницу игнорируя отступы по краям?

Можно, пакет pdfpages, команда \\includepdf.

## Заголовки секций, субсекций в оглавлении не влазят в строку, что делать?

Пример: \\section{ПАРАГЕНЕЗИС И ТИПОМОРФИЗМ МИНЕРАЛОВ В ПЕГМАТИТАХ НА
ПРИМЕРЕ ТИПОМОРФНЫХ АССОЦИАЦИЙ} на выходе обрезает до НА ПРИМЕРЕ...

Если сильно надо написать такое длиннющее название раздела, используйте
альтернативное сокращенное имя для оглавления и колонтитулов:

    \documentclass[a4paper,12pt,twoside]{article} 
    \usepackage{/Data/documents/LaTeX_Templates/ed} 
    \begin{document} 
    \tableofcontents 
    \newpage 
    \section[Короткое название]{Тупое длиннющее название, которое ни при каких условиях не влезет ни в оглавление, ни в колонтитул} 
    \newpage 
    Текст 
    \end{document} 

Также по сведениям товарища St\_MPA3b переносы строк в оглавлении
нормально работают в стиле disser.

## Как мне сделать текст повёрнутым на 90° (вертикальным)?

Для этого можно использовать пакет rotating. Пример:

    \begin{sideways} Вертикальный текст \end{sideways}

[Источник](http://www.linux.org.ru/forum/general/4578155)

## Как отобразить байтовые поля, например структуру IP-пакета?

Для этого можно использовать пакет bytefield. Пример:

    \usepackage{bytefield} 
     
    \begin{document} 
    \begin{bytefield}{32} 
            \bitheader{0,15,16,31}\\ 
            \wordgroupr{Заголовок} 
            \bitbox{16}{Destport}\bitbox{16}{Src port}\\ 
            \bitbox{16}{size}\bitbox{16}{CRC}\\ 
            \endwordgroupr\\ 
            \wordbox[lrt]{1}{Data}\\ 
            \skippedwords\\ 
            \wordbox[lrb]{1}{up to 65{.}527 bytes} 
    \end{bytefield} 

[Источник](http://www.linux.org.ru/forum/development/4751338)

## Как заставить TeX расставлять переносы в узких столбцах таблиц?

Добавляем в преамбуле:

    \usepackage{array}
    \usepackage{ragged2e}
    \newcolumntype{x}[1]{%
    >{\RaggedRight\hspace{0pt}}p{#1}}%

Далее вместо p{ширина} используем x{ширина} Итого: переносы есть, как и
выравнивание по левому краю.

[Источник](http://www.linux.org.ru/forum/general/9128352?cid=9134170)

[Category:Полезные советы](Category:Полезные_советы "wikilink")
