# The Go Programming Language

  - <https://golang.org/>
  - <https://godoc.org/>
  - <https://www.golang-book.com/>
  - <https://gobyexample.com/>
  - <https://github.com/miekg/gobook>

## Go Env

    mkdir -p ~/gocode/{bin,src,pkg}
    
    cat <<EOF>> ~/.profile
    if [ -d "$HOME/gocode" ]; then
            export GOPATH=$HOME/gocode
            export PATH=$PATH:$GOPATH/bin
    fi
    EOF

## Устанавливаем свежую версию для Debian/Ubuntu

    sudo apt-get install golang
    go get gopkg.in/niemeyer/godeb.v1/cmd/godeb
    sudo apt-get remove golang
    godeb install

Info: go из репозитория нужен только для bootstrap.

Альтернативно доступны так же готовые сборки инсталлятора для 32/64 бит:
смотри <https://github.com/niemeyer/godeb>

## Go-get'able Git repo и nginx

Если у вас есть свой Git репозиторий и вы хотите, что бы из него можно
было ставить пакеты по **go get**, то вот 2 примера, как можно
добиться этого при помощи nginx.

### CGit + https

Добавьте в конфиг cgit:

    clone-url=https://$HTTP_HOST$SCRIPT_NAME/$CGIT_REPO_URL

и в конфиг nginx:

    if ($arg_go-get = 1) {
            return 200 '<html><head><meta name="go-import" content="$host$uri git https://$host/cgi-bin/cgit.cgi$uri.git"/></head></html>\n';
    }

### GitLab + SSH

    if ($arg_go-get = 1) {
            return 200 '<html><head><meta name="go-import" content="$host$uri git ssh://git@$host$uri.git"/></head></html>\n';
    }

ref: <https://golang.org/cmd/go/#hdr-Remote_import_paths>
