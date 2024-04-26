package packet

import "github.com/Wyze3306/gophertunnel/minecraft/protocol"

// EducationResourceURI is a packet that transmits education resource settings to all clients.
type EducationResourceURI struct {
	// Resource is the resource that is being referenced.
	Resource protocol.EducationSharedResourceURI
}

// ID ...
func (*EducationResourceURI) ID() uint32 {
	return IDEducationResourceURI
}

func (pk *EducationResourceURI) Marshal(io protocol.IO) {
	protocol.Single(io, &pk.Resource)
}
