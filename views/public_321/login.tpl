<div class="hidden">{{create_captcha}}</div>

<div class="container custom-layout-container">
<div class="row custom-layout-row">
<div class="col-xs-offset-1 col-xs-10 col-sm-offset-3 col-sm-6 col-md-offset-4 col-md-4 custom-layout-form">
  <div class="widget-container-col no-padding">
    <div class="widget-box transparent no-padding">
      <div class="widget-body no-padding">
        <div class="widget-main no-padding">
          <div class="tab-content no-padding">
            <div id="box-login" class="tab-pane in active no-padding">
              <div class="col-xs-12">
                <h2 class="text-success"><b>{{i18n .Lang "site_name"}} - {{i18n .Lang "login"}} {{i18n .Lang "page"}}</b></h2>
                <hr/>
              </div>
              <div class="col-xs-12 no-padding">
                <form id="form-login" action="{{urlfor "LoginController.POST"}}" method="POST">
                    <div class="form-group col-xs-offset-1 col-xs-10">
                      <span class="input-icon width-100">
                        <input type="text" name="username" class="width-100" maxlength="32" placeholder="{{i18n .Lang "username"}}" required minlength="4" />
                        <i class="ace-icon fa fa-user green"></i>
                      </span>
                    </div>
                    <div class="form-group col-xs-offset-1 col-xs-10">
                      <span class="input-icon width-100">
                        <input type="password" name="password" class="width-100" maxlength="32" placeholder="{{i18n .Lang "password"}}" required minlength="4" />
                        <i class="ace-icon fa fa-key red"></i>
                      </span>
                    </div>
                    <div class="form-group col-xs-offset-1 col-xs-10">
                        <div class="col-xs-6 no-padding">
                          <input type="hidden" id="id-login" name="captcha_id" value="">
                          <span class="input-icon">
                            <input type="number" name="captcha" class="width-100" maxlength="4" placeholder="{{i18n .Lang "captcha"}}" required minlength="4" />
                            <i class="ace-icon fa fa-shield blue"></i>
                          </span>
                        </div>
                        <div class="col-xs-6 no-padding">
                          <a class="captcha" href="javascript:">
                            <img class="captcha-img" id="captcha-login" onclick="refresh_captcha(true)" />
                          </a>
                        </div>
                    </div>
                    <div class="col-xs-12">
                      <button class="btn btn-success" type="submit">{{i18n .Lang "login"}}<i class="ace-icon fa fa-arrow-right icon-on-right bigger-120"></i></button>
                      &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                      <button class="btn btn-info" type="reset">{{i18n .Lang "reset"}}<i class="ace-icon fa fa-undo icon-on-right bigger-120"></i></button>
                    </div>
                </form>
              </div>
              <div class="col-xs-12 no-padding">
                <hr/>
                <div class="col-xs-6 text-left">
                  <i class="ace-icon fa fa-qq bigger-160"></i>
                </div>
                <div class="col-xs-6 text-right">
                  <a data-toggle="tab" href="#box-forget">{{i18n .Lang "forget_pwd"}}</a>&nbsp;&nbsp;&nbsp;&nbsp;
                  <a data-toggle="tab" href="#box-register">{{i18n .Lang "register"}}</a>
                </div>
              </div>
            </div>

            <div id="box-register" class="tab-pane no-padding">
              <div class="col-xs-12">
                <h2 class="text-success"><b>{{i18n .Lang "site_name"}} - {{i18n .Lang "register"}} {{i18n .Lang "page"}}</b></h2>
                <hr/>
              </div>
              <div class="col-xs-12 no-padding">
                <form id="form-register" action="{{urlfor "RegisterController.POST"}}" method="POST">
                  <fieldset>
                    <div class="form-group col-xs-offset-1 col-xs-10">
                      <span class="input-icon width-100">
                        <input type="text" name="username" class="width-100" maxlength="32" placeholder="{{i18n .Lang "username"}}" required minlength="4" />
                        <i class="ace-icon fa fa-user green"></i>
                      </span>
                    </div>
                    <div class="form-group col-xs-offset-1 col-xs-10">
                      <span class="input-icon width-100">
                        <input type="password" name="password" class="width-100" maxlength="32" placeholder="{{i18n .Lang "password"}}" required minlength="4" />
                        <i class="ace-icon fa fa-key red"></i>
                      </span>
                    </div>
                    <div class="form-group col-xs-offset-1 col-xs-10">
                      <span class="input-icon width-100">
                        <input type="password" name="pwd_confirm" class="width-100" maxlength="32" placeholder="{{i18n .Lang "pwd_confirm"}}" required equalTo="#register_password" />
                        <i class="ace-icon fa fa-key red"></i>
                      </span>
                    </div>
                    <div class="form-group col-xs-offset-1 col-xs-10">
                      <span class="input-icon width-100">
                        <input type="text" name="email" class="width-100" maxlength="128" placeholder="{{i18n .Lang "email"}}" required email="email" />
                        <i class="ace-icon fa fa-envelope green"></i>
                      </span>
                    </div>
                    <div class="form-group col-xs-offset-1 col-xs-10">
                      <div class="col-xs-6 no-padding">
                          <input type="hidden" id="id-register" name="captcha_id" value="">
                          <span class="input-icon">
                            <input type="number" name="captcha" class="width-100" maxlength="4" placeholder="{{i18n .Lang "captcha"}}" required minlength="4" />
                            <i class="ace-icon fa fa-shield blue"></i>
                          </span>
                        </div>
                        <div class="col-xs-6 no-padding">
                          <a class="captcha" href="javascript:">
                            <img class="captcha-img" id="captcha-register" onclick="refresh_captcha(true)" />
                          </a>
                        </div>
                    </div>
                    <div class="col-xs-12">
                      <button class="btn btn-success" type="submit">{{i18n .Lang "register"}}<i class="ace-icon fa fa-arrow-right icon-on-right bigger-120"></i></button>
                      &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                      <button class="btn btn-info" type="reset">{{i18n .Lang "reset"}}<i class="ace-icon fa fa-undo icon-on-right bigger-120"></i></button>
                    </div>
                  </fieldset>
                </form>
              </div>
              <div class="col-xs-12 no-padding text-right">
                <hr/>
                <a data-toggle="tab" href="#box-login">{{i18n .Lang "back_to"}}{{i18n .Lang "login"}}</a>
              </div>
            </div>

            <div id="box-forget" class="tab-pane no-padding">
              <div class="col-xs-12">
                <h2 class="text-success"><b>{{i18n .Lang "site_name"}} - {{i18n .Lang "find_pwd"}} {{i18n .Lang "page"}}</b></h2>
                <hr/>
              </div>
              <div class="col-xs-12 no-padding">
                <form id="form-forget" action="{{urlfor "FindPwdController.POST"}}" method="POST">
                  <fieldset>
                    <div class="form-group col-xs-offset-1 col-xs-10">
                      <span class="input-icon width-100">
                        <input type="text" name="email" class="width-100" maxlength="128" placeholder="{{i18n .Lang "email"}}" required email="email" />
                        <i class="ace-icon fa fa-envelope green"></i>
                      </span>
                    </div>
                    <div class="form-group col-xs-offset-1 col-xs-10">
                      <div class="col-xs-6 no-padding">
                          <input type="hidden" id="id-forget" name="captcha_id" value="">
                          <span class="input-icon">
                            <input type="number" name="captcha" class="width-100" maxlength="4" placeholder="{{i18n .Lang "captcha"}}" required minlength="4" />
                            <i class="ace-icon fa fa-shield blue"></i>
                          </span>
                        </div>
                        <div class="col-xs-6 no-padding">
                          <a class="captcha" href="javascript:">
                            <img class="captcha-img" id="captcha-forget" onclick="refresh_captcha(true)" />
                          </a>
                        </div>
                    </div>
                    <div class="col-xs-12">
                      <button class="btn btn-success" type="submit">{{i18n .Lang "send"}}<i class="ace-icon fa fa-arrow-right icon-on-right bigger-120"></i></button>
                    </div>
                  </fieldset>
                </form>
              </div>
              <div class="col-xs-12 no-padding text-right">
                <hr/>
                <a data-toggle="tab" href="#box-login">{{i18n .Lang "back_to"}}{{i18n .Lang "login"}}</a>
              </div>  
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

</div>
</div>
</div>

<!-- 模态框（Modal） -->
<div class="modal fade" id="msg-dialog" tabindex="-1" role="dialog" aria-labelledby="msg-title" aria-hidden="true"><br/><br/><br/><br/><br/><br/>
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header text-info bg-primary">
        <button type="button" class="close" data-dismiss="modal" aria-hidden="true">
          &times;
        </button>
        <h4 class="modal-title" id="msg-title">标题</h4>
      </div>
      <div class="modal-body lead text-center" id="msg-content">提示内容</div><br/>
    </div><!-- /.modal-content -->
  </div><!-- /.modal -->
</div>