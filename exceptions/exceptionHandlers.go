package exceptions

import (
	"fmt"
	"os"
)

// ExceptionWriteLog записывает инфу об ошибке в лог.
func ExceptionWriteLog(cmdData Element, err error, logFile *os.File, queue *Queue) error {
	err = queue.Put(Logs{}.WriteLog(logFile, fmt.Sprintf("Command failed, total retries: %d", cmdData.retrying), fmt.Errorf("%s", err.Error())), 0)
	return err
}

// ExceptionRetry повоторяет указанную команду, ставя ее в очередь.
func ExceptionRetry(cmdData Element, queue *Queue) error {
	err := queue.Put(Retry{}.Retry(queue, cmdData), cmdData.Retrying())
	return err
}

// ExceptionRetryAndWriteLog повторяет указанную команду и записывает лог при повторном исключении.
func ExceptionRetryAndWriteLog(cmdData Element, err error, logFile *os.File, queue *Queue) error {
	retrying := cmdData.Retrying()
	if retrying == 0 {
		ExceptionRetry(cmdData, queue)
	} else {
		// записываем инфу об ошибке в лог
		ExceptionWriteLog(cmdData, err, logFile, queue)
	}
	return nil
}

// ExceptionRetryTwiceAndWriteLog повторяет указанную команду дважды и записывает лог при получении исключения два раза подряд.
func ExceptionRetryTwiceAndWriteLog(cmdData Element, logFile *os.File, queue *Queue) error {
	retrying := cmdData.Retrying()
	if retrying == 0 {
		// повторяем команду первый раз
		queue.Put(Retry{}.Retry(queue, cmdData), retrying)
	} else if retrying == 1 {
		// второй раз
		queue.Put(SecondaryRetry{}.Retry(queue, cmdData), retrying)
	} else {
		// записываем инфу об ошибке в лог
		log := fmt.Sprintf("Command failed, total retries: %d", retrying)
		err := fmt.Errorf(SecondaryRetry{}.Error())
		queue.Put(Logs{}.WriteLog(logFile, log, err), 0)
	}
	return nil
}
