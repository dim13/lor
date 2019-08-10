# Монтирование разделов на сменных носителях при помощи udev

Существует множество способов монтирования сменных носителей. Один из
них - монтирование при помощи udev.

Редактируя правила udev, можно сделать один из двух вариантов
монтирования: вручную (здесь про него и будет рассказано) или
автоматически. Вариант монтирования автоматически я не буду описывать,
т.к. он неудобен тем, что требует для отмонтирования либо прав root'а;
либо настройки sudo; либо приходится извлекать носитель, не
отмонтировав его, что может быть чревато. Да и не всегда
нужно монтировать подключенный сменный носитель (например, если на
него нужно записать образ при помощи dd, отформатировать его и т.п.).

Описанный способ использует udev для того, чтобы при подключении
сменного носителя создать директорию в `/media` и
соответствующую запись в /etc/fstab. Для удобного
распознавания подключенного носителя имя директории в `/media`
состоит из названия файловой системы раздела и имени раздела в `/dev`.

Итак, создаем файл `/etc/udev/rules.d/mnt.rules` и заносим в него
следующее:

<code lang="C">

    KERNEL=="sd[a-z]", GOTO="do-disk-rules"
    KERNEL!="sd[a-z][0-9]", GOTO="end-of-file"
    LABEL="do-disk-rules"
    KERNEL=="sd[a-z]", GROUP="users"
    ACTION=="remove", ENV{ID_FS_TYPE}!="", RUN+="/bin/sed -i '/\/dev\/%k /d' /etc/fstab"
    ACTION=="remove", ENV{ID_FS_TYPE}!="", RUN+="/bin/rmdir /media/$env{ID_FS_TYPE}-%k"
    ACTION=="add", ENV{ID_FS_TYPE}!="", RUN+="/bin/mkdir -p /media/$env{ID_FS_TYPE}-%k"
    &#35; монтирование раздела fat32
    ACTION=="add", ENV{ID_FS_TYPE}=="vfat", RUN+="/bin/sed -i '$a\/dev/%k /media/$env{ID_FS_TYPE}-%k vfat rw,noauto,noatime,dmask=022,gid=user,user,fmask=133,iocharset=koi8-r 0 0' /etc/fstab", OPTIONS="last_rule"
    &#35; монтирование раздела ntfs
    ACTION=="add", ENV{ID_FS_TYPE}=="ntfs", RUN+="/bin/sed -i '$a\/dev/%k /media/$env{ID_FS_TYPE}-%k ntfs-3g rw,noauto,dmask=022,fmask=133,gid=user,user,iocharset=koi8-r 0 0' /etc/fstab", OPTIONS="last_rule"
    &#35; монтирование прочих ФС
    ACTION=="add", ENV{ID_FS_TYPE}!="", RUN+="/bin/sed -i '$a\/dev/%k /media/$env{ID_FS_TYPE}-%k $env{ID_FS_TYPE}  defaults,user 0 0' /etc/fstab"
    LABEL="end-of-file"

</code>

Эти правила будут работать в дистрибутивах с относительно старыми udev.
Для работы с более свежими версиями udev (например, 175-1) придется
немного изменить правила:

<code lang="C">

    &#35; монтирование съемных накопителей
    KERNEL=="sd[a-z]", GOTO="do-disk-rules"
    KERNEL!="sd[a-z][0-9]", GOTO="end-of-file"
    LABEL="do-disk-rules"
    ACTION=="add", ENV{ID_USB_DRIVER}="usb-storage", GROUP="storage"
    IMPORT{program}="/sbin/blkid -o udev -p %N"
    ACTION=="remove", ENV{ID_FS_TYPE}!="", RUN+="/bin/sed -i '/\/dev\/%k /d' /etc/fstab"
    ACTION=="remove", ENV{ID_FS_TYPE}!="", RUN+="/bin/rmdir /media/$env{ID_FS_TYPE}-%k"
    ACTION=="add", ENV{ID_FS_TYPE}!="", RUN+="/bin/mkdir -p /media/$env{ID_FS_TYPE}-%k"
    &#35; монтирование раздела fat32
    ACTION=="add", ENV{ID_FS_TYPE}=="vfat", RUN+="/bin/sed -i '$a\/dev/%k /media/$env{ID_FS_TYPE}-%k vfat rw,noauto,noatime,dmask=022,user,fmask=133,iocharset=koi8-r 0 0' /etc/fstab"
    &#35; монтирование раздела ntfs
    ACTION=="add", ENV{ID_FS_TYPE}=="ntfs", RUN+="/bin/sed -i '$a\/dev/%k /media/$env{ID_FS_TYPE}-%k ntfs-3g rw,noauto,dmask=000,fmask=111,user,locale=ru_RU.koi8-r,allow_other 0 0' /etc/fstab"
    &#35; монтирование прочих ФС
    ACTION=="add", ENV{ID_FS_TYPE}!="", ENV{ID_FS_TYPE}!="ntfs|vfat", RUN+="/bin/sed -i '$a\/dev/%k /media/$env{ID_FS_TYPE}-%k $env{ID_FS_TYPE}  defaults,user 0 0' /etc/fstab"
    LABEL="end-of-file"

</code>

Естественно, если у вас юникод, опцию locale указывать не надо. Если
кодировка другая - измените koi8-r на свою кодировку.

P.S. В новых версиях ntfs-3g "поломали" опцию user, поэтому для того,
чтобы пользователь мог монтировать флешки с ntfs, необходимо либо
скомпилировать ntfs-3g с поддержкой опции user, либо настроить sudo.
Для настройки sudo достаточно вписать в /etc/sudoers следующее:  
<code>

    ALL ALL=(ALL) NOPASSWD: /bin/ntfs-3g,/bin/mount,/bin/umount

</code>
