package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/google/uuid"

	"microblog/internal/auth"
	"microblog/internal/database"
	"microblog/internal/request"
	"microblog/internal/response"
)

func (app *application) status(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status": "OK",
	}

	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// authToken handles user login (POST /auth/token)
func (app *application) authToken(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Parse JSON request
	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	// Get user by username
	user, err := app.queries.GetUserByUsername(context.Background(), input.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			app.errorMessage(w, r, http.StatusUnauthorized, "invalid username or password", nil)
			return
		}
		app.serverError(w, r, err)
		return
	}

	// Compare password
	if !user.Password.Valid || !auth.ComparePassword(user.Password.String, input.Password) {
		app.errorMessage(w, r, http.StatusUnauthorized, "invalid username or password", nil)
		return
	}

	// Generate JWT token
	token, err := auth.GenerateToken(user.ID, app.config.jwtSecret)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Return token
	data := map[string]map[string]string{
		"data": {
			"token": token,
		},
	}

	err = response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

// createUser handles user registration (POST /users)
func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Parse JSON request
	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	// Hash password
	hashedPassword, err := auth.HashPassword(input.Password)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Generate unique user ID
	userID := uuid.New().String()

	// Create user in database
	err = app.queries.CreateUser(context.Background(), database.CreateUserParams{
		ID:       userID,
		Username: input.Username,
		Email:    sql.NullString{String: input.Email, Valid: input.Email != ""},
		Password: sql.NullString{String: hashedPassword, Valid: true},
	})
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Generate JWT token for auto-login
	token, err := auth.GenerateToken(userID, app.config.jwtSecret)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Return token
	data := map[string]map[string]string{
		"data": {
			"token": token,
		},
	}

	err = response.JSON(w, http.StatusCreated, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}
