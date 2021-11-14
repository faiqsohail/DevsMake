#!/bin/bash

# run frontend node server
./node_modules/.bin/next start &

# start go swagger server
./main --host=0.0.0.0 --port=8080 && kill $!
