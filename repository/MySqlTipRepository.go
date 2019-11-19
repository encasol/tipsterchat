package repository

import (
	"database/sql"
	"github.com/encasol/tipsterchat/model"
	_ "github.com/go-sql-driver/mysql"
)

type MySqlTipRepository struct {
	Connection *sql.DB
}

func (r MySqlTipRepository) All() ([]model.Tip, error) {

	selDB, err := r.Connection.Query("SELECT TipId, Bookie, Analysis, Rate, Stake, Pick FROM Tip")
	if err != nil {
		panic(err.Error())
	}

	tip := model.Tip{}
	tips := []model.Tip{}
	for selDB.Next() {
		err = selDB.Scan(&tip.Id, &tip.Bookie, &tip.Analysis, &tip.Rate, &tip.Stake, &tip.Pick)
		if err != nil {
			panic(err.Error())
		}

		tip.Match, err = r.GetRivals(tip.Id)
		if err != nil {
			panic(err.Error())
		}

		tips = append(tips, tip)
	}

	return tips, nil
}

func (r MySqlTipRepository) GetRivals(TipId int) ([]string, error) {
	selDB, err := r.Connection.Query("SELECT RivalName FROM Rivals Where TipId=?", TipId)
	if err != nil {
		panic(err.Error())
	}

	var rival string
	var rivals []string
	for selDB.Next() {
		err = selDB.Scan(&rival)
		if err != nil {
			panic(err.Error())
		}
		rivals = append(rivals, rival)
	}

	return rivals, nil
}

func (r MySqlTipRepository) Add(tip model.Tip) error {
	insert, err := r.Connection.Prepare("insert into Tip(Bookie, Analysis, Rate, Stake, Pick) Values(?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	result, err := insert.Exec(tip.Bookie, tip.Analysis, tip.Rate, tip.Stake, tip.Pick)
	if err != nil {
		panic(err.Error())
	}

	tipId, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	r.AddRivals(tip.Match, tipId)
	return nil
}

func (r MySqlTipRepository) AddRivals(Rivals []string, TipId int64) error {
	for _, rival := range Rivals {
		r.AddRival(rival, TipId)
	}

	return nil
}

func (r MySqlTipRepository) AddRival(Rival string, TipId int64) error {
	insert, err := r.Connection.Prepare("insert into Rivals(RivalName, TipId) Values(?, ?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = insert.Exec(Rival, TipId)
	return err
}
