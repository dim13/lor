Оригинал [здесь](http://0pointer.de/blog/projects/the-biggest-myths).

С того момента, когда было предложено включить systemd в дистрибутивы,
прошло много обсуждений на различных форумах, в списках рассылки и на
конференциях. В этих обсуждениях можно было часто услышать некоторые
мифы о systemd, которые повторялись снова и снова, но определенно не
содержали ни капли правды. Давайте потратим немного времени на и�
развенчание:

1.  Миф: systemd монолитен.
    Если вы соберете systemd со всеми включенными опциями, у вас
    появится 69 различных исполняемых файлов. Они все служат
    различным целям и аккуратно разделены по ряду причин. \<...\> На
    самом деле, они так хорошо разделены, что оказываются полезны и при
    использовании вне systemd.
    Пакет с 69 различными двоичными файлами с трудом можно назвать
    монолитным. Однако отличие от предыдущих решений в том, что мы
    поставляем больше компонентов в одном архиве и поддерживаем их в
    едином репозитории с объединенным жизненным циклом.
2.  Миф: systemd заточен под скорость.
    Да, systemd быстр (загрузка достаточно полного пользовательского
    окружения за \~900ms, кто лучше?), но это в основном всего лишь
    побочный эффект верного подхода (doing things right). На самом деле
    мы никогда не старались выжать последнюю каплю производительности из
    systemd. Напротив, мы часто выбирали чуть более медленный код, чтобы
    сохранить его читаемость. Это не означает, что нам безразлична
    скорость загрузки, но сводить systemd лишь к ускорению
    загрузки - это недопонимание, поскольку это не было среди
    наших приоритетных целей.
3.  Миф: время загрузки не критично для промышленных серверов.
    Это совсем не так. Многим администраторам желательна быстрая
    загрузка во время обслуживания серверов. В конфигурация�
    высокой доступности тоже хорошо, когда сбоившая машина быстро
    возвращается в строй. В облачных конфигурациях с большим
    количеством виртуальных машин стоимость медленной загрузки
    умножается на количество ВМ. Трата времени процессора и
    ввода-вывода на медленной загрузке сотен ВМ или
    контейнеров радикально снижает доступную плотность машин,
    в конце концов, растет счет за электроэнергию. \<...\>
    Конечно, во многих схемах скорость однократной загрузки не играет
    роли, но systemd должен покрывать всю ширину возможны�
    применений. И да, я знаю, что часто основное время
    загрузки серверов занимает прошивка, и ОС по сравнению с
    этим достаточно быстра, но ладно, все равно systemd должен
    покрывать всю ширину применений, и кроме того, не у все�
    серверов настолько плохи прошивки, и конечно же не у ВМ и
    контейнеров, которые также в каком-то роде являются
    серверами.\[2\]
4.  Миф: systemd не совместим со скриптами оболочки.
    Это фикция. Мы просто не используем их в процессе загрузки, потому
    что мы верим, что они не являются лучшим средством для этой цели,
    но это не значит, что systemd с ними не совместим. Вы можете легко
    запускать скрипты в качестве служб systemd, да хоть на любом языке
    программирования, systemd не интересуется содержимым вашего
    исполняемого файла. Более того, мы много используем скрипты
    для наших собственных целей: для установки, сборки, тестирования
    systemd. И вы можете прикрепить ваши скрипты в процесс ранней
    загрузки, использовать их в качестве обычных служб, можете
    запускать их в последней стадии выключения машины, ограничений
    практически нет.
5.  Миф: systemd сложен.
    Это также абсолютно бессмысленно. Платформа systemd на самом деле
    гораздо проще традиционных линуксов, потому что объединяет
    системные объекты и их зависимости в юниты. Язик
    конфигурационных файлов очень прост, и мы избавились от
    избыточных конфигов. Мы предоставляем единые средства для основной
    конфигурации ситсемы. The system is much less conglomerate than
    traditional Linuxes are. У нас также есть достаточно исчерпывающая
    документация, доступная с главной страницы, почти о каждой детали
    systemd, и она покрывает не только средства работы для
    пользователей, но также и API для разработчиков.
    Конечно, systemd приходится изучить. Так во всем новом. Однако, нам
    хотелось бы верить, что systemd проще для понимания, чем основанная
    на shell загрузка. Удивлены? Как выясняется, shell не самый красивый
    для изучения язык, его синтаксис сложен и неочевиден. Файлы юнитов
    systemd существенно проще для понимания, они не содержат языка
    программирования, но просты и декларативны. После сказанного,
    если вы опытны в shell, то да, работа с systemd потребует немного
    изучения.

`   To make learning easy we tried hard to provide the maximum compatibility to previous solutions. But not only that, on many distributions you'll find that some of the traditional tools will now even tell you -- while executing what you are asking for -- how you could do it with the newer tools instead, in a possibly nicer way. `
`В любом случае, результат таков, что systemd, возможно, настолько прост, насколько вообще может быть проста подобная система, и что мы стараемся сделать его легким для изучения. Но если вы знаток sysvinit, то адаптация systemd потребует изучения, но откровенно говоря, если уж вы разобрались с sysvinit, то systemd не должен составить проблему для вас.`

1.  Миф: systemd не модульный.
    Совсем неверно. Во время компиляции у вас есть ряд настроек для
    выбора того, что нужно собирать, а что нет. И мы документируем,
    как вы можете выбрать еще более подробно нужные вам компоненты,
    используя более глубокие средства.
    Эта модульность не слишком сильно отличается от модульности ядра
    Linux, где вы можете выбрать многие параметры во время компиляции.
    Если ядро достаточно модульно для вас, то systemd недалеко от него
    ушел.
2.  Миф: systemd только для десктопов.
    Это, конечно, не так. With systemd we try to cover pretty much the
    same bandwith as Linux itself does. While we care for desktop uses,
    we also care pretty much the same way for server uses, and embedded
    uses as well. You can bet that Red Hat wouldn't make it a core piece
    of RHEL7 if it wasn't the best option for managing services on
    servers.

`   People from numerous companies work on systemd. Car manufactureres build it into cars, Red Hat uses it for a server operating system, and GNOME uses many of its interfaces for improving the desktop. You find it in toys, in space telescopes, and in wind turbines.`

`   Most features I most recently worked on are probably relevant primarily on servers, such as container support, resource management or the security features. We cover desktop systems pretty well already, and there are number of companies doing systemd development for embedded, some even offer consulting services in it.`

1.  Миф: systemd - проявление NIH-синдрома.

`   This is not true. Before we began working on systemd we were pushing for Canonical's Upstart to be widely adopted (and Fedora/RHEL used it too for a while). However, we eventually came to the conclusion that its design was inherently flawed at its core (at least in our eyes: most fundamentally, it leaves dependency management to the admin/developer, instead of solving this hard problem in code), and if something's wrong in the core you better replace it, rather than fix it. This was hardly the only reason though, other things that came into play, such as the licensing/contribution agreement mess around it. NIH wasn't one of the reasons, though...`

1.  Миф: systemd - проект freedesktop.org.

`   Well, systemd is certainly hosted at fdo, but freedesktop.org is little else but a repository for code and documentation. Pretty much any coder can request a repository there and dump his stuff there (as long as it's somewhat relevant for the infrastructure of free systems). There's no cabale involved, no "standardization" scheme, no project vetting, nothing. It's just a nice, free, reliable place to have your repository. In that regard it's a bit like SourceForge, github, kernel.org, just not commercial and without over-the-top requirements, and hence a good place to keep our stuff.`

`   So yes, we host our stuff at fdo, but the implied assumption of this myth in that there was a group of people who meet and then agree on how the future free systems look like, is entirely bogus.`

1.  Миф: systemd не UNIX.

`   There's certainly some truth in that. systemd's sources do not contain a single line of code originating from original UNIX. However, we drive inspiration from UNIX, and thus there's a ton of UNIX in systemd. For example, the UNIX idea of "everything is a file" finds reflection in that in systemd all services are exposed at runtime in a kernel file system, the cgroupfs. Then, one of the original features of UNIX was multi-seat support, based on built-in terminal support. Text terminals are hardly the state of the art how you interface with your computer these days however. With systemd we brought native multi-seat support back, but this time with full support for today's hardware, covering graphics, mice, audio, webcams and more, and all that fully automatic, hotplug-capable and without configuration. In fact the design of systemd as a suite of integrated tools that each have their individual purposes but when used together are more than just the sum of the parts, that's pretty much at the core of UNIX philosophy. Then, the way our project is handled (i.e. maintaining much of the core OS in a single git repository) is much closer to the BSD model (which is a true UNIX, unlike Linux) of doing things (where most of the core OS is kept in a single CVS/SVN repository) than things on Linux ever where.`

`   Ultimately, UNIX is something different for everybody. For us systemd maintainers it is something we drive inspiration from. For others it is a religion, and much like the other world religions there are different readings and understandings of it. Some define UNIX based on specific pieces of code heritage, others see it just as a set of ideas, others as a set of commands or APIs, and even others as a definition of behaviours. Of course, it is impossible to ever make all these people happy.`

`   Ultimately the question whether something is UNIX or not matters very little. Being technically excellent is hardly exclusive to UNIX. For us, UNIX is a major influence (heck, the biggest one), but we also have other influences. Hence in some areas systemd will be very UNIXy, and in others a little bit less.`

1.  Миф: systemd сложен.

`   There's certainly some truth in that. Modern computers are complex beasts, and the OS running on it will hence have to be complex too. However, systemd is certainly not more complex than prior implementations of the same components. Much rather, it's simpler, and has less redundancy (see above). Moreover, building a simple OS based on systemd will involve much fewer packages than a traditional Linux did. Fewer packages makes it easier to build your system, gets rid of interdependencies and of much of the different behaviour of every component involved.`

1.  Миф: systemd раздут.

`   Well, bloated certainly has many different definitions. But in most definitions systemd is probably the opposite of bloat. Since systemd components share a common code base, they tend to share much more code for common code paths. Here's an example: in a traditional Linux setup, sysvinit, start-stop-daemon, inetd, cron, dbus, all implemented a scheme to execute processes with various configuration options in a certain, hopefully clean environment. On systemd the code paths for all of this, for the configuration parsing, as well as the actual execution is shared. This means less code, less place for mistakes, less memory and cache pressure, and is thus a very good thing. And as a side-effect you actually get a ton more functionality for it...`

`   As mentioned above, systemd is also pretty modular. You can choose at build time which components you need, and which you don't need. People can hence specifically choose the level of "bloat" they want.`

`   When you build systemd, it only requires three dependencies: glibc, libcap and dbus. That's it. It can make use of more dependencies, but these are entirely optional.`

`   So, yeah, whichever way you look at it, it's really not bloated.`

1.  Миф: для BSD привязка systemd только к Linux - это плохо.

`   Completely wrong. The BSD folks are pretty much uninterested in systemd. If systemd was portable, this would change nothing, they still wouldn't adopt it. And the same is true for the other Unixes in the world. Solaris has SMF, BSD has their own "rc" system, and they always maintained it separately from Linux. The init system is very close to the core of the entire OS. And these other operating systems hence define themselves among other things by their core userspace. The assumption that they'd adopt our core userspace if we just made it portable, is completely without any foundation.`

1.  Миф: привязка systemd только к Linux делает невозможным для Debian
    сделать systemd менеджером загрузки по-умолчанию.

`   Debian supports non-Linux kernels in their distribution. systemd won't run on those. Is that a problem though, and should that hinder them to adopt system as default? Not really. The folks who ported Debian to these other kernels were willing to invest time in a massive porting effort, they set up test and build system, and patched and built numerous packages for their goal. The maintainance of both a systemd unit file and a classic init script for the packaged services is a negligable amount of work compared to that, especially since those scripts tend to already exist.`

1.  Миф: systemd можно было бы портировать на другие ядра, если бы
    только его разработчики захотели.

`   That is simply not true. Porting systemd to other kernel is not feasible. We just use too many Linux-specific interfaces. For a few one might find replacements on other kernels, some features one might want to turn off, but for most this is nor really possible. Here's a small, very incomprehensive list: cgroups, fanotify, umount2(), /proc/self/mountinfo (including notification), /dev/swaps (same), udev, netlink, the structure of /sys, /proc/$PID/comm, /proc/$PID/cmdline, /proc/$PID/loginuid, /proc/$PID/stat, /proc/$PID/session, /proc/$PID/exe, /proc/$PID/fd, tmpfs, devtmpfs, capabilities, namespaces of all kinds, various prctl()s, numerous ioctls, the mount() system call and its semantics, selinux, audit, inotify, statfs, O_DIRECTORY, O_NOATIME, /proc/$PID/root, waitid(), SCM_CREDENTIALS, SCM_RIGHTS, mkostemp(), /dev/input, ...`

`   And no, if you look at this list and pick out the few where you can think of obvious counterparts on other kernels, then think again, and look at the others you didn't pick, and the complexity of replacing them.`

1.  Миф: systemd беспричинно непереносим.

`   Non-sense! We use the Linux-specific functionality because we need it to implement what we want. Linux has so many features that UNIX/POSIX didn't have, and we want to empower the user with them. These features are incredibly useful, but only if they are actually exposed in a friendly way to the user, and that's what we do with systemd.`

1.  Миф: systemd использует двоичные конфигурационные файлы.

`   No idea who came up with this crazy myth, but it's absolutely not true. systemd is configured pretty much exclusively via simple text files. A few settings you can also alter with the kernel command line and via environment variables. There's nothing binary in its configuration (not even XML). Just plain, simple, easy-to-read text files.`

`   Myth: systemd is a feature creep.`

`   Well, systemd certainly covers more ground that it used to. It's not just an init system anymore, but the basic userspace building block to build an OS from, but we carefully make sure to keep most of the features optional. You can turn a lot off at compile time, and even more at runtime. Thus you can choose freely how much feature creeping you want.`

1.  Миф: systemd заставляет вас что-то делать.

`   systemd is not the mafia. It's Free Software, you can do with it whatever you want, and that includes not using it. That's pretty much the opposite of "forcing".`

1.  Миф: systemd делает невозможным использование syslog.

`   Not true, we carefully made sure when we introduced the journal that all data is also passed on to any syslog daemon running. In fact, if something changed, then only that syslog gets more complete data now than it got before, since we now cover early boot stuff as well as STDOUT/STDERR of any system service.`

1.  Миф: systemd несовместим.

`   We try very hard to provide the best possible compatibility with sysvinit. In fact, the vast majority of init scripts should work just fine on systemd, unmodified. However, there actually are indeed a few incompatibilities, but we try to document these and explain what to do about them. Ultimately every system that is not actually sysvinit itself will have a certain amount of incompatibilities with it since it will not share the exect same code paths.`

`   It is our goal to ensure that differences between the various distributions are kept at a minimum. That means unit files usually work just fine on a different distribution than you wrote it on, which is a big improvement over classic init scripts which are very hard to write in a way that they run on multiple Linux distributions, due to numerous incompatibilities between them.`

1.  Миф: systemd не скриптуем из-за использования D-Bus.

`   Not true. Pretty much every single D-Bus interface systemd provides is also available in a command line tool, for example in systemctl, loginctl, timedatectl, hostnamectl, localectl and suchlike. You can easily call these tools from shell scripts, they open up pretty much the entire API from the command line with easy-to-use commands.`

`   That said, D-Bus actually has bindings for almost any scripting language this world knows. Even from the shell you can invoke arbitrary D-Bus methods with dbus-send or gdbus. If anything, this improves scriptability due to the good support of D-Bus in the various scripting languages.`

1.  Миф: systemd требует от вас использования специальны�
    конфигурационных средств вместо того, чтобы разрешить
    править конфигурационные файлы напрямую.

`   Not true at all. We offer some configuration tools, and using them gets you a bit of additional functionality (for example, command line completion for all settings!), but there's no need at all to use them. You can always edit the files in question directly if you wish, and that's fully supported. Of course sometimes you need to explicitly reload configuration of some daemon after editing the configuration, but that's pretty much true for most UNIX services.`

1.  Миф: systemd нестабилен и полон ошибок.

`   Certainly not according to our data. We have been monitoring the Fedora bug tracker (and some others) closely for a long long time. The number of bugs is very low for such a central component of the OS, especially if you discount the numerous RFE bugs we track for the project. We are pretty good in keeping systemd out of the list of blocker bugs of the distribution. We have a relatively fast development cycle with mostly incremental changes to keep quality and stability high.`

1.  Миф: systemd не отлаживаем.

`   False. Some people try to imply that the shell was a good debugger. Well, it isn't really. In systemd we provide you with actual debugging features instead. For example: interactive debugging, verbose tracing, the ability to mask any component during boot, and more. Also, we provide documentation for it.`

`   It's certainly well debuggable, we needed that for our own development work, after all. But we'll grant you one thing: it uses different debugging tools, we believe more appropriate ones for the purpose, though.`

1.  Миф: systemd делает изменения ради изменений.

`   Very much untrue. We pretty much exclusively have technical reasons for the changes we make, and we explain them in the various pieces of documentation, wiki pages, blog articles, mailing list announcements. We try hard to avoid making incompatible changes, and if we do we try to document the why and how in detail. And if you wonder about something, just ask us!`

1.  Миф: systemd - проект только Red-Hat, является частной
    собственностью каких-то умнозадых разработчиков,
    которые используют его для продвижения своего взгляда на мир.

`   Not true. Currently, there are 16 hackers with commit powers to the systemd git tree. Of these 16 only six are employed by Red Hat. The 10 others are folks from ArchLinux, from Debian, from Intel, even from Canonical, Mandriva, Pantheon and a number of community folks with full commit rights. And they frequently commit big stuff, major changes. Then, there are 374 individuals with patches in our tree, and they too came from a number of different companies and backgrounds, and many of those have way more than one patch in the tree. The discussions about where we want to take systemd are done in the open, on our IRC channel (#systemd on freenode, you are always weclome), on our mailing list, and on public hackfests (such as our next one in Brno, you are invited). We regularly attend various conferences, to collect feedback, to explain what we are doing and why, like few others do. We maintain blogs, engage in social networks (we actually have some pretty interesting content on Google+, and our Google+ Community is pretty alive, too.), and try really hard to explain the why and the how how we do things, and to listen to feedback and figure out where the current issues are (for example, from that feedback we compiled this lists of often heard myths about systemd...).`

`   What most systemd contributors probably share is a rough idea how a good OS should look like, and the desire to make it happen. However, by the very nature of the project being Open Source, and rooted in the community systemd is just what people want it to be, and if it's not what they want then they can drive the direction with patches and code, and if that's not feasible, then there are numerous other options to use, too, systemd is never exclusive.`

`   One goal of systemd is to unify the dispersed Linux landscape a bit. We try to get rid of many of the more pointless differences of the various distributions in various areas of the core OS. As part of that we sometimes adopt schemes that were previously used by only one of the distributions and push it to a level where it's the default of systemd, trying to gently push everybody towards the same set of basic configuration. This is never exclusive though, distributions can continue to deviate from that if they wish, however, if they end-up using the well-supported default their work becomes much easier and they might gain a feature or two. Now, as it turns out, more frequently than not we actually adopted schemes that where Debianisms, rather than Fedoraisms/Redhatisms as best supported scheme by systemd. For example, systems running systemd now generally store their hostname in /etc/hostname, something that used to be specific to Debian and now is used across distributions.`

`   One thing we'll grant you though, we sometimes can be smart-asses. We try to be prepared whenever we open our mouth, in order to be able to back-up with facts what we claim. That might make us appear as smart-asses.`

`   But in general, yes, some of the more influental contributors of systemd work for Red Hat, but they are in the minority, and systemd is a healthy, open community with different interests, different backgrounds, just unified by a few rough ideas where the trip should go, a community where code and its design counts, and certainly not company affiliation.`

1.  Миф: systemd не поддреживает отдельный от корня /usr.

`   Non-sense. Since its beginnings systemd supports the --with-rootprefix= option to its configure script which allows you to tell systemd to neatly split up the stuff needed for early boot and the stuff needed for later on. All this logic is fully present and we keep it up-to-date right there in systemd's build system.`

`   Of course, we still don't think that actually booting with /usr unavailable is a good idea, but we support this just fine in our build system. This won't fix the inherent problems of the scheme that you'll encounter all across the board, but you can't blame that on systemd, because in systemd we support this just fine.`

1.  Миф: systemd не позволяет вам заменить его компоненты.

`   Not true, you can turn off and replace pretty much any part of systemd, with very few exceptions. And those exceptions (such as journald) generally allow you to run an alternative side by side to it, while cooperating nicely with it.`

1.  Миф: использование в systemd D-Bus вместо сокетов делает его
    непрозрачным.

`   This claim is already contradictory in itself: D-Bus uses sockets as transport, too. Hence whenever D-Bus is used to send something around, a socket is used for that too. D-Bus is mostly a standardized serialization of messages to send over these sockets. If anything this makes it more transparent, since this serialization is well documented, understood and there are numerous tracing tools and language bindings for it. This is very much unlike the usual homegrown protocols the various classic UNIX daemons use to communicate locally.`

Hmm, did I write I just wanted to debunk a "few" myths? Maybe these were
more than just a few... Anyway, I hope I managed to clear up a couple of
misconceptions. Thanks for your time.

Footnotes

\[1\] For example, systemd-detect-virt, systemd-tmpfiles, systemd-udevd
are.

\[2\] Also, we are trying to do our little part on maybe making this
better. By exposing boot-time performance of the firmware more
prominently in systemd's boot output we hope to shame the firmware
writers to clean up their stuff.