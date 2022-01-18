package serverbox

type State struct {
	enabled bool
}

func InitializeState(sbc *SbContext) (state *State, err error) {
	err = nil
	state = new(State)
	return state, err
}
