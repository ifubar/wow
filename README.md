# wow
### server of sacred knowledge protected with POW algorithm

## how to run
~~~ 
docker-compose up
~~~

## POW algorithm
- client get a task along with jwt token 
- client have to obtain the solution:
  - starting with `task.InputPrefix` 
  - has `task.MustContainLeadingZeros` leading zeros under applying `task.Alg` algorithm 
- send the solution back along with header `Jwt: {token}`
- ???
- PROFIT
