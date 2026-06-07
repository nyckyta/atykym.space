const handleTerminalSwitch = event => {
    let currentInput = document.getElementById('tty-main')
    let ttyClone = currentInput.cloneNode(true)
    ttyClone.setAttribute('id', '')
    ttyClone.classList.add('tty-history')

    let editableArea = ttyClone.children[1]
    editableArea.setAttribute('contenteditable', false)
    editableArea.setAttribute('id', '')
    editableArea.setAttribute('autofocus', false)

    currentInput.parentElement.insertBefore(ttyClone, currentInput)
    currentInput.children[1].textContent = ''

    // remove carret
    let carret = document.getElementById('carret')
    carret.innerHTML = ''
    carret.setAttribute('id', '')
}

const clearTerminal = () => {
    document.querySelectorAll('.tty-history').forEach(history => history.remove())

    // deferred so it runs after the keydown handler that appends the <br> on Enter
    setTimeout(() => {
        let input = document.getElementById('tty-input')
        input.textContent = ''

        let carret = document.getElementById('carret')
        carret.style.right = '0px'

        input.focus()
    }, 0)
}

const calculateAndChangeCarretPosition = selection => {
    let lengthOfInput = document.getElementById('tty-input').textContent.length
    let startOffset = selection.getRangeAt(0).startOffset
    let endOffset = selection.getRangeAt(0).endOffset
    if (startOffset !== endOffset) {
        return
    }

    let shift = lengthOfInput - (startOffset > lengthOfInput ? lengthOfInput : startOffset)
    let carret = document.getElementById('carret')
    let carretStyles = carret.style
    let carretWidth = carret.getBoundingClientRect().width
    carretStyles.right = `${shift * carretWidth}px`
}

window.addEventListener('load', () => {
    // makes focusing on the input simpler
    window.addEventListener('click', () => {
        document.getElementById('tty-input').focus()
    })

    // handles move of the carret without input change
    document.addEventListener("selectionchange", event => {
        let selection = window.getSelection()
        let isTextselectionRelatedToInput = selection.focusNode.parentElement.isSameNode(document.getElementById('tty-input'))
        if (isTextselectionRelatedToInput) {
            calculateAndChangeCarretPosition(selection)
        }
    })

    // handles redraw of the terminal after response from the server is received
    document.getElementById('tty-input').addEventListener('htmx:afterRequest', event => {
        handleTerminalSwitch(event)
    })

    // handles 'clear' locally instead of sending it to the server
    document.getElementById('tty-input').addEventListener('htmx:confirm', event => {
        let cmd = event.target.textContent.trim()
        if (cmd === 'clear') {
            event.preventDefault()
            clearTerminal()
        }
    })

    // handles UI behavior when enter is clicked and input is processed
    document.getElementById('tty-input').addEventListener('keydown', event => {
        if (event.code === 'Enter') {
            event.preventDefault()
            document.getElementById('tty-input').insertAdjacentHTML('beforeend', "<br>")
        }
    })

    // handles move of the carret when text is changed
    document.getElementById('tty-input').addEventListener('input', e => {
        let selection = window.getSelection()
        calculateAndChangeCarretPosition(selection)
    })
})