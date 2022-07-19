const axios = require('axios');

const baseURL = 'https://www.googleapis.com/books/v1';
const volumes = '/volumes'  

module.exports = async function getBooks(author) {
    try {
        const response = await axios(baseURL + volumes, {
            method: 'get',
            params: {
                key: 'AIzaSyDe0FuNsS4t3jFkroYofXH1_67OsVQ7yEQ',
                q: `inauthor:${author}`, 
                projection: 'lite',
                orderBy: 'newest'
            },
            timeout: 15000
        });

        console.log(response.status);
    } catch (err) {
        console.error(err);
    }
};

