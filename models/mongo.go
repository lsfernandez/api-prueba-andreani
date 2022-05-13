package models

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var err error
var transacciones *mongo.Collection

func init() {
	mongoURL := os.Getenv("MONGO_CONNECTION")
	mongoDB := os.Getenv("MONGO_DB")
	mongoCollection := os.Getenv("MONGO_COLLECTION")
	clientOptions := options.Client().ApplyURI(mongoURL)
	mongoClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Error,NO se conectó a MongoDB", err)
	} else {
		db := mongoClient.Database(mongoDB)
		transacciones = db.Collection(mongoCollection)
		fmt.Println("Conectado con éxito a", db.Name())
	}
}

func GuardarPedido(pedidoCorrecto Pedido) (primitive.ObjectID, error) {
	result, err := transacciones.InsertOne(context.TODO(), pedidoCorrecto)
	insertID := result.InsertedID.(primitive.ObjectID)
	return insertID, err
}

func ConsultarPedido(idConsulta string) (ResponsePedido, error) {
	objID, _ := primitive.ObjectIDFromHex(idConsulta)
	var consultaMapeo ResponsePedido
	result := transacciones.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&consultaMapeo)
	consultaMapeo.IdPedido = objID.Hex()
	return consultaMapeo, result
}

func ModificarPedido(idConsulta string, datosParaActualizar ResponsePedido) error {
	objID, _ := primitive.ObjectIDFromHex(idConsulta)
	update := bson.M{"nombrecompleto": datosParaActualizar.NombreCompleto, "color": datosParaActualizar.Color, "coloringles": datosParaActualizar.ColorIngles, "fechapedido": datosParaActualizar.FechaPedido}
	_, err := transacciones.UpdateOne(context.TODO(), bson.M{"_id": objID}, bson.M{"$set": update})
	return err
}
