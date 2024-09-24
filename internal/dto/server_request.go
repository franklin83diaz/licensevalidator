package dto

type ServerRequest interface {
	GetSerialNumber() string
	GetProtectedId() string
}

type ServerRequestImpl struct {
	SerialNumber string
	ProtectedId  string
}

func (s *ServerRequestImpl) GetSerialNumber() string {
	return s.SerialNumber
}

func (s *ServerRequestImpl) GetProtectedId() string {
	return s.ProtectedId
}

func NewServerRequest(serialNumber, protectedId string) ServerRequest {
	return &ServerRequestImpl{
		SerialNumber: serialNumber,
		ProtectedId:  protectedId,
	}
}
