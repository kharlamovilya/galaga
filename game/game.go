package game

import (
	"awesomeProject/exceptions"
	"fmt"
)

func Start() {
	queue := &exceptions.Queue{}

	// Создаем корабль с начальным топливом initialFuelLevel
	initialFuelLevel := 100
	s := NewSpaceship(Vector2D{}, Vector2D{}, 0, initialFuelLevel)

	//Создаем макро команду для движения с сжиганием топлива, где кол-во сжигаемого топлива fuelConsumption
	fuelConsumption := 10
	moveAndBurn := new(MoveAndBurn).MoveAndBurnFuel(&s, &s, fuelConsumption)

	// Заготавливаем очередь, кладя туда 1 команду движения
	queue.Put(moveAndBurn, 0)

	fmt.Printf("Fuel left: %d\n", initialFuelLevel)

	// Начинаем длительное движение и движемся, допустим, пока есть топливо
	// Когда топливо кончится, то макро команда выбросит исключение, а цикл завершится
	for {
		// достаем команду
		cmdData, err := queue.Take()
		// проверяем, нормально ли досталась
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		// выполняем команду
		err = cmdData.Cmd().Execute()
		// проверяем, не выбросила ли она исключение
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		// Если нужная команда выполнилась, то кладем команду-повторюшку, чтобы реализовать длительное выполнение
		if cmdData.Cmd().GetName() == "MoveAndBurnCommand" {
			fuel, _ := s.Fuel()
			fmt.Printf("Fuel left: %d\n", fuel)
			// Кладем саму команду
			queue.Put(Repeat{}.Repeat(queue, cmdData), 0)
		}
	}
}
