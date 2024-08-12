# Unescape HTML strings with html2utf8

Convert HTML characters to UTF-8.

This UDF example can be used in MariaDB or from a Python script.

Build and install:
   make
   make install

Usage:

    MariaDB> SELECT html2utf8(my_field);

The `main.py` script demonstrates re-using the same code again from Python.
