package config

import conf "github.com/quixote-liu/config"

func init() {
	LoadServerGroup()
}

func LoadServerGroup() {
	group := conf.NewGroup("server")

	group.SetString("host", "127.0.0.1")
	group.SetString("port", "8989")
	conf.CONF.RegisterGroup(group)
}
