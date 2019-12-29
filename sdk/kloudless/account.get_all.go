package kloudless

import (
	"context"
	"errors"
	"net/http"

	"github.com/cymonevo/cloud-api/handler"
	"github.com/cymonevo/cloud-api/internal/log"
)

const getAllAccounts = "/accounts"

func (c *clientImpl) GetAccounts(ctx context.Context) (interface{}, error) {
	resp, err := c.client.Get(getAllAccounts, nil, c.headers(nil))
	if err != nil {
		log.ErrorDetail("GetAccounts", "error get accounts", err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Warnf("GetAccounts", "status %d %s", resp.StatusCode, resp.Status)
		return nil, errors.New(resp.Status)
	}
	var data interface{}
	err = handler.ParseBody(resp.Body, &data)
	if err != nil {
		log.ErrorDetail("GetAccounts", "error unmarshal data", err)
		return nil, err
	}
	log.Infof("SDK", "success get accounts %+v", data)
	return data, nil
}
