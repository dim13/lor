Дистрибутив Fedora (ранее Fedora Core) является свободным дистрибутивом,
который активно поддерживается компанией
[RedHat](http://www.redhat.com). В результате, именно в этом
дистрибутиве первыми появляются новшества, которые
впоследствии становятся "фактическим стандартом" промышленных
Linux-дистрибутивов.

RedHat Enterprise Linux - дистрибутив Linux компании RedHat, в настоящий
момент де-факто являющийся стандартом Linux для промышленного
использования. На его основе созданы открытые (бесплатные)
дистрибутивы [CentOS](http://www.centos.org/) и [Scientific
Linux](http://www.scientificlinux.org/)

[Официальный сайт компании RedHat](http://www.redhat.com/)

[Официальный сайт дистрибутива Fedora](http://fedoraproject.org/)

## Как назначить нескольким сетевым картам нужные алиасы при использовании демона network

Раньше за все псевдонимы (алиасы) отвечал файлик /etc/modprobe.conf (или
его инкарнация - каталог с одноименным названием). Теперь, в век
интерактивности, этим ведает udev.

Идем в /etc/udev/rules.d, открываем файл 70-persistent-net.rules и
меняем параметр NAME для каждой так, как нам нужно. Затем как
обычно прописываем параметры подключения карт в
/etc/sysconfig/network-scripts/ifcfg-ethX, где X - номер карты. Убираем
из загрузки NetworkManager, ставим вместо него network и перегружаемся.

## Как создать свой репозиторий для Fedora

Все очень просто:

1\. Создаем каталог на диске, в котором размещаем RPM-файлы нашего
репозитория

2\. Зайдя в указанный каталог, запускаем команду createrepo

    root@linux# createrepo .

3\. В каталоге /etc/yum.repos.d создаем файл myrepo.conf следующего
содержания

    [myrepo]
    name=My fedora repository
    baseurl=file:/home/ftp/f10/myrepo/
    enabled=1
    gpgcheck=0
    cost=100

## Проверка самодостаточности репозитория

Если вам необходимо создать "самодостаточный" репозиторий - например, вы
разворачиваете основной репозиторий во внутренней сети, с которого будут
устанавливаться пакеты на другие компьютеры, и при этом эти компьютеры
не должны выходить в интернет в основные репозитории
Fedora/CentOS/RHEL, можно воспользоваться командой repoclosure

repoclosure проведет проверку "замкнутости" всех активных в конфигурации
yum репозиториев, и выведет список зависимостей, которые необходимо
удовлетворить (например, докачать пакеты) чтобы получить
самодостаточный репозиторий.

## Интеграция "внешнего" JDK/JRE в систему

Если возникла необходимость использовать "внешний" JDK/JRE вместо
входящего в штатную поставку (например, для запуска Oracle
JDeveloper требуется Java2 SDK версии 1.6.0update4), можно
воспользоваться утилитой alternatives. Предположим, что
JDK установлен в каталог /usr/java/jdk1.6.0\_u11. Тогда, подадим
следующую команду:

    root@linux# alternatives --install /usr/bin/java java /usr/java/jdk1.6.0_11/jre/bin/java 3 \
        --slave /usr/bin/keytool keytool /usr/java/jdk1.6.0_11/jre/bin/keytool \
        --slave /usr/bin/rmiregistry rmiregistry /usr/java/jdk1.6.0_11/jre/bin/rmiregistry \
        --slave /usr/bin/rmid rmid /usr/java/jdk1.6.0_11/jre/bin/rmid \
        --slave /usr/bin/policytool policytool /usr/java/jdk1.6.0_11/jre/bin/policytool \
        --slave /usr/bin/tnameserv tnameserv /usr/java/jdk1.6.0_11/jre/bin/tnameserv \
        --slave /usr/bin/javac javac /usr/java/jdk1.6.0_11/jre/bin/javac \
        --slave /usr/bin/jmap jmap /usr/java/jdk1.6.0_11/jre/bin/jmap

После этого, можно переключить все приложения на использование
альтернативного JDK:

    root@linux# alternatives --config java
    There are 3 programs which provide 'java'.
      Selection    Command
    -----------------------------------------------
    *+ 1           /usr/lib/jvm/jre-1.6.0-openjdk.x86_64/bin/java
       2           /usr/lib/jvm/jre-1.5.0-gcj/bin/java
       3           /usr/java/jdk1.6.0_11/bin/java
    Enter to keep the current selection[+], or type selection number: 3
    root@linux# alternatives --config java
    There are 3 programs which provide 'java'.
      Selection    Command
    -----------------------------------------------
    *  1           /usr/lib/jvm/jre-1.6.0-openjdk.x86_64/bin/java
       2           /usr/lib/jvm/jre-1.5.0-gcj/bin/java
     + 3           /usr/java/jdk1.6.0_11/bin/java
    Enter to keep the current selection[+], or type selection number:

## Как получить графическую загрузку Fedora 10 на видеокартах NVidia?

Указать правильное значение параметра vga для ядра. Вот
[тут](http://dalth.livejournal.com/59451.html) перечислено
много-премного значений этого параметра для всяких разных
режимов. В тому числе, для режима 1280x800, выбирающего размер
экрана типичного современного ноутбука.

## Проблемы с PulseAudio в Fedora 11

[1](http://forums.fedoraforum.org/showthread.php?t=225660&highlight=pulse)
(позаимствовано
[здесь](http://www.linux.org.ru/view-message.jsp?msgid=3905302#comment-3907752),
мне помогло)

## Как отключить новое именование интерфейсов в Fedora15 и вернуться к ethX

Передать ядру параметр biosdevname=0 в загрузчике.

## Ссылки

[Неофициальный FAQ по Fedora Core](http://www.fedorafaq.org/) - хорошая
вещь, но на английском

[Часто Задаваемые Вопросы по Fedora](http://www.fedoracenter.ru)

[Сборник статей по настройке Fedora](http://www.mjmwired.net/resources/)
- отличный ресурс, где автор описывает "затачивание" Fedora после
установки. Не всё, конечно, следует перенимать, но есть и много
полезного. Описание к свежему выпуску Fedora обычно появляется в
течении недели.

[Category:Дистрибутивы](Category:Дистрибутивы "wikilink")
