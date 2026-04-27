package bettercimgui

/*
#define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
#define CIMGUI_USE_OPENGL3
#define CIMGUI_USE_GLFW
#include "imgui/cimgui.h"
#include "imgui/cimgui_impl.h"
#include <stdlib.h>
*/
import "C"
import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/go-gl/glfw/v3.4/glfw"
)

func init() {
	runtime.LockOSThread()
}

var Context *C.ImGuiContext
var IO *C.ImGuiIO

func GetGLFWWindowCPtr(window *glfw.Window) *C.GLFWwindow {
	w := reflect.ValueOf(window).Elem()
	wData := w.Field(0)

	return (*C.GLFWwindow)(unsafe.Pointer(wData.UnsafeAddr()))
}

func Init(window *C.GLFWwindow) {
	Context = C.igCreateContext(nil)
	IO = C.igGetIO_Nil()

	glslVer := C.CString("#version 410 core")
	C.ImGui_ImplGlfw_InitForOpenGL(window, C.bool(true))
	C.ImGui_ImplOpenGL3_Init(glslVer)
	C.free(unsafe.Pointer(glslVer))

	C.igStyleColorsDark(nil)
}

func Terminate() {
	C.ImGui_ImplOpenGL3_Shutdown()
	C.ImGui_ImplGlfw_Shutdown()
	C.igDestroyContext(Context)
}

func Update() {
	C.ImGui_ImplOpenGL3_NewFrame()
	C.ImGui_ImplGlfw_NewFrame()
	C.igNewFrame()
}

func Render() {
	C.igRender()
	C.ImGui_ImplOpenGL3_RenderDrawData(C.igGetDrawData())
}
