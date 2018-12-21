package salt

import "gobalt/src/transport"

type Salt struct {
	Auth      transport.Session
	Transport transport.Transport
}
