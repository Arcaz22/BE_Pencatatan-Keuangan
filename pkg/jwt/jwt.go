package jwt

import (
    "errors"
    "time"
    "github.com/golang-jwt/jwt/v5"
    "github.com/google/uuid"
    "github.com/spf13/viper"
)

func init() {
    viper.SetConfigFile(".env")
    viper.AutomaticEnv()

    err := viper.ReadInConfig()
    if err != nil {
        panic("Failed to load environment variables")
    }
}

func GenerateToken(userID uuid.UUID) (string, error) {
    secretKey := []byte(viper.GetString("JWT_SECRET_KEY"))

    claims := jwt.MapClaims{
        "user_id": userID.String(),
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) (uuid.UUID, error) {
    secretKey := []byte(viper.GetString("JWT_SECRET_KEY"))

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })

    if err != nil {
        return uuid.Nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        userIDStr, ok := claims["user_id"].(string)
        if !ok {
            return uuid.Nil, errors.New("invalid user ID in token")
        }

        userID, err := uuid.Parse(userIDStr)
        if err != nil {
            return uuid.Nil, errors.New("invalid UUID format in token")
        }

        return userID, nil
    }

    return uuid.Nil, errors.New("invalid token")
}
