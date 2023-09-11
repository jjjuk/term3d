package win

/*
#include <stdio.h>
#include <stdlib.h>
#include <windows.h>

void writeToTerminal(HANDLE handle, int x, int y, LPCSTR str, int len) {
  COORD zero = {x, y};
  DWORD written = 0;

  WriteConsoleOutputCharacterA(handle, str, len, zero, &written);
}
*/
import "C"
import (
	"unsafe"
)

type Win struct {
	handle C.HANDLE
}

func New() *Win {
	return &Win{handle: C.GetStdHandle(C.STD_OUTPUT_HANDLE)}
}

func (w *Win) WriteTermOutput(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.writeToTerminal(w.handle, C.int(0), C.int(0), c_text, C.int(len(text)))
}
