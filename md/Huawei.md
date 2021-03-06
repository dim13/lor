# Автоподключение модема huawei при помощи udev

Если вы являетесь "счастливым обладателем" МТС-модема, то прежде всего
вам понадобится установить пакеты wvdial и usb\_modeswitch. Первый
нужен для поднятия ppp-соединения, а второй — для переключения этого
"свистка" из режима эмуляции CD-ROM в режим модема.

Затем нужно при помощи lsusb после подключения модема узнать его VID и
PID (они нужны будут для usb\_modeswitch).

После того, как это готово, создаем такое правило udev: <code>

    SUBSYSTEM=="usb", ACTION=="add", ATTRS{idProduct}=="1436", ATTRS{idVendor}=="12d1", MODE:="0666", RUN+="/etc/udev/rules.d/huawei on"
    SUBSYSTEM=="usb", ACTION=="remove", ATTRS{idProduct}=="1436", ATTRS{idVendor}=="12d1", RUN+="/etc/udev/rules.d/huawei off"

</code>

И создаем исполняемый файл /etc/udev/rules.d/huawei с содержимым <code>

    #!/bin/bash
    STATUS=0
    
    if [ "$1" = "on" ]; then
            rc.d stop network
            sleep 1
            aplay /usr/share/skype/sounds/CallConnecting.wav
            usb_modeswitch -v 0x12d1 -p 0x1436 -H
            echo "AT^SYSCFG=14,2,3fffffff,0,1" > /dev/ttyUSB2
            while [ $STATUS = 0 ]; do
                    aplay /usr/share/skype/sounds/CallConnecting.wav
                    sleep 1
                    wvdial &
                    PID=$!
                    sleep 1
                    ps $PID && STATUS=1 || aplay /usr/share/skype/sounds/CallFailed.wav
            done
            aplay /usr/share/skype/sounds/CallRingingIn.wav
    else
            killall -9 wvdial
            killall -9 pppd
            rc.d start network
    
    fi

</code> (Если у вас не установлен skype, измените звуки на какие-нибудь
другие).

Все, теперь при подключении модема будут проигрываться характерные звуки
и подниматься ppp-соединение, а при отключении — соединение будет
разрываться.

P.S. Для МТС-модема нужен еще файл настройки wvdial (/etc/wvdial.conf):
<code>

    [Dialer Defaults]
    Init1 = ATZ
    Init2 = AT+CGDCONT=1, "IP", "internet.mts.ru"
    Modem Type = USB Modem
    Modem = /dev/ttyUSB0
    Baud = 57600
    New PPPD = yes
    Phone = *99***1#
    Password = mts
    Username = mts
    Stupid Mode = yes

</code>
