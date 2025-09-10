# bb-go

A bitboard-based chess engine implementation in Go!

## TODO
- make sure checking if king is in check is working properly

## Building and Running

### Prerequisites
- Go 1.18 or later

### Build
```bash
go build -o bb-go main.go
```

### Run
```bash
./bb-go
```

Or run directly without building:
```bash
go run main.go
```

## Bitboards Used

### Core Board Representation
- **bitboards[12]**: Individual piece bitboards (P, N, B, R, Q, K, p, n, b, r, q, k)
- **occupancies[3]**: Combined occupancy bitboards (WHITE, BLACK, BOTH)

### Attack Tables
- **pawnAttacks[2][64]**: Precomputed pawn attack patterns for each side and square
- **knightAttacks[64]**: Precomputed knight attack patterns for each square
- **kingAttacks[64]**: Precomputed king attack patterns for each square
- **bishopAttacks[64][512]**: Magic bitboard lookup table for bishop attacks
- **rookAttacks[64][4096]**: Magic bitboard lookup table for rook attacks

### Attack Masks
- **bishopMasks[64]**: Relevant occupancy masks for bishop attacks on each square
- **rookMasks[64]**: Relevant occupancy masks for rook attacks on each square

### File Masks
- **NOT_A_FILE**: Mask excluding A file (prevents wrap-around attacks)
- **NOT_H_FILE**: Mask excluding H file (prevents wrap-around attacks)
- **NOT_HG_FILE**: Mask excluding H and G files
- **NOT_AB_FILE**: Mask excluding A and B files

## Function Reference

### Random Number Generation
- **getRandom32BitUnsignedNumber()**: XOR-shift PRNG for 32-bit numbers
- **getRandom64BitUnsignedNumber()**: Combines four 32-bit numbers into 64-bit
- **generateMagicNumber()**: Generates sparse random numbers for magic bitboard candidates

### Bit Manipulation
- **getBit(bitboard, square)**: Check if bit is set at given square
- **setBit(bitboard, square)**: Set bit at given square
- **popBit(bitboard, square)**: Clear bit at given square
- **countBits(bitboard)**: Count number of set bits using Brian Kernighan's algorithm
- **getLeastSignificantFirstBitIndex(bitboard)**: Find index of least significant set bit

### Board Display
- **printBitboard(bitboard)**: Display bitboard as 8x8 grid with coordinates
- **printBoard()**: Display current board position with pieces, side to move, castling rights, and en passant

### FEN Parsing
- **parseFEN(fen)**: Parse Forsyth-Edwards Notation string into board state
- **isUpperCase(character)**: Check if character is uppercase letter
- **isAlpha(character)**: Check if character is alphabetic
- **isNumeric(character)**: Check if character is numeric digit

### Attack Pattern Generation
- **maskPawnAttacks(side, square)**: Generate pawn attack pattern for given side and square
- **maskKnightAttacks(square)**: Generate knight attack pattern for given square
- **maskKingAttacks(square)**: Generate king attack pattern for given square
- **maskBishopAttacks(square)**: Generate bishop attack mask (excluding edges)
- **maskRookAttacks(square)**: Generate rook attack mask (excluding edges)

### On-the-Fly Attack Generation
- **bishopAttacksOnTheFly(square, block)**: Calculate bishop attacks with blockers
- **rookAttacksOnTheFly(square, block)**: Calculate rook attacks with blockers

### Attack Table Initialization
- **initLeapersAttacks()**: Initialize attack tables for pawns, knights, and kings
- **initSlidersAttacks(bishop)**: Initialize magic bitboard attack tables for sliding pieces
- **initAll()**: Initialize all attack tables and data structures

### Magic Bitboard System
- **setOccupancy(variation, bitsInMask, attackMask)**: Generate occupancy variation from index
- **findMagicNumber(square, bishop)**: Find magic number for given square and piece type
- **initMagicNumbers()**: Generate and print magic numbers (development utility)

### Attack Lookup
- **getBishopAttacks(square, occupancy)**: Get bishop attacks using magic bitboards
- **getRookAttacks(square, occupancy)**: Get rook attacks using magic bitboards
- **getQueenAttacks(square, occupancy)**: Get queen attacks (combines bishop and rook)

## Bitwise Operations Notes

### Basic Operations
- `|=` performs logical OR and assigns result to left operand
- `&` returns 1 only when both operands have 1 in same position

### Least Significant Bit Isolation
```
block & (~block + 1)
```
- `~block` flips all bits
- `~block + 1` creates two's complement
- `&` operation isolates the least significant set bit

### Bit Counting
Uses Brian Kernighan's algorithm: `bitboard &= bitboard - 1` repeatedly clears the least significant set bit.

## Development Notes
- Consider hardware BMI (Bit Manipulation Instructions) for faster operations
- Magic numbers provide O(1) sliding piece attack lookup
- Bitboards enable parallel processing of multiple squares simultaneously

### Finding out whether a square is attacked
- **Pawns**: we're checking the inverse:
```Go
	// attacked by white pawns
	if side == WHITE && (pawnAttacks[BLACK][square]&bitboards[P]) > 0 {
		return 1
  // else, attacked by black pawns
	} else if side == BLACK && (pawnAttacks[WHITE][square]&bitboards[p] > 0) {
		return 1
	}
```
- to see if a square is attacked by a black pawn, we look at white pawn attacks for a square and & it with our black pawns occupancy. If the white pawn can attack our black pawn, the black pawn can also attack that square!
