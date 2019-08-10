Для не понимающих разницы между cat и ed:

    # cat <<EOF>> /etc/hostname.ath0
    rtsol
    EOF

добавляет в файл /etc/hostname.ath0 строку rtsol одной коммандой.

    # ed /etc/sysctl.conf <<EOF
    /net.inet6.icmp6.rediraccept/s/^#//
    w
    EOF

Заскриптенный ed ищет нужную строку и изменяет её, если она есть.

понял. извините. не знал о коментариях.
