package main

/*
This is here so that Go's static check doesn't give me warnings
for unused constants or functions in the editor :)
*/
var _ = a8
var _ = b8
var _ = c8
var _ = d8
var _ = e8
var _ = f8
var _ = g8
var _ = h8
var _ = a7
var _ = b7
var _ = c7
var _ = d7
var _ = e7
var _ = f7
var _ = g7
var _ = h7
var _ = a6
var _ = b6
var _ = c6
var _ = d6
var _ = e6
var _ = f6
var _ = g6
var _ = h6
var _ = a5
var _ = b5
var _ = c5
var _ = d5
var _ = e5
var _ = f5
var _ = g5
var _ = h5
var _ = a4
var _ = b4
var _ = c4
var _ = d4
var _ = e4
var _ = f4
var _ = g4
var _ = h4
var _ = a3
var _ = b3
var _ = c3
var _ = d3
var _ = e3
var _ = f3
var _ = g3
var _ = h3
var _ = a2
var _ = b2
var _ = c2
var _ = d2
var _ = e2
var _ = f2
var _ = g2
var _ = h2
var _ = a1
var _ = b1
var _ = c1
var _ = d1
var _ = e1
var _ = f1
var _ = g1
var _ = h1
var _ = WHITE
var _ = BLACK
var _ = ROOK
var _ = BISHOP
var _ = bishopRelevantBits[1]
var _ = rookRelevantBits[1]
var _ = bishopMagicNumbers[1]
var _ = rookMagicNumbers[1]

// Suppress unused function warnings
var _ = findMagicNumber
var _ = initMagicNumbers
var _ = printBitboard
var _ = setOccupancy
var _ = bishopAttacksOnTheFly
var _ = rookAttacksOnTheFly
var _ = maskBishopAttacks
var _ = maskRookAttacks
var _ = maskPawnAttacks
var _ = maskKnightAttacks
var _ = maskKingAttacks
var _ = getRandom32BitUnsignedNumber
var _ = getRandom64BitUnsignedNumber
var _ = generateMagicNumber
var _ = getBit
var _ = setBit
var _ = popBit
var _ = countBits
var _ = getLeastSignificantFirstBitIndex
var _ = initLeapersAttacks
var _ = initAll
var _ = asciiPieces
var _ = unicodePieces
var _ = charPieces
var _ = bitboards
var _ = occupancies
var _ = side
var _ = enpassant
var _ = castle
var _ = printBoard
var _ = parseFEN
var _ = isUpperCase
var _ = isAlpha
var _ = isNumeric
var _ = getBishopAttacks
var _ = getRookAttacks
var _ = getQueenAttacks
var _ = isSquareAttacked
var _ = printAttackedSquares
var _ = generateMoves
var _ = encodeMove
var _ = getMoveSourceSquare
var _ = getMoveTargetSquare
var _ = getMovePiece
var _ = getMovePromotedPiece
var _ = getMoveCaptureFlag
var _ = getMoveDoublePawnPushFlag
var _ = getMoveEnpassantFlag
var _ = getMoveCastlingFlag
var _ = addMove
var _ = printMove
var _ = printMoveList
var _ = copyBoardState
var _ = restorePreviousBoardState
var _ = initSlidersAttacks

// Suppress unused variable warnings
var _ = promotedPieces
var _ = SquareToBigInt
var _ = SquareToCoordinates
var _ = bitboardsCopy
var _ = occupanciesCopy
var _ = sideCopy
var _ = enpassantCopy
var _ = castleCopy
var _ = randomState
var _ = bishopRelevantBits
var _ = rookRelevantBits
var _ = rookMagicNumbers
var _ = bishopMagicNumbers
var _ = pawnAttacks
var _ = knightAttacks
var _ = kingAttacks
var _ = bishopMasks
var _ = rookMasks
var _ = bishopAttacks
var _ = rookAttacks

// func makeMove(move, moveFlag int) int {
// 	// quiet moves
// 	if moveFlag == allMoves {
// 		copyBoardState()

// 		sourceSquare := getMoveSourceSquare(move)
// 		targetSquare := getMoveTargetSquare(move)
// 		piece := getMovePiece(move)
// 		promotedPiece := getMovePromotedPiece(move)
// 		capture := getMoveCaptureFlag(move)
// 		double := getMoveDoublePawnPushFlag(move)
// 		enpass := getMoveEnpassantFlag(move)
// 		castling := getMoveCastlingFlag(move)

// 		// remove piece from source square and update occupancies
// 		bitboards[piece] = popBit(bitboards[piece], sourceSquare)
// 		occupancies[BOTH] = popBit(occupancies[BOTH], sourceSquare)
// 		if side == WHITE {
// 			occupancies[WHITE] = popBit(occupancies[WHITE], sourceSquare)
// 		} else {
// 			occupancies[BLACK] = popBit(occupancies[BLACK], sourceSquare)
// 		}

// 		// handle capture moves
// 		if capture > 0 {
// 			var startPiece int
// 			var endPiece int

// 			// loop over piece bitboards to remove captured piece from board
// 			if side == WHITE {
// 				startPiece = p
// 				endPiece = k
// 			} else {
// 				startPiece = P
// 				endPiece = K
// 			}
// 			for bbPiece := startPiece; bbPiece <= endPiece; bbPiece++ {
// 				// if there is a piece on target square,
// 				if getBit(bitboards[bbPiece], targetSquare) > 0 {
// 					bitboards[bbPiece] = popBit(bitboards[bbPiece], targetSquare)
// 					occupancies[BOTH] = popBit(occupancies[BOTH], targetSquare)
// 					if bbPiece >= P && bbPiece <= K {
// 						occupancies[WHITE] = popBit(occupancies[WHITE], targetSquare)
// 					} else {
// 						occupancies[BLACK] = popBit(occupancies[BLACK], targetSquare)
// 					}
// 					break
// 				}
// 			}
// 		}

// 		// pawn promotions
// 		if promotedPiece > 0 {
// 			bitboards[promotedPiece] = setBit(bitboards[promotedPiece], targetSquare)
// 			occupancies[BOTH] = setBit(occupancies[BOTH], targetSquare)
// 			if side == WHITE {
// 				occupancies[WHITE] = setBit(occupancies[WHITE], targetSquare)
// 			} else {
// 				occupancies[BLACK] = setBit(occupancies[BLACK], targetSquare)
// 			}
// 		} else {
// 			bitboards[piece] = setBit(bitboards[piece], targetSquare)
// 			occupancies[BOTH] = setBit(occupancies[BOTH], targetSquare)
// 			if side == WHITE {
// 				occupancies[WHITE] = setBit(occupancies[WHITE], targetSquare)
// 			} else {
// 				occupancies[BLACK] = setBit(occupancies[BLACK], targetSquare)
// 			}
// 		}

// 		// en passant capture
// 		if enpass > 0 {
// 			var captureSquare int
// 			var capturePiece int
// 			if side == WHITE {
// 				captureSquare = targetSquare + 8
// 				capturePiece = p
// 			} else {
// 				captureSquare = targetSquare - 8
// 				capturePiece = P
// 			}
// 			bitboards[capturePiece] = popBit(bitboards[capturePiece], captureSquare)
// 			occupancies[BOTH] = popBit(occupancies[BOTH], captureSquare)
// 			if side == WHITE {
// 				occupancies[BLACK] = popBit(occupancies[BLACK], captureSquare)
// 			} else {
// 				occupancies[WHITE] = popBit(occupancies[WHITE], captureSquare)
// 			}
// 		}
// 		enpassant = NO_SQ

// 		// double pawn push
// 		if double > 0 {
// 			// setup enpassant square
// 			if side == WHITE {
// 				enpassant = targetSquare + 8
// 			} else {
// 				enpassant = targetSquare - 8
// 			}
// 		}

// 		// castling - move rook and update occupancies
// 		if castling > 0 {
// 			switch targetSquare {
// 			case g1:
// 				bitboards[R] = popBit(bitboards[R], h1)
// 				bitboards[R] = setBit(bitboards[R], f1)
// 				occupancies[BOTH] = popBit(occupancies[BOTH], h1)
// 				occupancies[BOTH] = setBit(occupancies[BOTH], f1)
// 				occupancies[WHITE] = popBit(occupancies[WHITE], h1)
// 				occupancies[WHITE] = setBit(occupancies[WHITE], f1)
// 			case c1:
// 				bitboards[R] = popBit(bitboards[R], a1)
// 				bitboards[R] = setBit(bitboards[R], d1)
// 				occupancies[BOTH] = popBit(occupancies[BOTH], a1)
// 				occupancies[BOTH] = setBit(occupancies[BOTH], d1)
// 				occupancies[WHITE] = popBit(occupancies[WHITE], a1)
// 				occupancies[WHITE] = setBit(occupancies[WHITE], d1)
// 			case g8:
// 				bitboards[r] = popBit(bitboards[r], h8)
// 				bitboards[r] = setBit(bitboards[r], f8)
// 				occupancies[BOTH] = popBit(occupancies[BOTH], h8)
// 				occupancies[BOTH] = setBit(occupancies[BOTH], f8)
// 				occupancies[BLACK] = popBit(occupancies[BLACK], h8)
// 				occupancies[BLACK] = setBit(occupancies[BLACK], f8)
// 			case c8:
// 				bitboards[r] = popBit(bitboards[r], a8)
// 				bitboards[r] = setBit(bitboards[r], d8)
// 				occupancies[BOTH] = popBit(occupancies[BOTH], a8)
// 				occupancies[BOTH] = setBit(occupancies[BOTH], d8)
// 				occupancies[BLACK] = popBit(occupancies[BLACK], a8)
// 				occupancies[BLACK] = setBit(occupancies[BLACK], d8)
// 			}
// 		}

// 		// update castling rights
// 		castle &= castlingRights[sourceSquare]
// 		castle &= castlingRights[targetSquare]

// 		// check if king is in check (WIP)!
// 		side ^= 1

// 		if side == WHITE && isSquareAttacked(getLeastSignificantFirstBitIndex(bitboards[k]), side) > 0 {
// 			restorePreviousBoardState()
// 			return -1
// 		} else if side == BLACK && isSquareAttacked(getLeastSignificantFirstBitIndex(bitboards[K]), side) > 0 {
// 			restorePreviousBoardState()
// 			return -1
// 		} else {
// 			return 1
// 		}
// 	} else {
// 		// capture moves
// 		// make sure move is capture
// 		if getMoveCaptureFlag(move) > 0 {
// 			makeMove(move, allMoves)
// 		} else {
// 			return -1
// 		}
// 	}
// 	return 1
// }
