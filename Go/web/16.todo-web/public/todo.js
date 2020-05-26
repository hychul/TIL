(function($) {
    'use strict';
    $(function() {
        var todoListItem = $('.todo-list');
        var todoListInput = $('.todo-list-input');
        $('.todo-list-add-btn').on("click", function(event) {
            event.preventDefault();
            
            var item = $(this).prevAll('.todo-list-input').val();
            
            if (item) {
                $.post("todos", {name:item}, addItem);
                // todoListItem.append("<li><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' />" + item + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
                todoListInput.val("");
            }
        });

        var addItem = function(item) {
            if (item.completed) {
                todoListItem.append("<li" + " id='" + item.id + " 'class='completed'><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' checked='checked'/>" + item.name + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
            } else {
                todoListItem.append("<li" + " id='" + item.id + "'><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' />" + item.name + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
            }
        };

        $.get('/todos', function(items) {
            items.forEach(item => {
                addItem(item)
            });
        });
        
        todoListItem.on('change', '.checkbox', function() {
            if ($(this).attr('checked')) {
                $(this).removeAttr('checked');
            } else {
                $(this).attr('checked', 'checked');
            }
        
            $(this).closest("li").toggleClass('completed');
        });
        
        todoListItem.on('click', '.remove', function() {
            var id = $(this).closest("li").attr('id')
            var $self = $(this)
            $.ajax({
                url: "todos/" + id,
                type: "DELETE",
                success: function(e) {
                    if (e.success) {
                        $self.parent().remove();
                    }
                }
            })
        });
    });
})(jQuery);