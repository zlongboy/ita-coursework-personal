const axios = require('axios');

const config = require('../config');
const baseURL = 'https://www.googleapis.com/books/v1';
const endpoint = '/volumes'  

module.exports = async function getBooks(author) {
    try {
        const response = await axios(baseURL + endpoint, {
            method: 'get',
            params: {
                key: config.API_KEY,
                q: `inauthor:${author}`, 
            },
            timeout: 15000
        });
        //console.log(response.status)
        return response.data.items;
    } catch (err) {
        console.error(err);
    }
};

