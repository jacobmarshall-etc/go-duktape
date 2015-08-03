package duktape

/*
# include "duktape.h"
static void _duk_eval_string(duk_context *ctx, const char *str) {
  return duk_eval_string(ctx, str);
}
static void _duk_compile(duk_context *ctx, duk_uint_t flags) {
  return duk_compile(ctx, flags);
}
static void _duk_compile_file(duk_context *ctx, duk_uint_t flags, const char *path) {
  return duk_compile_file(ctx, flags, path);
}
static void _duk_compile_lstring(duk_context *ctx, duk_uint_t flags, const char *src, duk_size_t len) {
	return duk_compile_lstring(ctx, flags, src, len);
}
static void _duk_compile_lstring_filename(duk_context *ctx, duk_uint_t flags, const char *src, duk_size_t len) {
	return duk_compile_lstring_filename(ctx, flags, src, len);
}
static void _duk_compile_string(duk_context *ctx, duk_uint_t flags, const char *src) {
	return duk_compile_string(ctx, flags, src);
}
static void _duk_compile_string_filename(duk_context *ctx, duk_uint_t flags, const char *src) {
	return duk_compile_string_filename(ctx, flags, src);
}
static void _duk_dump_context_stderr(duk_context *ctx) {
	return duk_dump_context_stderr(ctx);
}
static void _duk_dump_context_stdout(duk_context *ctx) {
	return duk_dump_context_stdout(ctx);
}
static void _duk_eval(duk_context *ctx) {
	return duk_eval(ctx);
}
static void _duk_eval_file(duk_context *ctx, const char *path) {
	return duk_eval_file(ctx, path);
}
static void _duk_eval_file_noresult(duk_context *ctx, const char *path) {
	return duk_eval_file_noresult(ctx, path);
}
static void _duk_eval_lstring(duk_context *ctx, const char *src, duk_size_t len) {
	return duk_eval_lstring(ctx, src, len);
}
static void _duk_eval_lstring_noresult(duk_context *ctx, const char *src, duk_size_t len) {
	return duk_eval_lstring_noresult(ctx, src, len);
}
static void _duk_eval_noresult(duk_context *ctx) {
	return duk_eval_noresult(ctx);
}
static void _duk_eval_string_noresult(duk_context *ctx, const char *src) {
	return duk_eval_string_noresult(ctx, src);
}
static duk_bool_t _duk_is_error(duk_context *ctx, duk_idx_t index) {
	return duk_is_error(ctx, index);
}
static duk_bool_t _duk_is_object_coercible(duk_context *ctx, duk_idx_t index) {
	return duk_is_object_coercible(ctx, index);
}
static duk_int_t _duk_pcompile(duk_context *ctx, duk_uint_t flags) {
	return duk_pcompile(ctx, flags);
}
static duk_int_t _duk_pcompile_file(duk_context *ctx, duk_uint_t flags, const char *path) {
	return duk_pcompile_file(ctx, flags, path);
}
static duk_int_t _duk_pcompile_lstring(duk_context *ctx, duk_uint_t flags, const char *src, duk_size_t len) {
	return duk_pcompile_lstring(ctx, flags, src, len);
}
static duk_int_t _duk_pcompile_lstring_filename(duk_context *ctx, duk_uint_t flags, const char *src, duk_size_t len) {
	return duk_pcompile_lstring_filename(ctx, flags, src, len);
}
static duk_int_t _duk_pcompile_string(duk_context *ctx, duk_uint_t flags, const char *src) {
	return duk_pcompile_string(ctx, flags, src);
}
static duk_int_t _duk_pcompile_string_filename(duk_context *ctx, duk_uint_t flags, const char *src) {
	return duk_pcompile_string_filename(ctx, flags, src);
}
static duk_int_t _duk_peval(duk_context *ctx) {
	return duk_peval(ctx);
}
static duk_int_t _duk_peval_file(duk_context *ctx, const char *path) {
	return duk_peval_file(ctx, path);
}
static duk_int_t _duk_peval_file_noresult(duk_context *ctx, const char *path) {
	return duk_peval_file_noresult(ctx, path);
}
static duk_int_t _duk_peval_lstring(duk_context *ctx, const char *src, duk_size_t len) {
	return duk_peval_lstring(ctx, src, len);
}
static duk_int_t _duk_peval_lstring_noresult(duk_context *ctx, const char *src, duk_size_t len) {
	return duk_peval_lstring_noresult(ctx, src, len);
}
static duk_int_t _duk_peval_noresult(duk_context *ctx) {
	return duk_peval_noresult(ctx);
}
static duk_int_t _duk_peval_string(duk_context *ctx, const char *src) {
	return duk_peval_string(ctx, src);
}
static duk_int_t _duk_peval_string_noresult(duk_context *ctx, const char *src) {
	return duk_peval_string_noresult(ctx, src);
}
static const char *_duk_push_string_file(duk_context *ctx, const char *path) {
	return duk_push_string_file(ctx, path);
}
static duk_idx_t _duk_push_thread(duk_context *ctx) {
	return duk_push_thread(ctx);
}
static duk_idx_t _duk_push_thread_new_globalenv(duk_context *ctx) {
	return duk_push_thread_new_globalenv(ctx);
}
static void _duk_require_object_coercible(duk_context *ctx, duk_idx_t index) {
	return duk_require_object_coercible(ctx, index);
}
static void _duk_require_type_mask(duk_context *ctx, duk_idx_t index, duk_uint_t mask) {
	return duk_require_type_mask(ctx, index, mask);
}
static const char *_duk_safe_to_string(duk_context *ctx, duk_idx_t index) {
	return duk_safe_to_string(ctx, index);
}
static void _duk_xcopy_top(duk_context *to_ctx, duk_context *from_ctx, duk_idx_t count) {
	return duk_xcopy_top(to_ctx, from_ctx, count);
}
static void _duk_xmove_top(duk_context *to_ctx, duk_context *from_ctx, duk_idx_t count) {
	return duk_xmove_top(to_ctx, from_ctx, count);
}
static void *_duk_to_buffer(duk_context *ctx, duk_idx_t index, duk_size_t *out_size) {
	return duk_to_buffer(ctx, index, out_size);
}
static void *_duk_to_dynamic_buffer(duk_context *ctx, duk_idx_t index, duk_size_t *out_size) {
	return duk_to_dynamic_buffer(ctx, index, out_size);
}
static void *_duk_to_fixed_buffer(duk_context *ctx, duk_idx_t index, duk_size_t *out_size) {
	return duk_to_fixed_buffer(ctx, index, out_size);
}
static void *_duk_push_buffer(duk_context *ctx, duk_size_t size, duk_bool_t dynamic) {
	return duk_push_buffer(ctx, size, dynamic);
}
static void *_duk_push_fixed_buffer(duk_context *ctx, duk_size_t size) {
	return duk_push_fixed_buffer(ctx, size);
}
static void *_duk_push_dynamic_buffer(duk_context *ctx, duk_size_t size) {
	return duk_push_dynamic_buffer(ctx, size);
}
static void _duk_error(duk_context *ctx, duk_errcode_t err_code, const char *str) {
	duk_error(ctx, err_code, "%s", str);
}
static void _duk_push_error_object(duk_context *ctx, duk_errcode_t err_code, const char *str) {
	duk_push_error_object(ctx, err_code, "%s", str);
}
static void _duk_error_raw(duk_context *ctx, duk_errcode_t err_code, const char *filename, duk_int_t line, const char *text) {
	duk_error_raw(ctx, err_code, filename, line, text);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// See: http://duktape.org/api.html#duk_alloc
func (d *Context) Alloc(size int) {
	mustNotNil(d)
	C.duk_alloc(d.duk_context, C.duk_size_t(size))
}

// See: http://duktape.org/api.html#duk_alloc_raw
func (d *Context) AllocRaw(size int) {
	mustNotNil(d)
	C.duk_alloc_raw(d.duk_context, C.duk_size_t(size))
}

// See: http://duktape.org/api.html#duk_base64_decode
func (d *Context) Base64Decode(index int) {
	mustNotNil(d)
	C.duk_base64_decode(d.duk_context, C.duk_idx_t(index))
}

// See: http://duktape.org/api.html#duk_base64_encode
func (d *Context) Base64Encode(index int) string {
	mustNotNil(d)
	if s := C.duk_base64_encode(d.duk_context, C.duk_idx_t(index)); s != nil {
		return C.GoString(s)
	}
	return ""
}

// See: http://duktape.org/api.html#duk_call
func (d *Context) Call(nargs int) {
	mustNotNil(d)
	C.duk_call(d.duk_context, C.duk_idx_t(nargs))
}

// See: http://duktape.org/api.html#duk_call_method
func (d *Context) CallMethod(nargs int) {
	mustNotNil(d)
	C.duk_call_method(d.duk_context, C.duk_idx_t(nargs))
}

// See: http://duktape.org/api.html#duk_call_prop
func (d *Context) CallProp(objIndex int, nargs int) {
	mustNotNil(d)
	C.duk_call_prop(d.duk_context, C.duk_idx_t(objIndex), C.duk_idx_t(nargs))
}

// See: http://duktape.org/api.html#duk_check_stack
func (d *Context) CheckStack(extra int) bool {
	mustNotNil(d)
	return int(C.duk_check_stack(d.duk_context, C.duk_idx_t(extra))) == 1
}

// See: http://duktape.org/api.html#duk_check_stack_top
func (d *Context) CheckStackTop(top int) bool {
	mustNotNil(d)
	return int(C.duk_check_stack_top(d.duk_context, C.duk_idx_t(top))) == 1
}

// See: http://duktape.org/api.html#duk_check_type
func (d *Context) CheckType(index int, typ int) bool {
	mustNotNil(d)
	return int(C.duk_check_type(d.duk_context, C.duk_idx_t(index), C.duk_int_t(typ))) == 1
}

// See: http://duktape.org/api.html#duk_check_type_mask
func (d *Context) CheckTypeMask(index int, mask uint) bool {
	mustNotNil(d)
	return int(C.duk_check_type_mask(d.duk_context, C.duk_idx_t(index), C.duk_uint_t(mask))) == 1
}

// See: http://duktape.org/api.html#duk_compact
func (d *Context) Compact(objIndex int) {
	mustNotNil(d)
	C.duk_compact(d.duk_context, C.duk_idx_t(objIndex))
}

// See: http://duktape.org/api.html#duk_compile
func (d *Context) Compile(flags uint) {
	mustNotNil(d)
	C._duk_compile(d.duk_context, C.duk_uint_t(flags))
}

// See: http://duktape.org/api.html#duk_compile_file
func (d *Context) CompileFile(flags uint, path string) {
	mustNotNil(d)
	__path__ := C.CString(path)
	defer C.free(unsafe.Pointer(__path__))
	C._duk_compile_file(d.duk_context, C.duk_uint_t(flags), __path__)
}

// See: http://duktape.org/api.html#duk_compile_lstring
func (d *Context) CompileLstring(flags uint, src string, len int) {
	mustNotNil(d)
	__src__ := C.CString(src)
	defer C.free(unsafe.Pointer(__src__))
	C._duk_compile_lstring(d.duk_context, C.duk_uint_t(flags), __src__, C.duk_size_t(len))
}

// See: http://duktape.org/api.html#duk_compile_lstring_filename
func (d *Context) CompileLstringFilename(flags uint, src string, len int) {
	mustNotNil(d)
	__src__ := C.CString(src)
	defer C.free(unsafe.Pointer(__src__))
	C._duk_compile_lstring_filename(d.duk_context, C.duk_uint_t(flags), __src__, C.duk_size_t(len))
}

// See: http://duktape.org/api.html#duk_compile_string
func (d *Context) CompileString(flags uint, src string) {
	mustNotNil(d)
	__src__ := C.CString(src)
	defer C.free(unsafe.Pointer(__src__))
	C._duk_compile_string(d.duk_context, C.duk_uint_t(flags), __src__)
}

// See: http://duktape.org/api.html#duk_compile_string_filename
func (d *Context) CompileStringFilename(flags uint, src string) {
	mustNotNil(d)
	__src__ := C.CString(src)
	defer C.free(unsafe.Pointer(__src__))
	C._duk_compile_string_filename(d.duk_context, C.duk_uint_t(flags), __src__)
}

// See: http://duktape.org/api.html#duk_concat
func (d *Context) Concat(count int) {
	mustNotNil(d)
	C.duk_concat(d.duk_context, C.duk_idx_t(count))
}

// See: http://duktape.org/api.html#duk_copy
func (d *Context) Copy(fromIndex int, toIndex int) {
	mustNotNil(d)
	C.duk_copy(d.duk_context, C.duk_idx_t(fromIndex), C.duk_idx_t(toIndex))
}

// See: http://duktape.org/api.html#duk_del_prop
func (d *Context) DelProp(objIndex int) bool {
	mustNotNil(d)
	return int(C.duk_del_prop(d.duk_context, C.duk_idx_t(objIndex))) == 1
}

// See: http://duktape.org/api.html#duk_del_prop_index
func (d *Context) DelPropIndex(objIndex int, arrIndex uint) bool {
	mustNotNil(d)
	return int(C.duk_del_prop_index(d.duk_context, C.duk_idx_t(objIndex), C.duk_uarridx_t(arrIndex))) == 1
}

// See: http://duktape.org/api.html#duk_del_prop_string
func (d *Context) DelPropString(objIndex int, key string) bool {
	mustNotNil(d)
	__key__ := C.CString(key)
	defer C.free(unsafe.Pointer(__key__))
	return int(C.duk_del_prop_string(d.duk_context, C.duk_idx_t(objIndex), __key__)) == 1
}

// See: http://duktape.org/api.html#duk_def_prop
func (d *Context) DefProp(objIndex int, flags uint) {
	mustNotNil(d)
	C.duk_def_prop(d.duk_context, C.duk_idx_t(objIndex), C.duk_uint_t(flags))
}

// See: http://duktape.org/api.html#duk_destroy_heap
func (d *Context) DestroyHeap() {
	d.Gc(0)
	C.duk_destroy_heap(d.duk_context)
	d.duk_context = nil
}

// See: http://duktape.org/api.html#duk_dump_context_stderr
func (d *Context) DumpContextStderr() {
	mustNotNil(d)
	C._duk_dump_context_stderr(d.duk_context)
}

// See: http://duktape.org/api.html#duk_dump_context_stdout
func (d *Context) DumpContextStdout() {
	mustNotNil(d)
	C._duk_dump_context_stdout(d.duk_context)
}

// See: http://duktape.org/api.html#duk_dup
func (d *Context) Dup(fromIndex int) {
	mustNotNil(d)
	C.duk_dup(d.duk_context, C.duk_idx_t(fromIndex))
}

// See: http://duktape.org/api.html#duk_dup_top
func (d *Context) DupTop() {
	mustNotNil(d)
	C.duk_dup_top(d.duk_context)
}

// See: http://duktape.org/api.html#duk_enum
func (d *Context) Enum(objIndex int, enumFlags uint) {
	mustNotNil(d)
	C.duk_enum(d.duk_context, C.duk_idx_t(objIndex), C.duk_uint_t(enumFlags))
}

// See: http://duktape.org/api.html#duk_equals
func (d *Context) Equals(index1 int, index2 int) bool {
	mustNotNil(d)
	return int(C.duk_equals(d.duk_context, C.duk_idx_t(index1), C.duk_idx_t(index2))) == 1
}

// Error pushes a new Error object to the stack and throws it. This will call
// fmt.Sprint, forwarding arguments after the error code, to produce the
// Error's message.
//
// See: http://duktape.org/api.html#duk_error
func (d *Context) Error(errCode int, a ...interface{}) {
	mustNotNil(d)
	str := fmt.Sprint(a...)
	__str__ := C.CString(str)
	defer C.free(unsafe.Pointer(__str__))
	C._duk_error(d.duk_context, C.duk_errcode_t(errCode), __str__)
}

func (d *Context) ErrorRaw(errCode int, filename string, line int, errMsg string) {
	mustNotNil(d)
	__filename__ := C.CString(filename)
	__errMsg__ := C.CString(errMsg)
	defer C.free(unsafe.Pointer(__filename__))
	defer C.free(unsafe.Pointer(__errMsg__))

	C._duk_error_raw(d.duk_context, C.duk_errcode_t(errCode), __filename__, C.duk_int_t(line), __errMsg__)
}

// Errorf pushes a new Error object to the stack and throws it. This will call
// fmt.Sprintf, forwarding the format string and additional arguments, to
// produce the Error's message.
//
// See: http://duktape.org/api.html#duk_error
func (d *Context) Errorf(errCode int, format string, a ...interface{}) {
	mustNotNil(d)
	str := fmt.Sprintf(format, a...)
	__str__ := C.CString(str)
	defer C.free(unsafe.Pointer(__str__))
	C._duk_error(d.duk_context, C.duk_errcode_t(errCode), __str__)
}

// See: http://duktape.org/api.html#duk_eval
func (d *Context) Eval() {
	mustNotNil(d)
	C._duk_eval(d.duk_context)
}

// See: http://duktape.org/api.html#duk_eval_file
func (d *Context) EvalFile(path string) {
	mustNotNil(d)
	__path__ := C.CString(path)
	defer C.free(unsafe.Pointer(__path__))
	C._duk_eval_file(d.duk_context, __path__)
}

// See: http://duktape.org/api.html#duk_eval_file_noresult
func (d *Context) EvalFileNoresult(path string) {
	mustNotNil(d)
	__path__ := C.CString(path)
	defer C.free(unsafe.Pointer(__path__))
	C._duk_eval_file_noresult(d.duk_context, __path__)
}

// See: http://duktape.org/api.html#duk_eval_lstring
func (d *Context) EvalLstring(src string, len int) {
	mustNotNil(d)
	__src__ := C.CString(src)
	defer C.free(unsafe.Pointer(__src__))
	C._duk_eval_lstring(d.duk_context, __src__, C.duk_size_t(len))
}

// See: http://duktape.org/api.html#duk_eval_lstring_noresult
func (d *Context) EvalLstringNoresult(src string, len int) {
	mustNotNil(d)
	__src__ := C.CString(src)
	defer C.free(unsafe.Pointer(__src__))
	C._duk_eval_lstring_noresult(d.duk_context, __src__, C.duk_size_t(len))
}

// See: http://duktape.org/api.html#duk_eval_noresult
func (d *Context) EvalNoresult() {
	mustNotNil(d)
	C._duk_eval_noresult(d.duk_context)
}

// See: http://duktape.org/api.html#duk_eval_string
func (d *Context) EvalString(src string) {
	mustNotNil(d)
	__src__ := C.CString(src)
	defer C.free(unsafe.Pointer(__src__))
	C._duk_eval_string(d.duk_context, __src__)
}

// See: http://duktape.org/api.html#duk_eval_string_noresult
func (d *Context) EvalStringNoresult(src string) {
	mustNotNil(d)
	__src__ := C.CString(src)
	defer C.free(unsafe.Pointer(__src__))
	C._duk_eval_string_noresult(d.duk_context, __src__)
}

// See: http://duktape.org/api.html#duk_fatal
func (d *Context) Fatal(errCode int, errMsg string) {
	mustNotNil(d)
	__errMsg__ := C.CString(errMsg)
	defer C.free(unsafe.Pointer(__errMsg__))
	C.duk_fatal(d.duk_context, C.duk_errcode_t(errCode), __errMsg__)
}

// See: http://duktape.org/api.html#duk_gc
func (d *Context) Gc(flags uint) {
	mustNotNil(d)
	C.duk_gc(d.duk_context, C.duk_uint_t(flags))
}

// See: http://duktape.org/api.html#duk_get_boolean
func (d *Context) GetBoolean(index int) bool {
	mustNotNil(d)
	return int(C.duk_get_boolean(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_get_buffer
func (d *Context) GetBuffer(index int, outSize int) {
	mustNotNil(d)
	C.duk_get_buffer(d.duk_context, C.duk_idx_t(index), (*C.duk_size_t)(unsafe.Pointer(&outSize)))
}

// See: http://duktape.org/api.html#duk_get_context
func (d *Context) GetContext(index int) *Context {
	mustNotNil(d)
	return contextFromPointer(C.duk_get_context(d.duk_context, C.duk_idx_t(index)))
}

// See: http://duktape.org/api.html#duk_get_current_magic
func (d *Context) GetCurrentMagic() int {
	mustNotNil(d)
	return int(C.duk_get_current_magic(d.duk_context))
}

// See: http://duktape.org/api.html#duk_get_error_code
func (d *Context) GetErrorCode(index int) int {
	mustNotNil(d)
	code := int(C.duk_get_error_code(d.duk_context, C.duk_idx_t(index)))
	return code
}

// See: http://duktape.org/api.html#duk_get_finalizer
func (d *Context) GetFinalizer(index int) {
	mustNotNil(d)
	C.duk_get_finalizer(d.duk_context, C.duk_idx_t(index))
}

// See: http://duktape.org/api.html#duk_get_global_string
func (d *Context) GetGlobalString(key string) bool {
	mustNotNil(d)
	__key__ := C.CString(key)
	defer C.free(unsafe.Pointer(__key__))
	return int(C.duk_get_global_string(d.duk_context, __key__)) == 1
}

// See: http://duktape.org/api.html#duk_get_heapptr
func (d *Context) GetHeapptr(index int) unsafe.Pointer {
	mustNotNil(d)
	return unsafe.Pointer(C.duk_get_heapptr(d.duk_context, C.duk_idx_t(index)))
}

// See: http://duktape.org/api.html#duk_get_int
func (d *Context) GetInt(index int) int {
	mustNotNil(d)
	return int(C.duk_get_int(d.duk_context, C.duk_idx_t(index)))
}

// See: http://duktape.org/api.html#duk_get_length
func (d *Context) GetLength(index int) int {
	mustNotNil(d)
	return int(C.duk_get_length(d.duk_context, C.duk_idx_t(index)))
}

// See: http://duktape.org/api.html#duk_get_lstring
func (d *Context) GetLstring(index int, outLen int) string {
	mustNotNil(d)
	if s := C.duk_get_lstring(d.duk_context, C.duk_idx_t(index), (*C.duk_size_t)(unsafe.Pointer(&outLen))); s != nil {
		return C.GoString(s)
	}
	return ""
}

// See: http://duktape.org/api.html#duk_get_magic
func (d *Context) GetMagic(index int) int {
	mustNotNil(d)
	return int(C.duk_get_magic(d.duk_context, C.duk_idx_t(index)))
}

// See: http://duktape.org/api.html#duk_get_number
func (d *Context) GetNumber(index int) float64 {
	mustNotNil(d)
	return float64(C.duk_get_number(d.duk_context, C.duk_idx_t(index)))
}

// See: http://duktape.org/api.html#duk_get_pointer
func (d *Context) GetPointer(index int) unsafe.Pointer {
	mustNotNil(d)
	return C.duk_get_pointer(d.duk_context, C.duk_idx_t(index))
}

// See: http://duktape.org/api.html#duk_get_prop
func (d *Context) GetProp(objIndex int) bool {
	mustNotNil(d)
	return int(C.duk_get_prop(d.duk_context, C.duk_idx_t(objIndex))) == 1
}

// See: http://duktape.org/api.html#duk_get_prop_index
func (d *Context) GetPropIndex(objIndex int, arrIndex uint) bool {
	mustNotNil(d)
	return int(C.duk_get_prop_index(d.duk_context, C.duk_idx_t(objIndex), C.duk_uarridx_t(arrIndex))) == 1
}

// See: http://duktape.org/api.html#duk_get_prop_string
func (d *Context) GetPropString(objIndex int, key string) bool {
	mustNotNil(d)
	__key__ := C.CString(key)
	defer C.free(unsafe.Pointer(__key__))
	return int(C.duk_get_prop_string(d.duk_context, C.duk_idx_t(objIndex), __key__)) == 1
}

// See: http://duktape.org/api.html#duk_get_prototype
func (d *Context) GetPrototype(index int) {
	mustNotNil(d)
	C.duk_get_prototype(d.duk_context, C.duk_idx_t(index))
}

// See: http://duktape.org/api.html#duk_get_string
func (d *Context) GetString(i int) string {
	mustNotNil(d)
	if s := C.duk_get_string(d.duk_context, C.duk_idx_t(i)); s != nil {
		return C.GoString(s)
	}
	return ""
}

// See: http://duktape.org/api.html#duk_get_top
func (d *Context) GetTop() int {
	mustNotNil(d)
	return int(C.duk_get_top(d.duk_context))
}

// See: http://duktape.org/api.html#duk_get_top_index
func (d *Context) GetTopIndex() int {
	mustNotNil(d)
	return int(C.duk_get_top_index(d.duk_context))
}

// See: http://duktape.org/api.html#duk_get_type
func (d *Context) GetType(index int) Type {
	mustNotNil(d)
	return Type(C.duk_get_type(d.duk_context, C.duk_idx_t(index)))
}

// See: http://duktape.org/api.html#duk_get_type_mask
func (d *Context) GetTypeMask(index int) uint {
	mustNotNil(d)
	return uint(C.duk_get_type_mask(d.duk_context, C.duk_idx_t(index)))
}

// See: http://duktape.org/api.html#duk_get_uint
func (d *Context) GetUint(index int) uint {
	mustNotNil(d)
	return uint(C.duk_get_uint(d.duk_context, C.duk_idx_t(index)))
}

// See: http://duktape.org/api.html#duk_has_prop
func (d *Context) HasProp(objIndex int) bool {
	mustNotNil(d)
	return int(C.duk_has_prop(d.duk_context, C.duk_idx_t(objIndex))) == 1
}

// See: http://duktape.org/api.html#duk_has_prop_index
func (d *Context) HasPropIndex(objIndex int, arrIndex uint) bool {
	mustNotNil(d)
	return int(C.duk_has_prop_index(d.duk_context, C.duk_idx_t(objIndex), C.duk_uarridx_t(arrIndex))) == 1
}

// See: http://duktape.org/api.html#duk_has_prop_string
func (d *Context) HasPropString(objIndex int, key string) bool {
	mustNotNil(d)
	__key__ := C.CString(key)
	defer C.free(unsafe.Pointer(__key__))
	return int(C.duk_has_prop_string(d.duk_context, C.duk_idx_t(objIndex), __key__)) == 1
}

// See: http://duktape.org/api.html#duk_hex_decode
func (d *Context) HexDecode(index int) {
	mustNotNil(d)
	C.duk_hex_decode(d.duk_context, C.duk_idx_t(index))
}

// See: http://duktape.org/api.html#duk_hex_encode
func (d *Context) HexEncode(index int) string {
	mustNotNil(d)
	if s := C.duk_hex_encode(d.duk_context, C.duk_idx_t(index)); s != nil {
		return C.GoString(s)
	}
	return ""
}

// See: http://duktape.org/api.html#duk_insert
func (d *Context) Insert(toIndex int) {
	mustNotNil(d)
	C.duk_insert(d.duk_context, C.duk_idx_t(toIndex))
}

// See: http://duktape.org/api.html#duk_is_array
func (d *Context) IsArray(index int) bool {
	mustNotNil(d)
	return int(C.duk_is_array(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_is_boolean
func (d *Context) IsBoolean(index int) bool {
	mustNotNil(d)
	return int(C.duk_is_boolean(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_is_bound_function
func (d *Context) IsBoundFunction(index int) bool {
	mustNotNil(d)
	return int(C.duk_is_bound_function(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_is_buffer
func (d *Context) IsBuffer(index int) bool {
	mustNotNil(d)
	return int(C.duk_is_buffer(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_is_c_function
func (d *Context) IsCFunction(index int) bool {
	mustNotNil(d)
	return int(C.duk_is_c_function(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_is_callable
func (d *Context) IsCallable(index int) bool {
	mustNotNil(d)
	return int(C.duk_is_callable(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_is_constructor_call
func (d *Context) IsConstructorCall() bool {
	mustNotNil(d)
	return int(C.duk_is_constructor_call(d.duk_context)) == 1
}

// See: http://duktape.org/api.html#duk_is_dynamic_buffer
func (d *Context) IsDynamicBuffer(index int) bool {
	mustNotNil(d)
	return int(C.duk_is_dynamic_buffer(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_is_ecmascript_function
func (d *Context) IsEcmascriptFunction(index int) bool {
	mustNotNil(d)
	return int(C.duk_is_ecmascript_function(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_is_fixed_buffer
func (d *Context) IsFixedBuffer(index int) bool {
	mustNotNil(d)
	return int(C.duk_is_fixed_buffer(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_is_function
func (d *Context) IsFunction(index int) bool {
	mustNotNil(d)
	return int(C.duk_is_function(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_is_nan
func (d *Context) IsNan(index int) bool {
	mustNotNil(d)
	return int(C.duk_is_nan(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_is_null
func (d *Context) IsNull(index int) bool {
	mustNotNil(d)
	return int(C.duk_is_null(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_is_null_or_undefined
func (d *Context) IsNullOrUndefined(index int) bool {
	mustNotNil(d)
	return int(C.duk_is_null_or_undefined(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_is_number
func (d *Context) IsNumber(index int) bool {
	mustNotNil(d)
	return int(C.duk_is_number(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_is_object
func (d *Context) IsObject(index int) bool {
	mustNotNil(d)
	return int(C.duk_is_object(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_is_error
func (d *Context) IsError(index int) bool {
	mustNotNil(d)
	return int(C._duk_is_error(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_is_object_coercible
func (d *Context) IsObjectCoercible(index int) bool {
	mustNotNil(d)
	return int(C._duk_is_object_coercible(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_is_pointer
func (d *Context) IsPointer(index int) bool {
	mustNotNil(d)
	return int(C.duk_is_pointer(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_is_primitive
func (d *Context) IsPrimitive(index int) bool {
	mustNotNil(d)
	return int(C.duk_is_primitive(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_is_strict_call
func (d *Context) IsStrictCall() bool {
	mustNotNil(d)
	return int(C.duk_is_strict_call(d.duk_context)) == 1
}

// See: http://duktape.org/api.html#duk_is_string
func (d *Context) IsString(index int) bool {
	mustNotNil(d)
	return int(C.duk_is_string(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_is_thread
func (d *Context) IsThread(index int) bool {
	mustNotNil(d)
	return int(C.duk_is_thread(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_is_undefined
func (d *Context) IsUndefined(index int) bool {
	mustNotNil(d)
	return int(C.duk_is_undefined(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_is_valid_index
func (d *Context) IsValidIndex(index int) bool {
	mustNotNil(d)
	return int(C.duk_is_valid_index(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_join
func (d *Context) Join(count int) {
	mustNotNil(d)
	C.duk_join(d.duk_context, C.duk_idx_t(count))
}

// See: http://duktape.org/api.html#duk_json_decode
func (d *Context) JsonDecode(index int) {
	mustNotNil(d)
	C.duk_json_decode(d.duk_context, C.duk_idx_t(index))
}

// See: http://duktape.org/api.html#duk_json_encode
func (d *Context) JsonEncode(index int) string {
	mustNotNil(d)
	if s := C.duk_json_encode(d.duk_context, C.duk_idx_t(index)); s != nil {
		return C.GoString(s)
	}
	return ""
}

// See: http://duktape.org/api.html#duk_new
func (d *Context) New(nargs int) {
	C.duk_new(d.duk_context, C.duk_idx_t(nargs))
}

// See: http://duktape.org/api.html#duk_next
func (d *Context) Next(enumIndex int, getValue bool) bool {
	mustNotNil(d)
	var __getValue__ int
	if getValue {
		__getValue__ = 1
	}
	return int(C.duk_next(d.duk_context, C.duk_idx_t(enumIndex), C.duk_bool_t(__getValue__))) == 1
}

// See: http://duktape.org/api.html#duk_normalize_index
func (d *Context) NormalizeIndex(index int) int {
	mustNotNil(d)
	return int(C.duk_normalize_index(d.duk_context, C.duk_idx_t(index)))
}

// See: http://duktape.org/api.html#duk_pcall
func (d *Context) Pcall(nargs int) int {
	mustNotNil(d)
	return int(C.duk_pcall(d.duk_context, C.duk_idx_t(nargs)))
}

// See: http://duktape.org/api.html#duk_pcall_method
func (d *Context) PcallMethod(nargs int) int {
	mustNotNil(d)
	return int(C.duk_pcall_method(d.duk_context, C.duk_idx_t(nargs)))
}

// See: http://duktape.org/api.html#duk_pcall_prop
func (d *Context) PcallProp(objIndex int, nargs int) int {
	mustNotNil(d)
	return int(C.duk_pcall_prop(d.duk_context, C.duk_idx_t(objIndex), C.duk_idx_t(nargs)))
}

// See: http://duktape.org/api.html#duk_pcompile
func (d *Context) Pcompile(flags uint) error {
	mustNotNil(d)
	result := int(C._duk_pcompile(d.duk_context, C.duk_uint_t(flags)))
	return d.castStringToError(result)
}

// See: http://duktape.org/api.html#duk_pcompile_file
func (d *Context) PcompileFile(flags uint, path string) error {
	mustNotNil(d)
	__path__ := C.CString(path)
	defer C.free(unsafe.Pointer(__path__))

	result := int(C._duk_pcompile_file(d.duk_context, C.duk_uint_t(flags), __path__))
	return d.castStringToError(result)
}

// See: http://duktape.org/api.html#duk_pcompile_lstring
func (d *Context) PcompileLstring(flags uint, src string, len int) error {
	mustNotNil(d)
	__src__ := C.CString(src)
	defer C.free(unsafe.Pointer(__src__))

	result := int(C._duk_pcompile_lstring(d.duk_context, C.duk_uint_t(flags), __src__, C.duk_size_t(len)))
	return d.castStringToError(result)
}

// See: http://duktape.org/api.html#duk_pcompile_lstring_filename
func (d *Context) PcompileLstringFilename(flags uint, src string, len int) error {
	mustNotNil(d)
	__src__ := C.CString(src)
	defer C.free(unsafe.Pointer(__src__))

	result := int(C._duk_pcompile_lstring_filename(d.duk_context, C.duk_uint_t(flags), __src__, C.duk_size_t(len)))
	return d.castStringToError(result)
}

// See: http://duktape.org/api.html#duk_pcompile_string
func (d *Context) PcompileString(flags uint, src string) error {
	mustNotNil(d)
	__src__ := C.CString(src)
	defer C.free(unsafe.Pointer(__src__))

	result := int(C._duk_pcompile_string(d.duk_context, C.duk_uint_t(flags), __src__))
	return d.castStringToError(result)
}

// See: http://duktape.org/api.html#duk_pcompile_string_filename
func (d *Context) PcompileStringFilename(flags uint, src string) error {
	mustNotNil(d)
	__src__ := C.CString(src)
	defer C.free(unsafe.Pointer(__src__))

	result := int(C._duk_pcompile_string_filename(d.duk_context, C.duk_uint_t(flags), __src__))
	return d.castStringToError(result)
}

// See: http://duktape.org/api.html#duk_peval
func (d *Context) Peval() error {
	mustNotNil(d)
	result := int(C._duk_peval(d.duk_context))
	return d.castStringToError(result)
}

// See: http://duktape.org/api.html#duk_peval_file
func (d *Context) PevalFile(path string) error {
	mustNotNil(d)
	__path__ := C.CString(path)
	defer C.free(unsafe.Pointer(__path__))

	result := int(C._duk_peval_file(d.duk_context, __path__))
	return d.castStringToError(result)
}

// See: http://duktape.org/api.html#duk_peval_file_noresult
func (d *Context) PevalFileNoresult(path string) int {
	mustNotNil(d)
	__path__ := C.CString(path)
	defer C.free(unsafe.Pointer(__path__))
	return int(C._duk_peval_file_noresult(d.duk_context, __path__))
}

// See: http://duktape.org/api.html#duk_peval_lstring
func (d *Context) PevalLstring(src string, len int) error {
	mustNotNil(d)
	__src__ := C.CString(src)
	defer C.free(unsafe.Pointer(__src__))

	result := int(C._duk_peval_lstring(d.duk_context, __src__, C.duk_size_t(len)))
	return d.castStringToError(result)

}

// See: http://duktape.org/api.html#duk_peval_lstring_noresult
func (d *Context) PevalLstringNoresult(src string, len int) int {
	mustNotNil(d)
	__src__ := C.CString(src)
	defer C.free(unsafe.Pointer(__src__))
	return int(C._duk_peval_lstring_noresult(d.duk_context, __src__, C.duk_size_t(len)))
}

// See: http://duktape.org/api.html#duk_peval_noresult
func (d *Context) PevalNoresult() int {
	mustNotNil(d)
	return int(C._duk_peval_noresult(d.duk_context))
}

// See: http://duktape.org/api.html#duk_peval_string
func (d *Context) PevalString(src string) error {
	mustNotNil(d)
	__src__ := C.CString(src)
	defer C.free(unsafe.Pointer(__src__))

	result := int(C._duk_peval_string(d.duk_context, __src__))
	return d.castStringToError(result)
}

// See: http://duktape.org/api.html#duk_peval_string_noresult
func (d *Context) PevalStringNoresult(src string) int {
	mustNotNil(d)
	__src__ := C.CString(src)
	defer C.free(unsafe.Pointer(__src__))

	return int(C._duk_peval_string_noresult(d.duk_context, __src__))
}

func (d *Context) castStringToError(result int) error {
	mustNotNil(d)
	if result == 0 {
		return nil
	}

	err := &Error{}
	for _, key := range []string{"name", "message", "fileName", "lineNumber", "stack"} {
		d.GetPropString(-1, key)

		switch key {
		case "name":
			err.Type = d.SafeToString(-1)
		case "message":
			err.Message = d.SafeToString(-1)
		case "fileName":
			err.FileName = d.SafeToString(-1)
		case "lineNumber":
			if d.IsNumber(-1) {
				err.LineNumber = d.GetInt(-1)
			}
		case "stack":
			err.Stack = d.SafeToString(-1)
		}

		d.Pop()
	}

	return err
}

// See: http://duktape.org/api.html#duk_pop
func (d *Context) Pop() {
	mustNotNil(d)
	if d.GetTop() == 0 {
		return
	}
	C.duk_pop(d.duk_context)
}

// See: http://duktape.org/api.html#duk_pop_2
func (d *Context) Pop2() {
	mustNotNil(d)
	d.PopN(2)
}

// See: http://duktape.org/api.html#duk_pop_3
func (d *Context) Pop3() {
	mustNotNil(d)
	d.PopN(3)
}

// See: http://duktape.org/api.html#duk_pop_n
func (d *Context) PopN(count int) {
	mustNotNil(d)
	if d.GetTop() < count || count < 1 {
		return
	}
	C.duk_pop_n(d.duk_context, C.duk_idx_t(count))
}

// See: http://duktape.org/api.html#duk_push_array
func (d *Context) PushArray() int {
	mustNotNil(d)
	return int(C.duk_push_array(d.duk_context))
}

// See: http://duktape.org/api.html#duk_push_boolean
func (d *Context) PushBoolean(val bool) {
	mustNotNil(d)
	var __val__ int
	if val {
		__val__ = 1
	}
	C.duk_push_boolean(d.duk_context, C.duk_bool_t(__val__))
}

// See: http://duktape.org/api.html#duk_push_buffer
func (d *Context) PushBuffer(size int, dynamic bool) {
	mustNotNil(d)
	var __dynamic__ int
	if dynamic {
		__dynamic__ = 1
	}
	C._duk_push_buffer(d.duk_context, C.duk_size_t(size), C.duk_bool_t(__dynamic__))
}

// See: http://duktape.org/api.html#duk_push_c_function
func (d *Context) PushCFunction(fn *[0]byte, nargs int) int {
	mustNotNil(d)
	return int(C.duk_push_c_function(d.duk_context, fn, C.duk_idx_t(nargs)))
}

// See: http://duktape.org/api.html#duk_push_context_dump
func (d *Context) PushContextDump() {
	mustNotNil(d)
	C.duk_push_context_dump(d.duk_context)
}

// See: http://duktape.org/api.html#duk_push_current_function
func (d *Context) PushCurrentFunction() {
	mustNotNil(d)
	C.duk_push_current_function(d.duk_context)
}

// See: http://duktape.org/api.html#duk_push_current_thread
func (d *Context) PushCurrentThread() {
	mustNotNil(d)
	C.duk_push_current_thread(d.duk_context)
}

// See: http://duktape.org/api.html#duk_push_dynamic_buffer
func (d *Context) PushDynamicBuffer(size int) {
	mustNotNil(d)
	C._duk_push_dynamic_buffer(d.duk_context, C.duk_size_t(size))
}

// PushErrorObject pushes a new Error object to the stack. This will call
// fmt.Sprint, forwarding arguments after the error code, to produce the
// Error's message.
//
// See: http://duktape.org/api.html#duk_push_error_object
func (d *Context) PushErrorObject(errCode int, a ...interface{}) {
	mustNotNil(d)
	str := fmt.Sprint(a...)
	__str__ := C.CString(str)
	defer C.free(unsafe.Pointer(__str__))
	C._duk_push_error_object(d.duk_context, C.duk_errcode_t(errCode), __str__)
}

// PushErrorObjectf pushes a new Error object to the stack. This will call
// fmt.Sprintf, forwarding the format string and additional arguments, to
// produce the Error's message.
//
// See: http://duktape.org/api.html#duk_push_error_object
func (d *Context) PushErrorObjectf(errCode int, format string, a ...interface{}) {
	mustNotNil(d)
	str := fmt.Sprintf(format, a...)
	__str__ := C.CString(str)
	defer C.free(unsafe.Pointer(__str__))
	C._duk_push_error_object(d.duk_context, C.duk_errcode_t(errCode), __str__)
}

// See: http://duktape.org/api.html#duk_push_false
func (d *Context) PushFalse() {
	mustNotNil(d)
	C.duk_push_false(d.duk_context)
}

// See: http://duktape.org/api.html#duk_push_fixed_buffer
func (d *Context) PushFixedBuffer(size int) {
	mustNotNil(d)
	C._duk_push_fixed_buffer(d.duk_context, C.duk_size_t(size))
}

// See: http://duktape.org/api.html#duk_push_global_object
func (d *Context) PushGlobalObject() {
	mustNotNil(d)
	C.duk_push_global_object(d.duk_context)
}

// See: http://duktape.org/api.html#duk_push_global_stash
func (d *Context) PushGlobalStash() {
	mustNotNil(d)
	C.duk_push_global_stash(d.duk_context)
}

// See: http://duktape.org/api.html#duk_push_heapptr
func (d *Context) PushHeapptr(ptr unsafe.Pointer) {
	mustNotNil(d)
	C.duk_push_heapptr(d.duk_context, ptr)
}

// See: http://duktape.org/api.html#duk_push_heap_stash
func (d *Context) PushHeapStash() {
	mustNotNil(d)
	C.duk_push_heap_stash(d.duk_context)
}

// See: http://duktape.org/api.html#duk_push_int
func (d *Context) PushInt(val int) {
	mustNotNil(d)
	C.duk_push_int(d.duk_context, C.duk_int_t(val))
}

// See: http://duktape.org/api.html#duk_push_lstring
func (d *Context) PushLstring(str string, len int) string {
	mustNotNil(d)
	__str__ := C.CString(str)
	defer C.free(unsafe.Pointer(__str__))
	if s := C.duk_push_lstring(d.duk_context, __str__, C.duk_size_t(len)); s != nil {
		return C.GoString(s)
	}
	return ""
}

// See: http://duktape.org/api.html#duk_push_nan
func (d *Context) PushNan() {
	mustNotNil(d)
	C.duk_push_nan(d.duk_context)
}

// See: http://duktape.org/api.html#duk_push_null
func (d *Context) PushNull() {
	mustNotNil(d)
	C.duk_push_null(d.duk_context)
}

// See: http://duktape.org/api.html#duk_push_number
func (d *Context) PushNumber(val float64) {
	mustNotNil(d)
	C.duk_push_number(d.duk_context, C.duk_double_t(val))
}

// See: http://duktape.org/api.html#duk_push_object
func (d *Context) PushObject() int {
	mustNotNil(d)
	return int(C.duk_push_object(d.duk_context))
}

// See: http://duktape.org/api.html#duk_push_string
func (d *Context) PushString(str string) string {
	mustNotNil(d)
	__str__ := C.CString(str)
	defer C.free(unsafe.Pointer(__str__))
	if s := C.duk_push_string(d.duk_context, __str__); s != nil {
		return C.GoString(s)
	}
	return ""
}

// See: http://duktape.org/api.html#duk_push_string_file
func (d *Context) PushStringFile(path string) string {
	mustNotNil(d)
	__path__ := C.CString(path)
	defer C.free(unsafe.Pointer(__path__))
	if s := C._duk_push_string_file(d.duk_context, __path__); s != nil {
		return C.GoString(s)
	}
	return ""
}

// See: http://duktape.org/api.html#duk_push_this
func (d *Context) PushThis() {
	mustNotNil(d)
	C.duk_push_this(d.duk_context)
}

// See: http://duktape.org/api.html#duk_push_thread
func (d *Context) PushThread() int {
	mustNotNil(d)
	return int(C._duk_push_thread(d.duk_context))
}

// See: http://duktape.org/api.html#duk_push_thread_new_globalenv
func (d *Context) PushThreadNewGlobalenv() int {
	mustNotNil(d)
	return int(C._duk_push_thread_new_globalenv(d.duk_context))
}

// See: http://duktape.org/api.html#duk_push_thread_stash
func (d *Context) PushThreadStash(targetCtx *Context) {
	mustNotNil(d)
	C.duk_push_thread_stash(d.duk_context, targetCtx.duk_context)
}

// See: http://duktape.org/api.html#duk_push_true
func (d *Context) PushTrue() {
	mustNotNil(d)
	C.duk_push_true(d.duk_context)
}

// See: http://duktape.org/api.html#duk_push_uint
func (d *Context) PushUint(val uint) {
	mustNotNil(d)
	C.duk_push_uint(d.duk_context, C.duk_uint_t(val))
}

// See: http://duktape.org/api.html#duk_push_undefined
func (d *Context) PushUndefined() {
	mustNotNil(d)
	C.duk_push_undefined(d.duk_context)
}

// See: http://duktape.org/api.html#duk_put_global_string
func (d *Context) PutGlobalString(key string) bool {
	mustNotNil(d)
	__key__ := C.CString(key)
	defer C.free(unsafe.Pointer(__key__))
	return int(C.duk_put_global_string(d.duk_context, __key__)) == 1
}

// See: http://duktape.org/api.html#duk_put_prop
func (d *Context) PutProp(objIndex int) bool {
	mustNotNil(d)
	return int(C.duk_put_prop(d.duk_context, C.duk_idx_t(objIndex))) == 1
}

// See: http://duktape.org/api.html#duk_put_prop_index
func (d *Context) PutPropIndex(objIndex int, arrIndex uint) bool {
	mustNotNil(d)
	return int(C.duk_put_prop_index(d.duk_context, C.duk_idx_t(objIndex), C.duk_uarridx_t(arrIndex))) == 1
}

// See: http://duktape.org/api.html#duk_put_prop_string
func (d *Context) PutPropString(objIndex int, key string) bool {
	mustNotNil(d)
	__key__ := C.CString(key)
	defer C.free(unsafe.Pointer(__key__))
	return int(C.duk_put_prop_string(d.duk_context, C.duk_idx_t(objIndex), __key__)) == 1
}

// See: http://duktape.org/api.html#duk_remove
func (d *Context) Remove(index int) {
	mustNotNil(d)
	C.duk_remove(d.duk_context, C.duk_idx_t(index))
}

// See: http://duktape.org/api.html#duk_replace
func (d *Context) Replace(toIndex int) {
	mustNotNil(d)
	C.duk_replace(d.duk_context, C.duk_idx_t(toIndex))
}

// See: http://duktape.org/api.html#duk_require_boolean
func (d *Context) RequireBoolean(index int) bool {
	mustNotNil(d)
	return int(C.duk_require_boolean(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_require_buffer
func (d *Context) RequireBuffer(index int, outSize int) {
	mustNotNil(d)
	C.duk_require_buffer(d.duk_context, C.duk_idx_t(index), (*C.duk_size_t)(unsafe.Pointer(&outSize)))
}

// See: http://duktape.org/api.html#duk_require_context
func (d *Context) RequireContext(index int) *Context {
	mustNotNil(d)
	return contextFromPointer(C.duk_require_context(d.duk_context, C.duk_idx_t(index)))
}

// See: http://duktape.org/api.html#duk_require_heapptr
func (d *Context) RequireHeapptr(index int) unsafe.Pointer {
	mustNotNil(d)
	return unsafe.Pointer(C.duk_require_heapptr(d.duk_context, C.duk_idx_t(index)))
}

// See: http://duktape.org/api.html#duk_require_int
func (d *Context) RequireInt(index int) int {
	mustNotNil(d)
	return int(C.duk_require_int(d.duk_context, C.duk_idx_t(index)))
}

// See: http://duktape.org/api.html#duk_require_lstring
func (d *Context) RequireLstring(index int, outLen int) string {
	mustNotNil(d)
	if s := C.duk_require_lstring(d.duk_context, C.duk_idx_t(index), (*C.duk_size_t)(unsafe.Pointer(&outLen))); s != nil {
		return C.GoString(s)
	}
	return ""
}

// See: http://duktape.org/api.html#duk_require_normalize_index
func (d *Context) RequireNormalizeIndex(index int) int {
	mustNotNil(d)
	return int(C.duk_require_normalize_index(d.duk_context, C.duk_idx_t(index)))
}

// See: http://duktape.org/api.html#duk_require_null
func (d *Context) RequireNull(index int) {
	mustNotNil(d)
	C.duk_require_null(d.duk_context, C.duk_idx_t(index))
}

// See: http://duktape.org/api.html#duk_require_number
func (d *Context) RequireNumber(index int) float64 {
	mustNotNil(d)
	return float64(C.duk_require_number(d.duk_context, C.duk_idx_t(index)))
}

// See: http://duktape.org/api.html#duk_require_object_coercible
func (d *Context) RequireObjectCoercible(index int) {
	mustNotNil(d)
	C._duk_require_object_coercible(d.duk_context, C.duk_idx_t(index))
}

// See: http://duktape.org/api.html#duk_require_pointer
func (d *Context) RequirePointer(index int) {
	mustNotNil(d)
	C.duk_require_pointer(d.duk_context, C.duk_idx_t(index))
}

// See: http://duktape.org/api.html#duk_require_stack
func (d *Context) RequireStack(extra int) {
	mustNotNil(d)
	C.duk_require_stack(d.duk_context, C.duk_idx_t(extra))
}

// See: http://duktape.org/api.html#duk_require_stack_top
func (d *Context) RequireStackTop(top int) {
	mustNotNil(d)
	C.duk_require_stack_top(d.duk_context, C.duk_idx_t(top))
}

// See: http://duktape.org/api.html#duk_require_string
func (d *Context) RequireString(index int) string {
	mustNotNil(d)
	if s := C.duk_require_string(d.duk_context, C.duk_idx_t(index)); s != nil {
		return C.GoString(s)
	}
	return ""
}

// See: http://duktape.org/api.html#duk_require_top_index
func (d *Context) RequireTopIndex() int {
	mustNotNil(d)
	return int(C.duk_require_top_index(d.duk_context))
}

// See: http://duktape.org/api.html#duk_require_type_mask
func (d *Context) RequireTypeMask(index int, mask uint) {
	mustNotNil(d)
	C._duk_require_type_mask(d.duk_context, C.duk_idx_t(index), C.duk_uint_t(mask))
}

// See: http://duktape.org/api.html#duk_require_uint
func (d *Context) RequireUint(index int) uint {
	mustNotNil(d)
	return uint(C.duk_require_uint(d.duk_context, C.duk_idx_t(index)))
}

// See: http://duktape.org/api.html#duk_require_undefined
func (d *Context) RequireUndefined(index int) {
	mustNotNil(d)
	C.duk_require_undefined(d.duk_context, C.duk_idx_t(index))
}

// See: http://duktape.org/api.html#duk_require_valid_index
func (d *Context) RequireValidIndex(index int) {
	mustNotNil(d)
	C.duk_require_valid_index(d.duk_context, C.duk_idx_t(index))
}

// See: http://duktape.org/api.html#duk_resize_buffer
func (d *Context) ResizeBuffer(index int, newSize int) {
	mustNotNil(d)
	C.duk_resize_buffer(d.duk_context, C.duk_idx_t(index), C.duk_size_t(newSize))
}

// See: http://duktape.org/api.html#duk_safe_call
func (d *Context) SafeCall(fn *[0]byte, nargs, nrets int) int {
	mustNotNil(d)
	return int(C.duk_safe_call(
		d.duk_context,
		fn,
		C.duk_idx_t(nargs),
		C.duk_idx_t(nrets),
	))
}

// See: http://duktape.org/api.html#duk_safe_to_lstring
func (d *Context) SafeToLstring(index int, outLen int) string {
	mustNotNil(d)
	if s := C.duk_safe_to_lstring(d.duk_context, C.duk_idx_t(index), (*C.duk_size_t)(unsafe.Pointer(&outLen))); s != nil {
		return C.GoString(s)
	}
	return ""
}

// See: http://duktape.org/api.html#duk_safe_to_string
func (d *Context) SafeToString(index int) string {
	mustNotNil(d)
	if s := C._duk_safe_to_string(d.duk_context, C.duk_idx_t(index)); s != nil {
		return C.GoString(s)
	}
	return ""
}

// See: http://duktape.org/api.html#duk_set_finalizer
func (d *Context) SetFinalizer(index int) {
	mustNotNil(d)
	C.duk_set_finalizer(d.duk_context, C.duk_idx_t(index))
}

// See: http://duktape.org/api.html#duk_set_global_object
func (d *Context) SetGlobalObject() {
	mustNotNil(d)
	C.duk_set_global_object(d.duk_context)
}

// See: http://duktape.org/api.html#duk_set_magic
func (d *Context) SetMagic(index int, magic int) {
	mustNotNil(d)
	C.duk_set_magic(d.duk_context, C.duk_idx_t(index), C.duk_int_t(magic))
}

// See: http://duktape.org/api.html#duk_set_prototype
func (d *Context) SetPrototype(index int) {
	mustNotNil(d)
	C.duk_set_prototype(d.duk_context, C.duk_idx_t(index))
}

// See: http://duktape.org/api.html#duk_set_top
func (d *Context) SetTop(index int) {
	mustNotNil(d)
	C.duk_set_top(d.duk_context, C.duk_idx_t(index))
}

func (d *Context) StrictEquals(index1 int, index2 int) bool {
	mustNotNil(d)
	return int(C.duk_strict_equals(d.duk_context, C.duk_idx_t(index1), C.duk_idx_t(index2))) == 1
}

// See: http://duktape.org/api.html#duk_substring
func (d *Context) Substring(index int, startCharOffset int, endCharOffset int) {
	mustNotNil(d)
	C.duk_substring(d.duk_context, C.duk_idx_t(index), C.duk_size_t(startCharOffset), C.duk_size_t(endCharOffset))
}

// See: http://duktape.org/api.html#duk_swap
func (d *Context) Swap(index1 int, index2 int) {
	mustNotNil(d)
	C.duk_swap(d.duk_context, C.duk_idx_t(index1), C.duk_idx_t(index2))
}

// See: http://duktape.org/api.html#duk_swap_top
func (d *Context) SwapTop(index int) {
	mustNotNil(d)
	C.duk_swap_top(d.duk_context, C.duk_idx_t(index))
}

// See: http://duktape.org/api.html#duk_throw
func (d *Context) Throw() {
	mustNotNil(d)
	C.duk_throw(d.duk_context)
}

// See: http://duktape.org/api.html#duk_to_boolean
func (d *Context) ToBoolean(index int) bool {
	mustNotNil(d)
	return int(C.duk_to_boolean(d.duk_context, C.duk_idx_t(index))) == 1
}

// See: http://duktape.org/api.html#duk_to_buffer
func (d *Context) ToBuffer(index int, outSize int) {
	mustNotNil(d)
	C._duk_to_buffer(d.duk_context, C.duk_idx_t(index), (*C.duk_size_t)(unsafe.Pointer(&outSize)))
}

// See: http://duktape.org/api.html#duk_to_defaultvalue
func (d *Context) ToDefaultvalue(index int, hint int) {
	mustNotNil(d)
	C.duk_to_defaultvalue(d.duk_context, C.duk_idx_t(index), C.duk_int_t(hint))
}

// See: http://duktape.org/api.html#duk_to_dynamic_buffer
func (d *Context) ToDynamicBuffer(index int, outSize int) {
	mustNotNil(d)
	C._duk_to_dynamic_buffer(d.duk_context, C.duk_idx_t(index), (*C.duk_size_t)(unsafe.Pointer(&outSize)))
}

// See: http://duktape.org/api.html#duk_to_fixed_buffer
func (d *Context) ToFixedBuffer(index int, outSize int) {
	mustNotNil(d)
	C._duk_to_fixed_buffer(d.duk_context, C.duk_idx_t(index), (*C.duk_size_t)(unsafe.Pointer(&outSize)))
}

// See: http://duktape.org/api.html#duk_to_int
func (d *Context) ToInt(index int) int {
	mustNotNil(d)
	return int(C.duk_to_int(d.duk_context, C.duk_idx_t(index)))
}

// See: http://duktape.org/api.html#duk_to_int32
func (d *Context) ToInt32(index int) int32 {
	mustNotNil(d)
	return int32(C.duk_to_int32(d.duk_context, C.duk_idx_t(index)))
}

// See: http://duktape.org/api.html#duk_to_lstring
func (d *Context) ToLstring(index int, outLen int) string {
	mustNotNil(d)
	if s := C.duk_to_lstring(d.duk_context, C.duk_idx_t(index), (*C.duk_size_t)(unsafe.Pointer(&outLen))); s != nil {
		return C.GoString(s)
	}
	return ""
}

// See: http://duktape.org/api.html#duk_to_null
func (d *Context) ToNull(index int) {
	mustNotNil(d)
	C.duk_to_null(d.duk_context, C.duk_idx_t(index))
}

// See: http://duktape.org/api.html#duk_to_number
func (d *Context) ToNumber(index int) float64 {
	mustNotNil(d)
	return float64(C.duk_to_number(d.duk_context, C.duk_idx_t(index)))
}

// See: http://duktape.org/api.html#duk_to_object
func (d *Context) ToObject(index int) {
	mustNotNil(d)
	C.duk_to_object(d.duk_context, C.duk_idx_t(index))
}

// See: http://duktape.org/api.html#duk_to_pointer
func (d *Context) ToPointer(index int) {
	mustNotNil(d)
	C.duk_to_pointer(d.duk_context, C.duk_idx_t(index))
}

// See: http://duktape.org/api.html#duk_to_primitive
func (d *Context) ToPrimitive(index int, hint int) {
	mustNotNil(d)
	C.duk_to_primitive(d.duk_context, C.duk_idx_t(index), C.duk_int_t(hint))
}

// See: http://duktape.org/api.html#duk_to_string
func (d *Context) ToString(index int) string {
	mustNotNil(d)
	if s := C.duk_to_string(d.duk_context, C.duk_idx_t(index)); s != nil {
		return C.GoString(s)
	}
	return ""
}

// See: http://duktape.org/api.html#duk_to_uint
func (d *Context) ToUint(index int) uint {
	mustNotNil(d)
	return uint(C.duk_to_uint(d.duk_context, C.duk_idx_t(index)))
}

// See: http://duktape.org/api.html#duk_to_uint16
func (d *Context) ToUint16(index int) uint16 {
	mustNotNil(d)
	return uint16(C.duk_to_uint16(d.duk_context, C.duk_idx_t(index)))
}

// See: http://duktape.org/api.html#duk_to_uint32
func (d *Context) ToUint32(index int) uint32 {
	mustNotNil(d)
	return uint32(C.duk_to_uint32(d.duk_context, C.duk_idx_t(index)))
}

// See: http://duktape.org/api.html#duk_to_undefined
func (d *Context) ToUndefined(index int) {
	mustNotNil(d)
	C.duk_to_undefined(d.duk_context, C.duk_idx_t(index))
}

// See: http://duktape.org/api.html#duk_trim
func (d *Context) Trim(index int) {
	mustNotNil(d)
	C.duk_trim(d.duk_context, C.duk_idx_t(index))
}

// See: http://duktape.org/api.html#duk_xcopy_top
func (d *Context) XcopyTop(fromCtx *Context, count int) {
	mustNotNil(d)
	C._duk_xcopy_top(d.duk_context, fromCtx.duk_context, C.duk_idx_t(count))
}

// See: http://duktape.org/api.html#duk_xmove_top
func (d *Context) XmoveTop(fromCtx *Context, count int) {
	mustNotNil(d)
	C._duk_xmove_top(d.duk_context, fromCtx.duk_context, C.duk_idx_t(count))
}

// See: http://duktape.org/api.html#duk_push_pointer
func (d *Context) PushPointer(p unsafe.Pointer) {
	mustNotNil(d)
	C.duk_push_pointer(d.duk_context, p)
}

/**
 * Unimplemented.
 *
 * CharCodeAt see: http://duktape.org/api.html#duk_char_code_at
 * CreateHeap see: http://duktape.org/api.html#duk_create_heap
 * DecodeString see: http://duktape.org/api.html#duk_decode_string
 * Free see: http://duktape.org/api.html#duk_free
 * FreeRaw see: http://duktape.org/api.html#duk_free_raw
 * GetMemoryFunctions see: http://duktape.org/api.html#duk_get_memory_functions
 * MapString see: http://duktape.org/api.html#duk_map_string
 * PushSprintf see: http://duktape.org/api.html#duk_push_sprintf
 * PushVsprintf see: http://duktape.org/api.html#duk_push_vsprintf
 * PutFunctionList see: http://duktape.org/api.html#duk_put_function_list
 * PutNumberList see: http://duktape.org/api.html#duk_put_number_list
 * Realloc see: http://duktape.org/api.html#duk_realloc
 * ReallocRaw see: http://duktape.org/api.html#duk_realloc_raw
 * GetCFunction see: http://duktape.org/api.html#duk_get_c_function
 * RequireCFunction see: http://duktape.org/api.html#duk_require_c_function
 */
