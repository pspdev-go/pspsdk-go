package gu

// Common pspgu.h macros used to configure 3D rendering.
const (
	False int32 = 0
	True  int32 = 1

	Triangles int32 = 3

	DepthTest   int32 = 1
	ScissorTest int32 = 2
	CullFace    int32 = 5
	ClipPlanes  int32 = 8

	Projection int32 = 0
	View       int32 = 1
	Model      int32 = 2

	Color8888    int32 = 7 << 2
	Vertex32BitF int32 = 3 << 7
	Transform3D  int32 = 0 << 23

	DisplayOn int32 = 1

	PSM4444 int32 = 2
	PSM8888 int32 = 3

	Smooth int32 = 1
	CW     int32 = 0
	GEqual int32 = 7

	ColorBufferBit int32 = 1
	DepthBufferBit int32 = 4

	Direct       int32 = 0
	SyncFinish   int32 = 0
	SyncWhatDone int32 = 0
)

const (
	ScreenWidth  int32 = 480
	ScreenHeight int32 = 272
	VRAMWidth    int32 = 512
)

// PSPSDK-compatible macro names.
const (
	GU_FALSE            = False
	GU_TRUE             = True
	GU_TRIANGLES        = Triangles
	GU_DEPTH_TEST       = DepthTest
	GU_SCISSOR_TEST     = ScissorTest
	GU_CULL_FACE        = CullFace
	GU_CLIP_PLANES      = ClipPlanes
	GU_PROJECTION       = Projection
	GU_VIEW             = View
	GU_MODEL            = Model
	GU_COLOR_8888       = Color8888
	GU_VERTEX_32BITF    = Vertex32BitF
	GU_TRANSFORM_3D     = Transform3D
	GU_DISPLAY_ON       = DisplayOn
	GU_PSM_4444         = PSM4444
	GU_PSM_8888         = PSM8888
	GU_SMOOTH           = Smooth
	GU_CW               = CW
	GU_GEQUAL           = GEqual
	GU_COLOR_BUFFER_BIT = ColorBufferBit
	GU_DEPTH_BUFFER_BIT = DepthBufferBit
	GU_DIRECT           = Direct
	GU_SYNC_FINISH      = SyncFinish
	GU_SYNC_WHAT_DONE   = SyncWhatDone
)
