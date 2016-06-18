# gosmf [![GoDoc](https://godoc.org/github.com/anthonyrego/gosmf?status.svg)](https://godoc.org/github.com/anthonyrego/gosmf)

WIP golang simple media framework

## Requirements

Golang 1.5+

[SDL2](https://www.libsdl.org):
The window package requires SDL2 Development libraries/framework to be installed in order the compile. SDL2 is used for it's cross platform window management and input handling.

### Windows only

MinGW: Make sure you have the bin folder set in your PATH. Ensure you can run gcc and that the SDL2 include folder and libraries are in the Mingw include and lib folders. The arch will for mingw and sdl2 must match up in order to compile properly.

After successful compiling, you will need to have the SDL2.dll in the same folder as the .exe
