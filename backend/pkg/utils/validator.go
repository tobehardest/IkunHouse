package utils

//const (
//	ValidatorKey  = "ValidatorKey"
//	TranslatorKey = "TranslatorKey"
//	local         = "chinese"
//)
//
//func TransInit(ctx context.Context) context.Context {
//	chinese := zh.New()
//	english := en.New()
//	// 设置国际化翻译器
//	uni := ut.New(chinese, chinese, english)
//	// 设置验证器
//	val := validator.New()
//	// 根据参数取翻译器实例
//	trans, _ := uni.GetTranslator(local)
//
//	switch local {
//	case "chinese":
//		zhTranslations.RegisterDefaultTranslations(val, trans)
//		// 使用fld.Tag.Get("")注册一个获取tag的自定义方法
//		val.RegisterTagNameFunc(func(fld reflect.StructField) string {
//			return fld.Tag.Get("comment")
//		})
//	}
//}
