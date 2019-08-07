Postfix - агент передачи почты (MTA — mail transfer agent).

## Архитектура

## Базовые настройки

## Настройки внешних фильтров для Postfix

Сам Postfix занимается лишь одной задачей - передачей почты. Установка
всевозможных фильтров в систему реализуется через управление потоками
данных между компонентами (демонами), входящих в Postfix, и внешними
приложениями.

##### Пример простого фильтра содержимого почты (ClamAV & SpamAssassin), основанного на pipe.

<http://www.postfix.org/FILTER_README.html#simple_filter>)

Сценарий обработки почты таков: smtpd, слушающий tcp/25, принимает почту
из сети. smtpd запущен с выставленным параметров content_filter. После
проверки содержимого на корректность демоном cleanup оно передаётся
фильтру, выполненному в виде процесса pipe, вызывающего внешнюю
программу. После соответствующей обработки программа должна с
помощью sendmail отправить письмо в очередь сообщений.

В файле master.cf сделайте правку записи сервера smtpd (это можно
сделать и в main.cf):

    smtp      inet  n       -       -       -       -       smtpd
       -o content_filter=mailfilter:dummy

... и определите программу-фильтр:

    mailfilter  unix    -   n   n   -   -   pipe
        flags=R user=filter argv=/usr/local/bin/mailfilter -f ${sender} -- ${recipient}

Не забудьте создать пользователя, указанного в параметре user.

Теперь нам осталось побеспокоиться о конечной программе
(подразумевается, что ClamAV & Spamassassin уже
установлены и настроены). У меня сей скрипт появился откуда-то
из недр Интернета и был немного мной подредактирован. Мой дистрибутив -
Debian lenny.

    #!/bin/sh

    MYDOMAIN=example.org
    INSPECT_DIR=/tmp
    #Не забываем создать директорию карантина
    QUARANTINE_DIR=/var/lib/mailfilter/quarantine

    SENDMAIL="/usr/sbin/sendmail -i"
    SPAMC="/usr/bin/spamc"
    FILTER_SPAMC="$SPAMC -u postfix -U /var/run/spamd.sock"
    CLAMSCAN="/usr/bin/clamscan"

    ORIG_MESS="$INSPECT_DIR/orig.$$"
    AV_OUT="$INSPECT_DIR/av_out.$$"
    SA_OUT="$INSPECT_DIR/sa_out.$$"
    TOADMIN_MESS="$INSPECT_DIR/to_admin.$$"

    VIRUSADMIN="postmaster@$MYDOMAIN"

    EX_TEMPFAIL=75
    EX_UNAVAILABLE=69

    trap "rm -f $ORIG_MESS $TOADMIN_MESS $SA_OUT $AV_OUT" 0 1 2 3 15




    cat > $ORIG_MESS || { echo Cannot save mail to file; exit $EX_TEMPFAIL; }
    $CLAMSCAN --no-summary --stdout $ORIG_MESS > $AV_OUT
    AV_RESULT=$?

    case "$AV_RESULT" in
    0)
      $FILTER_SPAMC < $ORIG_MESS > $SA_OUT || { echo Cannot save mail to file; exit $EX_TEMPFAIL; }
      $SENDMAIL "$@" < $SA_OUT
      exit 0
      ;;
    1)
      QUARANTINE_FILENAME=`date +%Y%m%d`_$$
      cp $ORIG_MESS $QUARANTINE_DIR/virus/$QUARANTINE_FILENAME

      echo "Subject: VIRUS FOUND" >> $TOADMIN_MESS
      echo >> $TOADMIN_MESS
      echo >> $TOADMIN_MESS
      grep Received $ORIG_MESS >> $TOADMIN_MESS
      echo "From: $2 (unverified)" >> $TOADMIN_MESS
      echo "To: $4" >> $TOADMIN_MESS
      grep Subject $ORIG_MESS >> $TOADMIN_MESS
      echo >> $TOADMIN_MESS
      cat $AV_OUT | awk '{print $2 " " $3}' >> $TOADMIN_MESS
      echo >> $TOADMIN_MESS
      echo "Saved as $QUARANTINE_DIR/virus/$QUARANTINE_FILENAME" >> $TOADMIN_MESS
        echo >> $TOADMIN_MESS

      $SENDMAIL -f $VIRUSADMIN -r $VIRUSADMIN -F "Antivirus" $VIRUSADMIN < $TOADMIN_MESS

      exit 0
      ;;
    *)
      echo "Subject: ANTIVIRUS FAILED" >> $TOADMIN_MESS
      echo >> $TOADMIN_MESS
      echo "Antivirus Failed with next problem:" >> $TOADMIN_MESS
      echo >> $TOADMIN_MESS
      case "$AV_RESULT" in
    40)
           echo "Unknown option passed." >> $TOADMIN_MESS
           ;;
        50)
           echo "Database initialization error." >> $TOADMIN_MESS
           ;;
        52)
           echo "Not supported file type." >> $TOADMIN_MESS
           ;;
        53)
           echo "Can't open directory." >> $TOADMIN_MESS
           ;;
        54)
           echo "Can't open file. (ofm)" >> $TOADMIN_MESS
           ;;
        55)
           echo "Error reading file. (ofm)" >> $TOADMIN_MESS
           ;;
        56)
           echo "Can't stat input file / directory." >> $TOADMIN_MESS
           ;;
        57)
           echo "Can't get absolute path name of current" >> $TOADMIN_MESS
           echo "working directory." >> $TOADMIN_MESS
           ;;
        58)
           echo "I/O error, please check your filesystem." >> $TOADMIN_MESS
           ;;
        59)
           echo "Can't get information about current user" >> $TOADMIN_MESS
           echo "from /etc/passwd." >> $TOADMIN_MESS
           ;;
        60)
           echo "Can't  get  information  about  user" >> $TOADMIN_MESS
           echo "clamav (default name) from /etc/passwd." >> $TOADMIN_MESS
           ;;
        61)
           echo "Can't fork." >> $TOADMIN_MESS
           ;;
        63)
           echo "Can't create temporary files/directories" >> $TOADMIN_MESS
           echo "(check permissions)." >> $TOADMIN_MESS
           ;;
        64)
           echo "Can't write to temporary directory (please" >> $TOADMIN_MESS
           echo "specify another one)." >> $TOADMIN_MESS
           ;;
        70)
           echo "Can't allocate and clear memory (calloc)." >> $TOADMIN_MESS
           ;;
        71)
           echo "Can't allocate memory (malloc)." >> $TOADMIN_MESS
           ;;
        *)
           echo "Unknown error $AV_RESULT" >> $TOADMIN_MESS
           ;;
      esac
           echo "************************************************" >> $TOADMIN_MESS
      $SENDMAIL -f $VIRADMIN -r $VIRADMIN -F "Antivirus" "$VIRADMIN" < $TOADMIN_MESS
      exit $EX_TEMPFAIL
      ;;
    esac

    exit 0

Данный вариант фильтра подходит для не очень загруженных почтовых
серверов (таких - большинство).

## Ссылки

[1](http://www.postfix.org/documentation.html)
[2](http://www.books.ru/shop/books/561452)

[Category:Администрирование](Category:Администрирование "wikilink")