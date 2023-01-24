package players

type Player struct {
	x, y           int32
	token, texture string
	tick           int64
}

func (player *Player) SetX(x int32) {
	player.x = x
}

func (player *Player) SetY(y int32) {
	player.y = y
}

func (player *Player) GetX() int32 {
	return player.x
}

func (player *Player) SetTick(tick int64) {
	player.tick = tick
}

func (player *Player) GetTick() int64 {
	return player.tick
}

func (player *Player) GetY() int32 {
	return player.y
}

func (player *Player) SetToken(token string) {
	player.token = token
}

func (player *Player) GetToken() string {
	return player.token
}

func (player *Player) SetTexture(texture string) {
	player.texture = texture
}

func (player *Player) GetTexture() string {
	return player.texture
}

func (player *Player) Update(x, y int32, texture, token string) {
	player.SetX(x)
	player.SetY(y)
	player.SetTexture(texture)
	player.SetToken(token)
	player.tick++
}
