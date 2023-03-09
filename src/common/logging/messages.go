package logging

import "go.uber.org/zap/zapcore"

type messageID string

type message struct {
	level   zapcore.Level
	message string
}

const (
	ConvertCaseStart    messageID = "ConvertCaseStart"
	ConvertCaseEnd      messageID = "ConvertCaseEnd"
	EncodeUrlStart      messageID = "EncodeUrlStart"
	EncodeUrlEnd        messageID = "EncodeUrlEnd"
	DecodeUrlStart      messageID = "DecodeUrlStart"
	DecodeUrlEnd        messageID = "DecodeUrlEnd"
	FormatJsonStart     messageID = "FormatJsonStart"
	FormatJsonEnd       messageID = "FormatJsonEnd"
	GenerateRandomStart messageID = "GenerateRandomStart"
	GenerateRandomEnd   messageID = "GenerateRandomEnd"
	HashStart           messageID = "HashStart"
	HashEnd             messageID = "HashEnd"
	EncodeBaseStart     messageID = "EncodeBaseStart"
	EncodeBaseEnd       messageID = "EncodeBaseEnd"
	EncodeBaseError     messageID = "EncodeBaseError"
	DecodeBaseStart     messageID = "DecodeBaseStart"
	DecodeBaseEnd       messageID = "DecodeBaseEnd"
	DecodeBaseError     messageID = "DecodeBaseError"
	DecodeJwtStart      messageID = "DecodeJwtStart"
	DecodeJwtEnd        messageID = "DecodeJwtEnd"
	EncryptAESStart     messageID = "EncryptAESStart"
	EncryptAESEnd       messageID = "EncryptAESEnd"
	DecryptAESStart     messageID = "DecryptAESStart"
	DecryptAESEnd       messageID = "DecryptAESEnd"
)

var messages = map[messageID]message{
	ConvertCaseStart:    {zapcore.InfoLevel, "Convert Case Start"},
	ConvertCaseEnd:      {zapcore.InfoLevel, "Convert Case End"},
	EncodeUrlStart:      {zapcore.InfoLevel, "Encode Url Start"},
	EncodeUrlEnd:        {zapcore.InfoLevel, "Encode Url End"},
	DecodeUrlStart:      {zapcore.InfoLevel, "Decode Url Start"},
	DecodeUrlEnd:        {zapcore.InfoLevel, "Decode Url SEnd"},
	FormatJsonStart:     {zapcore.InfoLevel, "Format Json Start"},
	FormatJsonEnd:       {zapcore.InfoLevel, "Format Json End"},
	GenerateRandomStart: {zapcore.InfoLevel, "Generate Random Start"},
	GenerateRandomEnd:   {zapcore.InfoLevel, "Generate Random End"},
	HashStart:           {zapcore.InfoLevel, "Hash Start"},
	HashEnd:             {zapcore.InfoLevel, "Hash End"},
	EncodeBaseStart:     {zapcore.InfoLevel, "Encode Base Start"},
	EncodeBaseEnd:       {zapcore.InfoLevel, "Encode Base End"},
	EncodeBaseError:     {zapcore.WarnLevel, "Encode Base Error"},
	DecodeBaseStart:     {zapcore.InfoLevel, "Decode Base Start"},
	DecodeBaseEnd:       {zapcore.InfoLevel, "Decode Base End"},
	DecodeBaseError:     {zapcore.WarnLevel, "Decode Base Error"},
	DecodeJwtStart:      {zapcore.InfoLevel, "Decode JWT Start"},
	DecodeJwtEnd:        {zapcore.InfoLevel, "Decode JWT End"},
	EncryptAESStart:     {zapcore.InfoLevel, "Encrypt AES256-GCM Start"},
	EncryptAESEnd:       {zapcore.InfoLevel, "Encrypt AES256-GCM End"},
	DecryptAESStart:     {zapcore.InfoLevel, "Decrypt AES256-GCM Start"},
	DecryptAESEnd:       {zapcore.InfoLevel, "Decrypt AES256-GCM End"},
}
