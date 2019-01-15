// content of index.js
const http = require('http')
const puppeteer = require('puppeteer');
const port = 3000

const requestHandler = (request, response) => {
    let request_url = request.url.substring(1, request.url.length);
    console.log(request_url);
    if (request_url.startsWith("http")) {
        response.writeHead(200, { 'Content-Type': 'text/plain' });
        response.end(request_url);

    } else {
        response.writeHead(500, { 'Content-Type': 'text/plain' });
        response.end("error");
    }
}

// (async() => {
// 	const browser = await puppeteer.launch({headless:true});
// 	const page = await browser.newPage();
// 	await page.goto(process.argv[2]);
//     await page.waitFor(500);
// 	let content = await page.content();
// 	console.log(content);
// 	browser.close();
// })();

const server = http.createServer(requestHandler)

server.listen(port, (err) => {
    if (err) {
        return console.log('something bad happened', err)
    }

    console.log(`server is listening on ${port}`)
})