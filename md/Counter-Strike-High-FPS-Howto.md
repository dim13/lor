Данная статья ориентирована на пользователей со слабыми видеокартами.
Описанным ниже способом я добился стабильных 60 FPS на ноутбуке с
видеокартой Intel GMA 900. Поскольку Counter-Strike распространяется
только через Steam, а Valve пока что официально поддерживают только
Ubuntu, то гайд ориентирован в первую очередь на пользователей Ubuntu.
Все манипуляции тестировались на Ubuntu 12.04.

# Настройка ОС

Нужно по-максимуму разгрузить видеокарту. Для этого, если вы пользуетесь
Unity, то переключитесь на Unity-2D. Если же вы пользуетесь другим WM,
то просто отключите композитинг.

# Настройка игры

## Штатные средства

Заходим в Options, на вкладку Video. Там выбираем как можно меньшее
разрешение. Точнее минимально возможное для комфортной игры с вашей
точки зрения. Главное - помните простое правило: выше разрешение - ниже
FPS. Я остановился на 800x600. Далее снимаем галочки с опций "Enable HD
models if available", "Allow custom addon content" и ставим галочку на
"Low video quality". Применяем настройки, заходим в игру и проверяем. На
новых Intel'овских видеокартах этого должно хватить для стабильных 100
FPS. Если вы счастливый обладатель видеокарт GMA9XX, то для комфортной
игры нужно будет еще больше исковеркать графику конфигом.

## High FPS Config

**userconfig.cfg**

    fps_max "60" //observed frames per second on client
    cl_cmdrate "65" //fps + 5
    cl_updaterate "30" //half observed frames per second on client 
    ex_extrapmax "1.2" //default Half-Life value
    cl_cmdbackup "2"
    rate "15000"
    ex_interp "0.05"
    cl_download_ingame "0"
    
    //Fps maxing
    cl_himodels "0"
    gl_monolights "1"
    cl_highdetail "0"
    gl_dither "0"
    gl_cull "1"
    gl_keeptjunctions "0"
    gl_lightholes "0"
    r_lightmap "0"
    gl_max_size "128"
    gl_playermip "1"
    gl_picmip "2"
    cl_shadows "0"
    cl_weather "0"
    r_mmx "1"
    gl_round_down "5" //high val - ugly textures, higher fps
    gl_spriteblend "0"
    gl_texturemode "GL_NEAREST_MIPMAP_NEAREST"
    gl_wateramp "0"
    gl_ztrick "0"
    r_dynamic "0"
    r_decals "10" //time before decals disapear
    r_mirroralpha "0"
    violence_ablood "0"
    violence_hblood "0"
    violence_agibs "0"
    violence_hgibs "0"
    cl_minmodels "1"
    mp_decals "5"
    r_detailtextures "0"
    fastsprites "2"
    cl_corpsestay "10"
    max_shells "1"
    max_smokepuffs "1"

Открываем каталог *cstrike*, создаем там файл *userconfig.cfg* и
копируем в него приведенный выше конфиг. Описание использованных
в нем cvar'ов можно найти
[здесь](https://developer.valvesoftware.com/wiki/NS_Variables) и
[здесь](http://www.elxdraco.net/cvarlist/).
