package gum

import "unsafe"

// Draw3D describes a non-indexed 3D draw and its perspective/model transform.
// Its layout is shared with the PSP GCC ABI adapter.
type Draw3D struct {
	Prim        int32
	VType       int32
	Count       int32
	Vertices    unsafe.Pointer
	FOVY        float32
	Aspect      float32
	Near        float32
	Far         float32
	Translation Vector3
	Rotation    Vector3
}

// DrawArray draws a non-indexed vertex array through an ABI adapter.
//
// PSP's GCC ABI passes the fifth C argument in t0, whereas TinyGo's MIPS ABI
// places it on the stack. The adapter keeps the Go-facing call at four
// arguments so vertex pointers are passed correctly.
func DrawArray(prim int32, vtype int32, count int32, vertices unsafe.Pointer) {
	drawArrayABI(prim, vtype, count, vertices)
}

// DrawArray3D applies the configured perspective/model transform and draws a
// non-indexed vertex array. All GUM calls execute on the GCC side so both
// floating-point and extended argument conventions match the PSP ABI.
func DrawArray3D(draw *Draw3D) {
	drawArray3DABI(unsafe.Pointer(draw))
}

//go:linkname drawArrayABI pspsdk_go_gum_draw_array
func drawArrayABI(prim int32, vtype int32, count int32, vertices unsafe.Pointer)

//go:linkname drawArray3DABI pspsdk_go_gum_draw_array_3d
func drawArray3DABI(draw unsafe.Pointer)
