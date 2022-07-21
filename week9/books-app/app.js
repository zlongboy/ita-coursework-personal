const express = require('express');
const bodyParser = require('body-parser');

const app = express();
const port = 8080;

// FOR TESTING //
//console.clear();
// const getBooks = require('./integrations/books');
// (async function () {
//     try {
//         const result = await getBooks('michael-lewis')
//         console.log('awaited', result)
//     } catch (err) {
//         console.log(err);
//     }
// })();
// END //

app.set('view engine', 'pug');
app.set('views', 'views');
app.use(bodyParser.urlencoded({ extended: true }));
//Load public files here (e.g. CSS - video 75)

const addRoutes = require('./routes/add');
const displayRoutes = require('./routes/display');
const errorController = require('./controllers/errors');

app.use(addRoutes);
app.use(displayRoutes);
app.use(errorController.get404);

app.listen(port);

