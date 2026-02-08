package ytbgo

import (
	"errors"
	"net/http"
)

// InnerTube struct
type InnerTube struct {
	Adaptor Adaptor
}

// Adaptor interface
type Adaptor interface {
	Dispatch(endpoint string, params map[string]string, body map[string]any) (map[string]any, error)
}

// NewInnerTube creates a new InnerTube instance
func NewInnerTube(httpClient *http.Client, clientName, clientVersion string, apiKey, userAgent, referer string, locale *Locale, auto bool, cookie string, xGoogVisitorId string) (*InnerTube, error) {
	if clientName == "" {
		return nil, errors.New("precondition failed: Missing client name")
	}

	if clientVersion == "" {
		return nil, errors.New("precondition failed: Missing client version")
	}
	var context ClientContext
	if auto {
		context = GetContext(clientName)
	} else {
		context = ClientContext{
			ClientName:     clientName,
			ClientVersion:  clientVersion,
			APIKey:         apiKey,
			UserAgent:      userAgent,
			Referer:        referer,
			Locale:         locale,
			Cookie:         cookie,
			XGoogVisitorId: xGoogVisitorId,
		}
	}

	return &InnerTube{
		Adaptor: NewInnerTubeAdaptor(context, httpClient),
	}, nil
}

// Call method to make requests
func (it *InnerTube) Call(endpoint string, params map[string]string, body map[string]any) (map[string]any, error) {
	response, err := it.Adaptor.Dispatch(endpoint, params, body)
	if err != nil {
		return nil, err
	}

	delete(response, "responseContext") // Remove responseContext if exists
	return response, nil
}

// Example API call methods
func (it *InnerTube) Config() (map[string]any, error) {
	return it.Call("CONFIG", nil, nil)
}

func (it *InnerTube) Guide() (map[string]any, error) {
	return it.Call("GUIDE", nil, nil)
}

func (it *InnerTube) Player(videoID string) (map[string]any, error) {
	return it.Call("PLAYER", nil, Filter(map[string]any{
		"videoId": videoID,
	}))
}

func (it *InnerTube) Browse(browseID *string, params *string, continuation *string) (map[string]any, error) {
	body := map[string]any{
		"browseId":     browseID,
		"params":       params,
		"continuation": continuation,
	}
	//fmt.Println("body: ", body)
	//fmt.Println("Filter(body): ", Filter(body))
	return it.Call("BROWSE", nil, Filter(body))
}

func (it *InnerTube) Search(query *string, params *string, continuation *string) (map[string]any, error) {
	body := map[string]any{
		"query":        query,
		"params":       params,
		"continuation": continuation,
	}
	//fmt.Println("body: ", body)
	//fmt.Println("Filter(body): ", Filter(body))
	return it.Call("SEARCH", nil, Filter(body))
}

func (it *InnerTube) Next(videoId *string, playlistId *string, params *string, index *int, continuation *string) (map[string]any, error) {
	body := map[string]any{
		"videoId":       videoId,
		"playlistId":    playlistId,
		"params":        params,
		"playlistIndex": index,
		"continuation":  continuation,
	}
	//fmt.Println("body: ", body)
	//fmt.Println("Filter(body): ", Filter(body))
	return it.Call("NEXT", nil, Filter(body))
}

func (it *InnerTube) GetTranscript(params *string) (map[string]any, error) {
	body := map[string]any{
		"params": params,
	}
	//fmt.Println("body: ", body)
	//fmt.Println("Filter(body): ", Filter(body))
	return it.Call("GET_TRANSCRIPT", nil, Filter(body))
}

func (it *InnerTube) MusicGetSearchSuggestions(input *string) (map[string]any, error) {
	body := map[string]any{
		"input": input,
	}
	//fmt.Println("body: ", body)
	//fmt.Println("Filter(body): ", Filter(body))
	return it.Call("MUSIC/GET_SEARCH_SUGGESTIONS", nil, Filter(body))
}

func (it *InnerTube) MusicGetQueue(videoIds *[]string, playlistId *string) (map[string]any, error) {
	body := map[string]any{
		"playlistId": playlistId,
		"videoIds":   videoIds,
	}
	//fmt.Println("body: ", body)
	//fmt.Println("Filter(body): ", Filter(body))
	return it.Call("MUSIC/GET_QUEUE", nil, Filter(body))
}
