#!/bin/bash
clear
echo "Recargando go..."
docker-compose stop
docker-compose up &
echo -n "Servidor iniciandose"
while ! httping -qc1 localhost:3000/api/ping
do
    echo -n "."
    sleep 1
done
echo "Go Iniciado🚀"
watch -d -t -g ls -lR &>/dev/null  && bash go_init.sh