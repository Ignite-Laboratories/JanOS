package std

// ChannelAction provides either the impulse context or an action to perform.
type ChannelAction struct {
	Context Context
	Action  func()
}
