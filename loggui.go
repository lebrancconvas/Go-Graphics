package main

import (
	"log"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
);

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil;
	}

	messageLog := "Test Debug.";

	ebitenutil.DebugPrint(screen, messageLog);

	return nil;
}

func main() {
	appTitle := "Log GUI";

	if err := ebiten.Run(update, 400, 300, 2, appTitle); err != nil {
		log.Fatal(err);
	}
}