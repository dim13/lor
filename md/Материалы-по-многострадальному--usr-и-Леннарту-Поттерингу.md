#### Основное

  - [Блог Поттеринга](http://0pointer.de/blog)
  - [Lennart's posts on
    LWN](http://www.google.com/search?q=site%3Alwn.net+mezcalero)
  - [systemd announcment](http://0pointer.de/blog/projects/systemd.html)
      - [Сравнение систем инициализации systemd, upstart и
        SysVinit](http://www.opennet.ru/opennews/art.shtml?num=30412)
      - [Introducing the
        Journal](https://docs.google.com/document/pub?id=1IC9yOXj7j6cdLLxWEBAGRL6wl97tFxgjLUEHIX3MSTs)
      - [systemd status update
        \#1](http://0pointer.de/blog/projects/systemd-update.html)
      - [systemd status update
        \#2](http://0pointer.de/blog/projects/systemd-update-2.html)
      - [systemd status update
        \#3](http://0pointer.de/blog/projects/systemd-update-3.html)

#### Freedesktop

  - [PulseAudio](http://www.freedesktop.org/wiki/Software/PulseAudio)
  - [The Case For The /usr
    Merge](http://www.freedesktop.org/wiki/Software/systemd/TheCaseForTheUsrMerge)
  - [Separate /usr is
    broken](http://www.freedesktop.org/wiki/Software/systemd/separate-usr-is-broken)
  - [systemd homepage](http://www.freedesktop.org/wiki/Software/systemd)
  - [Multi-Seat on Linux
    (systemd)](http://www.freedesktop.org/wiki/Software/systemd/multiseat)
  - [Wayland (non-Lennart with comparable
    butthurt)](http://wayland.freedesktop.org/)

#### LWN

  - [Introducing /run](http://lwn.net/Articles/436012/)
  - [Udev and systemd to merge (source tree
    only)](http://lwn.net/Articles/490413/)
  - [Choosing between portability and
    innovation](http://lwn.net/Articles/430598/)
    [(ист)](http://www.linux.org.ru/forum/talks/7584989)
  - [Systemd and ConsoleKit](https://lwn.net/Articles/441328/)
  - [LPC: Booting and systemd](https://lwn.net/Articles/458789/)
  - [The case for the /usr merge](https://lwn.net/Articles/477467/)
  - [Interview with Lennart Poettering
    (LinuxFR.org)](https://lwn.net/Articles/450299/)
  - [Quotes of the week](https://lwn.net/Articles/429695/)
  - [The logger meets linux-kernel](https://lwn.net/Articles/473999/)
  - [Fedora rejects Journal-only logging, ponders
    ARM](https://lwn.net/Articles/487342/)
  - [Upstart: please update to latest upstream
    version](https://lwn.net/Articles/484168/)
  - [Systemd & the tightly couple core band vs a world of many inits
    (discussion)](https://lwn.net/Articles/494095/)

#### Fedora

  - [Fedora: SysVinit to Systemd
    Cheatsheet](http://fedoraproject.org/wiki/SysVinit_to_Systemd_Cheatsheet)
  - [/usr move in Fedora
    (accepted)](https://fedoraproject.org/wiki/Features/UsrMove)
  - [Fedora: Systemd Journal as only default
    syslog](https://fedoraproject.org/wiki/Features/systemd-journal)
      - [Fedora 18: systemd-journal ONLY is
        rejected](https://fedorahosted.org/fesco/ticket/828),
        [подробности](http://meetbot.fedoraproject.org/fedora-meeting/2012-03-19/fesco.2012-03-19-18.00.log.html)
  - [Fedora: discussion about UsrMove
    feature](http://thread.gmane.org/gmane.linux.redhat.fedora.devel/155511/focus=155547)
  - [ConsoleKit Removal in
    Fedora](http://fedoraproject.org/wiki/Features/ckremoval)

#### Debian (є канонічним)

  - [Debian: /run in
    Wheezy](http://wiki.debian.org/ReleaseGoals/RunDirectory)
  - [Debian: 'Move all to /usr'
    discussion](http://thread.gmane.org/gmane.linux.debian.devel.general/165785)
  - [Debian: Deprecating /usr as a standalone
    filesystem?](http://lists.debian.org/debian-devel/2009/05/msg00075.html)
  - [Debian: no deprecation of /usr as a standalone
    filesystem](http://lists.debian.org/debian-devel/2009/05/msg00935.html)
  - [Debian: Read only partitions](http://wiki.debian.org/ReadonlyRoot)
  - [Debian: Installing systemd in
    Debian](http://wiki.debian.org/systemd)
  - [Debian mail list: A few observations about
    systemd](http://lists.debian.org/debian-devel/2011/07/msg00269.html)
  - [Debian debates systemd](https://lwn.net/Articles/452865/)

#### OpenSUSE

  - [Opensuse-factory: Proposal for 12.2, move all binaries under
    /usr](http://lists.opensuse.org/opensuse-factory/2011-11/msg01431.html)
  - [/usr merge in openSUSE](http://en.opensuse.org/openSUSE:Usr_merge)
  - [openSUSE. Зарисовки для новелл: возвращение
    sysvinit](http://fossbook.info/subproj/opensuse/1370)
  - [openSUSE: Vim-base-systemd Epic
    Win](http://rpmfind.net/linux/RPM/opensuse/factory/x86_64/vim-base-7.3.456-1.1.x86_64.html)
    (см. ниже)
    [(ист)](https://www.linux.org.ru/news/linux-general/7602530?cid=7604028)

<code>

`* Wed Feb 29 2012 mvyskocil@suse.cz`  
` - remove pointless systemd dependency and run the tmpfiles binary`  
`   only in case it exists`  
`...`  
`* Wed Dec 07 2011 aj@suse.de`  
` - Move require of systemd to base package since the base postinstall`  
`   needs it.`

</code>

#### Gentoo

  - [Gentoo: Locations of binaries and separate
    /usr](http://thread.gmane.org/gmane.linux.gentoo.devel/74464)
  - [Gentoo: Warn users not to do separate /usr partition without proper
    initramfs in the
    handbook?](http://thread.gmane.org/gmane.linux.gentoo.devel/72128)
  - [Gentoo: sys-fs/udev-171-r5 - problem running /usr/sbin/alsactl when
    /usr is not mounted](https://bugs.gentoo.org/show_bug.cgi?id=403493)
  - [Gentoo: udev 181 unmasking
    news](http://archives.gentoo.org/gentoo-dev/msg_690a54fc89fa435270f66c8585f118d1.xml)
  - [Gentoo: busybox(sep-usr) support for mounting /usr w/out hassle
    (sep /usr without
    initrd)](http://archives.gentoo.org/gentoo-dev/msg_978b2222cb87ec028fee65dc29eedbdb.xml)
  - [Gentoo: mdev wiki](https://wiki.gentoo.org/wiki/Mdev)

#### Arch

  - [Arch: Перенос /sbin в /usr/bin после обновления udev
    до 181](http://archlinux.org.ru/forum/viewtopic.php?f=6&t=8239)
  - [Arch: /usr на отдельном разделе.UPD проблема
    вернулась\!](http://archlinux.org.ru/forum/viewtopic.php?f=8&t=8133)
  - [Arch moves to
    systemd](https://mailman.archlinux.org/pipermail/arch-dev-public/2012-August/023389.html)
  - [Arch Linux прощается с
    /lib](https://www.linux.org.ru/news/gnome/8036367)
  - [Are We Removing What Defines Arch
    Linux?](http://allanmcrae.com/2012/08/are-we-removing-what-defines-arch-linux/)

#### Разное (w/ критика)

  - [Understanding the bin, sbin, usr/bin, usr/sbin
    split](http://lists.busybox.net/pipermail/busybox/2010-December/074114.html)
    и [Про каталоги в
    юниксе](http://www.muromec.org.ua/2012/02/blog-post.html)
  - [Lennart's advertisement](http://i.imgur.com/usftZ.png)
  - [systemd as external dependency
    (GNOME 3.2)](https://mail.gnome.org/archives/desktop-devel-list/2011-May/msg00427.html)
  - [Rewriting
    udisks](http://davidz25.blogspot.com/2012/03/simpler-faster-better.html)
      - [Udisks2: Another Loss For
        Linux](http://igurublog.wordpress.com/2012/03/11/udisks2-another-loss-for-linux/)
  - journald discussion
      - [Rainer Gerhards: journald and
        rsyslog](http://blog.gerhards.net/2011/11/journald-and-rsyslog.html)
      - [Rainer Gerhards: What I don't like about
        journald](http://blog.gerhards.net/2011/11/what-i-dont-like-about-journald.html)
      - [Rainer Gerhards: Serious syslog
        problems?](http://blog.gerhards.net/2011/11/serious-syslog-problems.html)
      - [Journald: just a reinvention of the
        wheel?](http://www.itwire.com/opinion-and-analysis/open-sauce/52112-journald-just-a-reinvention-of-the-wheel)
  - [Lennart upon your internets](http://rusty.ozlabs.org/?p=236)
  - [The systemd
    fallacy](http://monolight.cc/2011/05/the-systemd-fallacy/)
  - [RFC: User sessions in
    systemd](http://lists.freedesktop.org/archives/systemd-devel/2012-April/004974.html)
    и [Часть функциональности Gnome, KDE и Xfce переносят в
    systemd](http://russianfedora.ru/content/%D0%A7%D0%B0%D1%81%D1%82%D1%8C-%D1%84%D1%83%D0%BD%D0%BA%D1%86%D0%B8%D0%BE%D0%BD%D0%B0%D0%BB%D1%8C%D0%BD%D0%BE%D1%81%D1%82%D0%B8-gnome-kde-%D0%B8-xfce-%D0%BF%D0%B5%D1%80%D0%B5%D0%BD%D0%BE%D1%81%D1%8F%D1%82-%D0%B2-systemd)
  - [Lennart reply to
    Canonical](https://plus.google.com/115547683951727699051/posts/X3fUhyJREKq)
  - [Greg Kroah-Hartman on future of Linux and its
    ecosystem](https://plus.google.com/111049168280159033135/posts/V2t57Efkf1s)
  - [Lennart on
    procps-ng](http://www.phoronix.com/scan.php?page=news_item&px=MTEwMzQ)
  - [Идет работа по удалению виртуальных терминалов из ядра (
    CONFIG\_VT=n
    )](http://rfremix.ru/content/%D0%98%D0%B4%D0%B5%D1%82-%D1%80%D0%B0%D0%B1%D0%BE%D1%82%D0%B0-%D0%BF%D0%BE-%D1%83%D0%B4%D0%B0%D0%BB%D0%B5%D0%BD%D0%B8%D1%8E-%D0%B2%D0%B8%D1%80%D1%82%D1%83%D0%B0%D0%BB%D1%8C%D0%BD%D1%8B%D1%85-%D1%82%D0%B5%D1%80%D0%BC%D0%B8%D0%BD%D0%B0%D0%BB%D0%BE%D0%B2-%D0%B8%D0%B7-%D1%8F%D0%B4%D1%80%D0%B0-configvtn)
    и [kmscon FAQ](https://github.com/dvdhrm/kmscon/wiki/FAQ)
  - [GNOME (et al): Rotting In
    Threes](https://igurublog.wordpress.com/2012/11/05/gnome-et-al-rotting-in-threes/)
  - Цитаты:

`"As a trivial example: systemd creates user session information in
/run/user/$user. I brought up with lennart the fact that this would only
permit one session per user. He rejected out of hand the fact that more
than one session would ever be needed, because Gnome only allowed one
session per user".`

еще

`Polkit is an example of this: we have a patch to make systemd optional`
`at runtime, we request users to test it, and instead of testing it we`
`end up with a 300+ posts thread about how bad Lennart is, with nearly`
`no-one trying to investigate what is wrong about the patch and in
which` `situations it doesn't work.`

#### Linux.org.ru

  - [Разработчики дистрибутива Fedora Linux рассматривают возможность
    перемещения всех имеющихся в системе исполняемых файлов в одну
    директорию](http://www.linux.org.ru/forum/talks/6949377)
  - [Разработчики Fedora обсуждают объединение каталогов для исполняемых
    файлов](http://www.linux.org.ru/news/redhat/6950484)
  - [Для Fedora 17 утверждён план по переносу компонентов из корня в
    /usr и переход на
    Btrfs](http://www.linux.org.ru/news/redhat/7094857)
  - [Gentoo: прилетела
    новостишка](http://www.linux.org.ru/forum/talks/7541891)
  - [The Journal: жизнь после
    syslog](http://www.linux.org.ru/news/opensource/7029354)
  - [Поддержка /usr на отдельном
    разделе](http://www.linux.org.ru/forum/talks/7268879)
  - [Охота за зависимостями от
    /usr](http://www.linux.org.ru/forum/talks/7371996)
  - [Релиз systemd v38 c поддержкой Journal, замены системе
    syslog](http://www.linux.org.ru/news/opensource/7255258)
  - [Gentoo: initrd для монтирования
    /usr](http://www.linux.org.ru/forum/desktop/7557733)
  - [Слияние кодовой базы udev и
    systemd](http://www.linux.org.ru/news/linux-general/7602530)
  - [Новости systemd](http://www.linux.org.ru/forum/talks/7670147)
  - [Один из разработчиков GNOME осветил проблемы развития
    проекта](https://www.linux.org.ru/news/gnome/8036367)
  - [Xfce 4 будет окружением рабочего стола по умолчанию в Debian 7
    «Wheezy»](https://www.linux.org.ru/news/debian/8084439)

#### Немного о важных пакетах

  - **sysvinit** - System V style init programs original written by
    Miquel van Smoorenburg that control the booting and shutdown of your
    system. init (short for initialization) is a program for Unix-based
    computer operating systems that spawns all other processes. It runs
    as a daemon and typically has PID 1. The boot loader starts the
    kernel and the kernel starts init.

<!-- end list -->

  -   - **systemd** is a system and service manager for Linux,
        compatible with SysV and LSB init scripts. systemd provides
        aggressive parallelization capabilities, uses socket and D-Bus
        activation for starting services, offers on-demand starting of
        daemons, keeps track of processes using Linux cgroups, supports
        snapshotting and restoring of the system state, maintains mount
        and automount points and implements an elaborate transactional
        dependency-based service control logic. It can work as a drop-in
        replacement for sysvinit.

<!-- end list -->

  -   - **upstart** is an event-based replacement for the /sbin/init
        daemon which handles starting of tasks and services during boot,
        stopping them during shutdown and supervising them while the
        system is running. It was originally developed for the Ubuntu
        distribution, but is intended to be suitable for deployment in
        all Linux distributions as a replacement for the venerable
        System-V init.

<!-- end list -->

  - **initramfs-tools** - This package contains tools to create and boot
    an initramfs for prepackaged 2.6 Linux kernel. The initramfs is an
    cpio archive. At boot time, the kernel unpacks that archive into
    ram, mounts and uses it as initial root file system. From there on
    the mounting of the real root file system occurs in user space.
    klibc handles the boot-time networking setup. Supports nfs root
    system. Any boot loader with initrd support is able to load an
    initramfs archive. Having the root on MD, LVM2, LUKS or NFS is also
    supported.

<!-- end list -->

  -   - **dracut** - unlike existing initramfs's, this is an attempt at
        having as little as possible hard-coded into the initramfs as
        possible. The initramfs has (basically) one purpose in life --
        getting the rootfs mounted so that we can transition to the real
        rootfs. This is all driven off of device availability.
        Therefore, instead of scripts hard-coded to do various things,
        we depend on udev to create device nodes for us and then when we
        have the rootfs's device node, we mount and carry on. Having the
        root on MD, LVM2, LUKS is supported as well as NFS, iSCSI, NBD
        and FCOE with dracut-network.

<!-- end list -->

  - **udev** is the device manager for the Linux kernel. Primarily, it
    manages device nodes in /dev. It is the successor of devfs and
    hotplug, which means that it handles the /dev directory and all user
    space actions when adding/removing devices, including firmware load.

<!-- end list -->

  -   - ***hotplug** is a system for managing devices that can be
        dynamically attached to and removed from the system while it's
        running. The most obvious use for this system is handling USB
        and firewire devices, though it also handles PCI (32-bit PCMCIA
        - or CardBus - devices are really PCI in disguise), tape drives,
        SCSI devices, devices requiring firmware to be loaded into them,
        input devices, and more. Obsolete.*

<!-- end list -->

  -   - ***HAL** was a software project providing a hardware abstraction
        layer for Unix-like computer systems. It aimed to allow desktop
        applications to discover and use the hardware of the host system
        through a simple, portable and abstract API, regardless of the
        type of the underlying hardware. HAL is now deprecated on
        GNU/Linux systems, with functionality being merged into udev as
        of 2008–2010. HAL is in maintenance mode - no new features are
        added. All future development focuses on udisks, upower and
        other parts of the stack. HAL -\> udev/udisks/other*

<!-- end list -->

  - **Rsyslog** is an enhanced multi-threaded syslogd with a focus on
    security and reliability. Among others, it offers support for
    on-demand disk buffering, reliable syslog over TCP, SSL, TLS and
    RELP, writing to databases, email alerting. It is a drop-in
    replacement for syslogd.

<!-- end list -->

  - **D-Bus** is a message bus system, a simple way for applications to
    talk to one another. In addition to interprocess communication,
    D-Bus helps coordinate process lifecycle; it makes it simple and
    reliable to code a "single instance" application or daemon, and to
    launch applications and daemons on demand when their services are
    needed.

<!-- end list -->

  - **PolicyKit** is an application-level toolkit for defining and
    handling the policy that allows unprivileged processes to speak to
    privileged processes: It is a framework for centralizing the
    decision making process with respect to granting access to
    privileged operations for unprivileged applications. PolicyKit is
    specifically targeting applications in rich desktop environments on
    multi-user UNIX-like operating systems. It does not imply or rely on
    any exotic kernel features.

<!-- end list -->

  - **ConsoleKit** is a framework for defining and tracking users, login
    sessions, and seats. ConsoleKit is currently not actively
    maintained. The focus has shifted to the built-in seat/user/session
    management of Software/systemd called systemd-loginctl

<!-- end list -->

  - **udisks** project provides: a daemon, udisksd, that implements
    well-defined D-Bus interfaces that can be used to query and
    manipulate storage devices; a command-line tool, **udisksctl
    (udisks2)**, that can be used to query and use the daemon.

<!-- end list -->

  - **upower** is an abstraction for enumerating power devices,
    listening to device events and querying history and statistics.
    UPower aims to make a large chunk of HAL redundant, as HAL is
    officially deprecated. UPower is also useful to control the latency
    of different operations on your computer, which enables you to save
    significant amounts of power.

<!-- end list -->

  - **Avahi** is a system which facilitates service discovery on a local
    network via the mDNS/DNS-SD protocol suite. This enables you to plug
    your laptop or computer into a network and instantly be able to view
    other people who you can chat with, find printers to print to or
    find files being shared. Compatible technology is found in Apple
    MacOS X (branded Bonjour and sometimes Zeroconf).

<!-- end list -->

  - **PulseAudio** is a sound system for POSIX OSes, meaning that it is
    a proxy for your sound applications. It allows you to do advanced
    operations on your sound data as it passes between your application
    and your hardware. Things like transferring the audio to a different
    machine, changing the sample format or channel count and mixing
    several sounds into one are easily achieved using a sound server.

[category:Системное ПО](category:Системное_ПО "wikilink")
