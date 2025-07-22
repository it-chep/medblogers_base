package cache

//
////go:generate mockgen -destination=mocks/mocks.go -package=mocks . Cacher
//type Cacher interface {
//	GetMulti(ctx context.Context, keys []string) (map[string]*memcache.Item, error)
//	Get(ctx context.Context, key string) (*memcache.Item, error)
//	SetMultiV2(ctx context.Context, items []*memcache.Item) error
//	Set(ctx context.Context, items *memcache.Item) error
//}
//
//type GenericMemcached[TValue any] struct {
//	mc Cacher
//}
//
//func NewGenericMemcached[TValue any](ctx context.Context, host string) *GenericMemcached[TValue] {
//	mc, err := memcache.New([]string{host},
//		memcache.NewManualConfigProvider(),
//		memcache.WithMaxInFlightRequestsPerNode(150),
//		memcache.WithMaxInFlightRequests(100),
//	)
//
//	if err != nil {
//		logger.FatalKV(ctx, fmt.Sprintf("[FATAL] не удалось создать клиент memcached: %s", err.Error()))
//	}
//
//	return &GenericMemcached[TValue]{mc: mc}
//}
//
//func WrapCached[TValue any](mc Cacher) *GenericMemcached[TValue] {
//	return &GenericMemcached[TValue]{mc: mc}
//}
//
//func (c *GenericMemcached[TValue]) Set(ctx context.Context, key string, value TValue, ttl time.Duration) {
//	bytes, _ := json.Marshal(value)
//
//	item := &memcache.Item{Key: fmt.Sprint(key), Value: bytes, Expiration: int32(ttl.Seconds())}
//
//	if err := c.mc.Set(ctx, item); err != nil {
//		logger.Warnf(ctx, "ошибка записи в кэш [%v]: %s", key, err.Error())
//	}
//}
//
//func (c *GenericMemcached[TValue]) Get(ctx context.Context, key string) (TValue, bool) {
//	var value TValue
//	item, err := c.mc.Get(ctx, fmt.Sprint(key))
//	if err != nil {
//		return value, false
//	}
//
//	if err = json.Unmarshal(item.Value, &value); err != nil {
//		return value, false
//	}
//
//	return value, true
//}
//
//func (c *GenericMemcached[TValue]) GetMulti(ctx context.Context, key ...string) (map[string]TValue, bool) {
//	items, err := c.mc.GetMulti(ctx, key)
//	if err != nil {
//		return nil, false
//	}
//
//	result := make(map[string]TValue)
//	for k, v := range items {
//		var value TValue
//		if err = json.Unmarshal(v.Value, &value); err != nil {
//			return nil, false
//		}
//		result[k] = value
//	}
//
//	return result, true
//}
//
//func (c *GenericMemcached[TValue]) SetMulti(ctx context.Context, values []lo.Tuple2[string, TValue], ttl time.Duration) {
//	items := lo.Map(values, func(item lo.Tuple2[string, TValue], _ int) *memcache.Item {
//		bytes, _ := json.Marshal(item.B)
//
//		return &memcache.Item{Key: item.A, Value: bytes, Expiration: int32(ttl.Seconds())}
//	})
//
//	if err := c.mc.SetMultiV2(ctx, items); err != nil {
//		logger.Error(ctx, "Ошибка сохранения в кеш: %s", err)
//	}
//
//}
