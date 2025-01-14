package dtos

type CreateMinecraftServerDTO struct {
	OwnerEmail      string
	ServerName      string
	MaxPlayerAmount int
	ModesNames      []string
	PluginsNames    []string
	AllocatedRamMB  int
}
