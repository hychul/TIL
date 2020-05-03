$(function() {
    if (!window.EventSource) {
        alert("This browser doesn't support event source")
        return
    }

    var $chatlog = $('#chat-log')
    var $chatmsg = $('#chat-msg')

    var isBlank = function(string) {
        return string == null || string.trim() === "";
    };

    var username;
    while (isBlank(username)) {
        username = prompt("What's user name?");
        if (!isBlank(username)) {
            $('#user-name').html('<b>' + username + '</b>')
        }
    }

    $('#input-form').on('submit', function(e) {
        $.post('/messages', {
            sender: username,
            msg: $chatmsg.val()
        });

        $chatmsg.val("")
        $chatmsg.focus()
    });
})