package mongodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/diSpector/incrowd.git/internal/models/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDbStorage struct {
	client     *mongo.Client
	database   string
	collection string
}

func New(userName, password, host string, port int, database, collection string) (MongoDbStorage, error) {
	connStr := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authSource=%s", userName, password, host, port, database, database)
	client, err := mongo.Connect(options.Client().ApplyURI(connStr))
	if err != nil {
		return MongoDbStorage{}, err
	}

	return MongoDbStorage{
		client:     client,
		database:   database,
		collection: collection,
	}, nil
}

func (s MongoDbStorage) Close(ctx context.Context) error {
	return s.client.Disconnect(ctx)
}

func (s MongoDbStorage) GetArticles(ctx context.Context, limit, offset int64) ([]domain.Article, error) {
	collection := s.client.Database(s.database).Collection(s.collection)
	opts := options.Find()
	opts.SetLimit(limit)
	opts.SetSkip(offset)

	cursor, err := collection.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var articles []domain.Article
	for cursor.Next(ctx) {
		var article domain.Article
		if err := cursor.Decode(&article); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}

func (s MongoDbStorage) GetArticlesCount(ctx context.Context) (int64, error) {
	collection := s.client.Database(s.database).Collection(s.collection)
	count, err := collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s MongoDbStorage) GetArticleById(ctx context.Context, id string) (*domain.Article, error) {
	collection := s.client.Database(s.database).Collection(s.collection)
	var article domain.Article
	filter := bson.D{{Key: "id", Value: id}}
	err := collection.FindOne(ctx, filter).Decode(&article)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &article, nil
}

func (s MongoDbStorage) GetLastNIdsModTimeBySource(ctx context.Context, n int, source string) ([]domain.ArticleOriginMod, error) {
	coll := s.client.Database(s.database).Collection(s.collection)
	opts := options.Find().
		SetSort(bson.D{{Key: "_id", Value: -1}}).
		SetLimit(int64(n)).
		SetProjection(bson.D{{Key: "id", Value: 1}, {Key: "lastModified", Value: 1}, {Key: "source", Value: 1}})

	cursor, err := coll.Find(ctx, bson.D{{Key: "source.sourceSystem", Value: source}}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var articles []domain.ArticleOriginMod
	for cursor.Next(ctx) {
		var article domain.ArticleOriginMod
		if err := cursor.Decode(&article); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}

func (s MongoDbStorage) GetAtricleBySource(ctx context.Context, sourceId, sourceSystem string) (*domain.Article, error) {
	coll := s.client.Database(s.database).Collection(s.collection)
	var article domain.Article
	err := coll.FindOne(
		ctx,
		bson.D{{Key: "source.sourceId", Value: sourceId},
			{Key: "source.sourceSystem", Value: sourceSystem}},
		nil,
	).Decode(&article)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &article, nil
}

func (s MongoDbStorage) SaveArticle(ctx context.Context, article domain.Article) error {
	collection := s.client.Database(s.database).Collection(s.collection)
	_, err := collection.InsertOne(ctx, article)
	if err != nil {
		return err
	}

	return nil
}

func (s MongoDbStorage) ReplaceArticleById(ctx context.Context, id string, article domain.Article) error {
	collection := s.client.Database(s.database).Collection(s.collection)
	filter := bson.M{"id": id}

	_, err := collection.ReplaceOne(ctx, filter, article)
	if err != nil {
		return err
	}

	return nil
}
