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

        // verify data
        refresh_captcha(false);
        $.validator.setDefaults({
            submitHandler: function(form) {
                convpwd_submit(form, submit_callback);
            },
            errorPlacement : function (error, element) {
                element.tooltip("destroy");    //清除以前的气泡
                element.tooltip({      //重新生成气泡
                    title: error.text(),
                    placement : 'top',
                })
                var par = $(element).parent();       //获取父元素，为了加上红色边框使错误更加醒目
                par.addClass("has-error");
                var info = $('<span class="glyphicon glyphicon-remove form-control-feedback" aria-hidden="true"></span>');//就是文本框右边那个红色的X
                if(par.find("span").length != 0)//清除之前红色的X
                {
                par.children("span").remove();
                }
                info.appendTo(par);//将红色的X追加到文本框上
            },
            success : function (str,element) {
                // element.tooltip("destroy");          // 清除气泡
                var par = $(element).parent();     // 获取元素父元素便于使文本框有红变绿
                par.removeClass("has-error");   // 恢复文本框边框颜色
                par.addClass("has-success");    // 使文本框边框颜色变成绿色
                if (par.find("span").length != 0) {  // 清除错误时文本框里面的红色X
                    par.children("span").remove();
                }
                var info = $('<span class="glyphicon glyphicon-ok form-control-feedback" aria-hidden="true"></span>');//就是一个勾
                info.appendTo(par);
            }
        });
        $("#form-login").validate();
        $("#form-forget").validate();
        $("#form-register").validate();

        // $("#form-login").submit( function(){
        //     // submit_ajax($("#form-login"), submit_callback);
        //     convpwd_submit($("#form-login"), submit_callback);
        //     return false;
        // });
        // $("#form-forget").submit( function(){
        //     submit_ajax($("#form-forget"), submit_callback);
        //     return false;
        // });
        // $("#form-register").submit( function(){
        //     // submit_ajax($("#form-register"), submit_callback);
        //     convpwd_submit($("#form-register"), submit_callback);
        //     return false;
        // });
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

    function submit_callback(data){
        if (data.code) {
            $('#msg-title').html('{{i18n .Lang "login"}}{{i18n .Lang "success"}}');
            $('#msg-content').html(data.msg);
            $('#msg-dialog').modal('show');
            window.location.href = data.url;
        } else {
            $('#msg-title').html('{{i18n .Lang "login"}}{{i18n .Lang "failed"}}');
            $('#msg-content').html(data.error);
            $('#msg-dialog').modal('show');
            refresh_captcha(true);
        }
        return false;
    }
</script>