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
}

window.addEventListener('load', () => {
    window.addEventListener('click', () => {
        document.getElementById('tty-input').focus()
    }, false)

    document.getElementById('tty-input').addEventListener('htmx:afterRequest', event => {
        handleTerminalSwitch(event)
    })
    
    document.getElementById('tty-input').addEventListener('keydown', event => {
        if (event.code === 'Enter') {
            event.preventDefault()
            document.getElementById('tty-input').insertAdjacentHTML('beforeend', "<br>")
        }
    })
})