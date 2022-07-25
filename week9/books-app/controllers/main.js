const clean = require('../util/clean');
const data = require('../util/data')
const getBooks = require('../integrations/books');

exports.getSearchBooks = (req, res) => {
    res.render('search-book', {
        pageTitle: 'Search Books',
    })
}

exports.postSearchBooks = (req, res) => {
    (async function () {
            res.render('book-results', {
            pageTitle: 'Book Results',
            recs: await data.fetchBook(req.body.id)
        })
    })();   
}

exports.getAuthor = (req, res) => {
    res.render('add-author', {
        pageTitle: 'Add Author'
    });
};

exports.postAuthor = (req, res, next) => {
    (async function () {
        res.render('author-results', {
            pageTitle: 'All Results',
            recs: clean.volumes(await getBooks(clean.author(req.body.name)))
            //TODO: How can I speed up this render?
        });
    })();
};
