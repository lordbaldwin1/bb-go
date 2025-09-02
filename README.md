# bb-go

## Bitwise notes for bitboards
`|=` will logical OR all of the bits from the two boards and assign that result to the bitboard left of the operator.

`&` result is one if both operands have 1 in the same position!
0 & 0 = 0
0 & 1 = 0
1 & 0 = 0
1 & 1 = 1

Finding the index of the least significant bit!

First, singling out the least significant bit:
`block & (~block + 1n)`
`~block` will flip all of the bits of block
`~block + 1n` will, as before, flip all of the bits of block. This means when we add `1n`, all of the `1n` bits until the first `0n` will be flipped to zero and the first zero will be flipped to 1. 
`block & ~block + 1n` So, when we `&` block with it, the only remaining bit that is equal between the two is the least significant one (important: 0 & 0 = 0).
`(block & ~block + 1n) - 1n` basically will flip our singled out bit to 0 and turn all the preceding bits to 1.
```TS
  let block = 0n;
  block = setBit(block, f5);
  block = setBit(block, d8);
  block = setBit(block, c5);
  block = setBit(block, d2);
  block = setBit(block, d3);
  printBitboard(~block);
  printBitboard(~block + 1n);
  printBitboard((block & ~block + 1n) - 1n);
``` 

- look into faster bit manipulations using hardware implementations? BMI?
