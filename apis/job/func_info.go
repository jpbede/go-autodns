package job

import (
	"context"
	"errors"
	"strconv"
)

func (c *client) Info(jobID int, ctx context.Context) (*ObjectJob, error) {
	var out ObjectJobResponse
	if err := c.transport.Get(ctx, "/job/"+strconv.Itoa(jobID), &out); err != nil {
		return nil, err
	}
	if len(out.Data) > 0 {
		return out.Data[0], nil
	}
	return nil, errors.New("no job found")
}
