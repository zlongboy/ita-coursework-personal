//import model

import { __dirname, __filename } from '../app.js'

export const getAddRecords = (req, res) => {
    res.render('add-record', {
        path: '/add-record'
    });
};

export const postAddRecords = (req, res) => {
    res.redirect('/');
};

export const getAllRecords = (req, res) => {
    res.render('all-records', {
        path: '/'
    });
};