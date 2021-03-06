import ctypes
import os

lib = ctypes.CDLL(os.path.abspath("./mylib.so"))


class GoString(ctypes.Structure):
    _fields_ = [("p", ctypes.c_char_p), ("n", ctypes.c_longlong)]


class Coords(ctypes.Structure):
    _fields_ = [("x", ctypes.c_longlong), ("y", ctypes.c_longlong)]


msg = GoString("Hello Python!", 13)

lib.PrintGoStr(msg)
lib.PrintInt(123)
lib.PrintCStr("Hello C!")
lib.PrintStruct(Coords(1, 2))
lib.PrintStructMeth(Coords(3, 4))
# lib.PrintMap({1: 2, 3: 4})

file = os.path.abspath(os.path.join(
    os.path.dirname(os.path.realpath(__file__)),
    "../testdata.json"))

# https://docs.python.org/3/library/ctypes.html#return-types
lib.GetFirstJSONElement.restype = ctypes.c_char_p

print lib.GetFirstJSONElement(file)
