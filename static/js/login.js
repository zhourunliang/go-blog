var do_login = function(form) {

    var data = JSON.stringify(form)
    var request = {
        method: 'POST',
        url: '/api/user/do_login',
        data: data,
        contentType: 'application/json',
        callback: function(response) {
            // log('响应', response)
            var res = JSON.parse(response)
            if (res.status) {
                showAlert(res.msg)
                self.location.href="/"
            }else{
                showAlert(res.msg, false)
            }
        }
    }
    ajax(request)
}

var e = function(selector) {
    return document.querySelector(selector)
}

var bindEvents = function() {
    // 绑定发表新留言事件
    var button = e('#id-button-submit')
    button.addEventListener('click', function(event){
        // log('click new')
        // 得到用户填写的数据
        var form = {
            name: e('#id-input-name').value,
            password: e('#id-input-password').value,
        }
        if (strIsEmpty(form.name)) {
            showAlert('请输入姓名', false)
            return false
        }else if(strIsEmpty(form.password)){
            showAlert('请输入密码', false)
            return false
        }
        do_login(form)
    })
}

var __main = function() {
    // 绑定事件
    bindEvents()
}

__main()
