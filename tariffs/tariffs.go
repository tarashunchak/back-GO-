package TARIFFS

import DB_CONN "NeoCom/connection"

type Tariff struct {
	ID          int     `json:"id"`
	TName       string  `json:"tariff_name"`
	MPrice      float32 `json:"monthly_price"`
	DPrice      float32 `json:"daily_price"`
	CallMinutes float32 `json:"call_minutes"`
	InternetGB  float32 `json:"internet_GB"`
	Messages    int     `json:"messages"`
	IsActive    bool    `json:"is_active"`
}

func SelectAllTariffs() []Tariff {
	rows, err := DB_CONN.Conn.DB.Query("SELECT id, tariff_name, monthly_price, daily_price " +
		", call_minutes, internet_GB, messages, is_active FROM tariffs WHERE is_visible = true")
	if err != nil {
		return nil
	}

	var tariffs []Tariff
	for rows.Next() {
		var tariff Tariff
		err = rows.Scan(&tariff.ID, &tariff.TName, &tariff.MPrice, &tariff.DPrice, &tariff.CallMinutes, &tariff.InternetGB, &tariff.Messages, &tariff.IsActive)

		if err != nil {
			return nil
		}

		tariffs = append(tariffs, tariff)
	}

	return tariffs
}
