package event

type EventConsumer interface {
	SendJSON(eventtype string, o interface{})
}
