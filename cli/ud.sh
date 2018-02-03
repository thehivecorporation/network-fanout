#!/bin/bash

echo -n "$1" | nc -u -c -w1 localhost "$2"
