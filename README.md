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
$ dualbox enc --file0 <file_path_0> --file1 <file_path_1> --key0 <key_0> --key1 <key_1>
```

To decrypt file run:

```bash
$ dualbox dec --file <file_path> --key <key>
```

To see more command information run:

```bash
$ dualbox --help
```