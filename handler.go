package mdns

type Handler interface {
	OnPeerFound(m Message, r Response)
	OnNetworkInterrupt()
} 

var defaultHandler Handler

func setExecutor(handlerImpl Handler) {
	defaultHandler = handlerImpl
}