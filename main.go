package main

import (
	"log"
    "image/color"   
    _ "image/png"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
);

const screenWidth int = 800;
const screenHeight int = 600; 

var img *ebiten.Image;

// Create Game Struct 
type Game struct{};

func init() {
    var err error;
    img, _, err = ebitenutil.NewImageFromFile("./assets/visual/sprite/character/Cat_32.png");

    if err != nil {
        log.Fatal(err);
    }
}

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

    screen.DrawImage(img, nil);
}


// Layout 
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240;
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight);
	ebiten.SetWindowTitle("Test Ebiten GUI");
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err);
	}
}