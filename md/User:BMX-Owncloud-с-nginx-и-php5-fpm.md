    server {
        listen 80;
        server_name www.example.com example.com;
        return 301 https://$server_name$request_uri;
    }
    
    server {
        listen 443 ssl;
        server_name www.example.com example.com;
    
        root /usr/share/nginx/html;
        index index.php index.html index.htm;
    
        ssl on;
        ssl_certificate /etc/ssl/certs/www.example.com.pem;
        ssl_certificate_key /etc/ssl/private/www.example.com.key;
    
        access_log /var/log/nginx/access.log;
        error_log /var/log/nginx/error.log;
    
        rewrite ^/owncloud/caldav((/|$).*)$ /owncloud/remote.php/caldav$1 last;
        rewrite ^/owncloud/carddav((/|$).*)$ /owncloud/remote.php/carddav$1 last;
        rewrite ^/owncloud/webdav((/|$).*)$ /owncloud/remote.php/webdav$1 last;
    
        location ~ ^/owncloud/(data|config|\.ht|db_structure.xml|README) {
            deny all;
        }
    
        location /owncloud {
            root /usr/share/;
            rewrite ^/owncloud/.well-known/host-meta /public.php?service=host-meta last;
            rewrite ^/owncloud/.well-known/host-meta.json /public.php?service=host-meta-json last;
            rewrite ^/owncloud/.well-known/carddav /remote.php/carddav/ redirect;
            rewrite ^/owncloud/.well-known/caldav /remote.php/caldav/ redirect;
            rewrite ^/owncloud/apps/calendar/caldav.php /remote.php/caldav/ last;
            rewrite ^/owncloud/apps/contacts/carddav.php /remote.php/carddav/ last;
            rewrite ^/owncloud/apps/([^/]*)/(.*\.(css|php))$ /index.php?app=$1&getfile=$2 last;
            rewrite ^(/owncloud/core/doc[^\/]+/)$ $1/index.html;
            try_files $uri $uri/ index.php;
        }
    
        location ~ ^(?<script_name>.+?\.php)(?<path_info>/.*)?$ {
            root /usr/share/;
            try_files $script_name = 404;
            fastcgi_pass unix:/var/run/php5-fpm.sock;
            fastcgi_param PATH_INFO $path_info;
            fastcgi_param HTTPS on;
            # This one is a little bit tricky, you need to pass all parameters in a single line, separating them with newline (\n)
            fastcgi_param PHP_VALUE "upload_max_filesize = 1024M \n post_max_size = 1024M"; # This finishes the max upload size settings
            fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name; # On some systems OC will work without this setting, but it doesn't hurt to leave it here
            include fastcgi_params;
        }
    
        location ~/(\.ht|README|AUTHORS|INSTALL|LICENSE|CONFIG|ChangeLog) {
            deny all;
        }
    }
