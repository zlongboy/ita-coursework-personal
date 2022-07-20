const axios = require('axios');

const googleBooks = require('../config');
const baseURL = 'https://www.googleapis.com/books/v1';
const volumes = '/volumes'  

module.exports = async function getBooks(author) {
    try {
        const response = await axios(baseURL + volumes, {
            method: 'get',
            params: {
                key: googleBooks.API_KEY,
                q: `inauthor:${author}`, 
                projection: 'lite',
                orderBy: 'newest'
            },
            timeout: 15000
        });
        
        //console.log('From function...');
        // console.log(response.status);
        //console.log(response.data.items[0].id);
        return response.data.items;
    } catch (err) {
        console.error(err);
    }
};

