// content of index.js
const http = require('http')
const port = 3000

const requestHandler = (request, response) => {
  if (request.url[1:].startsWith(http)) {
    response.writeHead(200,{'Content-Type': 'text/plain'});
    response.end(request.url);

  } else {
    response.writeHead(500,{'Content-Type': 'text/plain'});
    response.end("error");

  }
  console.log(request.url)
}

const server = http.createServer(requestHandler)

server.listen(port, (err) => {
  if (err) {
    return console.log('something bad happened', err)
  }

  console.log(`server is listening on ${port}`)
})