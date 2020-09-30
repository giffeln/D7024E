#!/bin/bash

amount=$1

docker-compose up --scale testrun=$1
