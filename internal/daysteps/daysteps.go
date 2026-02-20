package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	//
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid data format")
	}
	// Преобразую кол-во шагов в целое число
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid steps")
	}
	// Ошибка, кол-во шагов доожно быть положительным
	if steps <= 0{
		return 0, 0, fmt.Errorf("wrong steps value")
	}
	// Преобразую строку лительности
	dur, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, err
	}
	// Ошибка длительность должна быть положительной
	if dur <= 0 {
		return 0, 0, fmt.Errorf("wrond duration value")
	}
	return steps, dur, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// Разбираю входные данные
	steps, duration, err := parsePackage(data)
	if err != nil { // Возврат пустой строки при неккоректных данных.
		log.Println(err)
		return ""
	}
	if steps <= 0 { // доп проверка корректности кол-ва шагов.
		log.Println("invalid steps value")
		return ""
	}
	distanceMeters := float64(steps) * stepLength // Расчитываю пройденную дистанцию в метрах.
	distanceKm := distanceMeters / float64(mInKm) // перевел дистанцию в километры

	calories, err:= spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil{
		log.Println(err)
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceKm, calories)
}
