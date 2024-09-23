package internal

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

func (s *ServerRequestImpl) GetProtectedID() string {
	return s.ProtectedId
}
