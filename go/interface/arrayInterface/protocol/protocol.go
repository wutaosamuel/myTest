package protocol

type Protocol interface {
	GetProtocol() Protocol
	Print()
}

// NOTE: ptl *Protocol cannot used in Protocol interface
func ProtocolPrint(ptl Protocol) {
	ptl.Print()
}
