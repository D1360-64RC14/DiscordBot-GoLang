package youtube

import "fmt"

// SearchOut :
// Estrutura do JSON de Search
type SearchOut struct {
	Error         *youtubeErrorOut `json:",omitempty"`
	NextPageToken string
	PrevPageToken string
	Items         []struct {
		ID struct {
			VideoID string
		}
		Snippet struct {
			ChannelID  string
			Title      string
			Thumbnails struct {
				Medium struct {
					URL    string
					Height uint16
				}
				High struct {
					URL    string
					Height uint16
				}
			}
			ChannelTitle string
		}
	}
}

// ChannelsOut :
// Estrutura do JSON de Channels
type ChannelsOut struct {
	Error *youtubeErrorOut `json:",omitempty"`
	Items []struct {
		Snippet struct {
			Title       string
			Description string
			PublishedAt string
			Thumbnails  struct {
				High struct {
					URL string
				}
			}
			Country string
		}
	}
}

type youtubeErrorOut struct {
	Code    uint16
	Message string
}

// String :
// Método de `Error` que retorna o
// error code e message formatados.
func (i youtubeErrorOut) String() string {
	return fmt.Sprintf("`ERROR:\nCode: %s\nMessage: %s`",
		string(i.Code),
		i.Message,
	)
}
