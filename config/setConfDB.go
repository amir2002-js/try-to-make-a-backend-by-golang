package config

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ConfigSet struct {
	Host     string
	Port     uint16
	User     string
	Password string
	DBName   string
}

func (set *ConfigSet) SetConf(conf *pgxpool.Config) (dbPool *pgxpool.Pool, err error) {
	conf.ConnConfig.Database = set.DBName
	conf.ConnConfig.Host = set.Host
	conf.ConnConfig.Port = set.Port
	conf.ConnConfig.User = set.User
	conf.ConnConfig.Password = set.Password

	dbPool, err = pgxpool.NewWithConfig(context.Background(), conf)
	if err != nil {
		return nil, err
	}
	return dbPool, nil
}
