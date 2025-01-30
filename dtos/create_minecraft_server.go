package dtos

type CreateMinecraftServerDTO struct {
	OwnerEmail      string `json:"ownerEmail"`
	ServerName      string `json:"serverName"`
	MaxPlayerAmount int    `json:"maxPlayerAmount"`
	AllocatedRamMB  int    `json:"allocatedRamMB"`
}

type CreateMinecraftServerResponseDTO struct {
	Id int32 `json:"id"`
}
