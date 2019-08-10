\[/usr broken\] \[Поттеринг\] \[Arch Linux\] \[Костыли\]

Если /usr нужен системе с самого начала, нужно перенести его
монтирование в initrd.

`Понадобится загрузочный диск, флешка с дистрибутивом Linux, желательно Archlinux. После загрузки с нее нужно примонтировать умершую систему. Например:`

    mount /dev/sdb8 /mnt
    mount /dev/sdb10 /mnt/usr
    mount /dev/sdb6 /mnt/boot
    mount -t proc none /mnt/proc
    mount -o bind /dev /mnt/dev
    mount -o bind /sys /mnt/sys

`Произвести chroot в систему:`

chroot /mnt

`Создать хук для mkinitcpio. Для этого создадим два файла: `

/lib/initcpio/install/usr

    #!/bin/bash
    build(){
     add_binary /bin/mount
     SCRIPT="usr"
    }

/lib/initcpio/hooks

    # vim:set ft=sh:
    run_hook ()
    {
     mount_handler=usr_mount_handler
    }
    
    usr_mount_handler(){<----->default_mount_handler "$@"
    <----->msg -n ":: Mounting /usr..."
    <----->cp /new_root/etc/fstab /etc/fstab
    <----->mount /usr
    <----->mount --move /usr /new_root/usr
    <----->msg "done."
    }

`Добавляем хук usr в /etc/mkinitcpio.conf в параметр HOOKS сразу после filesystems и пересобираем initrd командой`

mkinitcpio -p linux

В /etc/fstab у раздела /usr должна быть отключена автопроверка, то есть
последней цифрой в строке с ним должен быть 0.
