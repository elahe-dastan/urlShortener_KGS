package KGS

import "urlShortener/models"

// it should be 8 but for simplicty I put it 2
const LengthOfShortURL  =  2

const source  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var randomShortURLs []models.RandomShortURL

func GenerateAllRandomShortURLsRec(prefix string,k int )  {
	if k == 0 {
		randomShortURLs = append(randomShortURLs, models.RandomShortURL{ShortURL:prefix, IsUsed:false})
		return
	}

	for i := range source {
		newPrefix := prefix + string(source[i])
		GenerateAllRandomShortURLsRec(newPrefix, k-1)
	}
}

func GenerateAllRandomShortURLs()  []models.RandomShortURL{
	GenerateAllRandomShortURLsRec("", LengthOfShortURL)
	return randomShortURLs
}

