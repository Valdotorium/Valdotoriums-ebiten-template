#!/bin/sh
#this file builds the game for web and puts all necessary files into the docs folder
echo "Building wasm executable..."
cd src
env GOOS=js GOARCH=wasm go build -o main.wasm .
echo " 1/4 Building wasm executable done"
cp $(go env GOROOT)/misc/wasm/wasm_exec.js .
cd ..
echo " 2/4 Fetched JavaScript wasm runner"
#clearing docs
rmdir docs/assets
rm docs/*
echo " 3/4 Cleared assets"
#building the docs
mv src/main.wasm docs
cp src/assets docs
cp src/index.html docs
mv src/wasm_exec.js docs
cp src/assets docs
echo " 4/4 Collected and moved all necessary files into docs"
echo "Done!"


