# How to pack

## Bytes Layout

| nonce_1 (12 bytes) | checksum_0 (24 bytes) | nonce_2 (12 bytes) | checksum_1 (24 bytes) | cipher_1 | cipher_2 |

## Checksum

encrypted(<cipher_0_from> + <cipher_0_to> (4+4 bytes)) = checksum_0 (24 bytes)
