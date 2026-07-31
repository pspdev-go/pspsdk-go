#include <stddef.h>
#include <pspkernel.h>
#include <pspgum.h>
#include <vlf.h>

int pspsdk_go_vlf_add_text(int x, int y, const char *text) {
    return vlfGuiAddTextF(x, y, "%s", text);
}

int pspsdk_go_vlf_set_text(int text, const char *value) {
    return vlfGuiSetTextF(text, "%s", value);
}
