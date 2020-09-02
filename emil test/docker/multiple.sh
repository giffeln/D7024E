#!/bin/bash

amount=$1

docker-compose up -d --scale testrun=$1