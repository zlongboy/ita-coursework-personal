const Record = require('../models/record');
const clean = require('../util/clean');
const getBooks = require('../integrations/books')

// RECORDS //
exports.getAddRecords = (req, res) => {
    res.render('add-record', {
        pageTitle: 'Add Record'
    });
};

exports.postAddRecords = (req, res) => {
    const record = new Record(req.body.name);
    record
        .save()
        .then(() => {
            res.redirect('/');
        })
        .catch(err => console.log(err))
};

exports.getAllRecords = (req, res) => {
    Record.fetchRecords()
        .then(([rows]) => {
            res.render('all-records', {
                pageTitle: 'All Records',
                recs: rows
            });
        })
        .catch(err => console.log(err))
}; 


// BOOKS //
exports.getAuthor = (req, res) => {
    res.render('add-book', {
        pageTitle: 'Add Book'
    });
};

exports.postAuthor = (req, res, next) => {
    getBooks(clean.author(req.body.name));
    res.redirect('/add-book');
};

