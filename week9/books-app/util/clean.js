//two functions clean.author(), clean.volumes()

exports.author = (input) => {
    return input.toLowerCase().replaceAll(' ', '-');
};

exports.volumes = (volumes) => {
    volumes.forEach(e => {
        //TODO loop through array and extract volumes. Call this function in the model to assign.
    });
};
