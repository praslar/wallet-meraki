package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"wallet/internal/model"
	"wallet/internal/repo"
)

type TokenService struct {
	Repo repo.IRepo
}

func NewTokenService(repo repo.IRepo) TokenService {
	return TokenService{
		Repo: repo,
	}
}

func (s *TokenService) TriggerCrawl() error {
	response, err := http.Get("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1&sparkline=false&locale=en")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var rawToken []model.RawToken
	err = json.Unmarshal(responseData, &rawToken)
	if err != nil {
		fmt.Println("Can't unmarshal the byte array")
		return err
	}

	var TotalToken []model.Token
	var baseToken model.Token
	for _, v := range rawToken {
		baseToken.ID = v.ID
		baseToken.Symbol = v.Symbol
		baseToken.Name = v.Name
		baseToken.Image = v.Image
		baseToken.CurrentPrice = v.CurrentPrice
		baseToken.MarketCap = v.MarketCap
		baseToken.MarketCapRank = v.MarketCapRank
		baseToken.FullyDilutedValuation = v.FullyDilutedValuation
		baseToken.TotalVolume = v.TotalVolume
		baseToken.High24H = v.High24H
		baseToken.Low24H = v.Low24H
		baseToken.PriceChange24H = v.PriceChange24H
		baseToken.PriceChangePercentage24H = v.PriceChangePercentage24H
		baseToken.MarketCapChange24H = v.MarketCapChange24H
		baseToken.MarketCapChangePercentage24H = v.MarketCapChangePercentage24H
		baseToken.CirculatingSupply = v.CirculatingSupply
		baseToken.TotalSupply = v.TotalSupply
		baseToken.MaxSupply = v.MaxSupply
		baseToken.Ath = v.Ath
		baseToken.AthChangePercentage = v.AthChangePercentage
		baseToken.AthDate = v.AthDate
		baseToken.Roi = fmt.Sprint(v.Roi)
		baseToken.LastUpdated = v.LastUpdated
		TotalToken = append(TotalToken, baseToken)
	}

	err = s.Repo.CreateToken(TotalToken)
	if err != nil {
		return err
	}
	return nil
}
