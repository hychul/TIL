$(function() {
    if (!window.EventSource) {
        alert("This browser doesn't support event source");
        return;
    }

    var $chatlog = $('#chat-log');
    var $chatmsg = $('#chat-msg');

    var isBlank = function(string) {
        return string == null || string.trim() === "";
    };

    var username;
    while (isBlank(username)) {
        username = prompt("What's user name?");
        if (!isBlank(username)) {
            $('#user-name').html('<b>' + username + '</b>');
        }
    }

    $('#input-form').on('submit', function(e) {
        $.post('/messages', {
            sender: username,
            msg: $chatmsg.val()
        });

        $chatmsg.val("");
        $chatmsg.focus();
        return false;
    });

    var addMessage = function(data) {
        var text = "";

        if (!isBlank(data.sender)) {
            text = '<strong>' + data.sender + ':</strong>';
        }
        text += data.msg;

        $chatlog.prepend('<div><span>' + text + '</span></div>');
    }

    var es = new EventSource('/stream')
    es.onopen = function(e) {
        $.post('/users', {
            name: username
        })
    }
    es.onmessage = function(e) {
        var msg = JSON.parse(e.data);
        addMessage(msg);
    }

    window.onbeforeunload = function() {
        $.ajax({
            url: "/users?name=" + username,
            type: "DELETE"
        });
        es.close();
    }
})