## Назначение

В данной статье описывается настройка Subversion для авторизации
пользователей через протокол LDAP. В качестве связующего звена
между двумя сервисами будет использоваться демон saslauthd,
предоставляющий простой уровень аутентификации и безопасности
- SASL.

## Целевая аудитория

Данное руководство расчитано на опытных системных администраторов и/или
опытных пользователей \*nix-систем, которые настраивали раз в жизни
subversion и openldap дома или на предприятии. Не смотря на такие
жесткие требования, обычный copy-paste в 95% случаев приведет к
желаемому результату. В статье акцент делается на настройку дву�
сервисов: subversion и saslauthd. Вопросы по администрированию и
управлению openldap не рассматриваются, т.к. выходят за рамки данной
статьи.

## Стендовое оборудование

Работоспособность методики проверена на двух системах: Fedora 14 и
Debian Squeeze. Ключевые моменты, коих немного, будут описаны по мере
необходимости. Тем не менее в других дистрибутивах GNU/Linux или ОС
придется догадываться в чем может быть проблема, когда ничего не
заработает.

## Установка ПО

Программное обеспечение ставится из репозиториев по умолчанию. При
самостоятельной сборки из исходников работоспособность метода не
гарантируется.

**Debian Squeeze**:

`# apt-get install subversion subversion-tools`
`# apt-get install ldap-utils slapd`
`# apt-get install libsasl2-modules-ldap sasl2-bin`
`# apt-get install libgsasl7 libsasl2-2 libsasl2-modules libsasl2-modules-gssapi-heimdal`

**Fedora**:

`# yum install openldap openldap-servers openldap-clients`
`# yum install cyrus-sasl cyrus-sasl-md5 cyrus-sasl-gssapi cyrus-sasl-lib`
`# yum install cyrus-sasl-ldap cyrus-sasl-plain`

## Подготовка

Первым делом вам необходимо настроить сервер openldap. К этому моменту
считаем, что вы это уже сделали. В кратце вам необходимо иметь хотя бы
одну учетную запись с objectClass organizationalPerson и/или
inetOrgPerson.

## Настройка saslauthd

Зададим параметры, по которым saslauthd будет искать и проверять
введенные данные от пользователей. Файл */etc/saslauthd.conf*

`ldap_servers: `<ldap://>**<SERVER IP ADDRESS>**`/`
`ldap_search_base: `**<search DN>**
`ldap_auth_method: bind`
`ldap_bind_dn: `**<bind DN>**
`ldap_bind_pw: `**<bind DN password>**
`ldap_filter: uid=%u`
`ldap_password_attr: userPassword`

Описание полей:

  - **<SERVER IP ADDRESS>** - IP адрес сервера openldap, например,
    localhost.
  - **<search DN>** - Корень каталога для поиска, например,
    dc=localdomain.
  - **<bind DN>** - учетная запись, у которой есть права на поиск в
    каталоге **<search DN>**. Опытным путем выяснилось, что ей не
    нужны права на чтение userPassword учетных записей пользователей.
  - **<bind DN password>** - пароль для **<bind DN>**

Более подробную информацию вы можете найти на страницах интернета,
например, здесь: [auth_ldap module for
saslauthd](http://www.opensource.apple.com/source/passwordserver_sasl/passwordserver_sasl-159/cyrus_sasl/saslauthd/LDAP_SASLAUTHD?txt)

Запускаем демон и проверяем, что все в порядке:

`# /etc/init.d/saslauthd restart`
`# testsaslauthd -u gh0stwizard -p my_super_password`
`0: OK "Success."`

Для Subversion также необходимо указать настройки, исходя из которы�
будет сделан выбран метод работы с saslauthd. Грубо говоря, это
опции для плагина sasl внутри subversion.

**Fedora** Файл */etc/sasl2/svn.conf*

`pwcheck_method: saslauthd`
`auxprop_plugin: ldap`
`mech_list: PLAIN`
`saslauthd_path: /var/run/saslauthd/mux`
`ldapdb_uri: `<ldap://>

**Debian Squeeze** использует иную конфигурацию. Файл
*/usr/lib/sasl2/svn.conf*

`pwcheck_method: saslauthd`
`auxprop_plugin: ldap`
`mech_list: PLAIN`
`saslauthd_path: /var/run/saslauthd/mux`
`ldapdb_uri: `<ldap://>

**Debian Wheezy** работает с *auxprop_plugin: ldapdb* Файл
*/usr/lib/sasl2/svn.conf*

`pwcheck_method: saslauthd`
`auxprop_plugin: ldapdb`
`mech_list: PLAIN`
`saslauthd_path: /var/run/saslauthd/mux`
`ldapdb_uri: `<ldap://>

## Настройка Subversion

Запускаем svnserve и создаем репозиторий:

`# svnserve --daemon --root /var/svn`
`# svnadmin create /var/svn/test_repo`

Файл */var/svn/test_repo/conf/svnserve.conf*

`[general]`
`anon-access = none`
`auth-access = write`
`authz-db = /var/svn/svn-access`

`[sasl]`
`use-sasl = true`
`min-encryption = 0`
`max-encryption = 256`

Остальные настройки в данном файле не играют никакой роли. Файл
*/var/svn/svn-access*

`[test_repo:/]`
`gh0stwizard = rw`

## Проверка и диагностика проблем

`# svn info `<svn://localhost/test_repo>` --username gh0stwizard`

`Path: test_repo`
`URL: `<svn://localhost/test_repo>
`Repository Root: `<svn://localhost/test_repo>
`Repository UUID: e1991669-3dcc-4b64-9eac-4219ddd374b6`
`Revision: 0`
`Node Kind: directory`
`Last Changed Rev: 0`
`Last Changed Date: 2012-05-27 10:08:45 -0400 (Sun, 27 May 2012)`

Наиболее частая проблема заключается в том, что плагин sasl из
subversion не находит корректный файл svn.conf. Симптомы:

1\. При запуске subversion, в логе /var/log/auth.log видно:

`May 27 10:47:51 tvv svnserve: auxpropfunc error invalid parameter supplied`
`May 27 10:47:51 tvv svnserve: _sasl_plugin_load failed on sasl_auxprop_plug_init for plugin: ldapdb`

2\. При попытке авторизации через svn, в том же файле:

`May 27 10:43:14 tvv svnserve: unable to open Berkeley db /etc/sasldb2: Permission denied`
`May 27 10:43:14 tvv svn: DIGEST-MD5 client step 2`

Для решения проблемы необходимо выяснить в какой именно директории
данная сборка subversion ищет файл svn.conf, как было указано
выше:

Debian использует /usr/lib/sasl2/svn.conf

Fedora использует /etc/sasl2/svn.conf

## Литература

1.  [Описание модуля auth_ldap для
    saslauthd](http://www.opensource.apple.com/source/passwordserver_sasl/passwordserver_sasl-159/cyrus_sasl/saslauthd/LDAP_SASLAUTHD?txt)
2.  [Тема на форуме: subversion +
    sasl](http://www.linux.org.ru/forum/development/7622495)
3.  [Тема на форуме: svnserver &
    sasl+ldap](http://www.linux.org.ru/forum/admin/7689522)