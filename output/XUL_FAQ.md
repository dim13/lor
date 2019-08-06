## Как сделать окно на весь экран?

В 1.8 версии движка - никак. Можно установить размеры окна равными
размеру экрана, но даже в этом случае при запуске в xulrunner,
окно коллапсирует к размерам по умолчанию. Поэтому в 1.8 версии
нужно отнимать 1 от размера экрана (тем не менее, на новых версия�
движка 1.8, стоит всё же проверить способ и без вычитания единицы,
возможно Mozilla наконец-то это починила):

    <nowiki>
        window.moveTo(0, 0);
        window.resizeTo(screen.width-1, screen.height-1);
    </nowiki>

В версии 1.9 всё намного лучше, но тоже не идеально, приходится
пользоваться таймером:

    <nowiki>
        setTimeout('window.fullScreen = true;', 1);
    </nowiki>

## XMLHttpRequest не хочет ходить в сеть, раньше ходил\!

Сеть в 1.9 по умолчанию находится в режиме «offline» (в 1.8 - в
«online»). Поэтому всё что связано с сетью, работать не будет. Нужно
принудительно включить режим "online", используя nsIIOService2:

    <nowiki>
        const mIOService = Components.classes["@mozilla.org/network/io-service;1"]
                                          .getService(Components.interfaces.nsIIOService2);

        try
        {
            mIOService.offline = false;
        }
        catch(ex) { }
    </nowiki>

## Как сделать UDP Server Socket в JavaScript?

[Никак](https://bugzilla.mozilla.org/show_bug.cgi?id=433102). Если
потратить время, можно попытаться написать XPCOM компонент, либо
[использовать
Python](http://aspn.activestate.com/ASPN/Downloads/Komodo/PyXPCOM/).

## Как получить путь к профайлу пользователя?

    <nowiki>
    var path = "";
    const mDirService = new Components.Constructor("@mozilla.org/file/directory_service;1","nsIProperties");

    try
    {
        path = (new mDirService()).get("ProfD", Components.interfaces.nsIFile).path;
    }
    catch(ex)
    {
        alert("Error: \n\n" + ex);
        return; // если внутри функции
    }

    if(path.search(/\\/) != -1)
        path += "\\";
    else
        path += "/";

    alert("Profile path: " + path);
    </nowiki>

## Как записать локальный файл?

    <nowiki>
    function save(filepath, data)
    {
        var mLocalFile = Components.classes["@mozilla.org/file/local;1"]
                            .createInstance(Components.interfaces.nsILocalFile);

        mLocalFile.initWithPath(filepath);

        if(!mLocalFile.exists())
            mLocalFile.create(Components.interfaces.nsIFile.NORMAL_FILE_TYPE, 420);

        var mFileOutputStream = Components.classes["@mozilla.org/network/file-output-stream;1"]
                                    .createInstance( Components.interfaces.nsIFileOutputStream);

        /*
         * Флаги открытия, из исходников XULRunner
         *
         *  #define PR_RDONLY       0x01
         *  #define PR_WRONLY       0x02
         *  #define PR_RDWR         0x04
         *  #define PR_CREATE_FILE  0x08
         *  #define PR_APPEND       0x10
         *  #define PR_TRUNCATE     0x20
         *  #define PR_SYNC         0x40
         *  #define PR_EXCL         0x80
         *
         * Режимы открытия, сходны с правами файлов UNIX, и могут быть
         * проигнорированы на других платформах.
         *
         *   00400   Read by owner
         *   00200   Write by owner
         *   00100   Execute by owner
         *   00040   Read by group
         *   00020   Write by group
         *   00010   Execute by group
         *   00004   Read by others
         *   00002   Write by others
         *   00001   Execute by others
         *
         */
        mFileOutputStream.init(mLocalFile, 0x04 | 0x08 | 0x20, 420, 0);

        mFileOutputStream.write(data, data.length);

        mFileOutputStream.close();
    }
    </nowiki>

## Как прочитать локальный файл?

    <nowiki>
    function read(filepath)
    {
        var mLocalFile = Components.classes["@mozilla.org/file/local;1"]
                             .createInstance(Components.interfaces.nsILocalFile);

        mLocalFile.initWithPath(filepath);

        if(!mLocalFile.exists())
            return "";

        var mFileInputStream = Components.classes["@mozilla.org/network/file-input-stream;1"]
                                   .createInstance(Components.interfaces.nsIFileInputStream);

        // 0x01 - открываем только на чтение
        // 0004 - чтение для всех (аналогично правам файлов в UNIX)
        // null - флаги поведения, не используются
        mFileInputStream.init(mLocalFile, 0x01, 00004, null);

        // для бинарных файлов используем nsIBinaryInputStream
        var mInputStream = Components.classes["@mozilla.org/scriptableinputstream;1"]
                               .createInstance(Components.interfaces.nsIScriptableInputStream);

        mInputStream.init(mFileInputStream);

        return mInputStream.read(mInputStream.available());
    }
    </nowiki>

## Как распарсить формат календаря iCal?

Использовать
[iCalReader](http://code.google.com/p/kisaragi/source/browse/trunk/src/iCalReader.js?r=13).

## Как сериализовать/десериализовать XML?

Имхо, лучший \[де\]сериализатор -
[1](http://www.kawa.net/works/js/xml/objtree-e.html).

## Как написать клиента погодного сервиса?

  - Изучаем документацию на [Yahoo](http://developer.yahoo.com/weather/)
  - Берём базы международных идентификаторов на
    [Intellicast](http://www.intellicast.com/FFC/SupportFiles)
  - Используем XML десериализатор для получения RSS данных о погоде и
    превращения их в JavaScript объект
  - Выводим полученные данные на экран

## Как отключить кеширование в XUL?

Директивами в Вашем скрипте настроек проекта
(defaults/preferences/prefs.js):

    <nowiki>
    pref("nglayout.debug.disable_xul_cache", true);
    pref("nglayout.debug.disable_xul_fastload", true);
    </nowiki>

## Как выводить отладочные сообщения в шелл?

Директивой в Вашем скрипте настроек проекта
(defaults/preferences/prefs.js) разрешить использование dump():

    <nowiki>
    pref("browser.dom.window.dump.enabled", true);
    </nowiki>

Теперь в коде можно писать

    <nowiki>
    dump("Values: " + a + ", " + b + "\n");
    </nowiki>

## Как получить favicon из загруженной страницы в <browser>?

Favicon обычно хранится в html документе как

    <nowiki>
    <link rel="shortcut icon" href="/images/favicon.png" />
    </nowiki>

или

    <nowiki>
    <link rel="icon" href="/images/favicon.png" />
    </nowiki>

Если таких элементов нету, то подразумевается, что favicon по умолчанию
расположен на HOST/favicon.ico, где HOST - доменное имя, например
«**google.com**».

Поэтому алгоритм примерно такой:

    <nowiki>
    function getFavIcon()
    {
        var rel;
        var favicon = null;
        var firstIconIndex = -1;

        // получаем массив тегов <link>
        var links = content.document.getElementsByTagName("link");

        for(var i in links)
        {
            if(typeof(links[i].rel) != "undefined")
            {
                rel = links[i].rel.toLowerCase();

                // Отлично, сразу нашли favicon.
                if(rel == "shortcut icon")
                {
                    favicon = links[i].href;
                    break;
                }
                // Возможно, favicon передаётся как rel="icon".
                // При этом не прекращаем поиск.
                else if(rel == "icon" && firstIconIndex == -1)
                    firstIconIndex = i;
            }
        }

        // Не нашли подходящих тегов с rel="shortcut icon"
        if(!favicon)
        {
            // Возможно, нашли с атрибутом rel="icon"
            if(firstIconIndex != -1)
                favicon = links[firstIconIndex].href;
            else
            {
                var uri = browser.currentURI;

                // Ничего подходящего не нашли.
                // Пробуем загрузить напрямую из HOST/favicon.ico
                try
                {
                    favicon = uri.scheme + "://" + uri.host + "/favicon.ico";
                }
                catch(e) { favicon = ""; }
            }
        }

        alert(favicon);
    }
    </nowiki>

Полный XUL проект для демонстрации получения favicon лежит
![здесь](example5.tar.bz2 "здесь") (переименуйте в .tar.bz2). После
запуска примера, дождитесь загрузки адреса в браузере, и нажмите кнопку
вверху окна. Полученный URL favicon'а передаётся в элемент *image*,
который загружает и отображает иконку. Сам элемент *image*
расположен слева внизу окна.

## Как внедрить SVG?

SVG картинка в качестве *src* атрибута элементов пока не поддерживается.
Можно только внедрить SVG напрямую, вот например несколько ссылок по
этому поводу:

  - [2](http://developer.mozilla.org/en/Code_snippets/Embedding_SVG)
  - [3](http://weblogs.mozillazine.org/weirdal/archives/015917.html)
  - [4](http://www.josephguhlin.com/blog/archives/2006/04/28/playing-around-with-xul-and-svg/)
  - [5](http://kb.mozillazine.org/Using_SVG_with_XBL_in_XUL)

## Как создать iconview?

С помощью HTML элемента *div*. Не забываем, что для работы HTML
элементов нужно импортировать соответствующее пространство имён
(*xmlns*):

    <nowiki>
    <?xml version="1.0"?>
    <?xml-stylesheet href="chrome://global/skin/" type="text/css"?>

    <window
        title="Example"
        width="320"
        height="200"
        ns="http://www.mozilla.org/keymaster/gatekeeper/there.is.only.xul"
        xmlns:html="http://www.w3.org/1999/xhtml">

        <html:div>
            <button label="1" />
            <button label="2" />
            <button label="3" />
            <button label="4" />
            <button label="5" />
            <button label="6" />
            <button label="7" />
            <button label="8" />
        </html:div>

    </window>
    </nowiki>

Чтобы динамически создать объект *html:div*, необходимо также указать
методу *document.createElementNS()* необходимое пространство имён:

    <nowiki>
    const XHTML_NS = "http://www.w3.org/1999/xhtml";
    var div = document.createElementNS(XHTML_NS, "html:div");
    </nowiki>

## Что такое Firefox?

Firefox - это расширяемый плагинами браузер, написанный на XUL, и
позволяющий также выполнять локальное и удалённое XUL содержимое
(см. [Введение](XUL_Введение "wikilink")). Плагины представляют собой
устанавливаемое XUL содержимое.

## Есть ли русскоязычный форум по XUL?

[Есть](http://xpoint.ru/forums/programming/XUL/forum.xhtml).

[Category:XUL](Category:XUL "wikilink")