package audio

/*
#cgo darwin CFLAGS: -D_GOSMF_OSX_
#cgo darwin LDFLAGS: -framework OpenAL

#cgo darwin CFLAGS: -D_GOSMF_LINUX_
#cgo linux LDFLAGS: -lopenal

#cgo windows CFLAGS: -D_GOSMF_WINDOWS_
*/
import "C"
