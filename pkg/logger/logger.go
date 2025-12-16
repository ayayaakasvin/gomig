package logger

import (
	"log"
	"log/slog"
	"os"
)

func NewLogger(toLogAfter bool) *slog.Logger {
	logger, err := setupLogger(toLogAfter);
	if err != nil {
		log.Fatalf("failed to setup logger: %s", err.Error())
		return nil
	}

	return logger
}

func setupLogger(toLogAfter bool) (*slog.Logger, error) {
    textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
        Level: slog.LevelDebug,
    })

    jsonH, file, err := jsonHandler(toLogAfter)
    if err != nil {
        return nil, err
    }

    handlers := []slog.Handler{textHandler}
    if jsonH != nil {
        handlers = append(handlers, jsonH)
    }

    handler := &MultiHandler{
        handlers: handlers,
    }

    logger := slog.New(handler).With(
        "cli-tool", "gomig",
    )

	if file != nil {
		logger.Info("Log file", "filename", file.Name())
	}

    return logger, nil
}

func jsonHandler(toLogAfter bool) (*slog.JSONHandler, *os.File, error) {
    if !toLogAfter {
        return nil, nil, nil
    }

    file, err := os.CreateTemp("", "gomig-*.log")
    if err != nil {
        return nil, nil, err
    }

    jsonHandler := slog.NewJSONHandler(file, &slog.HandlerOptions{
        Level: slog.LevelInfo,
    })

    return jsonHandler, file, nil
}