const Book = require('../models/record');
const clean = require('../util/clean');
const getBooks = require('../integrations/books')

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

exports.getAuthor = (req, res) => {
    res.render('add-book', {
        pageTitle: 'Add Book'
    });
};

exports.postAuthor = (req, res, next) => {
    getBooks(clean.author(req.body.name));
    res.redirect('/add-book');
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
    