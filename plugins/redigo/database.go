package redigo

type DatabaseInfo struct {
	Network string
	Addr    string
}

func (d *DatabaseInfo) Peer() string {
	return d.Addr
}
