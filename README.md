# sqlite-uuidv7

An SQLite extension for generating
[UUIDv7s](https://datatracker.ietf.org/doc/html/draft-peabody-dispatch-new-uuid-format).
It wraps https://github.com/gofrs/uuid.

## Usage

You will need a version of SQLite that allows loading of extensions. If you are
on a Mac, you can `brew install sqlite`.

```text
$ make 
$ sqlite3
> .load uuidv7.so
> select uuid_generate_v7();
01889914-efa9-7ec3-997b-c2b1de310c56
> select hex(uuid_generate_v7_bytes());
0188991513087EC3B78D3FD9361C5B32
> .quit
```
