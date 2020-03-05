#!/bin/bash

URL="localhost:1123"

XML="$(curl 'http://'$URL'/WebTelys?Request=Mesure' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:73.0) Gecko/20100101 Firefox/73.0' -H 'Accept: */*' -H 'Accept-Language: ru-RU,ru;q=0.8,en-US;q=0.5,en;q=0.3' --compressed -H 'X-Requested-With: XMLHttpRequest' -H 'Connection: keep-alive' -H 'Referer: http://'$URL'/Conduit.asp' -H 'Cookie: Langue=En')"

#XML="$(curl 'http://localhost:1123/WebTelys?Request=Mesure' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:73.0) Gecko/20100101 Firefox/73.0' -H 'Accept: */*' -H 'Accept-Language: ru-RU,ru;q=0.8,en-US;q=0.5,en;q=0.3' --compressed -H 'X-Requested-With: XMLHttpRequest' -H 'Connection: keep-alive' -H 'Referer: http://localhost:1123/Conduit.asp' -H 'Cookie: Langue=En')"

echo $XML