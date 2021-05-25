package una

var _ option = (*optionAppKey)(nil)

type optionAppKey struct {
	// 应用
	key string
}

// AppKey 配置应用
func AppKey(key string) *optionAppKey {
	return &optionAppKey{
		key: key,
	}
}

func (as *optionAppKey) apply(options *options) {
	options.chuangcache.appKey = as.key
}
