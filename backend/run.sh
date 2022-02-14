#!/usr/bin/env bash

echo "Start migrate..."
./app migrate up

echo "Start server..."
./app
