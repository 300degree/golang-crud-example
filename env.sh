#! /bin/bash

rm -rf .env
touch .env

DB_CONFIG=( DB_HOST=$1 DB_PORT=$2 DB_USER=$3 DB_PASS=$4 DB_NAME=$5 )
if [ $# -eq 5 ]; then
    for i in ${!DB_CONFIG[@]}; do
        echo ${DB_CONFIG[$i]} >> .env
    done

    echo -e "generated .env file \e[1;92msuccessfully\e[0m\n"

else
    echo -e "Usage: sh ./env.sh <DB_HOST> <DB_PORT> <DB_USER> <DB_PASS> <DB_NAME>\n"
    echo -e "generated .env file \e[1;91mnot successfully\e[0m"

    echo "Usage: sh ./env.sh <DB_HOST> <DB_PORT> <DB_USER> <DB_PASS> <DB_NAME>" >> .env
    echo "Example: sh ./env.sh localhost 5432 postgres 123456 postgres" >> .env

fi