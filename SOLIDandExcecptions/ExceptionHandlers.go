package main

import (
	"fmt"
	"os"
)

func ExceptionHandler5(cmdData Element, err error, logFile *os.File, queue *Queue) error {
	err = queue.Put(Logs{}.WriteLog(logFile, fmt.Sprintf("Command failed, total retries: %d", cmdData.retrying), fmt.Errorf("%s", err.Error())), 0)
	return err
}

func ExceptionHandler7(cmdData Element, queue *Queue) error {
	err := queue.Put(Retry{}.Retry(queue, cmdData), cmdData.retrying)
	return err
}

func ExceptionHandler8(cmdData Element, err error, logFile *os.File, queue *Queue) error {
	if cmdData.retrying == 0 {
		// обработчик из 7 пункта
		ExceptionHandler7(cmdData, queue)
	} else {
		// обработчик исключения из 5-го пункта дз
		ExceptionHandler5(cmdData, err, logFile, queue)
	}
	return nil
}

func ExceptionHandler9(cmdData Element, logFile *os.File, queue *Queue) error {
	if cmdData.retrying == 0 {
		// повторяем команду первый раз
		queue.Put(Retry{}.Retry(queue, cmdData), cmdData.retrying)
	} else if cmdData.retrying == 1 {
		// второй раз
		queue.Put(SecondaryRetry{}.Retry(queue, cmdData), cmdData.retrying)
	} else {
		// записываем инфу об ошибке в лог
		log := fmt.Sprintf("Command failed, total retries: %d", cmdData.retrying)
		err := fmt.Errorf(SecondaryRetry{}.Error())
		queue.Put(Logs{}.WriteLog(logFile, log, err), 0)
	}
	return nil
}
