// RSA公钥加密
function rsa_encrypt(info) {
    var method = new JSEncrypt();
    method.setPublicKey('-----BEGIN PUBLIC KEY-----MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC02RFpif6PfAWPGz0lb5ZMEDikUfBy0uUvz2WyT4IL/DSvHOmsBweAmapnD+9z9L9zc8e79PRPHYVDn6XmTlTQi7+B+hluQe9juddFCJzCOyblt5lAYtcxokb/TnZYDkeAb8V+Yn8k3WGeV4E2QzhF6M1+bk35rSkwGXSJ//ZXcwIDAQAB-----END PUBLIC KEY-----');
    return method.encrypt(info);
}

// 转换密码后提交
function convpwd_submit(formOld, callback) {
    // 创建formNew
    var formNew = $('<form></form>', { style: "display:none;" }).appendTo('body');
    // 设置属性
    formNew.attr('action', $(formOld).attr('action'));
    formNew.attr('method', $(formOld).attr('method'));
    formNew.attr('target', $(formOld).attr('target'));
    // formNew.hide();

    // 添加控件到formNew
    var data = $(formOld).serializeArray();
    $.each(data, function(i, field) {
        var ctrl = $('<input type="hidden" />');
        ctrl.attr('name', field.name);
		if (field.name == 'password' || field.name == 'pwd_confirm') {
			ctrl.attr('value', rsa_encrypt(field.value));
		} else {
			ctrl.attr('value', field.value);
		}
        formNew.append(ctrl);
    } );

    // 提交
    submit_ajax(formNew, callback);
    return false;
}
