# Bytes Layout

| nonce_1 (12 bytes) | checksum_0 (26 bytes) | nonce_2 (12 bytes) | checksum_1 (16 bytes) | cipher_1 | cipher_2 |

## checksum

encrypted("ok" + <cipher_0_from> + <cipher_0_to> (2+4+4 bytes)) = checksum_0 26 bytes