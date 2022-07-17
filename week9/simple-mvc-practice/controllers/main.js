//import model

import { __dirname, __filename } from '../app.js'

export let records = [];

export const getAddRecords = (req, res) => {
    res.render('add-record', {
        path: '/add-record',
        pageTitle: 'Add Record'
    });
};

export const postAddRecords = (req, res) => {
    records.push({ title: req.body.name });
    console.log(records);
    res.redirect('/');
};

export const getAllRecords = (req, res) => {
    res.render('all-records', {
        path: '/',
        pageTitle: 'All Records',
        recs: records
    });
};