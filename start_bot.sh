#!/bin/bash

export GorillaToken=token
export GorillaType=url_verification

./gorilla_bot 8080 >> gorilla_bot.log &
