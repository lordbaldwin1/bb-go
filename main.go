package main

import (
	"fmt"
	"strings"
	"time"
)

const EMPTY_BOARD = "8/8/8/8/8/8/8/8 b - - "
const START_POSITION = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1 "
const TRICKY_POSITION = "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1 "
const KILLER_POSITION = "rnbqkb1r/pp1p1pPp/8/2p1pP2/1P1P4/3P3P/P1P1P3/RNBQKBNR w KQkq e6 0 1 "
const CMK_POSITION = "r2q1rk1/ppp2ppp/2n1bn2/2b1p3/3pP3/3P1NPP/PPP1NPB1/R1BQ1RK1 b - - 0 9 "

// board squares
const (
	a8 = 0
	b8 = 1
	c8 = 2
	d8 = 3
	e8 = 4
	f8 = 5
	g8 = 6
	h8 = 7
)

const (
	a7 = 8
	b7 = 9
	c7 = 10
	d7 = 11
	e7 = 12
	f7 = 13
	g7 = 14
	h7 = 15
)

const (
	a6 = 16
	b6 = 17
	c6 = 18
	d6 = 19
	e6 = 20
	f6 = 21
	g6 = 22
	h6 = 23
)

const (
	a5 = 24
	b5 = 25
	c5 = 26
	d5 = 27
	e5 = 28
	f5 = 29
	g5 = 30
	h5 = 31
)

const (
	a4 = 32
	b4 = 33
	c4 = 34
	d4 = 35
	e4 = 36
	f4 = 37
	g4 = 38
	h4 = 39
)

const (
	a3 = 40
	b3 = 41
	c3 = 42
	d3 = 43
	e3 = 44
	f3 = 45
	g3 = 46
	h3 = 47
)

const (
	a2 = 48
	b2 = 49
	c2 = 50
	d2 = 51
	e2 = 52
	f2 = 53
	g2 = 54
	h2 = 55
)

const (
	a1 = 56
	b1 = 57
	c1 = 58
	d1 = 59
	e1 = 60
	f1 = 61
	g1 = 62
	h1 = 63
)

const NO_SQ = 64

const WHITE = 0
const BLACK = 1
const BOTH = 2

const ROOK = 0
const BISHOP = 1

/*
	 Castling bits binary representation

	 0001	1	white king can castle to the king side
	 0010	2	white king can castle to the queene side
	 0100	4	black king can castle to the king side
	 1000	8	black king can castle to the queen side

	 examples

	 1111		both sides can castle both directions
	 1001		black king => queen side
				white king => king side
*/
// castling bits
const (
	WK = 1
	WQ = 2
	BK = 4
	BQ = 8
)

// encode pieces uppercase = white, lower = black
const (
	P = 0
	N = 1
	B = 2
	R = 3
	Q = 4
	K = 5
	p = 6
	n = 7
	b = 8
	r = 9
	q = 10
	k = 11
)

// ASCII pieces as a string
var asciiPieces = "PNBRQKpnbrqk"

// Unicode pieces as array
var unicodePieces = []string{
	"♙", "♘", "♗", "♖", "♕", "♔",
	"♟", "♞", "♝", "♜", "♛", "♚",
}

// convert ASCII character pieces to encoded constants
var charPieces = map[byte]int{
	'P': P, 'N': N, 'B': B, 'R': R, 'Q': Q, 'K': K,
	'p': p, 'n': n, 'b': b, 'r': r, 'q': q, 'k': k,
}

var promotedPieces = map[int]byte{
	Q: 'q', R: 'r', B: 'b', N: 'n', q: 'q', r: 'r', b: 'b', n: 'n',
}

var SquareToBigInt = []uint64{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47,
	48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63,
}

// for future use
var SquareToCoordinates = []string{
	"a8", "b8", "c8", "d8", "e8", "f8", "g8", "h8",
	"a7", "b7", "c7", "d7", "e7", "f7", "g7", "h7",
	"a6", "b6", "c6", "d6", "e6", "f6", "g6", "h6",
	"a5", "b5", "c5", "d5", "e5", "f5", "g5", "h5",
	"a4", "b4", "c4", "d4", "e4", "f4", "g4", "h4",
	"a3", "b3", "c3", "d3", "e3", "f3", "g3", "h3",
	"a2", "b2", "c2", "d2", "e2", "f2", "g2", "h2",
	"a1", "b1", "c1", "d1", "e1", "f1", "g1", "h1",
}

// states
var bitboards = [12]uint64{}
var occupancies = [3]uint64{}
var side int
var enpassant int = NO_SQ
var castle int

// copy of previous states
var bitboardsCopy [12]uint64
var occupanciesCopy [3]uint64
var sideCopy int
var enpassantCopy int
var castleCopy int

/*********************************************************\
===========================================================

                      Random numbers

===========================================================
\*********************************************************/

// pseudo random number state
var randomState uint64 = 1804289383

// generate 32-bit pseudo legal number
func getRandom32BitUnsignedNumber() uint32 {
	// get current randomState
	num := uint32(randomState)

	// XOR shift algorithm
	num ^= num << 13
	num ^= num >> 17
	num ^= num << 5

	// update random number randomState
	randomState = uint64(num)

	return num
}

// generate 64-bit pseudo legal numbers
func getRandom64BitUnsignedNumber() uint64 {
	// define 4 random numbers
	var n1, n2, n3, n4 uint64

	// init random numbers slicing 16 bits from MS1B side
	n1 = uint64(getRandom32BitUnsignedNumber() & 0xffff)
	n2 = uint64(getRandom32BitUnsignedNumber() & 0xffff)
	n3 = uint64(getRandom32BitUnsignedNumber() & 0xffff)
	n4 = uint64(getRandom32BitUnsignedNumber() & 0xffff)

	return n1 | (n2 << 16) | (n3 << 32) | (n4 << 48)
}

// generate magic number candidate
func generateMagicNumber() uint64 {
	n1 := getRandom64BitUnsignedNumber()
	n2 := getRandom64BitUnsignedNumber()
	n3 := getRandom64BitUnsignedNumber()
	return n1 & n2 & n3
}

/*********************************************************\
===========================================================

                    Bit manipulations

===========================================================
\*********************************************************/

// Bit manipulation functions for bitboards
func getBit(bitboard uint64, square int) uint64 {
	return bitboard & (1 << SquareToBigInt[square])
}

func setBit(bitboard uint64, square int) uint64 {
	return bitboard | (1 << SquareToBigInt[square])
}

func popBit(bitboard uint64, square int) uint64 {
	if getBit(bitboard, square) != 0 {
		bitboard &= ^(1 << square)
	}
	return bitboard
}

func countBits(bitboard uint64) int {
	count := 0

	for bitboard > 0 {
		// reset least significant first bit
		bitboard &= bitboard - 1
		count++
	}
	return count
}

func getLeastSignificantFirstBitIndex(bitboard uint64) int {
	if bitboard == 0 {
		return -1
	}

	return countBits((bitboard & (^bitboard + 1)) - 1)
}

/*********************************************************\
===========================================================

                    Input/Output

===========================================================
\*********************************************************/

func printBitboard(bitboard uint64) {
	fmt.Println()

	for rank := range 8 {
		for file := range 8 {
			// convert file & rank into square index
			square := rank*8 + file

			// print ranks
			if file == 0 {
				fmt.Printf("  %d  ", 8-rank)
			}
			// print bit state (either 1 or 0)
			if getBit(bitboard, square) != 0 {
				fmt.Print(" 1 ")
			} else {
				fmt.Print(" 0 ")
			}
		}
		fmt.Println()
	}

	// print board files
	fmt.Println("\n      a  b  c  d  e  f  g  h ")

	// print bitboard as unsigned decimal number
	fmt.Printf("\n      Bitboard: %d\n", bitboard)
}

// print board
func printBoard() {
	fmt.Println()
	for rank := range 8 {
		for file := range 8 {
			square := rank*8 + file

			if file == 0 {
				fmt.Printf("  %d ", 8-rank)
			}
			piece := -1

			// loop over all piece bitboards
			for bbPiece := P; bbPiece <= k; bbPiece++ {
				if getBit(bitboards[bbPiece], square) != 0 {
					piece = bbPiece
				}
			}

			if piece == -1 {
				fmt.Printf("  %c", '.')
			} else {
				fmt.Printf("  %c", asciiPieces[piece])
			}
		}
		fmt.Println()
	}
	fmt.Println("\n      a  b  c  d  e  f  g  h")

	fmt.Println()
	if side != 1 {
		fmt.Println("      Side:      white")
	} else {
		fmt.Println("      Side:      black")
	}

	if enpassant != NO_SQ {
		fmt.Println("      Enpassant:   ", SquareToCoordinates[enpassant])
	} else {
		fmt.Println("      Enpassant:  none")
	}

	fmt.Printf("      Castling:   ")

	if castle&WK != 0 {
		fmt.Print("K")
	} else {
		fmt.Print("-")
	}
	if castle&WQ != 0 {
		fmt.Print("Q")
	} else {
		fmt.Print("-")
	}
	if castle&BK != 0 {
		fmt.Print("k")
	} else {
		fmt.Print("-")
	}
	if castle&BQ != 0 {
		fmt.Print("q")
	} else {
		fmt.Print("-")
	}
	fmt.Print("\n\n")
}

// add handling of malformed strings?
func parseFEN(fen string) {
	// reset board position and state variables
	for i := range bitboards {
		bitboards[i] = 0
	}
	for i := range occupancies {
		occupancies[i] = 0
	}
	side = WHITE
	enpassant = NO_SQ
	castle = 0

	// setup board
	i, square := 0, 0
	for ; i < len(fen) && square < 64; i++ {
		if fen[i] == '/' {
			continue
		} else if isAlpha(fen[i]) {
			piece := charPieces[fen[i]]
			bitboards[piece] = setBit(bitboards[piece], square)

			// set occupancies
			if isUpperCase(fen[i]) {
				occupancies[WHITE] = setBit(occupancies[WHITE], square)
			} else {
				occupancies[BLACK] = setBit(occupancies[BLACK], square)
			}
			occupancies[BOTH] = setBit(occupancies[BOTH], square)
			square++
		} else if isNumeric(fen[i]) {
			empty := int(fen[i] - '0')
			square += empty
		}
	}

	// starting side
	for ; fen[i] == ' '; i++ {
	}
	splitFen := strings.Split(fen[i:], " ")

	if splitFen[0] == "w" {
		side = WHITE
	} else {
		side = BLACK
	}

	// castling rights
	for _, c := range splitFen[1] {
		if c == '-' {
			break
		}

		switch c {
		case 'K':
			castle |= WK
		case 'Q':

			castle |= WQ
		case 'k':
			castle |= BK
		case 'q':

			castle |= BQ
		default:
			continue
		}
	}

	// enpassant square
	if splitFen[2] != "-" {
		file := int(splitFen[2][0] - 'a')
		rank := 8 - int(splitFen[2][1]-'0')
		enpassant = rank*8 + file
	}
}

func isUpperCase(character byte) bool {
	return (character >= 'A' && character <= 'Z')
}

func isAlpha(character byte) bool {
	return (character >= 'a' && character <= 'z') || (character >= 'A' && character <= 'Z')
}

func isNumeric(character byte) bool {
	return character >= '0' && character <= '9'
}

/*********************************************************\
===========================================================

                        Attacks

===========================================================
\*********************************************************/

/*
          not A file

  8   0  1  1  1  1  1  1  1
  7   0  1  1  1  1  1  1  1
  6   0  1  1  1  1  1  1  1
  5   0  1  1  1  1  1  1  1
  4   0  1  1  1  1  1  1  1
  3   0  1  1  1  1  1  1  1
  2   0  1  1  1  1  1  1  1
  1   0  1  1  1  1  1  1  1

      a  b  c  d  e  f  g  h

          not H file

  8   1  1  1  1  1  1  1  0
  7   1  1  1  1  1  1  1  0
  6   1  1  1  1  1  1  1  0
  5   1  1  1  1  1  1  1  0
  4   1  1  1  1  1  1  1  0
  3   1  1  1  1  1  1  1  0
  2   1  1  1  1  1  1  1  0
  1   1  1  1  1  1  1  1  0

      a  b  c  d  e  f  g  h


          not HG file
  8   1  1  1  1  1  1  0  0
  7   1  1  1  1  1  1  0  0
  6   1  1  1  1  1  1  0  0
  5   1  1  1  1  1  1  0  0
  4   1  1  1  1  1  1  0  0
  3   1  1  1  1  1  1  0  0
  2   1  1  1  1  1  1  0  0
  1   1  1  1  1  1  1  0  0

      a  b  c  d  e  f  g  h

          not AB file

  8   0  0  1  1  1  1  1  1
  7   0  0  1  1  1  1  1  1
  6   0  0  1  1  1  1  1  1
  5   0  0  1  1  1  1  1  1
  4   0  0  1  1  1  1  1  1
  3   0  0  1  1  1  1  1  1
  2   0  0  1  1  1  1  1  1
  1   0  0  1  1  1  1  1  1

      a  b  c  d  e  f  g  h
*/

// All zero file constants
const NOT_A_FILE = uint64(18374403900871474942)
const NOT_H_FILE = uint64(9187201950435737471)
const NOT_HG_FILE = uint64(4557430888798830399)
const NOT_AB_FILE = uint64(18229723555195321596)

// Relevant occupancy bit count for every square on board
var bishopRelevantBits = []int{
	6, 5, 5, 5, 5, 5, 5, 6, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7, 7, 7, 7, 5, 5, 5, 5,
	7, 9, 9, 7, 5, 5, 5, 5, 7, 9, 9, 7, 5, 5, 5, 5, 7, 7, 7, 7, 5, 5, 5, 5, 5, 5,
	5, 5, 5, 5, 6, 5, 5, 5, 5, 5, 5, 6,
}

var rookRelevantBits = []int{
	12, 11, 11, 11, 11, 11, 11, 12, 11, 10, 10, 10, 10, 10, 10, 11, 11, 10, 10,
	10, 10, 10, 10, 11, 11, 10, 10, 10, 10, 10, 10, 11, 11, 10, 10, 10, 10, 10,
	10, 11, 11, 10, 10, 10, 10, 10, 10, 11, 11, 10, 10, 10, 10, 10, 10, 11, 12,
	11, 11, 11, 11, 11, 11, 12,
}

// magic numbers
var rookMagicNumbers = []uint64{
	0x8a80104000800020,
	0x140002000100040,
	0x2801880a0017001,
	0x100081001000420,
	0x200020010080420,
	0x3001c0002010008,
	0x8480008002000100,
	0x2080088004402900,
	0x800098204000,
	0x2024401000200040,
	0x100802000801000,
	0x120800800801000,
	0x208808088000400,
	0x2802200800400,
	0x2200800100020080,
	0x801000060821100,
	0x80044006422000,
	0x100808020004000,
	0x12108a0010204200,
	0x140848010000802,
	0x481828014002800,
	0x8094004002004100,
	0x4010040010010802,
	0x20008806104,
	0x100400080208000,
	0x2040002120081000,
	0x21200680100081,
	0x20100080080080,
	0x2000a00200410,
	0x20080800400,
	0x80088400100102,
	0x80004600042881,
	0x4040008040800020,
	0x440003000200801,
	0x4200011004500,
	0x188020010100100,
	0x14800401802800,
	0x2080040080800200,
	0x124080204001001,
	0x200046502000484,
	0x480400080088020,
	0x1000422010034000,
	0x30200100110040,
	0x100021010009,
	0x2002080100110004,
	0x202008004008002,
	0x20020004010100,
	0x2048440040820001,
	0x101002200408200,
	0x40802000401080,
	0x4008142004410100,
	0x2060820c0120200,
	0x1001004080100,
	0x20c020080040080,
	0x2935610830022400,
	0x44440041009200,
	0x280001040802101,
	0x2100190040002085,
	0x80c0084100102001,
	0x4024081001000421,
	0x20030a0244872,
	0x12001008414402,
	0x2006104900a0804,
	0x1004081002402,
}

var bishopMagicNumbers = []uint64{
	0x40040844404084,
	0x2004208a004208,
	0x10190041080202,
	0x108060845042010,
	0x581104180800210,
	0x2112080446200010,
	0x1080820820060210,
	0x3c0808410220200,
	0x4050404440404,
	0x21001420088,
	0x24d0080801082102,
	0x1020a0a020400,
	0x40308200402,
	0x4011002100800,
	0x401484104104005,
	0x801010402020200,
	0x400210c3880100,
	0x404022024108200,
	0x810018200204102,
	0x4002801a02003,
	0x85040820080400,
	0x810102c808880400,
	0xe900410884800,
	0x8002020480840102,
	0x220200865090201,
	0x2010100a02021202,
	0x152048408022401,
	0x20080002081110,
	0x4001001021004000,
	0x800040400a011002,
	0xe4004081011002,
	0x1c004001012080,
	0x8004200962a00220,
	0x8422100208500202,
	0x2000402200300c08,
	0x8646020080080080,
	0x80020a0200100808,
	0x2010004880111000,
	0x623000a080011400,
	0x42008c0340209202,
	0x209188240001000,
	0x400408a884001800,
	0x110400a6080400,
	0x1840060a44020800,
	0x90080104000041,
	0x201011000808101,
	0x1a2208080504f080,
	0x8012020600211212,
	0x500861011240000,
	0x180806108200800,
	0x4000020e01040044,
	0x300000261044000a,
	0x802241102020002,
	0x20906061210001,
	0x5a84841004010310,
	0x4010801011c04,
	0xa010109502200,
	0x4a02012000,
	0x500201010098b028,
	0x8040002811040900,
	0x28000010020204,
	0x6000020202d0240,
	0x8918844842082200,
	0x4010011029020020,
}

// pawn attacks table [side][square]
var pawnAttacks = [2][64]uint64{}

// knight attacks table [square]
var knightAttacks = [64]uint64{}

// king attacks table [square]
var kingAttacks = [64]uint64{}

// bishop attack masks
var bishopMasks = [64]uint64{}

// rook attacks mask
var rookMasks = [64]uint64{}

// bishop attacks table [square][occupancies]
var bishopAttacks = [64][512]uint64{}

// rook attacks table [square][occupancies]
var rookAttacks = [64][4096]uint64{}

func maskPawnAttacks(side int, square int) uint64 {
	// result attacks bitboard
	var attacks uint64 = 0
	// piece bitboard
	var bitboard uint64 = 0
	// set piece on board
	bitboard = setBit(bitboard, square)

	if side == WHITE {
		if (bitboard>>7)&NOT_A_FILE != 0 {
			attacks |= bitboard >> 7
		}
		if (bitboard>>9)&NOT_H_FILE != 0 {
			attacks |= bitboard >> 9
		}
	} else {
		// a8
		if (bitboard<<7)&NOT_H_FILE != 0 {
			attacks |= bitboard << 7
		}
		if (bitboard<<9)&NOT_A_FILE != 0 {
			attacks |= bitboard << 9
		}
	}

	return attacks
}

func maskKnightAttacks(square int) uint64 {
	var attacks uint64 = 0
	var bitboard uint64 = 0
	bitboard = setBit(bitboard, square)

	// up1 right2
	if (bitboard>>6)&NOT_AB_FILE != 0 {
		attacks |= bitboard >> 6
	}
	// up2 right1
	if (bitboard>>15)&NOT_A_FILE != 0 {
		attacks |= bitboard >> 15
	}
	// up2 left1
	if (bitboard>>17)&NOT_H_FILE != 0 {
		attacks |= bitboard >> 17
	}
	// up1 left2
	if (bitboard>>10)&NOT_HG_FILE != 0 {
		attacks |= bitboard >> 10
	}

	// down1 left2
	if (bitboard<<6)&NOT_HG_FILE != 0 {
		attacks |= bitboard << 6
	}
	// down2 left1
	if (bitboard<<15)&NOT_H_FILE != 0 {
		attacks |= bitboard << 15
	}
	// down2 right1
	if (bitboard<<17)&NOT_A_FILE != 0 {
		attacks |= bitboard << 17
	}
	//down1 right 2
	if (bitboard<<10)&NOT_AB_FILE != 0 {
		attacks |= bitboard << 10
	}

	return attacks
}

func maskKingAttacks(square int) uint64 {
	var attacks uint64 = 0
	var bitboard uint64 = 0
	bitboard = setBit(bitboard, square)

	if (bitboard>>1)&NOT_H_FILE != 0 {
		attacks |= bitboard >> 1
	}
	if (bitboard>>7)&NOT_A_FILE != 0 {
		attacks |= bitboard >> 7
	}
	if (bitboard>>9)&NOT_H_FILE != 0 {
		attacks |= bitboard >> 9
	}
	if bitboard>>8 != 0 {
		attacks |= bitboard >> 8
	}

	if (bitboard<<1)&NOT_A_FILE != 0 {
		attacks |= bitboard << 1
	}
	if (bitboard<<7)&NOT_H_FILE != 0 {
		attacks |= bitboard << 7
	}
	if (bitboard<<9)&NOT_A_FILE != 0 {
		attacks |= bitboard << 9
	}
	if bitboard<<8 != 0 {
		attacks |= bitboard << 8
	}

	return attacks
}

func maskBishopAttacks(square int) uint64 {
	var attacks uint64 = 0

	// init ranks & files
	var r, f int

	// init target rank & files
	tr := square / 8
	tf := square % 8

	// mask relevant bishop occupancy bits
	for r, f = tr+1, tf+1; r <= 6 && f <= 6; r, f = r+1, f+1 {
		attacks |= 1 << uint(r*8+f)
	}
	for r, f = tr+1, tf-1; r <= 6 && f >= 1; r, f = r+1, f-1 {
		attacks |= 1 << uint(r*8+f)
	}
	for r, f = tr-1, tf-1; r >= 1 && f >= 1; r, f = r-1, f-1 {
		attacks |= 1 << uint(r*8+f)
	}
	for r, f = tr-1, tf+1; r >= 1 && f <= 6; r, f = r-1, f+1 {
		attacks |= 1 << uint(r*8+f)
	}

	return attacks
}

func maskRookAttacks(square int) uint64 {
	var attacks uint64 = 0

	// rank & file
	var r, f int

	// target rank & file
	tr := square / 8
	tf := square % 8

	// mask relevant rook occupancy bits
	for r = tr + 1; r <= 6; r++ {
		attacks |= 1 << uint(r*8+tf)
	}
	for r = tr - 1; r >= 1; r-- {
		attacks |= 1 << uint(r*8+tf)
	}
	for f = tf + 1; f <= 6; f++ {
		attacks |= 1 << uint(tr*8+f)
	}
	for f = tf - 1; f >= 1; f-- {
		attacks |= 1 << uint(tr*8+f)
	}

	return attacks
}

// generate bishop attacks on the fly
func bishopAttacksOnTheFly(square int, block uint64) uint64 {
	var attacks uint64 = 0

	var r, f int

	tr := square / 8
	tf := square % 8

	// mask attacks, if we hit a blocker, don't go any further
	for r, f = tr+1, tf+1; r <= 7 && f <= 7; r, f = r+1, f+1 {
		attacks |= 1 << uint(r*8+f)
		if (1<<uint(r*8+f))&block != 0 {
			break
		}
	}
	for r, f = tr+1, tf-1; r <= 7 && f >= 0; r, f = r+1, f-1 {
		attacks |= 1 << uint(r*8+f)
		if (1<<uint(r*8+f))&block != 0 {
			break
		}
	}
	for r, f = tr-1, tf-1; r >= 0 && f >= 0; r, f = r-1, f-1 {
		attacks |= 1 << uint(r*8+f)
		if (1<<uint(r*8+f))&block != 0 {
			break
		}
	}
	for r, f = tr-1, tf+1; r >= 0 && f <= 7; r, f = r-1, f+1 {
		attacks |= 1 << uint(r*8+f)
		if (1<<uint(r*8+f))&block != 0 {
			break
		}
	}

	return attacks
}

func rookAttacksOnTheFly(square int, block uint64) uint64 {
	var attacks uint64 = 0

	var r, f int

	tr := square / 8
	tf := square % 8

	// mask attacks, if we hit a blocker, don't go any further
	for r = tr + 1; r <= 7; r++ {
		attacks |= 1 << uint(r*8+tf)
		if (1<<uint(r*8+tf))&block != 0 {
			break
		}
	}
	for r = tr - 1; r >= 0; r-- {
		attacks |= 1 << uint(r*8+tf)
		if (1<<uint(r*8+tf))&block != 0 {
			break
		}
	}
	for f = tf + 1; f <= 7; f++ {
		attacks |= 1 << uint(tr*8+f)
		if (1<<uint(tr*8+f))&block != 0 {
			break
		}
	}
	for f = tf - 1; f >= 0; f-- {
		attacks |= 1 << uint(tr*8+f)
		if (1<<uint(tr*8+f))&block != 0 {
			break
		}
	}

	return attacks
}

func initLeapersAttacks() {
	for square := range 64 {
		// init pawn attacks
		pawnAttacks[WHITE][square] = maskPawnAttacks(WHITE, square)
		pawnAttacks[BLACK][square] = maskPawnAttacks(BLACK, square)

		// init knight attacks
		knightAttacks[square] = maskKnightAttacks(square)

		// init king attacks
		kingAttacks[square] = maskKingAttacks(square)
	}
}

// index = 1, bitsInMask = 10, attackMask = rook d4
// square idx = 11

// Index is the configuration, if bitsInMask = 10, there are 2**10 - 1 = 1023 combinations
// of variations in bits set for a 10 bit number. If we iterate 1...1023,
// index will represent every combination by which bits are on/off. count helps us
// "loop" over and check each bit of Index. If that bit is set, we set our occupancy bit on at
// the associated square. We get the LS1B index so we know the offset for index->square.
// We pop the bit because otherwise we would be trying to place each square in the same
// place regardless of it's association with a particular index bit.
func setOccupancy(
	occupancyVariation int,
	bitsInMask int,
	attackMask uint64,
) uint64 {
	// occupancy map
	var occupancy uint64 = 0

	for idx := range bitsInMask {
		// get LS1B index of attack mask
		square := getLeastSignificantFirstBitIndex(attackMask)

		// pop LS1B in attack mask
		attackMask = popBit(attackMask, square)

		// check if bit in variation at idx offset is turned on
		if occupancyVariation&(1<<idx) != 0 {
			// populate occupancy map
			occupancy |= 1 << uint(square)
		}
	}

	return occupancy
}

/*********************************************************\
===========================================================

                        Magics

===========================================================
\*********************************************************/

/*
A magic number is a a number which will product unique indices for all of our attacks
that take into account the blocking occupancies.
*/
func findMagicNumber(square int, bishop int) uint64 {
	occupancies := make([]uint64, 4096)
	attacks := make([]uint64, 4096)

	var attackMask uint64
	// number of squares that can potentially block the piece's attacks from square
	var numBitsInMask int
	if bishop != 0 {
		attackMask = maskBishopAttacks(square)
		numBitsInMask = bishopRelevantBits[square]
	} else {
		attackMask = maskRookAttacks(square)
		numBitsInMask = rookRelevantBits[square]
	}

	// 2^relevantBits, we will loop over all possible occupancy configurations
	occupancyConfigurations := 1 << numBitsInMask

	// init our arrays with all occupancy variations and attacks for rook or bishop
	for i := range occupancyConfigurations {

		occupancies[i] = setOccupancy(
			i,
			numBitsInMask,
			attackMask,
		)

		if bishop != 0 {
			attacks[i] = bishopAttacksOnTheFly(square, occupancies[i])
		} else {
			attacks[i] = rookAttacksOnTheFly(square, occupancies[i])
		}
	}

	// trying 1 million magic numbers
	for range 10000000 {

		// get pseudo random uint64 number
		magicNumber := generateMagicNumber()
		usedAttacks := make([]uint64, 4096)

		// skip bad magic numbers
		if countBits((attackMask*magicNumber)&0xff00000000000000) < 6 {
			continue
		}

		failed := false
		for i := range occupancyConfigurations {
			magicIndex := int((occupancies[i] * magicNumber) >> (64 - numBitsInMask))

			if usedAttacks[magicIndex] == 0 {
				usedAttacks[magicIndex] = attacks[i]
			} else if usedAttacks[magicIndex] != attacks[i] {
				failed = true
				break
			}
		}

		if !failed {
			return magicNumber
		}
	}
	return 0
}

// init magic numbers
func initMagicNumbers() {
	for i := range 64 {
		// init rook magic numbers
		fmt.Printf("0x%x,\n", findMagicNumber(i, ROOK))
	}

	fmt.Println()

	for i := range 64 {
		// init bishop magic numbers
		fmt.Printf("0x%x,\n", findMagicNumber(i, BISHOP))
	}
}

/*
For each square, get the attack mask and number of bits in that
mask to calculate how many occupancy variations there are.
Go over each occupancy variation, get the occupancy for that square,
get the magic index for that square, and set our piece attacks
table for that square and index to the on the fly attack for the
occupancy
*/
func initSlidersAttacks(bishop int) {
	// loop over 64 board squares
	for square := range 64 {
		// init bishop & rook masks
		bishopMasks[square] = maskBishopAttacks(square)
		rookMasks[square] = maskRookAttacks(square)

		// init current mask
		var attackMask uint64
		if bishop == 1 {
			attackMask = bishopMasks[square]
		} else {
			attackMask = rookMasks[square]
		}

		// init num bits in mask
		numBitsInMask := countBits(attackMask)

		// init occupancy indices
		occupancyIndices := (1 << numBitsInMask)

		for i := range occupancyIndices {
			// bishop
			if bishop == 1 {
				// init current occupancy
				occupancy := setOccupancy(i, numBitsInMask, attackMask)

				// init magic index
				magicIndex := (occupancy * bishopMagicNumbers[square]) >> (64 - bishopRelevantBits[square])

				// init bishop attacks
				bishopAttacks[square][magicIndex] = bishopAttacksOnTheFly(square, occupancy)
			} else {
				// init current occupancy
				occupancy := setOccupancy(i, numBitsInMask, attackMask)

				// init magic index
				magicIndex := (occupancy * rookMagicNumbers[square]) >> (64 - rookRelevantBits[square])

				// init bishop attacks
				rookAttacks[square][magicIndex] = rookAttacksOnTheFly(square, occupancy)
			}
		}
	}
}

/*
These just get the magic hash index so we can look up attacks quickly
*/
func getBishopAttacks(square int, occupancy uint64) uint64 {
	// get bishop attacks assuming current board occupancy
	occupancy &= bishopMasks[square]
	occupancy *= bishopMagicNumbers[square]
	occupancy >>= 64 - bishopRelevantBits[square]

	return bishopAttacks[square][occupancy]
}

func getRookAttacks(square int, occupancy uint64) uint64 {
	// get bishop attacks assuming current board occupancy
	occupancy &= rookMasks[square]
	occupancy *= rookMagicNumbers[square]
	occupancy >>= 64 - rookRelevantBits[square]

	return rookAttacks[square][occupancy]
}

/*
Just get attacks for both rook/bishop and bitwise OR them
*/
func getQueenAttacks(square int, occupancy uint64) uint64 {
	return getBishopAttacks(square, occupancy) | getRookAttacks(square, occupancy)
}

/*********************************************************\
===========================================================

                        Move Generator

===========================================================
\*********************************************************/

// is given square attacked by the given side?
func isSquareAttacked(square, side int) int {
	if side == WHITE && pawnAttacks[BLACK][square]&bitboards[P] > 0 {
		return 1
	} else if side == BLACK && pawnAttacks[WHITE][square]&bitboards[p] > 0 {
		return 1
	}

	if side == WHITE && (knightAttacks[square]&bitboards[N]) > 0 {
		return 1
	} else if side == BLACK && (knightAttacks[square]&bitboards[n]) > 0 {
		return 1
	}

	if side == WHITE && (getBishopAttacks(square, occupancies[BOTH])&bitboards[B]) > 0 {
		return 1
	} else if side == BLACK && (getBishopAttacks(square, occupancies[BOTH])&bitboards[b]) > 0 {
		return 1
	}

	if side == WHITE && (getRookAttacks(square, occupancies[BOTH])&bitboards[R]) > 0 {
		return 1
	} else if side == BLACK && (getRookAttacks(square, occupancies[BOTH])&bitboards[r]) > 0 {
		return 1
	}

	if side == WHITE && (getQueenAttacks(square, occupancies[BOTH])&bitboards[Q]) > 0 {
		return 1
	} else if side == BLACK && (getQueenAttacks(square, occupancies[BOTH])&bitboards[q]) > 0 {
		return 1
	}

	if side == WHITE && (kingAttacks[square]&bitboards[K]) > 0 {
		return 1
	} else if side == BLACK && (kingAttacks[square]&bitboards[k]) > 0 {
		return 1
	}

	return 0
}

func printAttackedSquares(side int) {
	fmt.Println()

	for rank := range 8 {
		for file := range 8 {
			square := rank*8 + file

			if file == 0 {
				fmt.Printf("  %d  ", 8-rank)
			}

			fmt.Printf(" %d ", isSquareAttacked(square, side))
		}
		fmt.Println()
	}

	fmt.Println("\n      a  b  c  d  e  f  g  h ")
}

const (
	allMoves     = 0
	onlyCaptures = 1
)

var castlingRights = [64]int{
	7, 15, 15, 15, 3, 15, 15, 11,
	15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15,
	13, 15, 15, 15, 12, 15, 15, 14,
}

func makeMove(move, moveFlag int) int {
	// quiet moves
	if moveFlag == allMoves {
		var copy BoardStateCopy
		copyBoardState(&copy)

		sourceSquare := getMoveSourceSquare(move)
		targetSquare := getMoveTargetSquare(move)
		piece := getMovePiece(move)
		promotedPiece := getMovePromotedPiece(move)
		capture := getMoveCaptureFlag(move)
		double := getMoveDoublePawnPushFlag(move)
		enpass := getMoveEnpassantFlag(move)
		castling := getMoveCastlingFlag(move)

		// remove piece from source square and update occupancies
		bitboards[piece] = popBit(bitboards[piece], sourceSquare)
		occupancies[BOTH] = popBit(occupancies[BOTH], sourceSquare)
		if side == WHITE {
			occupancies[WHITE] = popBit(occupancies[WHITE], sourceSquare)
		} else {
			occupancies[BLACK] = popBit(occupancies[BLACK], sourceSquare)
		}

		// handle capture moves
		if capture > 0 {
			var startPiece int
			var endPiece int

			// loop over piece bitboards to remove captured piece from board
			if side == WHITE {
				startPiece = p
				endPiece = k
			} else {
				startPiece = P
				endPiece = K
			}
			for bbPiece := startPiece; bbPiece <= endPiece; bbPiece++ {
				// if there is a piece on target square,
				if getBit(bitboards[bbPiece], targetSquare) > 0 {
					bitboards[bbPiece] = popBit(bitboards[bbPiece], targetSquare)
					occupancies[BOTH] = popBit(occupancies[BOTH], targetSquare)
					if bbPiece >= P && bbPiece <= K {
						occupancies[WHITE] = popBit(occupancies[WHITE], targetSquare)
					} else {
						occupancies[BLACK] = popBit(occupancies[BLACK], targetSquare)
					}
					break
				}
			}
		}

		// pawn promotions
		if promotedPiece > 0 {
			bitboards[promotedPiece] = setBit(bitboards[promotedPiece], targetSquare)
			occupancies[BOTH] = setBit(occupancies[BOTH], targetSquare)
			if side == WHITE {
				occupancies[WHITE] = setBit(occupancies[WHITE], targetSquare)
			} else {
				occupancies[BLACK] = setBit(occupancies[BLACK], targetSquare)
			}
		} else {
			bitboards[piece] = setBit(bitboards[piece], targetSquare)
			occupancies[BOTH] = setBit(occupancies[BOTH], targetSquare)
			if side == WHITE {
				occupancies[WHITE] = setBit(occupancies[WHITE], targetSquare)
			} else {
				occupancies[BLACK] = setBit(occupancies[BLACK], targetSquare)
			}
		}

		// en passant capture
		if enpass > 0 {
			var captureSquare int
			var capturePiece int
			if side == WHITE {
				captureSquare = targetSquare + 8
				capturePiece = p
			} else {
				captureSquare = targetSquare - 8
				capturePiece = P
			}
			bitboards[capturePiece] = popBit(bitboards[capturePiece], captureSquare)
			occupancies[BOTH] = popBit(occupancies[BOTH], captureSquare)
			if side == WHITE {
				occupancies[BLACK] = popBit(occupancies[BLACK], captureSquare)
			} else {
				occupancies[WHITE] = popBit(occupancies[WHITE], captureSquare)
			}
		}
		enpassant = NO_SQ

		// double pawn push
		if double > 0 {
			// setup enpassant square
			if side == WHITE {
				enpassant = targetSquare + 8
			} else {
				enpassant = targetSquare - 8
			}
		}

		// castling - move rook and update occupancies
		if castling > 0 {
			switch targetSquare {
			case g1:
				bitboards[R] = popBit(bitboards[R], h1)
				bitboards[R] = setBit(bitboards[R], f1)
				occupancies[BOTH] = popBit(occupancies[BOTH], h1)
				occupancies[BOTH] = setBit(occupancies[BOTH], f1)
				occupancies[WHITE] = popBit(occupancies[WHITE], h1)
				occupancies[WHITE] = setBit(occupancies[WHITE], f1)
			case c1:
				bitboards[R] = popBit(bitboards[R], a1)
				bitboards[R] = setBit(bitboards[R], d1)
				occupancies[BOTH] = popBit(occupancies[BOTH], a1)
				occupancies[BOTH] = setBit(occupancies[BOTH], d1)
				occupancies[WHITE] = popBit(occupancies[WHITE], a1)
				occupancies[WHITE] = setBit(occupancies[WHITE], d1)
			case g8:
				bitboards[r] = popBit(bitboards[r], h8)
				bitboards[r] = setBit(bitboards[r], f8)
				occupancies[BOTH] = popBit(occupancies[BOTH], h8)
				occupancies[BOTH] = setBit(occupancies[BOTH], f8)
				occupancies[BLACK] = popBit(occupancies[BLACK], h8)
				occupancies[BLACK] = setBit(occupancies[BLACK], f8)
			case c8:
				bitboards[r] = popBit(bitboards[r], a8)
				bitboards[r] = setBit(bitboards[r], d8)
				occupancies[BOTH] = popBit(occupancies[BOTH], a8)
				occupancies[BOTH] = setBit(occupancies[BOTH], d8)
				occupancies[BLACK] = popBit(occupancies[BLACK], a8)
				occupancies[BLACK] = setBit(occupancies[BLACK], d8)
			}
		}

		// update castling rights
		castle &= castlingRights[sourceSquare]
		castle &= castlingRights[targetSquare]

		// check if king is in check (WIP)!
		side ^= 1

		if side == WHITE && isSquareAttacked(getLeastSignificantFirstBitIndex(bitboards[k]), side) > 0 {
			restorePreviousBoardState(copy)
			return -1
		} else if side == BLACK && isSquareAttacked(getLeastSignificantFirstBitIndex(bitboards[K]), side) > 0 {
			restorePreviousBoardState(copy)
			return -1
		} else {
			return 1
		}
	} else {
		// capture moves
		// make sure move is capture
		if getMoveCaptureFlag(move) > 0 {
			makeMove(move, allMoves)
		} else {
			return -1
		}
	}
	return 1
}

func generateMoves(moveList *Moves) {
	moveList.count = 0

	// define source & target squares
	var sourceSquare int
	var targetSquare int

	// define current pieces bitboard copy & its attacks
	var bitboard uint64
	var attacks uint64

	for piece := P; piece <= k; piece++ {
		bitboard = bitboards[piece]

		// gen pawn moves
		if side == WHITE {
			if piece == P {
				for bitboard > 0 {
					sourceSquare = getLeastSignificantFirstBitIndex(bitboard)
					targetSquare = sourceSquare - 8
					if targetSquare >= a8 && getBit(occupancies[BOTH], targetSquare) == 0 {
						if sourceSquare >= a7 && sourceSquare <= h7 {
							// promotions
							addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, Q, 0, 0, 0, 0))
							addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, R, 0, 0, 0, 0))
							addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, B, 0, 0, 0, 0))
							addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, N, 0, 0, 0, 0))
						} else {
							// single/double push
							addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, 0, 0, 0, 0, 0))

							if (sourceSquare >= a2 && sourceSquare <= h2) && getBit(occupancies[BOTH], targetSquare-8) == 0 {
								addMove(moveList, encodeMove(sourceSquare, targetSquare-8, piece, 0, 0, 1, 0, 0))
							}
						}
					}
					// white captures
					attacks = pawnAttacks[side][sourceSquare] & occupancies[BLACK]
					for attacks > 0 {
						targetSquare = getLeastSignificantFirstBitIndex(attacks)

						if sourceSquare >= a7 && sourceSquare <= h7 {
							// promotion capture
							addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, Q, 1, 0, 0, 0))
							addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, R, 1, 0, 0, 0))
							addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, B, 1, 0, 0, 0))
							addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, N, 1, 0, 0, 0))
						} else {
							// capture
							addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, 0, 1, 0, 0, 0))
						}
						attacks = popBit(attacks, targetSquare)
					}

					// white enpassant
					if enpassant != NO_SQ {
						enpassantAttacks := pawnAttacks[side][sourceSquare] & (1 << enpassant) // ?

						if enpassantAttacks > 0 {
							targetEnpassant := getLeastSignificantFirstBitIndex(enpassantAttacks)
							addMove(moveList, encodeMove(sourceSquare, targetEnpassant, piece, 0, 1, 0, 1, 0))
						}
					}
					bitboard = popBit(bitboard, sourceSquare)
				}
			}
			// white castling
			if piece == K {
				if castle&WK != 0 {
					if getBit(occupancies[BOTH], f1) == 0 && getBit(occupancies[BOTH], g1) == 0 {
						if isSquareAttacked(e1, BLACK) == 0 && isSquareAttacked(f1, BLACK) == 0 {
							addMove(moveList, encodeMove(e1, g1, piece, 0, 0, 0, 0, 1))
						}
					}
				}

				if castle&WQ != 0 {
					if getBit(occupancies[BOTH], d1) == 0 && getBit(occupancies[BOTH], c1) == 0 && getBit(occupancies[BOTH], b1) == 0 {
						if isSquareAttacked(e1, BLACK) == 0 && isSquareAttacked(d1, BLACK) == 0 {
							addMove(moveList, encodeMove(e1, c1, piece, 0, 0, 0, 0, 1))
						}
					}
				}
			}
			// black pawns & king castling
		} else {
			if piece == p {
				for bitboard > 0 {
					sourceSquare = getLeastSignificantFirstBitIndex(bitboard)
					targetSquare = sourceSquare + 8

					if targetSquare <= h1 && getBit(occupancies[BOTH], targetSquare) == 0 {
						if sourceSquare >= a2 && sourceSquare <= h2 {
							// black pawn promotions
							addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, q, 0, 0, 0, 0))
							addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, r, 0, 0, 0, 0))
							addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, b, 0, 0, 0, 0))
							addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, n, 0, 0, 0, 0))
						} else {
							// black push
							addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, 0, 0, 0, 0, 0))

							// black double pawn push
							if (sourceSquare >= a7 && sourceSquare <= h7) && getBit(occupancies[BOTH], targetSquare+8) == 0 {
								addMove(moveList, encodeMove(sourceSquare, targetSquare+8, piece, 0, 0, 1, 0, 0))
							}
						}
					}
					// black captures
					attacks = pawnAttacks[side][sourceSquare] & occupancies[WHITE]

					for attacks > 0 {
						targetSquare = getLeastSignificantFirstBitIndex(attacks)

						if sourceSquare >= a2 && sourceSquare <= h2 {
							addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, q, 1, 0, 0, 0))
							addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, r, 1, 0, 0, 0))
							addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, b, 1, 0, 0, 0))
							addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, n, 1, 0, 0, 0))
						} else {
							addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, 0, 1, 0, 0, 0))
						}
						attacks = popBit(attacks, targetSquare)
					}

					if enpassant != NO_SQ {
						enpassantAttacks := pawnAttacks[side][sourceSquare] & (1 << enpassant)

						if enpassantAttacks > 0 {
							targetEnpassant := getLeastSignificantFirstBitIndex(enpassantAttacks)
							addMove(moveList, encodeMove(sourceSquare, targetEnpassant, piece, 0, 1, 0, 1, 0))
						}
					}
					bitboard = popBit(bitboard, sourceSquare)
				}
			}

			// black castling
			if piece == k {
				if castle&BK != 0 {
					if getBit(occupancies[BOTH], f8) == 0 && getBit(occupancies[BOTH], g8) == 0 {
						if isSquareAttacked(e8, WHITE) == 0 && isSquareAttacked(f8, WHITE) == 0 {
							addMove(moveList, encodeMove(e8, g8, piece, 0, 0, 0, 0, 1))
						}
					}
				}

				if castle&BQ != 0 {
					if getBit(occupancies[BOTH], d8) == 0 && getBit(occupancies[BOTH], c8) == 0 && getBit(occupancies[BOTH], b8) == 0 {
						if isSquareAttacked(e8, WHITE) == 0 && isSquareAttacked(d8, WHITE) == 0 {
							addMove(moveList, encodeMove(e8, c8, piece, 0, 0, 0, 0, 1))
						}
					}
				}
			}
		}

		// gen knight moves
		if side == WHITE && piece == N {
			for bitboard > 0 {
				sourceSquare = getLeastSignificantFirstBitIndex(bitboard)

				attacks = knightAttacks[sourceSquare] & ^occupancies[WHITE]

				for attacks > 0 {
					targetSquare = getLeastSignificantFirstBitIndex(attacks)

					// quiet move, else capture
					if getBit(occupancies[BLACK], targetSquare) == 0 {
						addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, 0, 0, 0, 0, 0))
					} else {
						addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, 0, 1, 0, 0, 0))
					}

					attacks = popBit(attacks, targetSquare)
				}

				bitboard = popBit(bitboard, sourceSquare)
			}
		} else if side == BLACK && piece == n {
			for bitboard > 0 {
				sourceSquare = getLeastSignificantFirstBitIndex(bitboard)

				attacks = knightAttacks[sourceSquare] & ^occupancies[BLACK]

				for attacks > 0 {
					targetSquare = getLeastSignificantFirstBitIndex(attacks)

					// quiet move, else capture
					if getBit(occupancies[WHITE], targetSquare) == 0 {
						addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, 0, 0, 0, 0, 0))
					} else {
						addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, 0, 1, 0, 0, 0))
					}

					attacks = popBit(attacks, targetSquare)
				}

				bitboard = popBit(bitboard, sourceSquare)
			}
		}

		// gen bishop moves
		if side == WHITE && piece == B {
			for bitboard > 0 {
				sourceSquare = getLeastSignificantFirstBitIndex(bitboard)

				attacks = getBishopAttacks(sourceSquare, occupancies[BOTH]) & ^occupancies[WHITE]

				for attacks > 0 {
					targetSquare = getLeastSignificantFirstBitIndex(attacks)

					// quiet move, else capture
					if getBit(occupancies[BLACK], targetSquare) == 0 {
						addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, 0, 0, 0, 0, 0))
					} else {
						addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, 0, 1, 0, 0, 0))
					}

					attacks = popBit(attacks, targetSquare)
				}

				bitboard = popBit(bitboard, sourceSquare)
			}
		} else if side == BLACK && piece == b {
			for bitboard > 0 {
				sourceSquare = getLeastSignificantFirstBitIndex(bitboard)

				attacks = getBishopAttacks(sourceSquare, occupancies[BOTH]) & ^occupancies[BLACK]

				for attacks > 0 {
					targetSquare = getLeastSignificantFirstBitIndex(attacks)

					// quiet move, else capture
					if getBit(occupancies[WHITE], targetSquare) == 0 {
						addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, 0, 0, 0, 0, 0))
					} else {
						addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, 0, 1, 0, 0, 0))
					}

					attacks = popBit(attacks, targetSquare)
				}

				bitboard = popBit(bitboard, sourceSquare)
			}
		}

		// gen rook moves
		if side == WHITE && piece == R {
			for bitboard > 0 {
				sourceSquare = getLeastSignificantFirstBitIndex(bitboard)

				attacks = getRookAttacks(sourceSquare, occupancies[BOTH]) & ^occupancies[WHITE]

				for attacks > 0 {
					targetSquare = getLeastSignificantFirstBitIndex(attacks)

					// quiet move, else capture
					if getBit(occupancies[BLACK], targetSquare) == 0 {
						addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, 0, 0, 0, 0, 0))
					} else {
						addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, 0, 1, 0, 0, 0))
					}

					attacks = popBit(attacks, targetSquare)
				}

				bitboard = popBit(bitboard, sourceSquare)
			}
		} else if side == BLACK && piece == r {
			for bitboard > 0 {
				sourceSquare = getLeastSignificantFirstBitIndex(bitboard)

				attacks = getRookAttacks(sourceSquare, occupancies[BOTH]) & ^occupancies[BLACK]

				for attacks > 0 {
					targetSquare = getLeastSignificantFirstBitIndex(attacks)

					// quiet move, else capture
					if getBit(occupancies[WHITE], targetSquare) == 0 {
						addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, 0, 0, 0, 0, 0))
					} else {
						addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, 0, 1, 0, 0, 0))
					}

					attacks = popBit(attacks, targetSquare)
				}

				bitboard = popBit(bitboard, sourceSquare)
			}
		}

		// gen queen moves
		if side == WHITE && piece == Q {
			for bitboard > 0 {
				sourceSquare = getLeastSignificantFirstBitIndex(bitboard)

				attacks = getQueenAttacks(sourceSquare, occupancies[BOTH]) & ^occupancies[WHITE]

				for attacks > 0 {
					targetSquare = getLeastSignificantFirstBitIndex(attacks)

					// quiet move, else capture
					if getBit(occupancies[BLACK], targetSquare) == 0 {
						addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, 0, 0, 0, 0, 0))
					} else {
						addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, 0, 1, 0, 0, 0))
					}

					attacks = popBit(attacks, targetSquare)
				}

				bitboard = popBit(bitboard, sourceSquare)
			}
		} else if side == BLACK && piece == q {
			for bitboard > 0 {
				sourceSquare = getLeastSignificantFirstBitIndex(bitboard)

				attacks = getQueenAttacks(sourceSquare, occupancies[BOTH]) & ^occupancies[BLACK]

				for attacks > 0 {
					targetSquare = getLeastSignificantFirstBitIndex(attacks)

					// quiet move, else capture
					if getBit(occupancies[WHITE], targetSquare) == 0 {
						addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, 0, 0, 0, 0, 0))
					} else {
						addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, 0, 1, 0, 0, 0))
					}

					attacks = popBit(attacks, targetSquare)
				}

				bitboard = popBit(bitboard, sourceSquare)
			}
		}

		// gen non-castling king moves
		if side == WHITE && piece == K {
			for bitboard > 0 {
				sourceSquare = getLeastSignificantFirstBitIndex(bitboard)

				attacks = kingAttacks[sourceSquare] & ^occupancies[WHITE]

				for attacks > 0 {
					targetSquare = getLeastSignificantFirstBitIndex(attacks)

					// quiet move, else capture
					if getBit(occupancies[BLACK], targetSquare) == 0 {
						addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, 0, 0, 0, 0, 0))
					} else {
						addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, 0, 1, 0, 0, 0))
					}

					attacks = popBit(attacks, targetSquare)
				}

				bitboard = popBit(bitboard, sourceSquare)
			}
		} else if side == BLACK && piece == k {
			for bitboard > 0 {
				sourceSquare = getLeastSignificantFirstBitIndex(bitboard)

				attacks = kingAttacks[sourceSquare] & ^occupancies[BLACK]

				for attacks > 0 {
					targetSquare = getLeastSignificantFirstBitIndex(attacks)

					// quiet move, else capture
					if getBit(occupancies[WHITE], targetSquare) == 0 {
						addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, 0, 0, 0, 0, 0))
					} else {
						addMove(moveList, encodeMove(sourceSquare, targetSquare, piece, 0, 1, 0, 0, 0))
					}

					attacks = popBit(attacks, targetSquare)
				}

				bitboard = popBit(bitboard, sourceSquare)
			}
		}
	}
}

/*
binary move bits																					hexadecimal constants

0000 0000 0000 0000 0011 1111		source square							0x3f
0000 0000 0000 1111 1100 0000		target square							0xfc0
0000 0000 1111 0000 0000 0000		piece											0xf000
0000 1111 0000 0000 0000 0000		promoted piece						0xf0000
0001 0000 0000 0000 0000 0000		capture flag							0x100000
0010 0000 0000 0000 0000 0000		double pawn push flag			0x200000
0100 0000 0000 0000 0000 0000		enpassant capture					0x400000
1000 0000 0000 0000 0000 0000		castling flag							0x800000
*/

func encodeMove(source, target, piece, promoted, capture, double, enpassant, castling int) int {
	return source |
		target<<6 |
		piece<<12 |
		promoted<<16 |
		capture<<20 |
		double<<21 |
		enpassant<<22 |
		castling<<23
}

func getMoveSourceSquare(move int) int {
	return move & 0x3f
}

func getMoveTargetSquare(move int) int {
	return (move & 0xfc0) >> 6
}

func getMovePiece(move int) int {
	return (move & 0xf000) >> 12
}

func getMovePromotedPiece(move int) int {
	return (move & 0xf0000) >> 16
}

// not shifted! check == 0 or > 0 for non-flagged/flagged respectively
func getMoveCaptureFlag(move int) int {
	return (move & 0x100000)
}

func getMoveDoublePawnPushFlag(move int) int {
	return (move & 0x200000)
}

func getMoveEnpassantFlag(move int) int {
	return (move & 0x400000)
}

func getMoveCastlingFlag(move int) int {
	return (move & 0x800000)
}

type Moves struct {
	moves [256]int
	count int
}

func addMove(moveList *Moves, move int) {
	moveList.moves[moveList.count] = move
	moveList.count++
}

// for UCI
func printMove(move int) {
	if getMovePromotedPiece(move) > 0 {
		fmt.Printf("%s%s%c\n",
			SquareToCoordinates[getMoveSourceSquare(move)],
			SquareToCoordinates[getMoveTargetSquare(move)],
			promotedPieces[getMovePromotedPiece(move)])
	} else {
		fmt.Printf("%s%s\n",
			SquareToCoordinates[getMoveSourceSquare(move)],
			SquareToCoordinates[getMoveTargetSquare(move)])
	}
}

// for debugging
func printMoveList(moveList *Moves) {
	if moveList.count == 0 {
		fmt.Printf("\n    No moves in the move list!\n")
		return
	}

	fmt.Printf("\n    move     piece    capture    double    enpassant    castling\n")

	for i := range moveList.count {
		move := moveList.moves[i]

		capture := 0
		doublePawn := 0
		enpassant := 0
		castling := 0
		promotedPiece := 0
		if getMoveCaptureFlag(move) > 0 {
			capture = 1
		}
		if getMoveDoublePawnPushFlag(move) > 0 {
			doublePawn = 1
		}
		if getMoveEnpassantFlag(move) > 0 {
			enpassant = 1
		}
		if getMoveCastlingFlag(move) > 0 {
			castling = 1
		}
		if getMovePromotedPiece(move) > 0 {
			promotedPiece = int(promotedPieces[getMovePromotedPiece(move)])
		}

		promotedPieceChar := ' '
		if promotedPiece > 0 {
			promotedPieceChar = rune(promotedPiece)
		}

		fmt.Printf("    %s%s%c      %c         %d         %d           %d           %d\n",
			SquareToCoordinates[getMoveSourceSquare(move)],
			SquareToCoordinates[getMoveTargetSquare(move)],
			promotedPieceChar,
			asciiPieces[getMovePiece(move)],
			capture,
			doublePawn,
			enpassant,
			castling,
		)
	}
	fmt.Println("\n\n    total number of moves:", moveList.count)
}

/*********************************************************\
===========================================================

                        Perft

===========================================================
\*********************************************************/

func getTimeMS() int64 {
	return time.Now().UnixMilli()
}

func copyBoardState(state *BoardStateCopy) {
	copy(state.bitboards[:], bitboards[:])
	copy(state.occupancies[:], occupancies[:])
	state.side = side
	state.enpassant = enpassant
	state.castle = castle
}

func restorePreviousBoardState(state BoardStateCopy) {
	copy(bitboards[:], state.bitboards[:])
	copy(occupancies[:], state.occupancies[:])
	side = state.side
	enpassant = state.enpassant
	castle = state.castle
}

type BoardStateCopy struct {
	bitboards   [12]uint64
	occupancies [3]uint64
	side        int
	enpassant   int
	castle      int
}

var nodesCount uint64

func perftDriver(depth int) {
	if depth == 0 {
		nodesCount++
		return
	}

	var moveList Moves
	generateMoves(&moveList)

	for moveCount := 0; moveCount < moveList.count; moveCount++ {
		var copy BoardStateCopy
		copyBoardState(&copy)

		if makeMove(moveList.moves[moveCount], allMoves) != 1 {
			continue
		}

		perftDriver(depth - 1)

		restorePreviousBoardState(copy)
	}
}

func perftTest(depth int) {
	fmt.Printf("\n     Performance Test\n\n")
	var moveList Moves
	generateMoves(&moveList)

	startTime := getTimeMS()
	for moveCount := 0; moveCount < moveList.count; moveCount++ {
		var copy BoardStateCopy
		copyBoardState(&copy)

		if makeMove(moveList.moves[moveCount], allMoves) != 1 {
			continue
		}

		var cumulativeNodes uint64 = nodesCount

		perftDriver(depth - 1)

		oldNodes := nodesCount - cumulativeNodes

		restorePreviousBoardState(copy)

		fmt.Printf("     move: %s%s%c   nodes: %d\n",
			SquareToCoordinates[getMoveSourceSquare(moveList.moves[moveCount])],
			SquareToCoordinates[getMoveTargetSquare(moveList.moves[moveCount])],
			promotedPieces[getMovePromotedPiece(moveList.moves[moveCount])],
			oldNodes)
	}
	endTime := getTimeMS()
	fmt.Printf("\n     Depth: %d\n", depth)
	fmt.Printf("     Nodes: %d\n", nodesCount)
	fmt.Printf("      Time: %dms\n\n", endTime-startTime)
}

/*********************************************************\
===========================================================

                        UCI

===========================================================
\*********************************************************/

// format for moves "e7e8q" or "e2e3"
func parseMove(moveString string) int {
	var moveList Moves
	generateMoves(&moveList)

	sourceSquare := (moveString[0] - 'a') + (8-(moveString[1]-'0'))*8
	targetSquare := (moveString[2] - 'a') + (8-(moveString[3]-'0'))*8

	for moveCount := 0; moveCount < moveList.count; moveCount++ {
		// get move from generated moves based on source/target squares
		move := moveList.moves[moveCount]

		if sourceSquare == byte(getMoveSourceSquare(move)) && targetSquare == byte(getMoveTargetSquare(move)) {
			// possibly useless? only function is to check if moveString is invalid
			promotedPiece := getMovePromotedPiece(move)

			if getMovePromotedPiece(move) > 0 {
				if len(moveString) < 5 {
					return 0
				}
				if (promotedPiece == Q || promotedPiece == q) && moveString[4] == 'q' {
					return move
				} else if (promotedPiece == R || promotedPiece == r) && moveString[4] == 'r' {
					return move
				} else if (promotedPiece == B || promotedPiece == b) && moveString[4] == 'b' {
					return move
				} else if (promotedPiece == N || promotedPiece == n) && moveString[4] == 'n' {
					return move
				}
				continue
			}

			return move
		}
	}
	return 0
}

/*********************************************************\
===========================================================

                        Init all

===========================================================
\*********************************************************/

func initAll() {
	initLeapersAttacks()
	initSlidersAttacks(BISHOP)
	initSlidersAttacks(ROOK)
}

/*********************************************************\
===========================================================

                        Main driver

===========================================================
\*********************************************************/

func main() {
	initAll()

	parseFEN("r3k2r/pPppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1 ")
	printBoard()

	move := parseMove("b7b8q")
	if move > 0 {
		makeMove(move, allMoves)
	} else {
		fmt.Println("illegal move!")
	}
	printBoard()
}
