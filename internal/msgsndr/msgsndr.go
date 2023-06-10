package msgsndr

import (
	"context"
	"net"

	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"

	"github.com/v1tbrah/relation-service/config"
)

type Sender struct {
	liveCtx context.Context

	topicFriendAddedConn   *kafka.Conn
	topicFriendRemovedConn *kafka.Conn

	cfg config.Kafka
}

func New(ctx context.Context, cfg config.Kafka) (*Sender, error) {
	if !cfg.Enable {
		return nil, nil
	}

	topicFriendAddedConn, err := kafka.DialLeader(ctx, "tcp", net.JoinHostPort(cfg.Host, cfg.Port), cfg.TopicFriendAdded, 0)
	if err != nil {
		return nil, errors.Wrapf(err, "kafka.DialLeader, address (%s), topic (%s)", net.JoinHostPort(cfg.Host, cfg.Port), cfg.TopicFriendAdded)
	}

	topicFriendRemovedConn, err := kafka.DialLeader(ctx, "tcp", net.JoinHostPort(cfg.Host, cfg.Port), cfg.TopicFriendRemoved, 0)
	if err != nil {
		return nil, errors.Wrapf(err, "kafka.DialLeader, address (%s), topic (%s)", net.JoinHostPort(cfg.Host, cfg.Port), cfg.TopicFriendRemoved)
	}

	return &Sender{
		liveCtx:                ctx,
		topicFriendAddedConn:   topicFriendAddedConn,
		topicFriendRemovedConn: topicFriendRemovedConn,
		cfg:                    cfg,
	}, nil
}

func (ms *Sender) Close(ctx context.Context) (err error) {
	if ms == nil {
		return nil
	}

	closed := make(chan struct{})

	go func() {
		if err = ms.topicFriendAddedConn.Close(); err != nil {
			err = errors.Wrap(err, "topicFriendAddedConn.Close")
			closed <- struct{}{}
			return
		}

		if err = ms.topicFriendRemovedConn.Close(); err != nil {
			err = errors.Wrap(err, "topicFriendRemovedConn.Close")
			closed <- struct{}{}
			return
		}

		closed <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-closed:
		return err
	}
}
