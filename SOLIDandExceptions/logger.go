package main

import (
	"fmt"
	"os"
)

type ILogger interface {
	WriteLog(*os.File, string, error) ICommand
}

type Logs struct {
}

// Команда записи логов в файл

type WriteLogCommand struct {
	file *os.File
	log  string
	err  error
}

func (c *WriteLogCommand) Execute() error {
	_, err := c.file.Write([]byte(fmt.Sprintf("%s\nError message: %s\n\n", c.log, c.err.Error())))
	return err
}

// Создание экзмепляров команд

func (l Logs) WriteLog(logFile *os.File, log string, err error) ICommand {
	return &WriteLogCommand{file: logFile, log: log, err: err}
}
