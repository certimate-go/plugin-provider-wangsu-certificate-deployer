package certificate

import (
	common "github.com/certimate-go/plugin-provider-wangsu-certificate-deployer/internal/wangsusdk/zz-shared-common"
)

func WithAkSk(ak, sk string) common.OptionsFunc {
	return common.WithAkSk(ak, sk)
}
