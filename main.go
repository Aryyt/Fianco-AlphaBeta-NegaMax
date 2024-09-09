package main

import (
	"fianco/gameLogic"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	"log"
)

type Sprite struct {
	Img  *ebiten.Image
	X, Y float64
}

type Game struct {
	GameBoard   *Sprite
	WhitePieces []*Sprite
	BlackPieces []*Sprite
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.WhitePieces = g.WhitePieces[0+1:]
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{100, 100, 100, 0xff})
	opts := &ebiten.DrawImageOptions{}

	opts.GeoM.Translate(g.GameBoard.X, g.GameBoard.Y)
	screen.DrawImage(g.GameBoard.Img, opts)
	opts.GeoM.Reset()

	for _, piece := range g.WhitePieces {
		opts.GeoM.Translate(piece.X, piece.Y)

		screen.DrawImage(piece.Img, opts)
		opts.GeoM.Reset()
	}
	for _, piece := range g.BlackPieces {
		opts.GeoM.Translate(piece.X, piece.Y)

		screen.DrawImage(piece.Img, opts)
		opts.GeoM.Reset()

	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}
func main() {
	test := gameLogic.InitializeState()
	fmt.Println(test.Board)
	fmt.Println(gameLogic.CheckForWinner(test))
	gameLogic.ChangePieceAtPosition(&test, gameLogic.Position{0, 0}, gameLogic.Position{0, 0}, 1)
	fmt.Println(test.Board)
	fmt.Println(gameLogic.CheckForWinner(test))

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Fianco")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	boardImage, _, err := ebitenutil.NewImageFromFile("assets/images/FiancoBoard.png")
	if err != nil {
		log.Fatal(err)
	}
	whiteImage, _, err := ebitenutil.NewImageFromFile("assets/images/FiancoPieceWhite.png")
	if err != nil {
		log.Fatal(err)
	}
	blackImage, _, err := ebitenutil.NewImageFromFile("assets/images/FiancoPieceBlack.png")
	if err != nil {
		log.Fatal(err)
	}

	whitePiecesSprites := make([]*Sprite, 0)
	whitePiecesPos := test.GetWhitePieces()
	for _, po := range whitePiecesPos {
		whitePiecesSprites = append(whitePiecesSprites, &Sprite{
			Img: whiteImage,
			X:   float64(po.X) * 50.0,
			Y:   float64(po.Y) * 50.0,
		})
	}

	blackPiecesSprites := make([]*Sprite, 0)
	blackPiecesPos := test.GetBlackPieces()
	for _, po := range blackPiecesPos {
		blackPiecesSprites = append(blackPiecesSprites, &Sprite{
			Img: blackImage,
			X:   float64(po.X) * 50.0,
			Y:   float64(po.Y) * 50.0,
		})
	}
	game := Game{
		GameBoard: &Sprite{
			Img: boardImage,
			X:   0.0,
			Y:   0.0,
		},
		WhitePieces: whitePiecesSprites,
		BlackPieces: blackPiecesSprites,
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}

}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
