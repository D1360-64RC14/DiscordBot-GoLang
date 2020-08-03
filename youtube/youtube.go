package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/D1360-64RC14/config"
)

// GetSearch :
// Retorna um struct com o JSON do Search
func GetSearch(search string) SearchOut {
	// Query values
	var urlQueryParams = url.Values{}
	urlQueryParams.Add("part", "snippet,id")
	urlQueryParams.Add("type", "video")
	urlQueryParams.Add("maxResults", "1")
	urlQueryParams.Add("key", config.Data.YoutubeAPI.Token)
	urlQueryParams.Add("q", search)

	var requestURL = fmt.Sprintf("%s?%s",
		config.Data.YoutubeAPI.RequestURLs.SearchAPI,
		urlQueryParams.Encode(),
	)

	var requestBody = makeRequest(requestURL)
	var searchDataStruct SearchOut

	json.Unmarshal(requestBody, &searchDataStruct)

	if searchDataStruct.Error != nil {
		return SearchOut{}
	}

	return searchDataStruct
}

// GetChannels :
// Retorna um struct com o JSON do Canal
func GetChannels(channelID string) ChannelsOut {
	// Query values
	var urlQueryParams = url.Values{}
	urlQueryParams.Add("part", "snippet,id")
	urlQueryParams.Add("maxResults", "1")
	urlQueryParams.Add("id", channelID)
	urlQueryParams.Add("key", config.Data.YoutubeAPI.Token)

	var requestURL = fmt.Sprintf("%s?%s",
		config.Data.YoutubeAPI.RequestURLs.ChannelsAPI,
		urlQueryParams.Encode(),
	)

	var requestBody = makeRequest(requestURL)
	var channelsDataStruct ChannelsOut

	json.Unmarshal(requestBody, &channelsDataStruct)

	if channelsDataStruct.Error != nil {
		return channelsDataStruct
	}

	return channelsDataStruct
}

func makeRequest(requestURL string) []byte {
	var response, responseERR = http.Get(requestURL)
	if responseERR != nil {
		fmt.Println(responseERR.Error())
	}
	var responseBody, responseBodyERR = ioutil.ReadAll(response.Body)
	if responseBodyERR != nil {
		fmt.Println(responseBodyERR.Error())
	}
	response.Body.Close()
	return responseBody
}
