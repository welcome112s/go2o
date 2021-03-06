package impl

/**
 * Copyright 2015 @ to2.net.
 * name : platform_service
 * author : jarryliu
 * date : 2020-09-05 15:30
 * description :
 * history :
 */

import (
	"context"
	"github.com/ixre/gof"
	"go2o/core/domain/interface/mss/notify"
	"go2o/core/domain/interface/registry"
	"go2o/core/domain/interface/valueobject"
	"go2o/core/infrastructure/domain"
	"go2o/core/infrastructure/format"
	"go2o/core/module"
	"go2o/core/module/bank"
	"go2o/core/service/proto"
)

var _ proto.FoundationServiceServer = new(foundationService)

// 基础服务
type foundationService struct {
	_rep         valueobject.IValueRepo
	registryRepo registry.IRegistryRepo
	notifyRepo   notify.INotifyRepo
	serviceUtil
}

func NewFoundationService(rep valueobject.IValueRepo, registryRepo registry.IRegistryRepo, notifyRepo notify.INotifyRepo) *foundationService {
	return &foundationService{
		_rep:         rep,
		notifyRepo:   notifyRepo,
		registryRepo: registryRepo,
	}
}

// 保存短信API凭据
func (s *foundationService) SaveSmsApi(_ context.Context, r *proto.SmsApiSaveRequest) (*proto.Result, error) {
	manager := s.notifyRepo.Manager()
	a := r.Api
	perm := &notify.SmsApiPerm{
		Key:         a.Key,
		Secret:      a.Secret,
		ApiUrl:      a.ApiUrl,
		Params:      a.Params,
		Method:      a.Method,
		Charset:     a.Charset,
		Signature:   a.Signature,
		SuccessChar: a.SuccessChar,
	}
	if err := manager.SaveSmsApiPerm(r.Provider, perm); err != nil {
		return s.error(err), nil
	}
	return s.success(nil), nil
}

// 获取短信API凭据, @provider 短信服务商, 默认:http
func (s *foundationService) GetSmsApi(_ context.Context, provider *proto.String) (*proto.SSmsApi, error) {
	manager := s.notifyRepo.Manager()
	perm := manager.GetSmsApiPerm(provider.Value)
	if perm != nil {
		return &proto.SSmsApi{
			ApiUrl:      perm.ApiUrl,
			Key:         perm.Key,
			Secret:      perm.Secret,
			Params:      perm.Params,
			Method:      perm.Method,
			Charset:     perm.Charset,
			Signature:   perm.Signature,
			SuccessChar: perm.SuccessChar,
		}, nil
	}
	return &proto.SSmsApi{
		ApiUrl:      "",
		Key:         "",
		Secret:      "",
		Params:      "",
		Method:      "",
		Charset:     "",
		Signature:   "",
		SuccessChar: "",
	}, nil
}

// 保存面板HOOK数据,这通常是在第三方应用中初始化或调用,参见文档：BoardHooks
func (s *foundationService) SaveBoardHook(_ context.Context, request *proto.BoardHookSaveRequest) (*proto.Result, error) {

	mp := map[string]string{
		registry.BoardHookURL:   request.HookURL,
		registry.BoardHookToken: request.Token,
	}
	for k, v := range mp {
		if ir := s.registryRepo.Get(k); ir != nil {
			if err := ir.Update(v); err != nil {
				return s.error(err), nil
			}
		}
	}
	return s.success(nil), nil
}

//
//// 删除值
//func (s *foundationService) DeleteValue(_ context.Context, s2 *proto.String) (*proto.Result, error) {
//	err := s._rep.DeleteValue(s2.Value)
//	return s.result(err), nil
//}
//
//// 根据前缀获取值
//func (s *foundationService) GetValuesByPrefix(_ context.Context, s2 *proto.String) (*proto.StringMap, error) {
//	return &proto.StringMap{
//		Value: s._rep.GetValues(s2.Value),
//	}, nil
//}

// 验证超级用户账号和密码
func (s *foundationService) SuperValidate(_ context.Context, user *proto.UserPwd) (*proto.Bool, error) {
	superPwd := gof.CurrentApp.Config().Get("super_login_md5")
	encPwd := domain.Sha1Pwd(user.Pwd + user.User)
	return &proto.Bool{Value: superPwd == encPwd}, nil
}

// 保存超级用户账号和密码
func (s *foundationService) FlushSuperPwd(_ context.Context, user *proto.UserPwd) (*proto.Empty, error) {
	conf := gof.CurrentApp.Config()
	encPwd := domain.Sha1Pwd(user.Pwd + user.User)
	conf.Set("super_login_md5", encPwd)
	//conf.Flush()
	//todo:
	panic("暂不支持保存")
}

// 注册单点登录应用,返回值：
//   -  1. 成功，并返回token
//   - -1. 接口地址不正确
//   - -2. 已经注册
func (s *foundationService) RegisterApp(_ context.Context, app *proto.SSsoApp) (*proto.String, error) {
	sso := module.Get(module.SSO).(*module.SSOModule)
	token, err := sso.Register(app)
	if err == nil {
		return &proto.String{
			Value: "1:" + token,
		}, nil
	}
	return &proto.String{
		Value: err.Error(),
	}, nil
}

// 获取应用信息
func (s *foundationService) GetApp(_ context.Context, s2 *proto.String) (*proto.SSsoApp, error) {
	sso := module.Get(module.SSO).(*module.SSOModule)
	return sso.Get(s2.Value), nil
}

// 获取单点登录应用
func (s *foundationService) GetAllSsoApp(_ context.Context, _ *proto.Empty) (*proto.StringListResponse, error) {
	sso := module.Get(module.SSO).(*module.SSOModule)
	return &proto.StringListResponse{
		Value: sso.Array(),
	}, nil
}

// 创建同步登录的地址
func (s *foundationService) GetSyncLoginUrl(_ context.Context, s2 *proto.String) (*proto.String, error) {
	panic("not implement")
	//return fmt.Sprintf("%s://%s%s/auth?return_url=%s",
	//	consts.DOMAIN_PASSPORT_PROTO, consts.DOMAIN_PREFIX_PASSPORT,
	//	variable.Domain, returnUrl), nil
}

// 获取移动应用设置
func (s *foundationService) GetMoAppConf() valueobject.MoAppConf {
	return s._rep.GetMoAppConf()
}

// 保存移动应用设置
func (s *foundationService) SaveMoAppConf(r *valueobject.MoAppConf) error {
	return s._rep.SaveMoAppConf(r)
}

// 获取微信接口配置
func (s *foundationService) GetWxApiConfig() valueobject.WxApiConfig {
	return s._rep.GetWxApiConfig()
}

// 保存微信接口配置
func (s *foundationService) SaveWxApiConfig(v *valueobject.WxApiConfig) error {
	return s._rep.SaveWxApiConfig(v)
}

// 获取资源地址
func (s *foundationService) ResourceUrl(_ context.Context, s2 *proto.String) (*proto.String, error) {
	return &proto.String{Value: format.GetResUrl(s2.Value)}, nil
}

// 获取全局商户销售设置
func (s *foundationService) GetGlobMchSaleConf() valueobject.GlobMchSaleConf {
	return s._rep.GetGlobMchSaleConf()
}

// 保存全局商户销售设置
func (s *foundationService) SaveGlobMchSaleConf(v *valueobject.GlobMchSaleConf) error {
	return s._rep.SaveGlobMchSaleConf(v)
}

// 获取短信设置
func (s *foundationService) GetSmsApiSet() notify.SmsApiSet {
	//return s._rep.GetSmsApiSet()
	//todo: will remove
	return notify.SmsApiSet{}
}

// 获取下级区域
func (s *foundationService) GetChildAreas(_ context.Context, code *proto.Int32) (*proto.AreaListResponse, error) {
	var arr []*proto.SArea
	for _, v := range s._rep.GetChildAreas(code.Value) {
		arr = append(arr, &proto.SArea{
			Code:   int32(v.Code),
			Parent: int32(v.Parent),
			Name:   v.Name,
		})
	}
	return &proto.AreaListResponse{
		Value: arr,
	}, nil
}

// 获取地区名称
func (s *foundationService) GetAreaNames(_ context.Context, request *proto.GetAreaNamesRequest) (*proto.StringListResponse, error) {
	return &proto.StringListResponse{
		Value: s._rep.GetAreaNames(request.Value),
	}, nil
}

// 获取省市区字符串
func (s *foundationService) GetAreaString(province, city, district int32) string {
	if province == 0 || city == 0 || district == 0 {
		return ""
	}
	return s._rep.GetAreaString(province, city, district)
}

// 获取支付平台
func (s *foundationService) GetPayPlatform() []*bank.PaymentPlatform {
	m := module.Get(module.PAY).(*module.PaymentModule)
	return m.GetPayPlatform()
}
