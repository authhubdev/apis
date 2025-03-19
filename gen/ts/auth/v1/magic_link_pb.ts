// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file auth/v1/magic_link.proto (package auth.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { AuthResponse, OTPVerification } from "./common_pb";
import { file_auth_v1_common } from "./common_pb";
import { file_google_protobuf_wrappers } from "@bufbuild/protobuf/wkt";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file auth/v1/magic_link.proto.
 */
export const file_auth_v1_magic_link: GenFile = /*@__PURE__*/
  fileDesc("ChhhdXRoL3YxL21hZ2ljX2xpbmsucHJvdG8SB2F1dGgudjEihQEKFFNlbmRNYWdpY0xpbmtSZXF1ZXN0Eg0KBWVtYWlsGAEgASgJEioKBG5hbWUYAiABKAsyHC5nb29nbGUucHJvdG9idWYuU3RyaW5nVmFsdWUSMgoMcmVkaXJlY3RfdXJsGAMgASgLMhwuZ29vZ2xlLnByb3RvYnVmLlN0cmluZ1ZhbHVlIkUKFVNlbmRNYWdpY0xpbmtSZXNwb25zZRIXCg92ZXJpZmljYXRpb25faWQYASABKAkSEwoLaXNfbmV3X3VzZXIYAiABKAgiSAoWVmVyaWZ5TWFnaWNMaW5rUmVxdWVzdBIuCgx2ZXJpZmljYXRpb24YASABKAsyGC5hdXRoLnYxLk9UUFZlcmlmaWNhdGlvbiI+ChdWZXJpZnlNYWdpY0xpbmtSZXNwb25zZRIjCgRhdXRoGAEgASgLMhUuYXV0aC52MS5BdXRoUmVzcG9uc2UyvAEKEE1hZ2ljTGlua1NlcnZpY2USUAoNU2VuZE1hZ2ljTGluaxIdLmF1dGgudjEuU2VuZE1hZ2ljTGlua1JlcXVlc3QaHi5hdXRoLnYxLlNlbmRNYWdpY0xpbmtSZXNwb25zZSIAElYKD1ZlcmlmeU1hZ2ljTGluaxIfLmF1dGgudjEuVmVyaWZ5TWFnaWNMaW5rUmVxdWVzdBogLmF1dGgudjEuVmVyaWZ5TWFnaWNMaW5rUmVzcG9uc2UiAEIyWjBnaXRodWIuY29tL2F1dGhodWJkZXYvYXBpcy9nZW4vZ28vYXV0aC92MTthdXRodjFiBnByb3RvMw", [file_auth_v1_common, file_google_protobuf_wrappers]);

/**
 * SendMagicLinkRequest is used to request a magic link for authentication.
 *
 * @generated from message auth.v1.SendMagicLinkRequest
 */
export type SendMagicLinkRequest = Message<"auth.v1.SendMagicLinkRequest"> & {
  /**
   * Email address to send the magic link to.
   *
   * @generated from field: string email = 1;
   */
  email: string;

  /**
   * Optional name for new users.
   *
   * @generated from field: google.protobuf.StringValue name = 2;
   */
  name?: string;

  /**
   * Optional redirect URL after successful verification.
   *
   * @generated from field: google.protobuf.StringValue redirect_url = 3;
   */
  redirectUrl?: string;
};

/**
 * Describes the message auth.v1.SendMagicLinkRequest.
 * Use `create(SendMagicLinkRequestSchema)` to create a new message.
 */
export const SendMagicLinkRequestSchema: GenMessage<SendMagicLinkRequest> = /*@__PURE__*/
  messageDesc(file_auth_v1_magic_link, 0);

/**
 * SendMagicLinkResponse confirms that a magic link was sent.
 *
 * @generated from message auth.v1.SendMagicLinkResponse
 */
export type SendMagicLinkResponse = Message<"auth.v1.SendMagicLinkResponse"> & {
  /**
   * Verification ID to be used with the OTP.
   *
   * @generated from field: string verification_id = 1;
   */
  verificationId: string;

  /**
   * Indicates if a new user was created.
   *
   * @generated from field: bool is_new_user = 2;
   */
  isNewUser: boolean;
};

/**
 * Describes the message auth.v1.SendMagicLinkResponse.
 * Use `create(SendMagicLinkResponseSchema)` to create a new message.
 */
export const SendMagicLinkResponseSchema: GenMessage<SendMagicLinkResponse> = /*@__PURE__*/
  messageDesc(file_auth_v1_magic_link, 1);

/**
 * VerifyMagicLinkRequest is used to verify a magic link.
 *
 * @generated from message auth.v1.VerifyMagicLinkRequest
 */
export type VerifyMagicLinkRequest = Message<"auth.v1.VerifyMagicLinkRequest"> & {
  /**
   * OTP verification data.
   *
   * @generated from field: auth.v1.OTPVerification verification = 1;
   */
  verification?: OTPVerification;
};

/**
 * Describes the message auth.v1.VerifyMagicLinkRequest.
 * Use `create(VerifyMagicLinkRequestSchema)` to create a new message.
 */
export const VerifyMagicLinkRequestSchema: GenMessage<VerifyMagicLinkRequest> = /*@__PURE__*/
  messageDesc(file_auth_v1_magic_link, 2);

/**
 * VerifyMagicLinkResponse returns authentication tokens after successful verification.
 *
 * @generated from message auth.v1.VerifyMagicLinkResponse
 */
export type VerifyMagicLinkResponse = Message<"auth.v1.VerifyMagicLinkResponse"> & {
  /**
   * Authentication tokens for the user.
   *
   * @generated from field: auth.v1.AuthResponse auth = 1;
   */
  auth?: AuthResponse;
};

/**
 * Describes the message auth.v1.VerifyMagicLinkResponse.
 * Use `create(VerifyMagicLinkResponseSchema)` to create a new message.
 */
export const VerifyMagicLinkResponseSchema: GenMessage<VerifyMagicLinkResponse> = /*@__PURE__*/
  messageDesc(file_auth_v1_magic_link, 3);

/**
 * MagicLinkService provides authentication via magic links.
 *
 * @generated from service auth.v1.MagicLinkService
 */
export const MagicLinkService: GenService<{
  /**
   * SendMagicLink sends a magic link to the specified email.
   * If the email doesn't exist, a new account will be created.
   *
   * @generated from rpc auth.v1.MagicLinkService.SendMagicLink
   */
  sendMagicLink: {
    methodKind: "unary";
    input: typeof SendMagicLinkRequestSchema;
    output: typeof SendMagicLinkResponseSchema;
  },
  /**
   * VerifyMagicLink validates the magic link OTP and returns auth tokens.
   *
   * @generated from rpc auth.v1.MagicLinkService.VerifyMagicLink
   */
  verifyMagicLink: {
    methodKind: "unary";
    input: typeof VerifyMagicLinkRequestSchema;
    output: typeof VerifyMagicLinkResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_auth_v1_magic_link, 0);

