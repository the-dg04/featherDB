### A simple and lightweight JSON server written in Go

#### Installation
Clone this repo
```bash
git clone https://github.com/the-dg04/featherDB.git
cd featherDB
```
Install dependencies ```github.com/gin-gonic/gin``` and ```github.com/google/uuid```

#### Run shit

specify ```PORT``` in ```server.go``` to run on a custom port (default: 6969)

```bash
go run .
```

#### API reference
- ```/api/record/:id``` [GET] Get record by id
- ```/api/new``` [POST] Create new record [send record as json in body] [returns ```_id``` of created record] 
- ```/api/filter``` [POST] Filter records [send filter as json in body] [returns filtered records] 
- ```/api/patch/:id``` [PATCH] Update record by id [send updated fields as json in body] 
- ```/api/delete/:id``` [DELETE] Delete record by id
- ```/api/clear``` [DELETE] Delete the fucking database
