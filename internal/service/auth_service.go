package service

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/bcrypt"

	"dubai-auto/internal/config"
	"dubai-auto/internal/model"
	"dubai-auto/internal/repository"
	"dubai-auto/internal/utils"
	"dubai-auto/pkg/auth"
	"dubai-auto/pkg/files"
)

type AuthService struct {
	repo   *repository.AuthRepository
	config *config.Config
}

func NewAuthService(repo *repository.AuthRepository, config *config.Config) *AuthService {
	return &AuthService{repo, config}
}

func (s *AuthService) UserRegisterDevice(ctx *fasthttp.RequestCtx, userID int, req model.UserRegisterDevice) model.Response {

	err := s.repo.UserRegisterDevice(ctx, userID, req)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return model.Response{Data: model.Success{Message: "Device registered successfully"}}
}

func (s *AuthService) Application(ctx *fasthttp.RequestCtx, req model.UserApplication) model.Response {
	u, err := s.repo.Application(ctx, req)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	accessToken, refreshToken := auth.CreateRefreshAccsessToken(u.ID, req.RoleID)
	return model.Response{
		Data: model.LoginFiberResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}

func (s *AuthService) ApplicationDocuments(ctx *fasthttp.RequestCtx, id int, licence, memorandum, copyOfID *multipart.FileHeader) model.Response {
	documents := model.UserApplicationDocuments{}
	ext := strings.ToLower(filepath.Ext(licence.Filename))

	if ext != ".pdf" {
		return model.Response{Error: errors.New("only PDF files are allowed"), Status: http.StatusBadRequest}
	}

	if !utils.IsPDF(licence) {
		return model.Response{Error: errors.New("file is not a valid PDF"), Status: http.StatusBadRequest}
	}

	path, err := files.SaveOriginal(licence, config.ENV.STATIC_PATH+"documents/"+strconv.Itoa(id))

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	documents.Licence = path
	ext = strings.ToLower(filepath.Ext(memorandum.Filename))

	if ext != ".pdf" {
		return model.Response{Error: errors.New("only PDF files are allowed"), Status: http.StatusBadRequest}
	}

	if !utils.IsPDF(memorandum) {
		return model.Response{Error: errors.New("file is not a valid PDF"), Status: http.StatusBadRequest}
	}

	path, err = files.SaveOriginal(memorandum, config.ENV.STATIC_PATH+"documents/"+strconv.Itoa(id))

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	documents.Memorandum = path
	ext = strings.ToLower(filepath.Ext(copyOfID.Filename))

	if ext != ".pdf" {
		return model.Response{Error: errors.New("only PDF files are allowed"), Status: http.StatusBadRequest}
	}

	if !utils.IsPDF(copyOfID) {
		return model.Response{Error: errors.New("file is not a valid PDF"), Status: http.StatusBadRequest}
	}

	path, err = files.SaveOriginal(copyOfID, config.ENV.STATIC_PATH+"documents/"+strconv.Itoa(id))

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	documents.CopyOfID = path
	err = s.repo.ApplicationDocuments(ctx, id, documents)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return model.Response{Data: model.Success{Message: "Application documents sent successfully"}}
}

func (s *AuthService) UserLoginGoogle(ctx *fasthttp.RequestCtx, tokenID string) model.Response {
	req, _ := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	req.Header.Set("Authorization", "Bearer "+tokenID)
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusBadRequest}
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Response{Error: errors.New("failed to get user info"), Status: http.StatusBadRequest}
	}

	var userInfo model.GoogleUserInfo

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return model.Response{Error: err, Status: http.StatusBadRequest}
	}

	u, err := s.repo.UserLoginGoogle(ctx, userInfo)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	accessToken, refreshToken := auth.CreateRefreshAccsessToken(u.ID, 1)
	return model.Response{
		Data: model.LoginFiberResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}

func (s *AuthService) DeleteAccount(ctx *fasthttp.RequestCtx, userID int) model.Response {
	err := s.repo.DeleteAccount(ctx, userID)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	// todo: delete associated files
	return model.Response{Data: model.Success{Message: "Account deleted successfully"}}
}

func (s *AuthService) UserEmailConfirmation(ctx *fasthttp.RequestCtx, user *model.UserEmailConfirmationRequest) model.Response {
	u, err := s.repo.TempUserByEmail(ctx, &user.Email)

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusNotFound,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.OTP), []byte(user.OTP))

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusBadRequest,
		}
	}

	u.ID, err = s.repo.UserEmailGetOrRegister(ctx, u.Username, user.Email, u.OTP)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	accessToken, refreshToken := auth.CreateRefreshAccsessToken(u.ID, 1)

	return model.Response{
		Data: model.LoginFiberResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}

func (s *AuthService) UserPhoneConfirmation(ctx *fasthttp.RequestCtx, user *model.UserPhoneConfirmationRequest) model.Response {
	u, err := s.repo.TempUserByPhone(ctx, &user.Phone)

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusNotFound,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.OTP), []byte(user.OTP))

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusBadRequest,
		}
	}

	u.ID, err = s.repo.UserPhoneGetOrRegister(ctx, u.Username, user.Phone, u.OTP)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	accessToken, refreshToken := auth.CreateRefreshAccsessToken(u.ID, 1)
	return model.Response{
		Data: model.LoginFiberResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}

func (s *AuthService) UserLoginEmail(ctx *fasthttp.RequestCtx, user *model.UserLoginEmail) model.Response {
	otp := utils.RandomOTP()
	username := ""
	// for google play store testing
	if user.Email == "berdalyyew99@gmail.com" {
		otp = 123456
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("%d", otp)), bcrypt.DefaultCost)

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusInternalServerError,
		}
	}

	err = s.repo.TempUserEmailGetOrRegister(ctx, username, user.Email, string(hashedPassword))

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusInternalServerError,
		}
	}

	err = utils.SendEmail("OTP", fmt.Sprintf("Your OTP is: %d", otp), user.Email)

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusInternalServerError,
		}
	}

	return model.Response{
		Data: model.Success{Message: "Successfully created the user"},
	}
}

func (s *AuthService) UserForgetPassword(ctx *fasthttp.RequestCtx, user *model.UserForgetPasswordReq) model.Response {
	u, err := s.repo.UserByEmail(ctx, &user.Email)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusNotFound}
	}

	otp := utils.RandomOTP()
	otpHash, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("%d", otp)), bcrypt.DefaultCost)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	err = utils.SendEmail("Password Reset", fmt.Sprintf("Your new password is: %d", otp), user.Email)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	err = s.repo.UpdateUserTempPassword(ctx, u.ID, string(otpHash))

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return model.Response{Data: model.Success{Message: "Confirmation code sent successfully"}}
}

func (s *AuthService) UserResetPassword(ctx *fasthttp.RequestCtx, user *model.UserResetPasswordReq) model.Response {
	u, err := s.repo.UserByEmail(ctx, &user.Email)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusNotFound}
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.OTP), []byte(user.OTP))

	if err != nil {
		return model.Response{Error: err, Status: http.StatusBadRequest}
	}

	newPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	err = s.repo.UpdateUserPassword(ctx, u.ID, string(newPassword))

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return model.Response{Data: model.Success{Message: "New password reset successfully"}}
}

func (s *AuthService) ThirdPartyLogin(ctx *fasthttp.RequestCtx, user *model.ThirdPartyLoginReq) model.Response {
	u, err := s.repo.ThirdPartyLogin(ctx, user.Email)

	// this is for appStore and playStore testing
	if user.Email == "danisultan2021@gmail.com" && user.Password == "123456" {
		accessToken, refreshToken := auth.CreateRefreshAccsessToken(u.ID, u.RoleID)

		return model.Response{
			Data: model.ThirdPartyLoginFiberResponse{
				AccessToken:    accessToken,
				RefreshToken:   refreshToken,
				FirstTimeLogin: u.FirstTimeLogin,
			},
		}
	}

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusNotFound,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusBadRequest,
		}
	}

	accessToken, refreshToken := auth.CreateRefreshAccsessToken(u.ID, u.RoleID)
	return model.Response{
		Data: model.ThirdPartyLoginFiberResponse{
			AccessToken:    accessToken,
			RefreshToken:   refreshToken,
			FirstTimeLogin: u.FirstTimeLogin,
		},
	}
}

func (s *AuthService) UserLoginPhone(ctx *fasthttp.RequestCtx, user *model.UserLoginPhone) model.Response {
	otp := utils.RandomOTP()
	username := ""
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("%d", otp)), bcrypt.DefaultCost)

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusInternalServerError,
		}
	}

	err = s.repo.TempUserPhoneGetOrRegister(ctx, username, user.Phone, string(hashedPassword))

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusInternalServerError,
		}
	}

	err = utils.SendOtp(user.Phone, otp)

	if err != nil {
		// Log error but don't fail the request - OTP sending failure shouldn't block user creation
		log.Printf("Warning: Failed to send OTP to %s: %v", user.Phone, err)
	}

	return model.Response{
		Data: model.Success{Message: "Successfully created the user."},
	}
}

func (s *AuthService) AdminLogin(ctx *fasthttp.RequestCtx, userReq *model.AdminLoginReq) model.Response {
	u, err := s.repo.AdminLogin(ctx, userReq.Email)

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusNotFound,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(userReq.Password))

	if err != nil {
		return model.Response{
			Error:  err,
			Status: http.StatusBadRequest,
		}
	}

	accessToken, refreshToken := auth.CreateRefreshAccsessToken(u.ID, auth.ADMIN_ROLE)
	return model.Response{
		Data: model.LoginFiberResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}

func (s *AuthService) UserLoginApple(ctx *fasthttp.RequestCtx, authorizationCode string) model.Response {
	// Step 1: Generate client secret using .p8 file
	clientSecret, err := utils.GenerateAppleClientSecret(
		s.config.APPLE_TEAM_ID,
		s.config.APPLE_KEY_ID,
		s.config.APPLE_CLIENT_ID,
		s.config.APPLE_KEY_PATH,
	)
	if err != nil {
		return model.Response{Error: fmt.Errorf("failed to generate Apple client secret: %w", err), Status: http.StatusInternalServerError}
	}

	// Step 2: Exchange authorization code for access token
	tokenData := url.Values{}
	tokenData.Set("client_id", s.config.APPLE_CLIENT_ID)
	tokenData.Set("client_secret", clientSecret)
	tokenData.Set("code", authorizationCode)
	tokenData.Set("grant_type", "authorization_code")

	tokenReq, err := http.NewRequest("POST", "https://appleid.apple.com/auth/token", strings.NewReader(tokenData.Encode()))
	if err != nil {
		return model.Response{Error: err, Status: http.StatusBadRequest}
	}
	tokenReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	tokenResp, err := client.Do(tokenReq)
	if err != nil {
		return model.Response{Error: err, Status: http.StatusBadRequest}
	}
	defer tokenResp.Body.Close()

	if tokenResp.StatusCode != http.StatusOK {
		bodyBytes := make([]byte, 1024)
		tokenResp.Body.Read(bodyBytes)
		return model.Response{Error: fmt.Errorf("failed to exchange code for token: %s", string(bodyBytes)), Status: http.StatusBadRequest}
	}

	var tokenResponse struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		IDToken      string `json:"id_token"`
	}

	if err := json.NewDecoder(tokenResp.Body).Decode(&tokenResponse); err != nil {
		return model.Response{Error: err, Status: http.StatusBadRequest}
	}

	// Step 3: Decode ID token to get user info (Apple provides user info in ID token)
	// The ID token is a JWT that contains user information
	// We'll decode it without verification for now (in production, verify with Apple's public keys)
	parts := strings.Split(tokenResponse.IDToken, ".")
	if len(parts) != 3 {
		return model.Response{Error: errors.New("invalid ID token format"), Status: http.StatusBadRequest}
	}

	// Decode base64 URL encoded payload (second part)
	// Note: In production, you should verify the JWT signature using Apple's public keys
	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return model.Response{Error: fmt.Errorf("failed to decode ID token payload: %w", err), Status: http.StatusBadRequest}
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return model.Response{Error: fmt.Errorf("failed to parse ID token payload: %w", err), Status: http.StatusBadRequest}
	}

	var userInfo model.AppleUserInfo
	if sub, ok := payload["sub"].(string); ok {
		userInfo.Sub = sub
	}
	if email, ok := payload["email"].(string); ok {
		userInfo.Email = email
	}

	// If email is not in ID token, try to get it from userinfo endpoint
	if userInfo.Email == "" {
		userInfoReq, err := http.NewRequest("GET", "https://appleid.apple.com/auth/userinfo", nil)
		if err == nil {
			userInfoReq.Header.Set("Authorization", "Bearer "+tokenResponse.AccessToken)
			userInfoResp, err := client.Do(userInfoReq)
			if err == nil && userInfoResp.StatusCode == http.StatusOK {
				var userInfoFromAPI model.AppleUserInfo
				if err := json.NewDecoder(userInfoResp.Body).Decode(&userInfoFromAPI); err == nil {
					if userInfoFromAPI.Email != "" {
						userInfo.Email = userInfoFromAPI.Email
					}
				}
				userInfoResp.Body.Close()
			}
		}
	}

	u, err := s.repo.UserLoginApple(ctx, userInfo)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	accessToken, refreshToken := auth.CreateRefreshAccsessToken(u.ID, 1)

	return model.Response{
		Data: model.LoginFiberResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}
