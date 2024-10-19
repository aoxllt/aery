package interceptorController

import (
	"aery-go/api/interceptor"
	"aery-go/internal/service"
	"context"
)

type ChangepasswordController struct{}

var ChangePasswordController ChangepasswordController

func (c *ChangepasswordController) Get(ctx context.Context, req *interceptor.ChangePasswordReq) (res *interceptor.ChangePasswordRes, err error) {
	res = &interceptor.ChangePasswordRes{}
	res.Status = service.InterceptorServiceforController().Changepassword(ctx)
	return
}
