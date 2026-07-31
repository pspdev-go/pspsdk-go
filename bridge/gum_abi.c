#include <stddef.h>

#include <pspgu.h>
#include <pspgum.h>

typedef struct PspsdkGoGumDraw3D {
    int prim;
    int vtype;
    int count;
    const void *vertices;
    float fovy;
    float aspect;
    float near_z;
    float far_z;
    ScePspFVector3 translation;
    ScePspFVector3 rotation;
} PspsdkGoGumDraw3D;

/*
 * TinyGo's MIPS calling convention places the fifth argument on the stack,
 * while the PSP toolchain passes arguments five through eight in t0-t3.
 * Keep the Go-facing entry point at four arguments and let GCC perform the
 * PSP ABI call to sceGumDrawArray.
 */
void pspsdk_go_gum_draw_array(int prim, int vtype, int count,
                              const void *vertices) {
    sceGumDrawArray(prim, vtype, count, NULL, vertices);
}

void pspsdk_go_gum_draw_array_3d(const PspsdkGoGumDraw3D *draw) {
    sceGumMatrixMode(GU_PROJECTION);
    sceGumLoadIdentity();
    sceGumPerspective(draw->fovy, draw->aspect, draw->near_z, draw->far_z);

    sceGumMatrixMode(GU_VIEW);
    sceGumLoadIdentity();

    sceGumMatrixMode(GU_MODEL);
    sceGumLoadIdentity();
    sceGumTranslate(&draw->translation);
    sceGumRotateXYZ(&draw->rotation);

    sceGumDrawArray(
        draw->prim, draw->vtype, draw->count, NULL, draw->vertices);
}
