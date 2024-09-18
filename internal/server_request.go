package internal

type ServerRequest interface {
	GetSerialNumber() string
	GetProtectedID() string
}

type ServerRequestImpl struct {
	SerialNumber string
	ProtectedID  string
}

func (s *ServerRequestImpl) GetSerialNumber() string {
	return s.SerialNumber
}

func (s *ServerRequestImpl) GetProtectedID() string {
	return s.ProtectedID
}
