### Training tracker
Super simple activity tracker.
1. You can add activity through POST request
2. Posted activity is added to queue
3. Then it's being consumed and saved into DB
4. Then you can fetch specific training by it's id or userId

#### DB
```sql
CREATE TABLE trainings (
  id SERIAL PRIMARY KEY,
  user_id VARCHAR NOT NULL,
  title VARCHAR NOT NULL,
  "type" VARCHAR NOT NULL,
  distance INT NOT NULL,
  "time" INT NOT NULL,
  "date" timestamptz NOT NULL
);
```

#### Requests
1. Add training
```
curl --location --request POST 'http://localhost:1323/training' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userId": "fgw2f2",
    "title": "My run",
    "type": "run",
    "distance": 2131,
    "time": 1231,
    "date": "2022-04-11T15:58:01.511Z"
}'
```
2. Fetch by id `curl --location --request GET 'http://localhost:1323/training/1243'`
3. Fetch by userId `curl --location --request GET 'http://localhost:1323/training/user/1eoineo1'`

#### additional info
- all need services are in docker-compose
- I didn't add go service in docker-compose as I was running it locally
- all need env variables are in .env.dist

