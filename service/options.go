package service

// Option describes a functional option for configuring the Client.
type Option func(*service)

// MaxReconnects sets max reconnect attempts for Client.
// 0 means reconnect forever.
func SetMethodVersion(methodName MethodName, version MethodVersion) Option {
	return func(s *service) {
		s.mm[methodName] = version
	}
}
