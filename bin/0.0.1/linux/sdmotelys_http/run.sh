#!/bin/bash

URL=$1

current_time=$(date "+%Y.%m.%d-%H.%M.%S")

new_fileName=out.$current_time.xml

XML="$(curl 'http://'$URL'/WebTelys?Request=Mesure' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:73.0) Gecko/20100101 Firefox/73.0' -H 'Accept: */*' -H 'Accept-Language: ru-RU,ru;q=0.8,en-US;q=0.5,en;q=0.3' -H 'X-Requested-With: XMLHttpRequest' -H 'Connection: keep-alive' -H 'Referer: http://'$URL'/Conduit.asp' -H 'Cookie: Langue=En' -o $new_fileName -s)"


/mnt/data/mek/apps/sdmotelys_http/sdmotelys_http -filename=$new_fileName

rm $new_fileName -f
