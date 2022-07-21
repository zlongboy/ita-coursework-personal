// const db = require('../util/records-schema');

// module.exports = class Record {
//     constructor(name) {
//         this.title = name; 
//     }

//     save() {
//         return db.execute('INSERT INTO records (title) VALUES (?)', [this.title]);
//     }

//     static fetchRecords() {
//         return db.execute('SELECT * FROM records')
//     }
// };