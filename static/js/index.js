const handleTerminalSwitch = event => {
    let currentInput = document.getElementById('tty-main')
    let ttyClone = currentInput.cloneNode(true)
    ttyClone.setAttribute('id', '')

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

window.addEventListener('load', () => {
    window.addEventListener('click', () => {
        document.getElementById('tty-input').focus()
    })

    document.addEventListener("selectionchange", event => {
        let selection = window.getSelection()
        let isTextselectionRelatedToInput = selection.focusNode.parentElement.isSameNode(document.getElementById('tty-input'))
        if (isTextselectionRelatedToInput) {
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
    })

    document.getElementById('tty-input').addEventListener('htmx:afterRequest', event => {
        handleTerminalSwitch(event)
    })

    document.getElementById('tty-input').addEventListener('keydown', event => {
        if (event.code === 'Enter') {
            event.preventDefault()
            document.getElementById('tty-input').insertAdjacentHTML('beforeend', "<br>")
        }
    })

    document.getElementById('tty-input').addEventListener('input', e => {
        let range = window.getSelection().getRangeAt(0)
        // let editable = document.getElementById('tty-input')
        // let textNode = editable.firstChild
        // if (textNode === undefined || textNode === null) {
        //     return
        // }
        // let range = document.createRange()
        // let sel = window.getSelection()
        // let contentLen = editable.textContent.length
        // range.setStart(textNode, contentLen === 0 ? 0 : contentLen)
        // range.collapse(true)
    
        // sel.removeAllRanges()
        // sel.addRange(range)
    })
})