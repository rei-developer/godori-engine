package game

type GameMap struct {
	Name       string  `json: "name"`
	BGM        string  `json: "bgm"`
	Width      int     `json: "width"`
	Height     int     `json: "height"`
	Data       [][]int `json: "data"`
	Collisions []int   `json: "collisions"`
	Priorities []int   `json: "priorities"`
	Portals    []*Portal
}
