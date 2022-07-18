const Record = require('../models/record');

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
    