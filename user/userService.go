package user

type Player struct {
	player UserI
}

func (p Player) GetName(name string) string {
	return p.player.GetName(name)
}
func NewUser(p UserI) *Player {
	return &Player{
		player: p,
	}
}
