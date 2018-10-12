function submit_ajax(form, callback) {
	submit_custom($(form).attr('action'), $(form).attr('method'), $(form).serialize(), callback)
}
function submit_custom(urlPath, ajaxType, formData, callback) {
	$.ajax({
		async: true,
		url: urlPath,
		type: ajaxType,
		dataType: 'json',
		contentType: 'application/x-www-form-urlencoded',
		data: formData,
		beforeSend: function (xhr) {
		},
		complete: function (xhr, ts) {
		},
		success: function (data) {
			var info = {};
			if (typeof(data) == "object") {
				info = data;
			} else {
				info.code = false;
				info.error = 'Rensponse data type error!!!';
			}
			callback(info);
		},
		error: function (xhr, msg, e) {
			console.log(xhr.status);      // 状态码
			console.log(xhr.readyState);   // 状态
			console.log(msg);             // 错误信息
			var info = {};
			info.code = false;
			info.error = '' + xhr.status + msg;
			callback(info);
		}
	});
}
