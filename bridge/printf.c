#include <pspdebug.h>
#include <pspkdebug.h>

void psp_debugscreen_kputs(const char *text) {
    pspDebugScreenKprintf("%s", text);
}

void psp_kdebug_puts(const char *text) {
    Kprintf("%s", text);
}
