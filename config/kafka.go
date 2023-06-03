package config

import (
	"os"
)

const (
	defaultKafkaEnable        = true
	defaultKafkaHost          = "127.0.0.1"
	defaultKafkaPort          = "9092"
	defaultTopicFriendAdded   = "friend_added"
	defaultTopicFriendRemoved = "friend_removed"
)
const (
	envNameKafkaEnable        = "KAFKA_ENABLE"
	envNameKafkaHost          = "KAFKA_HOST"
	envNameKafkaPort          = "KAFKA_PORT"
	envNameTopicFriendAdded   = "TOPIC_FRIEND_ADDED"
	envNameTopicFriendRemoved = "TOPIC_FRIEND_REMOVED"
)

type KafkaConfig struct {
	Enable bool

	Host string
	Port string

	TopicFriendAdded   string
	TopicFriendRemoved string
}

func newDefaultKafkaConfig() KafkaConfig {
	return KafkaConfig{
		Enable:             defaultKafkaEnable,
		Host:               defaultKafkaHost,
		Port:               defaultKafkaPort,
		TopicFriendAdded:   defaultTopicFriendAdded,
		TopicFriendRemoved: defaultTopicFriendRemoved,
	}
}

func (c *KafkaConfig) parseEnv() {
	envKafkaEnable := os.Getenv(envNameKafkaEnable)
	if envKafkaEnable != "" {
		c.Enable = envKafkaEnable == "true"
	}

	envKafkaHost := os.Getenv(envNameKafkaHost)
	if envKafkaHost != "" {
		c.Host = envKafkaHost
	}

	envKafkaPort := os.Getenv(envNameKafkaPort)
	if envKafkaPort != "" {
		c.Port = envKafkaPort
	}

	envTopicFriendAdded := os.Getenv(envNameTopicFriendAdded)
	if envTopicFriendAdded != "" {
		c.TopicFriendAdded = envTopicFriendAdded
	}

	envTopicFriendRemoved := os.Getenv(envNameTopicFriendRemoved)
	if envTopicFriendRemoved != "" {
		c.TopicFriendRemoved = envTopicFriendRemoved
	}
}
