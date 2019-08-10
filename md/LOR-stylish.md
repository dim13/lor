Некоторые удобные настройки для плагина stylish, помогающие
по-человечески отображать ЛОР. Для того, чтобы настройки
действовали только на ЛОР, необходимо в "шапку" настроек
пользовательского стиля добавить

`   @-moz-document domain("linux.org.ru"){}`

(а в фигурных скобках - уже сами стили).

1.  Форматирование блока code:

`   code { border-left: 5px solid  !important;}`

1.  Форматирование цитат:

`   div.quote p cite{text-decoration: none !important; font-weight: bold !important;}`  
`   .quote{border: 1px dotted !important;}`

1.  Моноширинный шрифт в текстовой форме:

`   #msg{font-family: monospace !important; font-size: 12px !important;}`

1.  Форматирование меню в моем форке скрипта LOR-panel:

`   .SortTrackMenu{position: absolute;  margin: auto; background: none repeat scroll 0 0  #0000ff !important; text-align: left;}`  
`   .SortTrackMenuItem{left: 0px; margin: 1px; background-color: #c0c0c0 !important; color: black;}`

1.  перенос тегов в списке тем (треккере и форумах) в право, так что
    название темы выровнено по левому краю, а теги по правому

`   .message-table .tag { float: right; margin: 0 1px; }`

1.  лор шириной во весь экран (для темы tango)

`   body { max-width: 100% !important; }`

1.  подчистка мусора:

`   #bd div.messages div.infoblock {display: none !important;}`  
`   p[style*="margin-top: 0"] > em {display: none !important;}`  
`   #ft-back-button{display: none !important;}`  
`   a:hover{text-decoration: none !important;}`

1.  Приличное оформление подписи к посту:

`   .sign{margin-left: 0 !important; padding-bottom: 0 !important; margin-bottom: 0 !important;}`  
`   footer{padding-bottom: 0 !important; margin-bottom: 0 !important;}`

1.  Стандартный размер шрифтов (без ctrl-), до изменений на сайте.

`   body {font-size: medium !important;}`  
`   code {font-size: medium !important;}`  
`   pre {font-size: medium !important;}`  
`   label {font-size: medium !important;}`  
`   textarea {font-size: medium !important;}`  
`   button {font-size: medium !important;}`
