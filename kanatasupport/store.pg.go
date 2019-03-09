package kanatasupport

import (
	"strconv"

	"github.com/premeidoworks/kanata/api"

	"github.com/jackc/pgx"
)

var (
	defaultPgStore = new(PostgresStore)
)

type PostgresStore struct {
	db *pgx.ConnPool
}

func (this *PostgresStore) Init(config *api.StoreInitConfig) error {
	pgxConfig := pgx.ConnPoolConfig{}
	pgxConfig.Host = config.Details["Host"]

	port, err := strconv.Atoi(config.Details["Port"])
	if err != nil {
		return err
	}
	pgxConfig.Port = uint16(port)

	pgxConfig.Database = config.Details["Database"]
	pgxConfig.User = config.Details["User"]
	pgxConfig.Password = config.Details["Password"]

	maxConnections, err := strconv.Atoi(config.Details["MaxConnections"])
	if err != nil {
		return err
	}
	pgxConfig.MaxConnections = maxConnections
	pool, err := pgx.NewConnPool(pgxConfig)
	if err != nil {
		return err
	}
	this.db = pool
	return nil
}

func (this *PostgresStore) SaveMessage(msg *api.Message) error {
	var scheduleTime interface{}
	if msg.ScheduleTime <= 0 {
		scheduleTime = nil
	} else {
		scheduleTime = msg.ScheduleTime
	}

	var outId interface{}
	if len(msg.OutId) == 0 {
		outId = nil
	} else {
		outId = msg.OutId
	}

	_, err := this.db.Exec(
		"INSERT INTO public.messages(message_id, queue, topic, body, schedule_time, status, out_id, create_time)"+
			" VALUES ($1, $2, $3, $4, $5, $6, $7, now())",
		msg.MessageId, msg.Queue, msg.Topic, msg.Body, scheduleTime, msg.Status, outId,
	)
	if err != nil {
		return err
	}
	return nil
}
