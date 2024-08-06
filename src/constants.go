//this file stores constant values, such as file paths that need to be loaded
package main
//can be used to describe positions in 2d space with integers
type Vector2i struct {
	x int
	y int
}
//can be used to describe positions in 2d space with floats
type Vector2 struct {
	x float64
	y float64
}
//dimensions of the window
var WindowWidth int = 720
var WindowHeight int = 480
