# go-minesweeper
Minesweeper API made in Golang

## Important notes

* The architecture in this API was based on the article [Clean architecture in GO](https://medium.com/@hatajoe/clean-architecture-in-go-4030f11ec1b1)
* Did not put any test because of the time given
* Dockerize the API for easy uploading to the Cloud, the API was deployed to a EC2 instance inside it's own cluster with load balancer (even dough we dont really need a balancer, because i have only one instance running this, just did it for the fun of it)
* For the Click call, i've decided to not put it as a PUT because i didnt want to send all the game info in the body (what happens if i a have a 100x100 game? that would be a lot of info to pass), and i didnt put it as a PATCH because im not changing anything in the game, just a cell value (maybe i could've created a resource /cell and then just update that one cell, but for the simplicity of the API i did not persist Cell values  - only in the game -)
	
## Endpoints

| endpoint      | HTTP |description                       |
|:--------------|:----------------------------------|:----------------------------------|
| `/ping`      |GET|Returns a "pong" string to notice that server is still responding |
| `/games`    | GET|Returns a list of all Created Games |
| `/game/{id}`    | GET|Returns a Game or 404 if not found |
| `/game` |POST| Creates a new Game |
| `/game/{id}/click` |POST| Register a click on the game board |


### `/ping` resource
Check server is online
```
curl -X GET /ping -H 'Content-Type: application/json'
```

### `/games` resource
Retrieve a list of all saved Games
```
curl -X GET /games -H 'Content-Type: application/json' 
```
| Http Code      | description                       |
|:--------------|:----------------------------------|
| 200 | returns a list of saved Games |
| 500 | Server error |

### `/game/{id}` resource
Retrieve a saved Game
```
curl -X GET /game/{id} -H 'Content-Type: application/json' 
```
| Http Code      | description                       |
|:--------------|:----------------------------------|
| 200 | returns a GAME |
| 404 | Game not fond |
| 500 | Server error |

#### Response Body expected example
    {
        "id": 1,
        "rows": 3,
        "cols": 3,
        "mines": 1,
        "cellsRevealed": 0,
        "status": 0,
        "result": 2,
        "startTime": "2019-10-04T14:23:58.005756108Z",
        "grid": [...
                [{"mine": true,"clicked": false,"value": 0},...}]
                ...]
    }

### `/game` resource
Create a new game
```
curl -X POST /game -H 'Content-Type: application/json' -d '{"mines":1, "cols":3, "rows":3}'
```
| Http Code      | description                       |
|:--------------|:----------------------------------|
| 201 | returns a New GAME |
| 400 | Bad Request |
| 500 | Server error |

#### Request Body expected example
| Property      | description                       |
|:--------------|:----------------------------------|
| rows | quantity of rows to the game - min 1: max 100 |
| cols | quantity of cols to the game - min 1: max 100 |
| mines | quantity of mines in the game, can not be more than (rows*cols) |
    {
        "rows": 3,
        "cols": 3,
        "mines": 1,
    }


#### Response Body expected example
    {
        "id": 1,
        "rows": 3,
        "cols": 3,
        "mines": 1,
        "cellsRevealed": 0,
        "status": 0,
        "result": 2,
        "startTime": "2019-10-04T14:23:58.005756108Z",
        "grid": [...
                [{"mine": true,"clicked": false,"value": 0},...}]
                ...]
    }

### `/game/{id}/click` resource
Create a new game
```
curl -X POST /game{id}/click -H 'Content-Type: application/json' -d '{"col":3, "rows":1}'
```
| Http Code      | description                       |
|:--------------|:----------------------------------|
| 200 | return Game updated |
| 400 | Bad Request |
| 404 | Game not found |
| 500 | Server error |

#### Request Body expected example
| Property      | description                       |
|:--------------|:----------------------------------|
| row | number of row - it must be within the boundaries of the game, min:0 max:quantity of rows in game |
| col | number of col - it must be within the boundaries of the game, min:0 max:quantity of cols in game |
    {
        "row": 0,
        "col": 0,
    }


#### Response Body expected example
    {
        "id": 1,
        "rows": 3,
        "cols": 3,
        "mines": 1,
        "cellsRevealed": 0,
        "status": 0,
        "result": 2,
        "startTime": "2019-10-04T14:23:58.005756108Z",
        "grid": [...
                [{"mine": true,"clicked": false,"value": 0},...}]
                ...]
    }


### Game
| Property      | description                       |
|:--------------|:----------------------------------|
| id | id of the game |
| rows | quantity of rows to the game - min 1: max 100 |
| cols | quantity of cols to the game - min 1: max 100 |
| mines | quantity of mines in the game, can not be more than (rows*cols) |
| cellsRevealed | number of cells that were already clicked |
| status | [Status](###Status) of the game |
| result | [Result](###Result) of the game |
| startTime | date time the game was created |
| grid | matrix of [Cell](###Cell) for de board of the game |

### Cell 
| Property      | description                       |
|:--------------|:----------------------------------|
| Mine | bool indicating if there is a mine or not |
| Clicked | bool indicates if the cell was clicked|
| Value | quantity of mines that sorrounds this cell | 

### Status constant
| index      | description                       |
|:--------------|:----------------------------------|
| 0 | New Game |
| 1 | On going game|
| 2 | Paused Game |
| 3 | Finished game |

### Result constant
| index      | description                       |
|:--------------|:----------------------------------|
| 0 | Win |
| 1 | Lose|
| 2 | Undefined |
