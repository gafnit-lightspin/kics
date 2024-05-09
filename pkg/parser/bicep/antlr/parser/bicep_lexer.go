// Code generated from bicep.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type bicepLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var BicepLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func biceplexerLexerInit() {
	staticData := &BicepLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "'@'", "','", "'['", "']'", "'('", "')'", "'.'", "'|'", "",
		"'='", "'{'", "'}'", "'param'", "'var'", "'true'", "'false'", "'null'",
		"'array'", "'object'", "'resource'", "'output'", "'targetScope'", "'import'",
		"'with'", "'as'", "'metadata'", "'existing'", "'type'", "'module'",
		"", "", "", "", "'string'", "'int'", "'bool'", "'if'", "'for'", "'in'",
		"'?'", "'>'", "'>='", "'<'", "'<='", "'=='", "'!='",
	}
	staticData.SymbolicNames = []string{
		"", "MULTILINE_STRING", "AT", "COMMA", "OBRACK", "CBRACK", "OPAR", "CPAR",
		"DOT", "PIPE", "COL", "ASSIGN", "OBRACE", "CBRACE", "PARAM", "VAR",
		"TRUE", "FALSE", "NULL", "ARRAY", "OBJECT", "RESOURCE", "OUTPUT", "TARGET_SCOPE",
		"IMPORT", "WITH", "AS", "METADATA", "EXISTING", "TYPE", "MODULE", "STRING_LEFT_PIECE",
		"STRING_MIDDLE_PIECE", "STRING_RIGHT_PIECE", "STRING_COMPLETE", "STRING",
		"INT", "BOOL", "IF", "FOR", "IN", "QMARK", "GT", "GTE", "LT", "LTE",
		"EQ", "NEQ", "IDENTIFIER", "NUMBER", "NL", "SINGLE_LINE_COMMENT", "MULTI_LINE_COMMENT",
		"SPACES", "UNKNOWN",
	}
	staticData.RuleNames = []string{
		"MULTILINE_STRING", "AT", "COMMA", "OBRACK", "CBRACK", "OPAR", "CPAR",
		"DOT", "PIPE", "COL", "ASSIGN", "OBRACE", "CBRACE", "PARAM", "VAR",
		"TRUE", "FALSE", "NULL", "ARRAY", "OBJECT", "RESOURCE", "OUTPUT", "TARGET_SCOPE",
		"IMPORT", "WITH", "AS", "METADATA", "EXISTING", "TYPE", "MODULE", "STRING_LEFT_PIECE",
		"STRING_MIDDLE_PIECE", "STRING_RIGHT_PIECE", "STRING_COMPLETE", "STRING",
		"INT", "BOOL", "IF", "FOR", "IN", "QMARK", "GT", "GTE", "LT", "LTE",
		"EQ", "NEQ", "IDENTIFIER", "NUMBER", "NL", "SINGLE_LINE_COMMENT", "MULTI_LINE_COMMENT",
		"SPACES", "UNKNOWN", "STRINGCHAR", "ESCAPE", "HEX",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 54, 429, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 121, 8, 0, 10, 0, 12, 0, 124, 9, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 3, 9, 149,
		8, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 5, 30, 271, 8, 30, 10, 30, 12, 30,
		274, 9, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 5, 31, 281, 8, 31, 10, 31,
		12, 31, 284, 9, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 5, 32, 291, 8, 32,
		10, 32, 12, 32, 294, 9, 32, 1, 32, 1, 32, 1, 33, 1, 33, 5, 33, 300, 8,
		33, 10, 33, 12, 33, 303, 9, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39,
		1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1,
		44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47,
		5, 47, 353, 8, 47, 10, 47, 12, 47, 356, 9, 47, 1, 48, 4, 48, 359, 8, 48,
		11, 48, 12, 48, 360, 1, 48, 1, 48, 4, 48, 365, 8, 48, 11, 48, 12, 48, 366,
		3, 48, 369, 8, 48, 1, 49, 4, 49, 372, 8, 49, 11, 49, 12, 49, 373, 1, 50,
		1, 50, 1, 50, 1, 50, 5, 50, 380, 8, 50, 10, 50, 12, 50, 383, 9, 50, 1,
		50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 5, 51, 391, 8, 51, 10, 51, 12, 51,
		394, 9, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 52, 4, 52, 402, 8, 52,
		11, 52, 12, 52, 403, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54, 3, 54, 412,
		8, 54, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 4, 55, 420, 8, 55, 11,
		55, 12, 55, 421, 1, 55, 1, 55, 3, 55, 426, 8, 55, 1, 56, 1, 56, 2, 122,
		392, 0, 57, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37,
		19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55,
		28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73,
		37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91,
		46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54,
		109, 0, 111, 0, 113, 0, 1, 0, 8, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48,
		57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 2, 0, 10, 10, 13, 13, 2, 0,
		9, 9, 32, 32, 5, 0, 9, 10, 13, 13, 36, 36, 39, 39, 92, 92, 6, 0, 36, 36,
		39, 39, 92, 92, 110, 110, 114, 114, 116, 116, 3, 0, 48, 57, 65, 70, 97,
		102, 442, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0,
		0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1,
		0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61,
		1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0,
		69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0,
		0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0,
		0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0,
		0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1,
		0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0,
		107, 1, 0, 0, 0, 1, 115, 1, 0, 0, 0, 3, 129, 1, 0, 0, 0, 5, 131, 1, 0,
		0, 0, 7, 133, 1, 0, 0, 0, 9, 135, 1, 0, 0, 0, 11, 137, 1, 0, 0, 0, 13,
		139, 1, 0, 0, 0, 15, 141, 1, 0, 0, 0, 17, 143, 1, 0, 0, 0, 19, 148, 1,
		0, 0, 0, 21, 150, 1, 0, 0, 0, 23, 152, 1, 0, 0, 0, 25, 154, 1, 0, 0, 0,
		27, 156, 1, 0, 0, 0, 29, 162, 1, 0, 0, 0, 31, 166, 1, 0, 0, 0, 33, 171,
		1, 0, 0, 0, 35, 177, 1, 0, 0, 0, 37, 182, 1, 0, 0, 0, 39, 188, 1, 0, 0,
		0, 41, 195, 1, 0, 0, 0, 43, 204, 1, 0, 0, 0, 45, 211, 1, 0, 0, 0, 47, 223,
		1, 0, 0, 0, 49, 230, 1, 0, 0, 0, 51, 235, 1, 0, 0, 0, 53, 238, 1, 0, 0,
		0, 55, 247, 1, 0, 0, 0, 57, 256, 1, 0, 0, 0, 59, 261, 1, 0, 0, 0, 61, 268,
		1, 0, 0, 0, 63, 278, 1, 0, 0, 0, 65, 288, 1, 0, 0, 0, 67, 297, 1, 0, 0,
		0, 69, 306, 1, 0, 0, 0, 71, 313, 1, 0, 0, 0, 73, 317, 1, 0, 0, 0, 75, 322,
		1, 0, 0, 0, 77, 325, 1, 0, 0, 0, 79, 329, 1, 0, 0, 0, 81, 332, 1, 0, 0,
		0, 83, 334, 1, 0, 0, 0, 85, 336, 1, 0, 0, 0, 87, 339, 1, 0, 0, 0, 89, 341,
		1, 0, 0, 0, 91, 344, 1, 0, 0, 0, 93, 347, 1, 0, 0, 0, 95, 350, 1, 0, 0,
		0, 97, 358, 1, 0, 0, 0, 99, 371, 1, 0, 0, 0, 101, 375, 1, 0, 0, 0, 103,
		386, 1, 0, 0, 0, 105, 401, 1, 0, 0, 0, 107, 407, 1, 0, 0, 0, 109, 411,
		1, 0, 0, 0, 111, 413, 1, 0, 0, 0, 113, 427, 1, 0, 0, 0, 115, 116, 5, 39,
		0, 0, 116, 117, 5, 39, 0, 0, 117, 118, 5, 39, 0, 0, 118, 122, 1, 0, 0,
		0, 119, 121, 9, 0, 0, 0, 120, 119, 1, 0, 0, 0, 121, 124, 1, 0, 0, 0, 122,
		123, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 123, 125, 1, 0, 0, 0, 124, 122,
		1, 0, 0, 0, 125, 126, 5, 39, 0, 0, 126, 127, 5, 39, 0, 0, 127, 128, 5,
		39, 0, 0, 128, 2, 1, 0, 0, 0, 129, 130, 5, 64, 0, 0, 130, 4, 1, 0, 0, 0,
		131, 132, 5, 44, 0, 0, 132, 6, 1, 0, 0, 0, 133, 134, 5, 91, 0, 0, 134,
		8, 1, 0, 0, 0, 135, 136, 5, 93, 0, 0, 136, 10, 1, 0, 0, 0, 137, 138, 5,
		40, 0, 0, 138, 12, 1, 0, 0, 0, 139, 140, 5, 41, 0, 0, 140, 14, 1, 0, 0,
		0, 141, 142, 5, 46, 0, 0, 142, 16, 1, 0, 0, 0, 143, 144, 5, 124, 0, 0,
		144, 18, 1, 0, 0, 0, 145, 149, 5, 58, 0, 0, 146, 147, 5, 58, 0, 0, 147,
		149, 5, 58, 0, 0, 148, 145, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 149, 20,
		1, 0, 0, 0, 150, 151, 5, 61, 0, 0, 151, 22, 1, 0, 0, 0, 152, 153, 5, 123,
		0, 0, 153, 24, 1, 0, 0, 0, 154, 155, 5, 125, 0, 0, 155, 26, 1, 0, 0, 0,
		156, 157, 5, 112, 0, 0, 157, 158, 5, 97, 0, 0, 158, 159, 5, 114, 0, 0,
		159, 160, 5, 97, 0, 0, 160, 161, 5, 109, 0, 0, 161, 28, 1, 0, 0, 0, 162,
		163, 5, 118, 0, 0, 163, 164, 5, 97, 0, 0, 164, 165, 5, 114, 0, 0, 165,
		30, 1, 0, 0, 0, 166, 167, 5, 116, 0, 0, 167, 168, 5, 114, 0, 0, 168, 169,
		5, 117, 0, 0, 169, 170, 5, 101, 0, 0, 170, 32, 1, 0, 0, 0, 171, 172, 5,
		102, 0, 0, 172, 173, 5, 97, 0, 0, 173, 174, 5, 108, 0, 0, 174, 175, 5,
		115, 0, 0, 175, 176, 5, 101, 0, 0, 176, 34, 1, 0, 0, 0, 177, 178, 5, 110,
		0, 0, 178, 179, 5, 117, 0, 0, 179, 180, 5, 108, 0, 0, 180, 181, 5, 108,
		0, 0, 181, 36, 1, 0, 0, 0, 182, 183, 5, 97, 0, 0, 183, 184, 5, 114, 0,
		0, 184, 185, 5, 114, 0, 0, 185, 186, 5, 97, 0, 0, 186, 187, 5, 121, 0,
		0, 187, 38, 1, 0, 0, 0, 188, 189, 5, 111, 0, 0, 189, 190, 5, 98, 0, 0,
		190, 191, 5, 106, 0, 0, 191, 192, 5, 101, 0, 0, 192, 193, 5, 99, 0, 0,
		193, 194, 5, 116, 0, 0, 194, 40, 1, 0, 0, 0, 195, 196, 5, 114, 0, 0, 196,
		197, 5, 101, 0, 0, 197, 198, 5, 115, 0, 0, 198, 199, 5, 111, 0, 0, 199,
		200, 5, 117, 0, 0, 200, 201, 5, 114, 0, 0, 201, 202, 5, 99, 0, 0, 202,
		203, 5, 101, 0, 0, 203, 42, 1, 0, 0, 0, 204, 205, 5, 111, 0, 0, 205, 206,
		5, 117, 0, 0, 206, 207, 5, 116, 0, 0, 207, 208, 5, 112, 0, 0, 208, 209,
		5, 117, 0, 0, 209, 210, 5, 116, 0, 0, 210, 44, 1, 0, 0, 0, 211, 212, 5,
		116, 0, 0, 212, 213, 5, 97, 0, 0, 213, 214, 5, 114, 0, 0, 214, 215, 5,
		103, 0, 0, 215, 216, 5, 101, 0, 0, 216, 217, 5, 116, 0, 0, 217, 218, 5,
		83, 0, 0, 218, 219, 5, 99, 0, 0, 219, 220, 5, 111, 0, 0, 220, 221, 5, 112,
		0, 0, 221, 222, 5, 101, 0, 0, 222, 46, 1, 0, 0, 0, 223, 224, 5, 105, 0,
		0, 224, 225, 5, 109, 0, 0, 225, 226, 5, 112, 0, 0, 226, 227, 5, 111, 0,
		0, 227, 228, 5, 114, 0, 0, 228, 229, 5, 116, 0, 0, 229, 48, 1, 0, 0, 0,
		230, 231, 5, 119, 0, 0, 231, 232, 5, 105, 0, 0, 232, 233, 5, 116, 0, 0,
		233, 234, 5, 104, 0, 0, 234, 50, 1, 0, 0, 0, 235, 236, 5, 97, 0, 0, 236,
		237, 5, 115, 0, 0, 237, 52, 1, 0, 0, 0, 238, 239, 5, 109, 0, 0, 239, 240,
		5, 101, 0, 0, 240, 241, 5, 116, 0, 0, 241, 242, 5, 97, 0, 0, 242, 243,
		5, 100, 0, 0, 243, 244, 5, 97, 0, 0, 244, 245, 5, 116, 0, 0, 245, 246,
		5, 97, 0, 0, 246, 54, 1, 0, 0, 0, 247, 248, 5, 101, 0, 0, 248, 249, 5,
		120, 0, 0, 249, 250, 5, 105, 0, 0, 250, 251, 5, 115, 0, 0, 251, 252, 5,
		116, 0, 0, 252, 253, 5, 105, 0, 0, 253, 254, 5, 110, 0, 0, 254, 255, 5,
		103, 0, 0, 255, 56, 1, 0, 0, 0, 256, 257, 5, 116, 0, 0, 257, 258, 5, 121,
		0, 0, 258, 259, 5, 112, 0, 0, 259, 260, 5, 101, 0, 0, 260, 58, 1, 0, 0,
		0, 261, 262, 5, 109, 0, 0, 262, 263, 5, 111, 0, 0, 263, 264, 5, 100, 0,
		0, 264, 265, 5, 117, 0, 0, 265, 266, 5, 108, 0, 0, 266, 267, 5, 101, 0,
		0, 267, 60, 1, 0, 0, 0, 268, 272, 5, 39, 0, 0, 269, 271, 3, 109, 54, 0,
		270, 269, 1, 0, 0, 0, 271, 274, 1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 272,
		273, 1, 0, 0, 0, 273, 275, 1, 0, 0, 0, 274, 272, 1, 0, 0, 0, 275, 276,
		5, 36, 0, 0, 276, 277, 5, 123, 0, 0, 277, 62, 1, 0, 0, 0, 278, 282, 5,
		125, 0, 0, 279, 281, 3, 109, 54, 0, 280, 279, 1, 0, 0, 0, 281, 284, 1,
		0, 0, 0, 282, 280, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 285, 1, 0, 0,
		0, 284, 282, 1, 0, 0, 0, 285, 286, 5, 36, 0, 0, 286, 287, 5, 123, 0, 0,
		287, 64, 1, 0, 0, 0, 288, 292, 5, 125, 0, 0, 289, 291, 3, 109, 54, 0, 290,
		289, 1, 0, 0, 0, 291, 294, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 292, 293,
		1, 0, 0, 0, 293, 295, 1, 0, 0, 0, 294, 292, 1, 0, 0, 0, 295, 296, 5, 39,
		0, 0, 296, 66, 1, 0, 0, 0, 297, 301, 5, 39, 0, 0, 298, 300, 3, 109, 54,
		0, 299, 298, 1, 0, 0, 0, 300, 303, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0, 301,
		302, 1, 0, 0, 0, 302, 304, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 304, 305,
		5, 39, 0, 0, 305, 68, 1, 0, 0, 0, 306, 307, 5, 115, 0, 0, 307, 308, 5,
		116, 0, 0, 308, 309, 5, 114, 0, 0, 309, 310, 5, 105, 0, 0, 310, 311, 5,
		110, 0, 0, 311, 312, 5, 103, 0, 0, 312, 70, 1, 0, 0, 0, 313, 314, 5, 105,
		0, 0, 314, 315, 5, 110, 0, 0, 315, 316, 5, 116, 0, 0, 316, 72, 1, 0, 0,
		0, 317, 318, 5, 98, 0, 0, 318, 319, 5, 111, 0, 0, 319, 320, 5, 111, 0,
		0, 320, 321, 5, 108, 0, 0, 321, 74, 1, 0, 0, 0, 322, 323, 5, 105, 0, 0,
		323, 324, 5, 102, 0, 0, 324, 76, 1, 0, 0, 0, 325, 326, 5, 102, 0, 0, 326,
		327, 5, 111, 0, 0, 327, 328, 5, 114, 0, 0, 328, 78, 1, 0, 0, 0, 329, 330,
		5, 105, 0, 0, 330, 331, 5, 110, 0, 0, 331, 80, 1, 0, 0, 0, 332, 333, 5,
		63, 0, 0, 333, 82, 1, 0, 0, 0, 334, 335, 5, 62, 0, 0, 335, 84, 1, 0, 0,
		0, 336, 337, 5, 62, 0, 0, 337, 338, 5, 61, 0, 0, 338, 86, 1, 0, 0, 0, 339,
		340, 5, 60, 0, 0, 340, 88, 1, 0, 0, 0, 341, 342, 5, 60, 0, 0, 342, 343,
		5, 61, 0, 0, 343, 90, 1, 0, 0, 0, 344, 345, 5, 61, 0, 0, 345, 346, 5, 61,
		0, 0, 346, 92, 1, 0, 0, 0, 347, 348, 5, 33, 0, 0, 348, 349, 5, 61, 0, 0,
		349, 94, 1, 0, 0, 0, 350, 354, 7, 0, 0, 0, 351, 353, 7, 1, 0, 0, 352, 351,
		1, 0, 0, 0, 353, 356, 1, 0, 0, 0, 354, 352, 1, 0, 0, 0, 354, 355, 1, 0,
		0, 0, 355, 96, 1, 0, 0, 0, 356, 354, 1, 0, 0, 0, 357, 359, 7, 2, 0, 0,
		358, 357, 1, 0, 0, 0, 359, 360, 1, 0, 0, 0, 360, 358, 1, 0, 0, 0, 360,
		361, 1, 0, 0, 0, 361, 368, 1, 0, 0, 0, 362, 364, 5, 46, 0, 0, 363, 365,
		7, 2, 0, 0, 364, 363, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0, 366, 364, 1, 0,
		0, 0, 366, 367, 1, 0, 0, 0, 367, 369, 1, 0, 0, 0, 368, 362, 1, 0, 0, 0,
		368, 369, 1, 0, 0, 0, 369, 98, 1, 0, 0, 0, 370, 372, 7, 3, 0, 0, 371, 370,
		1, 0, 0, 0, 372, 373, 1, 0, 0, 0, 373, 371, 1, 0, 0, 0, 373, 374, 1, 0,
		0, 0, 374, 100, 1, 0, 0, 0, 375, 376, 5, 47, 0, 0, 376, 377, 5, 47, 0,
		0, 377, 381, 1, 0, 0, 0, 378, 380, 8, 3, 0, 0, 379, 378, 1, 0, 0, 0, 380,
		383, 1, 0, 0, 0, 381, 379, 1, 0, 0, 0, 381, 382, 1, 0, 0, 0, 382, 384,
		1, 0, 0, 0, 383, 381, 1, 0, 0, 0, 384, 385, 6, 50, 0, 0, 385, 102, 1, 0,
		0, 0, 386, 387, 5, 47, 0, 0, 387, 388, 5, 42, 0, 0, 388, 392, 1, 0, 0,
		0, 389, 391, 9, 0, 0, 0, 390, 389, 1, 0, 0, 0, 391, 394, 1, 0, 0, 0, 392,
		393, 1, 0, 0, 0, 392, 390, 1, 0, 0, 0, 393, 395, 1, 0, 0, 0, 394, 392,
		1, 0, 0, 0, 395, 396, 5, 42, 0, 0, 396, 397, 5, 47, 0, 0, 397, 398, 1,
		0, 0, 0, 398, 399, 6, 51, 0, 0, 399, 104, 1, 0, 0, 0, 400, 402, 7, 4, 0,
		0, 401, 400, 1, 0, 0, 0, 402, 403, 1, 0, 0, 0, 403, 401, 1, 0, 0, 0, 403,
		404, 1, 0, 0, 0, 404, 405, 1, 0, 0, 0, 405, 406, 6, 52, 0, 0, 406, 106,
		1, 0, 0, 0, 407, 408, 9, 0, 0, 0, 408, 108, 1, 0, 0, 0, 409, 412, 8, 5,
		0, 0, 410, 412, 3, 111, 55, 0, 411, 409, 1, 0, 0, 0, 411, 410, 1, 0, 0,
		0, 412, 110, 1, 0, 0, 0, 413, 425, 5, 92, 0, 0, 414, 426, 7, 6, 0, 0, 415,
		416, 5, 117, 0, 0, 416, 417, 5, 123, 0, 0, 417, 419, 1, 0, 0, 0, 418, 420,
		3, 113, 56, 0, 419, 418, 1, 0, 0, 0, 420, 421, 1, 0, 0, 0, 421, 419, 1,
		0, 0, 0, 421, 422, 1, 0, 0, 0, 422, 423, 1, 0, 0, 0, 423, 424, 5, 125,
		0, 0, 424, 426, 1, 0, 0, 0, 425, 414, 1, 0, 0, 0, 425, 415, 1, 0, 0, 0,
		426, 112, 1, 0, 0, 0, 427, 428, 7, 7, 0, 0, 428, 114, 1, 0, 0, 0, 18, 0,
		122, 148, 272, 282, 292, 301, 354, 360, 366, 368, 373, 381, 392, 403, 411,
		421, 425, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// bicepLexerInit initializes any static state used to implement bicepLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewbicepLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func BicepLexerInit() {
	staticData := &BicepLexerLexerStaticData
	staticData.once.Do(biceplexerLexerInit)
}

// NewbicepLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewbicepLexer(input antlr.CharStream) *bicepLexer {
	BicepLexerInit()
	l := new(bicepLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &BicepLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "bicep.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// bicepLexer tokens.
const (
	bicepLexerMULTILINE_STRING    = 1
	bicepLexerAT                  = 2
	bicepLexerCOMMA               = 3
	bicepLexerOBRACK              = 4
	bicepLexerCBRACK              = 5
	bicepLexerOPAR                = 6
	bicepLexerCPAR                = 7
	bicepLexerDOT                 = 8
	bicepLexerPIPE                = 9
	bicepLexerCOL                 = 10
	bicepLexerASSIGN              = 11
	bicepLexerOBRACE              = 12
	bicepLexerCBRACE              = 13
	bicepLexerPARAM               = 14
	bicepLexerVAR                 = 15
	bicepLexerTRUE                = 16
	bicepLexerFALSE               = 17
	bicepLexerNULL                = 18
	bicepLexerARRAY               = 19
	bicepLexerOBJECT              = 20
	bicepLexerRESOURCE            = 21
	bicepLexerOUTPUT              = 22
	bicepLexerTARGET_SCOPE        = 23
	bicepLexerIMPORT              = 24
	bicepLexerWITH                = 25
	bicepLexerAS                  = 26
	bicepLexerMETADATA            = 27
	bicepLexerEXISTING            = 28
	bicepLexerTYPE                = 29
	bicepLexerMODULE              = 30
	bicepLexerSTRING_LEFT_PIECE   = 31
	bicepLexerSTRING_MIDDLE_PIECE = 32
	bicepLexerSTRING_RIGHT_PIECE  = 33
	bicepLexerSTRING_COMPLETE     = 34
	bicepLexerSTRING              = 35
	bicepLexerINT                 = 36
	bicepLexerBOOL                = 37
	bicepLexerIF                  = 38
	bicepLexerFOR                 = 39
	bicepLexerIN                  = 40
	bicepLexerQMARK               = 41
	bicepLexerGT                  = 42
	bicepLexerGTE                 = 43
	bicepLexerLT                  = 44
	bicepLexerLTE                 = 45
	bicepLexerEQ                  = 46
	bicepLexerNEQ                 = 47
	bicepLexerIDENTIFIER          = 48
	bicepLexerNUMBER              = 49
	bicepLexerNL                  = 50
	bicepLexerSINGLE_LINE_COMMENT = 51
	bicepLexerMULTI_LINE_COMMENT  = 52
	bicepLexerSPACES              = 53
	bicepLexerUNKNOWN             = 54
)
