package main

import (
	"fmt"
	"os"
	"vyshkat/internal/di"
	"vyshkat/internal/ui"
)

func main() {
	container, err := di.BuildContainer()
	if err != nil {
		fmt.Fprintln(os.Stderr, "ошибка настройки приложения:", err)
		os.Exit(1)
	}
	if err := container.Invoke(func(cli *ui.CLI) { cli.Run() }); err != nil {
		fmt.Fprintln(os.Stderr, "ошибка запуска приложения:", err)
		os.Exit(1)
	}
}
