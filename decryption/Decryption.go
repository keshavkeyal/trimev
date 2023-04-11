package decryption

import (
	"fmt"
)

var (
	//define the 100 character key
	key = []rune("abcdefghiklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!§$%&/() =?*'<> #|;²³~ @`´.abcdefghiklmnkolijklmo")
	//Charaxter array
	//carray = []string{"😀", "😃", "😄", "😁", "😆", "😅", "😂", "🤣", "🥲", "😊", "😇", "🙂", "🙃", "😉", "😌", "😍", "🥰", "😘", "😗", "😙", "😚", "😋", "😛", "😝", "😜", "🤪", "🤨", "🧐", "🤓", "😎", "🥸", "🤩", "🥳", "😏", "😒", "😞", "😔", "😟", "😕", "🙁", "☹️", "😣", "😖", "😫", "😩", "🥺", "😢", "😭", "😤", "😠", "😡", "🤬", "🤯", "😳", "🥵", "🥶", "😱", "😨", "😰", "😥", "😓", "🤗", "🤔", "🤭", "🤫", "🤥", "😶", "😐", "😑", "😬", "🙄", "😯", "😦", "😧", "😮", "😲", "🥱", "😴", "🤤", "😪", "😵", "🤐", "🥴", "🤢", "🤮", "🤧", "😷", "🤒", "🤕", "🥺", "👶", "🧒", "👦", "👧", "🧑", "👱", "👨", "🧔", "👩", "👱‍♀️", "👱‍♂️", "🧓", "👴", "👵", "🙍", "🙎", "🙅", "🙆", "💁", "🙋", "🧏", "🤷", "🤦", "👮", "🕵️", "💂", "👷", "👸", "👳", "👲", "🧕", "🤴", "👰", "🤵", "👼", "🤰", "🦸", "🦹", "🧙", "🧚", "🧜", "🧞", "🧟", "💆", "💇", "🚶", "🧍", "🧎", "🏃", "💃", "🕺", "🕴️", "👯", "🧖", "🧘", "👫", "👭", "👬", "💏", "👩‍❤️‍💋‍👨", "👨‍❤️‍💋‍👨", "👩‍❤️‍💋‍👩", "💑", "👩‍❤️‍👨", "👨‍❤️‍👨", "👩‍❤️‍👩", "👪", "👨‍👩‍👦", "👨‍👩‍👧", "👨‍👩‍👧‍👦", "👨‍👩‍👦‍👦", "👨‍👩‍👧‍👧", "○", "◙", "♂", "♀", "♪", "♫", "☼", "►", "◄", "↕️", "‼️", "¶", "§", "▬", "↨", "↑", "↓", "→", "←", "∟", "↔️", "▲", "▼", " ", "!", "\"", "#", "$", "%", "&", "'", "(", ")", "*", "+", ",", "-", ".", "/", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ":", ";", "<", "=", ">", "?", "@", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "[", "\\", "]", "^", "_", "`", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "{", "|", "}", "~", "⌂", "Ç", "ü", "é", "â", "ä", "à", "å", "ç", "ê", "ë", "è", "ï", "î", "ì", "Ä", "Å", "É", "æ", "Æ", "ô", "ö", "ò", "û", "ù", "ÿ", "Ö", "Ü", "ø", "£", "Ø", "×", "ƒ", "á", "í", "ó", "ú", "ñ", "Ñ", "ª", "º", "¿", "®", "¬", "½", "¼", "¡", "«", "»", "░", "▒", "▓", "│", "┤", "╡", "╢", "╖", "╕", "╣", "║", "╗", "╝", "╜", "╛", "┐", "└", "┴", "┬", "├", "─", "┼", "╞", "╟", "╚", "╔", "╩", "╦", "╠"}
	// defining the keysymbols for expression key for dynamic key
	keysymbols = []string{"+", "-", "+", "+", "-", "^", "-", "+", "+", "-", "^", "-", "+", "+", "-", "+", "-", "+", "^", "/"}
	//define the plain text block size
	ptBlockSize = 10
	//define the key block size
	keyBlockSize = 5
	//
	blockSize = 10
	//key ascii code
	keyasciicode [][]rune
	//subkeys array
	subkeys []rune
	//Mathematical symbol array
	symbols []string
	//dynamic number
	dynamicnum = 8
	//variable which is globally used
	variable              = 0
	expresstionblockarray = []expressionblock{}
	//newley genrated expression
	expressionarray []expressionblock
	//String for decrytption
	encryptedText string
	intarray      []int
	const1        = 66
	//constant  value 2
	const2 = 45
	blocks [][]int
	bl     []string
)

type expressionblock struct {
	symbol string
	number int
}

func Decryption(k string) {
	encryptedText = k
	//fmt.Println(Gettingindexvalues())
	blocks = create2DArray(Gettingindexvalues())
	fmt.Println(blocks)
	CreatingSubKey()
	//Calling the CreatingExpressionBLocks function and creating the expression blocks
	CreatingExpressionBLocks()
	//Encoding with Railfence & calculating the variable
	variablecalc(EncodeWithRailFence(expresstionblockarray, dynamicnum))
	//Encrypting text
	DecrytptingText()

}
func Gettingindexvalues() []int {
	for _, char := range encryptedText {
		codePoint := int(char)
		fmt.Println(codePoint)
		intarray = append(intarray, codePoint)
	}
	return intarray
}
func create2DArray(inputArr []int) [][]int {
	var twoDArr [][]int

	for i := 0; i < len(inputArr); i += 10 {
		if i+10 <= len(inputArr) {
			twoDArr = append(twoDArr, inputArr[i:i+10])
		} else {
			twoDArr = append(twoDArr, inputArr[i:])
		}
	}

	return twoDArr
}
func CreatingSubKey() {
	for i := 0; i < len(key); i += keyBlockSize {
		end := i + keyBlockSize
		if end > len(key) {
			end = len(key)
		}
		keyasciicode = append(keyasciicode, key[i:end])
	}
	for k := 0; k < 20; k++ {
		subkey := (keyasciicode[k][0] ^ keyasciicode[k][1]) ^ (keyasciicode[k][2] ^ keyasciicode[k][3] ^ keyasciicode[k][4])
		subkeys = append(subkeys, subkey)
	}
	fmt.Println(subkeys)
}

// rail fence cipher coding
func fence(input []expressionblock, rail int) [][]expressionblock {
	matrix := make([][]expressionblock, rail)
	for i := range matrix {
		matrix[i] = make([]expressionblock, 0, len(input)/2)
	}
	var d bool
	i := 0
	for _, c := range input {
		matrix[i] = append(matrix[i], c)
		switch d {
		case false:
			i++
			if i == rail-1 {
				d = true
			}
		case true:
			i--
			if i == 0 {
				d = false
			}
		}
	}
	return matrix
}

// Encode returns rail fence encoding with railfence of subkeys
func EncodeWithRailFence(input []expressionblock, rail int) []expressionblock {
	matrix := fence(input, rail)
	for _, row := range matrix {
		expressionarray = append(expressionarray, row...)
	}
	expresstionblockarray = expressionarray
	expressionarray = nil
	//fmt.Println(expresstionblockarray)
	return expresstionblockarray
}

// function for creating expression blocks
func CreatingExpressionBLocks() {
	for i := range keysymbols {
		expresstionblockarray = append(expresstionblockarray, expressionblock{symbol: keysymbols[i], number: int(subkeys[i])})
	}
}

// function for getting new variable
func variablecalc(expresstionblockarr []expressionblock) int {
	variablevalue := 0
	for i := range expresstionblockarr {

		switch expresstionblockarr[i].symbol {
		case "+":
			variablevalue = variablevalue + expresstionblockarr[i].number
		case "-":
			variablevalue = variablevalue - expresstionblockarr[i].number
		case "/":
			variablevalue = variablevalue / expresstionblockarr[i].number

		case "*":
			variablevalue = variablevalue * expresstionblockarr[i].number

		case "^":
			variablevalue = variablevalue ^ expresstionblockarr[i].number
		default:
		}

	}
	variable = variablevalue ^ const1 ^ const2 ^ dynamicnum
	return variablevalue
}

// getting dynamic number
func gettingdynamicnum(num int) int {
	res := 0
	for num > 0 {
		res ^= num % 10
		num /= 10
	}
	dynamicnum = res
	return res
}

// Encrypting the text
func DecrytptingText() {
	for k := 0; k < len(blocks); k++ {
		for i := 0; i < len(blocks[k]); i++ {
			//fmt.Println((blocks[k][i]) ^ (variable) ^ (i))
			mon := blocks[k][i] ^ variable ^ i
			bl = append(bl, string(rune(int(blocks[k][i]^variable^i))))
			//fmt.Println(bl)
			if i == 9 {
				dynamicnum = gettingdynamicnum(int(mon ^ (variable) ^ (i)))
			}
		}
		variable = variablecalc(EncodeWithRailFence(expresstionblockarray, dynamicnum)) ^ const1 ^ const2 ^ dynamicnum
		fmt.Println("-----------------------------------------------------------")

	}
	fmt.Println(bl)
}
