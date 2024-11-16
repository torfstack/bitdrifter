package logging

import "context"

type dynamicLogTag string

const (
	dynamicLogTagGameSessionId dynamicLogTag = "game_session_id"
	dynamicLogTagUserId        dynamicLogTag = "user_id"
)

var dynamicTags = []dynamicLogTag{
	dynamicLogTagGameSessionId,
	dynamicLogTagUserId,
}

func SetGameSessionId(ctx context.Context, gameSessionId string) context.Context {
	return context.WithValue(ctx, dynamicLogTagGameSessionId, gameSessionId)
}

func SetUserId(ctx context.Context, userId string) context.Context {
	return context.WithValue(ctx, dynamicLogTagUserId, userId)
}
