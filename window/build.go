package window

/*
#cgo darwin CFLAGS:   -D_GOSMF_OSX_
#cgo darwin LDFLAGS:  -L/usr/local/lib -lSDL2

#cgo linux CFLAGS:    -D_GOSMF_LINUX_
#cgo linux LDFLAGS:   -lSDL2main -lSDL2

#cgo windows CFLAGS:  -D_GOSMF_WINDOWS_
#cgo windows LDFLAGS: -lSDL2main -lSDL2
*/
import "C"
