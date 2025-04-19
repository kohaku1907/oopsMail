package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage() *RedisStorage {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &RedisStorage{
		client: client,
	}
}

func (s *RedisStorage) CreateMailbox(ctx context.Context, id string, expiration time.Duration) error {
	key := fmt.Sprintf("mailbox:%s", id)
	return s.client.Set(ctx, key, "{}", expiration).Err()
}

func (s *RedisStorage) StoreEmail(ctx context.Context, mailboxID string, email *Email) error {
	key := fmt.Sprintf("mailbox:%s:emails", mailboxID)
	emailJSON, err := json.Marshal(email)
	if err != nil {
		return err
	}

	return s.client.RPush(ctx, key, emailJSON).Err()
}

func (s *RedisStorage) GetEmails(ctx context.Context, mailboxID string) ([]*Email, error) {
	key := fmt.Sprintf("mailbox:%s:emails", mailboxID)
	emailsJSON, err := s.client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	emails := make([]*Email, 0, len(emailsJSON))
	for _, emailJSON := range emailsJSON {
		var email Email
		if err := json.Unmarshal([]byte(emailJSON), &email); err != nil {
			return nil, err
		}
		emails = append(emails, &email)
	}

	return emails, nil
}

func (s *RedisStorage) DeleteMailbox(ctx context.Context, mailboxID string) error {
	mailboxKey := fmt.Sprintf("mailbox:%s", mailboxID)
	emailsKey := fmt.Sprintf("mailbox:%s:emails", mailboxID)

	pipe := s.client.Pipeline()
	pipe.Del(ctx, mailboxKey)
	pipe.Del(ctx, emailsKey)
	_, err := pipe.Exec(ctx)
	return err
}

func (s *RedisStorage) MailboxExists(ctx context.Context, mailboxID string) (bool, error) {
	key := fmt.Sprintf("mailbox:%s", mailboxID)
	exists, err := s.client.Exists(ctx, key).Result()
	return exists > 0, err
}
