**AJAX** - Асинхронный JavaScript и XML (Asynchronous JavaScript and
XML). Технология, позволяющая запрашивать с веб-сайтов (HTTP GET) или
отправлять веб-сайтам (HTTP POST) данные в виде XML. Естественно, со
своей стороны веб-сервер должен уметь отдавать или принимать
запрошенный XML. По своей сути, AJAX реализован как
стандартизированный JavaScript объект *XMLHttpRequest*,
который поддерживается большинством браузеров, и успешно
используется в XUL.

AJAX позволяет создавать интерактивные страницы в интернете, и
чрезвычайно популярен при реализации новомодных течений типа
[Web 2.0](http://ru.wikipedia.org/wiki/Web_2.0).

В XUL AJAX позволяет запрашивать какие-нибудь интересные для
пользователя данные с других сайтов, например галерею из
[Google Picasa](http://picasaweb.google.com), или погоду с [Yahoo
Weather](http://weather.yahoo.com/). Далее полученный XML преобразуется
десериализатором в JavaScript объект, и выводится на экран в какой-либо
форме.

## Спецификация

Объект XMLHttpRequest описан следующим образом:

    <nowiki>
    interface XMLHttpRequest
    {
      // обработчик событий
      attribute EventListener onreadystatechange;

      // константы состояний
      const unsigned short UNSENT = 0;
      const unsigned short OPENED = 1;
      const unsigned short HEADERS_RECEIVED = 2;
      const unsigned short LOADING = 3;
      const unsigned short DONE = 4;
      readonly attribute unsigned short readyState;

      // формирование запросов
      void open(in DOMString method, in DOMString url);
      void open(in DOMString method, in DOMString url, in boolean async);
      void open(in DOMString method, in DOMString url,
                in boolean async, in DOMString user);
      void open(in DOMString method, in DOMString url, in boolean async,
                in DOMString user, in DOMString password);
      void setRequestHeader(in DOMString header, in DOMString value);

      // запрос
      void send();
      void send(in DOMString data);
      void send(in Document data);
      void abort();

      // ответ
      DOMString getAllResponseHeaders();
      DOMString getResponseHeader(in DOMString header);

      // атрибуты ответа
      readonly attribute DOMString responseText;
      readonly attribute Document responseXML;
      readonly attribute unsigned short status;
      readonly attribute DOMString statusText;
    };
    </nowiki>

Методы *open(...)* служат для формирования запросов. Методы *send(...)*
- для отправки сформированных запросов серверу. Обработчик событий
*onreadystatechange* служит для того, чтобы узнать состояние нашего
запроса, и, возможно, получить ответ от сервера (в асинхронных
запросах).

Запрос можно отослать синхронно, или асинхронно. Если запрос отсылается
синхронно, то *send()* блокируется и освобождается только после
выполнения операции (или ошибки). Сразу после выполнения
*send()* в таком случае можно проверить статус ошибки и получить ответ.
Если же запрос отсылается асинхронно, то *send()* не блокируется, и
выполнение программы продолжается. В этом случае необходимо
установить обработчик *onreadystatechange*, чтобы узнать, когда
операция завершилась и проверить статус и состояние операции.

Состояние запроса можно узнать по значение атрибута *readyState*,
которое будет равно одной из констант состояний. Статус запроса
(HTTP статус) можно узнать по атрибуту *status*. Если запрос прошёл
успешно, *status* будет равен *200* (HTTP Ok).

С помощью *XMLHttpRequest* можно читать локальные файлы, в этом случае
запрашиваемый URL должен начинаться с локального протокола
*<file://>*. Индикатор успешного выполнения в таком случае будет
значение атрибута *status*, равное *0*.

*send()* в дополнение ко всему может отсылать данные, переданные ей как
аргумент. Для запросов типа *GET* отсылать никаких данных не нужно,
поэтому в качестве аргумента передаём *null* или *''* (пустую
строку). Для запросов типа *POST* в качестве аргумента передаём
данные, которые нужно отослать серверу.

## Использование

Пример использования асинхронного запроса:

    <nowiki>
    // создаём новый объект XMLHttpRequest
    var mXML = new XMLHttpRequest();

    // функция запроса данных
    function check(url)
    {
        try
        {

        dump("Checking " + url + "...\n");

        // установка обработчика событий
        mXML.onreadystatechange = function()
        {
            if(this.readyState == 4)
            {
                // 200 для удалённых файлов, 0 для локальных
                if(this.status == 200 || this.status == 0)
                    dosomething(this.responseText);
            }
        };

        // задаём параметры запроса:
        //
        // 'GET' - HTTP GET, будем скачивать данные
        // url - URL для скачивания
        // true - скачиваем асинхронно
        mXML.open('GET', url, true);

        // можем принимать XML
        mXML.setRequestHeader('Accept', 'text/xml');

        // интерпретируем полученный XML как текст
        // (возможно, в этих полях не будет необходимости)
        mXML.overrideMimeType('text/plain');
        mXML.setRequestHeader('Content-Type', 'text/plain');

        // отправляем запрос, результат
        // придёт в обработчик onreadystatechange
        mXML.send(null);

        }
        catch(e) { alert("Error:\n\n" + e); }
    }
    </nowiki>

Тоже самое, только синхронно:

    <nowiki>
    // создаём новый объект XMLHttpRequest
    var mXML = new XMLHttpRequest();

    function check(url)
    {
        try
        {

        dump("Checking " + url + "...\n");

        // задаём параметры запроса:
        //
        // 'GET' - HTTP GET, будем скачивать данные
        // url - URL для скачивания
        // false - скачиваем синхронно
        mXML.open('GET', url, false);

        mXML.setRequestHeader('Accept', 'text/xml');
        mXML.overrideMimeType('text/plain');
        mXML.setRequestHeader('Content-Type', 'text/plain');

        // отправляем запрос, тут же ждём результат
        mXML.send(null);

        // результат получен
        if(mXML.status == 200 || mXML.status == 0)
            dosomething(mXML.responseText);

        }
        catch(e) { alert("Error:\n\n" + e); }
    }
    </nowiki>

## Пример

В проекте *example6* есть полноценный пример как скачивать локальные или
удалённые файлы.

## Ссылки

  - ![Исходники example6 (переименуйте в .tar.bz2)](example6.tar.bz2
    "Исходники example6 (переименуйте в .tar.bz2)")
  - [Спецификация
    XMLHttpRequest](http://www.w3.org/TR/2007/WD-XMLHttpRequest-20071026/)
  - [XMLHttpRequest от
    Mozilla](http://developer.mozilla.org/en/docs/XMLHttpRequest)
  - [(Де)сериализатор
    XML](http://www.kawa.net/works/js/xml/objtree-e.html)

[Category:XUL](Category:XUL)