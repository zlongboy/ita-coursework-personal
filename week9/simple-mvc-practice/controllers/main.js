//import model
const Record = require('../models/record');

exports.getAddRecords = (req, res) => {
    res.render('add-record', {
        pageTitle: 'Add Record'
    });
};

exports.postAddRecords = (req, res) => {
    const record = new Record(req.body.name);
    record.save();
    res.redirect('/');
};

exports.getAllRecords = (req, res) => {
    Record.fetchRecords(records => {
        res.render('all-records', {
            pageTitle: 'All Records',
            recs: records
        });
    });
}; 
    