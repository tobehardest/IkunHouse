package uniqueid

type UniqueId interface {
	GetID() (id uint64, err error)
}
