package service

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"

	"github.com/ReneKroon/ttlcache/v2"
	"github.com/willroberts/minecraft-client"
)

type DispatchService struct {
	Cache *ttlcache.Cache
}

type DispatchRequest struct {
	Address  string   `json:"address" binding:"required"`
	RconPort string   `json:"rcon_port" binding:"required"`
	Password string   `json:"password" binding:"required"`
	Commands []string `json:"commands" binding:"required"`
}

type DispatchResponse struct {
	Status   int         `json:"status"`
	Response interface{} `json:"response"`
}

func NewDispatchService() *DispatchService {

	newCache := func() *ttlcache.Cache {
		expirationCallback := func(_ string, _ ttlcache.EvictionReason, iClient interface{}) {
			client, ok := iClient.(minecraft.Client)

			if !ok {
				logrus.Fatal("Unable to cast client in cache to client! (Shouldn't see this)")
			}

			client.Close() // close client on expiration in cache
		}

		cache := ttlcache.NewCache()
		_ = cache.SetTTL(2 * time.Minute)
		cache.SetExpirationReasonCallback(expirationCallback)
		cache.SetCacheSizeLimit(64)

		return cache
	}

	cache := newCache()
	return &DispatchService{
		Cache: cache,
	}
}

func (req *DispatchRequest) GetAddressPort() string {
	return req.Address + ":" + req.RconPort
}

func (s *DispatchService) DispatchCommands(request *DispatchRequest) DispatchResponse {
	var client *minecraft.Client
	var err error

	addressport := request.GetAddressPort()

	// item not in cache
	if iClient, exists := s.Cache.Get(addressport); exists == nil {
		cli, ok := iClient.(minecraft.Client)

		if !ok {
			logrus.Fatal("Unable to cast client in cache to client! (Shouldn't see this)")
		}

		client = &cli

	} else {
		client, err = minecraft.NewClient(addressport)

		if err != nil {
			return badRequest("Unable to connect to server!")
		}

		_ = s.Cache.Set(addressport, client)
	}

	var responses []string

	if err = client.Authenticate(request.Password); err != nil {
		return badRequest("Unable to authenticate!")
	}

	for _, command := range request.Commands {
		resp, err := client.SendCommand(command)

		if err != nil {
			return badRequest("Unable to execute command " + command + "!")
		}

		responses = append(responses, resp.Body)
	}

	return DispatchResponse{
		Status:   http.StatusOK,
		Response: request.Commands,
	}
}

func badRequest(message string) DispatchResponse {
	return DispatchResponse{
		Status:   http.StatusBadRequest,
		Response: message,
	}
}
