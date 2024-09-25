//You are tasked to create a function that will produce a Caesar Cipher of your text. 
// A Caesar Cipher works by replacing letters by a set number of spaces in the alphabet. 
// For example a Caesar Cipher with an offset of 1, would turn "ABC" into "BCD", and offset of 2 would turn "ABC" into "CDE". 
// If the letter is past the end of the alphabet, it loops around to the begging.
// You goal is have a function that can output the string "THIS IS MY CODING CHALLENGE" in both a 5 and 15 letter offset.

function ceasarCipher(text, offset) {
    const alphabets = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
    const offsetMOd = offset % 26;
    return text.toUpperCase().split('').map(char => {
        if (alphabets.includes(char)) {
            const index = alphabets.indexOf(char);
            const newIndex = (index + offsetMOd) % 26;
            return alphabets[newIndex < 0 ? newIndex + 26 : newIndex]
        }
        else {
            return char;
        }
    }).join('');
} 

const text = 'THIS IS MY CODING CHALLENEGE';

const offset5 = ceasarCipher(text, 5);
console.log(offset5);