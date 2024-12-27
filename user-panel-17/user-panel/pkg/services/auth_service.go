package services

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"user-panel/pkg/cache"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var validate = validator.New()
var jwtSecret = []byte("your-secret-key")

// HashPassword hashes the given password using bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyPassword compares a plain password with a hashed password
func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// GenerateJWT generates a JWT with an expiration time
func GenerateJWT(userID, role string) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"role":   role,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // 24 hours expiration
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateJWT validates a JWT and checks its expiration
func ValidateJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return nil, errors.New("token expired")
		}
		return token, nil
	}
	return nil, errors.New("invalid token")
}

// LogoutUser handles user logout by blacklisting the token
func LogoutUser(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}

	// Add the token to the blacklist
	err := cache.RedisClient.Set(cache.ctx, "blacklist:"+tokenString, true, time.Hour*24).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to blacklist token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user logged out successfully"})
}

// IsTokenBlacklisted checks if a token is blacklisted
func IsTokenBlacklisted(tokenString string) bool {
	_, err := cache.RedisClient.Get(cache.ctx, "blacklist:"+tokenString).Result()
	return err == nil
}

// GenerateOTP generates a random OTP and stores it in Redis
func GenerateOTP(userID string) (string, error) {
	otp := strconv.Itoa(rand.Intn(1000000))
	err := cache.RedisClient.Set(cache.ctx, "otp:"+userID, otp, time.Minute*5).Err()
	if err != nil {
		return "", err
	}
	return otp, nil
}

// VerifyOTP verifies the OTP entered by the user
func VerifyOTP(userID, otp string) bool {
	storedOTP, err := cache.RedisClient.Get(cache.ctx, "otp:"+userID).Result()
	return err == nil && storedOTP == otp
}

// RegisterInput represents user registration input
type RegisterInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// RegisterUser handles user registration
func RegisterUser(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// Validate input
	if err := validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	// Store the user in the database
	_, err = db.DB.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", input.Email, hashedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
}


// Policy represents a role-based access control policy
type Policy struct {
	Resource string
	Action   string
	Role     string
}

// AccessControlList defines the policies for different roles
var AccessControlList = []Policy{
	{"wallet", "view", "USER"},
	{"wallet", "update", "ADMIN"},
	{"refund", "approve", "ADMIN"},
	{"user", "delete", "SUPER_ADMIN"},
}

// IsAuthorized checks if a role has permission to perform an action on a resource
func IsAuthorized(role, resource, action string) bool {
	for _, policy := range AccessControlList {
		if policy.Resource == resource && policy.Action == action && policy.Role == role {
			return true
		}
	}
	return false
}

// LogAccess logs access attempts to the database
func LogAccess(userID, resource, action, status string) {
	_, err := db.DB.Exec("INSERT INTO access_logs (user_id, resource, action, status) VALUES ($1, $2, $3, $4)",
		userID, resource, action, status)
	if err != nil {
		log.Printf("Failed to log access: %v", err)
	}
}


var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path"},
	)
)

func init() {
	// Register custom metrics
	prometheus.MustRegister(httpRequestsTotal)
}

func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		httpRequestsTotal.WithLabelValues(c.Request.Method, c.FullPath()).Inc()
		c.Next()
	}
}
