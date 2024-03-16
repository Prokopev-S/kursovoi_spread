package candles

import (
    // "fmt"
    "kursovoi_spread/types"
)

func FetchAllData() ([]types.ExchangeFunding, []types.ExchangeFunding){
	spotData := make([]types.ExchangeFunding, 0, 6) 
    futuresData := make([]types.ExchangeFunding, 0, 6) 

	return spotData, futuresData
}  