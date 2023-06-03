package msgsndr

import (
	"encoding/json"
	"time"

	"github.com/avast/retry-go"
	"github.com/rs/zerolog/log"
)

const (
	friendRemoveRetryDelay    = time.Second * 10
	friendRemoveRetryAttempts = 5
)

type MsgFriendRemoved struct {
	UserID   int64
	FriendID int64
}

func (ms *Sender) SendMsgFriendRemoved(userID, friendID int64) {
	if ms == nil {
		return
	}

	msg := MsgFriendRemoved{UserID: userID, FriendID: friendID}

	msgBody, err := json.Marshal(msg)
	if err != nil {
		log.Warn().Err(err).Interface("msg", msg).Msg("json.Marshal, msg")
		return
	}

	err = retry.Do(func() error {
		_, err = ms.topicFriendRemovedConn.Write(msgBody)
		return err
	},
		retry.Context(ms.liveCtx),
		retry.Delay(friendRemoveRetryDelay),
		retry.Attempts(friendRemoveRetryAttempts),
	)

	if err != nil {
		log.Warn().Err(err).Interface("msg", msg).Msg("topicFriendRemovedConn.Write")
		return
	}
}
