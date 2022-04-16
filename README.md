# upwdemoinfocs

Sample of [demoinfocs-golang](https://github.com/markus-wa/demoinfocs-golang) library implementation.

This app will print the player who got damaged because of falling.

## Prerequisite
- Download CSGO demo file [here](https://gitlab.com/markus-wa/cs-demos-2/-/blob/master/default.7z),
  extract and place it within the root of this project folder (`/default.dem`)
- Install all dependencies with `go mod download`


## How to run
> *Assuming that this app will be running on Unix environment.*

`go run cmd/upwdemoinfocs/main.go -demo ./default.dem`

The output will looks like this:
```
-- Players damaged because of fall --


user_id, name, last_alive_pos
233,"xms*ASUS ♥ /F/","(499.030609130859375000000000, 717.356018066406250000000000, 1630.094482421875000000000000)"
233,"xms*ASUS ♥ /F/","(501.463623046875000000000000, 653.456665039062500000000000, 1626.368164062500000000000000)"
255,"kzy LJ∼","(524.833068847656250000000000, 585.918273925781250000000000, 1624.883911132812500000000000)"
257,"xms*ASUS ♥ /F/","(-88.858825683593750000000000, -517.630310058593750000000000, 1621.294921875000000000000000)"
```