package msgsndr

import (
	"encoding/json"
	"time"

	"github.com/avast/retry-go"
	"github.com/rs/zerolog/log"
)

const (
	friendAddRetryDelay    = time.Second * 10
	friendAddRetryAttempts = 5
)

type MsgFriendAdded struct {
	UserID   int64
	FriendID int64
}

func (ms *Sender) SendMsgFriendAdded(userID, friendID int64) {
	if ms == nil {
		return
	}

	msg := MsgFriendAdded{UserID: userID, FriendID: friendID}

	msgBody, err := json.Marshal(msg)
	if err != nil {
		log.Warn().Err(err).Interface("msg", msg).Msg("json.Marshal, msg")
		return
	}

	err = retry.Do(func() error {
		_, err = ms.topicFriendAddedConn.Write(msgBody)
		return err
	},
		retry.Context(ms.liveCtx),
		retry.Delay(friendAddRetryDelay),
		retry.Attempts(friendAddRetryAttempts),
	)

	if err != nil {
		log.Warn().Err(err).Int64("userID", userID).Int64("friendID", friendID).Msg("topicFriendAddedConn.Write")
		return
	}
}
