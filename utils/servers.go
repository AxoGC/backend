package utils

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mcstatus-io/mcutil/v4/query"
	"github.com/mcstatus-io/mcutil/v4/status"
)

var (
	ErrNotSupportedGame = errors.New("not supported game")
)

func (s *Server) GetOnlineCount(host string) (*int64, error) {

	ctx, canc := context.WithTimeout(context.Background(), time.Second*5)
	defer canc()

	switch s.GameID {
	case MinecraftBedrock:
		resp, err := status.Bedrock(ctx, host, s.Port)
		if err != nil {
			return nil, fmt.Errorf("failed to lookup: %w", err)
		}
		return resp.OnlinePlayers, nil
	case MinecraftJava:
		resp, err := status.Modern(ctx, host, s.Port)
		if err != nil {
			return nil, fmt.Errorf("failed to lookup: %w", err)
		}
		return resp.Players.Online, nil
	default:
		return nil, ErrNotSupportedGame
	}
}

func (s *Server) GetOnlineList(host string, bcs map[string][]BedrockCommand) ([]string, error) {

	ctx, canc := context.WithTimeout(context.Background(), time.Second*5)
	defer canc()

	switch s.GameID {
	case MinecraftBedrock:
	case MinecraftJava:
		resp, err := query.Full(ctx, host, s.Port)
		if err != nil {
			return nil, fmt.Errorf("failed to lookup: %w", err)
		}
		return resp.Players, nil
	default:
		return nil, nil
	}
	return nil, nil
}

func SetBedrockCommand(bcs map[string][]BedrockCommand) {

}
