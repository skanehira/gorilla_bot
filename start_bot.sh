#!/bin/bash

export URLVerificationToken=
export AuthorizationToken=

./gorilla_bot 8080 >> gorilla_bot.`date +%Y-%m-%d`.log &
