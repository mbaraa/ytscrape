# ytscrape

**ytscrape** is a YouTube scraper search library, with a REST API available at [yt.mbaraa.xyz](https://yt.mbaraa.xyz)

[![Deployment status](https://github.com/mbaraa/ytscrape/actions/workflows/rex-deploy.yml/badge.svg)](https://github.com/mbaraa/ytscrape/actions/workflows/rex-deploy.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/mbaraa/ytscrape)](https://goreportcard.com/report/github.com/mbaraa/ytscrape)
[![GoDoc](https://godoc.org/github.com/mbaraa/ytscrape?status.png)](https://godoc.org/github.com/mbaraa/ytscrape)

# Contributing

IDK, it would be really nice of you to contribute, check the poorly written [CONTRIBUTING.md](/CONTRIBUTING.md) for more info.

# Roadmap

- [x] Search YouTube for Videos
- [ ] Search YouTube for Channels
- [ ] Search YouTube for Playlists
- [ ] Search YouTube for Radios

# Usage

```go
import "github.com/mbaraa/ytscrape"

func main() {
    // results is of type []ytscrape.VideoResult
    results, err := ytscrape.Search("Volkswagon das auto")
    if err != nil {
        // handle error
    }
    // do something with the results
}
```

# REST API Docs

Well there's only has a single endpoint:

- **`GET /search`**: accepts a query `q` that has the search term, and responds with a body like this one when `q=Lana del rey`

```json
{
  "id": "TdrL3QxjyVw",
  "title": "Lana Del Rey - Summertime Sadness (Official Music Video)",
  "url": "https://youtube.com/watch?v=TdrL3QxjyVw",
  "duration": 266,
  "thumbnail_src": "https://i.ytimg.com/vi/TdrL3QxjyVw/hq720.jpg",
  "views": 574902158,
  "uploader": {
    "title": "Lana Del Rey",
    "url": "https://www.youtube.com/channel/UCqk3CdGN_j8IR9z4uBbVPSg"
  }
}
```

# Run REST API locally

1. Clone the repo.

```bash
git clone https://github.com/mbaraa/ytscrape
```

2. Run it with docker compose.

```bash
docker compose up
```

3. Visit http://localhost:20256
4. Don't ask why I chose this weird port.

---

Made with 🧉 by [Baraa Al-Masri](https://mbaraa.com)
