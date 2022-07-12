//import http from 'http';
import express from 'express';
import bodyParser from 'body-parser';
import path from 'path';
import { fileURLToPath } from 'url';

import { router as adminRoutes } from './routes.js';
import { router as shopRoutes } from './shop.js';
// import { requestHandler } from './routes.js';
// const server = http.createServer(requestHandler);

const __dirname = path.dirname(fileURLToPath(import.meta.url));

const app = express();

app.use(bodyParser.urlencoded({extended: false}));
//middleware from the body-parser package - parses data from form into key value pair
//app.use(express.static())

app.use(adminRoutes);
//automatically considers routes from routing file when running through middleware

// app.use('/add-product', (req, res, next) => {
//     res.send('<form action="/products" method="POST"><input type="text" name="name"><button type=submit>Add product</button></form>'
//     );
//     // next();
// });

// app.post('/products', (req, res, next) => {
//     console.log(req.body);
//     res.redirect('/');
//     // next();
// });


app.use(shopRoutes);
//Remember: order matters with the middleware, however shopRoutes uses a get request so will only listen for GETs

//app.use('/admin',shopRoutes);
//Can add a filter to usage of route (pull out of module) -- for example if all paths are in a subfolder

// app.use('/', (req, res, next) => {
//     res.send('<h1>You bet this is the homepage!</h1>');
// });

app.use((req, res, next) => {
    res.status(404).sendFile(path.join(__dirname, 'views', '404.html'));
});
//Catch-all 404 page, no route needed
//Can chain functions on res, as long as send() is last

app.listen(3000)

// const server = http.createServer(app);
// server.listen(3000);

export { __dirname };