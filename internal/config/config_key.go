package config

type ConfigKey string

func (c ConfigKey) String() string {
	return string(c)
}
