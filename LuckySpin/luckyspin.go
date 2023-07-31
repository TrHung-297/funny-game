package LuckySpin

import (
	"fmt"
	"github.com/gofiber/fiber"
	wr "github.com/mroth/weightedrand"
	"math/rand"
	"time"
)

func LuckySpin(ctx *fiber.Ctx) {
	// Khởi tạo danh sách các loại quả trong hộp

	var fruitId int
	// Khởi tạo seed ngẫu nhiên dựa trên thời gian để tăng tính ngẫu nhiên
	rand.Seed(time.Now().UnixNano())
	var fruitPick string
	var listOptions []wr.Choice
	for _, fruit := range fruits {
		if fruit.Quantity > 0 {
			choice := wr.Choice{
				Item:   fruit.Id,
				Weight: uint(fruit.Quantity),
			}
			listOptions = append(listOptions, choice)
		}
	}

	if len(listOptions) == 0 {
		fruitId = nonFruit.Id
	} else {
		c, _ := wr.NewChooser(listOptions...)
		fruitId = c.Pick().(int)
	}


	for _, v := range fruits {
		if fruitId == v.Id {
			fruitPick = fmt.Sprintf("quả %s",  v.Name)
			break
		} else if fruitId == nonFruit.Id {
			fruitPick = fmt.Sprintf("còn cái nịt quả")
		}
	}
	for j := 0; j < len(fruits); j++ {
		if fruits[j].Quantity > 0 && fruitId == fruits[j].Id {
			fruits[j].Quantity--
		}
	}
	WriteSuccess(ctx, map[string]interface{}{
		"Quả được chọn" : fruitPick,
		"Số lượng quả còn lại": fruits,
	})
}
