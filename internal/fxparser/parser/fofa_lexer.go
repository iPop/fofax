// Code generated from FOFA.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 14, 403,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 3, 5, 3, 39, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 58,
	10, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5,
	14, 385, 10, 14, 3, 15, 3, 15, 3, 15, 7, 15, 390, 10, 15, 12, 15, 14, 15,
	393, 11, 15, 3, 15, 3, 15, 3, 16, 6, 16, 398, 10, 16, 13, 16, 14, 16, 399,
	3, 16, 3, 16, 2, 2, 17, 3, 3, 5, 2, 7, 2, 9, 2, 11, 4, 13, 5, 15, 6, 17,
	7, 19, 8, 21, 9, 23, 10, 25, 11, 27, 12, 29, 13, 31, 14, 3, 2, 6, 10, 2,
	36, 36, 49, 49, 94, 94, 100, 100, 104, 104, 112, 112, 116, 116, 118, 118,
	5, 2, 50, 59, 67, 72, 99, 104, 4, 2, 36, 36, 94, 94, 4, 2, 11, 11, 34,
	34, 2, 448, 2, 3, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2,
	15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2,
	2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2,
	2, 2, 31, 3, 2, 2, 2, 3, 33, 3, 2, 2, 2, 5, 35, 3, 2, 2, 2, 7, 40, 3, 2,
	2, 2, 9, 46, 3, 2, 2, 2, 11, 57, 3, 2, 2, 2, 13, 59, 3, 2, 2, 2, 15, 62,
	3, 2, 2, 2, 17, 65, 3, 2, 2, 2, 19, 68, 3, 2, 2, 2, 21, 70, 3, 2, 2, 2,
	23, 73, 3, 2, 2, 2, 25, 75, 3, 2, 2, 2, 27, 384, 3, 2, 2, 2, 29, 386, 3,
	2, 2, 2, 31, 397, 3, 2, 2, 2, 33, 34, 7, 48, 2, 2, 34, 4, 3, 2, 2, 2, 35,
	38, 7, 94, 2, 2, 36, 39, 9, 2, 2, 2, 37, 39, 5, 7, 4, 2, 38, 36, 3, 2,
	2, 2, 38, 37, 3, 2, 2, 2, 39, 6, 3, 2, 2, 2, 40, 41, 7, 119, 2, 2, 41,
	42, 5, 9, 5, 2, 42, 43, 5, 9, 5, 2, 43, 44, 5, 9, 5, 2, 44, 45, 5, 9, 5,
	2, 45, 8, 3, 2, 2, 2, 46, 47, 9, 3, 2, 2, 47, 10, 3, 2, 2, 2, 48, 49, 7,
	118, 2, 2, 49, 50, 7, 116, 2, 2, 50, 51, 7, 119, 2, 2, 51, 58, 7, 103,
	2, 2, 52, 53, 7, 104, 2, 2, 53, 54, 7, 99, 2, 2, 54, 55, 7, 110, 2, 2,
	55, 56, 7, 117, 2, 2, 56, 58, 7, 103, 2, 2, 57, 48, 3, 2, 2, 2, 57, 52,
	3, 2, 2, 2, 58, 12, 3, 2, 2, 2, 59, 60, 7, 40, 2, 2, 60, 61, 7, 40, 2,
	2, 61, 14, 3, 2, 2, 2, 62, 63, 7, 126, 2, 2, 63, 64, 7, 126, 2, 2, 64,
	16, 3, 2, 2, 2, 65, 66, 7, 35, 2, 2, 66, 67, 7, 63, 2, 2, 67, 18, 3, 2,
	2, 2, 68, 69, 7, 63, 2, 2, 69, 20, 3, 2, 2, 2, 70, 71, 7, 63, 2, 2, 71,
	72, 7, 63, 2, 2, 72, 22, 3, 2, 2, 2, 73, 74, 7, 42, 2, 2, 74, 24, 3, 2,
	2, 2, 75, 76, 7, 43, 2, 2, 76, 26, 3, 2, 2, 2, 77, 78, 7, 118, 2, 2, 78,
	79, 7, 107, 2, 2, 79, 80, 7, 118, 2, 2, 80, 81, 7, 110, 2, 2, 81, 385,
	7, 103, 2, 2, 82, 83, 7, 106, 2, 2, 83, 84, 7, 103, 2, 2, 84, 85, 7, 99,
	2, 2, 85, 86, 7, 102, 2, 2, 86, 87, 7, 103, 2, 2, 87, 385, 7, 116, 2, 2,
	88, 89, 7, 100, 2, 2, 89, 90, 7, 113, 2, 2, 90, 91, 7, 102, 2, 2, 91, 385,
	7, 123, 2, 2, 92, 93, 7, 104, 2, 2, 93, 94, 7, 107, 2, 2, 94, 385, 7, 102,
	2, 2, 95, 96, 7, 102, 2, 2, 96, 97, 7, 113, 2, 2, 97, 98, 7, 111, 2, 2,
	98, 99, 7, 99, 2, 2, 99, 100, 7, 107, 2, 2, 100, 385, 7, 112, 2, 2, 101,
	102, 7, 107, 2, 2, 102, 103, 7, 101, 2, 2, 103, 385, 7, 114, 2, 2, 104,
	105, 7, 108, 2, 2, 105, 106, 7, 117, 2, 2, 106, 107, 7, 97, 2, 2, 107,
	108, 7, 112, 2, 2, 108, 109, 7, 99, 2, 2, 109, 110, 7, 111, 2, 2, 110,
	385, 7, 103, 2, 2, 111, 112, 7, 108, 2, 2, 112, 113, 7, 117, 2, 2, 113,
	114, 7, 97, 2, 2, 114, 115, 7, 111, 2, 2, 115, 116, 7, 102, 2, 2, 116,
	385, 7, 55, 2, 2, 117, 118, 7, 107, 2, 2, 118, 119, 7, 101, 2, 2, 119,
	120, 7, 113, 2, 2, 120, 121, 7, 112, 2, 2, 121, 122, 7, 97, 2, 2, 122,
	123, 7, 106, 2, 2, 123, 124, 7, 99, 2, 2, 124, 125, 7, 117, 2, 2, 125,
	385, 7, 106, 2, 2, 126, 127, 7, 106, 2, 2, 127, 128, 7, 113, 2, 2, 128,
	129, 7, 117, 2, 2, 129, 385, 7, 118, 2, 2, 130, 131, 7, 114, 2, 2, 131,
	132, 7, 113, 2, 2, 132, 133, 7, 116, 2, 2, 133, 385, 7, 118, 2, 2, 134,
	135, 7, 107, 2, 2, 135, 385, 7, 114, 2, 2, 136, 137, 7, 117, 2, 2, 137,
	138, 7, 118, 2, 2, 138, 139, 7, 99, 2, 2, 139, 140, 7, 118, 2, 2, 140,
	141, 7, 119, 2, 2, 141, 142, 7, 117, 2, 2, 142, 143, 7, 97, 2, 2, 143,
	144, 7, 101, 2, 2, 144, 145, 7, 113, 2, 2, 145, 146, 7, 102, 2, 2, 146,
	385, 7, 103, 2, 2, 147, 148, 7, 114, 2, 2, 148, 149, 7, 116, 2, 2, 149,
	150, 7, 113, 2, 2, 150, 151, 7, 118, 2, 2, 151, 152, 7, 113, 2, 2, 152,
	153, 7, 101, 2, 2, 153, 154, 7, 113, 2, 2, 154, 385, 7, 110, 2, 2, 155,
	156, 7, 101, 2, 2, 156, 157, 7, 113, 2, 2, 157, 158, 7, 119, 2, 2, 158,
	159, 7, 112, 2, 2, 159, 160, 7, 118, 2, 2, 160, 161, 7, 116, 2, 2, 161,
	385, 7, 123, 2, 2, 162, 163, 7, 116, 2, 2, 163, 164, 7, 103, 2, 2, 164,
	165, 7, 105, 2, 2, 165, 166, 7, 107, 2, 2, 166, 167, 7, 113, 2, 2, 167,
	385, 7, 112, 2, 2, 168, 169, 7, 101, 2, 2, 169, 170, 7, 107, 2, 2, 170,
	171, 7, 118, 2, 2, 171, 385, 7, 123, 2, 2, 172, 173, 7, 101, 2, 2, 173,
	174, 7, 103, 2, 2, 174, 175, 7, 116, 2, 2, 175, 385, 7, 118, 2, 2, 176,
	177, 7, 101, 2, 2, 177, 178, 7, 103, 2, 2, 178, 179, 7, 116, 2, 2, 179,
	180, 7, 118, 2, 2, 180, 181, 7, 48, 2, 2, 181, 182, 7, 117, 2, 2, 182,
	183, 7, 119, 2, 2, 183, 184, 7, 100, 2, 2, 184, 185, 7, 108, 2, 2, 185,
	186, 7, 103, 2, 2, 186, 187, 7, 101, 2, 2, 187, 385, 7, 118, 2, 2, 188,
	189, 7, 101, 2, 2, 189, 190, 7, 103, 2, 2, 190, 191, 7, 116, 2, 2, 191,
	192, 7, 118, 2, 2, 192, 193, 7, 48, 2, 2, 193, 194, 7, 107, 2, 2, 194,
	195, 7, 117, 2, 2, 195, 196, 7, 117, 2, 2, 196, 197, 7, 119, 2, 2, 197,
	198, 7, 103, 2, 2, 198, 385, 7, 116, 2, 2, 199, 200, 7, 101, 2, 2, 200,
	201, 7, 103, 2, 2, 201, 202, 7, 116, 2, 2, 202, 203, 7, 118, 2, 2, 203,
	204, 7, 48, 2, 2, 204, 205, 7, 107, 2, 2, 205, 206, 7, 117, 2, 2, 206,
	207, 7, 97, 2, 2, 207, 208, 7, 120, 2, 2, 208, 209, 7, 99, 2, 2, 209, 210,
	7, 110, 2, 2, 210, 211, 7, 107, 2, 2, 211, 385, 7, 102, 2, 2, 212, 213,
	7, 100, 2, 2, 213, 214, 7, 99, 2, 2, 214, 215, 7, 112, 2, 2, 215, 216,
	7, 112, 2, 2, 216, 217, 7, 103, 2, 2, 217, 385, 7, 116, 2, 2, 218, 219,
	7, 118, 2, 2, 219, 220, 7, 123, 2, 2, 220, 221, 7, 114, 2, 2, 221, 385,
	7, 103, 2, 2, 222, 223, 7, 113, 2, 2, 223, 385, 7, 117, 2, 2, 224, 225,
	7, 100, 2, 2, 225, 226, 7, 103, 2, 2, 226, 227, 7, 104, 2, 2, 227, 228,
	7, 113, 2, 2, 228, 229, 7, 116, 2, 2, 229, 385, 7, 103, 2, 2, 230, 231,
	7, 99, 2, 2, 231, 232, 7, 104, 2, 2, 232, 233, 7, 118, 2, 2, 233, 234,
	7, 103, 2, 2, 234, 385, 7, 116, 2, 2, 235, 236, 7, 113, 2, 2, 236, 237,
	7, 116, 2, 2, 237, 385, 7, 105, 2, 2, 238, 239, 7, 100, 2, 2, 239, 240,
	7, 99, 2, 2, 240, 241, 7, 117, 2, 2, 241, 242, 7, 103, 2, 2, 242, 243,
	7, 97, 2, 2, 243, 244, 7, 114, 2, 2, 244, 245, 7, 116, 2, 2, 245, 246,
	7, 113, 2, 2, 246, 247, 7, 118, 2, 2, 247, 248, 7, 113, 2, 2, 248, 249,
	7, 101, 2, 2, 249, 250, 7, 113, 2, 2, 250, 385, 7, 110, 2, 2, 251, 252,
	7, 117, 2, 2, 252, 253, 7, 103, 2, 2, 253, 254, 7, 116, 2, 2, 254, 255,
	7, 120, 2, 2, 255, 256, 7, 103, 2, 2, 256, 385, 7, 116, 2, 2, 257, 258,
	7, 99, 2, 2, 258, 259, 7, 114, 2, 2, 259, 385, 7, 114, 2, 2, 260, 261,
	7, 99, 2, 2, 261, 262, 7, 117, 2, 2, 262, 385, 7, 112, 2, 2, 263, 264,
	7, 107, 2, 2, 264, 265, 7, 117, 2, 2, 265, 266, 7, 97, 2, 2, 266, 267,
	7, 104, 2, 2, 267, 268, 7, 116, 2, 2, 268, 269, 7, 99, 2, 2, 269, 270,
	7, 119, 2, 2, 270, 385, 7, 102, 2, 2, 271, 272, 7, 107, 2, 2, 272, 273,
	7, 117, 2, 2, 273, 274, 7, 97, 2, 2, 274, 275, 7, 106, 2, 2, 275, 276,
	7, 113, 2, 2, 276, 277, 7, 112, 2, 2, 277, 278, 7, 103, 2, 2, 278, 279,
	7, 123, 2, 2, 279, 280, 7, 114, 2, 2, 280, 281, 7, 113, 2, 2, 281, 385,
	7, 118, 2, 2, 282, 283, 7, 107, 2, 2, 283, 284, 7, 117, 2, 2, 284, 285,
	7, 97, 2, 2, 285, 286, 7, 107, 2, 2, 286, 287, 7, 114, 2, 2, 287, 288,
	7, 120, 2, 2, 288, 385, 7, 56, 2, 2, 289, 290, 7, 107, 2, 2, 290, 291,
	7, 117, 2, 2, 291, 292, 7, 97, 2, 2, 292, 293, 7, 102, 2, 2, 293, 294,
	7, 113, 2, 2, 294, 295, 7, 111, 2, 2, 295, 296, 7, 99, 2, 2, 296, 297,
	7, 107, 2, 2, 297, 385, 7, 112, 2, 2, 298, 299, 7, 114, 2, 2, 299, 300,
	7, 113, 2, 2, 300, 301, 7, 116, 2, 2, 301, 302, 7, 118, 2, 2, 302, 303,
	7, 97, 2, 2, 303, 304, 7, 117, 2, 2, 304, 305, 7, 107, 2, 2, 305, 306,
	7, 124, 2, 2, 306, 385, 7, 103, 2, 2, 307, 308, 7, 114, 2, 2, 308, 309,
	7, 113, 2, 2, 309, 310, 7, 116, 2, 2, 310, 311, 7, 118, 2, 2, 311, 312,
	7, 97, 2, 2, 312, 313, 7, 117, 2, 2, 313, 314, 7, 107, 2, 2, 314, 315,
	7, 124, 2, 2, 315, 316, 7, 103, 2, 2, 316, 317, 7, 97, 2, 2, 317, 318,
	7, 105, 2, 2, 318, 385, 7, 118, 2, 2, 319, 320, 7, 114, 2, 2, 320, 321,
	7, 113, 2, 2, 321, 322, 7, 116, 2, 2, 322, 323, 7, 118, 2, 2, 323, 324,
	7, 97, 2, 2, 324, 325, 7, 117, 2, 2, 325, 326, 7, 107, 2, 2, 326, 327,
	7, 124, 2, 2, 327, 328, 7, 103, 2, 2, 328, 329, 7, 97, 2, 2, 329, 330,
	7, 110, 2, 2, 330, 385, 7, 118, 2, 2, 331, 332, 7, 107, 2, 2, 332, 333,
	7, 114, 2, 2, 333, 334, 7, 97, 2, 2, 334, 335, 7, 114, 2, 2, 335, 336,
	7, 113, 2, 2, 336, 337, 7, 116, 2, 2, 337, 338, 7, 118, 2, 2, 338, 385,
	7, 117, 2, 2, 339, 340, 7, 107, 2, 2, 340, 341, 7, 114, 2, 2, 341, 342,
	7, 97, 2, 2, 342, 343, 7, 101, 2, 2, 343, 344, 7, 113, 2, 2, 344, 345,
	7, 119, 2, 2, 345, 346, 7, 112, 2, 2, 346, 347, 7, 118, 2, 2, 347, 348,
	7, 116, 2, 2, 348, 385, 7, 123, 2, 2, 349, 350, 7, 107, 2, 2, 350, 351,
	7, 114, 2, 2, 351, 352, 7, 97, 2, 2, 352, 353, 7, 116, 2, 2, 353, 354,
	7, 103, 2, 2, 354, 355, 7, 105, 2, 2, 355, 356, 7, 107, 2, 2, 356, 357,
	7, 113, 2, 2, 357, 385, 7, 112, 2, 2, 358, 359, 7, 107, 2, 2, 359, 360,
	7, 114, 2, 2, 360, 361, 7, 97, 2, 2, 361, 362, 7, 101, 2, 2, 362, 363,
	7, 107, 2, 2, 363, 364, 7, 118, 2, 2, 364, 385, 7, 123, 2, 2, 365, 366,
	7, 107, 2, 2, 366, 367, 7, 114, 2, 2, 367, 368, 7, 97, 2, 2, 368, 369,
	7, 99, 2, 2, 369, 370, 7, 104, 2, 2, 370, 371, 7, 118, 2, 2, 371, 372,
	7, 103, 2, 2, 372, 385, 7, 116, 2, 2, 373, 374, 7, 107, 2, 2, 374, 375,
	7, 114, 2, 2, 375, 376, 7, 97, 2, 2, 376, 377, 7, 100, 2, 2, 377, 378,
	7, 103, 2, 2, 378, 379, 7, 104, 2, 2, 379, 380, 7, 113, 2, 2, 380, 381,
	7, 116, 2, 2, 381, 385, 7, 103, 2, 2, 382, 383, 7, 104, 2, 2, 383, 385,
	7, 122, 2, 2, 384, 77, 3, 2, 2, 2, 384, 82, 3, 2, 2, 2, 384, 88, 3, 2,
	2, 2, 384, 92, 3, 2, 2, 2, 384, 95, 3, 2, 2, 2, 384, 101, 3, 2, 2, 2, 384,
	104, 3, 2, 2, 2, 384, 111, 3, 2, 2, 2, 384, 117, 3, 2, 2, 2, 384, 126,
	3, 2, 2, 2, 384, 130, 3, 2, 2, 2, 384, 134, 3, 2, 2, 2, 384, 136, 3, 2,
	2, 2, 384, 147, 3, 2, 2, 2, 384, 155, 3, 2, 2, 2, 384, 162, 3, 2, 2, 2,
	384, 168, 3, 2, 2, 2, 384, 172, 3, 2, 2, 2, 384, 176, 3, 2, 2, 2, 384,
	188, 3, 2, 2, 2, 384, 199, 3, 2, 2, 2, 384, 212, 3, 2, 2, 2, 384, 218,
	3, 2, 2, 2, 384, 222, 3, 2, 2, 2, 384, 224, 3, 2, 2, 2, 384, 230, 3, 2,
	2, 2, 384, 235, 3, 2, 2, 2, 384, 238, 3, 2, 2, 2, 384, 251, 3, 2, 2, 2,
	384, 257, 3, 2, 2, 2, 384, 260, 3, 2, 2, 2, 384, 263, 3, 2, 2, 2, 384,
	271, 3, 2, 2, 2, 384, 282, 3, 2, 2, 2, 384, 289, 3, 2, 2, 2, 384, 298,
	3, 2, 2, 2, 384, 307, 3, 2, 2, 2, 384, 319, 3, 2, 2, 2, 384, 331, 3, 2,
	2, 2, 384, 339, 3, 2, 2, 2, 384, 349, 3, 2, 2, 2, 384, 358, 3, 2, 2, 2,
	384, 365, 3, 2, 2, 2, 384, 373, 3, 2, 2, 2, 384, 382, 3, 2, 2, 2, 385,
	28, 3, 2, 2, 2, 386, 391, 7, 36, 2, 2, 387, 390, 5, 5, 3, 2, 388, 390,
	10, 4, 2, 2, 389, 387, 3, 2, 2, 2, 389, 388, 3, 2, 2, 2, 390, 393, 3, 2,
	2, 2, 391, 389, 3, 2, 2, 2, 391, 392, 3, 2, 2, 2, 392, 394, 3, 2, 2, 2,
	393, 391, 3, 2, 2, 2, 394, 395, 7, 36, 2, 2, 395, 30, 3, 2, 2, 2, 396,
	398, 9, 5, 2, 2, 397, 396, 3, 2, 2, 2, 398, 399, 3, 2, 2, 2, 399, 397,
	3, 2, 2, 2, 399, 400, 3, 2, 2, 2, 400, 401, 3, 2, 2, 2, 401, 402, 8, 16,
	2, 2, 402, 32, 3, 2, 2, 2, 9, 2, 38, 57, 384, 389, 391, 399, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'.'", "", "'&&'", "'||'", "'!='", "'='", "'=='", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "", "BOOLEAN", "AND", "OR", "NOT", "EQ", "SEQ", "BR_OPEN", "BR_CLOSE",
	"FOFA_KEY", "STRING", "WS",
}

var lexerRuleNames = []string{
	"T__0", "ESC", "UNICODE", "HEX", "BOOLEAN", "AND", "OR", "NOT", "EQ", "SEQ",
	"BR_OPEN", "BR_CLOSE", "FOFA_KEY", "STRING", "WS",
}

type FOFALexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewFOFALexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *FOFALexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewFOFALexer(input antlr.CharStream) *FOFALexer {
	l := new(FOFALexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "FOFA.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// FOFALexer tokens.
const (
	FOFALexerT__0     = 1
	FOFALexerBOOLEAN  = 2
	FOFALexerAND      = 3
	FOFALexerOR       = 4
	FOFALexerNOT      = 5
	FOFALexerEQ       = 6
	FOFALexerSEQ      = 7
	FOFALexerBR_OPEN  = 8
	FOFALexerBR_CLOSE = 9
	FOFALexerFOFA_KEY = 10
	FOFALexerSTRING   = 11
	FOFALexerWS       = 12
)
