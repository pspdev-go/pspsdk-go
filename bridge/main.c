#include <pspkernel.h>
#include <pspuser.h>

PSP_MODULE_INFO("TinyGoPSP", 0, 1, 0);
PSP_MAIN_THREAD_ATTR(PSP_THREAD_ATTR_USER);

extern void tinygo_psp_start(void);

int main(void) {
    tinygo_psp_start();
    return 0;
}
