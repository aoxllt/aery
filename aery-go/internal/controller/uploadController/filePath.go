package uploadController

import (
	"aery-go/api/upload"
	"context"
	"fmt"
)

type FilePathController struct {
}

var FilePath FilePathController

func (c *FilePathController) Post(ctx context.Context, req *upload.FilePathReq) (res *upload.FilePathRes, err error) {
	res = &upload.FilePathRes{}
	fmt.Println(req.Paths)
	res.Status = true
	return res, nil
}
