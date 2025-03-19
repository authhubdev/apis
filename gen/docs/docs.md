# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [auth/v1/common.proto](#auth_v1_common-proto)
    - [AuthResponse](#auth-v1-AuthResponse)
    - [AuthTokens](#auth-v1-AuthTokens)
    - [SuccessResponse](#auth-v1-SuccessResponse)
    - [UserInfo](#auth-v1-UserInfo)
  
- [auth/v1/google_oauth.proto](#auth_v1_google_oauth-proto)
    - [ExchangeAuthCodeRequest](#auth-v1-ExchangeAuthCodeRequest)
    - [ExchangeAuthCodeResponse](#auth-v1-ExchangeAuthCodeResponse)
    - [GetAuthURLRequest](#auth-v1-GetAuthURLRequest)
    - [GetAuthURLResponse](#auth-v1-GetAuthURLResponse)
  
    - [AuthGoogleOAuthService](#auth-v1-AuthGoogleOAuthService)
  
- [auth/v1/magic_link.proto](#auth_v1_magic_link-proto)
    - [SendMagicLinkRequest](#auth-v1-SendMagicLinkRequest)
    - [SendMagicLinkResponse](#auth-v1-SendMagicLinkResponse)
    - [VerifyMagicLinkRequest](#auth-v1-VerifyMagicLinkRequest)
    - [VerifyMagicLinkResponse](#auth-v1-VerifyMagicLinkResponse)
  
    - [AuthMagicLinkService](#auth-v1-AuthMagicLinkService)
  
- [auth/v1/password.proto](#auth_v1_password-proto)
    - [ForgotPasswordRequest](#auth-v1-ForgotPasswordRequest)
    - [ForgotPasswordResponse](#auth-v1-ForgotPasswordResponse)
    - [LoginRequest](#auth-v1-LoginRequest)
    - [LoginResponse](#auth-v1-LoginResponse)
    - [RegisterRequest](#auth-v1-RegisterRequest)
    - [RegisterResponse](#auth-v1-RegisterResponse)
    - [ResetPasswordRequest](#auth-v1-ResetPasswordRequest)
    - [ResetPasswordResponse](#auth-v1-ResetPasswordResponse)
  
    - [AuthPasswordService](#auth-v1-AuthPasswordService)
  
- [auth/v1/session.proto](#auth_v1_session-proto)
    - [GetSessionRequest](#auth-v1-GetSessionRequest)
    - [GetSessionResponse](#auth-v1-GetSessionResponse)
    - [ListSessionsRequest](#auth-v1-ListSessionsRequest)
    - [ListSessionsResponse](#auth-v1-ListSessionsResponse)
    - [RefreshSessionRequest](#auth-v1-RefreshSessionRequest)
    - [RefreshSessionResponse](#auth-v1-RefreshSessionResponse)
    - [RevokeAllSessionsRequest](#auth-v1-RevokeAllSessionsRequest)
    - [RevokeAllSessionsResponse](#auth-v1-RevokeAllSessionsResponse)
    - [RevokeSessionRequest](#auth-v1-RevokeSessionRequest)
    - [RevokeSessionResponse](#auth-v1-RevokeSessionResponse)
    - [SessionInfo](#auth-v1-SessionInfo)
  
    - [DeviceType](#auth-v1-DeviceType)
    - [OperatingSystem](#auth-v1-OperatingSystem)
  
    - [AuthSessionService](#auth-v1-AuthSessionService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="auth_v1_common-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auth/v1/common.proto



<a name="auth-v1-AuthResponse"></a>

### AuthResponse
AuthResponse contains the complete authentication result.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tokens | [AuthTokens](#auth-v1-AuthTokens) |  | Authentication tokens for the user. |
| user | [UserInfo](#auth-v1-UserInfo) |  | User information. |
| is_new_user | [bool](#bool) |  | Indicates if this was a new user registration. |






<a name="auth-v1-AuthTokens"></a>

### AuthTokens
AuthTokens contains standard authentication tokens returned by auth operations.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  | Authentication token to be used for authorized requests. |
| expires_at | [int64](#int64) |  | Token expiration time in Unix timestamp format. |
| refresh_token | [string](#string) |  | Refresh token for obtaining a new access token when the current one expires. |






<a name="auth-v1-SuccessResponse"></a>

### SuccessResponse
SuccessResponse is a simple response indicating operation success.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  | Indicates whether the operation was successful. |
| message | [string](#string) | optional | Optional message providing additional context. |






<a name="auth-v1-UserInfo"></a>

### UserInfo
UserInfo contains basic user information returned by auth operations.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) |  | Unique identifier for the user. |
| email | [string](#string) |  | User's email address. |
| name | [string](#string) | optional | Optional user name. |
| email_verified | [bool](#bool) |  | Indicates if the user has a verified email. |
| created_at | [int64](#int64) |  | Creation timestamp of the user account. |
| updated_at | [int64](#int64) |  | Last update timestamp of the user account. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="auth_v1_google_oauth-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auth/v1/google_oauth.proto



<a name="auth-v1-ExchangeAuthCodeRequest"></a>

### ExchangeAuthCodeRequest
ExchangeAuthCodeRequest contains parameters for exchanging a Google authorization code.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [string](#string) |  | The authorization code received from Google |
| redirect_uri | [string](#string) |  | Redirect URI used in the authorization request |
| state | [string](#string) |  | State parameter from the original request for verification |






<a name="auth-v1-ExchangeAuthCodeResponse"></a>

### ExchangeAuthCodeResponse
ExchangeAuthCodeResponse contains the authentication result.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| auth | [AuthResponse](#auth-v1-AuthResponse) |  | Authentication data including tokens and user information |






<a name="auth-v1-GetAuthURLRequest"></a>

### GetAuthURLRequest
GetAuthURLRequest contains parameters for generating a Google OAuth URL.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| redirect_uri | [string](#string) |  | URI to redirect to after authorization |
| state | [string](#string) | optional | Optional state parameter for CSRF protection |






<a name="auth-v1-GetAuthURLResponse"></a>

### GetAuthURLResponse
GetAuthURLResponse contains the Google OAuth authorization URL.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| auth_url | [string](#string) |  | URL to redirect users to for Google authentication |
| state | [string](#string) |  | State parameter echoed back from the request or auto-generated if not provided |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="auth-v1-AuthGoogleOAuthService"></a>

### AuthGoogleOAuthService
AuthGoogleOAuthService provides Google OAuth authentication operations.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetAuthURL | [GetAuthURLRequest](#auth-v1-GetAuthURLRequest) | [GetAuthURLResponse](#auth-v1-GetAuthURLResponse) | GetAuthURL generates a Google OAuth authorization URL. This URL is used to redirect users to Google's authentication page. |
| ExchangeAuthCode | [ExchangeAuthCodeRequest](#auth-v1-ExchangeAuthCodeRequest) | [ExchangeAuthCodeResponse](#auth-v1-ExchangeAuthCodeResponse) | ExchangeAuthCode exchanges a Google OAuth authorization code for authentication tokens. It validates the authorization code and issues authentication tokens upon success. |

 <!-- end services -->



<a name="auth_v1_magic_link-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auth/v1/magic_link.proto



<a name="auth-v1-SendMagicLinkRequest"></a>

### SendMagicLinkRequest
SendMagicLinkRequest contains information for generating a magic link.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  | Email address of the user. Must be a valid email format between 5 and 255 characters. |
| name | [string](#string) | optional | Optional user name. If provided, must be between 2 and 100 characters. Only used when creating a new user account. |
| redirect_url | [string](#string) | optional | Optional redirect URL after successful authentication. If provided, user will be redirected to this URL with the token as a parameter. |
| ttl_seconds | [int32](#int32) | optional | Optional time-to-live for the magic link in seconds. If not provided, a default value will be used. |






<a name="auth-v1-SendMagicLinkResponse"></a>

### SendMagicLinkResponse
SendMagicLinkResponse indicates the magic link generation result.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [SuccessResponse](#auth-v1-SuccessResponse) |  | Indicates whether the magic link was successfully generated and sent. Note: For security reasons, this may return true even if the email doesn't exist. |
| expires_at | [int64](#int64) | optional | Optional expiration time in Unix timestamp format. |
| is_new_user | [bool](#bool) | optional | Indicates whether a new user account was created. Will be true for new registrations, false for existing users. |
| user_id | [string](#string) | optional | User ID of the user. Only provided for new user registrations. |






<a name="auth-v1-VerifyMagicLinkRequest"></a>

### VerifyMagicLinkRequest
VerifyMagicLinkRequest contains the token for verification.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  | Magic link token received via email. Must not be empty. |






<a name="auth-v1-VerifyMagicLinkResponse"></a>

### VerifyMagicLinkResponse
VerifyMagicLinkResponse contains the authentication result.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| auth | [AuthResponse](#auth-v1-AuthResponse) |  | Authentication data including tokens and user information. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="auth-v1-AuthMagicLinkService"></a>

### AuthMagicLinkService
AuthMagicLinkService provides magic link-based authentication operations.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SendMagicLink | [SendMagicLinkRequest](#auth-v1-SendMagicLinkRequest) | [SendMagicLinkResponse](#auth-v1-SendMagicLinkResponse) | SendMagicLink initiates the magic link authentication process. It generates a secure link and sends it to the user's email address. This method handles both login and registration - if the user doesn't exist, a new account will be created automatically. |
| VerifyMagicLink | [VerifyMagicLinkRequest](#auth-v1-VerifyMagicLinkRequest) | [VerifyMagicLinkResponse](#auth-v1-VerifyMagicLinkResponse) | VerifyMagicLink authenticates a user using a magic link token. It verifies the provided token and returns authentication tokens upon success. |

 <!-- end services -->



<a name="auth_v1_password-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auth/v1/password.proto



<a name="auth-v1-ForgotPasswordRequest"></a>

### ForgotPasswordRequest
ForgotPasswordRequest initiates password recovery.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  | Email address of the user requesting password recovery. Must be a valid email format between 5 and 255 characters. |






<a name="auth-v1-ForgotPasswordResponse"></a>

### ForgotPasswordResponse
ForgotPasswordResponse indicates the password recovery process has started.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [SuccessResponse](#auth-v1-SuccessResponse) |  | Indicates whether the password recovery request was successfully processed. Note: For security reasons, this may return true even if the email doesn't exist. |






<a name="auth-v1-LoginRequest"></a>

### LoginRequest
LoginRequest contains credentials for authentication.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  | Email address of the user. Must be a valid email format between 5 and 255 characters. |
| password | [string](#string) |  | User password. Must be between 6 and 100 characters. |






<a name="auth-v1-LoginResponse"></a>

### LoginResponse
LoginResponse contains the authentication result.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| auth | [AuthResponse](#auth-v1-AuthResponse) |  | Authentication data including tokens and user information. |






<a name="auth-v1-RegisterRequest"></a>

### RegisterRequest
RegisterRequest contains information for creating a new user account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  | Email address of the user. Must be a valid email format between 5 and 255 characters. |
| password | [string](#string) |  | User password. Must be between 6 and 100 characters. |
| name | [string](#string) | optional | Optional user name. If provided, must be between 2 and 100 characters. |






<a name="auth-v1-RegisterResponse"></a>

### RegisterResponse
RegisterResponse contains the registration result.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| auth | [AuthResponse](#auth-v1-AuthResponse) |  | Authentication data including tokens and user information. |






<a name="auth-v1-ResetPasswordRequest"></a>

### ResetPasswordRequest
ResetPasswordRequest contains information for resetting a password.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  | Password reset token received via email. Must not be empty. |
| password | [string](#string) |  | New password to set. Must be between 6 and 100 characters. |
| password_confirmation | [string](#string) |  | New password confirmation. Must match the password field and be between 6 and 100 characters. |






<a name="auth-v1-ResetPasswordResponse"></a>

### ResetPasswordResponse
ResetPasswordResponse indicates the result of the password reset.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [SuccessResponse](#auth-v1-SuccessResponse) |  | Indicates whether the password was successfully reset. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="auth-v1-AuthPasswordService"></a>

### AuthPasswordService
AuthPasswordService provides password-based authentication operations.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Login | [LoginRequest](#auth-v1-LoginRequest) | [LoginResponse](#auth-v1-LoginResponse) | Login authenticates a user using email and password. It verifies the provided credentials and returns authentication tokens upon success. |
| Register | [RegisterRequest](#auth-v1-RegisterRequest) | [RegisterResponse](#auth-v1-RegisterResponse) | Register creates a new user account with email and password. It performs validation on the provided data and creates a new user account. Returns authentication tokens upon successful registration. |
| ForgotPassword | [ForgotPasswordRequest](#auth-v1-ForgotPasswordRequest) | [ForgotPasswordResponse](#auth-v1-ForgotPasswordResponse) | ForgotPassword initiates the password recovery process. It sends a password reset link to the user's email address. |
| ResetPassword | [ResetPasswordRequest](#auth-v1-ResetPasswordRequest) | [ResetPasswordResponse](#auth-v1-ResetPasswordResponse) | ResetPassword changes a user's password using a reset token. It validates the token and updates the user's password. |

 <!-- end services -->



<a name="auth_v1_session-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auth/v1/session.proto



<a name="auth-v1-GetSessionRequest"></a>

### GetSessionRequest
GetSessionRequest contains the token to retrieve session information.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  | Authentication token for the session. Must not be empty. |






<a name="auth-v1-GetSessionResponse"></a>

### GetSessionResponse
GetSessionResponse contains information about the session.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session | [SessionInfo](#auth-v1-SessionInfo) |  | Session information. |
| user | [UserInfo](#auth-v1-UserInfo) |  | User information associated with the session. |






<a name="auth-v1-ListSessionsRequest"></a>

### ListSessionsRequest
ListSessionsRequest requests all active sessions for a user.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) | optional | User ID for which to list sessions. If not provided, the user associated with the token will be used. |
| token | [string](#string) | optional | Authentication token. If user_id is not provided, this must be provided. |






<a name="auth-v1-ListSessionsResponse"></a>

### ListSessionsResponse
ListSessionsResponse contains a list of active sessions.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sessions | [SessionInfo](#auth-v1-SessionInfo) | repeated | List of active sessions. |
| total | [int32](#int32) |  | Total number of active sessions. |






<a name="auth-v1-RefreshSessionRequest"></a>

### RefreshSessionRequest
RefreshSessionRequest contains the refresh token to get new auth tokens.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| refresh_token | [string](#string) |  | Refresh token received during previous authentication. Must not be empty. |






<a name="auth-v1-RefreshSessionResponse"></a>

### RefreshSessionResponse
RefreshSessionResponse contains new authentication tokens.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tokens | [AuthTokens](#auth-v1-AuthTokens) |  | New authentication tokens. |






<a name="auth-v1-RevokeAllSessionsRequest"></a>

### RevokeAllSessionsRequest
RevokeAllSessionsRequest contains information for revoking all user sessions.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [string](#string) | optional | User ID for which to revoke all sessions. If not provided, the user associated with the token will be used. |
| token | [string](#string) | optional | Authentication token. If user_id is not provided, this must be provided. This session will also be revoked unless keep_current is true. |
| keep_current | [bool](#bool) | optional | If true, the current session (identified by token) will not be revoked. |






<a name="auth-v1-RevokeAllSessionsResponse"></a>

### RevokeAllSessionsResponse
RevokeAllSessionsResponse indicates the result of revoking all sessions.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [SuccessResponse](#auth-v1-SuccessResponse) |  | Indicates whether all sessions were successfully revoked. |
| revoked_count | [int32](#int32) |  | Number of sessions revoked. |






<a name="auth-v1-RevokeSessionRequest"></a>

### RevokeSessionRequest
RevokeSessionRequest contains information for revoking a session.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session_id | [string](#string) | optional | Session ID to revoke. Either session_id or token must be provided. |
| token | [string](#string) | optional | Authentication token for the session to revoke. Either session_id or token must be provided. |






<a name="auth-v1-RevokeSessionResponse"></a>

### RevokeSessionResponse
RevokeSessionResponse indicates the result of the session revocation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [SuccessResponse](#auth-v1-SuccessResponse) |  | Indicates whether the session was successfully revoked. |






<a name="auth-v1-SessionInfo"></a>

### SessionInfo
SessionInfo contains information about an authentication session.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session_id | [string](#string) |  | Unique identifier for the session. |
| user_id | [string](#string) |  | User ID associated with the session. |
| ip_address | [string](#string) |  | IP address from which the session was created. |
| user_agent | [string](#string) |  | User agent information (browser, device, etc.). |
| device_type | [DeviceType](#auth-v1-DeviceType) |  | Device type (e.g., "Desktop", "Mobile", "Tablet"). |
| operating_system | [OperatingSystem](#auth-v1-OperatingSystem) |  | Operating system (e.g., "macOS", "Windows", "iOS", "Android"). |
| browser | [string](#string) |  | Browser name (e.g., "Chrome", "Firefox", "Safari"). |
| browser_version | [string](#string) |  | Browser version (e.g., "115.0.5790"). |
| device_model | [string](#string) | optional | Device model when available (e.g., "iPhone 14", "Pixel 7"). |
| location | [string](#string) | optional | Location information based on IP address (e.g., "San Francisco, US"). |
| created_at | [int64](#int64) |  | Creation timestamp of the session. |
| last_activity_at | [int64](#int64) |  | Last activity timestamp for the session. |
| expires_at | [int64](#int64) |  | Expiration timestamp of the session. |
| is_current | [bool](#bool) |  | Indicates whether this is the current active session. This is useful when listing multiple sessions to identify the current one. |





 <!-- end messages -->


<a name="auth-v1-DeviceType"></a>

### DeviceType
DeviceType represents the type of device used for authentication.

| Name | Number | Description |
| ---- | ------ | ----------- |
| DEVICE_TYPE_UNSPECIFIED | 0 | Default value for proto3 |
| DEVICE_TYPE_DESKTOP | 1 | Desktop computer |
| DEVICE_TYPE_MOBILE | 2 | Mobile phone |
| DEVICE_TYPE_TABLET | 3 | Tablet device |
| DEVICE_TYPE_OTHER | 4 | Other/unknown device type |



<a name="auth-v1-OperatingSystem"></a>

### OperatingSystem
OperatingSystem represents the operating system of the device.

| Name | Number | Description |
| ---- | ------ | ----------- |
| OPERATING_SYSTEM_UNSPECIFIED | 0 | Default value for proto3 |
| OPERATING_SYSTEM_MACOS | 1 | macOS operating system |
| OPERATING_SYSTEM_WINDOWS | 2 | Windows operating system |
| OPERATING_SYSTEM_IOS | 3 | iOS mobile operating system |
| OPERATING_SYSTEM_ANDROID | 4 | Android mobile operating system |
| OPERATING_SYSTEM_LINUX | 5 | Linux operating system |
| OPERATING_SYSTEM_CHROMEOS | 6 | ChromeOS operating system |
| OPERATING_SYSTEM_OTHER | 7 | Other/unknown operating system |


 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="auth-v1-AuthSessionService"></a>

### AuthSessionService
AuthSessionService provides session management operations.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetSession | [GetSessionRequest](#auth-v1-GetSessionRequest) | [GetSessionResponse](#auth-v1-GetSessionResponse) | GetSession retrieves information about the current authenticated session. It validates the provided token and returns session details. |
| RefreshSession | [RefreshSessionRequest](#auth-v1-RefreshSessionRequest) | [RefreshSessionResponse](#auth-v1-RefreshSessionResponse) | RefreshSession generates new authentication tokens using a refresh token. It validates the refresh token and issues a new access token. |
| RevokeSession | [RevokeSessionRequest](#auth-v1-RevokeSessionRequest) | [RevokeSessionResponse](#auth-v1-RevokeSessionResponse) | RevokeSession invalidates a specific session. This can be used when a user logs out from a specific device. |
| ListSessions | [ListSessionsRequest](#auth-v1-ListSessionsRequest) | [ListSessionsResponse](#auth-v1-ListSessionsResponse) | ListSessions returns all active sessions for a user. This can be used to show users all devices where they're currently logged in. |
| RevokeAllSessions | [RevokeAllSessionsRequest](#auth-v1-RevokeAllSessionsRequest) | [RevokeAllSessionsResponse](#auth-v1-RevokeAllSessionsResponse) | RevokeAllSessions invalidates all sessions for a user. This can be used for security purposes or when a user wants to log out from all devices. |

 <!-- end services -->



## Scalar Value Types

| .proto Type | Go | TypeScript |
| ----------- | -- | ---------- |
| double | float64 | number | 
| float | float32 | number | 
| int32 | int32 | number | 
| int64 | int64 | bigint | 
| uint32 | uint32 | number | 
| uint64 | uint64 | bigint | 
| sint32 | int32 | number | 
| sint64 | int64 | bigint | 
| fixed32 | uint32 | number | 
| fixed64 | uint64 | bigint | 
| sfixed32 | int32 | number | 
| sfixed64 | int64 | bigint | 
| bool | bool | boolean | 
| string | string | string | 
| bytes | []byte | Uint8Array | 