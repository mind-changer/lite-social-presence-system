package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/lite-social-presence-system/config"
	"github.com/sirupsen/logrus"
)

func main() {
	var err error
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	cfg := &config.Config{}
	err = config.LoadConfigFromEnv(cfg, "../.env")
	if err != nil {
		logrus.Fatal(" Error while loading env config ", err.Error())
		return
	}
	ctx := context.Background()
	dbConfig := cfg.Postgres
	url := fmt.Sprintf("postgresql://%s:%s@%s:%s",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port)
	fmt.Println("postgres url", url)
	conn, dbErr := pgx.Connect(ctx, url)
	if dbErr != nil {
		logrus.Fatalf("Unable to connect to server: %v\n", err)
		return
	}
	dbErr = createDB(ctx, conn)
	if dbErr != nil {
		logrus.Fatalf("Unable to create database: %v\n", err)
		return
	}
	conn, dbErr = pgx.Connect(ctx, url+"/"+dbConfig.Database)
	if dbErr != nil {
		logrus.Fatalf("Unable to connect to database: %v\n", err)
		return
	}
	dbErr = loadDB(ctx, conn)
	if dbErr != nil {
		return
	}
}

func createDB(ctx context.Context, conn *pgx.Conn) error {
	query := `create database mydb;`
	logrus.Info("query: ", query)
	rows, err := conn.Query(ctx, query)
	if err != nil {
		logrus.WithError(err).Error("Error while creating db")
		return err
	}
	rows.Close()
	return nil
}

func loadDB(ctx context.Context, conn *pgx.Conn) error {
	query := `
	create table users(
		id varchar(255) NOT NULL UNIQUE,
		status varchar(255) NOT NULL,
		PRIMARY KEY(id)
	);
	
	create table friends(
		user_id varchar(255) NOT NULL,
		friend_id varchar(255) NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (friend_id) REFERENCES users(id)
	);
	create table parties(
		id varchar(255) NOT NULL UNIQUE,
		owner_id varchar(255) NOT NULL,
		PRIMARY KEY(id),
		FOREIGN KEY (owner_id) REFERENCES users(id)
	);
	create table party_members(
		party_id varchar(255) NOT NULL,
		member_id varchar(255) NOT NULL,
		FOREIGN KEY (party_id) REFERENCES parties(id),
		FOREIGN KEY (member_id) REFERENCES users(id)
	);
	create table party_invitations(
		party_id varchar(255) NOT NULL,
		user_id varchar(255) NOT NULL,
		FOREIGN KEY (party_id) REFERENCES parties(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	create table friend_requests(
		user_id varchar(255) NOT NULL,
		requester_id varchar(255) NOT NULL,
		FOREIGN KEY (requester_id) REFERENCES users(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	logrus.Info("query ", query)
	ct, err := conn.Exec(ctx, query)
	if err != nil {
		logrus.WithError(err).Error("Error while loading db")
		return err
	}
	logrus.Info(ct)
	time.Sleep(2 * time.Second)
	query = `
	insert into users(id,status)
	values 
	 ('bnb','online'),
	 ('hillock123','online'),
	 ('amanora45','online'),
	 ('blankSpac3','online'),
	 ('supergamer','online');
	
	insert into friends(user_id,friend_id)
	values 
	 ('bnb','hillock123'),
	 ('hillock123','bnb'),
	 ('hillock123','amanora45');
	 ('amanora45','hillock123');
	`
	ct, err = conn.Exec(ctx, query)
	if err != nil {
		logrus.WithError(err).Error("Error while loading db")
		return err
	}
	logrus.Info(ct)
	logrus.Info("SUCCESS")
	return nil
}
