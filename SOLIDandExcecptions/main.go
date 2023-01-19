package main

import (
	"fmt"
	"os"
)

func main() {
	// указываем название файла, куда будут записываться логи
	logFileDestination := "logs.txt"
	// открываем этот файл, а если его нет, то создадим
	os.Remove(logFileDestination)
	logFile, _ := os.Create(logFileDestination)
	// создаем очередь для команд
	queue := Queue{}
	// добавим в очередь команду, приводящую к ошибке
	queue.Put(new(Command), 0)
	// запускаем цикл-исполнитель команд
	for {
		// достаем из очереди команду
		cmdData, err := queue.Take()
		// если очередь пуста, то завершаем программу
		if err != nil {
			fmt.Println("All commands are done, the queue is empty.")
			os.Exit(0)
		}
		// в go нет классического try except, поэтому проверяем функция вернула ошибку или нет
		if err = cmdData.cmd.Execute(); err != nil {
			// обрабатываем исключение, например, запишем его в лог
			ExceptionHandler5(cmdData, err, logFile, &queue)
		}
	}
}
