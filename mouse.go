package lgtv

func (l *LgTV) MouseMove(x, y int) error {
	return l.command(command{
		Name: tvCmdMouseMove,
		X:    x,
		Y:    y,
	})
}
