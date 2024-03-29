package controllers

import "go.uber.org/fx"

var Module = fx.Module(
	"controllers",
	fx.Provide(NewMainController),
	fx.Provide(NewAuthController),
	fx.Provide(NewUserController),
	fx.Provide(NewBlogController),
	fx.Provide(NewCommentController),
)
