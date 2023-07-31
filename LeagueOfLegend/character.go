package LeagueOfLegend

import (
	"github.com/gofiber/fiber"
)

type Level struct {
	LevelTft []RankTier `json:"LevelTft"`
}

type Character struct {
	Name string `json:"Name"`
	Id   int    `json:"Id"`
}

type RankTier struct {
	CostTier int `json:"CostTier"`
	Percent  int `json:"Percent"`
}

var mapQuantityByCost = map[int]int{
	1: 29,
	2: 22,
	3: 18,
	4: 12,
	5: 10,
}

var mapQuantityCharacter = map[int]int{}

var mapCharacterName = map[int]string{}

var listCostCharacter = map[int][]Character{
	1: {
		{Name: "Cassiopeia", Id: 1001},
		{Name: "Jhin", Id: 1002},
		{Name: "Kayle", Id: 1003},
		{Name: "Poppy", Id: 1004},
		{Name: "Samira", Id: 1005},
		{Name: "Cho'gath", Id: 1026},
		{Name: "Irelia", Id: 1027},
		{Name: "Mazaha", Id: 1028},
		{Name: "Maokai", Id: 1029},
		{Name: "Oriana", Id: 1030},
		{Name: "Renekton", Id: 1031},
		{Name: "Tristana", Id: 1032},
		{Name: "Viego", Id: 1033},
	},
	2: {
		{Name: "Galio", Id: 1006},
		{Name: "Sett", Id: 1007},
		{Name: "Teemo", Id: 1008},
		{Name: "Zed", Id: 1009},
		{Name: "Vi", Id: 1010},
		{Name: "Ashe", Id: 1034},
		{Name: "Jinx", Id: 1035},
		{Name: "Kasadin", Id: 1036},
		{Name: "Kled", Id: 1037},
		{Name: "Soraka", Id: 1038},
		{Name: "Swain", Id: 1039},
		{Name: "Talya", Id: 1040},
		{Name: "Warwick", Id: 1041},
	},
	3: {
		{Name: "Ekko", Id: 1011},
		{Name: "Garen", Id: 1012},
		{Name: "Sona", Id: 1013},
		{Name: "Taric", Id: 1014},
		{Name: "Karma", Id: 1015},
		{Name: "Ashan", Id: 1042},
		{Name: "Darius", Id: 1043},
		{Name: "Jayce", Id: 1044},
		{Name: "Kalista", Id: 1045},
		{Name: "Katarina", Id: 1046},
		{Name: "Lissandra", Id: 1047},
		{Name: "Rek'sai", Id: 1048},
		{Name: "Vel'kov", Id: 1049},
	},
	4: {
		{Name: "Gwen", Id: 1016},
		{Name: "Javan", Id: 1017},
		{Name: "Shen", Id: 1018},
		{Name: "Yasuo", Id: 1019},
		{Name: "Lux", Id: 1020},
		{Name: "Aphelios", Id: 1050},
		{Name: "Azir", Id: 1051},
		{Name: "Kai'sa", Id: 1052},
		{Name: "Nasus", Id: 1053},
		{Name: "Sejuani", Id: 1054},
		{Name: "Urgot", Id: 1055},
		{Name: "Zeri", Id: 1056},
	},
	5: {
		{Name: "Ryze", Id: 1021},
		{Name: "Sion", Id: 1022},
		{Name: "Ahri", Id: 1023},
		{Name: "Bel'veth", Id: 1024},
		{Name: "Senna", Id: 1025},
		{Name: "Aatrox", Id: 1057},
		{Name: "Heimerdinger", Id: 1058},
		{Name: "K'sante", Id: 1059},
	},
}

var level2 = Level{
	LevelTft: []RankTier{
		{CostTier: 1, Percent: 100},
		{CostTier: 2, Percent: 0},
		{CostTier: 3, Percent: 0},
		{CostTier: 4, Percent: 0},
		{CostTier: 5, Percent: 0},
	},
}

var level3 = Level{
	LevelTft: []RankTier{
		{CostTier: 1, Percent: 75},
		{CostTier: 2, Percent: 25},
		{CostTier: 3, Percent: 0},
		{CostTier: 4, Percent: 0},
		{CostTier: 5, Percent: 0},
	},
}

var level4 = Level{
	LevelTft: []RankTier{
		{CostTier: 1, Percent: 55},
		{CostTier: 2, Percent: 30},
		{CostTier: 3, Percent: 15},
		{CostTier: 4, Percent: 0},
		{CostTier: 5, Percent: 0},
	},
}

var level5 = Level{
	LevelTft: []RankTier{
		{CostTier: 1, Percent: 45},
		{CostTier: 2, Percent: 33},
		{CostTier: 3, Percent: 20},
		{CostTier: 4, Percent: 2},
		{CostTier: 5, Percent: 0},
	},
}

var level6 = Level{
	LevelTft: []RankTier{
		{CostTier: 1, Percent: 25},
		{CostTier: 2, Percent: 40},
		{CostTier: 3, Percent: 30},
		{CostTier: 4, Percent: 5},
		{CostTier: 5, Percent: 0},
	},
}

var level7 = Level{
	LevelTft: []RankTier{
		{CostTier: 1, Percent: 19},
		{CostTier: 2, Percent: 30},
		{CostTier: 3, Percent: 35},
		{CostTier: 4, Percent: 15},
		{CostTier: 5, Percent: 1},
	},
}

var level8 = Level{
	LevelTft: []RankTier{
		{CostTier: 1, Percent: 16},
		{CostTier: 2, Percent: 20},
		{CostTier: 3, Percent: 35},
		{CostTier: 4, Percent: 25},
		{CostTier: 5, Percent: 4},
	},
}

var level9 = Level{
	LevelTft: []RankTier{
		{CostTier: 1, Percent: 9},
		{CostTier: 2, Percent: 15},
		{CostTier: 3, Percent: 30},
		{CostTier: 4, Percent: 30},
		{CostTier: 5, Percent: 16},
	},
}

var mapLevelCharacter = map[int]Level{
	2: level2,
	3: level3,
	4: level4,
	5: level5,
	6: level6,
	7: level7,
	8: level8,
	9: level9,
}

type Response struct {
	Message string      `json:"Message,omitempty"`
	Data    interface{} `json:"Data,omitempty"`
	Status  int         `json:"Status,omitempty"`
}

func WriteSuccess(c *fiber.Ctx, v interface{}) error {
	res := Response{
		Message: "Success",
		Data:    v,
		Status:  200,
	}

	// Return
	return c.JSON(res)
}
