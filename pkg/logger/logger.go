package logger

import (
	"log"
	"os"
)

var Log *log.Logger

func InitLogger() {
	// Проверка и создание папки логс
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}

	// Открываем или создаём лог файл
	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Could not open log file: %v", err)
	}

	// Инициализируем логгер
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
}

// Уровни логов

func Info(format string, v ...interface{}) {
	Log.Printf("INFO: "+format, v...)
}

func Error(format string, v ...interface{}) {
	Log.Printf("ERROR: "+format, v...)
}

func Debug(format string, v ...interface{}) {
	Log.Printf("DEBUG: "+format, v...)
}
