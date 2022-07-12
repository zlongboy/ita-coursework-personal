import express from 'express';
import path from 'path';

import { router as homeRoutes } from './routes/home.js';
import { rootDir } from './util/path.js';
import { router as usersRoutes } from './routes/users.js';

const port = 8080;
const app = express();
const __dirname = path.dirname(rootDir)

app.use(express.static(path.join(__dirname, 'public')));
app.use(homeRoutes);
app.use(usersRoutes);
app.use((req, res) => {
    res.status(404).sendFile(path.join(__dirname,'views', '404.html'));
});

app.listen(port);