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
                  <i class="ace-icon fa fa-leaf green"></i>
                  {{i18n .Lang "login"}} {{i18n .Lang "page"}}
                </h4>

                <div class="space-6"></div>

                <form id="form-login" action="{{urlfor "LoginApi.Post"}}" method="POST">
                  <fieldset>
                    <label class="block clearfix">
                      <span class="block input-icon">
                        <i class="ace-icon fa fa-user"></i>
                        <input type="text" name="username" class="form-control" maxlength="32" placeholder="{{i18n .Lang "username"}}" required minlength="4" />
                      </span>
                    </label>

                    <label class="block clearfix">
                      <span class="block input-icon">
                        <i class="ace-icon fa fa-lock"></i>
                        <input type="password" name="password" class="form-control" maxlength="32" placeholder="{{i18n .Lang "password"}}" required minlength="4" />
                      </span>
                    </label>

                    <label class="block clearfix">
                      <div class="col-xs-6 no-padding">
                        <input type="hidden" id="id-login" name="captcha_id" value="">
                        <span class="input-icon">
                          <i class="ace-icon fa fa-shield"></i>
                          <input type="text" name="captcha" class="form-control" maxlength="4" placeholder="{{i18n .Lang "captcha"}}" required minlength="4" />
                        </span>
                      </div>
                      <div class="col-xs-6 no-padding">
                        <a class="captcha" href="javascript:void(0);">
                          <img class="captcha-img" id="captcha-login" onclick="refresh_captcha(true)" />
                        </a>
                      </div>
                    </label>

                    <div class="space"></div>

                    <div class="clearfix center">
                      <button type="sumbit" class="btn btn-sm btn-primary width-65">
                        <i class="ace-icon fa fa-key"></i>
                        <span class="bigger-110">{{i18n .Lang "login"}}</span>
                      </button>
                    </div>

                    <div class="space-4"></div>
                  </fieldset>
                </form>

                <div class="social-or-login center">
                  <span class="bigger-110"></span>
                </div>
                <div class="social-login center">
                  <a class="btn btn-success">
                    <i class="ace-icon fa fa-weixin"></i>
                  </a>
                  <a class="btn btn-primary">
                    <i class="ace-icon fa fa-qq"></i>
                  </a>
                </div>
              </div><!-- /.widget-main -->

              <div class="toolbar clearfix">
                <div>
                  <a href="#" data-target="#forgot-box" class="forgot-password-link">
                    <i class="ace-icon fa fa-arrow-left"></i>
                    {{i18n .Lang "forget_pwd"}}
                  </a>
                </div>

                <div>
                  <a href="#" data-target="#signup-box" class="user-signup-link">
                    {{i18n .Lang "register"}}
                    <i class="ace-icon fa fa-arrow-right"></i>
                  </a>
                </div>
              </div>
            </div><!-- /.widget-body -->
          </div><!-- /.login-box -->

          <div id="forgot-box" class="forgot-box widget-box no-border">
            <div class="widget-body">
              <div class="widget-main">
                <h4 class="header red lighter bigger center">
                  <i class="ace-icon fa fa-key"></i>
                  {{i18n .Lang "forget_pwd"}} {{i18n .Lang "page"}}
                </h4>

                <div class="space-6"></div>
                <p>
                  {{i18n .Lang "via"}}{{i18n .Lang "email"}}{{i18n .Lang "find_pwd"}}:
                </p>

                <form id="form-forget" action="{{urlfor "FindPwdApi.Post"}}" method="POST">
                  <fieldset>
                    <label class="block clearfix">
                      <span class="block input-icon">
                        <i class="ace-icon fa fa-envelope red"></i>
                        <input type="email" name="email" class="form-control" maxlength="128" placeholder="{{i18n .Lang "email"}}" required email />
                      </span>
                    </label>

                    <label class="block clearfix">
                      <div class="col-xs-6 no-padding">
                        <input type="hidden" id="id-forget" name="captcha_id" value="">
                        <span class="input-icon">
                          <i class="ace-icon fa fa-shield red"></i>
                          <input type="text" name="captcha" class="form-control" maxlength="4" placeholder="{{i18n .Lang "captcha"}}" required minlength="4" />
                        </span>
                      </div>
                      <div class="col-xs-6 no-padding">
                        <a class="captcha" href="javascript:void(0);">
                          <img class="captcha-img" id="captcha-forget" onclick="refresh_captcha(true)" />
                        </a>
                      </div>
                    </label>

                    <label class="block clearfix center">
                      <button type="sumbit" class="btn btn-sm btn-danger width-65">
                          <i class="ace-icon fa fa-lightbulb-o"></i>
                          <span class="bigger-110">{{i18n .Lang "send"}}</span>
                        </button>
                    </label>
                  </fieldset>
                </form>
              </div><!-- /.widget-main -->

              <div class="toolbar center">
                <a href="#" data-target="#login-box" class="back-to-login-link">
                  {{i18n .Lang "back_to"}}{{i18n .Lang "login"}}
                  <i class="ace-icon fa fa-arrow-right"></i>
                </a>
              </div>
            </div><!-- /.widget-body -->
          </div><!-- /.forgot-box -->

          <div id="signup-box" class="signup-box widget-box no-border">
            <div class="widget-body">
              <div class="widget-main">
                <h4 class="header green lighter bigger center">
                  <i class="ace-icon fa fa-users blue"></i>
                  {{i18n .Lang "register"}} {{i18n .Lang "page"}}
                </h4>

                <div class="space-6"></div>
                <p> {{i18n .Lang "via"}}{{i18n .Lang "email"}}{{i18n .Lang "register"}}{{i18n .Lang "user"}}: </p>

                <form id="form-register" action="{{urlfor "RegisterApi.Post"}}" method="POST" target="nm_iframe">
                  <fieldset>
                    <label class="block clearfix">
                      <span class="block input-icon">
                        <i class="ace-icon fa fa-envelope green"></i>
                        <input type="email" name="email" class="form-control" maxlength="128" placeholder="{{i18n .Lang "email"}}" required email />
                      </span>
                    </label>

                    <label class="block clearfix">
                      <span class="block input-icon">
                        <i class="ace-icon fa fa-user green"></i>
                        <input type="text" name="username" class="form-control" maxlength="32" placeholder="{{i18n .Lang "username"}}" required />
                      </span>
                    </label>

                    <label class="block clearfix">
                      <span class="block input-icon">
                        <i class="ace-icon fa fa-lock green"></i>
                        <input type="password" name="password" id="password" class="form-control" maxlength="32" placeholder="{{i18n .Lang "password"}}" required minlength="4" />
                      </span>
                    </label>

                    <label class="block clearfix">
                      <span class="block input-icon">
                        <i class="ace-icon fa fa-exchange green"></i>
                        <input type="password" name="pwd_confirm" class="form-control" maxlength="32" placeholder="{{i18n .Lang "pwd_confirm"}}" required minlength="4" equalTo="#password" />
                      </span>
                    </label>

                    <label class="block clearfix">
                      <div class="col-xs-6 no-padding">
                        <input type="hidden" id="id-register" name="captcha_id" value="">
                        <span class="input-icon">
                          <i class="ace-icon fa fa-shield green"></i>
                          <input type="text" name="captcha" class="form-control" maxlength="4" placeholder="{{i18n .Lang "captcha"}}" required minlength="4" />
                        </span>
                      </div>
                      <div class="col-xs-6 no-padding">
                        <a class="captcha" href="javascript:void(0);">
                          <img class="captcha-img" id="captcha-register" onclick="refresh_captcha(true)" />
                        </a>
                      </div>
                    </label>

                    <label class="block clearfix center">
                      <div class="col-xs-4">
                        <button type="reset" class="width-100 pull-left btn btn-sm">
                          <i class="ace-icon fa fa-refresh"></i>
                          <span class="bigger-110">{{i18n .Lang "reset"}}</span>
                        </button>
                      </div>
                      <div class="col-xs-8">
                        <button type="submit" class="width-100 pull-right btn btn-sm btn-success">
                          <span class="bigger-110">{{i18n .Lang "register"}}</span>
                          <i class="ace-icon fa fa-arrow-right icon-on-right"></i>
                        </button>
                      </div>
                    </label>
                  </fieldset>
                </form>
                <iframe id="id_iframe" name="nm_iframe" style="display:none;"></iframe>
              </div>

              <div class="toolbar center">
                <a href="#" data-target="#login-box" class="back-to-login-link">
                  <i class="ace-icon fa fa-arrow-left"></i>
                  {{i18n .Lang "back_to"}}{{i18n .Lang "login"}}
                </a>
              </div>
            </div><!-- /.widget-body -->
          </div><!-- /.signup-box -->
        </div><!-- /.position-relative -->

      </div><!-- /.login-container -->
    </div><!-- /.col -->
  </div><!-- /.row -->
</div><!-- /.main-content -->