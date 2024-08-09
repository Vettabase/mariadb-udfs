#!/usr/bin/env python3

import ctypes
import fileinput

# Load the shared object file
lib = ctypes.CDLL("build/html2utf8.so")

# Define the function prototype
lib.Html2utf8_raw.argtypes = [ctypes.c_char_p]
lib.Html2utf8_raw.restype = ctypes.c_char_p

for line in fileinput.input():
    l = line.encode()
    r = lib.Html2utf8_raw(l)
    d = r.decode()
    print(d)
