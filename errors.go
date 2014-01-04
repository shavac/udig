package graph

import "errors"

var VertexNullError = errors.New("Vertex cannot be null.")
var VertexNotExistError = errors.New("Vertex not exist.")
var EdgeNotExistError = errors.New("Edge not exist.")
var ConnectSelfError = errors.New("Cannot connect to self.")
var AlreadyConnectedError = errors.New("two vertex alread connected.")
