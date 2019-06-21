# 💠 50x50

[![Netlify Status](https://api.netlify.com/api/v1/badges/a244c9b2-a108-4d53-82d4-085ae5293cfd/deploy-status)](https://app.netlify.com/sites/50x50/deploys)
[![Go Report Card](https://goreportcard.com/badge/github.com/koenbollen/50x50)](https://goreportcard.com/report/github.com/koenbollen/50x50)

## About

This project was created as an assignment for Q42.nl

### Assignment

> Maak een 50x50 grid. Als je klikt op een cel wordt bij alle cellen in de rij en kolom van de cel 1 opgeteld. Was een cel leeg dan wordt die op 1 gezet. Na elke verandering licht een cel kort geel op. Als 5 elkaar in de Fibonacci-reeks opvolgende getallen naast elkaar staan, lichten deze cellen kort groen op en worden ze leeg gemaakt. Gebruik de programmeertaal die je het beste vindt passen.

### Key Acceptance Criteria

- There is a 50x50 grid of clickable cells
- Clicking on a set will increase the value of the entire row and column by `+1` (empty cells have a value of `0`).
- Cells pulse yellow when it's value increases
- If five cells in a row are part of the fibonacci sequence:  
    Pulse these cells green and reset the value to 0 (emptying them).

## Local Development

This project is created using Go and javscript. And following the 
[scripts-to-rule-them-all](https://github.com/github/scripts-to-rule-them-all) method:

To get all dependencies up and running just run:

```bash
$ script/bootstrap
```

Then start the server and client separate terminals:

```bash
$ script/server
```

```bash
$ script/client
```

Then open your browser at http://localhost:1234

### Tasks

These are roughly the steps I took, you can also checkout the [Git history](https://github.com/koenbollen/50x50/commits/master)!

- [x] Backend bootstrap
- [x] Basic calls
- [x] Client boilerplate
- [x] Client grid based on data
- [x] Clickable grid
- [x] Backend value increase
- [x] Client yellow pulse
- [x] Backend fibonacci search and action
- [x] Client green pulse
- [x] Documentation and deployment
- [ ] Extract handlers from main
- [ ] Better Error reporting
- [ ] Error handling clientside
