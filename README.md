# AOV ART CRAWLER

- Fast AOV Crawler, save image automatically
### Installations:
Add your art API to ``constants/api.go``

```go
ArtAPI = map[string]string{
"cn": "https://abc/HeroTrainingLoading/##ID##.jpg",
"tw": "https://abc/HeroTrainingLoading/##ID##.jpg",
"vn": "https://abc/HeroTrainingLoading/##ID##.jpg",
}
```

Install go modules:
```bash
go get
```

### Run crawler
```bash
go run main.go
```

OR

```bash
npm run dev
```
