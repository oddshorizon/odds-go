package util

const (
	PLATFORM_UKNOW  = "0"
	PLATFORM_MOBILE = "1"
	PLATFORM_PC     = "2"
	PLATFORM_WEB    = "3"
)

//
//  Mapper2PlatformId
//  @Description: 映射平台ID
//  @param platform
//  @return string
//
func Mapper2PlatformId(platform string) string {
	switch platform {
	case "ios":
		return PLATFORM_MOBILE
	case "android":
		return PLATFORM_MOBILE
	case "macos":
		return PLATFORM_PC
	case "linux":
		return PLATFORM_PC
	case "windows":
		return PLATFORM_PC
	case "fuchsia":
		return PLATFORM_PC
	case "web":
		return PLATFORM_WEB
	default:
		return PLATFORM_UKNOW
	}
}
