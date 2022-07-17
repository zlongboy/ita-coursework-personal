//import model
import { __dirname, __filename } from '../app.js'; 
import { Record } from '../models/record.js';

export const getAddRecords = (req, res) => {
    res.render('add-record', {
        pageTitle: 'Add Record'
    });
};

export const postAddRecords = (req, res) => {
    const record = new Record(req.body.name);
    record.save();
    res.redirect('/');
};

export const getAllRecords = (req, res) => {
    Record.fetchRecords(records => {
        res.render('all-records', {
            pageTitle: 'All Records',
            recs: records
        });
    });
}; 
    