package main

import (
	"log"
    "image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
);

// Create Game Struct 
type Game struct{};

// Update
func (g *Game) Update() error {
	return nil;
}


// Draw
func (g *Game) Draw(screen *ebiten.Image) {

    // Fill Screen. 
    screen.Fill(color.RGBA{0xff, 0, 0xff, 0xa3});

    // Text Log. 
    messageText := "Debug Log Test.";
	ebitenutil.DebugPrint(screen, messageText);
}


// Layout 
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240;
}

func main() {
	ebiten.SetWindowSize(640, 480);
	ebiten.SetWindowTitle("Test Ebiten GUI");
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err);
	}
}