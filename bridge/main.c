#include <pspkernel.h>
#include <pspuser.h>

#ifdef PSPSDK_GO_KERNEL_MODE
PSP_MODULE_INFO("TinyGoPSP", PSP_MODULE_KERNEL, 1, 0);
PSP_MAIN_THREAD_ATTR(0);
#else
PSP_MODULE_INFO("TinyGoPSP", PSP_MODULE_USER, 1, 0);
PSP_MAIN_THREAD_ATTR(PSP_THREAD_ATTR_USER);
#endif

extern void tinygo_psp_start(void);

int main(void) {
    tinygo_psp_start();
    return 0;
}
