## Как примонтировать память мобильного телефона Siemens в Linux?

  - Ставим пакеты fuse и fuse-dev для сборки драйвера файловой системы
  - Качаем siefs [отсюда](http://chaos.allsiemens.com/siefs/), и
    собираем. При сборке вылезает проблема линковки, чтобы
    исправить открываем Makefile, ищем **LDADD =
    $(fuseinst)/lib/libfuse.a -lpthread** и приводим ее к виду **LDADD =
    $(fuseinst)/lib/libfuse.a -lpthread -lrt -ldl**.
  - Монтируем:

<!-- end list -->

    mount -t siefs [-o options] /dev/ttyS0 /mnt/mobile

или, если телефон подключен через USB:

    mount -t siefs [-o options] /dev/ttyUSB0 /mnt/mobile

  - Наслаждаемся. Следует помнить, что по умолчанию файловые системы для
    FUSE доступны только для смонтировавшего их пользователя. Поэтому
    при монтировании от root надо добавить -o allow_other.

