## Что такое XPCOM

**XPCOM** - Кроссплатформенная Компонентная Объектная Модель (Cross
Platform Component Object Model). Данная модель схожа с небезызвестным
COM в ОС Windows, и предоставляет собой механизм создания и динамической
загрузки компонентов в Вашу программу. Компонент обычно представляет
собой некий класс, объект которого можно создать и использовать в
программе. Компоненты XPCOM могут быть использованы из языков C, C++,
JavaScript. В отличие от прямой линковки в библиотекой, как например это
обычно делается в языке C++ (g++ -o program -lsomelib), механизм XPCOM
полностью динамический, и позволяет подгружать новые типы во время
выполнения программы. XPCOM компоненты могут быть написаны как на
C++, так и на JavaScript. Большинство интерфейсов, используемых в
JavaScript в XUL есть ни что иное, как XPCOM компоненты.

При написании XPCOM компонентов нужно знать четыре понятия:

  - **интерфейс** - набор методов и атрибутов, которые как-то логически
    связаны друг с другом;
  - **компонент** - это объект, реализующий один или больше интерфейсов;
  - **фабрика** - объект (или функция), создающий компоненты, реализует
    интерфейс
    [nsIFactory](http://developer.mozilla.org/en/docs/nsIFactory);
  - **модуль** - объект, управлящий набором фабрик, реализует интерфейс
    [nsIModule](http://developer.mozilla.org/en/docs/nsIModule).

Стоит отметить, что эти понятия связаны не только с XPCOM, а с
компонентным программированием вообще. Эти же понятия вы
можете встретить, например, при работе с *D-Bus*. В *D-Bus* роль
модуля выполняет программа, создающая корневой dbus сервис (по
аналогии с фабрикой), роль компонентов выполняют дочерние dbus
объекты, которые аналогично с XPCOM могут реализовывать несколько
интерфейсов.

## Подготовка

Для компиляции XPCOM компонента необходим компилятор, набор стандартных
заголовочных файлов, набор для разработки (SDK) для XUL. Для
дистрибутивов, основанных на Debian, необходимо установить
пакеты build-essential, xulrunner-1.9-dev. Xulrunner-1.9 у Вас уже
должен стоять.

В Makefile в исходниках примера установите переменную GECKO_SDK_PATH
равной пути к заголовкам и IDL файлоам xulrunner. В KUbuntu Gutsy это
/usr/lib/xulrunner-devel-1.9a8. Посмотрите путь, куда установились файлы
из пакета xulrunner-1.9-dev (dpkg -L xulrunner-1.9-dev), там должны быть
каталоги include, idl, bin, lib - этот путь и должен быть в
GECKO_SDK_PATH. Для работы из командной строки, можно также сделать

    <nowiki>
    # export XUL=/usr/lib/xulrunner-devel-1.9a8
    </nowiki>

Теперь мы сможем быстро вызвать какую-нибудь утилиту из комплекта
разработчика XUL, например:

    <nowiki>
    # $XUL/bin/xpidl

    ERROR: must specify output mode
    Usage: /usr/lib/xulrunner-devel-1.9a8/bin/xpidl -m mode [-w] [-v] [-t version number] [-d filename.pp]
    ...
    </nowiki>

В данном примере мы рассмотрим простейший компонент, который складывает
два числа и возвращает результат.

## Создаём GUID

Для компонента и интерфейса необходимо создать уникальный идентификатор,
GUID. Используем утилиту **uuidgen** из пакета **e2fsprogs**.

## Создаём файл с описанием нашего компонента (IDL)

В XPCOM для описания компонентов используется формат файла, называемый
[XPIDL](http://www.mozilla.org/scriptable/xpidl/idl-authors-guide),
схожий с IDL из CORBA. В этом файле описывается имя интерфейса, его
методы, атрибуты, уникальный идентификатор, наследование и другое. В
нашем случае:

    <nowiki>
    #include "nsISupports.idl"

    [scriptable, uuid(_YOUR_INTERFACE_GUID_)]
    interface IMyComponent : nsISupports
    {
      // Только один метод, для сложения чисел
      long Add(in long a, in long b);
    };
    </nowiki>

Мы описали интерфейс с именем **IMyComponent**, имеющий всего лишь один
метод для сложения чисел. Вместо **_YOUR_INTERFACE_GUID_**
вставляем сгенерированный ранее GUID. Наш интерфейс должен
наследоваться от базового интерфейса
[nsISupports](http://developer.mozilla.org/en/nsISupports), который
имеет общую функциональность для всех интерфейсов - методы
**AddRef**, **QueryInterface** и **Release** (аналогично интерфейсу
**IUnknown** в COM). Метод **AddRef** увеличивает внутренний счётчик на
1 при дублировании указателя на интерфейс. **Release** служит для
уменьшения внутреннего счётчика на 1, когда один из дублированных
указателей больше не нужен. Если счётчик равен 0, то **Release**
безжалостно удаляет самого себя (объект this).
**QueryInterface** служит для получения указателя на один из
поддерживаемых интерфейсов (с увеличением счётчика, если
интерфейс найден), полагаясь на переданный в качестве аргумента
идентификтор.

Ключевое слово **scriptable** указывает на то, что мы можем использовать
этот интерфейс из скриптового языка, например JavaScript. Без указания
этого ключевого слова мы не сможет использовать наш компонент из
JavaScript.

Сохраняем этот файл как *IMyComponent.idl*.

## Генерация заголовка и xpt библиотеки

Теперь нам необходимо из описания нашего интрефейса, XPIDL файла,
сгенерировать заголовочный файл, который мы будем использовать
далее при реализации методов, и .xpt библиотеку. XPT библиотека - это
бинарное представления файла описания XPIDL, она аналогично содержит имя
интерфейса, его методы, атрибуты, уникальный идентификатор и т.д.,
только в бинарном виде, что облегчает жизнь загрузчику компонента
(XPConnect).

Следующие строки не являются обязательными, т.к. они уже внесены в
Makefile из примера, но всё-таки не помешает знать, как это делается.

Итак, генерируем заголовочный файл:

    <nowiki>
    # $XUL/bin/xpidl -m header -I$XUL/idl IMyComponent.idl
    </nowiki>

Создаётся файл *IMyComponent.h*.

    <nowiki>
    # $XUL/bin/xpidl -m typelib -I$XUL/idl IMyComponent.idl
    </nowiki>

Создаётся файл *IMyComponent.xpt*.

## Создаём реализацию компонента

Создаём файл *MyComponent.h* - здесь будет объявление (definition)
класса нашего компонента:

    <nowiki>
    #ifndef _MY_COMPONENT_H_
    #define _MY_COMPONENT_H_

    #include "IMyComponent.h"

    // Строка-идентификатор нашего компонента, т.н. Contract ID. Используя это строку мы сможем загрузить компонент.
    // Обычно эта строка имеет какой-то осмысленный вид, куда вносится домен, имя компонента, версия.
    #define MY_COMPONENT_CONTRACTID "@mydomain.com/XPCOMSample/MyComponent;1"

    // Имя класса
    #define MY_COMPONENT_CLASSNAME  "A Simple XPCOM Sample"

    // Уникальный 128-битный идентификатор компонента, аналогичный идентификатору интерфейса, только
    // в бинарном виде.
    #define MY_COMPONENT_CID        { 0x597a60b0, 0x5272, 0x4284, { 0x90, 0xf6, 0xe9, 0x6c, 0x24, 0x2d, 0x74, 0x6 } }

    // Объявляем класс компонента, наследуемся от сгенерированного класса IMyComponent.
    class MyComponent : public IMyComponent
    {
    public:

      // Этот макрос вставляет объявление стандартных трёх public методов из nsISupports,
      // и protected атрибута-счётика ссылок.
      NS_DECL_ISUPPORTS

      // Этот макрос вставляет объявление нашего метода Add().
      NS_DECL_IMYCOMPONENT

      MyComponent();
      virtual ~MyComponent();
    };

    #endif //_MY_COMPONENT_H_
    </nowiki>

Создаём файл *MyComponent.cpp* - здесь будет реализация класса нашего
компонента:

    <nowiki>
    #include "MyComponent.h"

    // Этот макрос вставляет стандартную реализацию методов AddRef(), Release() и QueryInterface().
    // Параметры - класс и интерфейс.
    NS_IMPL_ISUPPORTS1(MyComponent, IMyComponent)

    MyComponent::MyComponent()
    {}

    MyComponent::~MyComponent()
    {}

    // Реализация метода Add()
    NS_IMETHODIMP MyComponent::Add(PRInt32 a, PRInt32 b, PRInt32 *_retval)
    {
            *_retval = a + b;
            return NS_OK;
    }
    </nowiki>

## Создаём код для загрузчика

Создаём файл *MyComponentModule.cpp*, где будет находится информация для
загрузчика нашего компонента:

    <nowiki>
    #include "nsIGenericFactory.h"
    #include "MyComponent.h"

    // Создаём конструктор (фабрику) нашего компонента. Этот макрос
    // создаёт static функцию с именем <имя_аргумента>Constructor, в нашем случае
    // MyComponentConstructor, которая далее используется в структуре components.
    NS_GENERIC_FACTORY_CONSTRUCTOR(MyComponent)

    // Структура с описанием нашего компонента
    static nsModuleComponentInfo components[] =
    {
        {
           MY_COMPONENT_CLASSNAME,
           MY_COMPONENT_CID,
           MY_COMPONENT_CONTRACTID,
           MyComponentConstructor,
        }
    };

    // Создаём точку входа с необходимой информацией. Это аналогично extern "C" функциям
    // в С++ библиотеках, загружаемых через dlopen(), и которые создают и возвращают указатель
    // на объект класса.
    NS_IMPL_NSGETMODULE("MyComponentsModule", components)
    </nowiki>

Весь необходимый исходный код написан, теперь создадим файл сборки -
*Makefile*.

## Makefile

Создаём файл *Makefile*, не забываем поправить переменную
**GECKO_SDK_PATH**.

    <nowiki>
    CXX = g++
    CPPFLAGS += -fno-rtti -fno-exceptions -shared

    GECKO_SDK_PATH = /usr/lib/xulrunner-devel-1.9a8

    GECKO_CONFIG_INCLUDE = -include mozilla-config.h

    GECKO_DEFINES  = -DXPCOM_GLUE -DXPCOM_GLUE_USE_NSPR

    GECKO_INCLUDES = -I$(GECKO_SDK_PATH)/include -I$(GECKO_SDK_PATH)/sdk/include

    GECKO_LDFLAGS =  -L$(GECKO_SDK_PATH)/lib -lxpcomglue_s -lxul -lxpcom -lplds4 -lplc4 -lnspr4 -lpthread -ldl

    FILES = MyComponent.cpp MyComponentModule.cpp

    TARGET = MyComponent.so

    build:
            $(GECKO_SDK_PATH)/bin/xpidl -m header  -I$(GECKO_SDK_PATH)/idl IMyComponent.idl
            $(GECKO_SDK_PATH)/bin/xpidl -m typelib -I$(GECKO_SDK_PATH)/idl IMyComponent.idl
            $(CXX) -Wall -O2 -o $(TARGET) $(GECKO_CONFIG_INCLUDE) $(GECKO_DEFINES) $(GECKO_INCLUDES) $(CPPFLAGS) $(CXXFLAGS) $(FILES) $(GECKO_LDFLAGS)
            chmod +x $(TARGET)
            strip $(TARGET)

    clean:
            rm -f IMyComponent.xpt
            rm -f IMyComponent.h
            rm -f $(TARGET)
    </nowiki>

## Сборка и запуск

Мы создали всё необходимое для сборки XPCOM компонента. Теперь просто
используем утилиту *make* для сборки:

    <nowiki>
    # make
    </nowiki>

Если всё прошло удачно, создастся динамическая библиотека
*MyComponent.so*. Эту библиотеку и .xpt файл копируем в каталог
**components** нашего XUL проекта, и в проекте создаём .xul файл
наподобие:

    <nowiki>
    <?xml version="1.0"?>
    <?xml-stylesheet href="chrome://global/skin/" type="text/css"?>

    <window
        title="Example"
        width="320"
        height="200"
        ns="http://www.mozilla.org/keymaster/gatekeeper/there.is.only.xul">

    <label id="result" />

    <script type="application/x-javascript">

    function MyComponentTest()
    {
        var obj = null;

        try
        {
            // По Contract ID получаем доступ к модулю
            const cid = "@mydomain.com/XPCOMSample/MyComponent;1";
            obj = Components.classes[cid].createInstance();

            // Получаем указатель на объект нашего интерфейса
            obj = obj.QueryInterface(Components.interfaces.IMyComponent);
        }
        catch(err)
        {
            alert(err);
            return;
        }

        // Вызываем наш метод для сложения чисел
        var res = obj.Add(3, 4);

        document.getElementById("result").value = "3 + 4 = " + res;
    }

    MyComponentTest();

    </script>
    </window>
    </nowiki>

Запускаем:

    <nowiki>
    # xulrunner ./application.ini
    </nowiki>

Вы должны увидеть окно с текстом **3 + 4 = 7** внутри.

Исходники XUL проекта и XPCOM компонента приведены ниже.

## Ссылки

  - ![Исходники example3 с готовым компонентом (переименуйте в
    .tar.bz2)](example3.tar.bz2
    "Исходники example3 с готовым компонентом (переименуйте в .tar.bz2)")
  - ![Исходники компонента](xpcom-sample.tar.bz2 "Исходники компонента")
  - [Оригинал статьи на
    английском](http://www.iosart.com/firefox/xpcom/)
  - [Информация по XPConnect от
    Mozilla](http://www.mozilla.org/scriptable/)
  - [Информация по XPCOM от
    Mozilla](http://www.mozilla.org/projects/xpcom/index.html)
  - [Информация по
    XPCOM](http://www.mozilla.org/projects/xpcom/book/cxc/html/index.html)
  - [Ещё пример по
    XPCOM](http://lxr.mozilla.org/seamonkey/source/xpcom/sample/)
  - [Информация по XPIDL от
    Mozilla](http://www.mozilla.org/scriptable/xpidl/idl-authors-guide/)

*Это статья основана на примере по XPCOM от Alex Sirota. Редакция для
запуска под xulrunner-1.9, текст разделов и комментарии в С++ и
JavaScript коде на этой странице от пользователя Alex Custov.*

