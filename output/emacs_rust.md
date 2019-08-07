# Подсветка синтаксиса

Требуется установить emacs-пакет
[rust-mode](http://melpa.org/#/rust-mode) и дописать в файл
конфигурации:

    (autoload 'rust-mode "rust-mode" nil t)
    (add-to-list 'auto-mode-alist '("\\.rs\\'" . rust-mode))

# Автодополнение и поиск определения функции

Автодополнение осуществляется через библиотеку
[racer](https://github.com/phildawes/racer), перед ее компиляцией
необходимо установить:

  - Компилятор языка - [rust-lang](http://www.rust-lang.org/)

Желательно использовать последнюю ночную сборку, поскольку на данный
момент (02-05-2015) [racer](https://github.com/phildawes/racer)
отказывается компилироваться на более старых версиях

  - Пакетный менеджер для rust -
    [cargo](http://doc.crates.io/index.html)
  - Скачать исходные коды rust -
    [git](https://github.com/rust-lang/rust.git)

Для archlinux для установки можно воспользоваться готовыми бинарными
пакетами из AUR, с ними установка и получение исходников будет
выглядеть следующим образом:

    yaourt -S rust-nightly-bin
    yaourt -S racer-git
    git clone --recursive https://github.com/rust-lang/rust.git ~/.local/share/rust_src

Проверить правильность установки можно из командной строки указав
предварительно путь к скачанным исходным кодам rust:

    export RUST_SRC_PATH=~/.local/share/rust_src/src
    racer complete std::io::B

В ответ racer должен выдать список вариантов для автодополнения.

Для настройки автодополнения в emacs, необходим emacs-пакет
[company](http://melpa.org/#/company), после его установки и настройки
по необходимости, добавляем в файл конфигурации emacs строки:

    (setq racer-rust-src-path "~/.local/share/rust_src/src")
    (setq racer-cmd "/usr/bin/racer")
    ;; (add-to-list 'load-path "<path-to-racer>/editors")
    (eval-after-load "rust-mode" '(require 'racer))

Где закоментированная строка - указывает локальный путь к папке
[editors](https://github.com/phildawes/racer/tree/master/editors), если
racer был установлен вручную, ее необходимо раскоментировать и указать
актуальный путь, при установке racer через пакет, этот путь скорее
всего уже был добавлен в emacs.

После перезагрузки emacs должно заработать "go to the defenition" (по
умолчанию по M-.) и автодополнение (по умолчанию по Tab).

# Подсветка ошибок

Поиск ошибок во время редактирования кода - можно сделать либо через
emacs-пакет [flymake-rust](http://melpa.org/#/flymake-rust), либо при
помощи [flycheck-rust](http://melpa.org/#/flycheck-rust). Для
последнего достаточно добавить в конфиг:

    (eval-after-load 'flycheck
      '(add-hook 'flycheck-mode-hook #'flycheck-rust-setup))

# Компиляция и запуск

Компиляция одного исходного файла rust осуществляется командой "rustc",
а для проекта следует воспользоваться менеджером пакетов "cargo".
Приведенный ниже код, пытается найти файл проекта "Cargo.toml" в
текущей директории или выше по иерархии, если файл находится, то
компилируется и исполняется весь проект через "cargo", в противном
случае - компилируется и исполняется только текущий файл через "rustc".

    (defun lcl:rust-compile-hook ()
      (require 'compile)
      (set (make-local-variable 'compile-command)
           (if (locate-dominating-file (buffer-file-name) "Cargo.toml")
               "cargo run"
             (format "rustc %s && %s" (buffer-file-name)
                     (file-name-sans-extension (buffer-file-name))))))

    (setq-default compilation-read-command nil)
    (add-hook 'rust-mode-hook 'lcl:rust-compile-hook)

После добавления данного кода в файл инициализации, можно будет
запустить компиляцию на любом открытом файле rust вызвав
встроенную команду "compile" (M-x compile)

Более подробно про настройку можно прочитать
[тут](http://reangdblog.blogspot.com/2015/04/emacs-ide-rust.html)