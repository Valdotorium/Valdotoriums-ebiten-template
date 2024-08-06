#!/bin/sh
#this file clears all build folders
#clearing docs
cd docs
rm assets/*
rm *

cd ..
#clearing build
cd build
rm assets/*
rm *
cd ..