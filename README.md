# 50x50

[![Netlify Status](https://api.netlify.com/api/v1/badges/a244c9b2-a108-4d53-82d4-085ae5293cfd/deploy-status)](https://app.netlify.com/sites/50x50/deploys)

## Assignment

> Maak een 50x50 grid. Als je klikt op een cel wordt bij alle cellen in de rij en kolom van de cel 1 opgeteld. Was een cel leeg dan wordt die op 1 gezet. Na elke verandering licht een cel kort geel op. Als 5 elkaar in de Fibonacci-reeks opvolgende getallen naast elkaar staan, lichten deze cellen kort groen op en worden ze leeg gemaakt. Gebruik de programmeertaal die je het beste vindt passen.

### Disected

- 50x50 grid of clickable cells.
- click cell and the row and column gets +1 (empty = 0)
- cells pulse yellow on value increase

- if 5 cells in a row/col is in the fibonacci sequence:  
    pulse green and empty cells

## Local Development

This project is created using Go and javscript. To get all dependencies up and running just run:

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

### Tasks

- [x] Backend bootstrap
- [x] Basic calls
- [x] Client boilerplate
- [x] Client grid based on data
- [x] Clickable grid
- [x] Backend value increase
- [x] Client yellow pulse
- [x] Backend fibonacci search and action
- [x] Client green pulse
- [ ] Documentation and deployment
- [ ] Better Error reporting
- [ ] Extract handlers from main
- [ ] Error handling clientside
