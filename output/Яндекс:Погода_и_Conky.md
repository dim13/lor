*Написано по посту пользователя
[osabio](https://www.linux.org.ru/people/osabio/profile). Оригинальный
[пост](https://www.linux.org.ru/forum/desktop/11202205).* Встроенны�
средств conky (без неочевидного употребления внешни�
wget/grep/sed/awk) не хватает на вывод погоды с какого-либо сайта.
Пользователь osabio предложил скрипт для вывода текущей
температуры/ветра (в текстовом режиме) с Яндекс.Погоды.

Нужно:

  - Conky
  - Python 2.7

Открываем текстовый редактор, пишем туда код ниже, и сохраняем в
"\~/weather/weather.py".

    #!/usr/bin/python
    import re
    import sys
    import urllib

    from urllib.request import urlopen
    home='/home/osabio'
    html = urlopen("https://pogoda.yandex.ru/nizhny-novgorod")
    htmls = html.read().decode('utf-8')
    tempera=re.compile(r'<div\ class="current\-weather__thermometer\ current\-weather__thermometer_type_now">(.*?)</div>')
    temp=tempera.findall(htmls)[0].replace(" ","").replace("°C","")
    windera=re.compile(r'<div\ class="current\-weather__info\-row\ current\-weather__info\-row_type_wind"><span\ class="current\-weather__info\-label">Ветер:\ </span>\ (.*?)<abbr')
    wind=windera.findall(htmls)[0]
    kompasera=re.compile(r'<abbr title=".*?">(.*?)</abbr>')
    kompas=kompasera.findall(htmls)[0]
    ftemp = open(home+r'/weather/temp', "w+")
    ftemp.write(temp)
    ftemp.close()
    ftemp = open(home+r'/weather/wind', "w+")
    ftemp.write(wind)
    ftemp.close()
    ftemp = open(home+r'/weather/kompas', "w+")
    ftemp.write(kompas)
    ftemp.close()

Также нужно будет :

  - Исправить строчку, содержащую в себе ссылку (html=urlopen(...)), на
    ссылку для своего города
  - Сменить значение переменной home на домашнюю папку своего
    пользователя
  - Создать файлы для работы скрипта (выполнить нижеприведенный скрипт
    bash):

<!-- end list -->

    mkdir ~/weather; touch ~/weather/temp ~/weather/kompas ~/weather/wind

Настройки conkyrc в минимальном виде:

    ${execi 600 python ~/weather/weather.py}\
    ${execi 600 cat ~/weather/temp} \ #Вывод температуры (число без знака градуса)
    ${execi 600 cat ~/weather/wind} \ #Вывод скорости ветра (число без едениц измерения)
    ${execi 600 cat ~/weather/kompas} \ #Вывод направления ветра (текстом: Ю, СЗ и т.п)

С условиями conky отображения направления ветра стрелочками:

    ${execi 600 python ~/weather/weather.py}\
    \
    ${execi 600 cat ~/weather/temp} °C\
    ${execi 600 cat ~/weather/wind} М/С\
    ${if_match "${execi 600 cat ~/weather/kompas}" == "ЮВ"}↖${endif}\
    ${if_match "${execi 600 cat ~/weather/kompas}" == "СЗ"}↘${endif}\
    ${if_match "${execi 600 cat ~/weather/kompas}" == "ЮЗ"}↗${endif}\
    ${if_match "${execi 600 cat ~/weather/kompas}" == "СВ"}↙${endif}\
    ${if_match "${execi 600 cat ~/weather/kompas}" == "С"}↓${endif}\
    ${if_match "${execi 600 cat ~/weather/kompas}" == "В"}←${endif}\
    ${if_match "${execi 600 cat ~/weather/kompas}" == "Ю"}↑${endif}\
    ${if_match "${execi 600 cat ~/weather/kompas}" == "З"}→${endif}\
     ${execi 600 cat ~/weather/wind}\

Получаем обновление виджета погоды каждые 10 минут. Радуемся.

Заменить время можно по всему документу автозаменой, т.к. везде стоит
600 с. Не рекомендуется ставить время обновления меньше 61 с - Яндекс
будет банить.

P.S. Время работы скрипта - \~0,55 сек.