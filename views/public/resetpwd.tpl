<div class="hidden">{{create_captcha}}</div>

<div class="main-content">
  <div class="row">
    <div class="col-sm-10 col-sm-offset-1">
      <div class="login-container">
        <div class="center">
          <h1>
            <img src="/static/img/logo+title.png" width="120" height="45" />
            <span class="red">{{i18n .Lang "site_name"}}</span>
          </h1>
          <h4 class="blue" id="id-company-text">&copy; {{i18n .Lang "site_domain"}}</h4>
        </div>

        <div class="space-6"></div>

        <div class="position-relative">
          <div id="login-box" class="login-box visible widget-box no-border">
            <div class="widget-body">
              <div class="widget-main">
                <h4 class="header blue lighter bigger center">
                  <i class="ace-icon fa fa-key green"></i>
                  {{i18n .Lang "reset_pwd"}} {{i18n .Lang "page"}}
                </h4>

                <div class="space-6"></div>

                <form id="form-resetpwd" action="{{urlfor "ResetPwdController.POST"}}" method="POST">
                  <fieldset>
                    <label class="block clearfix">
                      <span class="block input-icon">
                        <i class="ace-icon fa fa-user red"></i>
                        <input type="text" name="username" class="form-control" maxlength="32" placeholder="{{i18n .Lang "username"}}" required minlength="4" />
                      </span>
                    </label>

                    <label class="block clearfix">
                      <span class="block input-icon">
                        <i class="ace-icon fa fa-lock red"></i>
                        <input type="password" name="password" class="form-control" maxlength="32" placeholder="{{i18n .Lang "password"}}" required minlength="4" />
                      </span>
                    </label>

                    <label class="block clearfix">
                      <span class="block input-icon">
                        <i class="ace-icon fa fa-retweet red"></i>
                        <input type="password" name="pwd_confirm" class="form-control" maxlength="32" placeholder="{{i18n .Lang "pwd_confirm"}}" required minlength="4" />
                      </span>
                    </label>

                    <label class="block clearfix">
                      <div class="col-xs-6 no-padding">
                        <input type="hidden" id="id-login" name="captcha_id" value="">
                        <span class="input-icon">
                          <i class="ace-icon fa fa-shield red"></i>
                          <input type="text" name="captcha" class="form-control" maxlength="4" placeholder="{{i18n .Lang "captcha"}}" required minlength="4" />
                        </span>
                      </div>
                      <div class="col-xs-6 no-padding">
                        <a class="captcha" href="javascript:">
                          <img class="captcha-img" id="captcha-login" onclick="refresh_captcha(true)" />
                        </a>
                      </div>
                    </label>

                    <div class="space"></div>

                    <div class="clearfix center">
                      <button type="sumbit" class="btn btn-sm btn-primary width-65">
                        <i class="ace-icon fa fa-key"></i>
                        <span class="bigger-110">{{i18n .Lang "ok"}}</span>
                      </button>
                    </div>

                    <div class="space-4"></div>
                  </fieldset>
                </form>
              </div><!-- /.widget-main -->
            </div><!-- /.widget-body -->
          </div><!-- /.login-box -->          
        </div><!-- /.position-relative -->

      </div><!-- /.login-container -->
    </div><!-- /.col -->
  </div><!-- /.row -->
</div><!-- /.main-content -->



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
      <div class="modal-body lead text-center red" id="msg-content">提示内容</div><br/>
    </div><!-- /.modal-content -->
  </div><!-- /.modal -->
</div>