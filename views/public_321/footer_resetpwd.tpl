<script src="/static/js/jquery.validate.msg.{{.Lang}}.js"></script>

<script type="text/javascript">
    $.validator.setDefaults({
        submitHandler: submit_ajax,
    });

	$().ready(function() {
		refresh_captcha(false);

		$("#form-login").validate();
		// $("#form-register").validate();
		// $("#form-forget").validate();
	});

	function refresh_captcha(random) {
		var obj = $("input[name='captcha_id']");
		var value = obj.val();
		if (random) {
			value = '/captcha/' + value + '.png?reload=' + (new Date()).getTime();
		} else {
			value = '/captcha/' + value + '.png';
		}
		$('#captcha-login').attr('src', value);
		$('#captcha-register').attr('src', value);
		$('#captcha-forget').attr('src', value);
	}

    function submit_ajax(form) {
        var obj = $(form);
        $.ajax({
            async: true,
            url: obj.attr('action'),
            type: obj.attr('method'),
            dataType: 'json',
            data: obj.serialize(),
            contentType: 'application/x-www-form-urlencoded; charset=UTF-8',
            beforeSend: function (xhr) {
            },
            complete: function (xhr, ts) {
            },
            success: function (data) {
                if (data.code) {
                    $('#msg-title').html('{{i18n .Lang "login"}}{{i18n .Lang "success"}}');
                    $('#msg-content').html(data.error);
                    $('#msg-dialog').modal('show');
                    // window.location.href = data.url;
                } else {
                    $('#msg-title').html('{{i18n .Lang "login"}}{{i18n .Lang "failed"}}');
                    $('#msg-content').html(data.error);
                    $('#msg-dialog').modal('show');
                }
            },
            error: function (xhr, msg, e) {
                console.log(xhr.status);          // 状态码
                console.log(xhr.readyState);  // 状态
                console.log(msg);                   // 错误信息
                alert("" + xhr.status + msg);
            }
        });
    }
</script>