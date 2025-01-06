package enums

type ConfigTypes string

const (
	ENV ConfigTypes = "env"
)

func (c ConfigTypes) String() string {
	return string(c)
}
