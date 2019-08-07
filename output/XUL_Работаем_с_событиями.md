С событиями многие уже хорошо знакомы. События - это некие уведомления
низлежащей системы тулкита (underlying system), которая информирует
нас о пользовательской активности, системной активности, событиях
таймера и т.д. События могут возникать спонтанно, или посылаться
целенаправлено самой программой.

## Термины

Обычно различают три момента, свойственные распространению событий в
приложении:

  - **Capturing** - захват события другими элементами на его пути к
    элементу-цели;
  - **Bubbling** - всплытие события вверх по дереву элементов, после
    того, как оно всё-таки добралось до элемента-цели;
  - **Canceling** - отмена реагирования элементом по умолчанию при
    приходе события.

Первые два пункта уже известны Qt программистам. При возникновении
события, низлежащая система создаёт объект, и выбрасывает его в
мир. Захват есть ни что иное, как перехват события на пути к элементу
другими обработчиками (аналогично фильтрам событий в Qt). Перехватчик
может запретить событию распространяться дальше. Как только событие
всё-таки пришло в пункт назначения, оно может быть обработано. После
обработки события тулкит направляет его вверх по дереву элементов,
если не указано другое. Это и есть всплытие. Мы можем явно указать
тулкиту, что не хотим, чтобы событие всплывало вверх (т.е.
утверждение, что *события не тонут*, неверно - тонут и ещё
как :) ).

Описанные здесь интерфейсы и типы событий относятся к DOM2.

При обработке событий в JavaScript используют первый способ,
*Capturing*.

## Базовый интерфейс элемента

Итак, рассмотрим базовый интерфейс для элемента, который хочет работать
с событиями:

    <nowiki>
    interface EventTarget
    {
      void       addEventListener(in DOMString type,
                                  in EventListener listener,
                                  in boolean useCapture);

      void       removeEventListener(in DOMString type,
                                     in EventListener listener,
                                     in boolean useCapture);

      boolean    dispatchEvent(in Event evt)
                                    raises(EventException);
    };
    </nowiki>

Этот интерфейс наследуется всеми XUL элементами, а это значит, что
каждый XUL элемент может подписываться на события. Если мы хотим
начать или перестать отслеживать события какого-то объекта, мы должны
вызвать именно его методы *addEventListener*, *removeEventListener* или
*dispatchEvent*.

Рассмотрим методы:

  - **addEventListener** - добавляет фильтр событий

В качестве параметров принимает тип события, обработчик, и булево
значение. Булево значение указывает, хочет ли данный обработчик
получать событие первым. Например:

    <nowiki>
    function func()
    {
        alert("We are here");
    }

    addEventListener("load", func, false);
    </nowiki>

Здесь мы установили на глобальный объект *window* фильтр на событие
*load* (загрузка документа) с именем *func*. Как только документ
загрузится, сгенерируется событие *load*, и выполнится наша
функция *func()*. Можно ставить два одинаковых обработчика на
одинаковые элементы, только если последний булев параметр
различается. В таком случае снимать обработчики нужно тоже по
отдельности (см. *removeEventListener*).

  - **removeEventListener** - снимает обработчик

Аргументы аналогичны.

  - **dispatchEvent** - создаёт событие, как будто оно возникло само

Объект, к которому применяется этот метод, является пунктом назначения
события (как уже говорилось выше). Например:

    <nowiki>
    function dispatch()
    {
        var event = document.createEvent('Events');
        event.initEvent('load', false, false);
        window.dispatchEvent(event);
    }

    function func()
    {
        alert("We are here");
    }

    addEventListener("load", func, false);

    setTimeout("dispatch();", 3000);
    </nowiki>

По таймеру через 3 секунды вызываем функцию *dispatch()*, которая
создаёт объект события, и передаёт его глобальному объекту
*window*. Т.к. на этом объекте уже стоит обработчик *func()*, то он и
вызывается повторно.

### Функция-обработчик

Функция-обработчик - это функция, которая автоматически вызывается при
возникновении события. Она может принимать в качестве аргумента
непосредственно сам объект события. Функция, объявленная без
параметров, не использует объект события. В этом примере мы
передаём объект события в функцию-обработчик:

    <nowiki>
    function func(evt)
    {
        alert(evt);
    }

    addEventListener("load", func, false);
    </nowiki>

Как только документ загрузится, сгенерируется событие *load*, и
выполнится наша функция *func()* с аргументом события, которое
мы и показываем в сообщении. Ещё один способ сделать это через
замыкание:

    <nowiki>
    function func(evt)
    {
        alert(evt);
    }

    // сложнее, но по сути тоже самое
    addEventListener("load", function(e) { func(e); }, false);
    </nowiki>

### Обработчики-атрибуты

XUL позволяет задавать обработчики для событий прямо внутри элементов в
XUL документе. В этом случае имя атрибута складывается из префикса
«*on*» и типа события (*load*, *mouseover*, *click* и т.д.). Например:

    <nowiki>
    <button label="click me!" onclick="alert('Clicked');" />
    </nowiki>

Здесь мы создали обработчик события *click*, который выводит на экран
сообщение *Clicked*. Другой пример:

    <nowiki>
    <button label="click me!" onmouseover="alert('Over');" />
    </nowiki>

Здесь мы создали обработчик события *mouseover*, который выводит на
экран сообщение *Over*, когда мышь визуально входит в пространство
над кнопкой.

### Объект event

Объект *event* создаётся тулкитом в обработчиках событий, заданных в
качестве атрибутов элементов. Например:

    <nowiki>
    <button label="click me!" onclick="alert(event);" />
    </nowiki>

Здесь при щелчке на кнопку мы выводим на экран сообщение с типом
полученного события (в данном случае это MouseEvent).

## Базовый интерфейс события

Базовый интерфейс события выглядит следующим образом:

    <nowiki>
    interface Event
    {
      const unsigned short      CAPTURING_PHASE        = 1;
      const unsigned short      AT_TARGET              = 2;
      const unsigned short      BUBBLING_PHASE         = 3;

      readonly attribute DOMString        type;
      readonly attribute EventTarget      target;
      readonly attribute EventTarget      currentTarget;
      readonly attribute unsigned short   eventPhase;
      readonly attribute boolean          bubbles;
      readonly attribute boolean          cancelable;
      readonly attribute DOMTimeStamp     timeStamp;

      void               stopPropagation();
      void               preventDefault();
      void               initEvent(in DOMString eventTypeArg,
                                   in boolean canBubbleArg,
                                   in boolean cancelableArg);
    };
    </nowiki>

Константы определяют в каком состоянии находится сейчас событие.
Атрибуты определяют тип события, конечную цель, текущую цель, и
др. Интерфейс содержит важные для нас методы:

  - **stopPropagation** - останавливает распространение события

Остановить распространение события можно как в случае с *Capturing*, так
и в случае с *Bubbling*. Вызов этого метода означает, что мы больше не
хотим, чтобы кто-то ещё обработал данное конкретное событие.

  - **preventDefault** - отмена действия по умолчанию

Если мы не хотим, чтобы тулкит выполнинл действие по умолчанию для
данного события, мы должны вызвать этот метод. Например, создадим
кнопку, которая визуально не нажимается, но событие нажатости всё-таки
обрабатывается:

    <nowiki>
    <script type="application/x-javascript">
    function func(evt)
    {
        evt.preventDefault();
    }
    </script>

    <button label="click me!"
            onmousedown="func(event);"
            onmouseup="func(event);"
            onclick="alert('Clicked');" />
    </nowiki>

Здесь мы запрещаем тулкиту обрабатывать действие по умолчанию для
событий *mousedown* и *mouseup*, которое обычно заключается в
отрисовке кнопки в нажатом и отжатом виде, и отрисовке фокуса.

  - **initEvent** - инициализирует ранее созданный объект события

В качестве аргументов принимает

\*\# тип события в виде строки («*click*», «*mousedown*», и т.д.);

\*\# булево значение, определяющее, может ли событие всплывать;

\*\# булево значение, определяющее, можно ли отменять действие по
умолчанию.

Например:

    <nowiki>
    // создали объект события
    var event = document.createEvent('Events');

    // инициализация
    event.initEvent('mouseover', true, true);

    // отсылка какому-либо объекту
    obj.dispatchEvent(event);
    </nowiki>

## Предопределённые события

DOM предоставляет пользователю несколько типов предопределённых событий,
каждое из которых имеет свои собственные свойства и методы. Рассмотрим
их подробнее.

### UIEvent (событие пользовательского интерфейса)

Имеет такое представление:

    <nowiki>
    interface UIEvent : Event
    {
      readonly attribute views::AbstractView  view;
      readonly attribute long             detail;

      void       initUIEvent(in DOMString typeArg,
                             in boolean canBubbleArg,
                             in boolean cancelableArg,
                             in views::AbstractView viewArg,
                             in long detailArg);
    };
    </nowiki>

Типы данного события могут быть «*DOMFocusIn*», «*DOMFocusOut*»,
«*DOMActivate*».

### MouseEvent (событие мыши)

Имеет такое представление:

    <nowiki>
    interface MouseEvent : UIEvent
    {
      readonly attribute long             screenX;
      readonly attribute long             screenY;
      readonly attribute long             clientX;
      readonly attribute long             clientY;
      readonly attribute boolean          ctrlKey;
      readonly attribute boolean          shiftKey;
      readonly attribute boolean          altKey;
      readonly attribute boolean          metaKey;
      readonly attribute unsigned short   button;
      readonly attribute EventTarget      relatedTarget;

      void       initMouseEvent(in DOMString typeArg,
                                in boolean canBubbleArg,
                                in boolean cancelableArg,
                                in views::AbstractView viewArg,
                                in long detailArg,
                                in long screenXArg,
                                in long screenYArg,
                                in long clientXArg,
                                in long clientYArg,
                                in boolean ctrlKeyArg,
                                in boolean altKeyArg,
                                in boolean shiftKeyArg,
                                in boolean metaKeyArg,
                                in unsigned short buttonArg,
                                in EventTarget relatedTargetArg);
    };
    </nowiki>

Как видно, событие мыши наследуется от события *UIEvent*. Событие мыши
включает в себя координаты курсора на экране, координаты курсора
относительно элемента, флаги, определяющие состояние клавиш
*Ctrl*, *Alt*, *Shift*, номер нажатой кнопки. Номер нажатой кнопки имеет
смысл только в обработчике события нажатия или отжатия кнопки
(«*mousedown*», «*mouseup*»). Например:

    <nowiki>
    <?xml version="1.0"?>
    <?xml-stylesheet href="chrome://global/skin/" type="text/css"?>

    <window
        title="Example"
        width="320"
        height="200"
        ns="http://www.mozilla.org/keymaster/gatekeeper/there.is.only.xul">

    <script type="application/x-javascript">

    function func(e)
    {
        // выводим в шелл сообщение с глобальными координатами
        // мыши и номером нажатой кнопки
        dump(e.screenX
                + ","
                + e.screenY
                + " Mouse button: "
                + e.button
                + "\n");
    }

    addEventListener("mousemove", func, false);
    addEventListener("mousedown", func, false);

    </script>

    </window>
    </nowiki>

Типы данного события могут быть «*click*», «*mousedown*», «*mouseup*»,
«*mouseover*», «*mousemove*», «*mouseout*».

### MutationEvent (событие изменения документа)

Имеет такое представление:

    <nowiki>
    interface MutationEvent : Event
    {

      // attrChangeType
      const unsigned short      MODIFICATION       = 1;
      const unsigned short      ADDITION           = 2;
      const unsigned short      REMOVAL            = 3;

      readonly attribute Node             relatedNode;
      readonly attribute DOMString        prevValue;
      readonly attribute DOMString        newValue;
      readonly attribute DOMString        attrName;
      readonly attribute unsigned short   attrChange;

      void       initMutationEvent(in DOMString typeArg,
                                   in boolean canBubbleArg,
                                   in boolean cancelableArg,
                                   in Node relatedNodeArg,
                                   in DOMString prevValueArg,
                                   in DOMString newValueArg,
                                   in DOMString attrNameArg,
                                   in unsigned short attrChangeArg);
    };
    </nowiki>

Константы определяют, что сейчас произошло - добавление, модификация или
удаление.

Свойства:

  - **attrChange** - тип модификации документа при событии
    *DOMAttrModified*

Может быть MODIFICATION, ADDITION, or REMOVAL.

  - **attrName** - имя модифицированного элемента при событии
    *DOMAttrModified*


  - **newValue** - новое значение модифицированного элемента при событии
    *DOMAttrModified* и *DOMCharDataModified*


  - **prevValue** - прошлое значение модифицированного элемента при
    событии *DOMAttrModified* и *DOMCharDataModified*


  - **relatedNode**

Типы данного события могут быть «*DOMSubtreeModified*»,
«*DOMNodeInserted*», «*DOMNodeRemoved*»,
«*DOMNodeRemovedFromDocument*», «*DOMNodeInsertedIntoDocument*»,
«*DOMAttrModified*», «*DOMCharacterDataModified*».

### HTML события

В XUL есть также возможность обрабатывать HTML события. Поддерживаются
следующие типы сообщений:

  - **load** - загрузка документа завершена;
  - **unload** - документ убран из окна или фрейма;
  - **abort** - неожиданная остановка загрузки, например остановка
    загрузки HTML страницы до того, как какое-то изображение
    успело загрузится;
  - **error** - произошла ошибка загрузки документа или выполнения
    скрипта;
  - **select** - пользователь выделил текст в текстовом поле ввода;
  - **change** - элемент потерял фокус, и при этом ещё изменился с
    момента получения фокуса (поле ввода например);
  - **submit** - отправка формы;
  - **reset** - возврат формы к значениям по умолчанию;
  - **focus** - элемент получил фокус;
  - **blur** - элемент потерял фокус;
  - **resize** - размер контейнера, в котором отображается документ,
    изменился;
  - **scroll** - прокрутка контейнера, в котором отображается документ.

Например:

    <nowiki>
    <?xml version="1.0"?>
    <?xml-stylesheet href="chrome://global/skin/" type="text/css"?>

    <window
        title="Example"
        width="320"
        height="200"
        ns="http://www.mozilla.org/keymaster/gatekeeper/there.is.only.xul">

    <button id="b" label="click me" />
    <textbox id="t" />

    <script type="application/x-javascript">

    function fFocus()
    {
        dump("Focus\n");
    }

    function fChange()
    {
        dump("Change\n");
    }

    // мы хотим получать события фокуса для кнопки
    document.getElementById("b").addEventListener("focus", fFocus, false);

    // мы также хотим получать события изменения для поля ввода,
    // не забываем, что оно приходит только после потери фокуса,
    // и только если содержимое поля ввода изменилось
    document.getElementById("t").addEventListener("change", fChange, false);

    </script>

    </window>
    </nowiki>

## Свои события

Помимо ручного создания и посылки стандартных типов событий, есть
возможность создавать свои типы событий, и посылать их элементам,
которые смогут эти события обработать. Например:

    <nowiki>
    function dispatch()
    {
        // создаём событие
        var event = document.createEvent('Events');
        event.initEvent('myevent', false, false);

        // добавили в событие своё свойство
        event.customdata = "Hello";

        // отослали событие
        window.dispatchEvent(event);
    }

    function func(evt)
    {
        // наше событие должно иметь свойство "customdata",
        // показываем его в сообщении
        alert(evt.customdata);
    }

    // добавили обработчик своего события
    addEventListener("myevent", func, false);

    // создаём и посылаем своё событие
    dispatch();
    </nowiki>

## Ссылки

  - [1](http://www.w3.org/TR/DOM-Level-2-Events/events.html)
  - [2](http://www.w3.org/TR/DOM-Level-3-Events/events.html)
  - [3](http://developer.mozilla.org/en/DOM/)
  - [4](http://developer.mozilla.org/en/Gecko_DOM_Reference/Examples#Example_5:_Event_Propagation)

