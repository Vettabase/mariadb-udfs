# MariaDB Utility Functions

Created in Go for MariaDB, a small collection of functions to interace with Linux and the Filesystem from within MariaDB.

Functions:

* `bytesfree(/path/to/mount)`, get the number of bytes free from the provided volume path. Returns `-1` if the path does not exist or is unreachable.
* `deletefile(/tmp/my_file)`, delete a file on disk! Must be within `/tmp/` and accessible by the server process user. Returns `0` on success.
* `gzipfile(/tmp/my_file)`, compress a file with gzip, file must be within `/tmp/`. Returns `0` on success.
* `filexists(/path/to/file)`, return `0` if an accessible file exists.
