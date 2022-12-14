# How to pack

## Bytes Diagram

```
  0                   1  
  0 1 2 3 4 5 6 7 8 9 0 1 2 (byte)
0 +-+-+-+-+-+-+-+-+-+-+-+-+
  |        nonce 0        |
1 +-+-+-+-+-+-+-+-+-+-+-+-+
  |                       |
2 +       checksum 0      +
  |                       |
3 +-+-+-+-+-+-+-+-+-+-+-+-+
  |        nonce 1        |
4 +-+-+-+-+-+-+-+-+-+-+-+-+
  |                       |
5 +       checksum 1      +
  |                       |
6 +-+-+-+-+-+-+-+-+-+-+-+-+ `from_0`
. |       cipher 0...     |
. +       arbitrary       +
. |        length         |
. +-+-+-+-+-+-+-+-+-+-+-+-+ `to_0` = `from_1`
  |       cipher 1...     |
  +       arbitrary       +
  |        length         |
  +-+-+-+-+-+-+-+-+-+-+-+-+ `to_1`
```

## Checksum

`checksum` is the part in the header. It is the encrypted result of cipher start index, `from`, and end index, `to`, in the final binary.

`from` and `to` are byte arrays which are converted from two uint32 variables, so once concatenate them, the result is an 8 bytes array (let's call it `cipher_info` here).

The 8 bytes `cipher_info` will be encrypted in the same encryption algorithm and nonce that corresponding cipher is used. The final result is the `checksum` which is 24 bytes long.

A simple equation is provided as following:

```
Enc(<cipher_0_from> | <cipher_0_to>) = checksum_0
```
