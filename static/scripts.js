
const socket = new WebSocket("ws://" + "localhost:8081" + "/ws");

socket.onopen = function (event) {
    console.log("WebSocket connection established");
};

socket.onclose = function (event) {
    console.log("WebSocket connection closed");
};

socket.onmessage = function (event) {
    const message = JSON.parse(event.data);
    const messagesContainer = document.getElementById("messages");
    const newMessage = document.createElement("div");

    if (message.err) {
        newMessage.className = "message outgoing";
        newMessage.innerHTML = `
            <p style="background-color: #ff90b5">${message.text}</p>
            <div class="message-info">
                <span class="message-name">For you</span>
            </div>
        `;
    } else if (message.outgoing) {
        newMessage.className = "message outgoing";
        newMessage.innerHTML = `
            <p style="background-color: #89dcf5">${message.text}</p>
            <div class="message-info">
                <span class="message-name">You</span>
            </div>
        `;
    } else {
        newMessage.className = "message incoming";
        newMessage.innerHTML = `
            <div class="message-info">
                <span class="message-name">${message.username}</span>
            </div>
            <p>${message.text}</p>
            <button class="Wbutton" data-message="${message.text}">W</button>
        `;
    }

    messagesContainer.appendChild(newMessage);
    handleMessageButtonClick();
};




function sendMessage() {
    const messageInput = document.getElementById("message-input");
    const text = messageInput.value;

    if (text.trim() !== "") {
        const message = {
            message: text,
            id: id,
            outgoing: true, 
        };

        socket.send(JSON.stringify(message));

        messageInput.value = "";
    }
}


function handleMessageButtonClick() {
    let buttons = document.querySelectorAll('.Wbutton');

    buttons.forEach(function (button) {
        button.addEventListener('click', function () {
            let message = this.dataset.message;

            const data = {
                message: message,
                id: id,
                outgoing: false,
            };
            socket.send(JSON.stringify(data));

            this.style.display = 'none';
        });
    });
}



const messagesContainer = document.getElementById("messages");
const newMessage = document.createElement("div");
newMessage.className = "message incoming";
newMessage.innerHTML = `
            <div class="message-info">
                <span class="message-name">${message.username}</span>
            </div>
            <p>${message.text}</p>
            <button class="Wbutton" data-message="${message.text}">W</button>
        `;
messagesContainer.appendChild(newMessage);