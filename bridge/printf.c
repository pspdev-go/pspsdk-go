#include <pspdebug.h>
#include <pspkdebug.h>
#include <pspstdio_kernel.h>

void psp_debugscreen_kputs(const char *text) {
    pspDebugScreenKprintf("%s", text);
}

void psp_kdebug_puts(const char *text) {
    Kprintf("%s", text);
}

int pspsdk_go_fdputs(int fd, const char *text) {
    return fdprintf(fd, "%s", text);
}
