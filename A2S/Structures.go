package A2S

type Address struct {
	IP   string
	Port uint16
}
type Info struct {
	Header     byte
	Protocol   byte
	Name       string
	GameMap    string
	Folder     string
	Game       string
	ID         uint16
	Players    byte
	MaxPlayers byte
	Bots       byte
	ServerType byte
	Enviroment byte
	Visibility byte
	Vac        byte
	Version    string
	EDF        byte
	GamePort   uint16
}
