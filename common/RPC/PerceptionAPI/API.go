package PerceptionAPI

type API struct {
	Server *Server
}

func (a *API) ProcessPacket(value string, reply *string) error {
	a.Server.PacketChannel <- value
	return nil
}
