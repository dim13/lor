Рецепт для **Chromium**:

1\) Качаешь pepperflash-plugin:

Если у тебя система x32 -
[тут](http://www.slackware.com/~alien/slackbuilds/chromium-pepperflash-plugin/pkg/current/)

Если у тебя система x64 -
[тут](http://www.slackware.com/~alien/slackbuilds/chromium-pepperflash-plugin/pkg64/current/)

Качать нужно файл который больше всего весит (с расширением .txz).

2\) Распаковываешь скачанный архив.

3\) В консоли

    sudo mkdir /usr/lib/pepperflashplugin-nonfree 

4\)

    sudo cp /адрескаталогасархивом/libpepflashplayer.so /usr/lib/pepperflashplugin-nonfree/ 

5\) Перезапускаешь Chromium

6\) PROFIT\!\!\!
