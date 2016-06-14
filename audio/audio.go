package audio

func Init() {
	initDevice()
}

func Cleanup() {
	destroyDevice()
}
