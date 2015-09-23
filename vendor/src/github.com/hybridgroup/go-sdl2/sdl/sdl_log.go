package sdl

// #include <SDL2/SDL_log.h>
import "C"
// import "unsafe"

const (
	LOG_CATEGORY_APPLICATION	= iota
	LOG_CATEGORY_ERROR
	LOG_CATEGORY_ASSERT
	LOG_CATEGORY_SYSTEM
	LOG_CATEGORY_AUDIO
	LOG_CATEGORY_VIDEO
	LOG_CATEGORY_RENDER
	LOG_CATEGORY_INPUT
	LOG_CATEGORY_TEST
	LOG_CATEGORY_RESERVED1
	LOG_CATEGORY_RESERVED2
	LOG_CATEGORY_RESERVED3
	LOG_CATEGORY_RESERVED4
	LOG_CATEGORY_RESERVED5
	LOG_CATEGORY_RESERVED6
	LOG_CATEGORY_RESERVED7
	LOG_CATEGORY_RESERVED8
	LOG_CATEGORY_RESERVED9
	LOG_CATEGORY_RESERVED10
	LOG_CATEGORY_CUSTOM
)

const (
	LOG_PRIORITY_VERBOSE = iota + 1
	LOG_PRIORITY_DEBUG
	LOG_PRIORITY_INFO
	LOG_PRIORITY_WARN
	LOG_PRIORITY_ERROR
	LOG_PRIORITY_CRITICAL
	NUM_LOG_PRIORITIES
)

type LogPriority C.SDL_LogPriority

func LogSetAllPriority(priority LogPriority) {
	_priority := (C.SDL_LogPriority) (priority)
	C.SDL_LogSetAllPriority(_priority)
}

func LogSetPriority(category int, priority LogPriority) {
	_category := (C.int) (category)
	_priority := (C.SDL_LogPriority) (priority)
	C.SDL_LogSetPriority(_category, _priority)
}

func LogGetPriority(category int) LogPriority {
	_category := (C.int) (category)
	return (LogPriority) (C.SDL_LogGetPriority(_category))
}

func LogResetPriorities() {
	C.SDL_LogResetPriorities()
}

/* TODO:
extern DECLSPEC void SDLCALL SDL_Log(const char *fmt, ...);
extern DECLSPEC void SDLCALL SDL_LogVerbose(int category, const char *fmt, ...);
extern DECLSPEC void SDLCALL SDL_LogDebug(int category, const char *fmt, ...);
extern DECLSPEC void SDLCALL SDL_LogInfo(int category, const char *fmt, ...);
extern DECLSPEC void SDLCALL SDL_LogWarn(int category, const char *fmt, ...);
extern DECLSPEC void SDLCALL SDL_LogError(int category, const char *fmt, ...);
extern DECLSPEC void SDLCALL SDL_LogCritical(int category, const char *fmt, ...);
extern DECLSPEC void SDLCALL SDL_LogMessage(int category,
                                            SDL_LogPriority priority,
                                            const char *fmt, ...);
extern DECLSPEC void SDLCALL SDL_LogMessageV(int category,
                                             SDL_LogPriority priority,
                                             const char *fmt, va_list ap);
extern DECLSPEC void SDLCALL SDL_LogGetOutputFunction(SDL_LogOutputFunction *callback, void **userdata);
extern DECLSPEC void SDLCALL SDL_LogSetOutputFunction(SDL_LogOutputFunction callback, void *userdata);
*/
