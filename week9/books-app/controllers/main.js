const Book = require('../models/record');
const clean = require('../util/clean');

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

exports.searchAuthor = (req, res) => {
    //TODO: req and/or its properties is not being resolved before creating variable -- videos?
    console.log(req.params.name);
    const cleanAuthor = clean.author(req.params.name)
    console.log(cleanAuthor);
    // .then(() => {
    //     res.redirect('/add-record');
    // })
}

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
    