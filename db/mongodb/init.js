db.createUser({
    user: _getEnv("MONGO_USER"),
    pwd: _getEnv("MONGO_PASSWORD"),
    roles: [{ role: "readWrite", db: _getEnv("MONGO_INITDB_DATABASE") }],
});

db.createCollection('articles');

db.articles.createIndex({ id: "text" })
db.articles.createIndex({ "source.sourceSystem": "text" })