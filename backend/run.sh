#!/usr/bin/env bash

echo "Start migrate..."
./gshop migrate up

echo "Start server..."
./gshop
