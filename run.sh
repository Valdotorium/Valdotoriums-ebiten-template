#!/bin/sh
#this file build the go executable, reloads the assets and puts them into the build folder
cd src
echo " Running and building mac executable..."
#ls 
go build .
cd ..
echo " 1/4 Go compiled successfully..."
#reload assets
cd build
rm assets/*
rmdir assets
rm *
echo " 2/4 Removed old assets..."
cd ..
mv src/gobird build
cp -R src/assets build/assets
cd build
echo " 3/4 Copied assets..."
chmod +x gobird
#run file
echo " 4/4 Finished building"
./gobird
echo "Stopped executing"
