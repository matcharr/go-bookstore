package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/matcharr/go-bookstore/pkg/utils"
	"github.com/matcharr/go-bookstore/pkg/models"
)

var NewBook models.Book