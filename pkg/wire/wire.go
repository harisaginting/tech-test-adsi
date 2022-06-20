//go:build wireinject
// +build wireinject

package wire

import (
	googleWire "github.com/google/wire"
	mCase "github.com/harisaginting/tech-test-adsi/api/v1/cases"
)

func ApiCases() mCase.Controller {
	googleWire.Build(
		mCase.ProviderController,
		mCase.ProviderService,
	)
	return mCase.Controller{}
}
