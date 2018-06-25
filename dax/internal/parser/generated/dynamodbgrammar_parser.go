// Code generated from DynamoDbGrammar.g4 by ANTLR 4.7.1. DO NOT EDIT.

package generated // DynamoDbGrammar

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)



//import com.amazon.dynamodb.grammar.exceptions.RedundantParenthesesException; JAVA


// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 73, 657, 
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9, 
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23, 
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4, 
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34, 
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9, 
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44, 
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4, 
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55, 
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 60, 9, 
	60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65, 9, 65, 
	4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9, 70, 4, 
	71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75, 4, 76, 
	9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4, 81, 9, 
	81, 4, 82, 9, 82, 4, 83, 9, 83, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 7, 
	3, 173, 10, 3, 12, 3, 14, 3, 176, 11, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 192, 10, 5, 
	12, 5, 14, 5, 195, 11, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 213, 10, 5, 3, 
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 221, 10, 5, 12, 5, 14, 5, 224, 11, 
	5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 6, 8, 235, 10, 
	8, 13, 8, 14, 8, 236, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 243, 10, 9, 12, 9, 
	14, 9, 246, 11, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 
	11, 7, 11, 256, 10, 11, 12, 11, 14, 11, 259, 11, 11, 3, 12, 3, 12, 3, 12, 
	3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 268, 10, 13, 12, 13, 14, 13, 271, 11, 
	13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 280, 10, 15, 
	12, 15, 14, 15, 283, 11, 15, 3, 16, 3, 16, 3, 17, 3, 17, 5, 17, 289, 10, 
	17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 
	300, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 
	19, 310, 10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20, 317, 10, 20, 
	12, 20, 14, 20, 320, 11, 20, 3, 20, 3, 20, 3, 21, 3, 21, 7, 21, 326, 10, 
	21, 12, 21, 14, 21, 329, 11, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 
	23, 3, 23, 5, 23, 338, 10, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 26, 
	3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 
	28, 3, 28, 3, 28, 3, 28, 5, 28, 360, 10, 28, 3, 29, 3, 29, 3, 29, 3, 29, 
	3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 
	32, 3, 32, 5, 32, 378, 10, 32, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 
	3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 
	37, 3, 37, 3, 37, 3, 37, 5, 37, 400, 10, 37, 3, 38, 3, 38, 3, 38, 3, 38, 
	3, 38, 5, 38, 407, 10, 38, 3, 39, 3, 39, 5, 39, 411, 10, 39, 3, 39, 3, 
	39, 5, 39, 415, 10, 39, 3, 39, 3, 39, 5, 39, 419, 10, 39, 3, 39, 5, 39, 
	422, 10, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 7, 40, 431, 
	10, 40, 12, 40, 14, 40, 434, 11, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 
	3, 40, 7, 40, 442, 10, 40, 12, 40, 14, 40, 445, 11, 40, 5, 40, 447, 10, 
	40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 
	5, 42, 459, 10, 42, 3, 42, 5, 42, 462, 10, 42, 3, 42, 5, 42, 465, 10, 42, 
	3, 43, 3, 43, 5, 43, 469, 10, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 7, 
	44, 476, 10, 44, 12, 44, 14, 44, 479, 11, 44, 3, 44, 3, 44, 3, 45, 3, 45, 
	5, 45, 485, 10, 45, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 
	48, 3, 49, 3, 49, 5, 49, 497, 10, 49, 3, 50, 3, 50, 3, 50, 5, 50, 502, 
	10, 50, 3, 51, 3, 51, 3, 52, 3, 52, 3, 53, 3, 53, 5, 53, 510, 10, 53, 3, 
	54, 3, 54, 3, 54, 3, 55, 3, 55, 3, 55, 3, 55, 5, 55, 519, 10, 55, 3, 56, 
	3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 5, 56, 529, 10, 56, 3, 
	57, 3, 57, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 60, 3, 60, 3, 60, 3, 60, 
	5, 60, 542, 10, 60, 3, 61, 5, 61, 545, 10, 61, 3, 61, 3, 61, 3, 61, 3, 
	61, 7, 61, 551, 10, 61, 12, 61, 14, 61, 554, 11, 61, 5, 61, 556, 10, 61, 
	3, 61, 3, 61, 3, 61, 5, 61, 561, 10, 61, 3, 61, 5, 61, 564, 10, 61, 3, 
	62, 3, 62, 3, 63, 3, 63, 3, 64, 3, 64, 5, 64, 572, 10, 64, 3, 65, 3, 65, 
	3, 65, 3, 65, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 7, 66, 583, 10, 66, 12, 
	66, 14, 66, 586, 11, 66, 3, 66, 3, 66, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 
	3, 67, 3, 67, 3, 68, 3, 68, 3, 69, 3, 69, 3, 70, 3, 70, 3, 70, 3, 70, 5, 
	70, 605, 10, 70, 3, 70, 3, 70, 3, 71, 3, 71, 3, 71, 3, 72, 3, 72, 3, 73, 
	3, 73, 3, 74, 3, 74, 3, 75, 3, 75, 3, 76, 3, 76, 3, 76, 3, 77, 3, 77, 3, 
	77, 3, 77, 3, 77, 3, 77, 3, 77, 3, 77, 3, 77, 5, 77, 632, 10, 77, 3, 78, 
	3, 78, 3, 78, 3, 78, 3, 78, 5, 78, 639, 10, 78, 3, 79, 3, 79, 3, 79, 3, 
	80, 3, 80, 3, 81, 3, 81, 5, 81, 648, 10, 81, 3, 82, 3, 82, 3, 83, 6, 83, 
	653, 10, 83, 13, 83, 14, 83, 654, 3, 83, 2, 3, 8, 84, 2, 4, 6, 8, 10, 12, 
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 
	50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 
	86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 
	118, 120, 122, 124, 126, 128, 130, 132, 134, 136, 138, 140, 142, 144, 146, 
	148, 150, 152, 154, 156, 158, 160, 162, 164, 2, 8, 3, 2, 11, 16, 3, 2, 
	17, 18, 3, 2, 69, 70, 3, 2, 30, 31, 3, 2, 52, 53, 3, 2, 47, 49, 2, 647, 
	2, 166, 3, 2, 2, 2, 4, 169, 3, 2, 2, 2, 6, 177, 3, 2, 2, 2, 8, 212, 3, 
	2, 2, 2, 10, 225, 3, 2, 2, 2, 12, 227, 3, 2, 2, 2, 14, 234, 3, 2, 2, 2, 
	16, 238, 3, 2, 2, 2, 18, 247, 3, 2, 2, 2, 20, 251, 3, 2, 2, 2, 22, 260, 
	3, 2, 2, 2, 24, 263, 3, 2, 2, 2, 26, 272, 3, 2, 2, 2, 28, 275, 3, 2, 2, 
	2, 30, 284, 3, 2, 2, 2, 32, 288, 3, 2, 2, 2, 34, 299, 3, 2, 2, 2, 36, 309, 
	3, 2, 2, 2, 38, 311, 3, 2, 2, 2, 40, 323, 3, 2, 2, 2, 42, 330, 3, 2, 2, 
	2, 44, 337, 3, 2, 2, 2, 46, 339, 3, 2, 2, 2, 48, 341, 3, 2, 2, 2, 50, 344, 
	3, 2, 2, 2, 52, 347, 3, 2, 2, 2, 54, 359, 3, 2, 2, 2, 56, 361, 3, 2, 2, 
	2, 58, 365, 3, 2, 2, 2, 60, 368, 3, 2, 2, 2, 62, 377, 3, 2, 2, 2, 64, 379, 
	3, 2, 2, 2, 66, 382, 3, 2, 2, 2, 68, 385, 3, 2, 2, 2, 70, 389, 3, 2, 2, 
	2, 72, 394, 3, 2, 2, 2, 74, 401, 3, 2, 2, 2, 76, 408, 3, 2, 2, 2, 78, 423, 
	3, 2, 2, 2, 80, 450, 3, 2, 2, 2, 82, 453, 3, 2, 2, 2, 84, 468, 3, 2, 2, 
	2, 86, 470, 3, 2, 2, 2, 88, 484, 3, 2, 2, 2, 90, 486, 3, 2, 2, 2, 92, 488, 
	3, 2, 2, 2, 94, 492, 3, 2, 2, 2, 96, 496, 3, 2, 2, 2, 98, 501, 3, 2, 2, 
	2, 100, 503, 3, 2, 2, 2, 102, 505, 3, 2, 2, 2, 104, 507, 3, 2, 2, 2, 106, 
	511, 3, 2, 2, 2, 108, 518, 3, 2, 2, 2, 110, 520, 3, 2, 2, 2, 112, 530, 
	3, 2, 2, 2, 114, 532, 3, 2, 2, 2, 116, 534, 3, 2, 2, 2, 118, 537, 3, 2, 
	2, 2, 120, 544, 3, 2, 2, 2, 122, 565, 3, 2, 2, 2, 124, 567, 3, 2, 2, 2, 
	126, 571, 3, 2, 2, 2, 128, 573, 3, 2, 2, 2, 130, 577, 3, 2, 2, 2, 132, 
	589, 3, 2, 2, 2, 134, 596, 3, 2, 2, 2, 136, 598, 3, 2, 2, 2, 138, 600, 
	3, 2, 2, 2, 140, 608, 3, 2, 2, 2, 142, 611, 3, 2, 2, 2, 144, 613, 3, 2, 
	2, 2, 146, 615, 3, 2, 2, 2, 148, 617, 3, 2, 2, 2, 150, 619, 3, 2, 2, 2, 
	152, 631, 3, 2, 2, 2, 154, 633, 3, 2, 2, 2, 156, 640, 3, 2, 2, 2, 158, 
	643, 3, 2, 2, 2, 160, 647, 3, 2, 2, 2, 162, 649, 3, 2, 2, 2, 164, 652, 
	3, 2, 2, 2, 166, 167, 5, 4, 3, 2, 167, 168, 7, 2, 2, 3, 168, 3, 3, 2, 2, 
	2, 169, 174, 5, 40, 21, 2, 170, 171, 7, 3, 2, 2, 171, 173, 5, 40, 21, 2, 
	172, 170, 3, 2, 2, 2, 173, 176, 3, 2, 2, 2, 174, 172, 3, 2, 2, 2, 174, 
	175, 3, 2, 2, 2, 175, 5, 3, 2, 2, 2, 176, 174, 3, 2, 2, 2, 177, 178, 5, 
	8, 5, 2, 178, 179, 7, 2, 2, 3, 179, 7, 3, 2, 2, 2, 180, 181, 8, 5, 1, 2, 
	181, 182, 5, 36, 19, 2, 182, 183, 5, 10, 6, 2, 183, 184, 5, 36, 19, 2, 
	184, 213, 3, 2, 2, 2, 185, 186, 5, 36, 19, 2, 186, 187, 7, 19, 2, 2, 187, 
	188, 7, 4, 2, 2, 188, 193, 5, 36, 19, 2, 189, 190, 7, 3, 2, 2, 190, 192, 
	5, 36, 19, 2, 191, 189, 3, 2, 2, 2, 192, 195, 3, 2, 2, 2, 193, 191, 3, 
	2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 196, 3, 2, 2, 2, 195, 193, 3, 2, 2, 
	2, 196, 197, 7, 5, 2, 2, 197, 213, 3, 2, 2, 2, 198, 199, 5, 36, 19, 2, 
	199, 200, 7, 20, 2, 2, 200, 201, 5, 36, 19, 2, 201, 202, 7, 22, 2, 2, 202, 
	203, 5, 36, 19, 2, 203, 213, 3, 2, 2, 2, 204, 213, 5, 38, 20, 2, 205, 206, 
	7, 4, 2, 2, 206, 207, 5, 8, 5, 2, 207, 208, 7, 5, 2, 2, 208, 209, 8, 5, 
	1, 2, 209, 213, 3, 2, 2, 2, 210, 211, 7, 21, 2, 2, 211, 213, 5, 8, 5, 5, 
	212, 180, 3, 2, 2, 2, 212, 185, 3, 2, 2, 2, 212, 198, 3, 2, 2, 2, 212, 
	204, 3, 2, 2, 2, 212, 205, 3, 2, 2, 2, 212, 210, 3, 2, 2, 2, 213, 222, 
	3, 2, 2, 2, 214, 215, 12, 4, 2, 2, 215, 216, 7, 22, 2, 2, 216, 221, 5, 
	8, 5, 5, 217, 218, 12, 3, 2, 2, 218, 219, 7, 23, 2, 2, 219, 221, 5, 8, 
	5, 4, 220, 214, 3, 2, 2, 2, 220, 217, 3, 2, 2, 2, 221, 224, 3, 2, 2, 2, 
	222, 220, 3, 2, 2, 2, 222, 223, 3, 2, 2, 2, 223, 9, 3, 2, 2, 2, 224, 222, 
	3, 2, 2, 2, 225, 226, 9, 2, 2, 2, 226, 11, 3, 2, 2, 2, 227, 228, 5, 14, 
	8, 2, 228, 229, 7, 2, 2, 3, 229, 13, 3, 2, 2, 2, 230, 235, 5, 16, 9, 2, 
	231, 235, 5, 20, 11, 2, 232, 235, 5, 24, 13, 2, 233, 235, 5, 28, 15, 2, 
	234, 230, 3, 2, 2, 2, 234, 231, 3, 2, 2, 2, 234, 232, 3, 2, 2, 2, 234, 
	233, 3, 2, 2, 2, 235, 236, 3, 2, 2, 2, 236, 234, 3, 2, 2, 2, 236, 237, 
	3, 2, 2, 2, 237, 15, 3, 2, 2, 2, 238, 239, 7, 24, 2, 2, 239, 244, 5, 18, 
	10, 2, 240, 241, 7, 3, 2, 2, 241, 243, 5, 18, 10, 2, 242, 240, 3, 2, 2, 
	2, 243, 246, 3, 2, 2, 2, 244, 242, 3, 2, 2, 2, 244, 245, 3, 2, 2, 2, 245, 
	17, 3, 2, 2, 2, 246, 244, 3, 2, 2, 2, 247, 248, 5, 40, 21, 2, 248, 249, 
	7, 11, 2, 2, 249, 250, 5, 32, 17, 2, 250, 19, 3, 2, 2, 2, 251, 252, 7, 
	25, 2, 2, 252, 257, 5, 22, 12, 2, 253, 254, 7, 3, 2, 2, 254, 256, 5, 22, 
	12, 2, 255, 253, 3, 2, 2, 2, 256, 259, 3, 2, 2, 2, 257, 255, 3, 2, 2, 2, 
	257, 258, 3, 2, 2, 2, 258, 21, 3, 2, 2, 2, 259, 257, 3, 2, 2, 2, 260, 261, 
	5, 40, 21, 2, 261, 262, 5, 46, 24, 2, 262, 23, 3, 2, 2, 2, 263, 264, 7, 
	26, 2, 2, 264, 269, 5, 26, 14, 2, 265, 266, 7, 3, 2, 2, 266, 268, 5, 26, 
	14, 2, 267, 265, 3, 2, 2, 2, 268, 271, 3, 2, 2, 2, 269, 267, 3, 2, 2, 2, 
	269, 270, 3, 2, 2, 2, 270, 25, 3, 2, 2, 2, 271, 269, 3, 2, 2, 2, 272, 273, 
	5, 40, 21, 2, 273, 274, 5, 46, 24, 2, 274, 27, 3, 2, 2, 2, 275, 276, 7, 
	27, 2, 2, 276, 281, 5, 30, 16, 2, 277, 278, 7, 3, 2, 2, 278, 280, 5, 30, 
	16, 2, 279, 277, 3, 2, 2, 2, 280, 283, 3, 2, 2, 2, 281, 279, 3, 2, 2, 2, 
	281, 282, 3, 2, 2, 2, 282, 29, 3, 2, 2, 2, 283, 281, 3, 2, 2, 2, 284, 285, 
	5, 40, 21, 2, 285, 31, 3, 2, 2, 2, 286, 289, 5, 36, 19, 2, 287, 289, 5, 
	34, 18, 2, 288, 286, 3, 2, 2, 2, 288, 287, 3, 2, 2, 2, 289, 33, 3, 2, 2, 
	2, 290, 291, 5, 36, 19, 2, 291, 292, 9, 3, 2, 2, 292, 293, 5, 36, 19, 2, 
	293, 300, 3, 2, 2, 2, 294, 295, 7, 4, 2, 2, 295, 296, 5, 34, 18, 2, 296, 
	297, 7, 5, 2, 2, 297, 298, 8, 18, 1, 2, 298, 300, 3, 2, 2, 2, 299, 290, 
	3, 2, 2, 2, 299, 294, 3, 2, 2, 2, 300, 35, 3, 2, 2, 2, 301, 310, 5, 40, 
	21, 2, 302, 310, 5, 46, 24, 2, 303, 310, 5, 38, 20, 2, 304, 305, 7, 4, 
	2, 2, 305, 306, 5, 36, 19, 2, 306, 307, 7, 5, 2, 2, 307, 308, 8, 19, 1, 
	2, 308, 310, 3, 2, 2, 2, 309, 301, 3, 2, 2, 2, 309, 302, 3, 2, 2, 2, 309, 
	303, 3, 2, 2, 2, 309, 304, 3, 2, 2, 2, 310, 37, 3, 2, 2, 2, 311, 312, 7, 
	69, 2, 2, 312, 313, 7, 4, 2, 2, 313, 318, 5, 36, 19, 2, 314, 315, 7, 3, 
	2, 2, 315, 317, 5, 36, 19, 2, 316, 314, 3, 2, 2, 2, 317, 320, 3, 2, 2, 
	2, 318, 316, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2, 319, 321, 3, 2, 2, 2, 320, 
	318, 3, 2, 2, 2, 321, 322, 7, 5, 2, 2, 322, 39, 3, 2, 2, 2, 323, 327, 5, 
	42, 22, 2, 324, 326, 5, 44, 23, 2, 325, 324, 3, 2, 2, 2, 326, 329, 3, 2, 
	2, 2, 327, 325, 3, 2, 2, 2, 327, 328, 3, 2, 2, 2, 328, 41, 3, 2, 2, 2, 
	329, 327, 3, 2, 2, 2, 330, 331, 9, 4, 2, 2, 331, 43, 3, 2, 2, 2, 332, 333, 
	7, 6, 2, 2, 333, 338, 5, 42, 22, 2, 334, 335, 7, 7, 2, 2, 335, 336, 7, 
	68, 2, 2, 336, 338, 7, 8, 2, 2, 337, 332, 3, 2, 2, 2, 337, 334, 3, 2, 2, 
	2, 338, 45, 3, 2, 2, 2, 339, 340, 7, 71, 2, 2, 340, 47, 3, 2, 2, 2, 341, 
	342, 7, 70, 2, 2, 342, 343, 7, 2, 2, 3, 343, 49, 3, 2, 2, 2, 344, 345, 
	7, 71, 2, 2, 345, 346, 7, 2, 2, 3, 346, 51, 3, 2, 2, 2, 347, 348, 5, 54, 
	28, 2, 348, 349, 7, 2, 2, 3, 349, 53, 3, 2, 2, 2, 350, 360, 5, 76, 39, 
	2, 351, 360, 5, 78, 40, 2, 352, 360, 5, 80, 41, 2, 353, 360, 5, 72, 37, 
	2, 354, 360, 5, 74, 38, 2, 355, 360, 5, 82, 42, 2, 356, 360, 5, 56, 29, 
	2, 357, 360, 5, 60, 31, 2, 358, 360, 5, 58, 30, 2, 359, 350, 3, 2, 2, 2, 
	359, 351, 3, 2, 2, 2, 359, 352, 3, 2, 2, 2, 359, 353, 3, 2, 2, 2, 359, 
	354, 3, 2, 2, 2, 359, 355, 3, 2, 2, 2, 359, 356, 3, 2, 2, 2, 359, 357, 
	3, 2, 2, 2, 359, 358, 3, 2, 2, 2, 360, 55, 3, 2, 2, 2, 361, 362, 7, 64, 
	2, 2, 362, 363, 7, 45, 2, 2, 363, 364, 5, 158, 80, 2, 364, 57, 3, 2, 2, 
	2, 365, 366, 7, 66, 2, 2, 366, 367, 5, 158, 80, 2, 367, 59, 3, 2, 2, 2, 
	368, 369, 7, 65, 2, 2, 369, 370, 7, 45, 2, 2, 370, 371, 5, 158, 80, 2, 
	371, 372, 5, 62, 32, 2, 372, 61, 3, 2, 2, 2, 373, 378, 5, 64, 33, 2, 374, 
	378, 5, 66, 34, 2, 375, 378, 5, 68, 35, 2, 376, 378, 5, 70, 36, 2, 377, 
	373, 3, 2, 2, 2, 377, 374, 3, 2, 2, 2, 377, 375, 3, 2, 2, 2, 377, 376, 
	3, 2, 2, 2, 378, 63, 3, 2, 2, 2, 379, 380, 7, 24, 2, 2, 380, 381, 5, 132, 
	67, 2, 381, 65, 3, 2, 2, 2, 382, 383, 7, 25, 2, 2, 383, 384, 5, 120, 61, 
	2, 384, 67, 3, 2, 2, 2, 385, 386, 7, 64, 2, 2, 386, 387, 7, 54, 2, 2, 387, 
	388, 5, 124, 63, 2, 388, 69, 3, 2, 2, 2, 389, 390, 7, 65, 2, 2, 390, 391, 
	7, 54, 2, 2, 391, 392, 5, 124, 63, 2, 392, 393, 5, 64, 33, 2, 393, 71, 
	3, 2, 2, 2, 394, 395, 7, 32, 2, 2, 395, 396, 5, 158, 80, 2, 396, 397, 5, 
	14, 8, 2, 397, 399, 5, 116, 59, 2, 398, 400, 5, 150, 76, 2, 399, 398, 3, 
	2, 2, 2, 399, 400, 3, 2, 2, 2, 400, 73, 3, 2, 2, 2, 401, 402, 7, 26, 2, 
	2, 402, 403, 7, 59, 2, 2, 403, 404, 5, 158, 80, 2, 404, 406, 5, 116, 59, 
	2, 405, 407, 5, 150, 76, 2, 406, 405, 3, 2, 2, 2, 406, 407, 3, 2, 2, 2, 
	407, 75, 3, 2, 2, 2, 408, 410, 7, 28, 2, 2, 409, 411, 7, 29, 2, 2, 410, 
	409, 3, 2, 2, 2, 410, 411, 3, 2, 2, 2, 411, 412, 3, 2, 2, 2, 412, 414, 
	5, 158, 80, 2, 413, 415, 9, 5, 2, 2, 414, 413, 3, 2, 2, 2, 414, 415, 3, 
	2, 2, 2, 415, 416, 3, 2, 2, 2, 416, 418, 5, 46, 24, 2, 417, 419, 5, 154, 
	78, 2, 418, 417, 3, 2, 2, 2, 418, 419, 3, 2, 2, 2, 419, 421, 3, 2, 2, 2, 
	420, 422, 5, 150, 76, 2, 421, 420, 3, 2, 2, 2, 421, 422, 3, 2, 2, 2, 422, 
	77, 3, 2, 2, 2, 423, 424, 7, 44, 2, 2, 424, 425, 7, 45, 2, 2, 425, 426, 
	5, 158, 80, 2, 426, 427, 7, 4, 2, 2, 427, 432, 5, 140, 71, 2, 428, 429, 
	7, 3, 2, 2, 429, 431, 5, 140, 71, 2, 430, 428, 3, 2, 2, 2, 431, 434, 3, 
	2, 2, 2, 432, 430, 3, 2, 2, 2, 432, 433, 3, 2, 2, 2, 433, 435, 3, 2, 2, 
	2, 434, 432, 3, 2, 2, 2, 435, 436, 7, 3, 2, 2, 436, 446, 5, 118, 60, 2, 
	437, 438, 7, 3, 2, 2, 438, 443, 5, 120, 61, 2, 439, 440, 7, 3, 2, 2, 440, 
	442, 5, 120, 61, 2, 441, 439, 3, 2, 2, 2, 442, 445, 3, 2, 2, 2, 443, 441, 
	3, 2, 2, 2, 443, 444, 3, 2, 2, 2, 444, 447, 3, 2, 2, 2, 445, 443, 3, 2, 
	2, 2, 446, 437, 3, 2, 2, 2, 446, 447, 3, 2, 2, 2, 447, 448, 3, 2, 2, 2, 
	448, 449, 7, 5, 2, 2, 449, 79, 3, 2, 2, 2, 450, 451, 7, 56, 2, 2, 451, 
	452, 7, 57, 2, 2, 452, 81, 3, 2, 2, 2, 453, 454, 7, 58, 2, 2, 454, 455, 
	5, 84, 43, 2, 455, 456, 7, 59, 2, 2, 456, 458, 5, 158, 80, 2, 457, 459, 
	5, 104, 53, 2, 458, 457, 3, 2, 2, 2, 458, 459, 3, 2, 2, 2, 459, 461, 3, 
	2, 2, 2, 460, 462, 5, 116, 59, 2, 461, 460, 3, 2, 2, 2, 461, 462, 3, 2, 
	2, 2, 462, 464, 3, 2, 2, 2, 463, 465, 5, 86, 44, 2, 464, 463, 3, 2, 2, 
	2, 464, 465, 3, 2, 2, 2, 465, 83, 3, 2, 2, 2, 466, 469, 5, 4, 3, 2, 467, 
	469, 5, 102, 52, 2, 468, 466, 3, 2, 2, 2, 468, 467, 3, 2, 2, 2, 469, 85, 
	3, 2, 2, 2, 470, 471, 7, 67, 2, 2, 471, 472, 7, 4, 2, 2, 472, 477, 5, 88, 
	45, 2, 473, 474, 7, 3, 2, 2, 474, 476, 5, 88, 45, 2, 475, 473, 3, 2, 2, 
	2, 476, 479, 3, 2, 2, 2, 477, 475, 3, 2, 2, 2, 477, 478, 3, 2, 2, 2, 478, 
	480, 3, 2, 2, 2, 479, 477, 3, 2, 2, 2, 480, 481, 7, 5, 2, 2, 481, 87, 3, 
	2, 2, 2, 482, 485, 5, 90, 46, 2, 483, 485, 5, 92, 47, 2, 484, 482, 3, 2, 
	2, 2, 484, 483, 3, 2, 2, 2, 485, 89, 3, 2, 2, 2, 486, 487, 5, 94, 48, 2, 
	487, 91, 3, 2, 2, 2, 488, 489, 5, 94, 48, 2, 489, 490, 7, 11, 2, 2, 490, 
	491, 5, 96, 49, 2, 491, 93, 3, 2, 2, 2, 492, 493, 5, 160, 81, 2, 493, 95, 
	3, 2, 2, 2, 494, 497, 5, 98, 50, 2, 495, 497, 5, 100, 51, 2, 496, 494, 
	3, 2, 2, 2, 496, 495, 3, 2, 2, 2, 497, 97, 3, 2, 2, 2, 498, 502, 5, 162, 
	82, 2, 499, 502, 7, 40, 2, 2, 500, 502, 7, 69, 2, 2, 501, 498, 3, 2, 2, 
	2, 501, 499, 3, 2, 2, 2, 501, 500, 3, 2, 2, 2, 502, 99, 3, 2, 2, 2, 503, 
	504, 7, 68, 2, 2, 504, 101, 3, 2, 2, 2, 505, 506, 7, 9, 2, 2, 506, 103, 
	3, 2, 2, 2, 507, 509, 5, 106, 54, 2, 508, 510, 5, 110, 56, 2, 509, 508, 
	3, 2, 2, 2, 509, 510, 3, 2, 2, 2, 510, 105, 3, 2, 2, 2, 511, 512, 7, 60, 
	2, 2, 512, 513, 5, 108, 55, 2, 513, 107, 3, 2, 2, 2, 514, 515, 7, 54, 2, 
	2, 515, 519, 5, 124, 63, 2, 516, 517, 7, 46, 2, 2, 517, 519, 7, 54, 2, 
	2, 518, 514, 3, 2, 2, 2, 518, 516, 3, 2, 2, 2, 519, 109, 3, 2, 2, 2, 520, 
	521, 7, 61, 2, 2, 521, 528, 7, 62, 2, 2, 522, 523, 7, 4, 2, 2, 523, 524, 
	5, 112, 57, 2, 524, 525, 7, 3, 2, 2, 525, 526, 5, 114, 58, 2, 526, 527, 
	7, 5, 2, 2, 527, 529, 3, 2, 2, 2, 528, 522, 3, 2, 2, 2, 528, 529, 3, 2, 
	2, 2, 529, 111, 3, 2, 2, 2, 530, 531, 7, 68, 2, 2, 531, 113, 3, 2, 2, 2, 
	532, 533, 7, 68, 2, 2, 533, 115, 3, 2, 2, 2, 534, 535, 7, 63, 2, 2, 535, 
	536, 5, 8, 5, 2, 536, 117, 3, 2, 2, 2, 537, 538, 7, 46, 2, 2, 538, 539, 
	7, 37, 2, 2, 539, 541, 5, 138, 70, 2, 540, 542, 5, 132, 67, 2, 541, 540, 
	3, 2, 2, 2, 541, 542, 3, 2, 2, 2, 542, 119, 3, 2, 2, 2, 543, 545, 5, 122, 
	62, 2, 544, 543, 3, 2, 2, 2, 544, 545, 3, 2, 2, 2, 545, 546, 3, 2, 2, 2, 
	546, 555, 7, 54, 2, 2, 547, 552, 5, 140, 71, 2, 548, 549, 7, 3, 2, 2, 549, 
	551, 5, 140, 71, 2, 550, 548, 3, 2, 2, 2, 551, 554, 3, 2, 2, 2, 552, 550, 
	3, 2, 2, 2, 552, 553, 3, 2, 2, 2, 553, 556, 3, 2, 2, 2, 554, 552, 3, 2, 
	2, 2, 555, 547, 3, 2, 2, 2, 555, 556, 3, 2, 2, 2, 556, 557, 3, 2, 2, 2, 
	557, 558, 5, 124, 63, 2, 558, 560, 5, 138, 70, 2, 559, 561, 5, 126, 64, 
	2, 560, 559, 3, 2, 2, 2, 560, 561, 3, 2, 2, 2, 561, 563, 3, 2, 2, 2, 562, 
	564, 5, 132, 67, 2, 563, 562, 3, 2, 2, 2, 563, 564, 3, 2, 2, 2, 564, 121, 
	3, 2, 2, 2, 565, 566, 9, 6, 2, 2, 566, 123, 3, 2, 2, 2, 567, 568, 5, 160, 
	81, 2, 568, 125, 3, 2, 2, 2, 569, 572, 5, 128, 65, 2, 570, 572, 5, 130, 
	66, 2, 571, 569, 3, 2, 2, 2, 571, 570, 3, 2, 2, 2, 572, 127, 3, 2, 2, 2, 
	573, 574, 7, 55, 2, 2, 574, 575, 7, 36, 2, 2, 575, 576, 7, 38, 2, 2, 576, 
	129, 3, 2, 2, 2, 577, 578, 7, 55, 2, 2, 578, 579, 7, 4, 2, 2, 579, 584, 
	5, 146, 74, 2, 580, 581, 7, 3, 2, 2, 581, 583, 5, 146, 74, 2, 582, 580, 
	3, 2, 2, 2, 583, 586, 3, 2, 2, 2, 584, 582, 3, 2, 2, 2, 584, 585, 3, 2, 
	2, 2, 585, 587, 3, 2, 2, 2, 586, 584, 3, 2, 2, 2, 587, 588, 7, 5, 2, 2, 
	588, 131, 3, 2, 2, 2, 589, 590, 7, 51, 2, 2, 590, 591, 7, 4, 2, 2, 591, 
	592, 5, 134, 68, 2, 592, 593, 7, 3, 2, 2, 593, 594, 5, 136, 69, 2, 594, 
	595, 7, 5, 2, 2, 595, 133, 3, 2, 2, 2, 596, 597, 7, 68, 2, 2, 597, 135, 
	3, 2, 2, 2, 598, 599, 7, 68, 2, 2, 599, 137, 3, 2, 2, 2, 600, 601, 7, 4, 
	2, 2, 601, 604, 5, 142, 72, 2, 602, 603, 7, 3, 2, 2, 603, 605, 5, 144, 
	73, 2, 604, 602, 3, 2, 2, 2, 604, 605, 3, 2, 2, 2, 605, 606, 3, 2, 2, 2, 
	606, 607, 7, 5, 2, 2, 607, 139, 3, 2, 2, 2, 608, 609, 5, 146, 74, 2, 609, 
	610, 5, 148, 75, 2, 610, 141, 3, 2, 2, 2, 611, 612, 5, 160, 81, 2, 612, 
	143, 3, 2, 2, 2, 613, 614, 5, 160, 81, 2, 614, 145, 3, 2, 2, 2, 615, 616, 
	5, 160, 81, 2, 616, 147, 3, 2, 2, 2, 617, 618, 9, 7, 2, 2, 618, 149, 3, 
	2, 2, 2, 619, 620, 7, 34, 2, 2, 620, 621, 5, 152, 77, 2, 621, 151, 3, 2, 
	2, 2, 622, 632, 7, 40, 2, 2, 623, 624, 7, 35, 2, 2, 624, 632, 7, 41, 2, 
	2, 625, 626, 7, 42, 2, 2, 626, 632, 7, 41, 2, 2, 627, 628, 7, 35, 2, 2, 
	628, 632, 7, 39, 2, 2, 629, 630, 7, 42, 2, 2, 630, 632, 7, 39, 2, 2, 631, 
	622, 3, 2, 2, 2, 631, 623, 3, 2, 2, 2, 631, 625, 3, 2, 2, 2, 631, 627, 
	3, 2, 2, 2, 631, 629, 3, 2, 2, 2, 632, 153, 3, 2, 2, 2, 633, 634, 7, 33, 
	2, 2, 634, 635, 7, 43, 2, 2, 635, 636, 7, 37, 2, 2, 636, 638, 7, 32, 2, 
	2, 637, 639, 5, 156, 79, 2, 638, 637, 3, 2, 2, 2, 638, 639, 3, 2, 2, 2, 
	639, 155, 3, 2, 2, 2, 640, 641, 7, 50, 2, 2, 641, 642, 5, 8, 5, 2, 642, 
	157, 3, 2, 2, 2, 643, 644, 5, 160, 81, 2, 644, 159, 3, 2, 2, 2, 645, 648, 
	7, 69, 2, 2, 646, 648, 5, 162, 82, 2, 647, 645, 3, 2, 2, 2, 647, 646, 3, 
	2, 2, 2, 648, 161, 3, 2, 2, 2, 649, 650, 7, 72, 2, 2, 650, 163, 3, 2, 2, 
	2, 651, 653, 7, 73, 2, 2, 652, 651, 3, 2, 2, 2, 653, 654, 3, 2, 2, 2, 654, 
	652, 3, 2, 2, 2, 654, 655, 3, 2, 2, 2, 655, 165, 3, 2, 2, 2, 54, 174, 193, 
	212, 220, 222, 234, 236, 244, 257, 269, 281, 288, 299, 309, 318, 327, 337, 
	359, 377, 399, 406, 410, 414, 418, 421, 432, 443, 446, 458, 461, 464, 468, 
	477, 484, 496, 501, 509, 518, 528, 541, 544, 552, 555, 560, 563, 571, 584, 
	604, 631, 638, 647, 654,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "','", "'('", "')'", "'.'", "'['", "']'", "'*'", "", "'='", "'<>'", 
	"'<'", "'<='", "'>'", "'>='", "'+'", "'-'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "WS", "EQ", "NE", "LT", "LE", "GT", "GE", 
	"PLUS", "MINUS", "IN", "BETWEEN", "NOT", "AND", "OR", "SET", "ADD", "DELETE", 
	"REMOVE", "INSERT", "INTO", "VALUE", "VALUES", "UPDATE", "ON", "RETURNING", 
	"ALL", "KEYS", "KEY", "ONLY", "NEW", "NONE", "OLD", "UPDATED", "DUPLICATE", 
	"CREATE", "TABLE", "PRIMARY", "NUMBER", "BINARY", "STRING", "IF", "CAPACITY", 
	"GLOBAL", "LOCAL", "INDEXKEYWORD", "PROJECTION", "SHOW", "TABLES", "SELECT", 
	"FROM", "USE", "ENABLE", "SCAN", "WHERE", "DROP", "ALTER", "DESCRIBE", 
	"OPTION", "INDEX", "ID", "ATTRIBUTE_NAME_SUB", "LITERAL_SUB", "STRING_LITERAL", 
	"UNKNOWN",
}

var ruleNames = []string{
	"projection_", "projection", "condition_", "condition", "comparator_symbol", 
	"update_", "update", "set_section", "set_action", "add_section", "add_action", 
	"delete_section", "delete_action", "remove_section", "remove_action", "set_value", 
	"arithmetic", "operand", "function", "path", "id", "dereference", "literal", 
	"expression_attr_names_sub", "expression_attr_values_sub", "statement_", 
	"statement", "dropTableStatement", "describeStatement", "alterTableStatement", 
	"alterTableStatementType", "setCapacity", "addIndex", "dropIndex", "alterIndex", 
	"updateStatement", "deleteStatement", "insertStatement", "createTableStatement", 
	"showTablesStatement", "selectStatement", "selectProjection", "optionBlock", 
	"option", "singleOption", "keyValueOption", "optionKey", "optionValue", 
	"optionValueString", "optionValueNumber", "star", "hint", "indexHint", 
	"indexHintName", "scanInfo", "totalSegment", "segment", "where", "primaryKeyDecl", 
	"secondaryIndexDecl", "secondaryIndexType", "indexName", "projectionIndex", 
	"projectionIndexKeysOnly", "projectionIndexVector", "capacity", "readUnits", 
	"writeUnits", "indexDecl", "attributeDecl", "hashKey", "rangeKey", "attributeName", 
	"attributeType", "returning", "returningValue", "onDuplicateKeyUpdate", 
	"ifClause", "tableName", "ddlName", "stringLiteral", "unknown",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type DynamoDbGrammarParser struct {
	*antlr.BaseParser
}

func NewDynamoDbGrammarParser(input antlr.TokenStream) *DynamoDbGrammarParser {
	this := new(DynamoDbGrammarParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "DynamoDbGrammar.g4"

	return this
}

func validateRedundantParentheses(redundantParens bool) {
	// Do nothing
}
    /*private static void validateRedundantParentheses(bool redundantParens) {	JAVA
        if (redundantParens) {
            throw new RedundantParenthesesException();
        }
    }*/



// DynamoDbGrammarParser tokens.
const (
	DynamoDbGrammarParserEOF = antlr.TokenEOF
	DynamoDbGrammarParserT__0 = 1
	DynamoDbGrammarParserT__1 = 2
	DynamoDbGrammarParserT__2 = 3
	DynamoDbGrammarParserT__3 = 4
	DynamoDbGrammarParserT__4 = 5
	DynamoDbGrammarParserT__5 = 6
	DynamoDbGrammarParserT__6 = 7
	DynamoDbGrammarParserWS = 8
	DynamoDbGrammarParserEQ = 9
	DynamoDbGrammarParserNE = 10
	DynamoDbGrammarParserLT = 11
	DynamoDbGrammarParserLE = 12
	DynamoDbGrammarParserGT = 13
	DynamoDbGrammarParserGE = 14
	DynamoDbGrammarParserPLUS = 15
	DynamoDbGrammarParserMINUS = 16
	DynamoDbGrammarParserIN = 17
	DynamoDbGrammarParserBETWEEN = 18
	DynamoDbGrammarParserNOT = 19
	DynamoDbGrammarParserAND = 20
	DynamoDbGrammarParserOR = 21
	DynamoDbGrammarParserSET = 22
	DynamoDbGrammarParserADD = 23
	DynamoDbGrammarParserDELETE = 24
	DynamoDbGrammarParserREMOVE = 25
	DynamoDbGrammarParserINSERT = 26
	DynamoDbGrammarParserINTO = 27
	DynamoDbGrammarParserVALUE = 28
	DynamoDbGrammarParserVALUES = 29
	DynamoDbGrammarParserUPDATE = 30
	DynamoDbGrammarParserON = 31
	DynamoDbGrammarParserRETURNING = 32
	DynamoDbGrammarParserALL = 33
	DynamoDbGrammarParserKEYS = 34
	DynamoDbGrammarParserKEY = 35
	DynamoDbGrammarParserONLY = 36
	DynamoDbGrammarParserNEW = 37
	DynamoDbGrammarParserNONE = 38
	DynamoDbGrammarParserOLD = 39
	DynamoDbGrammarParserUPDATED = 40
	DynamoDbGrammarParserDUPLICATE = 41
	DynamoDbGrammarParserCREATE = 42
	DynamoDbGrammarParserTABLE = 43
	DynamoDbGrammarParserPRIMARY = 44
	DynamoDbGrammarParserNUMBER = 45
	DynamoDbGrammarParserBINARY = 46
	DynamoDbGrammarParserSTRING = 47
	DynamoDbGrammarParserIF = 48
	DynamoDbGrammarParserCAPACITY = 49
	DynamoDbGrammarParserGLOBAL = 50
	DynamoDbGrammarParserLOCAL = 51
	DynamoDbGrammarParserINDEXKEYWORD = 52
	DynamoDbGrammarParserPROJECTION = 53
	DynamoDbGrammarParserSHOW = 54
	DynamoDbGrammarParserTABLES = 55
	DynamoDbGrammarParserSELECT = 56
	DynamoDbGrammarParserFROM = 57
	DynamoDbGrammarParserUSE = 58
	DynamoDbGrammarParserENABLE = 59
	DynamoDbGrammarParserSCAN = 60
	DynamoDbGrammarParserWHERE = 61
	DynamoDbGrammarParserDROP = 62
	DynamoDbGrammarParserALTER = 63
	DynamoDbGrammarParserDESCRIBE = 64
	DynamoDbGrammarParserOPTION = 65
	DynamoDbGrammarParserINDEX = 66
	DynamoDbGrammarParserID = 67
	DynamoDbGrammarParserATTRIBUTE_NAME_SUB = 68
	DynamoDbGrammarParserLITERAL_SUB = 69
	DynamoDbGrammarParserSTRING_LITERAL = 70
	DynamoDbGrammarParserUNKNOWN = 71
)

// DynamoDbGrammarParser rules.
const (
	DynamoDbGrammarParserRULE_projection_ = 0
	DynamoDbGrammarParserRULE_projection = 1
	DynamoDbGrammarParserRULE_condition_ = 2
	DynamoDbGrammarParserRULE_condition = 3
	DynamoDbGrammarParserRULE_comparator_symbol = 4
	DynamoDbGrammarParserRULE_update_ = 5
	DynamoDbGrammarParserRULE_update = 6
	DynamoDbGrammarParserRULE_set_section = 7
	DynamoDbGrammarParserRULE_set_action = 8
	DynamoDbGrammarParserRULE_add_section = 9
	DynamoDbGrammarParserRULE_add_action = 10
	DynamoDbGrammarParserRULE_delete_section = 11
	DynamoDbGrammarParserRULE_delete_action = 12
	DynamoDbGrammarParserRULE_remove_section = 13
	DynamoDbGrammarParserRULE_remove_action = 14
	DynamoDbGrammarParserRULE_set_value = 15
	DynamoDbGrammarParserRULE_arithmetic = 16
	DynamoDbGrammarParserRULE_operand = 17
	DynamoDbGrammarParserRULE_function = 18
	DynamoDbGrammarParserRULE_path = 19
	DynamoDbGrammarParserRULE_id = 20
	DynamoDbGrammarParserRULE_dereference = 21
	DynamoDbGrammarParserRULE_literal = 22
	DynamoDbGrammarParserRULE_expression_attr_names_sub = 23
	DynamoDbGrammarParserRULE_expression_attr_values_sub = 24
	DynamoDbGrammarParserRULE_statement_ = 25
	DynamoDbGrammarParserRULE_statement = 26
	DynamoDbGrammarParserRULE_dropTableStatement = 27
	DynamoDbGrammarParserRULE_describeStatement = 28
	DynamoDbGrammarParserRULE_alterTableStatement = 29
	DynamoDbGrammarParserRULE_alterTableStatementType = 30
	DynamoDbGrammarParserRULE_setCapacity = 31
	DynamoDbGrammarParserRULE_addIndex = 32
	DynamoDbGrammarParserRULE_dropIndex = 33
	DynamoDbGrammarParserRULE_alterIndex = 34
	DynamoDbGrammarParserRULE_updateStatement = 35
	DynamoDbGrammarParserRULE_deleteStatement = 36
	DynamoDbGrammarParserRULE_insertStatement = 37
	DynamoDbGrammarParserRULE_createTableStatement = 38
	DynamoDbGrammarParserRULE_showTablesStatement = 39
	DynamoDbGrammarParserRULE_selectStatement = 40
	DynamoDbGrammarParserRULE_selectProjection = 41
	DynamoDbGrammarParserRULE_optionBlock = 42
	DynamoDbGrammarParserRULE_option = 43
	DynamoDbGrammarParserRULE_singleOption = 44
	DynamoDbGrammarParserRULE_keyValueOption = 45
	DynamoDbGrammarParserRULE_optionKey = 46
	DynamoDbGrammarParserRULE_optionValue = 47
	DynamoDbGrammarParserRULE_optionValueString = 48
	DynamoDbGrammarParserRULE_optionValueNumber = 49
	DynamoDbGrammarParserRULE_star = 50
	DynamoDbGrammarParserRULE_hint = 51
	DynamoDbGrammarParserRULE_indexHint = 52
	DynamoDbGrammarParserRULE_indexHintName = 53
	DynamoDbGrammarParserRULE_scanInfo = 54
	DynamoDbGrammarParserRULE_totalSegment = 55
	DynamoDbGrammarParserRULE_segment = 56
	DynamoDbGrammarParserRULE_where = 57
	DynamoDbGrammarParserRULE_primaryKeyDecl = 58
	DynamoDbGrammarParserRULE_secondaryIndexDecl = 59
	DynamoDbGrammarParserRULE_secondaryIndexType = 60
	DynamoDbGrammarParserRULE_indexName = 61
	DynamoDbGrammarParserRULE_projectionIndex = 62
	DynamoDbGrammarParserRULE_projectionIndexKeysOnly = 63
	DynamoDbGrammarParserRULE_projectionIndexVector = 64
	DynamoDbGrammarParserRULE_capacity = 65
	DynamoDbGrammarParserRULE_readUnits = 66
	DynamoDbGrammarParserRULE_writeUnits = 67
	DynamoDbGrammarParserRULE_indexDecl = 68
	DynamoDbGrammarParserRULE_attributeDecl = 69
	DynamoDbGrammarParserRULE_hashKey = 70
	DynamoDbGrammarParserRULE_rangeKey = 71
	DynamoDbGrammarParserRULE_attributeName = 72
	DynamoDbGrammarParserRULE_attributeType = 73
	DynamoDbGrammarParserRULE_returning = 74
	DynamoDbGrammarParserRULE_returningValue = 75
	DynamoDbGrammarParserRULE_onDuplicateKeyUpdate = 76
	DynamoDbGrammarParserRULE_ifClause = 77
	DynamoDbGrammarParserRULE_tableName = 78
	DynamoDbGrammarParserRULE_ddlName = 79
	DynamoDbGrammarParserRULE_stringLiteral = 80
	DynamoDbGrammarParserRULE_unknown = 81
)

// IProjection_Context is an interface to support dynamic dispatch.
type IProjection_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProjection_Context differentiates from other interfaces.
	IsProjection_Context()
}

type Projection_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjection_Context() *Projection_Context {
	var p = new(Projection_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_projection_
	return p
}

func (*Projection_Context) IsProjection_Context() {}

func NewProjection_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Projection_Context {
	var p = new(Projection_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_projection_

	return p
}

func (s *Projection_Context) GetParser() antlr.Parser { return s.parser }

func (s *Projection_Context) Projection() IProjectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProjectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProjectionContext)
}

func (s *Projection_Context) EOF() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserEOF, 0)
}

func (s *Projection_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Projection_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Projection_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterProjection_(s)
	}
}

func (s *Projection_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitProjection_(s)
	}
}




func (p *DynamoDbGrammarParser) Projection_() (localctx IProjection_Context) {
	localctx = NewProjection_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DynamoDbGrammarParserRULE_projection_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Projection()
	}
	{
		p.SetState(165)
		p.Match(DynamoDbGrammarParserEOF)
	}



	return localctx
}


// IProjectionContext is an interface to support dynamic dispatch.
type IProjectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProjectionContext differentiates from other interfaces.
	IsProjectionContext()
}

type ProjectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjectionContext() *ProjectionContext {
	var p = new(ProjectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_projection
	return p
}

func (*ProjectionContext) IsProjectionContext() {}

func NewProjectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjectionContext {
	var p = new(ProjectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_projection

	return p
}

func (s *ProjectionContext) GetParser() antlr.Parser { return s.parser }

func (s *ProjectionContext) AllPath() []IPathContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPathContext)(nil)).Elem())
	var tst = make([]IPathContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPathContext)
		}
	}

	return tst
}

func (s *ProjectionContext) Path(i int) IPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPathContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *ProjectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProjectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ProjectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterProjection(s)
	}
}

func (s *ProjectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitProjection(s)
	}
}




func (p *DynamoDbGrammarParser) Projection() (localctx IProjectionContext) {
	localctx = NewProjectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DynamoDbGrammarParserRULE_projection)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Path()
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == DynamoDbGrammarParserT__0 {
		{
			p.SetState(168)
			p.Match(DynamoDbGrammarParserT__0)
		}
		{
			p.SetState(169)
			p.Path()
		}


		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// ICondition_Context is an interface to support dynamic dispatch.
type ICondition_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCondition_Context differentiates from other interfaces.
	IsCondition_Context()
}

type Condition_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondition_Context() *Condition_Context {
	var p = new(Condition_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_condition_
	return p
}

func (*Condition_Context) IsCondition_Context() {}

func NewCondition_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Condition_Context {
	var p = new(Condition_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_condition_

	return p
}

func (s *Condition_Context) GetParser() antlr.Parser { return s.parser }

func (s *Condition_Context) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *Condition_Context) EOF() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserEOF, 0)
}

func (s *Condition_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Condition_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Condition_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterCondition_(s)
	}
}

func (s *Condition_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitCondition_(s)
	}
}




func (p *DynamoDbGrammarParser) Condition_() (localctx ICondition_Context) {
	localctx = NewCondition_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DynamoDbGrammarParserRULE_condition_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.condition(0)
	}
	{
		p.SetState(176)
		p.Match(DynamoDbGrammarParserEOF)
	}



	return localctx
}


// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetHasOuterParens returns the hasOuterParens attribute.
	GetHasOuterParens() bool


	// SetHasOuterParens sets the hasOuterParens attribute.
	SetHasOuterParens(bool)


	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	hasOuterParens bool// TODO = false
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) GetHasOuterParens() bool { return s.hasOuterParens }


func (s *ConditionContext) SetHasOuterParens(v bool) { s.hasOuterParens = v }


func (s *ConditionContext) CopyFrom(ctx *ConditionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
	s.hasOuterParens = ctx.hasOuterParens
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}





type OrContext struct {
	*ConditionContext
}

func NewOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrContext {
	var p = new(OrContext)

	p.ConditionContext = NewEmptyConditionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionContext))

	return p
}

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) AllCondition() []IConditionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConditionContext)(nil)).Elem())
	var tst = make([]IConditionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConditionContext)
		}
	}

	return tst
}

func (s *OrContext) Condition(i int) IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *OrContext) OR() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserOR, 0)
}


func (s *OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterOr(s)
	}
}

func (s *OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitOr(s)
	}
}


type NegationContext struct {
	*ConditionContext
}

func NewNegationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegationContext {
	var p = new(NegationContext)

	p.ConditionContext = NewEmptyConditionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionContext))

	return p
}

func (s *NegationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegationContext) NOT() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserNOT, 0)
}

func (s *NegationContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}


func (s *NegationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterNegation(s)
	}
}

func (s *NegationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitNegation(s)
	}
}


type InContext struct {
	*ConditionContext
}

func NewInContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InContext {
	var p = new(InContext)

	p.ConditionContext = NewEmptyConditionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionContext))

	return p
}

func (s *InContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InContext) AllOperand() []IOperandContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOperandContext)(nil)).Elem())
	var tst = make([]IOperandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOperandContext)
		}
	}

	return tst
}

func (s *InContext) Operand(i int) IOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *InContext) IN() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserIN, 0)
}


func (s *InContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterIn(s)
	}
}

func (s *InContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitIn(s)
	}
}


type AndContext struct {
	*ConditionContext
}

func NewAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndContext {
	var p = new(AndContext)

	p.ConditionContext = NewEmptyConditionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionContext))

	return p
}

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) AllCondition() []IConditionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConditionContext)(nil)).Elem())
	var tst = make([]IConditionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConditionContext)
		}
	}

	return tst
}

func (s *AndContext) Condition(i int) IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *AndContext) AND() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserAND, 0)
}


func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitAnd(s)
	}
}


type BetweenContext struct {
	*ConditionContext
}

func NewBetweenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BetweenContext {
	var p = new(BetweenContext)

	p.ConditionContext = NewEmptyConditionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionContext))

	return p
}

func (s *BetweenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BetweenContext) AllOperand() []IOperandContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOperandContext)(nil)).Elem())
	var tst = make([]IOperandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOperandContext)
		}
	}

	return tst
}

func (s *BetweenContext) Operand(i int) IOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *BetweenContext) BETWEEN() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserBETWEEN, 0)
}

func (s *BetweenContext) AND() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserAND, 0)
}


func (s *BetweenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterBetween(s)
	}
}

func (s *BetweenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitBetween(s)
	}
}


type FunctionConditionContext struct {
	*ConditionContext
}

func NewFunctionConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionConditionContext {
	var p = new(FunctionConditionContext)

	p.ConditionContext = NewEmptyConditionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionContext))

	return p
}

func (s *FunctionConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionConditionContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}


func (s *FunctionConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterFunctionCondition(s)
	}
}

func (s *FunctionConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitFunctionCondition(s)
	}
}


type ComparatorContext struct {
	*ConditionContext
}

func NewComparatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparatorContext {
	var p = new(ComparatorContext)

	p.ConditionContext = NewEmptyConditionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionContext))

	return p
}

func (s *ComparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparatorContext) AllOperand() []IOperandContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOperandContext)(nil)).Elem())
	var tst = make([]IOperandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOperandContext)
		}
	}

	return tst
}

func (s *ComparatorContext) Operand(i int) IOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *ComparatorContext) Comparator_symbol() IComparator_symbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparator_symbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparator_symbolContext)
}


func (s *ComparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterComparator(s)
	}
}

func (s *ComparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitComparator(s)
	}
}


type ConditionGroupingContext struct {
	*ConditionContext
	c IConditionContext 
}

func NewConditionGroupingContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConditionGroupingContext {
	var p = new(ConditionGroupingContext)

	p.ConditionContext = NewEmptyConditionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionContext))

	return p
}


func (s *ConditionGroupingContext) GetC() IConditionContext { return s.c }


func (s *ConditionGroupingContext) SetC(v IConditionContext) { s.c = v }

func (s *ConditionGroupingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionGroupingContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}


func (s *ConditionGroupingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterConditionGrouping(s)
	}
}

func (s *ConditionGroupingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitConditionGrouping(s)
	}
}



func (p *DynamoDbGrammarParser) Condition() (localctx IConditionContext) {
	return p.condition(0)
}

func (p *DynamoDbGrammarParser) condition(_p int) (localctx IConditionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewConditionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IConditionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, DynamoDbGrammarParserRULE_condition, _p)
	var _la int


	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewComparatorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(179)
			p.Operand()
		}
		{
			p.SetState(180)
			p.Comparator_symbol()
		}
		{
			p.SetState(181)
			p.Operand()
		}


	case 2:
		localctx = NewInContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(183)
			p.Operand()
		}
		{
			p.SetState(184)
			p.Match(DynamoDbGrammarParserIN)
		}
		{
			p.SetState(185)
			p.Match(DynamoDbGrammarParserT__1)
		}
		{
			p.SetState(186)
			p.Operand()
		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == DynamoDbGrammarParserT__0 {
			{
				p.SetState(187)
				p.Match(DynamoDbGrammarParserT__0)
			}
			{
				p.SetState(188)
				p.Operand()
			}


			p.SetState(193)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(194)
			p.Match(DynamoDbGrammarParserT__2)
		}


	case 3:
		localctx = NewBetweenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(196)
			p.Operand()
		}
		{
			p.SetState(197)
			p.Match(DynamoDbGrammarParserBETWEEN)
		}
		{
			p.SetState(198)
			p.Operand()
		}
		{
			p.SetState(199)
			p.Match(DynamoDbGrammarParserAND)
		}
		{
			p.SetState(200)
			p.Operand()
		}


	case 4:
		localctx = NewFunctionConditionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(202)
			p.Function()
		}


	case 5:
		localctx = NewConditionGroupingContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(203)
			p.Match(DynamoDbGrammarParserT__1)
		}
		{
			p.SetState(204)

			var _x = p.condition(0)

			localctx.(*ConditionGroupingContext).c = _x
		}
		{
			p.SetState(205)
			p.Match(DynamoDbGrammarParserT__2)
		}

		            validateRedundantParentheses(localctx.(*ConditionGroupingContext).GetC().GetHasOuterParens());
		            localctx.(*ConditionGroupingContext).SetHasOuterParens(true)
		        


	case 6:
		localctx = NewNegationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(208)
			p.Match(DynamoDbGrammarParserNOT)
		}
		{
			p.SetState(209)
			p.condition(3)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(218)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndContext(p, NewConditionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DynamoDbGrammarParserRULE_condition)
				p.SetState(212)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(213)
					p.Match(DynamoDbGrammarParserAND)
				}
				{
					p.SetState(214)
					p.condition(3)
				}


			case 2:
				localctx = NewOrContext(p, NewConditionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DynamoDbGrammarParserRULE_condition)
				p.SetState(215)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(216)
					p.Match(DynamoDbGrammarParserOR)
				}
				{
					p.SetState(217)
					p.condition(2)
				}

			}

		}
		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}



	return localctx
}


// IComparator_symbolContext is an interface to support dynamic dispatch.
type IComparator_symbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparator_symbolContext differentiates from other interfaces.
	IsComparator_symbolContext()
}

type Comparator_symbolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparator_symbolContext() *Comparator_symbolContext {
	var p = new(Comparator_symbolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_comparator_symbol
	return p
}

func (*Comparator_symbolContext) IsComparator_symbolContext() {}

func NewComparator_symbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comparator_symbolContext {
	var p = new(Comparator_symbolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_comparator_symbol

	return p
}

func (s *Comparator_symbolContext) GetParser() antlr.Parser { return s.parser }
func (s *Comparator_symbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comparator_symbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Comparator_symbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterComparator_symbol(s)
	}
}

func (s *Comparator_symbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitComparator_symbol(s)
	}
}




func (p *DynamoDbGrammarParser) Comparator_symbol() (localctx IComparator_symbolContext) {
	localctx = NewComparator_symbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DynamoDbGrammarParserRULE_comparator_symbol)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		_la = p.GetTokenStream().LA(1)

		if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << DynamoDbGrammarParserEQ) | (1 << DynamoDbGrammarParserNE) | (1 << DynamoDbGrammarParserLT) | (1 << DynamoDbGrammarParserLE) | (1 << DynamoDbGrammarParserGT) | (1 << DynamoDbGrammarParserGE))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IUpdate_Context is an interface to support dynamic dispatch.
type IUpdate_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpdate_Context differentiates from other interfaces.
	IsUpdate_Context()
}

type Update_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdate_Context() *Update_Context {
	var p = new(Update_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_update_
	return p
}

func (*Update_Context) IsUpdate_Context() {}

func NewUpdate_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Update_Context {
	var p = new(Update_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_update_

	return p
}

func (s *Update_Context) GetParser() antlr.Parser { return s.parser }

func (s *Update_Context) Update() IUpdateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpdateContext)
}

func (s *Update_Context) EOF() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserEOF, 0)
}

func (s *Update_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Update_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Update_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterUpdate_(s)
	}
}

func (s *Update_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitUpdate_(s)
	}
}




func (p *DynamoDbGrammarParser) Update_() (localctx IUpdate_Context) {
	localctx = NewUpdate_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DynamoDbGrammarParserRULE_update_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)
		p.Update()
	}
	{
		p.SetState(226)
		p.Match(DynamoDbGrammarParserEOF)
	}



	return localctx
}


// IUpdateContext is an interface to support dynamic dispatch.
type IUpdateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpdateContext differentiates from other interfaces.
	IsUpdateContext()
}

type UpdateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdateContext() *UpdateContext {
	var p = new(UpdateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_update
	return p
}

func (*UpdateContext) IsUpdateContext() {}

func NewUpdateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdateContext {
	var p = new(UpdateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_update

	return p
}

func (s *UpdateContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdateContext) AllSet_section() []ISet_sectionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISet_sectionContext)(nil)).Elem())
	var tst = make([]ISet_sectionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISet_sectionContext)
		}
	}

	return tst
}

func (s *UpdateContext) Set_section(i int) ISet_sectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISet_sectionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISet_sectionContext)
}

func (s *UpdateContext) AllAdd_section() []IAdd_sectionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAdd_sectionContext)(nil)).Elem())
	var tst = make([]IAdd_sectionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAdd_sectionContext)
		}
	}

	return tst
}

func (s *UpdateContext) Add_section(i int) IAdd_sectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdd_sectionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAdd_sectionContext)
}

func (s *UpdateContext) AllDelete_section() []IDelete_sectionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDelete_sectionContext)(nil)).Elem())
	var tst = make([]IDelete_sectionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDelete_sectionContext)
		}
	}

	return tst
}

func (s *UpdateContext) Delete_section(i int) IDelete_sectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelete_sectionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDelete_sectionContext)
}

func (s *UpdateContext) AllRemove_section() []IRemove_sectionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRemove_sectionContext)(nil)).Elem())
	var tst = make([]IRemove_sectionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRemove_sectionContext)
		}
	}

	return tst
}

func (s *UpdateContext) Remove_section(i int) IRemove_sectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRemove_sectionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRemove_sectionContext)
}

func (s *UpdateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *UpdateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterUpdate(s)
	}
}

func (s *UpdateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitUpdate(s)
	}
}




func (p *DynamoDbGrammarParser) Update() (localctx IUpdateContext) {
	localctx = NewUpdateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DynamoDbGrammarParserRULE_update)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << DynamoDbGrammarParserSET) | (1 << DynamoDbGrammarParserADD) | (1 << DynamoDbGrammarParserDELETE) | (1 << DynamoDbGrammarParserREMOVE))) != 0) {
		p.SetState(232)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case DynamoDbGrammarParserSET:
			{
				p.SetState(228)
				p.Set_section()
			}


		case DynamoDbGrammarParserADD:
			{
				p.SetState(229)
				p.Add_section()
			}


		case DynamoDbGrammarParserDELETE:
			{
				p.SetState(230)
				p.Delete_section()
			}


		case DynamoDbGrammarParserREMOVE:
			{
				p.SetState(231)
				p.Remove_section()
			}



		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// ISet_sectionContext is an interface to support dynamic dispatch.
type ISet_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSet_sectionContext differentiates from other interfaces.
	IsSet_sectionContext()
}

type Set_sectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySet_sectionContext() *Set_sectionContext {
	var p = new(Set_sectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_set_section
	return p
}

func (*Set_sectionContext) IsSet_sectionContext() {}

func NewSet_sectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_sectionContext {
	var p = new(Set_sectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_set_section

	return p
}

func (s *Set_sectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Set_sectionContext) SET() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserSET, 0)
}

func (s *Set_sectionContext) AllSet_action() []ISet_actionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISet_actionContext)(nil)).Elem())
	var tst = make([]ISet_actionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISet_actionContext)
		}
	}

	return tst
}

func (s *Set_sectionContext) Set_action(i int) ISet_actionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISet_actionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISet_actionContext)
}

func (s *Set_sectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_sectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Set_sectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterSet_section(s)
	}
}

func (s *Set_sectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitSet_section(s)
	}
}




func (p *DynamoDbGrammarParser) Set_section() (localctx ISet_sectionContext) {
	localctx = NewSet_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DynamoDbGrammarParserRULE_set_section)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.Match(DynamoDbGrammarParserSET)
	}
	{
		p.SetState(237)
		p.Set_action()
	}
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == DynamoDbGrammarParserT__0 {
		{
			p.SetState(238)
			p.Match(DynamoDbGrammarParserT__0)
		}
		{
			p.SetState(239)
			p.Set_action()
		}


		p.SetState(244)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// ISet_actionContext is an interface to support dynamic dispatch.
type ISet_actionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSet_actionContext differentiates from other interfaces.
	IsSet_actionContext()
}

type Set_actionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySet_actionContext() *Set_actionContext {
	var p = new(Set_actionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_set_action
	return p
}

func (*Set_actionContext) IsSet_actionContext() {}

func NewSet_actionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_actionContext {
	var p = new(Set_actionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_set_action

	return p
}

func (s *Set_actionContext) GetParser() antlr.Parser { return s.parser }

func (s *Set_actionContext) Path() IPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *Set_actionContext) Set_value() ISet_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISet_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISet_valueContext)
}

func (s *Set_actionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_actionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Set_actionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterSet_action(s)
	}
}

func (s *Set_actionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitSet_action(s)
	}
}




func (p *DynamoDbGrammarParser) Set_action() (localctx ISet_actionContext) {
	localctx = NewSet_actionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DynamoDbGrammarParserRULE_set_action)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		p.Path()
	}
	{
		p.SetState(246)
		p.Match(DynamoDbGrammarParserEQ)
	}
	{
		p.SetState(247)
		p.Set_value()
	}



	return localctx
}


// IAdd_sectionContext is an interface to support dynamic dispatch.
type IAdd_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdd_sectionContext differentiates from other interfaces.
	IsAdd_sectionContext()
}

type Add_sectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdd_sectionContext() *Add_sectionContext {
	var p = new(Add_sectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_add_section
	return p
}

func (*Add_sectionContext) IsAdd_sectionContext() {}

func NewAdd_sectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Add_sectionContext {
	var p = new(Add_sectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_add_section

	return p
}

func (s *Add_sectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Add_sectionContext) ADD() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserADD, 0)
}

func (s *Add_sectionContext) AllAdd_action() []IAdd_actionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAdd_actionContext)(nil)).Elem())
	var tst = make([]IAdd_actionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAdd_actionContext)
		}
	}

	return tst
}

func (s *Add_sectionContext) Add_action(i int) IAdd_actionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdd_actionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAdd_actionContext)
}

func (s *Add_sectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Add_sectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Add_sectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterAdd_section(s)
	}
}

func (s *Add_sectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitAdd_section(s)
	}
}




func (p *DynamoDbGrammarParser) Add_section() (localctx IAdd_sectionContext) {
	localctx = NewAdd_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DynamoDbGrammarParserRULE_add_section)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.Match(DynamoDbGrammarParserADD)
	}
	{
		p.SetState(250)
		p.Add_action()
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == DynamoDbGrammarParserT__0 {
		{
			p.SetState(251)
			p.Match(DynamoDbGrammarParserT__0)
		}
		{
			p.SetState(252)
			p.Add_action()
		}


		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IAdd_actionContext is an interface to support dynamic dispatch.
type IAdd_actionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdd_actionContext differentiates from other interfaces.
	IsAdd_actionContext()
}

type Add_actionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdd_actionContext() *Add_actionContext {
	var p = new(Add_actionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_add_action
	return p
}

func (*Add_actionContext) IsAdd_actionContext() {}

func NewAdd_actionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Add_actionContext {
	var p = new(Add_actionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_add_action

	return p
}

func (s *Add_actionContext) GetParser() antlr.Parser { return s.parser }

func (s *Add_actionContext) Path() IPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *Add_actionContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *Add_actionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Add_actionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Add_actionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterAdd_action(s)
	}
}

func (s *Add_actionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitAdd_action(s)
	}
}




func (p *DynamoDbGrammarParser) Add_action() (localctx IAdd_actionContext) {
	localctx = NewAdd_actionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DynamoDbGrammarParserRULE_add_action)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(258)
		p.Path()
	}
	{
		p.SetState(259)
		p.Literal()
	}



	return localctx
}


// IDelete_sectionContext is an interface to support dynamic dispatch.
type IDelete_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelete_sectionContext differentiates from other interfaces.
	IsDelete_sectionContext()
}

type Delete_sectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelete_sectionContext() *Delete_sectionContext {
	var p = new(Delete_sectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_delete_section
	return p
}

func (*Delete_sectionContext) IsDelete_sectionContext() {}

func NewDelete_sectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Delete_sectionContext {
	var p = new(Delete_sectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_delete_section

	return p
}

func (s *Delete_sectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Delete_sectionContext) DELETE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserDELETE, 0)
}

func (s *Delete_sectionContext) AllDelete_action() []IDelete_actionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDelete_actionContext)(nil)).Elem())
	var tst = make([]IDelete_actionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDelete_actionContext)
		}
	}

	return tst
}

func (s *Delete_sectionContext) Delete_action(i int) IDelete_actionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDelete_actionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDelete_actionContext)
}

func (s *Delete_sectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Delete_sectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Delete_sectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterDelete_section(s)
	}
}

func (s *Delete_sectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitDelete_section(s)
	}
}




func (p *DynamoDbGrammarParser) Delete_section() (localctx IDelete_sectionContext) {
	localctx = NewDelete_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DynamoDbGrammarParserRULE_delete_section)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)
		p.Match(DynamoDbGrammarParserDELETE)
	}
	{
		p.SetState(262)
		p.Delete_action()
	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == DynamoDbGrammarParserT__0 {
		{
			p.SetState(263)
			p.Match(DynamoDbGrammarParserT__0)
		}
		{
			p.SetState(264)
			p.Delete_action()
		}


		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IDelete_actionContext is an interface to support dynamic dispatch.
type IDelete_actionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelete_actionContext differentiates from other interfaces.
	IsDelete_actionContext()
}

type Delete_actionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelete_actionContext() *Delete_actionContext {
	var p = new(Delete_actionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_delete_action
	return p
}

func (*Delete_actionContext) IsDelete_actionContext() {}

func NewDelete_actionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Delete_actionContext {
	var p = new(Delete_actionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_delete_action

	return p
}

func (s *Delete_actionContext) GetParser() antlr.Parser { return s.parser }

func (s *Delete_actionContext) Path() IPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *Delete_actionContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *Delete_actionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Delete_actionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Delete_actionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterDelete_action(s)
	}
}

func (s *Delete_actionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitDelete_action(s)
	}
}




func (p *DynamoDbGrammarParser) Delete_action() (localctx IDelete_actionContext) {
	localctx = NewDelete_actionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DynamoDbGrammarParserRULE_delete_action)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(270)
		p.Path()
	}
	{
		p.SetState(271)
		p.Literal()
	}



	return localctx
}


// IRemove_sectionContext is an interface to support dynamic dispatch.
type IRemove_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRemove_sectionContext differentiates from other interfaces.
	IsRemove_sectionContext()
}

type Remove_sectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRemove_sectionContext() *Remove_sectionContext {
	var p = new(Remove_sectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_remove_section
	return p
}

func (*Remove_sectionContext) IsRemove_sectionContext() {}

func NewRemove_sectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Remove_sectionContext {
	var p = new(Remove_sectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_remove_section

	return p
}

func (s *Remove_sectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Remove_sectionContext) REMOVE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserREMOVE, 0)
}

func (s *Remove_sectionContext) AllRemove_action() []IRemove_actionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRemove_actionContext)(nil)).Elem())
	var tst = make([]IRemove_actionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRemove_actionContext)
		}
	}

	return tst
}

func (s *Remove_sectionContext) Remove_action(i int) IRemove_actionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRemove_actionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRemove_actionContext)
}

func (s *Remove_sectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Remove_sectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Remove_sectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterRemove_section(s)
	}
}

func (s *Remove_sectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitRemove_section(s)
	}
}




func (p *DynamoDbGrammarParser) Remove_section() (localctx IRemove_sectionContext) {
	localctx = NewRemove_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, DynamoDbGrammarParserRULE_remove_section)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
		p.Match(DynamoDbGrammarParserREMOVE)
	}
	{
		p.SetState(274)
		p.Remove_action()
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == DynamoDbGrammarParserT__0 {
		{
			p.SetState(275)
			p.Match(DynamoDbGrammarParserT__0)
		}
		{
			p.SetState(276)
			p.Remove_action()
		}


		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IRemove_actionContext is an interface to support dynamic dispatch.
type IRemove_actionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRemove_actionContext differentiates from other interfaces.
	IsRemove_actionContext()
}

type Remove_actionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRemove_actionContext() *Remove_actionContext {
	var p = new(Remove_actionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_remove_action
	return p
}

func (*Remove_actionContext) IsRemove_actionContext() {}

func NewRemove_actionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Remove_actionContext {
	var p = new(Remove_actionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_remove_action

	return p
}

func (s *Remove_actionContext) GetParser() antlr.Parser { return s.parser }

func (s *Remove_actionContext) Path() IPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *Remove_actionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Remove_actionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Remove_actionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterRemove_action(s)
	}
}

func (s *Remove_actionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitRemove_action(s)
	}
}




func (p *DynamoDbGrammarParser) Remove_action() (localctx IRemove_actionContext) {
	localctx = NewRemove_actionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, DynamoDbGrammarParserRULE_remove_action)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.Path()
	}



	return localctx
}


// ISet_valueContext is an interface to support dynamic dispatch.
type ISet_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSet_valueContext differentiates from other interfaces.
	IsSet_valueContext()
}

type Set_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySet_valueContext() *Set_valueContext {
	var p = new(Set_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_set_value
	return p
}

func (*Set_valueContext) IsSet_valueContext() {}

func NewSet_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_valueContext {
	var p = new(Set_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_set_value

	return p
}

func (s *Set_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Set_valueContext) CopyFrom(ctx *Set_valueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Set_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type ArithmeticValueContext struct {
	*Set_valueContext
}

func NewArithmeticValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArithmeticValueContext {
	var p = new(ArithmeticValueContext)

	p.Set_valueContext = NewEmptySet_valueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Set_valueContext))

	return p
}

func (s *ArithmeticValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticValueContext) Arithmetic() IArithmeticContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithmeticContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithmeticContext)
}


func (s *ArithmeticValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterArithmeticValue(s)
	}
}

func (s *ArithmeticValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitArithmeticValue(s)
	}
}


type OperandValueContext struct {
	*Set_valueContext
}

func NewOperandValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OperandValueContext {
	var p = new(OperandValueContext)

	p.Set_valueContext = NewEmptySet_valueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Set_valueContext))

	return p
}

func (s *OperandValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandValueContext) Operand() IOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}


func (s *OperandValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterOperandValue(s)
	}
}

func (s *OperandValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitOperandValue(s)
	}
}



func (p *DynamoDbGrammarParser) Set_value() (localctx ISet_valueContext) {
	localctx = NewSet_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, DynamoDbGrammarParserRULE_set_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		localctx = NewOperandValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(284)
			p.Operand()
		}


	case 2:
		localctx = NewArithmeticValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(285)
			p.Arithmetic()
		}

	}


	return localctx
}


// IArithmeticContext is an interface to support dynamic dispatch.
type IArithmeticContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetHasOuterParens returns the hasOuterParens attribute.
	GetHasOuterParens() bool


	// SetHasOuterParens sets the hasOuterParens attribute.
	SetHasOuterParens(bool)


	// IsArithmeticContext differentiates from other interfaces.
	IsArithmeticContext()
}

type ArithmeticContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	hasOuterParens bool// TODO = false
}

func NewEmptyArithmeticContext() *ArithmeticContext {
	var p = new(ArithmeticContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_arithmetic
	return p
}

func (*ArithmeticContext) IsArithmeticContext() {}

func NewArithmeticContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArithmeticContext {
	var p = new(ArithmeticContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_arithmetic

	return p
}

func (s *ArithmeticContext) GetParser() antlr.Parser { return s.parser }

func (s *ArithmeticContext) GetHasOuterParens() bool { return s.hasOuterParens }


func (s *ArithmeticContext) SetHasOuterParens(v bool) { s.hasOuterParens = v }


func (s *ArithmeticContext) CopyFrom(ctx *ArithmeticContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
	s.hasOuterParens = ctx.hasOuterParens
}

func (s *ArithmeticContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type PlusMinusContext struct {
	*ArithmeticContext
}

func NewPlusMinusContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlusMinusContext {
	var p = new(PlusMinusContext)

	p.ArithmeticContext = NewEmptyArithmeticContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArithmeticContext))

	return p
}

func (s *PlusMinusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlusMinusContext) AllOperand() []IOperandContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOperandContext)(nil)).Elem())
	var tst = make([]IOperandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOperandContext)
		}
	}

	return tst
}

func (s *PlusMinusContext) Operand(i int) IOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}


func (s *PlusMinusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterPlusMinus(s)
	}
}

func (s *PlusMinusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitPlusMinus(s)
	}
}


type ArithmeticParensContext struct {
	*ArithmeticContext
	a IArithmeticContext 
}

func NewArithmeticParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArithmeticParensContext {
	var p = new(ArithmeticParensContext)

	p.ArithmeticContext = NewEmptyArithmeticContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArithmeticContext))

	return p
}


func (s *ArithmeticParensContext) GetA() IArithmeticContext { return s.a }


func (s *ArithmeticParensContext) SetA(v IArithmeticContext) { s.a = v }

func (s *ArithmeticParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticParensContext) Arithmetic() IArithmeticContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithmeticContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithmeticContext)
}


func (s *ArithmeticParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterArithmeticParens(s)
	}
}

func (s *ArithmeticParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitArithmeticParens(s)
	}
}



func (p *DynamoDbGrammarParser) Arithmetic() (localctx IArithmeticContext) {
	localctx = NewArithmeticContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, DynamoDbGrammarParserRULE_arithmetic)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPlusMinusContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(288)
			p.Operand()
		}
		{
			p.SetState(289)
			_la = p.GetTokenStream().LA(1)

			if !(_la == DynamoDbGrammarParserPLUS || _la == DynamoDbGrammarParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(290)
			p.Operand()
		}


	case 2:
		localctx = NewArithmeticParensContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(292)
			p.Match(DynamoDbGrammarParserT__1)
		}
		{
			p.SetState(293)

			var _x = p.Arithmetic()


			localctx.(*ArithmeticParensContext).a = _x
		}
		{
			p.SetState(294)
			p.Match(DynamoDbGrammarParserT__2)
		}

		            validateRedundantParentheses(localctx.(*ArithmeticParensContext).GetA().GetHasOuterParens());
		            localctx.(*ArithmeticParensContext).SetHasOuterParens(true)
		        

	}


	return localctx
}


// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetHasOuterParens returns the hasOuterParens attribute.
	GetHasOuterParens() bool


	// SetHasOuterParens sets the hasOuterParens attribute.
	SetHasOuterParens(bool)


	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	hasOuterParens bool// TODO = false
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_operand
	return p
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) GetHasOuterParens() bool { return s.hasOuterParens }


func (s *OperandContext) SetHasOuterParens(v bool) { s.hasOuterParens = v }


func (s *OperandContext) CopyFrom(ctx *OperandContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
	s.hasOuterParens = ctx.hasOuterParens
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type PathOperandContext struct {
	*OperandContext
}

func NewPathOperandContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PathOperandContext {
	var p = new(PathOperandContext)

	p.OperandContext = NewEmptyOperandContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperandContext))

	return p
}

func (s *PathOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathOperandContext) Path() IPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}


func (s *PathOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterPathOperand(s)
	}
}

func (s *PathOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitPathOperand(s)
	}
}


type LiteralOperandContext struct {
	*OperandContext
}

func NewLiteralOperandContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralOperandContext {
	var p = new(LiteralOperandContext)

	p.OperandContext = NewEmptyOperandContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperandContext))

	return p
}

func (s *LiteralOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralOperandContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}


func (s *LiteralOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterLiteralOperand(s)
	}
}

func (s *LiteralOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitLiteralOperand(s)
	}
}


type FunctionOperandContext struct {
	*OperandContext
}

func NewFunctionOperandContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionOperandContext {
	var p = new(FunctionOperandContext)

	p.OperandContext = NewEmptyOperandContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperandContext))

	return p
}

func (s *FunctionOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionOperandContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}


func (s *FunctionOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterFunctionOperand(s)
	}
}

func (s *FunctionOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitFunctionOperand(s)
	}
}


type ParenOperandContext struct {
	*OperandContext
	o IOperandContext 
}

func NewParenOperandContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenOperandContext {
	var p = new(ParenOperandContext)

	p.OperandContext = NewEmptyOperandContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperandContext))

	return p
}


func (s *ParenOperandContext) GetO() IOperandContext { return s.o }


func (s *ParenOperandContext) SetO(v IOperandContext) { s.o = v }

func (s *ParenOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenOperandContext) Operand() IOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}


func (s *ParenOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterParenOperand(s)
	}
}

func (s *ParenOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitParenOperand(s)
	}
}



func (p *DynamoDbGrammarParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, DynamoDbGrammarParserRULE_operand)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPathOperandContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(299)
			p.Path()
		}


	case 2:
		localctx = NewLiteralOperandContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(300)
			p.Literal()
		}


	case 3:
		localctx = NewFunctionOperandContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(301)
			p.Function()
		}


	case 4:
		localctx = NewParenOperandContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(302)
			p.Match(DynamoDbGrammarParserT__1)
		}
		{
			p.SetState(303)

			var _x = p.Operand()


			localctx.(*ParenOperandContext).o = _x
		}
		{
			p.SetState(304)
			p.Match(DynamoDbGrammarParserT__2)
		}

		            validateRedundantParentheses(localctx.(*ParenOperandContext).GetO().GetHasOuterParens());
		            localctx.(*ParenOperandContext).SetHasOuterParens(true)
		        

	}


	return localctx
}


// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) CopyFrom(ctx *FunctionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type FunctionCallContext struct {
	*FunctionContext
}

func NewFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ID() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserID, 0)
}

func (s *FunctionCallContext) AllOperand() []IOperandContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOperandContext)(nil)).Elem())
	var tst = make([]IOperandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOperandContext)
		}
	}

	return tst
}

func (s *FunctionCallContext) Operand(i int) IOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}


func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}



func (p *DynamoDbGrammarParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, DynamoDbGrammarParserRULE_function)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewFunctionCallContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
		p.Match(DynamoDbGrammarParserID)
	}
	{
		p.SetState(310)
		p.Match(DynamoDbGrammarParserT__1)
	}
	{
		p.SetState(311)
		p.Operand()
	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == DynamoDbGrammarParserT__0 {
		{
			p.SetState(312)
			p.Match(DynamoDbGrammarParserT__0)
		}
		{
			p.SetState(313)
			p.Operand()
		}


		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(319)
		p.Match(DynamoDbGrammarParserT__2)
	}



	return localctx
}


// IPathContext is an interface to support dynamic dispatch.
type IPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPathContext differentiates from other interfaces.
	IsPathContext()
}

type PathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathContext() *PathContext {
	var p = new(PathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_path
	return p
}

func (*PathContext) IsPathContext() {}

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	var p = new(PathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_path

	return p
}

func (s *PathContext) GetParser() antlr.Parser { return s.parser }

func (s *PathContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *PathContext) AllDereference() []IDereferenceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDereferenceContext)(nil)).Elem())
	var tst = make([]IDereferenceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDereferenceContext)
		}
	}

	return tst
}

func (s *PathContext) Dereference(i int) IDereferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDereferenceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDereferenceContext)
}

func (s *PathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterPath(s)
	}
}

func (s *PathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitPath(s)
	}
}




func (p *DynamoDbGrammarParser) Path() (localctx IPathContext) {
	localctx = NewPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, DynamoDbGrammarParserRULE_path)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(321)
		p.Id()
	}
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(322)
				p.Dereference()
			}


		}
		p.SetState(327)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}



	return localctx
}


// IIdContext is an interface to support dynamic dispatch.
type IIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdContext differentiates from other interfaces.
	IsIdContext()
}

type IdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdContext() *IdContext {
	var p = new(IdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_id
	return p
}

func (*IdContext) IsIdContext() {}

func NewIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdContext {
	var p = new(IdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_id

	return p
}

func (s *IdContext) GetParser() antlr.Parser { return s.parser }

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserID, 0)
}

func (s *IdContext) ATTRIBUTE_NAME_SUB() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserATTRIBUTE_NAME_SUB, 0)
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitId(s)
	}
}




func (p *DynamoDbGrammarParser) Id() (localctx IIdContext) {
	localctx = NewIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, DynamoDbGrammarParserRULE_id)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DynamoDbGrammarParserID || _la == DynamoDbGrammarParserATTRIBUTE_NAME_SUB) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IDereferenceContext is an interface to support dynamic dispatch.
type IDereferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDereferenceContext differentiates from other interfaces.
	IsDereferenceContext()
}

type DereferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDereferenceContext() *DereferenceContext {
	var p = new(DereferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_dereference
	return p
}

func (*DereferenceContext) IsDereferenceContext() {}

func NewDereferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DereferenceContext {
	var p = new(DereferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_dereference

	return p
}

func (s *DereferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *DereferenceContext) CopyFrom(ctx *DereferenceContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *DereferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DereferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type ListAccessContext struct {
	*DereferenceContext
}

func NewListAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListAccessContext {
	var p = new(ListAccessContext)

	p.DereferenceContext = NewEmptyDereferenceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DereferenceContext))

	return p
}

func (s *ListAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListAccessContext) INDEX() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserINDEX, 0)
}


func (s *ListAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterListAccess(s)
	}
}

func (s *ListAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitListAccess(s)
	}
}


type MapAccessContext struct {
	*DereferenceContext
}

func NewMapAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapAccessContext {
	var p = new(MapAccessContext)

	p.DereferenceContext = NewEmptyDereferenceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DereferenceContext))

	return p
}

func (s *MapAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapAccessContext) Id() IIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}


func (s *MapAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterMapAccess(s)
	}
}

func (s *MapAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitMapAccess(s)
	}
}



func (p *DynamoDbGrammarParser) Dereference() (localctx IDereferenceContext) {
	localctx = NewDereferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, DynamoDbGrammarParserRULE_dereference)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(335)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DynamoDbGrammarParserT__3:
		localctx = NewMapAccessContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(330)
			p.Match(DynamoDbGrammarParserT__3)
		}
		{
			p.SetState(331)
			p.Id()
		}


	case DynamoDbGrammarParserT__4:
		localctx = NewListAccessContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(332)
			p.Match(DynamoDbGrammarParserT__4)
		}
		{
			p.SetState(333)
			p.Match(DynamoDbGrammarParserINDEX)
		}
		{
			p.SetState(334)
			p.Match(DynamoDbGrammarParserT__5)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyFrom(ctx *LiteralContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type LiteralSubContext struct {
	*LiteralContext
}

func NewLiteralSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralSubContext {
	var p = new(LiteralSubContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *LiteralSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralSubContext) LITERAL_SUB() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserLITERAL_SUB, 0)
}


func (s *LiteralSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterLiteralSub(s)
	}
}

func (s *LiteralSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitLiteralSub(s)
	}
}



func (p *DynamoDbGrammarParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, DynamoDbGrammarParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewLiteralSubContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		p.Match(DynamoDbGrammarParserLITERAL_SUB)
	}



	return localctx
}


// IExpression_attr_names_subContext is an interface to support dynamic dispatch.
type IExpression_attr_names_subContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpression_attr_names_subContext differentiates from other interfaces.
	IsExpression_attr_names_subContext()
}

type Expression_attr_names_subContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_attr_names_subContext() *Expression_attr_names_subContext {
	var p = new(Expression_attr_names_subContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_expression_attr_names_sub
	return p
}

func (*Expression_attr_names_subContext) IsExpression_attr_names_subContext() {}

func NewExpression_attr_names_subContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_attr_names_subContext {
	var p = new(Expression_attr_names_subContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_expression_attr_names_sub

	return p
}

func (s *Expression_attr_names_subContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_attr_names_subContext) ATTRIBUTE_NAME_SUB() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserATTRIBUTE_NAME_SUB, 0)
}

func (s *Expression_attr_names_subContext) EOF() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserEOF, 0)
}

func (s *Expression_attr_names_subContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_attr_names_subContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Expression_attr_names_subContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterExpression_attr_names_sub(s)
	}
}

func (s *Expression_attr_names_subContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitExpression_attr_names_sub(s)
	}
}




func (p *DynamoDbGrammarParser) Expression_attr_names_sub() (localctx IExpression_attr_names_subContext) {
	localctx = NewExpression_attr_names_subContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, DynamoDbGrammarParserRULE_expression_attr_names_sub)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.Match(DynamoDbGrammarParserATTRIBUTE_NAME_SUB)
	}
	{
		p.SetState(340)
		p.Match(DynamoDbGrammarParserEOF)
	}



	return localctx
}


// IExpression_attr_values_subContext is an interface to support dynamic dispatch.
type IExpression_attr_values_subContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpression_attr_values_subContext differentiates from other interfaces.
	IsExpression_attr_values_subContext()
}

type Expression_attr_values_subContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_attr_values_subContext() *Expression_attr_values_subContext {
	var p = new(Expression_attr_values_subContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_expression_attr_values_sub
	return p
}

func (*Expression_attr_values_subContext) IsExpression_attr_values_subContext() {}

func NewExpression_attr_values_subContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_attr_values_subContext {
	var p = new(Expression_attr_values_subContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_expression_attr_values_sub

	return p
}

func (s *Expression_attr_values_subContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_attr_values_subContext) LITERAL_SUB() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserLITERAL_SUB, 0)
}

func (s *Expression_attr_values_subContext) EOF() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserEOF, 0)
}

func (s *Expression_attr_values_subContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_attr_values_subContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Expression_attr_values_subContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterExpression_attr_values_sub(s)
	}
}

func (s *Expression_attr_values_subContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitExpression_attr_values_sub(s)
	}
}




func (p *DynamoDbGrammarParser) Expression_attr_values_sub() (localctx IExpression_attr_values_subContext) {
	localctx = NewExpression_attr_values_subContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, DynamoDbGrammarParserRULE_expression_attr_values_sub)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(342)
		p.Match(DynamoDbGrammarParserLITERAL_SUB)
	}
	{
		p.SetState(343)
		p.Match(DynamoDbGrammarParserEOF)
	}



	return localctx
}


// IStatement_Context is an interface to support dynamic dispatch.
type IStatement_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatement_Context differentiates from other interfaces.
	IsStatement_Context()
}

type Statement_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatement_Context() *Statement_Context {
	var p = new(Statement_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_statement_
	return p
}

func (*Statement_Context) IsStatement_Context() {}

func NewStatement_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Statement_Context {
	var p = new(Statement_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_statement_

	return p
}

func (s *Statement_Context) GetParser() antlr.Parser { return s.parser }

func (s *Statement_Context) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *Statement_Context) EOF() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserEOF, 0)
}

func (s *Statement_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Statement_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Statement_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterStatement_(s)
	}
}

func (s *Statement_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitStatement_(s)
	}
}




func (p *DynamoDbGrammarParser) Statement_() (localctx IStatement_Context) {
	localctx = NewStatement_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, DynamoDbGrammarParserRULE_statement_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(345)
		p.Statement()
	}
	{
		p.SetState(346)
		p.Match(DynamoDbGrammarParserEOF)
	}



	return localctx
}


// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) InsertStatement() IInsertStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInsertStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInsertStatementContext)
}

func (s *StatementContext) CreateTableStatement() ICreateTableStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreateTableStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreateTableStatementContext)
}

func (s *StatementContext) ShowTablesStatement() IShowTablesStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShowTablesStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShowTablesStatementContext)
}

func (s *StatementContext) UpdateStatement() IUpdateStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdateStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpdateStatementContext)
}

func (s *StatementContext) DeleteStatement() IDeleteStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeleteStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeleteStatementContext)
}

func (s *StatementContext) SelectStatement() ISelectStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectStatementContext)
}

func (s *StatementContext) DropTableStatement() IDropTableStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDropTableStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDropTableStatementContext)
}

func (s *StatementContext) AlterTableStatement() IAlterTableStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlterTableStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlterTableStatementContext)
}

func (s *StatementContext) DescribeStatement() IDescribeStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescribeStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescribeStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitStatement(s)
	}
}




func (p *DynamoDbGrammarParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, DynamoDbGrammarParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(357)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DynamoDbGrammarParserINSERT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(348)
			p.InsertStatement()
		}


	case DynamoDbGrammarParserCREATE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(349)
			p.CreateTableStatement()
		}


	case DynamoDbGrammarParserSHOW:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(350)
			p.ShowTablesStatement()
		}


	case DynamoDbGrammarParserUPDATE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(351)
			p.UpdateStatement()
		}


	case DynamoDbGrammarParserDELETE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(352)
			p.DeleteStatement()
		}


	case DynamoDbGrammarParserSELECT:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(353)
			p.SelectStatement()
		}


	case DynamoDbGrammarParserDROP:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(354)
			p.DropTableStatement()
		}


	case DynamoDbGrammarParserALTER:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(355)
			p.AlterTableStatement()
		}


	case DynamoDbGrammarParserDESCRIBE:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(356)
			p.DescribeStatement()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IDropTableStatementContext is an interface to support dynamic dispatch.
type IDropTableStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDropTableStatementContext differentiates from other interfaces.
	IsDropTableStatementContext()
}

type DropTableStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDropTableStatementContext() *DropTableStatementContext {
	var p = new(DropTableStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_dropTableStatement
	return p
}

func (*DropTableStatementContext) IsDropTableStatementContext() {}

func NewDropTableStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropTableStatementContext {
	var p = new(DropTableStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_dropTableStatement

	return p
}

func (s *DropTableStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DropTableStatementContext) DROP() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserDROP, 0)
}

func (s *DropTableStatementContext) TABLE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserTABLE, 0)
}

func (s *DropTableStatementContext) TableName() ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *DropTableStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropTableStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DropTableStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterDropTableStatement(s)
	}
}

func (s *DropTableStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitDropTableStatement(s)
	}
}




func (p *DynamoDbGrammarParser) DropTableStatement() (localctx IDropTableStatementContext) {
	localctx = NewDropTableStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, DynamoDbGrammarParserRULE_dropTableStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(359)
		p.Match(DynamoDbGrammarParserDROP)
	}
	{
		p.SetState(360)
		p.Match(DynamoDbGrammarParserTABLE)
	}
	{
		p.SetState(361)
		p.TableName()
	}



	return localctx
}


// IDescribeStatementContext is an interface to support dynamic dispatch.
type IDescribeStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDescribeStatementContext differentiates from other interfaces.
	IsDescribeStatementContext()
}

type DescribeStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescribeStatementContext() *DescribeStatementContext {
	var p = new(DescribeStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_describeStatement
	return p
}

func (*DescribeStatementContext) IsDescribeStatementContext() {}

func NewDescribeStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescribeStatementContext {
	var p = new(DescribeStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_describeStatement

	return p
}

func (s *DescribeStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DescribeStatementContext) DESCRIBE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserDESCRIBE, 0)
}

func (s *DescribeStatementContext) TableName() ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *DescribeStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescribeStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DescribeStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterDescribeStatement(s)
	}
}

func (s *DescribeStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitDescribeStatement(s)
	}
}




func (p *DynamoDbGrammarParser) DescribeStatement() (localctx IDescribeStatementContext) {
	localctx = NewDescribeStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, DynamoDbGrammarParserRULE_describeStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.Match(DynamoDbGrammarParserDESCRIBE)
	}
	{
		p.SetState(364)
		p.TableName()
	}



	return localctx
}


// IAlterTableStatementContext is an interface to support dynamic dispatch.
type IAlterTableStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlterTableStatementContext differentiates from other interfaces.
	IsAlterTableStatementContext()
}

type AlterTableStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlterTableStatementContext() *AlterTableStatementContext {
	var p = new(AlterTableStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_alterTableStatement
	return p
}

func (*AlterTableStatementContext) IsAlterTableStatementContext() {}

func NewAlterTableStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlterTableStatementContext {
	var p = new(AlterTableStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_alterTableStatement

	return p
}

func (s *AlterTableStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AlterTableStatementContext) ALTER() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserALTER, 0)
}

func (s *AlterTableStatementContext) TABLE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserTABLE, 0)
}

func (s *AlterTableStatementContext) TableName() ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *AlterTableStatementContext) AlterTableStatementType() IAlterTableStatementTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlterTableStatementTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlterTableStatementTypeContext)
}

func (s *AlterTableStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlterTableStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AlterTableStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterAlterTableStatement(s)
	}
}

func (s *AlterTableStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitAlterTableStatement(s)
	}
}




func (p *DynamoDbGrammarParser) AlterTableStatement() (localctx IAlterTableStatementContext) {
	localctx = NewAlterTableStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, DynamoDbGrammarParserRULE_alterTableStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(366)
		p.Match(DynamoDbGrammarParserALTER)
	}
	{
		p.SetState(367)
		p.Match(DynamoDbGrammarParserTABLE)
	}
	{
		p.SetState(368)
		p.TableName()
	}
	{
		p.SetState(369)
		p.AlterTableStatementType()
	}



	return localctx
}


// IAlterTableStatementTypeContext is an interface to support dynamic dispatch.
type IAlterTableStatementTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlterTableStatementTypeContext differentiates from other interfaces.
	IsAlterTableStatementTypeContext()
}

type AlterTableStatementTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlterTableStatementTypeContext() *AlterTableStatementTypeContext {
	var p = new(AlterTableStatementTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_alterTableStatementType
	return p
}

func (*AlterTableStatementTypeContext) IsAlterTableStatementTypeContext() {}

func NewAlterTableStatementTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlterTableStatementTypeContext {
	var p = new(AlterTableStatementTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_alterTableStatementType

	return p
}

func (s *AlterTableStatementTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *AlterTableStatementTypeContext) SetCapacity() ISetCapacityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetCapacityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetCapacityContext)
}

func (s *AlterTableStatementTypeContext) AddIndex() IAddIndexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddIndexContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAddIndexContext)
}

func (s *AlterTableStatementTypeContext) DropIndex() IDropIndexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDropIndexContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDropIndexContext)
}

func (s *AlterTableStatementTypeContext) AlterIndex() IAlterIndexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlterIndexContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlterIndexContext)
}

func (s *AlterTableStatementTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlterTableStatementTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AlterTableStatementTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterAlterTableStatementType(s)
	}
}

func (s *AlterTableStatementTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitAlterTableStatementType(s)
	}
}




func (p *DynamoDbGrammarParser) AlterTableStatementType() (localctx IAlterTableStatementTypeContext) {
	localctx = NewAlterTableStatementTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, DynamoDbGrammarParserRULE_alterTableStatementType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(375)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DynamoDbGrammarParserSET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(371)
			p.SetCapacity()
		}


	case DynamoDbGrammarParserADD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(372)
			p.AddIndex()
		}


	case DynamoDbGrammarParserDROP:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(373)
			p.DropIndex()
		}


	case DynamoDbGrammarParserALTER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(374)
			p.AlterIndex()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// ISetCapacityContext is an interface to support dynamic dispatch.
type ISetCapacityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetCapacityContext differentiates from other interfaces.
	IsSetCapacityContext()
}

type SetCapacityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetCapacityContext() *SetCapacityContext {
	var p = new(SetCapacityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_setCapacity
	return p
}

func (*SetCapacityContext) IsSetCapacityContext() {}

func NewSetCapacityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetCapacityContext {
	var p = new(SetCapacityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_setCapacity

	return p
}

func (s *SetCapacityContext) GetParser() antlr.Parser { return s.parser }

func (s *SetCapacityContext) SET() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserSET, 0)
}

func (s *SetCapacityContext) Capacity() ICapacityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICapacityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICapacityContext)
}

func (s *SetCapacityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetCapacityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SetCapacityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterSetCapacity(s)
	}
}

func (s *SetCapacityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitSetCapacity(s)
	}
}




func (p *DynamoDbGrammarParser) SetCapacity() (localctx ISetCapacityContext) {
	localctx = NewSetCapacityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, DynamoDbGrammarParserRULE_setCapacity)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(377)
		p.Match(DynamoDbGrammarParserSET)
	}
	{
		p.SetState(378)
		p.Capacity()
	}



	return localctx
}


// IAddIndexContext is an interface to support dynamic dispatch.
type IAddIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAddIndexContext differentiates from other interfaces.
	IsAddIndexContext()
}

type AddIndexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddIndexContext() *AddIndexContext {
	var p = new(AddIndexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_addIndex
	return p
}

func (*AddIndexContext) IsAddIndexContext() {}

func NewAddIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddIndexContext {
	var p = new(AddIndexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_addIndex

	return p
}

func (s *AddIndexContext) GetParser() antlr.Parser { return s.parser }

func (s *AddIndexContext) ADD() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserADD, 0)
}

func (s *AddIndexContext) SecondaryIndexDecl() ISecondaryIndexDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISecondaryIndexDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISecondaryIndexDeclContext)
}

func (s *AddIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddIndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AddIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterAddIndex(s)
	}
}

func (s *AddIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitAddIndex(s)
	}
}




func (p *DynamoDbGrammarParser) AddIndex() (localctx IAddIndexContext) {
	localctx = NewAddIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, DynamoDbGrammarParserRULE_addIndex)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(380)
		p.Match(DynamoDbGrammarParserADD)
	}
	{
		p.SetState(381)
		p.SecondaryIndexDecl()
	}



	return localctx
}


// IDropIndexContext is an interface to support dynamic dispatch.
type IDropIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDropIndexContext differentiates from other interfaces.
	IsDropIndexContext()
}

type DropIndexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDropIndexContext() *DropIndexContext {
	var p = new(DropIndexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_dropIndex
	return p
}

func (*DropIndexContext) IsDropIndexContext() {}

func NewDropIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropIndexContext {
	var p = new(DropIndexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_dropIndex

	return p
}

func (s *DropIndexContext) GetParser() antlr.Parser { return s.parser }

func (s *DropIndexContext) DROP() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserDROP, 0)
}

func (s *DropIndexContext) INDEXKEYWORD() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserINDEXKEYWORD, 0)
}

func (s *DropIndexContext) IndexName() IIndexNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexNameContext)
}

func (s *DropIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropIndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DropIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterDropIndex(s)
	}
}

func (s *DropIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitDropIndex(s)
	}
}




func (p *DynamoDbGrammarParser) DropIndex() (localctx IDropIndexContext) {
	localctx = NewDropIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, DynamoDbGrammarParserRULE_dropIndex)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(383)
		p.Match(DynamoDbGrammarParserDROP)
	}
	{
		p.SetState(384)
		p.Match(DynamoDbGrammarParserINDEXKEYWORD)
	}
	{
		p.SetState(385)
		p.IndexName()
	}



	return localctx
}


// IAlterIndexContext is an interface to support dynamic dispatch.
type IAlterIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlterIndexContext differentiates from other interfaces.
	IsAlterIndexContext()
}

type AlterIndexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlterIndexContext() *AlterIndexContext {
	var p = new(AlterIndexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_alterIndex
	return p
}

func (*AlterIndexContext) IsAlterIndexContext() {}

func NewAlterIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlterIndexContext {
	var p = new(AlterIndexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_alterIndex

	return p
}

func (s *AlterIndexContext) GetParser() antlr.Parser { return s.parser }

func (s *AlterIndexContext) ALTER() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserALTER, 0)
}

func (s *AlterIndexContext) INDEXKEYWORD() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserINDEXKEYWORD, 0)
}

func (s *AlterIndexContext) IndexName() IIndexNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexNameContext)
}

func (s *AlterIndexContext) SetCapacity() ISetCapacityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetCapacityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetCapacityContext)
}

func (s *AlterIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlterIndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AlterIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterAlterIndex(s)
	}
}

func (s *AlterIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitAlterIndex(s)
	}
}




func (p *DynamoDbGrammarParser) AlterIndex() (localctx IAlterIndexContext) {
	localctx = NewAlterIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, DynamoDbGrammarParserRULE_alterIndex)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(387)
		p.Match(DynamoDbGrammarParserALTER)
	}
	{
		p.SetState(388)
		p.Match(DynamoDbGrammarParserINDEXKEYWORD)
	}
	{
		p.SetState(389)
		p.IndexName()
	}
	{
		p.SetState(390)
		p.SetCapacity()
	}



	return localctx
}


// IUpdateStatementContext is an interface to support dynamic dispatch.
type IUpdateStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpdateStatementContext differentiates from other interfaces.
	IsUpdateStatementContext()
}

type UpdateStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdateStatementContext() *UpdateStatementContext {
	var p = new(UpdateStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_updateStatement
	return p
}

func (*UpdateStatementContext) IsUpdateStatementContext() {}

func NewUpdateStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdateStatementContext {
	var p = new(UpdateStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_updateStatement

	return p
}

func (s *UpdateStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdateStatementContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserUPDATE, 0)
}

func (s *UpdateStatementContext) TableName() ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *UpdateStatementContext) Update() IUpdateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpdateContext)
}

func (s *UpdateStatementContext) Where() IWhereContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereContext)
}

func (s *UpdateStatementContext) Returning() IReturningContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturningContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturningContext)
}

func (s *UpdateStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdateStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *UpdateStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterUpdateStatement(s)
	}
}

func (s *UpdateStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitUpdateStatement(s)
	}
}




func (p *DynamoDbGrammarParser) UpdateStatement() (localctx IUpdateStatementContext) {
	localctx = NewUpdateStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, DynamoDbGrammarParserRULE_updateStatement)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(392)
		p.Match(DynamoDbGrammarParserUPDATE)
	}
	{
		p.SetState(393)
		p.TableName()
	}
	{
		p.SetState(394)
		p.Update()
	}
	{
		p.SetState(395)
		p.Where()
	}
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == DynamoDbGrammarParserRETURNING {
		{
			p.SetState(396)
			p.Returning()
		}

	}



	return localctx
}


// IDeleteStatementContext is an interface to support dynamic dispatch.
type IDeleteStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeleteStatementContext differentiates from other interfaces.
	IsDeleteStatementContext()
}

type DeleteStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeleteStatementContext() *DeleteStatementContext {
	var p = new(DeleteStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_deleteStatement
	return p
}

func (*DeleteStatementContext) IsDeleteStatementContext() {}

func NewDeleteStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteStatementContext {
	var p = new(DeleteStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_deleteStatement

	return p
}

func (s *DeleteStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteStatementContext) DELETE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserDELETE, 0)
}

func (s *DeleteStatementContext) FROM() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserFROM, 0)
}

func (s *DeleteStatementContext) TableName() ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *DeleteStatementContext) Where() IWhereContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereContext)
}

func (s *DeleteStatementContext) Returning() IReturningContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturningContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturningContext)
}

func (s *DeleteStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DeleteStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterDeleteStatement(s)
	}
}

func (s *DeleteStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitDeleteStatement(s)
	}
}




func (p *DynamoDbGrammarParser) DeleteStatement() (localctx IDeleteStatementContext) {
	localctx = NewDeleteStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, DynamoDbGrammarParserRULE_deleteStatement)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(399)
		p.Match(DynamoDbGrammarParserDELETE)
	}
	{
		p.SetState(400)
		p.Match(DynamoDbGrammarParserFROM)
	}
	{
		p.SetState(401)
		p.TableName()
	}
	{
		p.SetState(402)
		p.Where()
	}
	p.SetState(404)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == DynamoDbGrammarParserRETURNING {
		{
			p.SetState(403)
			p.Returning()
		}

	}



	return localctx
}


// IInsertStatementContext is an interface to support dynamic dispatch.
type IInsertStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInsertStatementContext differentiates from other interfaces.
	IsInsertStatementContext()
}

type InsertStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsertStatementContext() *InsertStatementContext {
	var p = new(InsertStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_insertStatement
	return p
}

func (*InsertStatementContext) IsInsertStatementContext() {}

func NewInsertStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InsertStatementContext {
	var p = new(InsertStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_insertStatement

	return p
}

func (s *InsertStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *InsertStatementContext) INSERT() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserINSERT, 0)
}

func (s *InsertStatementContext) TableName() ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *InsertStatementContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *InsertStatementContext) INTO() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserINTO, 0)
}

func (s *InsertStatementContext) OnDuplicateKeyUpdate() IOnDuplicateKeyUpdateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOnDuplicateKeyUpdateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOnDuplicateKeyUpdateContext)
}

func (s *InsertStatementContext) Returning() IReturningContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturningContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturningContext)
}

func (s *InsertStatementContext) VALUES() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserVALUES, 0)
}

func (s *InsertStatementContext) VALUE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserVALUE, 0)
}

func (s *InsertStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InsertStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *InsertStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterInsertStatement(s)
	}
}

func (s *InsertStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitInsertStatement(s)
	}
}




func (p *DynamoDbGrammarParser) InsertStatement() (localctx IInsertStatementContext) {
	localctx = NewInsertStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, DynamoDbGrammarParserRULE_insertStatement)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(406)
		p.Match(DynamoDbGrammarParserINSERT)
	}
	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == DynamoDbGrammarParserINTO {
		{
			p.SetState(407)
			p.Match(DynamoDbGrammarParserINTO)
		}

	}
	{
		p.SetState(410)
		p.TableName()
	}
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == DynamoDbGrammarParserVALUE || _la == DynamoDbGrammarParserVALUES {
		{
			p.SetState(411)
			_la = p.GetTokenStream().LA(1)

			if !(_la == DynamoDbGrammarParserVALUE || _la == DynamoDbGrammarParserVALUES) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(414)
		p.Literal()
	}
	p.SetState(416)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == DynamoDbGrammarParserON {
		{
			p.SetState(415)
			p.OnDuplicateKeyUpdate()
		}

	}
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == DynamoDbGrammarParserRETURNING {
		{
			p.SetState(418)
			p.Returning()
		}

	}



	return localctx
}


// ICreateTableStatementContext is an interface to support dynamic dispatch.
type ICreateTableStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreateTableStatementContext differentiates from other interfaces.
	IsCreateTableStatementContext()
}

type CreateTableStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateTableStatementContext() *CreateTableStatementContext {
	var p = new(CreateTableStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_createTableStatement
	return p
}

func (*CreateTableStatementContext) IsCreateTableStatementContext() {}

func NewCreateTableStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateTableStatementContext {
	var p = new(CreateTableStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_createTableStatement

	return p
}

func (s *CreateTableStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateTableStatementContext) CREATE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserCREATE, 0)
}

func (s *CreateTableStatementContext) TABLE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserTABLE, 0)
}

func (s *CreateTableStatementContext) TableName() ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *CreateTableStatementContext) AllAttributeDecl() []IAttributeDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttributeDeclContext)(nil)).Elem())
	var tst = make([]IAttributeDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttributeDeclContext)
		}
	}

	return tst
}

func (s *CreateTableStatementContext) AttributeDecl(i int) IAttributeDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttributeDeclContext)
}

func (s *CreateTableStatementContext) PrimaryKeyDecl() IPrimaryKeyDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryKeyDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryKeyDeclContext)
}

func (s *CreateTableStatementContext) AllSecondaryIndexDecl() []ISecondaryIndexDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISecondaryIndexDeclContext)(nil)).Elem())
	var tst = make([]ISecondaryIndexDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISecondaryIndexDeclContext)
		}
	}

	return tst
}

func (s *CreateTableStatementContext) SecondaryIndexDecl(i int) ISecondaryIndexDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISecondaryIndexDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISecondaryIndexDeclContext)
}

func (s *CreateTableStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateTableStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CreateTableStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterCreateTableStatement(s)
	}
}

func (s *CreateTableStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitCreateTableStatement(s)
	}
}




func (p *DynamoDbGrammarParser) CreateTableStatement() (localctx ICreateTableStatementContext) {
	localctx = NewCreateTableStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, DynamoDbGrammarParserRULE_createTableStatement)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(421)
		p.Match(DynamoDbGrammarParserCREATE)
	}
	{
		p.SetState(422)
		p.Match(DynamoDbGrammarParserTABLE)
	}
	{
		p.SetState(423)
		p.TableName()
	}
	{
		p.SetState(424)
		p.Match(DynamoDbGrammarParserT__1)
	}
	{
		p.SetState(425)
		p.AttributeDecl()
	}
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(426)
				p.Match(DynamoDbGrammarParserT__0)
			}
			{
				p.SetState(427)
				p.AttributeDecl()
			}


		}
		p.SetState(432)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}
	{
		p.SetState(433)
		p.Match(DynamoDbGrammarParserT__0)
	}
	{
		p.SetState(434)
		p.PrimaryKeyDecl()
	}
	p.SetState(444)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == DynamoDbGrammarParserT__0 {
		{
			p.SetState(435)
			p.Match(DynamoDbGrammarParserT__0)
		}
		{
			p.SetState(436)
			p.SecondaryIndexDecl()
		}
		p.SetState(441)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == DynamoDbGrammarParserT__0 {
			{
				p.SetState(437)
				p.Match(DynamoDbGrammarParserT__0)
			}
			{
				p.SetState(438)
				p.SecondaryIndexDecl()
			}


			p.SetState(443)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(446)
		p.Match(DynamoDbGrammarParserT__2)
	}



	return localctx
}


// IShowTablesStatementContext is an interface to support dynamic dispatch.
type IShowTablesStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShowTablesStatementContext differentiates from other interfaces.
	IsShowTablesStatementContext()
}

type ShowTablesStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShowTablesStatementContext() *ShowTablesStatementContext {
	var p = new(ShowTablesStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_showTablesStatement
	return p
}

func (*ShowTablesStatementContext) IsShowTablesStatementContext() {}

func NewShowTablesStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShowTablesStatementContext {
	var p = new(ShowTablesStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_showTablesStatement

	return p
}

func (s *ShowTablesStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ShowTablesStatementContext) SHOW() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserSHOW, 0)
}

func (s *ShowTablesStatementContext) TABLES() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserTABLES, 0)
}

func (s *ShowTablesStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShowTablesStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ShowTablesStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterShowTablesStatement(s)
	}
}

func (s *ShowTablesStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitShowTablesStatement(s)
	}
}




func (p *DynamoDbGrammarParser) ShowTablesStatement() (localctx IShowTablesStatementContext) {
	localctx = NewShowTablesStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, DynamoDbGrammarParserRULE_showTablesStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(448)
		p.Match(DynamoDbGrammarParserSHOW)
	}
	{
		p.SetState(449)
		p.Match(DynamoDbGrammarParserTABLES)
	}



	return localctx
}


// ISelectStatementContext is an interface to support dynamic dispatch.
type ISelectStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectStatementContext differentiates from other interfaces.
	IsSelectStatementContext()
}

type SelectStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectStatementContext() *SelectStatementContext {
	var p = new(SelectStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_selectStatement
	return p
}

func (*SelectStatementContext) IsSelectStatementContext() {}

func NewSelectStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectStatementContext {
	var p = new(SelectStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_selectStatement

	return p
}

func (s *SelectStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectStatementContext) SELECT() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserSELECT, 0)
}

func (s *SelectStatementContext) SelectProjection() ISelectProjectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectProjectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectProjectionContext)
}

func (s *SelectStatementContext) FROM() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserFROM, 0)
}

func (s *SelectStatementContext) TableName() ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *SelectStatementContext) Hint() IHintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHintContext)
}

func (s *SelectStatementContext) Where() IWhereContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereContext)
}

func (s *SelectStatementContext) OptionBlock() IOptionBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionBlockContext)
}

func (s *SelectStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SelectStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterSelectStatement(s)
	}
}

func (s *SelectStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitSelectStatement(s)
	}
}




func (p *DynamoDbGrammarParser) SelectStatement() (localctx ISelectStatementContext) {
	localctx = NewSelectStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, DynamoDbGrammarParserRULE_selectStatement)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(451)
		p.Match(DynamoDbGrammarParserSELECT)
	}
	{
		p.SetState(452)
		p.SelectProjection()
	}
	{
		p.SetState(453)
		p.Match(DynamoDbGrammarParserFROM)
	}
	{
		p.SetState(454)
		p.TableName()
	}
	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == DynamoDbGrammarParserUSE {
		{
			p.SetState(455)
			p.Hint()
		}

	}
	p.SetState(459)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == DynamoDbGrammarParserWHERE {
		{
			p.SetState(458)
			p.Where()
		}

	}
	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == DynamoDbGrammarParserOPTION {
		{
			p.SetState(461)
			p.OptionBlock()
		}

	}



	return localctx
}


// ISelectProjectionContext is an interface to support dynamic dispatch.
type ISelectProjectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectProjectionContext differentiates from other interfaces.
	IsSelectProjectionContext()
}

type SelectProjectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectProjectionContext() *SelectProjectionContext {
	var p = new(SelectProjectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_selectProjection
	return p
}

func (*SelectProjectionContext) IsSelectProjectionContext() {}

func NewSelectProjectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectProjectionContext {
	var p = new(SelectProjectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_selectProjection

	return p
}

func (s *SelectProjectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectProjectionContext) Projection() IProjectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProjectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProjectionContext)
}

func (s *SelectProjectionContext) Star() IStarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStarContext)
}

func (s *SelectProjectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectProjectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SelectProjectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterSelectProjection(s)
	}
}

func (s *SelectProjectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitSelectProjection(s)
	}
}




func (p *DynamoDbGrammarParser) SelectProjection() (localctx ISelectProjectionContext) {
	localctx = NewSelectProjectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, DynamoDbGrammarParserRULE_selectProjection)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(466)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DynamoDbGrammarParserID, DynamoDbGrammarParserATTRIBUTE_NAME_SUB:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(464)
			p.Projection()
		}


	case DynamoDbGrammarParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(465)
			p.Star()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IOptionBlockContext is an interface to support dynamic dispatch.
type IOptionBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionBlockContext differentiates from other interfaces.
	IsOptionBlockContext()
}

type OptionBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionBlockContext() *OptionBlockContext {
	var p = new(OptionBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_optionBlock
	return p
}

func (*OptionBlockContext) IsOptionBlockContext() {}

func NewOptionBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionBlockContext {
	var p = new(OptionBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_optionBlock

	return p
}

func (s *OptionBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionBlockContext) OPTION() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserOPTION, 0)
}

func (s *OptionBlockContext) AllOption() []IOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionContext)(nil)).Elem())
	var tst = make([]IOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionContext)
		}
	}

	return tst
}

func (s *OptionBlockContext) Option(i int) IOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *OptionBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OptionBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterOptionBlock(s)
	}
}

func (s *OptionBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitOptionBlock(s)
	}
}




func (p *DynamoDbGrammarParser) OptionBlock() (localctx IOptionBlockContext) {
	localctx = NewOptionBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, DynamoDbGrammarParserRULE_optionBlock)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(468)
		p.Match(DynamoDbGrammarParserOPTION)
	}
	{
		p.SetState(469)
		p.Match(DynamoDbGrammarParserT__1)
	}
	{
		p.SetState(470)
		p.Option()
	}
	p.SetState(475)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == DynamoDbGrammarParserT__0 {
		{
			p.SetState(471)
			p.Match(DynamoDbGrammarParserT__0)
		}
		{
			p.SetState(472)
			p.Option()
		}


		p.SetState(477)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(478)
		p.Match(DynamoDbGrammarParserT__2)
	}



	return localctx
}


// IOptionContext is an interface to support dynamic dispatch.
type IOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionContext differentiates from other interfaces.
	IsOptionContext()
}

type OptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionContext() *OptionContext {
	var p = new(OptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_option
	return p
}

func (*OptionContext) IsOptionContext() {}

func NewOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionContext {
	var p = new(OptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_option

	return p
}

func (s *OptionContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionContext) SingleOption() ISingleOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleOptionContext)
}

func (s *OptionContext) KeyValueOption() IKeyValueOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyValueOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyValueOptionContext)
}

func (s *OptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterOption(s)
	}
}

func (s *OptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitOption(s)
	}
}




func (p *DynamoDbGrammarParser) Option() (localctx IOptionContext) {
	localctx = NewOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, DynamoDbGrammarParserRULE_option)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(482)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(480)
			p.SingleOption()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(481)
			p.KeyValueOption()
		}

	}


	return localctx
}


// ISingleOptionContext is an interface to support dynamic dispatch.
type ISingleOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleOptionContext differentiates from other interfaces.
	IsSingleOptionContext()
}

type SingleOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleOptionContext() *SingleOptionContext {
	var p = new(SingleOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_singleOption
	return p
}

func (*SingleOptionContext) IsSingleOptionContext() {}

func NewSingleOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleOptionContext {
	var p = new(SingleOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_singleOption

	return p
}

func (s *SingleOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleOptionContext) OptionKey() IOptionKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionKeyContext)
}

func (s *SingleOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SingleOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterSingleOption(s)
	}
}

func (s *SingleOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitSingleOption(s)
	}
}




func (p *DynamoDbGrammarParser) SingleOption() (localctx ISingleOptionContext) {
	localctx = NewSingleOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, DynamoDbGrammarParserRULE_singleOption)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(484)
		p.OptionKey()
	}



	return localctx
}


// IKeyValueOptionContext is an interface to support dynamic dispatch.
type IKeyValueOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyValueOptionContext differentiates from other interfaces.
	IsKeyValueOptionContext()
}

type KeyValueOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyValueOptionContext() *KeyValueOptionContext {
	var p = new(KeyValueOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_keyValueOption
	return p
}

func (*KeyValueOptionContext) IsKeyValueOptionContext() {}

func NewKeyValueOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyValueOptionContext {
	var p = new(KeyValueOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_keyValueOption

	return p
}

func (s *KeyValueOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyValueOptionContext) OptionKey() IOptionKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionKeyContext)
}

func (s *KeyValueOptionContext) OptionValue() IOptionValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionValueContext)
}

func (s *KeyValueOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyValueOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *KeyValueOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterKeyValueOption(s)
	}
}

func (s *KeyValueOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitKeyValueOption(s)
	}
}




func (p *DynamoDbGrammarParser) KeyValueOption() (localctx IKeyValueOptionContext) {
	localctx = NewKeyValueOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, DynamoDbGrammarParserRULE_keyValueOption)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(486)
		p.OptionKey()
	}
	{
		p.SetState(487)
		p.Match(DynamoDbGrammarParserEQ)
	}
	{
		p.SetState(488)
		p.OptionValue()
	}



	return localctx
}


// IOptionKeyContext is an interface to support dynamic dispatch.
type IOptionKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionKeyContext differentiates from other interfaces.
	IsOptionKeyContext()
}

type OptionKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionKeyContext() *OptionKeyContext {
	var p = new(OptionKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_optionKey
	return p
}

func (*OptionKeyContext) IsOptionKeyContext() {}

func NewOptionKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionKeyContext {
	var p = new(OptionKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_optionKey

	return p
}

func (s *OptionKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionKeyContext) DdlName() IDdlNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDdlNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDdlNameContext)
}

func (s *OptionKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OptionKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterOptionKey(s)
	}
}

func (s *OptionKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitOptionKey(s)
	}
}




func (p *DynamoDbGrammarParser) OptionKey() (localctx IOptionKeyContext) {
	localctx = NewOptionKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, DynamoDbGrammarParserRULE_optionKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(490)
		p.DdlName()
	}



	return localctx
}


// IOptionValueContext is an interface to support dynamic dispatch.
type IOptionValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionValueContext differentiates from other interfaces.
	IsOptionValueContext()
}

type OptionValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionValueContext() *OptionValueContext {
	var p = new(OptionValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_optionValue
	return p
}

func (*OptionValueContext) IsOptionValueContext() {}

func NewOptionValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionValueContext {
	var p = new(OptionValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_optionValue

	return p
}

func (s *OptionValueContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionValueContext) OptionValueString() IOptionValueStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionValueStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionValueStringContext)
}

func (s *OptionValueContext) OptionValueNumber() IOptionValueNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionValueNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionValueNumberContext)
}

func (s *OptionValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OptionValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterOptionValue(s)
	}
}

func (s *OptionValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitOptionValue(s)
	}
}




func (p *DynamoDbGrammarParser) OptionValue() (localctx IOptionValueContext) {
	localctx = NewOptionValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, DynamoDbGrammarParserRULE_optionValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(494)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DynamoDbGrammarParserNONE, DynamoDbGrammarParserID, DynamoDbGrammarParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(492)
			p.OptionValueString()
		}


	case DynamoDbGrammarParserINDEX:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(493)
			p.OptionValueNumber()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IOptionValueStringContext is an interface to support dynamic dispatch.
type IOptionValueStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionValueStringContext differentiates from other interfaces.
	IsOptionValueStringContext()
}

type OptionValueStringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionValueStringContext() *OptionValueStringContext {
	var p = new(OptionValueStringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_optionValueString
	return p
}

func (*OptionValueStringContext) IsOptionValueStringContext() {}

func NewOptionValueStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionValueStringContext {
	var p = new(OptionValueStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_optionValueString

	return p
}

func (s *OptionValueStringContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionValueStringContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *OptionValueStringContext) NONE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserNONE, 0)
}

func (s *OptionValueStringContext) ID() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserID, 0)
}

func (s *OptionValueStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionValueStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OptionValueStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterOptionValueString(s)
	}
}

func (s *OptionValueStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitOptionValueString(s)
	}
}




func (p *DynamoDbGrammarParser) OptionValueString() (localctx IOptionValueStringContext) {
	localctx = NewOptionValueStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, DynamoDbGrammarParserRULE_optionValueString)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(499)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DynamoDbGrammarParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(496)
			p.StringLiteral()
		}


	case DynamoDbGrammarParserNONE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(497)
			p.Match(DynamoDbGrammarParserNONE)
		}


	case DynamoDbGrammarParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(498)
			p.Match(DynamoDbGrammarParserID)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IOptionValueNumberContext is an interface to support dynamic dispatch.
type IOptionValueNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionValueNumberContext differentiates from other interfaces.
	IsOptionValueNumberContext()
}

type OptionValueNumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionValueNumberContext() *OptionValueNumberContext {
	var p = new(OptionValueNumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_optionValueNumber
	return p
}

func (*OptionValueNumberContext) IsOptionValueNumberContext() {}

func NewOptionValueNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionValueNumberContext {
	var p = new(OptionValueNumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_optionValueNumber

	return p
}

func (s *OptionValueNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionValueNumberContext) INDEX() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserINDEX, 0)
}

func (s *OptionValueNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionValueNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OptionValueNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterOptionValueNumber(s)
	}
}

func (s *OptionValueNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitOptionValueNumber(s)
	}
}




func (p *DynamoDbGrammarParser) OptionValueNumber() (localctx IOptionValueNumberContext) {
	localctx = NewOptionValueNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, DynamoDbGrammarParserRULE_optionValueNumber)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(501)
		p.Match(DynamoDbGrammarParserINDEX)
	}



	return localctx
}


// IStarContext is an interface to support dynamic dispatch.
type IStarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStarContext differentiates from other interfaces.
	IsStarContext()
}

type StarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStarContext() *StarContext {
	var p = new(StarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_star
	return p
}

func (*StarContext) IsStarContext() {}

func NewStarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StarContext {
	var p = new(StarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_star

	return p
}

func (s *StarContext) GetParser() antlr.Parser { return s.parser }
func (s *StarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterStar(s)
	}
}

func (s *StarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitStar(s)
	}
}




func (p *DynamoDbGrammarParser) Star() (localctx IStarContext) {
	localctx = NewStarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, DynamoDbGrammarParserRULE_star)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(503)
		p.Match(DynamoDbGrammarParserT__6)
	}



	return localctx
}


// IHintContext is an interface to support dynamic dispatch.
type IHintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHintContext differentiates from other interfaces.
	IsHintContext()
}

type HintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHintContext() *HintContext {
	var p = new(HintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_hint
	return p
}

func (*HintContext) IsHintContext() {}

func NewHintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HintContext {
	var p = new(HintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_hint

	return p
}

func (s *HintContext) GetParser() antlr.Parser { return s.parser }

func (s *HintContext) IndexHint() IIndexHintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexHintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexHintContext)
}

func (s *HintContext) ScanInfo() IScanInfoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScanInfoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScanInfoContext)
}

func (s *HintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *HintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterHint(s)
	}
}

func (s *HintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitHint(s)
	}
}




func (p *DynamoDbGrammarParser) Hint() (localctx IHintContext) {
	localctx = NewHintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, DynamoDbGrammarParserRULE_hint)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(505)
		p.IndexHint()
	}
	p.SetState(507)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == DynamoDbGrammarParserENABLE {
		{
			p.SetState(506)
			p.ScanInfo()
		}

	}



	return localctx
}


// IIndexHintContext is an interface to support dynamic dispatch.
type IIndexHintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexHintContext differentiates from other interfaces.
	IsIndexHintContext()
}

type IndexHintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexHintContext() *IndexHintContext {
	var p = new(IndexHintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_indexHint
	return p
}

func (*IndexHintContext) IsIndexHintContext() {}

func NewIndexHintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexHintContext {
	var p = new(IndexHintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_indexHint

	return p
}

func (s *IndexHintContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexHintContext) USE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserUSE, 0)
}

func (s *IndexHintContext) IndexHintName() IIndexHintNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexHintNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexHintNameContext)
}

func (s *IndexHintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexHintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IndexHintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterIndexHint(s)
	}
}

func (s *IndexHintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitIndexHint(s)
	}
}




func (p *DynamoDbGrammarParser) IndexHint() (localctx IIndexHintContext) {
	localctx = NewIndexHintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, DynamoDbGrammarParserRULE_indexHint)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(509)
		p.Match(DynamoDbGrammarParserUSE)
	}
	{
		p.SetState(510)
		p.IndexHintName()
	}



	return localctx
}


// IIndexHintNameContext is an interface to support dynamic dispatch.
type IIndexHintNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexHintNameContext differentiates from other interfaces.
	IsIndexHintNameContext()
}

type IndexHintNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexHintNameContext() *IndexHintNameContext {
	var p = new(IndexHintNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_indexHintName
	return p
}

func (*IndexHintNameContext) IsIndexHintNameContext() {}

func NewIndexHintNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexHintNameContext {
	var p = new(IndexHintNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_indexHintName

	return p
}

func (s *IndexHintNameContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexHintNameContext) INDEXKEYWORD() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserINDEXKEYWORD, 0)
}

func (s *IndexHintNameContext) IndexName() IIndexNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexNameContext)
}

func (s *IndexHintNameContext) PRIMARY() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserPRIMARY, 0)
}

func (s *IndexHintNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexHintNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IndexHintNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterIndexHintName(s)
	}
}

func (s *IndexHintNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitIndexHintName(s)
	}
}




func (p *DynamoDbGrammarParser) IndexHintName() (localctx IIndexHintNameContext) {
	localctx = NewIndexHintNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, DynamoDbGrammarParserRULE_indexHintName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(516)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DynamoDbGrammarParserINDEXKEYWORD:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(512)
			p.Match(DynamoDbGrammarParserINDEXKEYWORD)
		}
		{
			p.SetState(513)
			p.IndexName()
		}


	case DynamoDbGrammarParserPRIMARY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(514)
			p.Match(DynamoDbGrammarParserPRIMARY)
		}
		{
			p.SetState(515)
			p.Match(DynamoDbGrammarParserINDEXKEYWORD)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IScanInfoContext is an interface to support dynamic dispatch.
type IScanInfoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScanInfoContext differentiates from other interfaces.
	IsScanInfoContext()
}

type ScanInfoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScanInfoContext() *ScanInfoContext {
	var p = new(ScanInfoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_scanInfo
	return p
}

func (*ScanInfoContext) IsScanInfoContext() {}

func NewScanInfoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScanInfoContext {
	var p = new(ScanInfoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_scanInfo

	return p
}

func (s *ScanInfoContext) GetParser() antlr.Parser { return s.parser }

func (s *ScanInfoContext) ENABLE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserENABLE, 0)
}

func (s *ScanInfoContext) SCAN() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserSCAN, 0)
}

func (s *ScanInfoContext) TotalSegment() ITotalSegmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITotalSegmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITotalSegmentContext)
}

func (s *ScanInfoContext) Segment() ISegmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISegmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISegmentContext)
}

func (s *ScanInfoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScanInfoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ScanInfoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterScanInfo(s)
	}
}

func (s *ScanInfoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitScanInfo(s)
	}
}




func (p *DynamoDbGrammarParser) ScanInfo() (localctx IScanInfoContext) {
	localctx = NewScanInfoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, DynamoDbGrammarParserRULE_scanInfo)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(518)
		p.Match(DynamoDbGrammarParserENABLE)
	}
	{
		p.SetState(519)
		p.Match(DynamoDbGrammarParserSCAN)
	}
	p.SetState(526)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == DynamoDbGrammarParserT__1 {
		{
			p.SetState(520)
			p.Match(DynamoDbGrammarParserT__1)
		}
		{
			p.SetState(521)
			p.TotalSegment()
		}
		{
			p.SetState(522)
			p.Match(DynamoDbGrammarParserT__0)
		}
		{
			p.SetState(523)
			p.Segment()
		}
		{
			p.SetState(524)
			p.Match(DynamoDbGrammarParserT__2)
		}

	}



	return localctx
}


// ITotalSegmentContext is an interface to support dynamic dispatch.
type ITotalSegmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTotalSegmentContext differentiates from other interfaces.
	IsTotalSegmentContext()
}

type TotalSegmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTotalSegmentContext() *TotalSegmentContext {
	var p = new(TotalSegmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_totalSegment
	return p
}

func (*TotalSegmentContext) IsTotalSegmentContext() {}

func NewTotalSegmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TotalSegmentContext {
	var p = new(TotalSegmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_totalSegment

	return p
}

func (s *TotalSegmentContext) GetParser() antlr.Parser { return s.parser }

func (s *TotalSegmentContext) INDEX() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserINDEX, 0)
}

func (s *TotalSegmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TotalSegmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TotalSegmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterTotalSegment(s)
	}
}

func (s *TotalSegmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitTotalSegment(s)
	}
}




func (p *DynamoDbGrammarParser) TotalSegment() (localctx ITotalSegmentContext) {
	localctx = NewTotalSegmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, DynamoDbGrammarParserRULE_totalSegment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(528)
		p.Match(DynamoDbGrammarParserINDEX)
	}



	return localctx
}


// ISegmentContext is an interface to support dynamic dispatch.
type ISegmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSegmentContext differentiates from other interfaces.
	IsSegmentContext()
}

type SegmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySegmentContext() *SegmentContext {
	var p = new(SegmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_segment
	return p
}

func (*SegmentContext) IsSegmentContext() {}

func NewSegmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SegmentContext {
	var p = new(SegmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_segment

	return p
}

func (s *SegmentContext) GetParser() antlr.Parser { return s.parser }

func (s *SegmentContext) INDEX() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserINDEX, 0)
}

func (s *SegmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SegmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SegmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterSegment(s)
	}
}

func (s *SegmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitSegment(s)
	}
}




func (p *DynamoDbGrammarParser) Segment() (localctx ISegmentContext) {
	localctx = NewSegmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, DynamoDbGrammarParserRULE_segment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(530)
		p.Match(DynamoDbGrammarParserINDEX)
	}



	return localctx
}


// IWhereContext is an interface to support dynamic dispatch.
type IWhereContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhereContext differentiates from other interfaces.
	IsWhereContext()
}

type WhereContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereContext() *WhereContext {
	var p = new(WhereContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_where
	return p
}

func (*WhereContext) IsWhereContext() {}

func NewWhereContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereContext {
	var p = new(WhereContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_where

	return p
}

func (s *WhereContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereContext) WHERE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserWHERE, 0)
}

func (s *WhereContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *WhereContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *WhereContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterWhere(s)
	}
}

func (s *WhereContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitWhere(s)
	}
}




func (p *DynamoDbGrammarParser) Where() (localctx IWhereContext) {
	localctx = NewWhereContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, DynamoDbGrammarParserRULE_where)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(532)
		p.Match(DynamoDbGrammarParserWHERE)
	}
	{
		p.SetState(533)
		p.condition(0)
	}



	return localctx
}


// IPrimaryKeyDeclContext is an interface to support dynamic dispatch.
type IPrimaryKeyDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryKeyDeclContext differentiates from other interfaces.
	IsPrimaryKeyDeclContext()
}

type PrimaryKeyDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryKeyDeclContext() *PrimaryKeyDeclContext {
	var p = new(PrimaryKeyDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_primaryKeyDecl
	return p
}

func (*PrimaryKeyDeclContext) IsPrimaryKeyDeclContext() {}

func NewPrimaryKeyDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryKeyDeclContext {
	var p = new(PrimaryKeyDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_primaryKeyDecl

	return p
}

func (s *PrimaryKeyDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryKeyDeclContext) PRIMARY() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserPRIMARY, 0)
}

func (s *PrimaryKeyDeclContext) KEY() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserKEY, 0)
}

func (s *PrimaryKeyDeclContext) IndexDecl() IIndexDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexDeclContext)
}

func (s *PrimaryKeyDeclContext) Capacity() ICapacityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICapacityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICapacityContext)
}

func (s *PrimaryKeyDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryKeyDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PrimaryKeyDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterPrimaryKeyDecl(s)
	}
}

func (s *PrimaryKeyDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitPrimaryKeyDecl(s)
	}
}




func (p *DynamoDbGrammarParser) PrimaryKeyDecl() (localctx IPrimaryKeyDeclContext) {
	localctx = NewPrimaryKeyDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, DynamoDbGrammarParserRULE_primaryKeyDecl)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(535)
		p.Match(DynamoDbGrammarParserPRIMARY)
	}
	{
		p.SetState(536)
		p.Match(DynamoDbGrammarParserKEY)
	}
	{
		p.SetState(537)
		p.IndexDecl()
	}
	p.SetState(539)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == DynamoDbGrammarParserCAPACITY {
		{
			p.SetState(538)
			p.Capacity()
		}

	}



	return localctx
}


// ISecondaryIndexDeclContext is an interface to support dynamic dispatch.
type ISecondaryIndexDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSecondaryIndexDeclContext differentiates from other interfaces.
	IsSecondaryIndexDeclContext()
}

type SecondaryIndexDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySecondaryIndexDeclContext() *SecondaryIndexDeclContext {
	var p = new(SecondaryIndexDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_secondaryIndexDecl
	return p
}

func (*SecondaryIndexDeclContext) IsSecondaryIndexDeclContext() {}

func NewSecondaryIndexDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SecondaryIndexDeclContext {
	var p = new(SecondaryIndexDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_secondaryIndexDecl

	return p
}

func (s *SecondaryIndexDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SecondaryIndexDeclContext) INDEXKEYWORD() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserINDEXKEYWORD, 0)
}

func (s *SecondaryIndexDeclContext) IndexName() IIndexNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexNameContext)
}

func (s *SecondaryIndexDeclContext) IndexDecl() IIndexDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexDeclContext)
}

func (s *SecondaryIndexDeclContext) SecondaryIndexType() ISecondaryIndexTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISecondaryIndexTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISecondaryIndexTypeContext)
}

func (s *SecondaryIndexDeclContext) AllAttributeDecl() []IAttributeDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttributeDeclContext)(nil)).Elem())
	var tst = make([]IAttributeDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttributeDeclContext)
		}
	}

	return tst
}

func (s *SecondaryIndexDeclContext) AttributeDecl(i int) IAttributeDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttributeDeclContext)
}

func (s *SecondaryIndexDeclContext) ProjectionIndex() IProjectionIndexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProjectionIndexContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProjectionIndexContext)
}

func (s *SecondaryIndexDeclContext) Capacity() ICapacityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICapacityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICapacityContext)
}

func (s *SecondaryIndexDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SecondaryIndexDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SecondaryIndexDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterSecondaryIndexDecl(s)
	}
}

func (s *SecondaryIndexDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitSecondaryIndexDecl(s)
	}
}




func (p *DynamoDbGrammarParser) SecondaryIndexDecl() (localctx ISecondaryIndexDeclContext) {
	localctx = NewSecondaryIndexDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, DynamoDbGrammarParserRULE_secondaryIndexDecl)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(542)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == DynamoDbGrammarParserGLOBAL || _la == DynamoDbGrammarParserLOCAL {
		{
			p.SetState(541)
			p.SecondaryIndexType()
		}

	}
	{
		p.SetState(544)
		p.Match(DynamoDbGrammarParserINDEXKEYWORD)
	}
	p.SetState(553)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(545)
			p.AttributeDecl()
		}
		p.SetState(550)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == DynamoDbGrammarParserT__0 {
			{
				p.SetState(546)
				p.Match(DynamoDbGrammarParserT__0)
			}
			{
				p.SetState(547)
				p.AttributeDecl()
			}


			p.SetState(552)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}


	}
	{
		p.SetState(555)
		p.IndexName()
	}
	{
		p.SetState(556)
		p.IndexDecl()
	}
	p.SetState(558)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == DynamoDbGrammarParserPROJECTION {
		{
			p.SetState(557)
			p.ProjectionIndex()
		}

	}
	p.SetState(561)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == DynamoDbGrammarParserCAPACITY {
		{
			p.SetState(560)
			p.Capacity()
		}

	}



	return localctx
}


// ISecondaryIndexTypeContext is an interface to support dynamic dispatch.
type ISecondaryIndexTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSecondaryIndexTypeContext differentiates from other interfaces.
	IsSecondaryIndexTypeContext()
}

type SecondaryIndexTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySecondaryIndexTypeContext() *SecondaryIndexTypeContext {
	var p = new(SecondaryIndexTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_secondaryIndexType
	return p
}

func (*SecondaryIndexTypeContext) IsSecondaryIndexTypeContext() {}

func NewSecondaryIndexTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SecondaryIndexTypeContext {
	var p = new(SecondaryIndexTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_secondaryIndexType

	return p
}

func (s *SecondaryIndexTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *SecondaryIndexTypeContext) GLOBAL() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserGLOBAL, 0)
}

func (s *SecondaryIndexTypeContext) LOCAL() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserLOCAL, 0)
}

func (s *SecondaryIndexTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SecondaryIndexTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SecondaryIndexTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterSecondaryIndexType(s)
	}
}

func (s *SecondaryIndexTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitSecondaryIndexType(s)
	}
}




func (p *DynamoDbGrammarParser) SecondaryIndexType() (localctx ISecondaryIndexTypeContext) {
	localctx = NewSecondaryIndexTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, DynamoDbGrammarParserRULE_secondaryIndexType)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(563)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DynamoDbGrammarParserGLOBAL || _la == DynamoDbGrammarParserLOCAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IIndexNameContext is an interface to support dynamic dispatch.
type IIndexNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexNameContext differentiates from other interfaces.
	IsIndexNameContext()
}

type IndexNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexNameContext() *IndexNameContext {
	var p = new(IndexNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_indexName
	return p
}

func (*IndexNameContext) IsIndexNameContext() {}

func NewIndexNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexNameContext {
	var p = new(IndexNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_indexName

	return p
}

func (s *IndexNameContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexNameContext) DdlName() IDdlNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDdlNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDdlNameContext)
}

func (s *IndexNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IndexNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterIndexName(s)
	}
}

func (s *IndexNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitIndexName(s)
	}
}




func (p *DynamoDbGrammarParser) IndexName() (localctx IIndexNameContext) {
	localctx = NewIndexNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, DynamoDbGrammarParserRULE_indexName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(565)
		p.DdlName()
	}



	return localctx
}


// IProjectionIndexContext is an interface to support dynamic dispatch.
type IProjectionIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProjectionIndexContext differentiates from other interfaces.
	IsProjectionIndexContext()
}

type ProjectionIndexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjectionIndexContext() *ProjectionIndexContext {
	var p = new(ProjectionIndexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_projectionIndex
	return p
}

func (*ProjectionIndexContext) IsProjectionIndexContext() {}

func NewProjectionIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjectionIndexContext {
	var p = new(ProjectionIndexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_projectionIndex

	return p
}

func (s *ProjectionIndexContext) GetParser() antlr.Parser { return s.parser }

func (s *ProjectionIndexContext) ProjectionIndexKeysOnly() IProjectionIndexKeysOnlyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProjectionIndexKeysOnlyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProjectionIndexKeysOnlyContext)
}

func (s *ProjectionIndexContext) ProjectionIndexVector() IProjectionIndexVectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProjectionIndexVectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProjectionIndexVectorContext)
}

func (s *ProjectionIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProjectionIndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ProjectionIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterProjectionIndex(s)
	}
}

func (s *ProjectionIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitProjectionIndex(s)
	}
}




func (p *DynamoDbGrammarParser) ProjectionIndex() (localctx IProjectionIndexContext) {
	localctx = NewProjectionIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, DynamoDbGrammarParserRULE_projectionIndex)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(569)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(567)
			p.ProjectionIndexKeysOnly()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(568)
			p.ProjectionIndexVector()
		}

	}


	return localctx
}


// IProjectionIndexKeysOnlyContext is an interface to support dynamic dispatch.
type IProjectionIndexKeysOnlyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProjectionIndexKeysOnlyContext differentiates from other interfaces.
	IsProjectionIndexKeysOnlyContext()
}

type ProjectionIndexKeysOnlyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjectionIndexKeysOnlyContext() *ProjectionIndexKeysOnlyContext {
	var p = new(ProjectionIndexKeysOnlyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_projectionIndexKeysOnly
	return p
}

func (*ProjectionIndexKeysOnlyContext) IsProjectionIndexKeysOnlyContext() {}

func NewProjectionIndexKeysOnlyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjectionIndexKeysOnlyContext {
	var p = new(ProjectionIndexKeysOnlyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_projectionIndexKeysOnly

	return p
}

func (s *ProjectionIndexKeysOnlyContext) GetParser() antlr.Parser { return s.parser }

func (s *ProjectionIndexKeysOnlyContext) PROJECTION() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserPROJECTION, 0)
}

func (s *ProjectionIndexKeysOnlyContext) KEYS() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserKEYS, 0)
}

func (s *ProjectionIndexKeysOnlyContext) ONLY() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserONLY, 0)
}

func (s *ProjectionIndexKeysOnlyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProjectionIndexKeysOnlyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ProjectionIndexKeysOnlyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterProjectionIndexKeysOnly(s)
	}
}

func (s *ProjectionIndexKeysOnlyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitProjectionIndexKeysOnly(s)
	}
}




func (p *DynamoDbGrammarParser) ProjectionIndexKeysOnly() (localctx IProjectionIndexKeysOnlyContext) {
	localctx = NewProjectionIndexKeysOnlyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, DynamoDbGrammarParserRULE_projectionIndexKeysOnly)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(571)
		p.Match(DynamoDbGrammarParserPROJECTION)
	}
	{
		p.SetState(572)
		p.Match(DynamoDbGrammarParserKEYS)
	}
	{
		p.SetState(573)
		p.Match(DynamoDbGrammarParserONLY)
	}



	return localctx
}


// IProjectionIndexVectorContext is an interface to support dynamic dispatch.
type IProjectionIndexVectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProjectionIndexVectorContext differentiates from other interfaces.
	IsProjectionIndexVectorContext()
}

type ProjectionIndexVectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjectionIndexVectorContext() *ProjectionIndexVectorContext {
	var p = new(ProjectionIndexVectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_projectionIndexVector
	return p
}

func (*ProjectionIndexVectorContext) IsProjectionIndexVectorContext() {}

func NewProjectionIndexVectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjectionIndexVectorContext {
	var p = new(ProjectionIndexVectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_projectionIndexVector

	return p
}

func (s *ProjectionIndexVectorContext) GetParser() antlr.Parser { return s.parser }

func (s *ProjectionIndexVectorContext) PROJECTION() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserPROJECTION, 0)
}

func (s *ProjectionIndexVectorContext) AllAttributeName() []IAttributeNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttributeNameContext)(nil)).Elem())
	var tst = make([]IAttributeNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttributeNameContext)
		}
	}

	return tst
}

func (s *ProjectionIndexVectorContext) AttributeName(i int) IAttributeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttributeNameContext)
}

func (s *ProjectionIndexVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProjectionIndexVectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ProjectionIndexVectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterProjectionIndexVector(s)
	}
}

func (s *ProjectionIndexVectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitProjectionIndexVector(s)
	}
}




func (p *DynamoDbGrammarParser) ProjectionIndexVector() (localctx IProjectionIndexVectorContext) {
	localctx = NewProjectionIndexVectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, DynamoDbGrammarParserRULE_projectionIndexVector)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(575)
		p.Match(DynamoDbGrammarParserPROJECTION)
	}
	{
		p.SetState(576)
		p.Match(DynamoDbGrammarParserT__1)
	}
	{
		p.SetState(577)
		p.AttributeName()
	}
	p.SetState(582)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == DynamoDbGrammarParserT__0 {
		{
			p.SetState(578)
			p.Match(DynamoDbGrammarParserT__0)
		}
		{
			p.SetState(579)
			p.AttributeName()
		}


		p.SetState(584)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(585)
		p.Match(DynamoDbGrammarParserT__2)
	}



	return localctx
}


// ICapacityContext is an interface to support dynamic dispatch.
type ICapacityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCapacityContext differentiates from other interfaces.
	IsCapacityContext()
}

type CapacityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCapacityContext() *CapacityContext {
	var p = new(CapacityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_capacity
	return p
}

func (*CapacityContext) IsCapacityContext() {}

func NewCapacityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CapacityContext {
	var p = new(CapacityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_capacity

	return p
}

func (s *CapacityContext) GetParser() antlr.Parser { return s.parser }

func (s *CapacityContext) CAPACITY() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserCAPACITY, 0)
}

func (s *CapacityContext) ReadUnits() IReadUnitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReadUnitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReadUnitsContext)
}

func (s *CapacityContext) WriteUnits() IWriteUnitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWriteUnitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWriteUnitsContext)
}

func (s *CapacityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CapacityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CapacityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterCapacity(s)
	}
}

func (s *CapacityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitCapacity(s)
	}
}




func (p *DynamoDbGrammarParser) Capacity() (localctx ICapacityContext) {
	localctx = NewCapacityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, DynamoDbGrammarParserRULE_capacity)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(587)
		p.Match(DynamoDbGrammarParserCAPACITY)
	}
	{
		p.SetState(588)
		p.Match(DynamoDbGrammarParserT__1)
	}
	{
		p.SetState(589)
		p.ReadUnits()
	}
	{
		p.SetState(590)
		p.Match(DynamoDbGrammarParserT__0)
	}
	{
		p.SetState(591)
		p.WriteUnits()
	}
	{
		p.SetState(592)
		p.Match(DynamoDbGrammarParserT__2)
	}



	return localctx
}


// IReadUnitsContext is an interface to support dynamic dispatch.
type IReadUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReadUnitsContext differentiates from other interfaces.
	IsReadUnitsContext()
}

type ReadUnitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReadUnitsContext() *ReadUnitsContext {
	var p = new(ReadUnitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_readUnits
	return p
}

func (*ReadUnitsContext) IsReadUnitsContext() {}

func NewReadUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReadUnitsContext {
	var p = new(ReadUnitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_readUnits

	return p
}

func (s *ReadUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *ReadUnitsContext) INDEX() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserINDEX, 0)
}

func (s *ReadUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ReadUnitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterReadUnits(s)
	}
}

func (s *ReadUnitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitReadUnits(s)
	}
}




func (p *DynamoDbGrammarParser) ReadUnits() (localctx IReadUnitsContext) {
	localctx = NewReadUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, DynamoDbGrammarParserRULE_readUnits)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(594)
		p.Match(DynamoDbGrammarParserINDEX)
	}



	return localctx
}


// IWriteUnitsContext is an interface to support dynamic dispatch.
type IWriteUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWriteUnitsContext differentiates from other interfaces.
	IsWriteUnitsContext()
}

type WriteUnitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWriteUnitsContext() *WriteUnitsContext {
	var p = new(WriteUnitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_writeUnits
	return p
}

func (*WriteUnitsContext) IsWriteUnitsContext() {}

func NewWriteUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WriteUnitsContext {
	var p = new(WriteUnitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_writeUnits

	return p
}

func (s *WriteUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *WriteUnitsContext) INDEX() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserINDEX, 0)
}

func (s *WriteUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WriteUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *WriteUnitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterWriteUnits(s)
	}
}

func (s *WriteUnitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitWriteUnits(s)
	}
}




func (p *DynamoDbGrammarParser) WriteUnits() (localctx IWriteUnitsContext) {
	localctx = NewWriteUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, DynamoDbGrammarParserRULE_writeUnits)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(596)
		p.Match(DynamoDbGrammarParserINDEX)
	}



	return localctx
}


// IIndexDeclContext is an interface to support dynamic dispatch.
type IIndexDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexDeclContext differentiates from other interfaces.
	IsIndexDeclContext()
}

type IndexDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexDeclContext() *IndexDeclContext {
	var p = new(IndexDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_indexDecl
	return p
}

func (*IndexDeclContext) IsIndexDeclContext() {}

func NewIndexDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexDeclContext {
	var p = new(IndexDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_indexDecl

	return p
}

func (s *IndexDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexDeclContext) HashKey() IHashKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHashKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHashKeyContext)
}

func (s *IndexDeclContext) RangeKey() IRangeKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeKeyContext)
}

func (s *IndexDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IndexDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterIndexDecl(s)
	}
}

func (s *IndexDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitIndexDecl(s)
	}
}




func (p *DynamoDbGrammarParser) IndexDecl() (localctx IIndexDeclContext) {
	localctx = NewIndexDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, DynamoDbGrammarParserRULE_indexDecl)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(598)
		p.Match(DynamoDbGrammarParserT__1)
	}
	{
		p.SetState(599)
		p.HashKey()
	}
	p.SetState(602)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == DynamoDbGrammarParserT__0 {
		{
			p.SetState(600)
			p.Match(DynamoDbGrammarParserT__0)
		}
		{
			p.SetState(601)
			p.RangeKey()
		}

	}
	{
		p.SetState(604)
		p.Match(DynamoDbGrammarParserT__2)
	}



	return localctx
}


// IAttributeDeclContext is an interface to support dynamic dispatch.
type IAttributeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttributeDeclContext differentiates from other interfaces.
	IsAttributeDeclContext()
}

type AttributeDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeDeclContext() *AttributeDeclContext {
	var p = new(AttributeDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_attributeDecl
	return p
}

func (*AttributeDeclContext) IsAttributeDeclContext() {}

func NewAttributeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeDeclContext {
	var p = new(AttributeDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_attributeDecl

	return p
}

func (s *AttributeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeDeclContext) AttributeName() IAttributeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributeNameContext)
}

func (s *AttributeDeclContext) AttributeType() IAttributeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttributeTypeContext)
}

func (s *AttributeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AttributeDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterAttributeDecl(s)
	}
}

func (s *AttributeDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitAttributeDecl(s)
	}
}




func (p *DynamoDbGrammarParser) AttributeDecl() (localctx IAttributeDeclContext) {
	localctx = NewAttributeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, DynamoDbGrammarParserRULE_attributeDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(606)
		p.AttributeName()
	}
	{
		p.SetState(607)
		p.AttributeType()
	}



	return localctx
}


// IHashKeyContext is an interface to support dynamic dispatch.
type IHashKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHashKeyContext differentiates from other interfaces.
	IsHashKeyContext()
}

type HashKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHashKeyContext() *HashKeyContext {
	var p = new(HashKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_hashKey
	return p
}

func (*HashKeyContext) IsHashKeyContext() {}

func NewHashKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HashKeyContext {
	var p = new(HashKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_hashKey

	return p
}

func (s *HashKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *HashKeyContext) DdlName() IDdlNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDdlNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDdlNameContext)
}

func (s *HashKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HashKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *HashKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterHashKey(s)
	}
}

func (s *HashKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitHashKey(s)
	}
}




func (p *DynamoDbGrammarParser) HashKey() (localctx IHashKeyContext) {
	localctx = NewHashKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 140, DynamoDbGrammarParserRULE_hashKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(609)
		p.DdlName()
	}



	return localctx
}


// IRangeKeyContext is an interface to support dynamic dispatch.
type IRangeKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangeKeyContext differentiates from other interfaces.
	IsRangeKeyContext()
}

type RangeKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeKeyContext() *RangeKeyContext {
	var p = new(RangeKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_rangeKey
	return p
}

func (*RangeKeyContext) IsRangeKeyContext() {}

func NewRangeKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeKeyContext {
	var p = new(RangeKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_rangeKey

	return p
}

func (s *RangeKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeKeyContext) DdlName() IDdlNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDdlNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDdlNameContext)
}

func (s *RangeKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RangeKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterRangeKey(s)
	}
}

func (s *RangeKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitRangeKey(s)
	}
}




func (p *DynamoDbGrammarParser) RangeKey() (localctx IRangeKeyContext) {
	localctx = NewRangeKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 142, DynamoDbGrammarParserRULE_rangeKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(611)
		p.DdlName()
	}



	return localctx
}


// IAttributeNameContext is an interface to support dynamic dispatch.
type IAttributeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttributeNameContext differentiates from other interfaces.
	IsAttributeNameContext()
}

type AttributeNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeNameContext() *AttributeNameContext {
	var p = new(AttributeNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_attributeName
	return p
}

func (*AttributeNameContext) IsAttributeNameContext() {}

func NewAttributeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeNameContext {
	var p = new(AttributeNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_attributeName

	return p
}

func (s *AttributeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeNameContext) DdlName() IDdlNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDdlNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDdlNameContext)
}

func (s *AttributeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AttributeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterAttributeName(s)
	}
}

func (s *AttributeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitAttributeName(s)
	}
}




func (p *DynamoDbGrammarParser) AttributeName() (localctx IAttributeNameContext) {
	localctx = NewAttributeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 144, DynamoDbGrammarParserRULE_attributeName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(613)
		p.DdlName()
	}



	return localctx
}


// IAttributeTypeContext is an interface to support dynamic dispatch.
type IAttributeTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttributeTypeContext differentiates from other interfaces.
	IsAttributeTypeContext()
}

type AttributeTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeTypeContext() *AttributeTypeContext {
	var p = new(AttributeTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_attributeType
	return p
}

func (*AttributeTypeContext) IsAttributeTypeContext() {}

func NewAttributeTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeTypeContext {
	var p = new(AttributeTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_attributeType

	return p
}

func (s *AttributeTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeTypeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserNUMBER, 0)
}

func (s *AttributeTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserSTRING, 0)
}

func (s *AttributeTypeContext) BINARY() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserBINARY, 0)
}

func (s *AttributeTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AttributeTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterAttributeType(s)
	}
}

func (s *AttributeTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitAttributeType(s)
	}
}




func (p *DynamoDbGrammarParser) AttributeType() (localctx IAttributeTypeContext) {
	localctx = NewAttributeTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 146, DynamoDbGrammarParserRULE_attributeType)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(615)
		_la = p.GetTokenStream().LA(1)

		if !(((((_la - 45)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 45))) & ((1 << (DynamoDbGrammarParserNUMBER - 45)) | (1 << (DynamoDbGrammarParserBINARY - 45)) | (1 << (DynamoDbGrammarParserSTRING - 45)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IReturningContext is an interface to support dynamic dispatch.
type IReturningContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturningContext differentiates from other interfaces.
	IsReturningContext()
}

type ReturningContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturningContext() *ReturningContext {
	var p = new(ReturningContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_returning
	return p
}

func (*ReturningContext) IsReturningContext() {}

func NewReturningContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturningContext {
	var p = new(ReturningContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_returning

	return p
}

func (s *ReturningContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturningContext) RETURNING() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserRETURNING, 0)
}

func (s *ReturningContext) ReturningValue() IReturningValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturningValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturningValueContext)
}

func (s *ReturningContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturningContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ReturningContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterReturning(s)
	}
}

func (s *ReturningContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitReturning(s)
	}
}




func (p *DynamoDbGrammarParser) Returning() (localctx IReturningContext) {
	localctx = NewReturningContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 148, DynamoDbGrammarParserRULE_returning)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(617)
		p.Match(DynamoDbGrammarParserRETURNING)
	}
	{
		p.SetState(618)
		p.ReturningValue()
	}



	return localctx
}


// IReturningValueContext is an interface to support dynamic dispatch.
type IReturningValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturningValueContext differentiates from other interfaces.
	IsReturningValueContext()
}

type ReturningValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturningValueContext() *ReturningValueContext {
	var p = new(ReturningValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_returningValue
	return p
}

func (*ReturningValueContext) IsReturningValueContext() {}

func NewReturningValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturningValueContext {
	var p = new(ReturningValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_returningValue

	return p
}

func (s *ReturningValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturningValueContext) NONE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserNONE, 0)
}

func (s *ReturningValueContext) ALL() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserALL, 0)
}

func (s *ReturningValueContext) OLD() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserOLD, 0)
}

func (s *ReturningValueContext) UPDATED() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserUPDATED, 0)
}

func (s *ReturningValueContext) NEW() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserNEW, 0)
}

func (s *ReturningValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturningValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ReturningValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterReturningValue(s)
	}
}

func (s *ReturningValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitReturningValue(s)
	}
}




func (p *DynamoDbGrammarParser) ReturningValue() (localctx IReturningValueContext) {
	localctx = NewReturningValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 150, DynamoDbGrammarParserRULE_returningValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(629)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(620)
			p.Match(DynamoDbGrammarParserNONE)
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(621)
			p.Match(DynamoDbGrammarParserALL)
		}
		{
			p.SetState(622)
			p.Match(DynamoDbGrammarParserOLD)
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(623)
			p.Match(DynamoDbGrammarParserUPDATED)
		}
		{
			p.SetState(624)
			p.Match(DynamoDbGrammarParserOLD)
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(625)
			p.Match(DynamoDbGrammarParserALL)
		}
		{
			p.SetState(626)
			p.Match(DynamoDbGrammarParserNEW)
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(627)
			p.Match(DynamoDbGrammarParserUPDATED)
		}
		{
			p.SetState(628)
			p.Match(DynamoDbGrammarParserNEW)
		}

	}


	return localctx
}


// IOnDuplicateKeyUpdateContext is an interface to support dynamic dispatch.
type IOnDuplicateKeyUpdateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOnDuplicateKeyUpdateContext differentiates from other interfaces.
	IsOnDuplicateKeyUpdateContext()
}

type OnDuplicateKeyUpdateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOnDuplicateKeyUpdateContext() *OnDuplicateKeyUpdateContext {
	var p = new(OnDuplicateKeyUpdateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_onDuplicateKeyUpdate
	return p
}

func (*OnDuplicateKeyUpdateContext) IsOnDuplicateKeyUpdateContext() {}

func NewOnDuplicateKeyUpdateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OnDuplicateKeyUpdateContext {
	var p = new(OnDuplicateKeyUpdateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_onDuplicateKeyUpdate

	return p
}

func (s *OnDuplicateKeyUpdateContext) GetParser() antlr.Parser { return s.parser }

func (s *OnDuplicateKeyUpdateContext) ON() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserON, 0)
}

func (s *OnDuplicateKeyUpdateContext) DUPLICATE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserDUPLICATE, 0)
}

func (s *OnDuplicateKeyUpdateContext) KEY() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserKEY, 0)
}

func (s *OnDuplicateKeyUpdateContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserUPDATE, 0)
}

func (s *OnDuplicateKeyUpdateContext) IfClause() IIfClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfClauseContext)
}

func (s *OnDuplicateKeyUpdateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OnDuplicateKeyUpdateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OnDuplicateKeyUpdateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterOnDuplicateKeyUpdate(s)
	}
}

func (s *OnDuplicateKeyUpdateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitOnDuplicateKeyUpdate(s)
	}
}




func (p *DynamoDbGrammarParser) OnDuplicateKeyUpdate() (localctx IOnDuplicateKeyUpdateContext) {
	localctx = NewOnDuplicateKeyUpdateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 152, DynamoDbGrammarParserRULE_onDuplicateKeyUpdate)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(631)
		p.Match(DynamoDbGrammarParserON)
	}
	{
		p.SetState(632)
		p.Match(DynamoDbGrammarParserDUPLICATE)
	}
	{
		p.SetState(633)
		p.Match(DynamoDbGrammarParserKEY)
	}
	{
		p.SetState(634)
		p.Match(DynamoDbGrammarParserUPDATE)
	}
	p.SetState(636)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == DynamoDbGrammarParserIF {
		{
			p.SetState(635)
			p.IfClause()
		}

	}



	return localctx
}


// IIfClauseContext is an interface to support dynamic dispatch.
type IIfClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfClauseContext differentiates from other interfaces.
	IsIfClauseContext()
}

type IfClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfClauseContext() *IfClauseContext {
	var p = new(IfClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_ifClause
	return p
}

func (*IfClauseContext) IsIfClauseContext() {}

func NewIfClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfClauseContext {
	var p = new(IfClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_ifClause

	return p
}

func (s *IfClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *IfClauseContext) IF() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserIF, 0)
}

func (s *IfClauseContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *IfClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IfClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterIfClause(s)
	}
}

func (s *IfClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitIfClause(s)
	}
}




func (p *DynamoDbGrammarParser) IfClause() (localctx IIfClauseContext) {
	localctx = NewIfClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 154, DynamoDbGrammarParserRULE_ifClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(638)
		p.Match(DynamoDbGrammarParserIF)
	}
	{
		p.SetState(639)
		p.condition(0)
	}



	return localctx
}


// ITableNameContext is an interface to support dynamic dispatch.
type ITableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableNameContext differentiates from other interfaces.
	IsTableNameContext()
}

type TableNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableNameContext() *TableNameContext {
	var p = new(TableNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_tableName
	return p
}

func (*TableNameContext) IsTableNameContext() {}

func NewTableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableNameContext {
	var p = new(TableNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_tableName

	return p
}

func (s *TableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TableNameContext) DdlName() IDdlNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDdlNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDdlNameContext)
}

func (s *TableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterTableName(s)
	}
}

func (s *TableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitTableName(s)
	}
}




func (p *DynamoDbGrammarParser) TableName() (localctx ITableNameContext) {
	localctx = NewTableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 156, DynamoDbGrammarParserRULE_tableName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(641)
		p.DdlName()
	}



	return localctx
}


// IDdlNameContext is an interface to support dynamic dispatch.
type IDdlNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDdlNameContext differentiates from other interfaces.
	IsDdlNameContext()
}

type DdlNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDdlNameContext() *DdlNameContext {
	var p = new(DdlNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_ddlName
	return p
}

func (*DdlNameContext) IsDdlNameContext() {}

func NewDdlNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DdlNameContext {
	var p = new(DdlNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_ddlName

	return p
}

func (s *DdlNameContext) GetParser() antlr.Parser { return s.parser }

func (s *DdlNameContext) ID() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserID, 0)
}

func (s *DdlNameContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *DdlNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DdlNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DdlNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterDdlName(s)
	}
}

func (s *DdlNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitDdlName(s)
	}
}




func (p *DynamoDbGrammarParser) DdlName() (localctx IDdlNameContext) {
	localctx = NewDdlNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 158, DynamoDbGrammarParserRULE_ddlName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(645)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DynamoDbGrammarParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(643)
			p.Match(DynamoDbGrammarParserID)
		}


	case DynamoDbGrammarParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(644)
			p.StringLiteral()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_stringLiteral
	return p
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserSTRING_LITERAL, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}




func (p *DynamoDbGrammarParser) StringLiteral() (localctx IStringLiteralContext) {
	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 160, DynamoDbGrammarParserRULE_stringLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(647)
		p.Match(DynamoDbGrammarParserSTRING_LITERAL)
	}



	return localctx
}


// IUnknownContext is an interface to support dynamic dispatch.
type IUnknownContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnknownContext differentiates from other interfaces.
	IsUnknownContext()
}

type UnknownContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnknownContext() *UnknownContext {
	var p = new(UnknownContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_unknown
	return p
}

func (*UnknownContext) IsUnknownContext() {}

func NewUnknownContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnknownContext {
	var p = new(UnknownContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_unknown

	return p
}

func (s *UnknownContext) GetParser() antlr.Parser { return s.parser }

func (s *UnknownContext) AllUNKNOWN() []antlr.TerminalNode {
	return s.GetTokens(DynamoDbGrammarParserUNKNOWN)
}

func (s *UnknownContext) UNKNOWN(i int) antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserUNKNOWN, i)
}

func (s *UnknownContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnknownContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *UnknownContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterUnknown(s)
	}
}

func (s *UnknownContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitUnknown(s)
	}
}




func (p *DynamoDbGrammarParser) Unknown() (localctx IUnknownContext) {
	localctx = NewUnknownContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 162, DynamoDbGrammarParserRULE_unknown)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(650)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == DynamoDbGrammarParserUNKNOWN {
		{
			p.SetState(649)
			p.Match(DynamoDbGrammarParserUNKNOWN)
		}


		p.SetState(652)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


func (p *DynamoDbGrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
			var t *ConditionContext = nil
			if localctx != nil { t = localctx.(*ConditionContext) }
			return p.Condition_Sempred(t, predIndex)


	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *DynamoDbGrammarParser) Condition_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
			return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
			return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

