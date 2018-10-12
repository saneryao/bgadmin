<!-- Custom Scripts -->
<script src="/static/assets/js/jquery.validate.min.js"></script>
<script src="/static/js/jquery.validate.{{.Lang}}.js"></script>
<!--
<script src="/static/js/jquery.serializejson.min.js"></script>
-->
<script src="/static/js/jquery.submit.js"></script>

<script src="/static/js/jsencrypt.min.js"></script>
<script src="/static/js/jquery.pwd.js"></script>

<!-- inline scripts related to this page -->
<script type="text/javascript">
    jQuery(function($) {
        $(document).on('click', '.toolbar a[data-target]', function(e) {
            e.preventDefault();
            var target = $(this).data('target');
            $('.widget-box.visible').removeClass('visible');//hide others
            $(target).addClass('visible');//show target
        });
    
        // Just used for changing background
        $('body').attr('class', 'login-layout light-login');
        $('#id-text2').attr('class', 'grey');
        $('#id-company-text').attr('class', 'blue');

        //
        refresh_captcha(false);
        $("#form-resetpwd").submit( function(){
            return submit_ajax($("#form-login"));
        });
    });

    function refresh_captcha(random) {
        var obj = $("input[name='captcha_id']");
        var id = obj.val();
        var captcha = '';
        if (random) {
            captcha = '/captcha/' + id + '.png?reload=' + (new Date()).getTime();
        } else {
            captcha = '/captcha/' + id + '.png';
        }
        $('#id-login').val(id);
        $('#id-register').val(id);
        $('#id-forget').val(id);
        $('#captcha-login').attr('src', captcha);
        $('#captcha-register').attr('src', captcha);
        $('#captcha-forget').attr('src', captcha);
    }

    function submit_ajax(form) {
        $.ajax({
            async: true,
            url: $(form).attr('action'),
            type: $(form).attr('method'),
            dataType: 'json',
            data: getFormJson($(form)),
            contentType: 'application/json; charset=UTF-8',
            beforeSend: function (xhr) {
            },
            complete: function (xhr, ts) {
            },
            success: function (data) {
                if (data.code) {
                    // $('#msg-title').html('{{i18n .Lang "login"}}{{i18n .Lang "success"}}');
                    // $('#msg-content').html(data.msg);
                    // $('#msg-dialog').modal('show');
                    window.location.href = data.url;
                } else {
                    $('#msg-title').html('{{i18n .Lang "login"}}{{i18n .Lang "failed"}}');
                    $('#msg-content').html(data.error);
                    $('#msg-dialog').modal('show');
                    refresh_captcha(true);
                }
            },
            error: function (xhr, msg, e) {
                console.log(xhr.status);          // 状态码
                console.log(xhr.readyState);  // 状态
                console.log(msg);                 // 错误信息
                alert("" + xhr.status + msg);
            }
        });
        return false;
    }

    function getFormJson($form) {
        var unindexArr = $form.serializeArray();
        var indexArr = {};
        $.map(unindexArr, function (n, i) {
            indexArr[n['name']] = n['value'];
        });
        return JSON.stringify(indexArr);
    }
</script>