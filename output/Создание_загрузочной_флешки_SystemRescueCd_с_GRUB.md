Предоставлено пользователем
[GArik](http://www.linux.org.ru/people/GArik/profile)

В нашем примере на USB-накопитель будет установлен SystemRescueCD,
дистрибутив Tails и будет место для хранения файлов. Вы можете
использовать нижеприведенную инструкцию под свои нужды.

Для начала необходимо разбить USB-накопитель в на разделы следующим
образом:

    sdc3            Основной  ext4             [SystemRescueCD]       524,29
    sdc2            Основной  ext4             [Tails]           734,01
    sdc1            Основной  vfat             [GArik]         14971,57

Вы можете использовать для этого любую программу для работы с разделами
дисков. Я использовал cfdisk.

sdc3 - для SystemRescueCD, на него хватит где-нибудь 500Мб. Основной
раздел должен быть обязательно первым, иначе Windows может не
увидеть флешку при подключении. После разбиения флешки
форматируем разделы: основной в fat, а под систему лучше в
ext4 без журнала:

    # mkdosfs -n GArik /dev/sdc1
    # mke2fs -L Tails -T floppy -m 0 -t ext4 -O ^has_journal /dev/sdc2
    # mke2fs -L SystemRescueCD -T floppy -m 0 -t ext4 -O ^has_journal /dev/sdc3

Потом монтируем образ с SystemRescueCD и копируем содержимое на свежий
раздел:

    # mount /dev/sdc3 /media/SystemRescueCD/
    # mount -o loop SystemRescueCD/systemrescuecd-x86-2.5.0.iso /media/hd/
    # cp -a /media/hd/* /media/SystemRescueCD/

Затем нужно установить GRUB (лучше хотя бы версии 1.99):

    # grub-install --no-floppy --root-directory=/media/SystemRescueCD/ /dev/sdc

После установки grub'а появится новый каталог
/media/SystemRescueCD/boot/grub/. Туда нужно положить конфиг примерно
такого вида:

    if [ -s $prefix/grubenv ]; then
      load_env
    fi
    set default="0"
    if [ ${prev_saved_entry} ]; then
      set saved_entry=${prev_saved_entry}
      save_env saved_entry
      set prev_saved_entry=
      save_env prev_saved_entry
      set boot_once=true
    fi

    function savedefault {
      if [ -z ${boot_once} ]; then
        saved_entry=${chosen}
        save_env saved_entry
      fi
    }
    insmod ext2
    set root='(hd0,3)'
    set menu_color_normal=white/black
    set menu_color_highlight=blue/light-gray
    set timeout=5

    submenu "SystemRescueCD" {
        menuentry "SystemRescueCD x86" {
            linux  /isolinux/rescuecd scandelay=1 setkmap=ru
            initrd /isolinux/initram.igz
        }

        menuentry "SystemRescueCD x86_64" {
            linux  /isolinux/rescue64 scandelay=1 setkmap=ru
            initrd /isolinux/initram.igz
        }

        menuentry "SystemRescueCD alt x86" {
            linux  /isolinux/altker32 scandelay=1 setkmap=ru
            initrd /isolinux/initram.igz
        }

        menuentry "SystemRescueCD alt x86_64" {
            linux  /isolinux/altker64 scandelay=1 setkmap=ru
            initrd /isolinux/initram.igz
        }
    }

    menuentry "Tails" {
        set root='(hd0,2)'
        linux   /live/vmlinuz root=/dev/sda2 boot=live config noswap live-media=removable nopersistent noprompt keyboard-layouts=us,ru
        initrd  /live/initrd.img
    }

    menuentry "FreeDOS" {
        linux16  /isolinux/memdisk floppy
        initrd16 /bootdisk/freedos.img
    }

    menuentry "MemTest+" {
        linux16  /bootdisk/memtestp
    }

    menuentry "MHDD (HDD diagnostic tool)" {
        linux16  /isolinux/memdisk floppy
        initrd16 /bootdisk/mhdd.img
    }

    menuentry "HDT (Hardware Detection Tool)" {
        linux16  /isolinux/memdisk floppy
        initrd16 /bootdisk/hdt.img
    }

    menuentry "Boot from HDD" {
        exit
    }

    menuentry "Reboot" {
        reboot
    }

    menuentry "Halt" {
        halt
    }

menuentry можно добавить свои, по желанию. Вот и всё, мы получили
загрузочную USB-флешку с SystemRescueCD и Tails.