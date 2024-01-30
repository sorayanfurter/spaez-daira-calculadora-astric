

export default function Mask(mask: string, callbac?: any) {
    let content = '';
    let maxChars = numberCharactersPattern(mask);

    return (e: any) => {
        e.preventDefault()
        if (isNumeric(e.key) && content.length < maxChars) {
            content += e.key;
        }

        if (e.keyCode == 8) {
            if (content.length > 0) {
                content = content.substr(0, content.length - 1);
            }
        }

        e.target.value = maskIt(mask, content)
        if (callbac) {
            callbac(content == "" ? 0 : parseFloat(content))
        }
        return
    }
}

function numberCharactersPattern(pattern: any) {
    let numberChars = 0;
    for (let i = 0; i < pattern.length; i++) {
        if (pattern[i] === '#') {
            numberChars++;
        }
    }
    return numberChars;
}

function maskIt(pattern: string, value: string) {
    let position = 0;
    let currentChar = 0;
    let masked = '';
    while (position < pattern.length && currentChar < value.length) {
        if (pattern[position] === '#') {
            masked += value[currentChar];
            currentChar++;
        } else {
            masked += pattern[position];
        }
        position++;
    }
    return masked;
}

function isNumeric(char: any) {
    return !isNaN(char - parseInt(char));
}


