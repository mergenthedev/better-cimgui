package bettercimgui

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -l:cimgui.so -lglfw
#define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
#define CIMGUI_USE_GLFW
#define CIMGUI_USE_OPENGL3
#include "cimgui.h"
#include "cimgui_impl.h"
#include <GLFW/glfw3.h>
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

	glslVer := C.CString("#version 330 core")
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

func Demo() {
	C.igBegin(C.CString("Test"), nil, 0)
	//C.igText(C.CString("Hello"))
	C.igEnd()
}

/*

//
#cgo CFLAGS: -x c++ -std=c++11 -fpermissive -I./imgui
#define CIMGUI_DEFINE_ENUMS_AND_STRUCTS
#define CIMGUI_USE_OPENGL3
#define CIMGUI_USE_GLFW
#define IMGUI_DEFINE_MATH_OPERATORS
#include "imgui/cimgui.h"
#include "imgui/cimgui_impl.h"
#include "imgui/cimgui.cpp"
#include "imgui/cimgui_impl.cpp"
#include "imgui/imgui/imgui.h"
#include "imgui/imgui/imgui.cpp"
#include "imgui/imgui_impl_glfw.h"
#include "imgui/imgui_impl_opengl3.h"
#include "imgui/imgui_demo.cpp"
#include "imgui/imgui_draw.cpp"
#include "imgui/imgui_impl_glfw.cpp"
#include "imgui/imgui_impl_opengl3.cpp"
#include "imgui/imgui_tables.cpp"
#include "imgui/imgui_widgets.cpp"
#include <stdlib.h>
*/
