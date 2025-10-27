# Customer labs recruitment task

## Ask
Design a server to accept a flat JSON, then pass it to a go routine via a channel for processing. After processing the data, send it to a webhook.

## Solution
- ``main.go`` contains the http server and the route handlers. The channel and the goroutine are also created in ``main()``.
- ``GET /`` endpoint returns a simple string to verify if the API is running or not
- ``POST /`` endpoint accepts a 1D JSON of ``string:string`` type. It converts it to a map and sends it to the channel for the goroutine to pick up. Status code 202 is returned to the client.
  - If the incoming payload is not of the expected type, 400 status code is returned.
  - If the channel is full, then a 503 status code is returned.
- The channel is a buffered channel with a capacity of 100.
- The goroutine ``dispatch()`` scans for new data in the channel. The keys are changed to more descriptive ones, the attributes and traits are made into a neat nested map (Identified using regex). Then the map is marshalled and sent to a webhook via POST request.
- Standard events such as endpoint hit, incoming and outgoing data are logged to stdout