var email = "{{.Email}}";

// Установка соединения WebSocket с сервером
const socket = new WebSocket("ws://" + location.host + "/ws");

// Обработчик события при открытии соединения
socket.onopen = function (event) {
    console.log("WebSocket connection established");
    handleMessageButtonClick();
};

// Обработчик события при закрытии соединения
socket.onclose = function (event) {
    console.log("WebSocket connection closed");
};

// Обработчик события при получении нового сообщения от сервера
socket.onmessage = function (event) {
    const message = JSON.parse(event.data);
    const messagesContainer = document.getElementById("messages");
    const newMessage = document.createElement("div");

    // Проверяем, является ли сообщение исходящим (от текущего пользователя)
    if (message.outgoing) {
        newMessage.className = "message outgoing";
        newMessage.innerHTML = `
            <p style="background-color: #ff90b5">${message.text}</p>
            <div class="message-info">
                <span class="message-name">For you</span>
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

    // Добавляем новое сообщение в контейнер
    messagesContainer.appendChild(newMessage);
};


  
  

// Код выше добавляет обработчики событий для отправки и получения сообщений чата на клиентской стороне. Когда пользователь отправляет новое сообщение, оно отправляется с использованием WebSocket на сервер, а затем сервер отправляет сообщение всем подключенным клиентам, которые показывают его на своих страницах.

// Обратите внимание, что код выше является примером и его нужно адаптировать к вашим конкретным потребностям и указаниям серверной части на Go.

// Не забудьте также обновить путь к статическим файлам (CSS и JS) в разделе <link> и <script>, чтобы они указывали на соответствующие файлы на вашем сервере.

// Учтите, что WebSocket требует поддержки со стороны сервера, поэтому вы также должны внести изменения на серверной стороне, чтобы обрабатывать WebSocket-запросы.


function sendMessage() {
    const messageInput = document.getElementById("message-input");
    const text = messageInput.value;

    // Проверяем, что текст сообщения не пустой
    if (text.trim() !== "" && text.length < 1500) {
        const message = {
            message: text,
            email: email,
            outgoing: true,  // Устанавливаем флаг исходящего сообщения
        };

        // Отправляем сообщение на сервер
        socket.send(JSON.stringify(message));

        // Очищаем поле ввода сообщения
        messageInput.value = "";
        
        const messagesContainer = document.getElementById("messages");
        const newMessage = document.createElement("div");
        newMessage.className = "message outgoing";
        newMessage.innerHTML = `
            <p>${text}</p>
            <div class="message-info">
                <span class="message-name">You</span>
            </div>
        `;
        messagesContainer.appendChild(newMessage);
    }
  }


  function handleMessageButtonClick() {
    // Получение кнопок
    let buttons = document.querySelectorAll('.Wbutton');
  
    buttons.forEach(function(button) {
        button.addEventListener('click', function() {
            let message = this.dataset.message;

            const data = {
                message: message,
                email: email, // Замените на получение email пользователя
                outgoing: false,
              };
            // Отправка данных на сервер
            socket.send(JSON.stringify(data));

            // Скрытие кнопки
            this.style.display = 'none';
        });
    });
}