Принцип написания XPCOM компонента на JavaScript схож с принципом
написания оного на C++. Однако, здесь в более явном виде
используются понятия модуля и фабрики. Алгоритм действий прост:

  - создать XPIDL описание интерфейса, сгенерировать .xpt библиотеку;
  - создать JavaScript реализацию интерфейса, фабрики, модуля;
  - положить .xpt библиотеку и .js реализацию в каталог **components**
    нашего проекта.

Итак, приступим.

## Подготовка

Подготовка включает в себя все разделы вплоть до **Генерация заголовка и
xpt библиотеки** включительно в статье
[XUL_XPCOM](XUL_XPCOM). Тут ничего не поменялось, кроме
разве что GUID.

## Пишем JavaScript XPCOM

Создаём файл *IMyComponent.js* с таким содержимым:

    <nowiki>
    const Ci = Components.interfaces;

    // константы, аналогично как в C++
    const CLASS_ID    = Components.ID("b058921b-cdc2-43b3-8515-e2b9d6133263");
    const CLASS_NAME  = "Simple JavaScript Componennt";
    const CONTRACT_ID = "@mydomain.com/XPCOMSample/MyComponent;1";
    const INTERFACE   = Ci.IMyComponent;

    // Собственно наш компонент
    function MyComponent() {}

    MyComponent.prototype =
    {
        // Метод сложения двух чисел, описанный в .idl файле
        Add: function(a, b)
        {
            return a + b;
        },

        // Мы также должны реализовать QueryInterface, чтобы
        // проверить правильность использования объекта. Это означает, что
        // мы не сможем использовать этот объект как указатель на интерфейс, который в нём
        // не реализован. Методы AddRef() и Release() из nsISupports не реализованы.
        QueryInterface: function(aIID)
        {
            // Мы знаем только о базовом интерфейсе nsISupports и нашем IMyComponent,
            // иначе выбрасываем исключение.
            if(!aIID.equals(INTERFACE) && !aIID.equals(Ci.nsISupports))
                throw Components.results.NS_ERROR_NO_INTERFACE;

            return this;
        }
    };

    // Наша фабрика, создаёт и возвращает синглетон MyComponent.
    // Частично реализует интерфейс nsIFactory.
    var MyComponentFactory =
    {
        singleton: null,

        // Вызывается из скрипта для создания объекта XPCOM компонента
        createInstance: function(aOuter, aIID)
        {
            if(aOuter != null)
                throw Components.results.NS_ERROR_NO_AGGREGATION;

            // объект уже создан?
            if(this.singleton == null)
                this.singleton = new MyComponent();

            return this.singleton.QueryInterface(aIID);
        }
    };

    // Модуль, реализует интерфейс nsIModule, позволяет работать с нашей фабрикой
    var MyComponentModule =
    {
        registerSelf: function(aCompMgr, aFileSpec, aLocation, aType)
        {
            aCompMgr = aCompMgr.QueryInterface(Components.interfaces.nsIComponentRegistrar);
            aCompMgr.registerFactoryLocation(CLASS_ID, CLASS_NAME, CONTRACT_ID, aFileSpec, aLocation, aType);
        },

        unregisterSelf: function(aCompMgr, aLocation, aType)
        {
            aCompMgr = aCompMgr.QueryInterface(Components.interfaces.nsIComponentRegistrar);
            aCompMgr.unregisterFactoryLocation(CLASS_ID, aLocation);
        },

        getClassObject: function(aCompMgr, aCID, aIID)
        {
            // Абсолютно неправильное использование объекта модуля.
            // В принципе, такого не должно возникать.
            if(!aIID.equals(Components.interfaces.nsIFactory))
                throw Components.results.NS_ERROR_NOT_IMPLEMENTED;

            // Запрос нашего интерфейса, которому однозначно сопоставлена
            // фабрика.
            if(aCID.equals(CLASS_ID))
                return MyComponentFactory;

            // Мы не смогли обработать запрос объекта, и ничего не знаем о запрошенном
            // интерфейсе.
            throw Components.results.NS_ERROR_NO_INTERFACE;
        },

        canUnload: function(aCompMgr) { return true; }
    };

    // Главный загрузчик, должен возвращать объект нашего модуля
    function NSGetModule(aCompMgr, aFileSpec) { return MyComponentModule; }
    </nowiki>

## Запуск

Копируем .xpt библиотеку и .js файл в каталог *components* нашего XUL
проекта, и создаём .xul файл абсолютно идентичный таковому из статьи
[Пишем XPCOM компонент на C++](XUL_XPCOM). Запуск и результат
также должны совпадать с результатами из статьи [Пишем XPCOM компонент
на C++](XUL_XPCOM) - Вы должны увидеть окно с текстом **3 +
4 = 7** внутри.

## Анализ

Как Вы видите, единственное отличие от предыдущей статьи в том, что
XPCOM компонент реализован не на C++, а на JavaScript. Естественно,
такие компоненты проще писать и сопровождать.

## Ссылки

  - ![Исходники example4 с готовым компонентом (переименуйте в
    .tar.bz2)](example4.tar.bz2
    "Исходники example4 с готовым компонентом (переименуйте в .tar.bz2)")
  - [Создание XPCOM компонента на
    JavaScript](http://kb.mozillazine.org/Implementing_XPCOM_components_in_JavaScript)
  - [Создание XPCOM компонента на
    JavaScript-2](http://www.builderau.com.au/program/soa/Creating_XPCOM_components_with_JavaScript/0,339024614,339206503,00.htm)

[Category:XUL](Category:XUL)