//go:build wireinject
// +build wireinject

package wire

import (
	googleWire "github.com/google/wire"
	mUser "github.com/harisaginting/ginting/api/v1/user"
	"gorm.io/gorm"
)

func ApiUser(db *gorm.DB) mUser.Controller {
	googleWire.Build(
		mUser.ProviderController,
		mUser.ProviderService,
		mUser.ProviderRepository,
	)
	return mUser.Controller{}
}
