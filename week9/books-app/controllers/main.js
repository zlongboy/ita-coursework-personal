const Book = require('../models/record');
const clean = require('../util/clean');

const bodyParser = require('body-parser');

app.use(bodyParser.urlencoded({ extended: true }));

exports.getAddRecords = (req, res) => {
    res.render('add-record', {
        pageTitle: 'Add Record'
    });
};

exports.postAddRecords = (req, res) => {
    const record = new Book(req.body.name);
    record
        .save()
        .then(() => {
            res.redirect('/');
        })
        .catch(err => console.log(err))
};

exports.getSearchAuthor = (req, res) => {
    res.render('add-book', {
        pageTitle: 'Add Book'
    });
};

exports.postSearchAuthor = (req, res, next) => {
    console.log(req.body.name);
    next();
    console.log(req.body.name);
};

exports.getAllRecords = (req, res) => {
    Book.fetchRecords()
        .then(([rows]) => {
            res.render('all-records', {
                pageTitle: 'All Records',
                recs: rows
            });
        })
        .catch(err => console.log(err))
}; 
    