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
018898f3-3067-703b-8292-e2def527c6b4
> select quote(uuid_generate_v7_bytes());
X'018898F33745703BA7618E0C33AECF3E'
> .quit
```
