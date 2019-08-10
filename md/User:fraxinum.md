# "magic folder" в любой системе

Суть: создание "сортировочной" директории. При попадании файла в такое
место, к нему применяется правило (или набор правил), согласно
заданным шаблонам. К примеру - архив будет автоматически
разархивирован, или видеофайл будет перемещён в каталог "video".
На что хватит фантазии.

Тут я приведу пример наибыстрейшего способа развёртывания такой штуки.

### *пред подготовка*

Установите в вашу систему incron. Это демон, что работает по "шаблонам"
cron'ов, только относительно изменений в файловой системе.

Для deb систем: **sudo apt-get install incron**.

Добавляем себя в список тех, кто может оперировать демоном: **sudo echo
$USER \>\> /etc/incron.allow**

О самой тулзе очень хорошо разложено тут:
<http://wiki.opennet.ru/Incrontab>

### *подготовка*

Определитесь с такими вопросами как: 1. место хранения скрипта с
правилами 2. директорией, которой вы дадите желаемый функционал.

Буду показывать на живом примере, ибо мне лень переписывать пути. У меня
это '/home/lord/Dropbox/Files/linux/scripts/sorter.sh' и
'/home/lord/Downloads', соответственно.

Создаём скрипт с таким содержимым (тут назначены мои правила, только для
перемещений, согласно типу файла):

    #!/bin/bash
    
    ROOT_DIR="$1"
    FILE="$2"
    INCRON_ACTION="$3"
    LOG="${ROOT_DIR}/SORTER.log"
    
    ## for case:
    # ext regex
    shopt -s extglob
    # case insensitive
    shopt -s nocasematch
    
    function log {
        action_status="$1"
        action_error="$2"
        case ${action_status} in
            "ok")
                echo "[ $(date +%Y.%m.%d-%H:%M:%S) ] >> '${FILE}' to '${ROOT_DIR}/${mask_dir}/' >> ${INCRON_ACTION}"\
                    >> "$LOG"
            ;;
            "err")
                echo "!! [ $(date +%Y.%m.%d-%H:%M:%S) ] >> '${FILE}' to '${ROOT_DIR}/${mask_dir}/' >> ${INCRON_ACTION}"\
                    >> "$LOG"
                echo "${action_error}" >> "$LOG"
            ;;
        esac
    }
    
    function standart_action {
        [ ! -d "${ROOT_DIR}/${mask_dir}" ] && mkdir "${ROOT_DIR}/${mask_dir}"
        sleep 1
        # fix for firefox downloads
        [ -f "${ROOT_DIR}/${FILE}.part" ] && return
        file_action=$(mv "${ROOT_DIR}/${FILE}" "${ROOT_DIR}/${mask_dir}/" 2>&1)
        if [ $? -eq 0 ]; then
            log 'ok'
        else
            log 'err' "${file_action}"
        fi
    }
    
    case "${FILE}" in
        # video
        *.avi|*.3gp|*.flv|*.mkv|*.mp4|*.mp?(e)g|*.mov)
            mask_dir="SORTER-videos"
            standart_action
        ;;
        # audio
        *.mp3|*.flac|*.wav)
            mask_dir="SORTER-audio"
            standart_action
        ;;
        # DEB's && RMP's
        *.deb|*.rpm)
            mask_dir="SORTER-install"
            standart_action
        ;;
        # pictures
        *.jp?(e)g|*.gif|*.png|*.raw|*.ico|*.xcf|*.svg)
            mask_dir="SORTER-pictures"
            standart_action
        ;;
        # images
        *.iso)
            mask_dir="SORTER-images"
            standart_action
        ;;
        # archives
        *.tar|*.bz2|*.gz|*.7z|*.rar|*.tbz2|*.zip|*.tgz)
            mask_dir="SORTER-archives"
            standart_action
        ;;
        # doc's
        *.pdf|*.djvu|*.doc?(x)|*.xm[ls]|*.od[stf]|*.rtf)
            mask_dir="SORTER-docs"
            standart_action
        ;;
    esac

Делаем его исполняемым (chmod +x).

Правила задаём в секции case, согласно типам (или оставляем как есть,
если вас всё устраивает из коробки).

После этого выполняем **incrontab -e**, и вписываем правило для нашей
директории.

Строка должна принять вид, подобный к:

    /home/lord/Downloads IN_CREATE,IN_MOVED_TO /home/lord/Dropbox/Files/linux/scripts/sorter.sh $@ $# $%

Вы должны поменять пути на свои. Кроме путей всё оставляем прежним.

### *в работ*е

После удачного применения правила incrontab'a - наша директория начала
свою работу, согласно правилам.

В данном скрипте ведётся лог всех действий. В корне магического каталога
файлик "SORTER.log".

### *послесловие*

Что бы применить наши правила к уже существующим файлам - можно
воспользоваться скриптиком-костыликом.

    #!/bin/bash
    
    ROOT_PATH="/home/lord/Downloads/"
    SORTER_PATH="/home/lord/Dropbox/Files/linux/scripts/sorter.sh"
    
    for i in $(find "${ROOT_PATH}" -maxdepth 1 -type f | sed 's/ /•/g'); do
        sleep 0.1
        filename="$(echo ${i} | sed 's/•/ /g')"
        "${SORTER_PATH}" "${ROOT_PATH}" "$(basename "${filename}")" "NULL" &
    done
    
    echo "DONE"

Меняем пути на нужные - и запускаем. Помним, что если файлов много -
перемещать так же будет не мгновенно.

Вот и всё, что хотел сказать по этой теме.

-----
