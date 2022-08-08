![dualbox logo](dualbox.jpg)

# dualbox

Encrypt given data all together by different keys and decrypt them depends on the given keys.

## How to install

Go to project root and run:

```bash
$ go install
```

## How to use

To encrypt files run:

```bash
$ dualbox enc --file0 $FILE_PATH_0 --file1 $FILE_PATH_1 --password0 $PASSWORD_0 --password1 $PASSWORD_1
```

To decrypt file run:

```bash
$ dualbox dec --file $FILE_PATH --password $PASSWORD
```

To see more command information run:

```bash
$ dualbox --help
```