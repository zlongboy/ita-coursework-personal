const clean = require('../util/clean');
const data = require('../util/data')
const getBooks = require('../integrations/books');

exports.getSearchBooks = (req, res) => {
    res.render('search-book', {
        pageTitle: 'Search Books',
    })
}

exports.postSearchBooks = (req, res) => {
    const result = (async function () {
        return data.fetchBook(req.body.id);
    })();
    console.log(result);
    res.render('book-results', {
        pageTitle: 'Book Results',
        //TODO: result not resolved yet
        recs: result
    })
}

exports.getAuthor = (req, res) => {
    res.render('add-author', {
        pageTitle: 'Add Author'
    });
};

exports.postAuthor = (req, res, next) => {
    const added = (async function () {
        return clean.volumes(await getBooks(clean.author(req.body.name)));
    })();

    res.render('/author-results', {
        pageTitle: 'All Results',
        //TODO: added not resolved yet
        recs: added
    });
};
