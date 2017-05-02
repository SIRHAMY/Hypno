package hypno

type InputModule interface {
	// CreateStream() string
	InputInfo() string
	// StartInput()
	// StopInput()
}

type InputObject struct {
	Name      string
	Magnitude float64
}

type OutputModule interface {
	SetInputModule(*InputModule)
	Process(InputObject)
}
