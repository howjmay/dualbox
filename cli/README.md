# CLI

## Commands

```bash
$ dualbox enc --file0 $TARGET_FILE_0_PATH --file1 $TARGET_FILE_1_PATH --password0 $PASSWORD_0 --password1 $PASSWORD_1
$ dualbox enc --file0 $TARGET_FILE_0_PATH --file1 $TARGET_FILE_1_PATH --password-file $ENCRYPTION_TARGETS_FILE_PATH
$ dualbox enc --read $ENCRYPTION_TARGETS_FILE_PATH
$ dualbox dec --file $TARGET_FILE_PATH --password $PASSWORD
```

### encryption targets file example
```yaml
files:
    # file 0
    - ./testdata/testdata0.jpg
    # file 1
    - ./testdata/testdata1.png
passwords:
    # password for file 0
    - this_is_password_0
    # password for file 1
    - this_is_password_1
# use either `password` of `passwords`
passwords:
    # password for file 0
    - f79681852aad0428a5005a4dc0c25404bb3c3c2b387410a53cf6253e09e416db
    # password for file 1
    - 3aaf57cdd0fc902048388f7ebf9fe4f175ca3182c23109734e97323b5d719ab7
```
