package disk

var d = &disk{}

type disk struct{}

func GetDisk() *disk {
	return d
}
