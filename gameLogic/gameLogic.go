package gameLogic

const winingRowWhite = 0
const winingRowBlack = 8
const White = 1
const Black = 2

type State struct {
	Board       [9][9]int8
	WhitePlayer bool
	WhitePieces []Position
	BlackPieces []Position
}

func (s State) getPieceAtPosition(position Position) int8 {
	return s.Board[position.Y][position.X]
}
func (s State) GetWhitePieces() []Position {
	return s.WhitePieces
}
func (s State) GetBlackPieces() []Position {
	return s.BlackPieces
}
func (s State) GetBoard() [9][9]int8 {
	return s.Board
}
func InitializeState() State {
	var board [9][9]int8
	blackPieces := []Position{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0}, {8, 0}, {1, 1}, {2, 2}, {3, 3}, {7, 1}, {6, 2}, {5, 3}}
	whitePieces := []Position{{0, 8}, {1, 8}, {2, 8}, {3, 8}, {4, 8}, {5, 8}, {6, 8}, {7, 8}, {8, 8}, {1, 7}, {2, 6}, {3, 5}, {5, 5}, {6, 6}, {7, 7}}
	state := State{board, true, whitePieces, blackPieces}
	UpdateWhitePieces(&state)
	UpdateBlackPieces(&state)

	return state
}

func ChangePieceAtPosition(state *State, positionOld Position, positionNew Position, player int8) {
	state.Board[positionOld.Y][positionOld.X] = 0
	state.Board[positionNew.Y][positionNew.X] = player
}
func UpdateWhitePieces(s *State) {
	for _, piece := range s.WhitePieces {
		s.Board[piece.Y][piece.X] = White
	}
}
func UpdateBlackPieces(s *State) {
	for _, piece := range s.BlackPieces {
		s.Board[piece.Y][piece.X] = Black
	}
}
func CheckForWinner(s State) (bool, int8) {
	for i := 0; i < 9; i++ {
		if s.Board[winingRowWhite][i] == 1 {
			return true, White
		}
		if s.Board[winingRowBlack][i] == 2 {
			return true, Black
		}
	}
	return false, 0
}

type Position struct {
	X uint8
	Y uint8
}
