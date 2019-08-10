Тем кому не хватает аналога кнопочки «турбо», которая присутствует в
Lenovo Power Manager. Можно использовать мой скрипт. При запуске он
смотрит какой говернор используется в данный момент. Если ondemand,
то включаем вентилятор на полную, переключаем говернор в userspace и
выкручиваем частоты на максимум. Если userspace, то частоты в
минимум, говернор в ondemand, вентилятор в auto. Использую его на
Ubuntu 12.04. В принципе без проблем затачивается под любой ноут.
Управление частотой процессора в таком виде, по идее, должно
работь где угодно, а вот управление вентилятором нужно будет
подредактировать под вашу модель ноута. Для включения возможности
ручного управления вентилятором в thinkpad'ах делаем так:

    To enable fan control, the module parameter fan_control=1 must be given to thinkpad-acpi. 
    For example, in Ubuntu 8.04 (Hardy Heron), add the following to /etc/modprobe.d/options: options thinkpad_acpi fan_control=1

Если у вас другая модель thinkpad'а, то подстройте скорость вентилятора
[в соответствии с вот этими
рекомендациями](http://www.thinkwiki.org/wiki/How_to_control_fan_speed).  
Сам скрипт:

    #!/bin/bash
    
    # Checking root user
    if [ "$(id -u)" != "0" ]; then
       echo "This script must be run as root" 1>&2
       exit 1
    fi
    # Checking current state
    if [ `cat /sys/devices/system/cpu/cpu0/cpufreq/scaling_governor` = 'userspace' ]
    then
        # Lowest speed and ondemand governor for all cores
        for CPU in /sys/devices/system/cpu/*/cpufreq/
        do
          echo `cat ${CPU}cpuinfo_min_freq` | tee ${CPU}scaling_setspeed
                echo ondemand | tee ${CPU}scaling_governor
        done
        # Automatic fan control
        echo level auto | tee /proc/acpi/ibm/fan 
    else
        # Otherwise full fan speed
        echo level full-speed | tee /proc/acpi/ibm/fan
        # And full speed for all cores
        for CPU in /sys/devices/system/cpu/*/cpufreq/
        do
                echo userspace | tee ${CPU}scaling_governor
          echo `cat ${CPU}cpuinfo_max_freq` | tee ${CPU}scaling_setspeed
        done
    fi
    TEMP=$((`cat /sys/class/thermal/thermal_zone0/temp`/1000))
    echo "CPU Temp: ${TEMP}C"
    
    # Developed by gear. You can contact me by misty.g3ar@gmail.com
