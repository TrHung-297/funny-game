package LeagueOfLegend

import (
	"fmt"
	"github.com/gofiber/fiber"
	"math/rand"
	"strconv"
	"time"
)

func GetRangesPercent(currentLV Level) []struct {
	start         int
	end           int
	costCharacter int
} {
	ranges := make([]struct {
		start         int
		end           int
		costCharacter int
	}, len(currentLV.LevelTft))
	percentStart := 0
	for i, tier := range currentLV.LevelTft {
		ranges[i].start = percentStart
		percentEnd := percentStart + tier.Percent
		ranges[i].end = percentEnd
		percentStart = percentEnd
		ranges[i].costCharacter = tier.CostTier
	}
	return ranges
}

func GenerateRandomArray() []int {
	rand.Seed(time.Now().UnixNano())
	array := make([]int, 5)
	for i := 0; i < 5; i++ {
		array[i] = rand.Intn(100) + 1
	}
	return array
}

func PickCharacter(levelUser int) []Character {
	// Khởi tạo seed ngẫu nhiên từ thời gian hiện tại
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	output := make([]Character, 0)
	var tuongRandom Character
	arrayRandom := GenerateRandomArray()
	fmt.Println("arrayRandom", arrayRandom)

	dataLevelCurrent := mapLevelCharacter[levelUser]
	rangePercentLevel := GetRangesPercent(dataLevelCurrent)
	fmt.Println("rangePercentLevel", rangePercentLevel)
	for i := 0; i < len(arrayRandom); i++ {
		for j := 0; j < len(rangePercentLevel); j++ {
			if arrayRandom[i] > rangePercentLevel[j].start && arrayRandom[i] <= rangePercentLevel[j].end {
				tuongRandom = RandomTuong(rangePercentLevel[j].costCharacter, randSource)
				break
			}
		}
		output = append(output, tuongRandom)
	}

	return output
}

func RandomTuong(costCharacter int, randSource *rand.Rand) Character {
	// Mảng gồm các tướng
	originalArray := listCostCharacter[costCharacter]

	// Lấy một số ngẫu nhiên từ 0 đến độ dài của mảng - 1
	randomIndex := randSource.Intn(len(originalArray))

	// Lấy phần tử từ mảng theo chỉ số ngẫu nhiên
	randomElement := originalArray[randomIndex]
	return randomElement
}

func ReRollCharacterLeagueOfLegend(levelUser int) ([]Character, error) {
	characterOutput := PickCharacter(levelUser)
	return characterOutput, nil
}

func ReRoll(c *fiber.Ctx) {
	userLevel, _ := strconv.ParseInt(c.Query("userLevel", "2"), 0, 64)
	messages, _ := ReRollCharacterLeagueOfLegend(int(userLevel))
	WriteSuccess(c, messages)
}

func Init() {
	for index, listCharacter := range listCostCharacter {
		for _, character := range listCharacter {
			mapQuantityCharacter[character.Id] = mapQuantityByCost[index]
			mapCharacterName[character.Id] = character.Name
		}
	}
}

func Pick(c *fiber.Ctx) {
	idCharacter, _ := strconv.ParseInt(c.Query("idCharacter", "0"), 0, 64)
	mapQuantityCharacter[int(idCharacter)]--
	WriteSuccess(c, map[string]interface{}{
		"Tướng pick là":       mapCharacterName[int(idCharacter)],
		"Số lượng còn lại là": mapQuantityCharacter[int(idCharacter)],
	})
}
