## description

Самый простой элемент XUL - текстовый блок. Текст может задаваться в
атрибуте *value* или быть внутри открывающего/зкрывающего элементов:

    <nowiki>
    <description value="some text1" />
    <description>some text2</description>
    </nowiki>

## label

Текстовое поле, но более сложное, чем *description*. Если пользователь
кликает по элементу *label*, то движок передаёт фокус элементу, чей id
указан в атрибуте *control*. Атрибут *accesskey* может содержать одну из
букв текста, указанного в атрибуте *value*, и служит горячей клавишей
для активации данного элемента (как будто пользователь кликнул по
элементу, нажимаем ALT+d):

    <nowiki>
    <label accesskey="d" value="Email address" control="email"/>
    <textbox id="email"/>
    </nowiki>

## hbox

Менеджер компоновки по горизонтали. Все дети компонуются по горизонтали
в ряд, размер каждого ребёнка зависит от его параметра *flex*:

    <nowiki>
    <hbox>
        <label value="1" />
        <label value="2" />
        <label value="3" />
        <label value="stretchable label" flex="1"/>
    </hbox>
    </nowiki>

Здесь четвёртый элемент будет изменять свой размер, при изменении
размера окна, а первые три элементы оставаться фиксированными.

## vbox

Менеджер компоновки по вертикали. Все дети компонуются по вертикали в
колонку, размер каждого ребёнка зависит от его параметра *flex*:

    <nowiki>
    <vbox flex="1">
        <label value="1" />
        <label value="2" />
        <label value="3" />
        <label value="stretchable label" flex="1" style="background-color: red;" />
    </vbox>
    </nowiki>

У окна компоновка по умолчанию - vbox. Поэтому в данном примере корневой
элемент vbox имеент атрибут *flex* равным 1 - чтобы он растягивался
вместе с окном, и соответственно, чтобы четвёртый элемент *label*
также растягивался.

## button

Обычная кнопка. Атрибут *label* задаёт текст для кнопки, *image* - URL
изображения для кнопки. Атрибут *type* управляет поведением кнопки. Он
может быть

  - *checkbox* - кнопка ведёт себя как элемент *checkbox*, она может
    быть в двух состояниях - нажата и отжата:

<!-- end list -->

    <nowiki>
    <button label="Press Me" type="checkbox" oncommand="alert('Clicked');"/>
    </nowiki>

  - *menu* - кнопка с выпадающим меню, когда на неё кликают. Само меню
    задаётся как элемент *menupopup* внутри кнопки:

<!-- end list -->

    <nowiki>
    <button label="Press Me" type="menu">
        <menupopup>
            <menuitem label="1" value="first"/>
            <menuitem label="2" value="second"/>
            <menuitem label="3" value="third"/>
        </menupopup>
    </button>
    </nowiki>

  - *menu-button* - аналогично предыдущему, но меню выпадает, когда
    кликают на стрелку.
  - *radio* - кнопка, аналогичная *checkbox*, за исключением того, что
    только одна кнопка типа *radio* может быть нажата в группе.
  - *repeat* - кнопка с автоповтором, пока мышь держится нажатой на ней.

