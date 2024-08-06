HOW TO USE THIS TEMPLATE

CODING

-> Put all files that belong to the package main into the src directory
-> Put all files that belong to other packages into the src/pkg/[package-name] directory

BUILDING FOR WEB

-> Run this command on Mac/Linux:
    env GOOS=js GOARCH=wasm go build -o yourgame.wasm github.com/yourname/yourgame

   Or on Windows PowerShell:

    $Env:GOOS = 'js'
    $Env:GOARCH = 'wasm'
    go build -o yourgame.wasm github.com/yourname/yourgame
    Remove-Item Env:GOOS
    Remove-Item Env:GOARCH

   (you do not need to copy wasm_exec.js since it is already contained in the docs directory, if it does not work on Windows run:
   $goroot = go env GOROOT
   cp $goroot\misc\wasm\wasm_exec.js .)

   (Commands from https://ebitengine.org/en/documents/webassembly.html, accessed on 6th August 2024)

   put the generated wasm file into the docs directory and put your games assets into the docs/assets folder.

BUILDING FOR DESKTOP

-> just use go build . from the src folder.
   put the executable with the assets of your game into the build folder.

HELPER FUNCTIONS

Texture loading:
put all of your games image paths that are required into the designated lines in load.go
The images can be later accessed in Game.Textures

TouchIDs:
using this file: https://github.com/mevdschee/ebiten-mines/blob/main/touch/touch.go
interactions.go stores the position of the latest touch as the position of Game.Mouse
Touches and Left mouse clicks set the boolean isDown of Game.Mouse to true to detect mouse clicks / touches.
UpdateMouse() in main checks for touches / mouse clicks

Vectors:
in constants.go, structs for 2d Vectors are defined.

Game struct:
Predefines a game struct including a list of textures and a mouse.



